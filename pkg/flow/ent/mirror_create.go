// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/inode"
	"github.com/direktiv/direktiv/pkg/flow/ent/mirror"
	"github.com/direktiv/direktiv/pkg/flow/ent/mirroractivity"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/google/uuid"
)

// MirrorCreate is the builder for creating a Mirror entity.
type MirrorCreate struct {
	config
	mutation *MirrorMutation
	hooks    []Hook
}

// SetURL sets the "url" field.
func (mc *MirrorCreate) SetURL(s string) *MirrorCreate {
	mc.mutation.SetURL(s)
	return mc
}

// SetRef sets the "ref" field.
func (mc *MirrorCreate) SetRef(s string) *MirrorCreate {
	mc.mutation.SetRef(s)
	return mc
}

// SetCron sets the "cron" field.
func (mc *MirrorCreate) SetCron(s string) *MirrorCreate {
	mc.mutation.SetCron(s)
	return mc
}

// SetPublicKey sets the "public_key" field.
func (mc *MirrorCreate) SetPublicKey(s string) *MirrorCreate {
	mc.mutation.SetPublicKey(s)
	return mc
}

// SetPrivateKey sets the "private_key" field.
func (mc *MirrorCreate) SetPrivateKey(s string) *MirrorCreate {
	mc.mutation.SetPrivateKey(s)
	return mc
}

// SetPassphrase sets the "passphrase" field.
func (mc *MirrorCreate) SetPassphrase(s string) *MirrorCreate {
	mc.mutation.SetPassphrase(s)
	return mc
}

// SetCommit sets the "commit" field.
func (mc *MirrorCreate) SetCommit(s string) *MirrorCreate {
	mc.mutation.SetCommit(s)
	return mc
}

// SetLastSync sets the "last_sync" field.
func (mc *MirrorCreate) SetLastSync(t time.Time) *MirrorCreate {
	mc.mutation.SetLastSync(t)
	return mc
}

