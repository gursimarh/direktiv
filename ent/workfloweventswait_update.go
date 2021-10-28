// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/ent/predicate"
	"github.com/direktiv/direktiv/ent/workflowevents"
	"github.com/direktiv/direktiv/ent/workfloweventswait"
)

// WorkflowEventsWaitUpdate is the builder for updating WorkflowEventsWait entities.
type WorkflowEventsWaitUpdate struct {
	config
	hooks    []Hook
	mutation *WorkflowEventsWaitMutation
}

// Where appends a list predicates to the WorkflowEventsWaitUpdate builder.
func (wewu *WorkflowEventsWaitUpdate) Where(ps ...predicate.WorkflowEventsWait) *WorkflowEventsWaitUpdate {
	wewu.mutation.Where(ps...)
	return wewu
}

// SetEvents sets the "events" field.
func (wewu *WorkflowEventsWaitUpdate) SetEvents(m map[string]interface{}) *WorkflowEventsWaitUpdate {
	wewu.mutation.SetEvents(m)
	return wewu
}

// SetWorkfloweventID sets the "workflowevent" edge to the WorkflowEvents entity by ID.
func (wewu *WorkflowEventsWaitUpdate) SetWorkfloweventID(id int) *WorkflowEventsWaitUpdate {
	wewu.mutation.SetWorkfloweventID(id)
	return wewu
}

// SetWorkflowevent sets the "workflowevent" edge to the WorkflowEvents entity.
func (wewu *WorkflowEventsWaitUpdate) SetWorkflowevent(w *WorkflowEvents) *WorkflowEventsWaitUpdate {
	return wewu.SetWorkfloweventID(w.ID)
}

// Mutation returns the WorkflowEventsWaitMutation object of the builder.
func (wewu *WorkflowEventsWaitUpdate) Mutation() *WorkflowEventsWaitMutation {
	return wewu.mutation
}

// ClearWorkflowevent clears the "workflowevent" edge to the WorkflowEvents entity.
func (wewu *WorkflowEventsWaitUpdate) ClearWorkflowevent() *WorkflowEventsWaitUpdate {
	wewu.mutation.ClearWorkflowevent()
	return wewu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wewu *WorkflowEventsWaitUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wewu.hooks) == 0 {
		if err = wewu.check(); err != nil {
			return 0, err
		}
		affected, err = wewu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkflowEventsWaitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wewu.check(); err != nil {
				return 0, err
			}
			wewu.mutation = mutation
			affected, err = wewu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wewu.hooks) - 1; i >= 0; i-- {
			if wewu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wewu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wewu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wewu *WorkflowEventsWaitUpdate) SaveX(ctx context.Context) int {
	affected, err := wewu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wewu *WorkflowEventsWaitUpdate) Exec(ctx context.Context) error {
	_, err := wewu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wewu *WorkflowEventsWaitUpdate) ExecX(ctx context.Context) {
	if err := wewu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wewu *WorkflowEventsWaitUpdate) check() error {
	if _, ok := wewu.mutation.WorkfloweventID(); wewu.mutation.WorkfloweventCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"workflowevent\"")
	}
	return nil
}

