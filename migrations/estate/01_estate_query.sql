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
id, urlStr, addressStr, surface, room_amount, price_per_m2, price, query
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: DeleteEstate :exec
DELETE FROM estates
WHERE id = $1;