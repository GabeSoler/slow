package dim

import (
	"context"
	"sync"
	"time"
)

func DimLoop(ctx context.Context, wg *sync.WaitGroup, cycle int, totalCycles int) {
	ticker := time.NewTicker(time.Duration(cycle) * time.Minute)
	defer ticker.Stop()
	total := time.Duration(totalCycles) * time.Minute
	start := time.Now()
	for {
		select {
		case <-ctx.Done():
			wg.Done()

		case t := <-ticker.C:
			if t.Sub(start) >= total {
				Dim()
			} else {
				Blink()
			}
		}
	}
}
