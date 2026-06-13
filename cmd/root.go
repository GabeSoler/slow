// Package cmd handles all the action of Slow, the CLI tool to slow down your computer.
package cmd

import (
	"context"
	"log"
	"os"
	"sync"

	_ "net/http/pprof"

	"github.com/gabesoler/slow/dim"
	"github.com/gabesoler/slow/track"
	"github.com/spf13/cobra"
)

// Define variables to hold flag values
var (
	duration float32
	cycle    int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "slow",
	Short: "an app to help you slow down",
	Long: `Slow runs a loop cycle that triggers a blink on the brightness.
	The default is 60 minutes and ending at 8 cycles, witha  full dim of the screen.
	it also tracks app usage, so you can see what have you been doing.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Initialize your context, cancel, and WaitGroup

		var wg sync.WaitGroup
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // Ensure resources are cleaned up when Run exits

		// 2. Fetch the flag values (already parsed by Cobra at this point)
		// You can use the bound variables directly, or use cmd.Flags().Get...
		log.Printf("Running with duration: %f mins, cycles: %d\n", duration, cycle)

		// 3. Call your two imported functions
		// Increments the WaitGroup for the goroutines inside your functions
		wg.Add(2)

		go dim.DimLoop(ctx, &wg, cycle, duration*60)
		go track.TrackLoop(ctx, &wg)

		// Wait for the background processes to finish
		wg.Wait()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(usageCmd)
	rootCmd.AddCommand(trackCmd)
	rootCmd.AddCommand(dimCmd)

	// Define your two custom flags and bind them to the variables
	rootCmd.Flags().Float32VarP(&duration, "duration", "d", 8, "Total hours of DimLoop (float32)")
	rootCmd.Flags().IntVarP(&cycle, "every", "e", 60, "Cycle lenght in minutes (int)")
}
