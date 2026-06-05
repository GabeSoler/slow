package track

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gabesoler/slow/data"
)

type CurrentTrack struct {
	AppName       string
	WindowName    string
	currentStart  time.Time
	SecondCounter int
}

func TrackLoop(ctx context.Context, wg *sync.WaitGroup, track *CurrentTrack, db *data.DBModule) {
	defer wg.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	log.Println("Starting tracking loop...")

	for {
		select {
		case <-ctx.Done():
			if track.AppName != "" {
				currentStart := track.currentStart
				duration := time.Since(currentStart)
				db.RecordUsage(track.AppName, track.WindowName, duration, time.Now())
				log.Printf("[Shutdown] Final save for %s. Duration: %v", track.AppName, duration)
			}
			log.Println("Tracking goroutine stopped cleanly.")
			return

		case <-ticker.C:
			app, window, err := TrackWindow()
			if err != nil {
				log.Printf("Error tracking window: %v", err)
			}

			if track.AppName != app || track.WindowName != window {

				if track.AppName != "" {
					duration := time.Since(track.currentStart)
					db.RecordUsage(track.AppName,
						track.WindowName,
						duration,
						time.Now())
					log.Printf("Switched app! Recorded %s for %v (Loop ticks: %d)",
						track.AppName, duration, track.SecondCounter)
				}

				track.AppName = app
				track.WindowName = window
				track.currentStart = time.Now()
				track.SecondCounter++
			} else {
				track.SecondCounter++
			}

			log.Printf("App: %s | Active Seconds: %d", track.AppName, track.SecondCounter)
		}
	}
}
