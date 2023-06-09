CREATE TYPE user_type AS ENUM ('influencer', 'sponsor');

CREATE TABLE
    IF NOT EXISTS users (
        id VARCHAR(255) PRIMARY KEY, -- Firebase User ID
        email VARCHAR(255) NOT NULL,
        type user_type NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW (),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
    );