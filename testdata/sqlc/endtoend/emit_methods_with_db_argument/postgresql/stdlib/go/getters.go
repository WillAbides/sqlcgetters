// Code generated by sqlcgetters; DO NOT EDIT.

package querytest

import "database/sql"

func (x User) GetID() int32 {
	return x.ID
}

func (x User) GetFirstName() string {
	return x.FirstName
}

func (x User) GetLastName() sql.NullString {
	return x.LastName
}

func (x User) GetAge() int32 {
	return x.Age
}
