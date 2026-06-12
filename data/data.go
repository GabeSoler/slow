// Package data provides the database module for tracking app usage.
// It defines the AppUse model and includes functions for recording and retrieving app usage data.
package data

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// AppUse represents the database model for tracking app usage.
type AppUse struct {
	ID         uint      `gorm:"primaryKey"`
	AppName    string    `gorm:"not null;index"`
	WindowName string    `gorm:"not null;index"`
	Duration   float64   `gorm:"type:bigint;not null"` // Stored as nanoseconds (int64)
	Date       time.Time `gorm:"type:date;not null;index"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

// DBModule handles the database operations for AppUse.
type DBModule struct {
	db *gorm.DB
}

// NewDBModule initializes the module and runs auto-migrations.
func NewDBModule(db *gorm.DB) (*DBModule, error) {
	// AutoMigrate will create the table and indexes if they don't exist
	err := db.AutoMigrate(&AppUse{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate AppUse table: %w", err)
	}

	return &DBModule{db: db}, nil
}

// RecordUsage inserts a new app usage entry into the database.
func (m *DBModule) RecordUsage(appName string, appWindow string, duration float64, date time.Time) error {
	// Truncate the date to remove time components (keeping it strictly a "Date")
	cleanDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	entry := AppUse{
		AppName:    appName,
		WindowName: appWindow,
		Duration:   duration,
		Date:       cleanDate,
	}

	result := m.db.Create(&entry)
	return result.Error
}

// GetUsageByApp retrieves all usage records for a specific app.
func (m *DBModule) GetUsageByApp(appName string) ([]AppUse, error) {
	var records []AppUse
	result := m.db.Where("app_name = ?", appName).Find(&records)
	return records, result.Error
}

func (m *DBModule) GetTodayUsage() ([]AppUse, error) {
	var records []AppUse

	// Get today's date strictly at midnight 00:00:00
	now := time.Now()
	today := time.Date(now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location())

	result := m.db.
		Where("date = ?", today).
		Find(&records)

	return records, result.Error
}

type WindowSummary struct {
	AppName    string
	WindowName string
	TotalUse   float64
	AverageUse float64
	Switches   int
}

func (m *DBModule) GetAggretatedUsage(daysBack int) []WindowSummary {
	var currentRecords []WindowSummary
	dateRef := time.Now().AddDate(0, 0, -daysBack)
	m.db.Select(`
		app_name,
							window_name, 
							SUM(CAST(duration AS DECIMAL(18,2))) as total_use,
							SUM(CAST(duration AS DECIMAL(18,2))) / COUNT(DISTINCT DATE(date)) as average_use,
							COUNT(id) as switches
		`).
		Table("app_uses").
		Where("date >= ?", dateRef.Format("2006-01-02")). // Filter for today's date
		Group("app_name, window_name").
		Order("total_use DESC").
		Scan(&currentRecords)
	return currentRecords
}

func SetUpDatabase() (DBModule, error) {
	// 1. Initialize GORM DB connection
	db, err := gorm.Open(sqlite.Open("slow.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Optional: logs SQL queries
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 2. Initialize your db.AppUse module
	dbModule, err := NewDBModule(db)
	if err != nil {
		log.Fatalf("failed to initialize module: %v", err)
	}
	return *dbModule, err
}
