// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/instanceruntime"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
)

// InstanceRuntimeUpdate is the builder for updating InstanceRuntime entities.
type InstanceRuntimeUpdate struct {
	config
	hooks    []Hook
	mutation *InstanceRuntimeMutation
}

// Where appends a list predicates to the InstanceRuntimeUpdate builder.
func (iru *InstanceRuntimeUpdate) Where(ps ...predicate.InstanceRuntime) *InstanceRuntimeUpdate {
	iru.mutation.Where(ps...)
	return iru
}

// SetData sets the "data" field.
func (iru *InstanceRuntimeUpdate) SetData(s string) *InstanceRuntimeUpdate {
	iru.mutation.SetData(s)
	return iru
}

// SetController sets the "controller" field.
func (iru *InstanceRuntimeUpdate) SetController(s string) *InstanceRuntimeUpdate {
	iru.mutation.SetController(s)
	return iru
}

// SetNillableController sets the "controller" field if the given value is not nil.
func (iru *InstanceRuntimeUpdate) SetNillableController(s *string) *InstanceRuntimeUpdate {
	if s != nil {
		iru.SetController(*s)
	}
	return iru
}

// ClearController clears the value of the "controller" field.
func (iru *InstanceRuntimeUpdate) ClearController() *InstanceRuntimeUpdate {
	iru.mutation.ClearController()
	return iru
}

// SetMemory sets the "memory" field.
func (iru *InstanceRuntimeUpdate) SetMemory(s string) *InstanceRuntimeUpdate {
	iru.mutation.SetMemory(s)
	return iru
}

// SetNillableMemory sets the "memory" field if the given value is not nil.
func (iru *InstanceRuntimeUpdate) SetNillableMemory(s *string) *InstanceRuntimeUpdate {
	if s != nil {
		iru.SetMemory(*s)
	}
	return iru
}

// ClearMemory clears the value of the "memory" field.
func (iru *InstanceRuntimeUpdate) ClearMemory() *InstanceRuntimeUpdate {
	iru.mutation.ClearMemory()
	return iru
}

// SetFlow sets the "flow" field.
func (iru *InstanceRuntimeUpdate) SetFlow(s []string) *InstanceRuntimeUpdate {
	iru.mutation.SetFlow(s)
	return iru
}

// ClearFlow clears the value of the "flow" field.
func (iru *InstanceRuntimeUpdate) ClearFlow() *InstanceRuntimeUpdate {
	iru.mutation.ClearFlow()
	return iru
}

// SetOutput sets the "output" field.
func (iru *InstanceRuntimeUpdate) SetOutput(s string) *InstanceRuntimeUpdate {
	iru.mutation.SetOutput(s)
	return iru
}

// SetNillableOutput sets the "output" field if the given value is not nil.
func (iru *InstanceRuntimeUpdate) SetNillableOutput(s *string) *InstanceRuntimeUpdate {
	if s != nil {
		iru.SetOutput(*s)
	}
	return iru
}

// ClearOutput clears the value of the "output" field.
func (iru *InstanceRuntimeUpdate) ClearOutput() *InstanceRuntimeUpdate {
	iru.mutation.ClearOutput()
	return iru
}

// SetStateBeginTime sets the "stateBeginTime" field.
func (iru *InstanceRuntimeUpdate) SetStateBeginTime(t time.Time) *InstanceRuntimeUpdate {
	iru.mutation.SetStateBeginTime(t)
	return iru
}

// SetNillableStateBeginTime sets the "stateBeginTime" field if the given value is not nil.
func (iru *InstanceRuntimeUpdate) SetNillableStateBeginTime(t *time.Time) *InstanceRuntimeUpdate {
	if t != nil {
		iru.SetStateBeginTime(*t)
	}
	return iru
}

// ClearStateBeginTime clears the value of the "stateBeginTime" field.
func (iru *InstanceRuntimeUpdate) ClearStateBeginTime() *InstanceRuntimeUpdate {
	iru.mutation.ClearStateBeginTime()
	return iru
}

