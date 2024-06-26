-- +goose Up
CREATE TABLE banners(
    id serial primary key,
    title text not null,
    text text not null,
    url varchar(255) not null,
    feature_id INTEGER NOT NULL,
    tag_id INTEGER[] NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    last_version BOOLEAN NOT NULL DEFAULT TRUE,
    active BOOLEAN NOT NULL DEFAULT TRUE
);

-- +goose Down
DROP TABLE banners;
