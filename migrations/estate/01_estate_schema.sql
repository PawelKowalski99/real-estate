-- +goose Up
CREATE TABLE IF NOT EXISTS estates (
id   UUID PRIMARY KEY,
urlStr text NOT NULL,
addressStr text NOT NULL,
surface text NOT NULL,
room_amount text NOT NULL,
price_per_m2 text NOT NULL,
price text NOT NULL,
query text NOT NULL
);