// Code generated by sqlcgetters; DO NOT EDIT.

package querytest

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

func (x ListAuthorsColumnSortParams) GetMinID() int64 {
	return x.MinID
}

func (x ListAuthorsColumnSortParams) GetSortColumn() interface{} {
	return x.SortColumn
}
