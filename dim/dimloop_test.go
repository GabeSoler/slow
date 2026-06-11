package dim

import (
	"context"
	"sync"
	"testing"
	"time"
)

// TestRunDimmerSolo lets you execute the GLFW overlay entirely on its own
func TestRunDimmLoop(t *testing.T) {
	t.Log("Launching dim loop...")

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go DimLoop(ctx, &wg, 3, 1)
	time.Sleep(3 * time.Minute)
	cancel()
	wg.Done()
}
