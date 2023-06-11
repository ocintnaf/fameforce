CREATE TABLE
    IF NOT EXISTS influencers (
        id SERIAL PRIMARY KEY,
        user_id VARCHAR(255) UNIQUE,
        created_at TIMESTAMP NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW (),
        FOREIGN KEY (user_id) REFERENCES users (id)
    );