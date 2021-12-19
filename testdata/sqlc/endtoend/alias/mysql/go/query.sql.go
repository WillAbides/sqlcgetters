// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
)

const aliasBar = `-- name: AliasBar :exec
DELETE FROM bar b
WHERE b.id = ?
`

func (q *Queries) AliasBar(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, aliasBar, id)
	return err
}
