// Code generated by sqlcgetters; DO NOT EDIT.

package primary_key_later

import "database/sql"

func (x Author) GetID() int64 {
	return x.ID
}

func (x Author) GetName() string {
	return x.Name
}

func (x Author) GetBio() sql.NullString {
	return x.Bio
}