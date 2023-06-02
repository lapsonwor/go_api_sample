-- name: GetMiniGameRank :many
SELECT * FROM public.miniGame_rank
ORDER BY id;

-- name: GetMiniGameRankByWallet :one
SELECT * FROM public.miniGame_rank
WHERE wallet_address = $1;

-- name: CreateMiniGameRank :one
INSERT INTO public.miniGame_rank (
  wallet_address,
  mark
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateMiniGameRank :one
UPDATE public.miniGame_rank
SET mark = $1, updated_at = $2
WHERE wallet_address = $3
RETURNING *;

-- name: GetRankByWallet :one
SELECT mgr.mark, rank FROM (SELECT wallet_address, mark, RANK () OVER ( 
	ORDER BY mark DESC
) rank FROM public.miniGame_rank) mgr
WHERE wallet_address = $1;

-- name: GetTop10Scores :many
SELECT wallet_address, mark FROM public.miniGame_rank
ORDER BY mark DESC LIMIT 10; 