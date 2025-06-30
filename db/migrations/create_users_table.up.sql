CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    spotify_id TEXT UNIQUE NOT NULL,
    display_name TEXT,
    access_token TEXT NO,
    refresh_token TEXT NOT NULL,
    token_expires_at TIMESTAMPT WITH TIME ZONE NOT NULL,
    created_at TIMESTAMPT DEFAULT CURRENT,

);