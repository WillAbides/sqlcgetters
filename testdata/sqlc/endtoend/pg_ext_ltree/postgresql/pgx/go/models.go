// Code generated by sqlc. DO NOT EDIT.

package querytest

import (
	"database/sql"
)

type Foo struct {
	QualifiedName sql.NullString
	NameQuery     sql.NullString
	FtsNameQuery  sql.NullString
}
