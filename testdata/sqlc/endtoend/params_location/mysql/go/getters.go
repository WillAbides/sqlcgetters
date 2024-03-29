// Code generated by sqlcgetters; DO NOT EDIT.

package querytest

import "database/sql"

func (x Order) GetID() int32 {
	return x.ID
}

func (x Order) GetPrice() string {
	return x.Price
}

func (x Order) GetUserID() int32 {
	return x.UserID
}

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

func (x User) GetJobStatus() string {
	return x.JobStatus
}

func (x GetUserByIDRow) GetFirstName() string {
	return x.FirstName
}

func (x GetUserByIDRow) GetID() int32 {
	return x.ID
}

func (x GetUserByIDRow) GetLastName() sql.NullString {
	return x.LastName
}

func (x InsertNewUserParams) GetFirstName() string {
	return x.FirstName
}

func (x InsertNewUserParams) GetLastName() sql.NullString {
	return x.LastName
}

func (x LimitSQLCArgRow) GetFirstName() string {
	return x.FirstName
}

func (x LimitSQLCArgRow) GetID() int32 {
	return x.ID
}

func (x ListUserOrdersRow) GetID() sql.NullInt32 {
	return x.ID
}

func (x ListUserOrdersRow) GetFirstName() sql.NullString {
	return x.FirstName
}

func (x ListUserOrdersRow) GetPrice() string {
	return x.Price
}

func (x ListUserParenExprParams) GetID() int32 {
	return x.ID
}

func (x ListUserParenExprParams) GetLimit() int32 {
	return x.Limit
}

func (x ListUsersByFamilyParams) GetMaxAge() int32 {
	return x.MaxAge
}

func (x ListUsersByFamilyParams) GetInFamily() sql.NullString {
	return x.InFamily
}

func (x ListUsersByFamilyRow) GetFirstName() string {
	return x.FirstName
}

func (x ListUsersByFamilyRow) GetLastName() sql.NullString {
	return x.LastName
}

func (x ListUsersByIDRow) GetFirstName() string {
	return x.FirstName
}

func (x ListUsersByIDRow) GetID() int32 {
	return x.ID
}

func (x ListUsersByIDRow) GetLastName() sql.NullString {
	return x.LastName
}

func (x ListUsersWithLimitRow) GetFirstName() string {
	return x.FirstName
}

func (x ListUsersWithLimitRow) GetLastName() sql.NullString {
	return x.LastName
}
