package data

import (
	"fmt"
	"log"
	"os"
	"strconv"

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
		var totalStr string
		var avgStr string

		if item.TotalUse > 60.0 {
			totalMinutes := item.TotalUse / 60
			avgMinutes := item.AverageUse / 60

			totalStr = fmt.Sprintf("%.1f m", totalMinutes)
			avgStr = fmt.Sprintf("%.1f m", avgMinutes)

		} else {

			totalStr = fmt.Sprintf("%.1f s", item.TotalUse)
			avgStr = fmt.Sprintf("%.1f s", item.AverageUse)

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
