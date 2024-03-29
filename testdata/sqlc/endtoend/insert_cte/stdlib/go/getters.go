// Code generated by sqlcgetters; DO NOT EDIT.

package querytest

import (
	"database/sql"
	"time"
)

func (x Td3Code) GetID() int32 {
	return x.ID
}

func (x Td3Code) GetTsCreated() time.Time {
	return x.TsCreated
}

func (x Td3Code) GetTsUpdated() time.Time {
	return x.TsUpdated
}

func (x Td3Code) GetCreatedBy() string {
	return x.CreatedBy
}

func (x Td3Code) GetUpdatedBy() string {
	return x.UpdatedBy
}

func (x Td3Code) GetCode() sql.NullString {
	return x.Code
}

func (x Td3Code) GetHash() sql.NullString {
	return x.Hash
}

func (x Td3Code) GetIsPrivate() sql.NullBool {
	return x.IsPrivate
}

func (x Td3TestCode) GetID() int32 {
	return x.ID
}

func (x Td3TestCode) GetTsCreated() time.Time {
	return x.TsCreated
}

func (x Td3TestCode) GetTsUpdated() time.Time {
	return x.TsUpdated
}

func (x Td3TestCode) GetCreatedBy() string {
	return x.CreatedBy
}

func (x Td3TestCode) GetUpdatedBy() string {
	return x.UpdatedBy
}

func (x Td3TestCode) GetTestID() int32 {
	return x.TestID
}

func (x Td3TestCode) GetCodeHash() string {
	return x.CodeHash
}

func (x InsertCodeParams) GetCreatedBy() string {
	return x.CreatedBy
}

func (x InsertCodeParams) GetCode() sql.NullString {
	return x.Code
}

func (x InsertCodeParams) GetHash() sql.NullString {
	return x.Hash
}

func (x InsertCodeParams) GetTestID() int32 {
	return x.TestID
}
