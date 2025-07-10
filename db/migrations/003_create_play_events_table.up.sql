CREATE TABLE IF NOT EXISTS play_events (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    track_id TEXT REFERENCES tracks(id) ON DELETE CASCADE,
    played_at TIMESTAMP WITH TIME ZONE NOT NULL,
    context TEXT,
    UNIQUE (user_id, track_id, played_at)
);