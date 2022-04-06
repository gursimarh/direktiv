// Code generated by entc, DO NOT EDIT.

package inode

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the inode type in the database.
	Label = "inode"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldAttributes holds the string denoting the attributes field in the database.
	FieldAttributes = "attributes"
	// EdgeNamespace holds the string denoting the namespace edge name in mutations.
	EdgeNamespace = "namespace"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeWorkflow holds the string denoting the workflow edge name in mutations.
	EdgeWorkflow = "workflow"
	// EdgeAnnotations holds the string denoting the annotations edge name in mutations.
	EdgeAnnotations = "annotations"
	// Table holds the table name of the inode in the database.
	Table = "inodes"
	// NamespaceTable is the table that holds the namespace relation/edge.
	NamespaceTable = "inodes"
	// NamespaceInverseTable is the table name for the Namespace entity.
	// It exists in this package in order to avoid circular dependency with the "namespace" package.
	NamespaceInverseTable = "namespaces"
	// NamespaceColumn is the table column denoting the namespace relation/edge.
	NamespaceColumn = "namespace_inodes"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "inodes"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "inode_children"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "inodes"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "inode_children"
	// WorkflowTable is the table that holds the workflow relation/edge.
	WorkflowTable = "workflows"
	// WorkflowInverseTable is the table name for the Workflow entity.
	// It exists in this package in order to avoid circular dependency with the "workflow" package.
	WorkflowInverseTable = "workflows"
	// WorkflowColumn is the table column denoting the workflow relation/edge.
	WorkflowColumn = "inode_workflow"
	// AnnotationsTable is the table that holds the annotations relation/edge.
	AnnotationsTable = "annotations"
	// AnnotationsInverseTable is the table name for the Annotation entity.
	// It exists in this package in order to avoid circular dependency with the "annotation" package.
	AnnotationsInverseTable = "annotations"
	// AnnotationsColumn is the table column denoting the annotations relation/edge.
	AnnotationsColumn = "inode_annotations"
)

// Columns holds all SQL columns for inode fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldType,
	FieldAttributes,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "inodes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"inode_children",
	"namespace_inodes",
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
