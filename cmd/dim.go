package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dimCmd represents the dim command
var dimCmd = &cobra.Command{
	Use:   "dim",
	Short: "Dim launches the dim loop without the tracking system",
	Long:  `Dim allow runs in the backround and dims or lightens the screen depending on the time used.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dim called")
	},
}

func init() {
	// Define your two custom flags and bind them to the variables
	dimCmd.Flags().Float32VarP(&duration, "duration", "d", 8, "Total hours of DimLoop (float32)")
	dimCmd.Flags().IntVarP(&cycle, "every", "e", 60, "Cycle lenght in minutes (int)")
}
