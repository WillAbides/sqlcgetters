// Code generated by sqlcgetters; DO NOT EDIT.

package override

import "github.com/kyleconroy/sqlc-testdata/pkg"

func (x Foo) GetOther() string {
	return x.Other
}

func (x Foo) GetTotal() int64 {
	return x.Total
}

func (x Foo) GetRetyped() pkg.CustomType {
	return x.Retyped
}