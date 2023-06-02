// Code generated by sqlc. DO NOT EDIT.
// source: building.sql

package sqlc

import (
	"context"
)

const createBuilding = `-- name: CreateBuilding :one
INSERT INTO public.building (
  name,
  level,
  rarity,
  owner,
  isUR,
  token_id,
  is_burned
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, name, level, rarity, owner, isur, token_id, is_burned, created_at, updated_at, deleted_at
`

type CreateBuildingParams struct {
	Name     string `json:"name"`
	Level    int32  `json:"level"`
	Rarity   string `json:"rarity"`
	Owner    string `json:"owner"`
	Isur     bool   `json:"isur"`
	TokenID  int32  `json:"token_id"`
	IsBurned bool   `json:"is_burned"`
}

func (q *Queries) CreateBuilding(ctx context.Context, arg CreateBuildingParams) (Building, error) {
	row := q.db.QueryRowContext(ctx, createBuilding,
		arg.Name,
		arg.Level,
		arg.Rarity,
		arg.Owner,
		arg.Isur,
		arg.TokenID,
		arg.IsBurned,
	)
	var i Building
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Level,
		&i.Rarity,
		&i.Owner,
		&i.Isur,
		&i.TokenID,
		&i.IsBurned,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getBuilding = `-- name: GetBuilding :one
SELECT id, name, level, rarity, owner, isur, token_id, is_burned, created_at, updated_at, deleted_at FROM public.building
WHERE token_id = $1 LIMIT 1
`

func (q *Queries) GetBuilding(ctx context.Context, tokenID int32) (Building, error) {
	row := q.db.QueryRowContext(ctx, getBuilding, tokenID)
	var i Building
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Level,
		&i.Rarity,
		&i.Owner,
		&i.Isur,
		&i.TokenID,
		&i.IsBurned,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listBuilding = `-- name: ListBuilding :many
SELECT id, name, level, rarity, owner, isur, token_id, is_burned, created_at, updated_at, deleted_at FROM public.building
ORDER BY id
`

func (q *Queries) ListBuilding(ctx context.Context) ([]Building, error) {
	rows, err := q.db.QueryContext(ctx, listBuilding)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Building
	for rows.Next() {
		var i Building
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Level,
			&i.Rarity,
			&i.Owner,
			&i.Isur,
			&i.TokenID,
			&i.IsBurned,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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
