package cfunc

import (
	"fmt"
	"testing"
)

func TestRunTrakLoop(t *testing.T) {
	app, window := GetActiveWindow()
	fmt.Printf(app, window)
}
