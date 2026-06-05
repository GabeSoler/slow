package main

import (
	"fmt"
	"log"

	"github.com/gabesoler/slow/dim"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gabesoler/slow/data"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// 1. Initialize the Fyne Application
	myApp := app.New()
	myWindow := myApp.NewWindow("App Usage Tracker")
	myWindow.Resize(fyne.NewSize(700, 450))

	// 1.1 Initialize data module (this will be used to fetch real tracking data later)

	// 1. Initialize GORM DB connection
	db, err := gorm.Open(sqlite.Open("apps.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Optional: logs SQL queries
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 2. Initialize your db.AppUse module
	appMod, err := data.NewDBModule(db)
	if err != nil {
		log.Fatalf("failed to initialize module: %v", err)
	}

	// 4. Retrieve and print the data
	records, err := appMod.GetTodayUsage()
	if err != nil {
		log.Fatalf("failed to fetch records: %v", err)
	}

	var data [][]string
	for _, app := range records {
		arr := []string{app.AppName, app.WindowName, app.Duration.String()}
		data = append(data, arr)
	}

	if len(data) == 0 {
		for range []int{0, 1, 2} {
			arr := []string{"App Name", "Window Title", "Duration"}
			data = append(data, arr)
		}
	}

	// 3. Build a stylized table to display the metrics
	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide text placeholder")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		},
	)

	// Set column widths so it looks organized
	table.SetColumnWidth(0, 200)
	table.SetColumnWidth(1, 120)
	table.SetColumnWidth(2, 150)

	// 4. Create Sidebar Controls
	statusLabel := widget.NewLabel("Status: Tracking Active")

	toggleBtn := widget.NewButton("Pause Tracking", func() {
		if statusLabel.Text == "Status: Tracking Active" {
			statusLabel.SetText("Status: Paused")
		} else {
			statusLabel.SetText("Status: Tracking Active")
		}
	})
	toggleBtn.Importance = widget.HighImportance

	// Sidebar container layout
	sidebar := container.NewVBox(
		widget.NewLabelWithStyle("TRACKER", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		statusLabel,
		widget.NewSeparator(),
		toggleBtn,
		widget.NewButton("Force Screen Dim", func() {
			// This button will eventually trigger your low-level click-through overlay!
			fmt.Println("Dimmer button clicked! This will trigger the overlay.")
			go dim.Dim()
		}),
	)

	// 5. Combine Sidebar and Table into a Split Layout
	splitLayout := container.NewHSplit(sidebar, table)
	splitLayout.Offset = 0.25 // Sidebar takes up 25% of the screen width

	// Set content and run the Fyne loop
	myWindow.SetContent(splitLayout)
	myWindow.ShowAndRun()
}
