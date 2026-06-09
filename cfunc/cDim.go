package cfunc

/*
#include "brightness.h"
*/
import "C"

func SetScreenBrightness(level float32) {
	C.SetBrightness(C.float(level))
}
