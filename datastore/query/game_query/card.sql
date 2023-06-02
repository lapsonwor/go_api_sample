-- name: GetCardByTokenID :one
SELECT *, r.name as rarity FROM public.card c
LEFT JOIN public.rarity r
ON c.rarity_id = r.id
LEFT JOIN public.c_e_file cef
ON c.c_e_file_id = cef.id
WHERE c.token_id = $1 
LIMIT 1;

-- name: ListAllCards :many
SELECT * FROM public.card
ORDER BY id;

-- name: CreateCard :one
INSERT INTO public.card (
  name,
  owner_address,
  equipment_id,
  character_id,
  level,
  token_id,
  rarity_id,
  is_ur,
  locked,
  number_of_skills,
  c_e_file_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: CreateCardFromWorker :one
INSERT INTO public.card (
  name,
  owner_address,
  equipment_id,
  character_id,
  level,
  token_id,
  rarity_id,
  is_ur,
  c_e_file_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateCardFromWorker :one
UPDATE public.card
SET owner_address = $1, updated_at = $2
WHERE token_id = $3
RETURNING *;