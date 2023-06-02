-- name: GetUserByID :one
SELECT * FROM public.user
WHERE id = $1 LIMIT 1;

-- name: GetUserByWallet :one
SELECT * FROM public.user
WHERE wallet_address = $1 LIMIT 1;

-- name: ListAllUser :many
SELECT * FROM public.user
ORDER BY id;

-- name: CreateUserSignup :one
INSERT INTO public.user (
  wallet_address,
  email,
  name,
  otp_auth_secret
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;
