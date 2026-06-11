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
	dimCmd.Flags().IntVarP(&duration, "duration", "d", 60, "Duration of the loop cycle in minutes")
	dimCmd.Flags().IntVarP(&cycles, "cycles", "c", 8, "Number of cycles before ending")
}
