// Code generated by sqlcgetters; DO NOT EDIT.

package querytest

import "database/sql"

func (x Bar) GetID() int32 {
	return x.ID
}

func (x Bar) GetTitle() sql.NullString {
	return x.Title
}

func (x Foo) GetID() int32 {
	return x.ID
}

func (x AliasExpandRow) GetID() int32 {
	return x.ID
}

func (x AliasExpandRow) GetID_2() int32 {
	return x.ID_2
}

func (x AliasExpandRow) GetTitle() sql.NullString {
	return x.Title
}

func (x AliasJoinRow) GetID() int32 {
	return x.ID
}

func (x AliasJoinRow) GetTitle() sql.NullString {
	return x.Title
}
