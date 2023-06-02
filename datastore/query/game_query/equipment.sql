-- name: GetEquipmentByID :one
SELECT * FROM public.equipment
WHERE id = $1 LIMIT 1;

-- name: ListAllEquipment :many
SELECT * FROM public.equipment
ORDER BY id;

-- name: GetEquipmentIDByName :one
SELECT id FROM public.equipment
WHERE name = $1 LIMIT 1;