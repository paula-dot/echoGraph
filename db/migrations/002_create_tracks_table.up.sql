CREATE TABLE tracks(
    id TEXT PRIMARY KEY, -- Spotify track ID
    title TEXT NOT NULL,
    artist TEXT NOT NULL,
    album TEXT,
    duration_seconds INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);