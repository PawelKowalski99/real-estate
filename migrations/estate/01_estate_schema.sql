-- +goose Up
CREATE TABLE IF NOT EXISTS estates (
id  UUID PRIMARY KEY,
id_estate text NOT NULL,
urlStr text NOT NULL,
addressStr text NOT NULL,
surface float8 NOT NULL,
room_amount text NOT NULL,
price_per_m2 float8 NOT NULL,
price float8 NOT NULL,
query text NOT NULL,
city text NOT NULL,
offer_date DATE NOT NULL DEFAULT CURRENT_DATE,
rent_price float8 NOT NULL
);