-- name: GetWhitelistByWallet :one
SELECT * FROM public.whitelist
WHERE wallet_address = $1 LIMIT 1;
