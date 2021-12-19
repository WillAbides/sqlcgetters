// Code generated by sqlcgetters; DO NOT EDIT.

package querytest

import "database/sql"

func (x Author) GetID() int64 {
	return x.ID
}

func (x Author) GetUsername() sql.NullString {
	return x.Username
}

func (x Author) GetEmail() sql.NullString {
	return x.Email
}

func (x Author) GetName() string {
	return x.Name
}

func (x Author) GetBio() sql.NullString {
	return x.Bio
}

func (x ListAuthorsParams) GetEmail() string {
	return x.Email
}

func (x ListAuthorsParams) GetUsername() string {
	return x.Username
}