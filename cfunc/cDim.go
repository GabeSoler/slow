package cfunc

/*
#include "brightness.h"
*/
import "C"
import "time"

// SetScreenBrightness adjusts the Mac display brightness (0.0 to 1.0)
func SetScreenBrightness(level float32) {
	C.SetBrightness(C.float(level))
}

// GetScreenBrightness reads the current Mac display brightness (0.0 to 1.0)
func GetScreenBrightness() float32 {
	return float32(C.GetBrightness())
}

// Blink drops the screen by decreace amount and restores it back to its original state
func Blink(duration time.Duration, decreace float32) {
	original := GetScreenBrightness()

	SetScreenBrightness(original - decreace)

	time.Sleep(duration)

	SetScreenBrightness(original)
}

// BlinkUp hightens the screen by decreace amount and restores it back to its original state
func BlinkUp(duration time.Duration, decreace float32) {
	original := GetScreenBrightness()

	SetScreenBrightness(original + decreace)

	time.Sleep(duration)

	SetScreenBrightness(original)
}

func Dim(decreace float32) {
	original := GetScreenBrightness()
	SetScreenBrightness(original - decreace)
}
