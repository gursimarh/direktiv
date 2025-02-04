// Code generated by entc, DO NOT EDIT.

package vardata

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the vardata type in the database.
	Label = "var_data"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// FieldMimeType holds the string denoting the mime_type field in the database.
	FieldMimeType = "mime_type"
	// EdgeVarrefs holds the string denoting the varrefs edge name in mutations.
	EdgeVarrefs = "varrefs"
	// Table holds the table name of the vardata in the database.
	Table = "var_data"
	// VarrefsTable is the table that holds the varrefs relation/edge.
	VarrefsTable = "var_refs"
	// VarrefsInverseTable is the table name for the VarRef entity.
	// It exists in this package in order to avoid circular dependency with the "varref" package.
	VarrefsInverseTable = "var_refs"
	// VarrefsColumn is the table column denoting the varrefs relation/edge.
	VarrefsColumn = "var_data_varrefs"
)

// Columns holds all SQL columns for vardata fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSize,
	FieldHash,
	FieldData,
	FieldMimeType,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultMimeType holds the default value on creation for the "mime_type" field.
	DefaultMimeType string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
