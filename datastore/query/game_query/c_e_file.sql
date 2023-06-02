-- name: GetCeFileByIpfsUri :one
SELECT id FROM public.c_e_file
WHERE ipfs_uri = $1 LIMIT 1;