package cfunc

import (
	"testing"
	"time"
)

func TestSetScreenBrightness(t *testing.T) {
	// 1. Drop brightness to 20%
	t.Log("Testing: Dimming screen to 20%...")
	SetScreenBrightness(0.2)

	// Hold it there for 1 second so you can visually see it happen
	time.Sleep(1 * time.Second)

	// 2. Restore brightness back to 80%
	t.Log("Testing: Restoring screen to 80%...")
	SetScreenBrightness(0.8)

	// Small pause to let the hardware state update
	time.Sleep(200 * time.Millisecond)

	// 3. Blink function
	t.Log("Testing: Blink function..")
	duration := time.Duration(500 * time.Millisecond)
	Blink(duration, 0.3)

	// Small pause to let the hardware state update
	time.Sleep(200 * time.Millisecond)
}
