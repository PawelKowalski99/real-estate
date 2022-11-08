// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: 01_estate_query.sql

package estate

import (
	"context"

	"github.com/google/uuid"
)

const createEstate = `-- name: CreateEstate :one
INSERT INTO estates (
id, urlStr, addressStr, surface, room_amount, price_per_m2, price, query, rent_price, id_estate, city
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING id, id_estate, urlstr, addressstr, surface, room_amount, price_per_m2, price, query, city, offer_date, rent_price
`

type CreateEstateParams struct {
	ID         uuid.UUID
	Urlstr     string
	Addressstr string
	Surface    float64
	RoomAmount string
	PricePerM2 float64
	Price      float64
	Query      string
	RentPrice  float64
	IDEstate   string
	City       string
}

func (q *Queries) CreateEstate(ctx context.Context, arg CreateEstateParams) (Estate, error) {
	row := q.db.QueryRowContext(ctx, createEstate,
		arg.ID,
		arg.Urlstr,
		arg.Addressstr,
		arg.Surface,
		arg.RoomAmount,
		arg.PricePerM2,
		arg.Price,
		arg.Query,
		arg.RentPrice,
		arg.IDEstate,
		arg.City,
	)
	var i Estate
	err := row.Scan(
		&i.ID,
		&i.IDEstate,
		&i.Urlstr,
		&i.Addressstr,
		&i.Surface,
		&i.RoomAmount,
		&i.PricePerM2,
		&i.Price,
		&i.Query,
		&i.City,
		&i.OfferDate,
		&i.RentPrice,
	)
	return i, err
}

const deleteEstate = `-- name: DeleteEstate :exec
DELETE FROM estates
WHERE id = $1
`

func (q *Queries) DeleteEstate(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteEstate, id)
	return err
}

const findEstates = `-- name: FindEstates :many
SELECT id, id_estate, urlstr, addressstr, surface, room_amount, price_per_m2, price, query, city, offer_date, rent_price FROM estates
WHERE query = $1
ORDER BY price
`

func (q *Queries) FindEstates(ctx context.Context, query string) ([]Estate, error) {
	rows, err := q.db.QueryContext(ctx, findEstates, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Estate
	for rows.Next() {
		var i Estate
		if err := rows.Scan(
			&i.ID,
			&i.IDEstate,
			&i.Urlstr,
			&i.Addressstr,
			&i.Surface,
			&i.RoomAmount,
			&i.PricePerM2,
			&i.Price,
			&i.Query,
			&i.City,
			&i.OfferDate,
			&i.RentPrice,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEstate = `-- name: GetEstate :one
SELECT id, id_estate, urlstr, addressstr, surface, room_amount, price_per_m2, price, query, city, offer_date, rent_price FROM estates
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEstate(ctx context.Context, id uuid.UUID) (Estate, error) {
	row := q.db.QueryRowContext(ctx, getEstate, id)
	var i Estate
	err := row.Scan(
		&i.ID,
		&i.IDEstate,
		&i.Urlstr,
		&i.Addressstr,
		&i.Surface,
		&i.RoomAmount,
		&i.PricePerM2,
		&i.Price,
		&i.Query,
		&i.City,
		&i.OfferDate,
		&i.RentPrice,
	)
	return i, err
}

const getEstateByUrl = `-- name: GetEstateByUrl :one
SELECT id, id_estate, urlstr, addressstr, surface, room_amount, price_per_m2, price, query, city, offer_date, rent_price FROM estates
WHERE urlStr = $1 LIMIT 1
`

func (q *Queries) GetEstateByUrl(ctx context.Context, urlstr string) (Estate, error) {
	row := q.db.QueryRowContext(ctx, getEstateByUrl, urlstr)
	var i Estate
	err := row.Scan(
		&i.ID,
		&i.IDEstate,
		&i.Urlstr,
		&i.Addressstr,
		&i.Surface,
		&i.RoomAmount,
		&i.PricePerM2,
		&i.Price,
		&i.Query,
		&i.City,
		&i.OfferDate,
		&i.RentPrice,
	)
	return i, err
}

const getPrices = `-- name: GetPrices :many
SELECT city, price FROM estates
`

type GetPricesRow struct {
	City  string
	Price float64
}

func (q *Queries) GetPrices(ctx context.Context) ([]GetPricesRow, error) {
	rows, err := q.db.QueryContext(ctx, getPrices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPricesRow
	for rows.Next() {
		var i GetPricesRow
		if err := rows.Scan(&i.City, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPricesPerM2 = `-- name: GetPricesPerM2 :many
SELECT city, price_per_m2 FROM estates
`

type GetPricesPerM2Row struct {
	City       string
	PricePerM2 float64
}

func (q *Queries) GetPricesPerM2(ctx context.Context) ([]GetPricesPerM2Row, error) {
	rows, err := q.db.QueryContext(ctx, getPricesPerM2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPricesPerM2Row
	for rows.Next() {
		var i GetPricesPerM2Row
		if err := rows.Scan(&i.City, &i.PricePerM2); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listEstates = `-- name: ListEstates :many
SELECT id, id_estate, urlstr, addressstr, surface, room_amount, price_per_m2, price, query, city, offer_date, rent_price FROM estates
ORDER BY price
`

func (q *Queries) ListEstates(ctx context.Context) ([]Estate, error) {
	rows, err := q.db.QueryContext(ctx, listEstates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Estate
	for rows.Next() {
		var i Estate
		if err := rows.Scan(
			&i.ID,
			&i.IDEstate,
			&i.Urlstr,
			&i.Addressstr,
			&i.Surface,
			&i.RoomAmount,
			&i.PricePerM2,
			&i.Price,
			&i.Query,
			&i.City,
			&i.OfferDate,
			&i.RentPrice,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
