package cmd

import (
	"context"
	"sync"

	"github.com/gabesoler/slow/track"
	"github.com/spf13/cobra"
)

// trackCmd represents the track command
var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Activates the tracking loop without the blink loop",
	Long:  `Track is checking for the windows you use and recording them in a sqlite local db`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Initialize your context, cancel, and WaitGroup
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel() // Ensure resources are cleaned up when Run exits

		var wg sync.WaitGroup

		// 3. Call your two imported functions
		// Increments the WaitGroup for the goroutines inside your functions
		wg.Add(1)

		go track.TrackLoop(ctx, &wg)

		// 4. Wait for the background processes to finish
		wg.Wait()
	},
}

func init() {
	// rootCmd.AddCommand(trackCmd)
}
