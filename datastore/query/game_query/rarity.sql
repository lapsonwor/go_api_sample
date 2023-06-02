-- name: CreateRarity :one
INSERT INTO public.rarity (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: GetRarityIdByName :one
SELECT id FROM public.rarity
WHERE name = $1 LIMIT 1; 