// SetDeadline sets the "deadline" field.
func (iru *InstanceRuntimeUpdate) SetDeadline(t time.Time) *InstanceRuntimeUpdate {
	iru.mutation.SetDeadline(t)
	return iru
}

// SetNillableDeadline sets the "deadline" field if the given value is not nil.
func (iru *InstanceRuntimeUpdate) SetNillableDeadline(t *time.Time) *InstanceRuntimeUpdate {
	if t != nil {
		iru.SetDeadline(*t)
	}
	return iru
}

// ClearDeadline clears the value of the "deadline" field.
func (iru *InstanceRuntimeUpdate) ClearDeadline() *InstanceRuntimeUpdate {
	iru.mutation.ClearDeadline()
	return iru
}

// SetAttempts sets the "attempts" field.
func (iru *InstanceRuntimeUpdate) SetAttempts(i int) *InstanceRuntimeUpdate {
	iru.mutation.ResetAttempts()
	iru.mutation.SetAttempts(i)
	return iru
}

// SetNillableAttempts sets the "attempts" field if the given value is not nil.
func (iru *InstanceRuntimeUpdate) SetNillableAttempts(i *int) *InstanceRuntimeUpdate {
	if i != nil {
		iru.SetAttempts(*i)
	}
	return iru
}

// AddAttempts adds i to the "attempts" field.
func (iru *InstanceRuntimeUpdate) AddAttempts(i int) *InstanceRuntimeUpdate {
	iru.mutation.AddAttempts(i)
	return iru
}

// ClearAttempts clears the value of the "attempts" field.
func (iru *InstanceRuntimeUpdate) ClearAttempts() *InstanceRuntimeUpdate {
	iru.mutation.ClearAttempts()
	return iru
}

// SetCallerData sets the "caller_data" field.
func (iru *InstanceRuntimeUpdate) SetCallerData(s string) *InstanceRuntimeUpdate {
	iru.mutation.SetCallerData(s)
	return iru
}

// SetNillableCallerData sets the "caller_data" field if the given value is not nil.
func (iru *InstanceRuntimeUpdate) SetNillableCallerData(s *string) *InstanceRuntimeUpdate {
	if s != nil {
		iru.SetCallerData(*s)
	}
	return iru
}

// ClearCallerData clears the value of the "caller_data" field.
func (iru *InstanceRuntimeUpdate) ClearCallerData() *InstanceRuntimeUpdate {
	iru.mutation.ClearCallerData()
	return iru
}

// SetInstanceContext sets the "instanceContext" field.
func (iru *InstanceRuntimeUpdate) SetInstanceContext(s string) *InstanceRuntimeUpdate {
	iru.mutation.SetInstanceContext(s)
	return iru
}

// SetNillableInstanceContext sets the "instanceContext" field if the given value is not nil.
func (iru *InstanceRuntimeUpdate) SetNillableInstanceContext(s *string) *InstanceRuntimeUpdate {
	if s != nil {
		iru.SetInstanceContext(*s)
	}
	return iru
}

// ClearInstanceContext clears the value of the "instanceContext" field.
func (iru *InstanceRuntimeUpdate) ClearInstanceContext() *InstanceRuntimeUpdate {
	iru.mutation.ClearInstanceContext()
	return iru
}

// SetStateContext sets the "stateContext" field.
func (iru *InstanceRuntimeUpdate) SetStateContext(s string) *InstanceRuntimeUpdate {
	iru.mutation.SetStateContext(s)
	return iru
}

// SetNillableStateContext sets the "stateContext" field if the given value is not nil.
func (iru *InstanceRuntimeUpdate) SetNillableStateContext(s *string) *InstanceRuntimeUpdate {
	if s != nil {
		iru.SetStateContext(*s)
	}
	return iru
}

