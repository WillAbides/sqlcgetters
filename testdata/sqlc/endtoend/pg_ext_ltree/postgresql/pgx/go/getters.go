// Code generated by sqlcgetters; DO NOT EDIT.

package querytest

import "database/sql"

func (x Foo) GetQualifiedName() sql.NullString {
	return x.QualifiedName
}

func (x Foo) GetNameQuery() sql.NullString {
	return x.NameQuery
}

func (x Foo) GetFtsNameQuery() sql.NullString {
	return x.FtsNameQuery
}
