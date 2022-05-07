// Code generated by entc, DO NOT EDIT.

package file

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStorageKey holds the string denoting the storage_key field in the database.
	FieldStorageKey = "storage_key"
	// FieldExpiresAt holds the string denoting the expires_at field in the database.
	FieldExpiresAt = "expires_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the file in the database.
	Table = "files"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "files"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_files"
)

// Columns holds all SQL columns for file fields.
var Columns = []string{
	FieldID,
	FieldStorageKey,
	FieldExpiresAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "files"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_files",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// StorageKeyValidator is a validator for the "storage_key" field. It is called by the builders before save.
	StorageKeyValidator func(string) error
	// DefaultExpiresAt holds the default value on creation for the "expires_at" field.
	DefaultExpiresAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
