-- +goose Up
CREATE TABLE repository (
                            id BIGSERIAL PRIMARY KEY,
                            name TEXT NOT NULL,
                            link TEXT NOT NULL
);