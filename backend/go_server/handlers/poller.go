package handlers

import (
	"database/sql"
	"log"
	"time"

	"echoGraph/backend/go_server/models"

	"github.com/zmb3/spotify"
)

func StartPolling(client spotify.Client, db *sql.DB) {
	for {
		recent, err := client.PlayerRecentlyPlayed()
		if err != nil {
			log.Printf("Error fetching recently played tracks: %v", err)
			time.Sleep(30 * time.Second) // Wait before retrying
			continue
		}

		for _, item := range recent {
			t := item.Track
			track := models.Track{
				Title:           t.Name,
				Artist:          t.Artists[0].Name,
				Album:           t.Album.Name,
				DurationSeconds: int(t.Duration / time.millisecond / 1000),
			}
			if err := models.saveTrack(db, track); err != nil {
				log.Println("Failed to save track:", err)
			}
		}

		time.Sleep(30 * time.Second) // Poll every 30 seconds
	}

}
