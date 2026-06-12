package track

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gabesoler/slow/cfunc"
	"github.com/gabesoler/slow/data"
)

type CurrentTrack struct {
	AppName      string
	WindowName   string
	currentStart time.Time
}

func TrackLoop(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	db, err := data.SetUpDatabase()
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}
	track := CurrentTrack{
		AppName:      "None",
		WindowName:   "None",
		currentStart: time.Now(),
	}

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	log.Println("Starting tracking loop...")

	for {
		select {
		case <-ctx.Done():
			if track.AppName != "" {
				currentStart := track.currentStart
				duration := time.Since(currentStart).Seconds()
				err := db.RecordUsage(track.AppName, track.WindowName, duration, time.Now())
				if err != nil {
					log.Printf("[Shutdown] Error saving final record for %s: %v", track.AppName, err)
				}
				log.Printf("[Shutdown] Final save for %s. Duration: %v", track.AppName, duration)
			}
			log.Println("Tracking goroutine stopped cleanly.")
			return

		case <-ticker.C:
			app, window := cfunc.GetActiveWindow()

			if track.AppName != app || track.WindowName != window {
				if track.AppName != "" {
					// Storing new window
					duration := time.Since(track.currentStart).Seconds()
					db.RecordUsage(track.AppName,
						track.WindowName,
						duration,
						time.Now())
					log.Printf("Switched app! Recorded %s for %v",
						track.AppName, duration)
				}
				// Reseting track
				track.currentStart = time.Now()
				track.AppName = app
				track.WindowName = window
			}
		}
	}
}