// ClearStateContext clears the value of the "stateContext" field.
func (iru *InstanceRuntimeUpdate) ClearStateContext() *InstanceRuntimeUpdate {
	iru.mutation.ClearStateContext()
	return iru
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (iru *InstanceRuntimeUpdate) SetInstanceID(id uuid.UUID) *InstanceRuntimeUpdate {
	iru.mutation.SetInstanceID(id)
	return iru
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (iru *InstanceRuntimeUpdate) SetNillableInstanceID(id *uuid.UUID) *InstanceRuntimeUpdate {
	if id != nil {
		iru = iru.SetInstanceID(*id)
	}
	return iru
}

// SetInstance sets the "instance" edge to the Instance entity.
func (iru *InstanceRuntimeUpdate) SetInstance(i *Instance) *InstanceRuntimeUpdate {
	return iru.SetInstanceID(i.ID)
}

// SetCallerID sets the "caller" edge to the Instance entity by ID.
func (iru *InstanceRuntimeUpdate) SetCallerID(id uuid.UUID) *InstanceRuntimeUpdate {
	iru.mutation.SetCallerID(id)
	return iru
}

// SetNillableCallerID sets the "caller" edge to the Instance entity by ID if the given value is not nil.
func (iru *InstanceRuntimeUpdate) SetNillableCallerID(id *uuid.UUID) *InstanceRuntimeUpdate {
	if id != nil {
		iru = iru.SetCallerID(*id)
	}
	return iru
}

// SetCaller sets the "caller" edge to the Instance entity.
func (iru *InstanceRuntimeUpdate) SetCaller(i *Instance) *InstanceRuntimeUpdate {
	return iru.SetCallerID(i.ID)
}

// Mutation returns the InstanceRuntimeMutation object of the builder.
func (iru *InstanceRuntimeUpdate) Mutation() *InstanceRuntimeMutation {
	return iru.mutation
}

// ClearInstance clears the "instance" edge to the Instance entity.
func (iru *InstanceRuntimeUpdate) ClearInstance() *InstanceRuntimeUpdate {
	iru.mutation.ClearInstance()
	return iru
}

// ClearCaller clears the "caller" edge to the Instance entity.
func (iru *InstanceRuntimeUpdate) ClearCaller() *InstanceRuntimeUpdate {
	iru.mutation.ClearCaller()
	return iru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iru *InstanceRuntimeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iru.hooks) == 0 {
		affected, err = iru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstanceRuntimeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iru.mutation = mutation
			affected, err = iru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iru.hooks) - 1; i >= 0; i-- {
			if iru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iru *InstanceRuntimeUpdate) SaveX(ctx context.Context) int {
	affected, err := iru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iru *InstanceRuntimeUpdate) Exec(ctx context.Context) error {
	_, err := iru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iru *InstanceRuntimeUpdate) ExecX(ctx context.Context) {
	if err := iru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iru *InstanceRuntimeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instanceruntime.Table,
			Columns: instanceruntime.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: instanceruntime.FieldID,
			},
		},
	}
	if ps := iru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iru.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldData,
		})
	}
	if value, ok := iru.mutation.Controller(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldController,
		})
	}
	if iru.mutation.ControllerCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instanceruntime.FieldController,
		})
	}
	if value, ok := iru.mutation.Memory(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldMemory,
		})
	}
	if iru.mutation.MemoryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instanceruntime.FieldMemory,
		})
	}
	if value, ok := iru.mutation.Flow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: instanceruntime.FieldFlow,
		})
	}
	if iru.mutation.FlowCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: instanceruntime.FieldFlow,
		})
	}
	if value, ok := iru.mutation.Output(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldOutput,
		})
	}
	if iru.mutation.OutputCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instanceruntime.FieldOutput,
		})
	}
	if value, ok := iru.mutation.StateBeginTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instanceruntime.FieldStateBeginTime,
		})
	}
	if iru.mutation.StateBeginTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: instanceruntime.FieldStateBeginTime,
		})
	}
	if value, ok := iru.mutation.Deadline(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instanceruntime.FieldDeadline,
		})
	}
	if iru.mutation.DeadlineCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: instanceruntime.FieldDeadline,
		})
	}
	if value, ok := iru.mutation.Attempts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: instanceruntime.FieldAttempts,
		})
	}
	if value, ok := iru.mutation.AddedAttempts(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: instanceruntime.FieldAttempts,
		})
	}
	if iru.mutation.AttemptsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: instanceruntime.FieldAttempts,
		})
	}
	if value, ok := iru.mutation.CallerData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldCallerData,
		})
	}
	if iru.mutation.CallerDataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instanceruntime.FieldCallerData,
		})
	}
	if value, ok := iru.mutation.InstanceContext(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldInstanceContext,
		})
	}
	if iru.mutation.InstanceContextCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instanceruntime.FieldInstanceContext,
		})
	}
	if value, ok := iru.mutation.StateContext(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldStateContext,
		})
	}
	if iru.mutation.StateContextCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instanceruntime.FieldStateContext,
		})
	}
	if iru.mutation.InstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   instanceruntime.InstanceTable,
			Columns: []string{instanceruntime.InstanceColumn},
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
	if nodes := iru.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   instanceruntime.InstanceTable,
			Columns: []string{instanceruntime.InstanceColumn},
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
	if iru.mutation.CallerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instanceruntime.CallerTable,
			Columns: []string{instanceruntime.CallerColumn},
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
	if nodes := iru.mutation.CallerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instanceruntime.CallerTable,
			Columns: []string{instanceruntime.CallerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, iru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{instanceruntime.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// InstanceRuntimeUpdateOne is the builder for updating a single InstanceRuntime entity.
type InstanceRuntimeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InstanceRuntimeMutation
}

// SetData sets the "data" field.
func (iruo *InstanceRuntimeUpdateOne) SetData(s string) *InstanceRuntimeUpdateOne {
	iruo.mutation.SetData(s)
	return iruo
}

// SetController sets the "controller" field.
func (iruo *InstanceRuntimeUpdateOne) SetController(s string) *InstanceRuntimeUpdateOne {
	iruo.mutation.SetController(s)
	return iruo
}

// SetNillableController sets the "controller" field if the given value is not nil.
func (iruo *InstanceRuntimeUpdateOne) SetNillableController(s *string) *InstanceRuntimeUpdateOne {
	if s != nil {
		iruo.SetController(*s)
	}
	return iruo
}

// ClearController clears the value of the "controller" field.
func (iruo *InstanceRuntimeUpdateOne) ClearController() *InstanceRuntimeUpdateOne {
	iruo.mutation.ClearController()
	return iruo
}

// SetMemory sets the "memory" field.
func (iruo *InstanceRuntimeUpdateOne) SetMemory(s string) *InstanceRuntimeUpdateOne {
	iruo.mutation.SetMemory(s)
	return iruo
}

// SetNillableMemory sets the "memory" field if the given value is not nil.
func (iruo *InstanceRuntimeUpdateOne) SetNillableMemory(s *string) *InstanceRuntimeUpdateOne {
	if s != nil {
		iruo.SetMemory(*s)
	}
	return iruo
}

// ClearMemory clears the value of the "memory" field.
func (iruo *InstanceRuntimeUpdateOne) ClearMemory() *InstanceRuntimeUpdateOne {
	iruo.mutation.ClearMemory()
	return iruo
}

// SetFlow sets the "flow" field.
func (iruo *InstanceRuntimeUpdateOne) SetFlow(s []string) *InstanceRuntimeUpdateOne {
	iruo.mutation.SetFlow(s)
	return iruo
}

// ClearFlow clears the value of the "flow" field.
func (iruo *InstanceRuntimeUpdateOne) ClearFlow() *InstanceRuntimeUpdateOne {
	iruo.mutation.ClearFlow()
	return iruo
}

// SetOutput sets the "output" field.
func (iruo *InstanceRuntimeUpdateOne) SetOutput(s string) *InstanceRuntimeUpdateOne {
	iruo.mutation.SetOutput(s)
	return iruo
}

// SetNillableOutput sets the "output" field if the given value is not nil.
func (iruo *InstanceRuntimeUpdateOne) SetNillableOutput(s *string) *InstanceRuntimeUpdateOne {
	if s != nil {
		iruo.SetOutput(*s)
	}
	return iruo
}

// ClearOutput clears the value of the "output" field.
func (iruo *InstanceRuntimeUpdateOne) ClearOutput() *InstanceRuntimeUpdateOne {
	iruo.mutation.ClearOutput()
	return iruo
}

// SetStateBeginTime sets the "stateBeginTime" field.
func (iruo *InstanceRuntimeUpdateOne) SetStateBeginTime(t time.Time) *InstanceRuntimeUpdateOne {
	iruo.mutation.SetStateBeginTime(t)
	return iruo
}

// SetNillableStateBeginTime sets the "stateBeginTime" field if the given value is not nil.
func (iruo *InstanceRuntimeUpdateOne) SetNillableStateBeginTime(t *time.Time) *InstanceRuntimeUpdateOne {
	if t != nil {
		iruo.SetStateBeginTime(*t)
	}
	return iruo
}

// ClearStateBeginTime clears the value of the "stateBeginTime" field.
func (iruo *InstanceRuntimeUpdateOne) ClearStateBeginTime() *InstanceRuntimeUpdateOne {
	iruo.mutation.ClearStateBeginTime()
	return iruo
}

// SetDeadline sets the "deadline" field.
func (iruo *InstanceRuntimeUpdateOne) SetDeadline(t time.Time) *InstanceRuntimeUpdateOne {
	iruo.mutation.SetDeadline(t)
	return iruo
}

// SetNillableDeadline sets the "deadline" field if the given value is not nil.
func (iruo *InstanceRuntimeUpdateOne) SetNillableDeadline(t *time.Time) *InstanceRuntimeUpdateOne {
	if t != nil {
		iruo.SetDeadline(*t)
	}
	return iruo
}

// ClearDeadline clears the value of the "deadline" field.
func (iruo *InstanceRuntimeUpdateOne) ClearDeadline() *InstanceRuntimeUpdateOne {
	iruo.mutation.ClearDeadline()
	return iruo
}

// SetAttempts sets the "attempts" field.
func (iruo *InstanceRuntimeUpdateOne) SetAttempts(i int) *InstanceRuntimeUpdateOne {
	iruo.mutation.ResetAttempts()
	iruo.mutation.SetAttempts(i)
	return iruo
}

// SetNillableAttempts sets the "attempts" field if the given value is not nil.
func (iruo *InstanceRuntimeUpdateOne) SetNillableAttempts(i *int) *InstanceRuntimeUpdateOne {
	if i != nil {
		iruo.SetAttempts(*i)
	}
	return iruo
}

// AddAttempts adds i to the "attempts" field.
func (iruo *InstanceRuntimeUpdateOne) AddAttempts(i int) *InstanceRuntimeUpdateOne {
	iruo.mutation.AddAttempts(i)
	return iruo
}

// ClearAttempts clears the value of the "attempts" field.
func (iruo *InstanceRuntimeUpdateOne) ClearAttempts() *InstanceRuntimeUpdateOne {
	iruo.mutation.ClearAttempts()
	return iruo
}

// SetCallerData sets the "caller_data" field.
func (iruo *InstanceRuntimeUpdateOne) SetCallerData(s string) *InstanceRuntimeUpdateOne {
	iruo.mutation.SetCallerData(s)
	return iruo
}

// SetNillableCallerData sets the "caller_data" field if the given value is not nil.
func (iruo *InstanceRuntimeUpdateOne) SetNillableCallerData(s *string) *InstanceRuntimeUpdateOne {
	if s != nil {
		iruo.SetCallerData(*s)
	}
	return iruo
}

// ClearCallerData clears the value of the "caller_data" field.
func (iruo *InstanceRuntimeUpdateOne) ClearCallerData() *InstanceRuntimeUpdateOne {
	iruo.mutation.ClearCallerData()
	return iruo
}

// SetInstanceContext sets the "instanceContext" field.
func (iruo *InstanceRuntimeUpdateOne) SetInstanceContext(s string) *InstanceRuntimeUpdateOne {
	iruo.mutation.SetInstanceContext(s)
	return iruo
}

// SetNillableInstanceContext sets the "instanceContext" field if the given value is not nil.
func (iruo *InstanceRuntimeUpdateOne) SetNillableInstanceContext(s *string) *InstanceRuntimeUpdateOne {
	if s != nil {
		iruo.SetInstanceContext(*s)
	}
	return iruo
}

// ClearInstanceContext clears the value of the "instanceContext" field.
func (iruo *InstanceRuntimeUpdateOne) ClearInstanceContext() *InstanceRuntimeUpdateOne {
	iruo.mutation.ClearInstanceContext()
	return iruo
}

// SetStateContext sets the "stateContext" field.
func (iruo *InstanceRuntimeUpdateOne) SetStateContext(s string) *InstanceRuntimeUpdateOne {
	iruo.mutation.SetStateContext(s)
	return iruo
}

// SetNillableStateContext sets the "stateContext" field if the given value is not nil.
func (iruo *InstanceRuntimeUpdateOne) SetNillableStateContext(s *string) *InstanceRuntimeUpdateOne {
	if s != nil {
		iruo.SetStateContext(*s)
	}
	return iruo
}

// ClearStateContext clears the value of the "stateContext" field.
func (iruo *InstanceRuntimeUpdateOne) ClearStateContext() *InstanceRuntimeUpdateOne {
	iruo.mutation.ClearStateContext()
	return iruo
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (iruo *InstanceRuntimeUpdateOne) SetInstanceID(id uuid.UUID) *InstanceRuntimeUpdateOne {
	iruo.mutation.SetInstanceID(id)
	return iruo
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (iruo *InstanceRuntimeUpdateOne) SetNillableInstanceID(id *uuid.UUID) *InstanceRuntimeUpdateOne {
	if id != nil {
		iruo = iruo.SetInstanceID(*id)
	}
	return iruo
}

// SetInstance sets the "instance" edge to the Instance entity.
func (iruo *InstanceRuntimeUpdateOne) SetInstance(i *Instance) *InstanceRuntimeUpdateOne {
	return iruo.SetInstanceID(i.ID)
}

// SetCallerID sets the "caller" edge to the Instance entity by ID.
func (iruo *InstanceRuntimeUpdateOne) SetCallerID(id uuid.UUID) *InstanceRuntimeUpdateOne {
	iruo.mutation.SetCallerID(id)
	return iruo
}

// SetNillableCallerID sets the "caller" edge to the Instance entity by ID if the given value is not nil.
func (iruo *InstanceRuntimeUpdateOne) SetNillableCallerID(id *uuid.UUID) *InstanceRuntimeUpdateOne {
	if id != nil {
		iruo = iruo.SetCallerID(*id)
	}
	return iruo
}

// SetCaller sets the "caller" edge to the Instance entity.
func (iruo *InstanceRuntimeUpdateOne) SetCaller(i *Instance) *InstanceRuntimeUpdateOne {
	return iruo.SetCallerID(i.ID)
}

// Mutation returns the InstanceRuntimeMutation object of the builder.
func (iruo *InstanceRuntimeUpdateOne) Mutation() *InstanceRuntimeMutation {
	return iruo.mutation
}

// ClearInstance clears the "instance" edge to the Instance entity.
func (iruo *InstanceRuntimeUpdateOne) ClearInstance() *InstanceRuntimeUpdateOne {
	iruo.mutation.ClearInstance()
	return iruo
}

// ClearCaller clears the "caller" edge to the Instance entity.
func (iruo *InstanceRuntimeUpdateOne) ClearCaller() *InstanceRuntimeUpdateOne {
	iruo.mutation.ClearCaller()
	return iruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iruo *InstanceRuntimeUpdateOne) Select(field string, fields ...string) *InstanceRuntimeUpdateOne {
	iruo.fields = append([]string{field}, fields...)
	return iruo
}

// Save executes the query and returns the updated InstanceRuntime entity.
func (iruo *InstanceRuntimeUpdateOne) Save(ctx context.Context) (*InstanceRuntime, error) {
	var (
		err  error
		node *InstanceRuntime
	)
	if len(iruo.hooks) == 0 {
		node, err = iruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstanceRuntimeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iruo.mutation = mutation
			node, err = iruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iruo.hooks) - 1; i >= 0; i-- {
			if iruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iruo *InstanceRuntimeUpdateOne) SaveX(ctx context.Context) *InstanceRuntime {
	node, err := iruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iruo *InstanceRuntimeUpdateOne) Exec(ctx context.Context) error {
	_, err := iruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iruo *InstanceRuntimeUpdateOne) ExecX(ctx context.Context) {
	if err := iruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iruo *InstanceRuntimeUpdateOne) sqlSave(ctx context.Context) (_node *InstanceRuntime, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instanceruntime.Table,
			Columns: instanceruntime.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: instanceruntime.FieldID,
			},
		},
	}
	id, ok := iruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing InstanceRuntime.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := iruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, instanceruntime.FieldID)
		for _, f := range fields {
			if !instanceruntime.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != instanceruntime.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iruo.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldData,
		})
	}
	if value, ok := iruo.mutation.Controller(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldController,
		})
	}
	if iruo.mutation.ControllerCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instanceruntime.FieldController,
		})
	}
	if value, ok := iruo.mutation.Memory(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldMemory,
		})
	}
	if iruo.mutation.MemoryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instanceruntime.FieldMemory,
		})
	}
	if value, ok := iruo.mutation.Flow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: instanceruntime.FieldFlow,
		})
	}
	if iruo.mutation.FlowCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: instanceruntime.FieldFlow,
		})
	}
	if value, ok := iruo.mutation.Output(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldOutput,
		})
	}
	if iruo.mutation.OutputCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instanceruntime.FieldOutput,
		})
	}
	if value, ok := iruo.mutation.StateBeginTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instanceruntime.FieldStateBeginTime,
		})
	}
	if iruo.mutation.StateBeginTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: instanceruntime.FieldStateBeginTime,
		})
	}
	if value, ok := iruo.mutation.Deadline(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instanceruntime.FieldDeadline,
		})
	}
	if iruo.mutation.DeadlineCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: instanceruntime.FieldDeadline,
		})
	}
	if value, ok := iruo.mutation.Attempts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: instanceruntime.FieldAttempts,
		})
	}
	if value, ok := iruo.mutation.AddedAttempts(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: instanceruntime.FieldAttempts,
		})
	}
	if iruo.mutation.AttemptsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: instanceruntime.FieldAttempts,
		})
	}
	if value, ok := iruo.mutation.CallerData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldCallerData,
		})
	}
	if iruo.mutation.CallerDataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instanceruntime.FieldCallerData,
		})
	}
	if value, ok := iruo.mutation.InstanceContext(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldInstanceContext,
		})
	}
	if iruo.mutation.InstanceContextCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instanceruntime.FieldInstanceContext,
		})
	}
	if value, ok := iruo.mutation.StateContext(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceruntime.FieldStateContext,
		})
	}
	if iruo.mutation.StateContextCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instanceruntime.FieldStateContext,
		})
	}
	if iruo.mutation.InstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   instanceruntime.InstanceTable,
			Columns: []string{instanceruntime.InstanceColumn},
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
	if nodes := iruo.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   instanceruntime.InstanceTable,
			Columns: []string{instanceruntime.InstanceColumn},
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
	if iruo.mutation.CallerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instanceruntime.CallerTable,
			Columns: []string{instanceruntime.CallerColumn},
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
	if nodes := iruo.mutation.CallerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instanceruntime.CallerTable,
			Columns: []string{instanceruntime.CallerColumn},
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
	_node = &InstanceRuntime{config: iruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{instanceruntime.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
