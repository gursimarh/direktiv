// Code generated by entc, DO NOT EDIT.

package namespace

import (
	"time"
)

const (
	// Label holds the string label denoting the namespace type in the database.
	Label = "namespace"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// EdgeWorkflows holds the string denoting the workflows edge name in mutations.
	EdgeWorkflows = "workflows"
	// Table holds the table name of the namespace in the database.
	Table = "namespaces"
	// WorkflowsTable is the table the holds the workflows relation/edge.
	WorkflowsTable = "workflows"
	// WorkflowsInverseTable is the table name for the Workflow entity.
	// It exists in this package in order to avoid circular dependency with the "workflow" package.
	WorkflowsInverseTable = "workflows"
	// WorkflowsColumn is the table column denoting the workflows relation/edge.
	WorkflowsColumn = "namespace_workflows"
)

// Columns holds all SQL columns for namespace fields.
var Columns = []string{
	FieldID,
	FieldCreated,
	FieldKey,
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
	// DefaultCreated holds the default value on creation for the "created" field.
	DefaultCreated func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
