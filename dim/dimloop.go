package dim

import (
	"context"
	"sync"
	"time"

	"github.com/gabesoler/slow/cfunc"
)

func DimLoop(ctx context.Context, wg *sync.WaitGroup, cycles int, duration int) {
	ticker := time.NewTicker(time.Duration(duration) * time.Minute)
	defer ticker.Stop()
	totalDuration := time.Duration(cycles) * time.Minute
	start := time.Now()
	blinkDuration := time.Duration(200 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			wg.Done()

		case t := <-ticker.C:
			switch timeRef := t.Sub(start); {
			case timeRef >= totalDuration:
				cfunc.Dim(.5)
			case timeRef <= totalDuration/2:
				cfunc.BlinkUp(blinkDuration, .3)
			default:
				cfunc.Blink(blinkDuration, .4)
			}
		}
	}
}
