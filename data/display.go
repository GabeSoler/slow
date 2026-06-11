package data

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

func DisplayUsage(daysBack int) {
	m, err := SetUpDatabase()
	if err != nil {
		log.Fatalf("error Initiatin Db, %v", err)
	}
	data := m.GetAggretatedUsage(daysBack)

	table := tablewriter.NewWriter(os.Stdout)

	table.Header("App", "Window", "TotalTime", "TimeAv", "Switches")

	maxWindowLen := 20
	for _, item := range data {
		displayWindow := item.WindowName
		if len(displayWindow) > maxWindowLen {
			displayWindow = displayWindow[:maxWindowLen-3] + "..."
		}

		totalMinutes := int(item.TotalUse.Minutes())
		avgMinutes := int(item.AverageUse.Minutes())

		totalStr := strconv.Itoa(totalMinutes) + "m"
		avgStr := strconv.Itoa(avgMinutes) + "m"

		if item.TotalUse < time.Minute {
			totalStr = item.TotalUse.Round(time.Second).String()
		}
		if item.AverageUse < time.Minute {
			avgStr = item.AverageUse.Round(time.Second).String()
		}
		row := []string{
			item.AppName,
			displayWindow,
			totalStr,
			avgStr,
			strconv.Itoa(item.Switches), // Convert the int to a string
		}
		table.Append(row)
	}

	// Render the table to the terminal
	table.Render()
}
