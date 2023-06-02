// Code generated by sqlc. DO NOT EDIT.
// source: c_e_file.sql

package sqlc

import (
	"context"
)

const getCeFileByIpfsUri = `-- name: GetCeFileByIpfsUri :one
SELECT id FROM public.c_e_file
WHERE ipfs_uri = $1 LIMIT 1
`

func (q *Queries) GetCeFileByIpfsUri(ctx context.Context, ipfsUri string) (int32, error) {
	row := q.db.QueryRowContext(ctx, getCeFileByIpfsUri, ipfsUri)
	var id int32
	err := row.Scan(&id)
	return id, err
}