package models

import (
	"database/sql"
	"log"
	"time"

	"github.com/zmb3/spotify"
)

// Track represents a music track.
type Track struct {
	ID       string
	Name     string
	Artist   string
	Album    string
	Duration int
	PlayedAt time.Time
}

// SaveTrack inserts a track play event into the PostgreSQL database.
// Assumes a unique (track ID, played at) constraint to avoid duplicates.
func SaveTrack(db *sql.DB, item *spotify.FullTrack, playedAt time.Time) error {
	const query = `
	INSERT INTO play_events 
		(track_id, name, artist, album, duration_ms, played_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (track_id, played_at) DO NOTHING;
	`
	var artistName string
	if len(item.Artists) > 0 {
		artistName = item.Artists[0].Name
	}

	_, err := db.Exec(
		query,
		item.ID.String(),
		item.Name,
		artistName,
		item.Album.Name,
		item.Duration,
		playedAt.UTC(), // Always store in UTC for consistency
	)
	if err != nil {
		log.Printf("Failed to insert track '%s': %v", item.Name, err)
	}

	return err
}
