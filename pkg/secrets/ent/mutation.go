// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/direktiv/direktiv/pkg/secrets/ent/namespacesecret"
	"github.com/direktiv/direktiv/pkg/secrets/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeNamespaceSecret = "NamespaceSecret"
)

// NamespaceSecretMutation represents an operation that mutates the NamespaceSecret nodes in the graph.
type NamespaceSecretMutation struct {
	config
	op            Op
	typ           string
	id            *int
	ns            *string
	name          *string
	secret        *[]byte
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*NamespaceSecret, error)
	predicates    []predicate.NamespaceSecret
}

var _ ent.Mutation = (*NamespaceSecretMutation)(nil)

// namespacesecretOption allows management of the mutation configuration using functional options.
type namespacesecretOption func(*NamespaceSecretMutation)

// newNamespaceSecretMutation creates new mutation for the NamespaceSecret entity.
func newNamespaceSecretMutation(c config, op Op, opts ...namespacesecretOption) *NamespaceSecretMutation {
	m := &NamespaceSecretMutation{
		config:        c,
		op:            op,
		typ:           TypeNamespaceSecret,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withNamespaceSecretID sets the ID field of the mutation.
func withNamespaceSecretID(id int) namespacesecretOption {
	return func(m *NamespaceSecretMutation) {
		var (
			err   error
			once  sync.Once
			value *NamespaceSecret
		)
		m.oldValue = func(ctx context.Context) (*NamespaceSecret, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().NamespaceSecret.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withNamespaceSecret sets the old NamespaceSecret of the mutation.
func withNamespaceSecret(node *NamespaceSecret) namespacesecretOption {
	return func(m *NamespaceSecretMutation) {
		m.oldValue = func(context.Context) (*NamespaceSecret, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m NamespaceSecretMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m NamespaceSecretMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *NamespaceSecretMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetNs sets the "ns" field.
func (m *NamespaceSecretMutation) SetNs(s string) {
	m.ns = &s
}

// Ns returns the value of the "ns" field in the mutation.
func (m *NamespaceSecretMutation) Ns() (r string, exists bool) {
	v := m.ns
	if v == nil {
		return
	}
	return *v, true
}

// OldNs returns the old "ns" field's value of the NamespaceSecret entity.
// If the NamespaceSecret object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *NamespaceSecretMutation) OldNs(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldNs is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldNs requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNs: %w", err)
	}
	return oldValue.Ns, nil
}

// ResetNs resets all changes to the "ns" field.
func (m *NamespaceSecretMutation) ResetNs() {
	m.ns = nil
}

// SetName sets the "name" field.
func (m *NamespaceSecretMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *NamespaceSecretMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the NamespaceSecret entity.
// If the NamespaceSecret object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *NamespaceSecretMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *NamespaceSecretMutation) ResetName() {
	m.name = nil
}

// SetSecret sets the "secret" field.
func (m *NamespaceSecretMutation) SetSecret(b []byte) {
	m.secret = &b
}

// Secret returns the value of the "secret" field in the mutation.
func (m *NamespaceSecretMutation) Secret() (r []byte, exists bool) {
	v := m.secret
	if v == nil {
		return
	}
	return *v, true
}

// OldSecret returns the old "secret" field's value of the NamespaceSecret entity.
// If the NamespaceSecret object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *NamespaceSecretMutation) OldSecret(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldSecret is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldSecret requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSecret: %w", err)
	}
	return oldValue.Secret, nil
}

// ResetSecret resets all changes to the "secret" field.
func (m *NamespaceSecretMutation) ResetSecret() {
	m.secret = nil
}

// Where appends a list predicates to the NamespaceSecretMutation builder.
func (m *NamespaceSecretMutation) Where(ps ...predicate.NamespaceSecret) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *NamespaceSecretMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (NamespaceSecret).
func (m *NamespaceSecretMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *NamespaceSecretMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.ns != nil {
		fields = append(fields, namespacesecret.FieldNs)
	}
	if m.name != nil {
		fields = append(fields, namespacesecret.FieldName)
	}
	if m.secret != nil {
		fields = append(fields, namespacesecret.FieldSecret)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *NamespaceSecretMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case namespacesecret.FieldNs:
		return m.Ns()
	case namespacesecret.FieldName:
		return m.Name()
	case namespacesecret.FieldSecret:
		return m.Secret()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *NamespaceSecretMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case namespacesecret.FieldNs:
		return m.OldNs(ctx)
	case namespacesecret.FieldName:
		return m.OldName(ctx)
	case namespacesecret.FieldSecret:
		return m.OldSecret(ctx)
	}
	return nil, fmt.Errorf("unknown NamespaceSecret field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *NamespaceSecretMutation) SetField(name string, value ent.Value) error {
	switch name {
	case namespacesecret.FieldNs:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNs(v)
		return nil
	case namespacesecret.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case namespacesecret.FieldSecret:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSecret(v)
		return nil
	}
	return fmt.Errorf("unknown NamespaceSecret field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *NamespaceSecretMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *NamespaceSecretMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *NamespaceSecretMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown NamespaceSecret numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *NamespaceSecretMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *NamespaceSecretMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *NamespaceSecretMutation) ClearField(name string) error {
	return fmt.Errorf("unknown NamespaceSecret nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *NamespaceSecretMutation) ResetField(name string) error {
	switch name {
	case namespacesecret.FieldNs:
		m.ResetNs()
		return nil
	case namespacesecret.FieldName:
		m.ResetName()
		return nil
	case namespacesecret.FieldSecret:
		m.ResetSecret()
		return nil
	}
	return fmt.Errorf("unknown NamespaceSecret field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *NamespaceSecretMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *NamespaceSecretMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *NamespaceSecretMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *NamespaceSecretMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *NamespaceSecretMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *NamespaceSecretMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *NamespaceSecretMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown NamespaceSecret unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *NamespaceSecretMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown NamespaceSecret edge %s", name)
}
