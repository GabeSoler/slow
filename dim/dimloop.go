package dim

import (
	"context"
	"sync"
	"time"

	"github.com/gabesoler/slow/cfunc"
)

func DimLoop(ctx context.Context, wg *sync.WaitGroup, cycleMinutes int, durationMinutes float32) {
	ticker := time.NewTicker(time.Duration(cycleMinutes) * time.Minute)
	defer ticker.Stop()
	totalDuration := time.Duration(durationMinutes) * time.Minute
	start := time.Now()
	blinkDuration := time.Duration(300 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			wg.Done()

		case t := <-ticker.C:
			switch timeRef := t.Sub(start); {
			case timeRef >= totalDuration:
				cfunc.Dim(.5)
			case timeRef <= totalDuration/2:
				cfunc.Blink(blinkDuration, .4)
			default:
				cfunc.Blink(blinkDuration, .6)
			}
		}
	}
}
