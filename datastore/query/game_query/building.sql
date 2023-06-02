-- name: CreateBuilding :one
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
RETURNING *;

-- name: GetBuilding :one
SELECT * FROM public.building
WHERE token_id = $1 LIMIT 1;

-- name: ListBuilding :many
SELECT * FROM public.building
ORDER BY id;