package dim

import (
	"testing"
	"time"
)

// TestRunDimmerSolo lets you execute the GLFW overlay entirely on its own
func TestRunDimmerSolo(t *testing.T) {
	t.Log("Launching dim...")
	time.Sleep(2 * time.Second) // Give the user a moment to prepare for the dimming
	Dim()
	time.Sleep(2 * time.Second) // Give the user a moment to recover from the dimming
}
