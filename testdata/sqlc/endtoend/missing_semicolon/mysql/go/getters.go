// Code generated by sqlcgetters; DO NOT EDIT.

package querytest

import "database/sql"

func (x Author) GetID() int32 {
	return x.ID
}

func (x Author) GetName() string {
	return x.Name
}

func (x Author) GetBio() sql.NullString {
	return x.Bio
}

func (x SetAuthorParams) GetName() string {
	return x.Name
}

func (x SetAuthorParams) GetID() int32 {
	return x.ID
}
