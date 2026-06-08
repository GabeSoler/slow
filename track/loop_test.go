package track

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/gabesoler/slow/data"
)

// TestRunDimmerSolo lets you execute the GLFW overlay entirely on its own
func TestRunTrakLoop(t *testing.T) {
	t.Log("Launching TrackWindow...")

	ctx, cancel := context.WithCancel(context.Background())

	data.SetUpDatabase()
	// 2. Use a WaitGroup so main doesn't exit before the goroutine finishes saving
	var wg sync.WaitGroup

	globalTracker := &CurrentTrack{}

	// 3. Spin it off as a goroutine
	wg.Add(1)
	go TrackLoop(ctx, &wg, globalTracker)

	// Let it run for 5 seconds to simulate your app doing work...
	time.Sleep(5 * time.Second)

	// 4. Close time! Call the cancel function to tell the goroutine to stop
	t.Log("Calling close function...")
	cancel()

	// 5. Wait for the goroutine to finish saving to the DB and exit
	wg.Wait()
	t.Log("application exited completely")
}
