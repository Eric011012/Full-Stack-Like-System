CREATE TABLE IF NOT EXISTS message (
    id SERIAL PRIMARY KEY,
    content text NOT NULL,
    user_id INT NOT NULL, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)

CREATE TABLE IF NOT EXISTS reactions (
    id SERIAL PRIMARY KEY, 
    message_id INT REFERENCES message(id) ON DELETE CASCADE,
    user_id INT NOT NULL,
    emoji TEXT NOT NULL,
    UNIQUE(message_id, user_id, emoji)
)