func (wewu *WorkflowEventsWaitUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workfloweventswait.Table,
			Columns: workfloweventswait.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workfloweventswait.FieldID,
			},
		},
	}
	if ps := wewu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wewu.mutation.Events(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workfloweventswait.FieldEvents,
		})
	}
	if wewu.mutation.WorkfloweventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workfloweventswait.WorkfloweventTable,
			Columns: []string{workfloweventswait.WorkfloweventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowevents.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wewu.mutation.WorkfloweventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workfloweventswait.WorkfloweventTable,
			Columns: []string{workfloweventswait.WorkfloweventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowevents.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wewu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workfloweventswait.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WorkflowEventsWaitUpdateOne is the builder for updating a single WorkflowEventsWait entity.
type WorkflowEventsWaitUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkflowEventsWaitMutation
}

// SetEvents sets the "events" field.
func (wewuo *WorkflowEventsWaitUpdateOne) SetEvents(m map[string]interface{}) *WorkflowEventsWaitUpdateOne {
	wewuo.mutation.SetEvents(m)
	return wewuo
}

// SetWorkfloweventID sets the "workflowevent" edge to the WorkflowEvents entity by ID.
func (wewuo *WorkflowEventsWaitUpdateOne) SetWorkfloweventID(id int) *WorkflowEventsWaitUpdateOne {
	wewuo.mutation.SetWorkfloweventID(id)
	return wewuo
}

// SetWorkflowevent sets the "workflowevent" edge to the WorkflowEvents entity.
func (wewuo *WorkflowEventsWaitUpdateOne) SetWorkflowevent(w *WorkflowEvents) *WorkflowEventsWaitUpdateOne {
	return wewuo.SetWorkfloweventID(w.ID)
}

// Mutation returns the WorkflowEventsWaitMutation object of the builder.
func (wewuo *WorkflowEventsWaitUpdateOne) Mutation() *WorkflowEventsWaitMutation {
	return wewuo.mutation
}

// ClearWorkflowevent clears the "workflowevent" edge to the WorkflowEvents entity.
func (wewuo *WorkflowEventsWaitUpdateOne) ClearWorkflowevent() *WorkflowEventsWaitUpdateOne {
	wewuo.mutation.ClearWorkflowevent()
	return wewuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wewuo *WorkflowEventsWaitUpdateOne) Select(field string, fields ...string) *WorkflowEventsWaitUpdateOne {
	wewuo.fields = append([]string{field}, fields...)
	return wewuo
}

// Save executes the query and returns the updated WorkflowEventsWait entity.
func (wewuo *WorkflowEventsWaitUpdateOne) Save(ctx context.Context) (*WorkflowEventsWait, error) {
	var (
		err  error
		node *WorkflowEventsWait
	)
	if len(wewuo.hooks) == 0 {
		if err = wewuo.check(); err != nil {
			return nil, err
		}
		node, err = wewuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkflowEventsWaitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wewuo.check(); err != nil {
				return nil, err
			}
			wewuo.mutation = mutation
			node, err = wewuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wewuo.hooks) - 1; i >= 0; i-- {
			if wewuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wewuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wewuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wewuo *WorkflowEventsWaitUpdateOne) SaveX(ctx context.Context) *WorkflowEventsWait {
	node, err := wewuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wewuo *WorkflowEventsWaitUpdateOne) Exec(ctx context.Context) error {
	_, err := wewuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wewuo *WorkflowEventsWaitUpdateOne) ExecX(ctx context.Context) {
	if err := wewuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wewuo *WorkflowEventsWaitUpdateOne) check() error {
	if _, ok := wewuo.mutation.WorkfloweventID(); wewuo.mutation.WorkfloweventCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"workflowevent\"")
	}
	return nil
}

func (wewuo *WorkflowEventsWaitUpdateOne) sqlSave(ctx context.Context) (_node *WorkflowEventsWait, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workfloweventswait.Table,
			Columns: workfloweventswait.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workfloweventswait.FieldID,
			},
		},
	}
	id, ok := wewuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing WorkflowEventsWait.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := wewuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workfloweventswait.FieldID)
		for _, f := range fields {
			if !workfloweventswait.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workfloweventswait.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wewuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wewuo.mutation.Events(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: workfloweventswait.FieldEvents,
		})
	}
	if wewuo.mutation.WorkfloweventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workfloweventswait.WorkfloweventTable,
			Columns: []string{workfloweventswait.WorkfloweventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowevents.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wewuo.mutation.WorkfloweventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workfloweventswait.WorkfloweventTable,
			Columns: []string{workfloweventswait.WorkfloweventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workflowevents.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WorkflowEventsWait{config: wewuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wewuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workfloweventswait.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
