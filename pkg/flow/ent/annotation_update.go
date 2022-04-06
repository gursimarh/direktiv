// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/annotation"
	"github.com/direktiv/direktiv/pkg/flow/ent/inode"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// AnnotationUpdate is the builder for updating Annotation entities.
type AnnotationUpdate struct {
	config
	hooks    []Hook
	mutation *AnnotationMutation
}

// Where appends a list predicates to the AnnotationUpdate builder.
func (au *AnnotationUpdate) Where(ps ...predicate.Annotation) *AnnotationUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AnnotationUpdate) SetName(s string) *AnnotationUpdate {
	au.mutation.SetName(s)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AnnotationUpdate) SetUpdatedAt(t time.Time) *AnnotationUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetSize sets the "size" field.
func (au *AnnotationUpdate) SetSize(i int) *AnnotationUpdate {
	au.mutation.ResetSize()
	au.mutation.SetSize(i)
	return au
}

// AddSize adds i to the "size" field.
func (au *AnnotationUpdate) AddSize(i int) *AnnotationUpdate {
	au.mutation.AddSize(i)
	return au
}

// SetHash sets the "hash" field.
func (au *AnnotationUpdate) SetHash(s string) *AnnotationUpdate {
	au.mutation.SetHash(s)
	return au
}

// SetData sets the "data" field.
func (au *AnnotationUpdate) SetData(b []byte) *AnnotationUpdate {
	au.mutation.SetData(b)
	return au
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (au *AnnotationUpdate) SetNamespaceID(id uuid.UUID) *AnnotationUpdate {
	au.mutation.SetNamespaceID(id)
	return au
}

// SetNillableNamespaceID sets the "namespace" edge to the Namespace entity by ID if the given value is not nil.
func (au *AnnotationUpdate) SetNillableNamespaceID(id *uuid.UUID) *AnnotationUpdate {
	if id != nil {
		au = au.SetNamespaceID(*id)
	}
	return au
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (au *AnnotationUpdate) SetNamespace(n *Namespace) *AnnotationUpdate {
	return au.SetNamespaceID(n.ID)
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (au *AnnotationUpdate) SetWorkflowID(id uuid.UUID) *AnnotationUpdate {
	au.mutation.SetWorkflowID(id)
	return au
}

// SetNillableWorkflowID sets the "workflow" edge to the Workflow entity by ID if the given value is not nil.
func (au *AnnotationUpdate) SetNillableWorkflowID(id *uuid.UUID) *AnnotationUpdate {
	if id != nil {
		au = au.SetWorkflowID(*id)
	}
	return au
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (au *AnnotationUpdate) SetWorkflow(w *Workflow) *AnnotationUpdate {
	return au.SetWorkflowID(w.ID)
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (au *AnnotationUpdate) SetInstanceID(id uuid.UUID) *AnnotationUpdate {
	au.mutation.SetInstanceID(id)
	return au
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (au *AnnotationUpdate) SetNillableInstanceID(id *uuid.UUID) *AnnotationUpdate {
	if id != nil {
		au = au.SetInstanceID(*id)
	}
	return au
}

// SetInstance sets the "instance" edge to the Instance entity.
func (au *AnnotationUpdate) SetInstance(i *Instance) *AnnotationUpdate {
	return au.SetInstanceID(i.ID)
}

// SetInodeID sets the "inode" edge to the Inode entity by ID.
func (au *AnnotationUpdate) SetInodeID(id uuid.UUID) *AnnotationUpdate {
	au.mutation.SetInodeID(id)
	return au
}

// SetNillableInodeID sets the "inode" edge to the Inode entity by ID if the given value is not nil.
func (au *AnnotationUpdate) SetNillableInodeID(id *uuid.UUID) *AnnotationUpdate {
	if id != nil {
		au = au.SetInodeID(*id)
	}
	return au
}

// SetInode sets the "inode" edge to the Inode entity.
func (au *AnnotationUpdate) SetInode(i *Inode) *AnnotationUpdate {
	return au.SetInodeID(i.ID)
}

// Mutation returns the AnnotationMutation object of the builder.
func (au *AnnotationUpdate) Mutation() *AnnotationMutation {
	return au.mutation
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (au *AnnotationUpdate) ClearNamespace() *AnnotationUpdate {
	au.mutation.ClearNamespace()
	return au
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (au *AnnotationUpdate) ClearWorkflow() *AnnotationUpdate {
	au.mutation.ClearWorkflow()
	return au
}

// ClearInstance clears the "instance" edge to the Instance entity.
func (au *AnnotationUpdate) ClearInstance() *AnnotationUpdate {
	au.mutation.ClearInstance()
	return au
}

// ClearInode clears the "inode" edge to the Inode entity.
func (au *AnnotationUpdate) ClearInode() *AnnotationUpdate {
	au.mutation.ClearInode()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AnnotationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	au.defaults()
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AnnotationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AnnotationUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AnnotationUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AnnotationUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AnnotationUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := annotation.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AnnotationUpdate) check() error {
	if v, ok := au.mutation.Name(); ok {
		if err := annotation.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Annotation.name": %w`, err)}
		}
	}
	return nil
}

func (au *AnnotationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   annotation.Table,
			Columns: annotation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: annotation.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: annotation.FieldName,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: annotation.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: annotation.FieldSize,
		})
	}
	if value, ok := au.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: annotation.FieldSize,
		})
	}
	if value, ok := au.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: annotation.FieldHash,
		})
	}
	if value, ok := au.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: annotation.FieldData,
		})
	}
	if au.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.NamespaceTable,
			Columns: []string{annotation.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.NamespaceTable,
			Columns: []string{annotation.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.WorkflowTable,
			Columns: []string{annotation.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.WorkflowTable,
			Columns: []string{annotation.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.InstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.InstanceTable,
			Columns: []string{annotation.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.InstanceTable,
			Columns: []string{annotation.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.InodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.InodeTable,
			Columns: []string{annotation.InodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: inode.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.InodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.InodeTable,
			Columns: []string{annotation.InodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: inode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{annotation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AnnotationUpdateOne is the builder for updating a single Annotation entity.
type AnnotationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AnnotationMutation
}

// SetName sets the "name" field.
func (auo *AnnotationUpdateOne) SetName(s string) *AnnotationUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AnnotationUpdateOne) SetUpdatedAt(t time.Time) *AnnotationUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetSize sets the "size" field.
func (auo *AnnotationUpdateOne) SetSize(i int) *AnnotationUpdateOne {
	auo.mutation.ResetSize()
	auo.mutation.SetSize(i)
	return auo
}

// AddSize adds i to the "size" field.
func (auo *AnnotationUpdateOne) AddSize(i int) *AnnotationUpdateOne {
	auo.mutation.AddSize(i)
	return auo
}

// SetHash sets the "hash" field.
func (auo *AnnotationUpdateOne) SetHash(s string) *AnnotationUpdateOne {
	auo.mutation.SetHash(s)
	return auo
}

// SetData sets the "data" field.
func (auo *AnnotationUpdateOne) SetData(b []byte) *AnnotationUpdateOne {
	auo.mutation.SetData(b)
	return auo
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (auo *AnnotationUpdateOne) SetNamespaceID(id uuid.UUID) *AnnotationUpdateOne {
	auo.mutation.SetNamespaceID(id)
	return auo
}

// SetNillableNamespaceID sets the "namespace" edge to the Namespace entity by ID if the given value is not nil.
func (auo *AnnotationUpdateOne) SetNillableNamespaceID(id *uuid.UUID) *AnnotationUpdateOne {
	if id != nil {
		auo = auo.SetNamespaceID(*id)
	}
	return auo
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (auo *AnnotationUpdateOne) SetNamespace(n *Namespace) *AnnotationUpdateOne {
	return auo.SetNamespaceID(n.ID)
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (auo *AnnotationUpdateOne) SetWorkflowID(id uuid.UUID) *AnnotationUpdateOne {
	auo.mutation.SetWorkflowID(id)
	return auo
}

// SetNillableWorkflowID sets the "workflow" edge to the Workflow entity by ID if the given value is not nil.
func (auo *AnnotationUpdateOne) SetNillableWorkflowID(id *uuid.UUID) *AnnotationUpdateOne {
	if id != nil {
		auo = auo.SetWorkflowID(*id)
	}
	return auo
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (auo *AnnotationUpdateOne) SetWorkflow(w *Workflow) *AnnotationUpdateOne {
	return auo.SetWorkflowID(w.ID)
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (auo *AnnotationUpdateOne) SetInstanceID(id uuid.UUID) *AnnotationUpdateOne {
	auo.mutation.SetInstanceID(id)
	return auo
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (auo *AnnotationUpdateOne) SetNillableInstanceID(id *uuid.UUID) *AnnotationUpdateOne {
	if id != nil {
		auo = auo.SetInstanceID(*id)
	}
	return auo
}

// SetInstance sets the "instance" edge to the Instance entity.
func (auo *AnnotationUpdateOne) SetInstance(i *Instance) *AnnotationUpdateOne {
	return auo.SetInstanceID(i.ID)
}

// SetInodeID sets the "inode" edge to the Inode entity by ID.
func (auo *AnnotationUpdateOne) SetInodeID(id uuid.UUID) *AnnotationUpdateOne {
	auo.mutation.SetInodeID(id)
	return auo
}

// SetNillableInodeID sets the "inode" edge to the Inode entity by ID if the given value is not nil.
func (auo *AnnotationUpdateOne) SetNillableInodeID(id *uuid.UUID) *AnnotationUpdateOne {
	if id != nil {
		auo = auo.SetInodeID(*id)
	}
	return auo
}

// SetInode sets the "inode" edge to the Inode entity.
func (auo *AnnotationUpdateOne) SetInode(i *Inode) *AnnotationUpdateOne {
	return auo.SetInodeID(i.ID)
}

// Mutation returns the AnnotationMutation object of the builder.
func (auo *AnnotationUpdateOne) Mutation() *AnnotationMutation {
	return auo.mutation
}

// ClearNamespace clears the "namespace" edge to the Namespace entity.
func (auo *AnnotationUpdateOne) ClearNamespace() *AnnotationUpdateOne {
	auo.mutation.ClearNamespace()
	return auo
}

// ClearWorkflow clears the "workflow" edge to the Workflow entity.
func (auo *AnnotationUpdateOne) ClearWorkflow() *AnnotationUpdateOne {
	auo.mutation.ClearWorkflow()
	return auo
}

// ClearInstance clears the "instance" edge to the Instance entity.
func (auo *AnnotationUpdateOne) ClearInstance() *AnnotationUpdateOne {
	auo.mutation.ClearInstance()
	return auo
}

// ClearInode clears the "inode" edge to the Inode entity.
func (auo *AnnotationUpdateOne) ClearInode() *AnnotationUpdateOne {
	auo.mutation.ClearInode()
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AnnotationUpdateOne) Select(field string, fields ...string) *AnnotationUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Annotation entity.
func (auo *AnnotationUpdateOne) Save(ctx context.Context) (*Annotation, error) {
	var (
		err  error
		node *Annotation
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AnnotationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AnnotationUpdateOne) SaveX(ctx context.Context) *Annotation {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AnnotationUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AnnotationUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AnnotationUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := annotation.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AnnotationUpdateOne) check() error {
	if v, ok := auo.mutation.Name(); ok {
		if err := annotation.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Annotation.name": %w`, err)}
		}
	}
	return nil
}

func (auo *AnnotationUpdateOne) sqlSave(ctx context.Context) (_node *Annotation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   annotation.Table,
			Columns: annotation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: annotation.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Annotation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, annotation.FieldID)
		for _, f := range fields {
			if !annotation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != annotation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: annotation.FieldName,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: annotation.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: annotation.FieldSize,
		})
	}
	if value, ok := auo.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: annotation.FieldSize,
		})
	}
	if value, ok := auo.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: annotation.FieldHash,
		})
	}
	if value, ok := auo.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: annotation.FieldData,
		})
	}
	if auo.mutation.NamespaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.NamespaceTable,
			Columns: []string{annotation.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.NamespaceTable,
			Columns: []string{annotation.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.WorkflowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.WorkflowTable,
			Columns: []string{annotation.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.WorkflowTable,
			Columns: []string{annotation.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.InstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.InstanceTable,
			Columns: []string{annotation.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.InstanceTable,
			Columns: []string{annotation.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.InodeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.InodeTable,
			Columns: []string{annotation.InodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: inode.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.InodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   annotation.InodeTable,
			Columns: []string{annotation.InodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: inode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Annotation{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{annotation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
