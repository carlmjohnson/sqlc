// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: query.sql

package querytest

import (
	"context"
)

const allAuthors = `-- name: AllAuthors :many
SELECT  a.id,
        a.name,
        p.id as alias_id,
        p.name as alias_name
FROM    authors AS a
        LEFT JOIN authors AS p
            ON (authors.parent_id = p.id)
`

type AllAuthorsRow struct {
	ID        int64
	Name      string
	AliasID   int64
	AliasName string
}

func (q *Queries) AllAuthors(ctx context.Context) ([]AllAuthorsRow, error) {
	rows, err := q.db.QueryContext(ctx, allAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AllAuthorsRow
	for rows.Next() {
		var i AllAuthorsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AliasID,
			&i.AliasName,
		); err != nil {
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
