CREATE TABLE IF NOT EXISTS tracks(
    id TEXT PRIMARY KEY, -- Spotify track ID
    name TEXT NOT NULL,
    artist TEXT NOT NULL,
    album TEXT NOT NULL,
    duration_ms INTEGER NOT NULL
);