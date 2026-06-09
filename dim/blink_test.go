package dim

import (
	"testing"
	"time"
)

// TestRunDimmerSolo lets you execute the GLFW overlay entirely on its own
func TestRunBlinkSolo(t *testing.T) {
	t.Log("Launching blink...")
	time.Sleep(2 * time.Second) // Give the user a moment to prepare for the blinking
	Blink()
	time.Sleep(2 * time.Second) // Give the user a moment to recover from the blinking
}
