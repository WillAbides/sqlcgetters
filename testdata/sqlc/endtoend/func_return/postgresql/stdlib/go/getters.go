// Code generated by sqlcgetters; DO NOT EDIT.

package querytest

import (
	"database/sql"

	"github.com/tabbed/pqtype"
)

func (x User) GetID() sql.NullInt32 {
	return x.ID
}

func (x User) GetFirstName() string {
	return x.FirstName
}

func (x GenerateSeriesParams) GetColumn1() pqtype.Inet {
	return x.Column1
}

func (x GenerateSeriesParams) GetColumn2() int32 {
	return x.Column2
}