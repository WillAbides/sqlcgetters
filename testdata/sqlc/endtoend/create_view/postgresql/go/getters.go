// Code generated by sqlcgetters; DO NOT EDIT.

package querytest

import "database/sql"

func (x FirstView) GetVal() string {
	return x.Val
}

func (x Foo) GetVal() string {
	return x.Val
}

func (x Foo) GetVal2() sql.NullInt32 {
	return x.Val2
}

func (x SecondView) GetVal() string {
	return x.Val
}

func (x SecondView) GetVal2() sql.NullInt32 {
	return x.Val2
}
