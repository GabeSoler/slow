package cfunc

import (
	"testing"
	"time"
)

func TestSetScreenBrightness(t *testing.T) {
	// Blink function integrates the other functions
	t.Log("Testing: Blink function..")
	duration := time.Duration(500 * time.Millisecond)
	Blink(duration, 0.3)

	// Small pause to let the hardware state update
	time.Sleep(900 * time.Millisecond)

	t.Log("Testing BlinkUp function")
	BlinkUp(duration, .3)
	// Small pause to let the hardware state update
	time.Sleep(900 * time.Millisecond)
	// Drop brightness to 20%
	t.Log("Ending with a 60% dim")
	Dim(.6)
	// Small pause to let the hardware state update
	time.Sleep(200 * time.Millisecond)
}
