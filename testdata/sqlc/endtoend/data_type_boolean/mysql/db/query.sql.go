// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package db

import (
	"context"
)

const listBar = `-- name: ListBar :many
SELECT col_a, col_b, col_c FROM bar
`

func (q *Queries) ListBar(ctx context.Context) ([]Bar, error) {
	rows, err := q.db.QueryContext(ctx, listBar)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Bar
	for rows.Next() {
		var i Bar
		if err := rows.Scan(&i.ColA, &i.ColB, &i.ColC); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listFoo = `-- name: ListFoo :many
SELECT col_a, col_b, col_c FROM foo
`

func (q *Queries) ListFoo(ctx context.Context) ([]Foo, error) {
	rows, err := q.db.QueryContext(ctx, listFoo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Foo
	for rows.Next() {
		var i Foo
		if err := rows.Scan(&i.ColA, &i.ColB, &i.ColC); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
