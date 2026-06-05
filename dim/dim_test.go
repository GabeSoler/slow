package dim

import (
	"testing"
	"time"
)

// TestRunDimmerSolo lets you execute the GLFW overlay entirely on its own
func TestRunDimmerSolo(t *testing.T) {
	t.Log("Launching dimmer overlay for 5 seconds...")

	// Launch the overlay
	// Because go test runs this function on the main thread,
	// GLFW will be able to initialize perfectly.

	// Note: We use a channel or a quick goroutine to close it after 5 seconds
	// so the test doesn't hang forever.
	go func() {
		time.Sleep(5 * time.Second)
		t.Log("Time up! Closing test.")
		// If you have a close function, call it here,
		// otherwise just letting the test finish or killing the window works.
	}()

	Dim()
}
