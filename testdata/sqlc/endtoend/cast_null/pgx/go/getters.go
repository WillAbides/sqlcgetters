// Code generated by sqlcgetters; DO NOT EDIT.

package querytest

import "database/sql"

func (x Foo) GetBar() sql.NullString {
	return x.Bar
}

func (x ListNullableRow) GetA() sql.NullString {
	return x.A
}

func (x ListNullableRow) GetB() sql.NullInt32 {
	return x.B
}

func (x ListNullableRow) GetC() sql.NullInt64 {
	return x.C
}

func (x ListNullableRow) GetD() sql.NullTime {
	return x.D
}
