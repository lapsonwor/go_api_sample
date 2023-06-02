-- name: GetCharacterByID :one
SELECT * FROM public.character
WHERE id = $1 LIMIT 1;

-- name: ListAllCharacter :many
SELECT * FROM public.character
ORDER BY id;

-- name: GetCharacterIDByName :one
SELECT id FROM public.character
WHERE name = $1 LIMIT 1;

-- name: CreateCharacter :one
INSERT INTO public.character (
  name,
  attack_range,
  equipment_slot,
  country_id,
  gender_id,
  attribute_id
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;