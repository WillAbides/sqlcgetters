// Code generated by sqlcgetters; DO NOT EDIT.

package querytest

import "database/sql"

func (x Foo) GetBar() sql.NullString {
	return x.Bar
}

func (x Foo) GetBat() string {
	return x.Bat
}

func (x CoalesceColumnsRow) GetBar() sql.NullString {
	return x.Bar
}

func (x CoalesceColumnsRow) GetBat() string {
	return x.Bat
}

func (x CoalesceColumnsRow) GetBar_2() string {
	return x.Bar_2
}