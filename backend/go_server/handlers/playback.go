package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/paula-dot/echoGraph/backend/go_server/models"
	"github.com/zmb3/spotify"
)

func StartPolling(client *spotify.Client, db *sql.DB) {
	ticker := time.NewTicker(30 * time.Second) // Poll every 30 seconds

	go func() {
		for range ticker.C {
			current, err := client.PlayerCurrentlyPlaying()
			if err != nil {
				log.Printf("Error getting current playback: %v", err)
				continue
			}

			if current == nil || !current.Playing || current.Item == nil {
				continue
			}

			playedAt := time.Now().UTC()

			fmt.Println("Now playing:", current.Item.Name, "-", current.Item.Artists[0].Name)
			err = models.SaveTrack(db, current.Item, playedAt)
			if err != nil {
				log.Println("Failed to save track:", err)
			}
		}
	}()
}
