-- USERS TABLE
`
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    spotify_id TEXT UNIQUE NOT NULL,
    display_name TEXT,
    access_token TEXT,
    refresh_token TEXT,
    token_expires_at TIMESTAMPT WITH TIME ZONE NOT NULL,
    created_at TIMESTAMPT WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMPT WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- TRACKS TABLE
CREATE TABLE IF NOT EXISTS tracks (
    id TEXT PRIMARY KEY, -- Spotify track ID
    name TEXT NOT NULL,
    artist TEXT NOT NULL,
    album TEXT NOT NULL,
    duration_ms INTEGER NOT NULL,
);

-- PLAY EVENTS TABLE
CREATE TABLE IF NOT EXISTS play_events (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    track_id TEXT NOT NULL REFERENCES tracks(id) ON DELETE CASCADE,
    played_at TIMESTAMPT WITH TIME ZONE NOT NULL,
    context TEXT,
    UNIQUE (user_id, track_id, played_at)
);