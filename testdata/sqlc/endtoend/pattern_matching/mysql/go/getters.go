// Code generated by sqlcgetters; DO NOT EDIT.

package querytest

import "database/sql"

func (x Pet) GetName() sql.NullString {
	return x.Name
}