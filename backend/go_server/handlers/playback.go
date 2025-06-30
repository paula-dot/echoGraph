package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	// Update the import path below to the correct location of your models package.
	// For example, if your module name is "github.com/yourusername/echoGraph", use:
	"github.com/Degrante77/echoGraph/backend/go_server/models"
	"github.com/zmb3/spotify"
)

func StartPolling(client *spotify.Client, db *sql.DB) {
	ticker := time.NewTicker(30 * time.Second) // Poll every 30 seconds

	go func() {
		for range ticker.C {
			current, err := client.PlayerCurrentlyPlaying(context.Background())
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
