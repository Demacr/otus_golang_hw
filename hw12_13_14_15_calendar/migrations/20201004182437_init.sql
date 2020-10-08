-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    UUID VARCHAR(36)
);

CREATE TABLE events (
    id            SERIAL PRIMARY KEY,
    UUID          VARCHAR(36),
    header        TEXT,
    dt            TIMESTAMP,
    duration      INTERVAL,
    description   TEXT,
    user_id       INT,
    notify_before INTERVAL,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
);
-- +goose Down
DROP TABLE events;
DROP TABLE users;
