// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/ent/namespace"
	"github.com/direktiv/direktiv/ent/workflow"
	"github.com/google/uuid"
)

// NamespaceCreate is the builder for creating a Namespace entity.
type NamespaceCreate struct {
	config
	mutation *NamespaceMutation
	hooks    []Hook
}

// SetCreated sets the "created" field.
func (nc *NamespaceCreate) SetCreated(t time.Time) *NamespaceCreate {
	nc.mutation.SetCreated(t)
	return nc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (nc *NamespaceCreate) SetNillableCreated(t *time.Time) *NamespaceCreate {
	if t != nil {
		nc.SetCreated(*t)
	}
	return nc
}

// SetID sets the "id" field.
func (nc *NamespaceCreate) SetID(s string) *NamespaceCreate {
	nc.mutation.SetID(s)
	return nc
}

// AddWorkflowIDs adds the "workflows" edge to the Workflow entity by IDs.
func (nc *NamespaceCreate) AddWorkflowIDs(ids ...uuid.UUID) *NamespaceCreate {
	nc.mutation.AddWorkflowIDs(ids...)
	return nc
}

// AddWorkflows adds the "workflows" edges to the Workflow entity.
func (nc *NamespaceCreate) AddWorkflows(w ...*Workflow) *NamespaceCreate {
	ids := make([]uuid.UUID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return nc.AddWorkflowIDs(ids...)
}

// Mutation returns the NamespaceMutation object of the builder.
func (nc *NamespaceCreate) Mutation() *NamespaceMutation {
	return nc.mutation
}

// Save creates the Namespace in the database.
func (nc *NamespaceCreate) Save(ctx context.Context) (*Namespace, error) {
	var (
		err  error
		node *Namespace
	)
	nc.defaults()
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NamespaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			if node, err = nc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			if nc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NamespaceCreate) SaveX(ctx context.Context) *Namespace {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NamespaceCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NamespaceCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NamespaceCreate) defaults() {
	if _, ok := nc.mutation.Created(); !ok {
		v := namespace.DefaultCreated()
		nc.mutation.SetCreated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NamespaceCreate) check() error {
	if _, ok := nc.mutation.Created(); !ok {
		return &ValidationError{Name: "created", err: errors.New(`ent: missing required field "created"`)}
	}
	if v, ok := nc.mutation.ID(); ok {
		if err := namespace.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "id": %w`, err)}
		}
	}
	return nil
}

func (nc *NamespaceCreate) sqlSave(ctx context.Context) (*Namespace, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(string)
	}
	return _node, nil
}

func (nc *NamespaceCreate) createSpec() (*Namespace, *sqlgraph.CreateSpec) {
	var (
		_node = &Namespace{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: namespace.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: namespace.FieldID,
			},
		}
	)
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := nc.mutation.Created(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: namespace.FieldCreated,
		})
		_node.Created = value
	}
	if nodes := nc.mutation.WorkflowsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.WorkflowsTable,
			Columns: []string{namespace.WorkflowsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NamespaceCreateBulk is the builder for creating many Namespace entities in bulk.
type NamespaceCreateBulk struct {
	config
	builders []*NamespaceCreate
}

// Save creates the Namespace entities in the database.
func (ncb *NamespaceCreateBulk) Save(ctx context.Context) ([]*Namespace, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Namespace, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NamespaceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NamespaceCreateBulk) SaveX(ctx context.Context) []*Namespace {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NamespaceCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NamespaceCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}
