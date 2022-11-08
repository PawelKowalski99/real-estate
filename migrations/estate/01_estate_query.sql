-- name: GetEstate :one
SELECT * FROM estates
WHERE id = $1 LIMIT 1;

-- name: GetEstateByUrl :one
SELECT * FROM estates
WHERE urlStr = $1 LIMIT 1;

-- name: ListEstates :many
SELECT * FROM estates
ORDER BY price;

-- name: FindEstates :many
SELECT * FROM estates
WHERE query = $1
ORDER BY price;

-- name: CreateEstate :one
INSERT INTO estates (
id, urlStr, addressStr, surface, room_amount, price_per_m2, price, query, rent_price, id_estate, city
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: GetPrices :many
SELECT city, price FROM estates;

-- name: GetPricesPerM2 :many
SELECT city, price_per_m2 FROM estates;

-- name: DeleteEstate :exec
DELETE FROM estates
WHERE id = $1;