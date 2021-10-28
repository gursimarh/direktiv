// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/direktiv/direktiv/pkg/flow/ent/ref"
	"github.com/direktiv/direktiv/pkg/flow/ent/revision"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
)

// Ref is the model entity for the Ref schema.
type Ref struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"-"`
	// Immutable holds the value of the "immutable" field.
	Immutable bool `json:"immutable,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RefQuery when eager-loading is set.
	Edges         RefEdges `json:"edges"`
	revision_refs *uuid.UUID
	workflow_refs *uuid.UUID
}

// RefEdges holds the relations/edges for other nodes in the graph.
type RefEdges struct {
	// Workflow holds the value of the workflow edge.
	Workflow *Workflow `json:"workflow,omitempty"`
	// Revision holds the value of the revision edge.
	Revision *Revision `json:"revision,omitempty"`
	// Routes holds the value of the routes edge.
	Routes []*Route `json:"routes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// WorkflowOrErr returns the Workflow value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RefEdges) WorkflowOrErr() (*Workflow, error) {
	if e.loadedTypes[0] {
		if e.Workflow == nil {
			// The edge workflow was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workflow.Label}
		}
		return e.Workflow, nil
	}
	return nil, &NotLoadedError{edge: "workflow"}
}

// RevisionOrErr returns the Revision value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RefEdges) RevisionOrErr() (*Revision, error) {
	if e.loadedTypes[1] {
		if e.Revision == nil {
			// The edge revision was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: revision.Label}
		}
		return e.Revision, nil
	}
	return nil, &NotLoadedError{edge: "revision"}
}

// RoutesOrErr returns the Routes value or an error if the edge
// was not loaded in eager-loading.
func (e RefEdges) RoutesOrErr() ([]*Route, error) {
	if e.loadedTypes[2] {
		return e.Routes, nil
	}
	return nil, &NotLoadedError{edge: "routes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ref) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case ref.FieldImmutable:
			values[i] = new(sql.NullBool)
		case ref.FieldName:
			values[i] = new(sql.NullString)
		case ref.FieldID:
			values[i] = new(uuid.UUID)
		case ref.ForeignKeys[0]: // revision_refs
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case ref.ForeignKeys[1]: // workflow_refs
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Ref", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ref fields.
func (r *Ref) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ref.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case ref.FieldImmutable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field immutable", values[i])
			} else if value.Valid {
				r.Immutable = value.Bool
			}
		case ref.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case ref.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field revision_refs", values[i])
			} else if value.Valid {
				r.revision_refs = new(uuid.UUID)
				*r.revision_refs = *value.S.(*uuid.UUID)
			}
		case ref.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_refs", values[i])
			} else if value.Valid {
				r.workflow_refs = new(uuid.UUID)
				*r.workflow_refs = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryWorkflow queries the "workflow" edge of the Ref entity.
func (r *Ref) QueryWorkflow() *WorkflowQuery {
	return (&RefClient{config: r.config}).QueryWorkflow(r)
}

// QueryRevision queries the "revision" edge of the Ref entity.
func (r *Ref) QueryRevision() *RevisionQuery {
	return (&RefClient{config: r.config}).QueryRevision(r)
}

// QueryRoutes queries the "routes" edge of the Ref entity.
func (r *Ref) QueryRoutes() *RouteQuery {
	return (&RefClient{config: r.config}).QueryRoutes(r)
}

// Update returns a builder for updating this Ref.
// Note that you need to call Ref.Unwrap() before calling this method if this Ref
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Ref) Update() *RefUpdateOne {
	return (&RefClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Ref entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Ref) Unwrap() *Ref {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ref is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Ref) String() string {
	var builder strings.Builder
	builder.WriteString("Ref(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", immutable=")
	builder.WriteString(fmt.Sprintf("%v", r.Immutable))
	builder.WriteString(", name=")
	builder.WriteString(r.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Refs is a parsable slice of Ref.
type Refs []*Ref

func (r Refs) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
