// Package cfunc manages mac controlls using C and Objective C
package cfunc

/*
#cgo LDFLAGS: -framework CoreGraphics -framework Foundation -framework IOKit -framework ApplicationServices -F /System/Library/PrivateFrameworks -framework DisplayServices
#include <stdlib.h>
#include "window.h"
*/
import "C"

import (
	"unsafe"
)

func GetActiveWindow() (string, string) {
	var cAppName *C.char
	var cWindowTitle *C.char

	// Calls the safe C layer we defined in window.h
	success := C.GetActiveWindowDetails(&cAppName, &cWindowTitle)
	if success == 0 {
		return "", ""
	}

	defer C.free(unsafe.Pointer(cAppName))
	defer C.free(unsafe.Pointer(cWindowTitle))

	return C.GoString(cAppName), C.GoString(cWindowTitle)
}