// SetNillableLastSync sets the "last_sync" field if the given value is not nil.
func (mc *MirrorCreate) SetNillableLastSync(t *time.Time) *MirrorCreate {
	if t != nil {
		mc.SetLastSync(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MirrorCreate) SetUpdatedAt(t time.Time) *MirrorCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MirrorCreate) SetNillableUpdatedAt(t *time.Time) *MirrorCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MirrorCreate) SetID(u uuid.UUID) *MirrorCreate {
	mc.mutation.SetID(u)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MirrorCreate) SetNillableID(u *uuid.UUID) *MirrorCreate {
	if u != nil {
		mc.SetID(*u)
	}
	return mc
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (mc *MirrorCreate) SetNamespaceID(id uuid.UUID) *MirrorCreate {
	mc.mutation.SetNamespaceID(id)
	return mc
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (mc *MirrorCreate) SetNamespace(n *Namespace) *MirrorCreate {
	return mc.SetNamespaceID(n.ID)
}

// SetInodeID sets the "inode" edge to the Inode entity by ID.
func (mc *MirrorCreate) SetInodeID(id uuid.UUID) *MirrorCreate {
	mc.mutation.SetInodeID(id)
	return mc
}

// SetNillableInodeID sets the "inode" edge to the Inode entity by ID if the given value is not nil.
func (mc *MirrorCreate) SetNillableInodeID(id *uuid.UUID) *MirrorCreate {
	if id != nil {
		mc = mc.SetInodeID(*id)
	}
	return mc
}

// SetInode sets the "inode" edge to the Inode entity.
func (mc *MirrorCreate) SetInode(i *Inode) *MirrorCreate {
	return mc.SetInodeID(i.ID)
}

// AddActivityIDs adds the "activities" edge to the MirrorActivity entity by IDs.
func (mc *MirrorCreate) AddActivityIDs(ids ...uuid.UUID) *MirrorCreate {
	mc.mutation.AddActivityIDs(ids...)
	return mc
}

// AddActivities adds the "activities" edges to the MirrorActivity entity.
func (mc *MirrorCreate) AddActivities(m ...*MirrorActivity) *MirrorCreate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddActivityIDs(ids...)
}

// Mutation returns the MirrorMutation object of the builder.
func (mc *MirrorCreate) Mutation() *MirrorMutation {
	return mc.mutation
}

// Save creates the Mirror in the database.
func (mc *MirrorCreate) Save(ctx context.Context) (*Mirror, error) {
	var (
		err  error
		node *Mirror
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MirrorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MirrorCreate) SaveX(ctx context.Context) *Mirror {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MirrorCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MirrorCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MirrorCreate) defaults() {
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := mirror.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := mirror.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MirrorCreate) check() error {
	if _, ok := mc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Mirror.url"`)}
	}
	if _, ok := mc.mutation.Ref(); !ok {
		return &ValidationError{Name: "ref", err: errors.New(`ent: missing required field "Mirror.ref"`)}
	}
	if _, ok := mc.mutation.Cron(); !ok {
		return &ValidationError{Name: "cron", err: errors.New(`ent: missing required field "Mirror.cron"`)}
	}
	if _, ok := mc.mutation.PublicKey(); !ok {
		return &ValidationError{Name: "public_key", err: errors.New(`ent: missing required field "Mirror.public_key"`)}
	}
	if _, ok := mc.mutation.PrivateKey(); !ok {
		return &ValidationError{Name: "private_key", err: errors.New(`ent: missing required field "Mirror.private_key"`)}
	}
	if _, ok := mc.mutation.Passphrase(); !ok {
		return &ValidationError{Name: "passphrase", err: errors.New(`ent: missing required field "Mirror.passphrase"`)}
	}
	if _, ok := mc.mutation.Commit(); !ok {
		return &ValidationError{Name: "commit", err: errors.New(`ent: missing required field "Mirror.commit"`)}
	}
	if _, ok := mc.mutation.NamespaceID(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`ent: missing required edge "Mirror.namespace"`)}
	}
	return nil
}

func (mc *MirrorCreate) sqlSave(ctx context.Context) (*Mirror, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (mc *MirrorCreate) createSpec() (*Mirror, *sqlgraph.CreateSpec) {
	var (
		_node = &Mirror{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mirror.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mirror.FieldID,
			},
		}
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := mc.mutation.Ref(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldRef,
		})
		_node.Ref = value
	}
	if value, ok := mc.mutation.Cron(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldCron,
		})
		_node.Cron = value
	}
	if value, ok := mc.mutation.PublicKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldPublicKey,
		})
		_node.PublicKey = value
	}
	if value, ok := mc.mutation.PrivateKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldPrivateKey,
		})
		_node.PrivateKey = value
	}
	if value, ok := mc.mutation.Passphrase(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldPassphrase,
		})
		_node.Passphrase = value
	}
	if value, ok := mc.mutation.Commit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mirror.FieldCommit,
		})
		_node.Commit = value
	}
	if value, ok := mc.mutation.LastSync(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mirror.FieldLastSync,
		})
		_node.LastSync = &value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mirror.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := mc.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mirror.NamespaceTable,
			Columns: []string{mirror.NamespaceColumn},
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
		_node.namespace_mirrors = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.InodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   mirror.InodeTable,
			Columns: []string{mirror.InodeColumn},
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
		_node.inode_mirror = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mirror.ActivitiesTable,
			Columns: []string{mirror.ActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: mirroractivity.FieldID,
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

// MirrorCreateBulk is the builder for creating many Mirror entities in bulk.
type MirrorCreateBulk struct {
	config
	builders []*MirrorCreate
}

// Save creates the Mirror entities in the database.
func (mcb *MirrorCreateBulk) Save(ctx context.Context) ([]*Mirror, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Mirror, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MirrorMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MirrorCreateBulk) SaveX(ctx context.Context) []*Mirror {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MirrorCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MirrorCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}