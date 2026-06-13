package cmd

import (
	"fmt"
	"log"

	"github.com/gabesoler/slow/data"
	"github.com/spf13/cobra"
)

var (
	days int
	md   bool
)

// useCmd represents the use command
var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Observe your app usage",
	Long:  `After tracking your work, have an idea of what you did through the day or week.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage called")
		data.DisplayUsage(days, md)

		log.Println("Display done")
	},
}

func init() {
	//	rootCmd.AddCommand(usageCmd)

	usageCmd.Flags().IntVarP(&days, "back", "b", 0, "Days back to display Usage")
	usageCmd.Flags().BoolVarP(&md, "markdown", "m", false, "Render markdown")
}
