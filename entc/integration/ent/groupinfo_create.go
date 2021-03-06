// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/group"
	"entgo.io/ent/entc/integration/ent/groupinfo"
	"entgo.io/ent/schema/field"
)

// GroupInfoCreate is the builder for creating a GroupInfo entity.
type GroupInfoCreate struct {
	config
	mutation *GroupInfoMutation
	hooks    []Hook
}

// SetDesc sets the "desc" field.
func (gic *GroupInfoCreate) SetDesc(s string) *GroupInfoCreate {
	gic.mutation.SetDesc(s)
	return gic
}

// SetMaxUsers sets the "max_users" field.
func (gic *GroupInfoCreate) SetMaxUsers(i int) *GroupInfoCreate {
	gic.mutation.SetMaxUsers(i)
	return gic
}

// SetNillableMaxUsers sets the "max_users" field if the given value is not nil.
func (gic *GroupInfoCreate) SetNillableMaxUsers(i *int) *GroupInfoCreate {
	if i != nil {
		gic.SetMaxUsers(*i)
	}
	return gic
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (gic *GroupInfoCreate) AddGroupIDs(ids ...int) *GroupInfoCreate {
	gic.mutation.AddGroupIDs(ids...)
	return gic
}

// AddGroups adds the "groups" edges to the Group entity.
func (gic *GroupInfoCreate) AddGroups(g ...*Group) *GroupInfoCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gic.AddGroupIDs(ids...)
}

// Mutation returns the GroupInfoMutation object of the builder.
func (gic *GroupInfoCreate) Mutation() *GroupInfoMutation {
	return gic.mutation
}

// Save creates the GroupInfo in the database.
func (gic *GroupInfoCreate) Save(ctx context.Context) (*GroupInfo, error) {
	var (
		err  error
		node *GroupInfo
	)
	gic.defaults()
	if len(gic.hooks) == 0 {
		if err = gic.check(); err != nil {
			return nil, err
		}
		node, err = gic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gic.check(); err != nil {
				return nil, err
			}
			gic.mutation = mutation
			node, err = gic.sqlSave(ctx)
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gic.hooks) - 1; i >= 0; i-- {
			mut = gic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gic *GroupInfoCreate) SaveX(ctx context.Context) *GroupInfo {
	v, err := gic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (gic *GroupInfoCreate) defaults() {
	if _, ok := gic.mutation.MaxUsers(); !ok {
		v := groupinfo.DefaultMaxUsers
		gic.mutation.SetMaxUsers(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gic *GroupInfoCreate) check() error {
	if _, ok := gic.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New("ent: missing required field \"desc\"")}
	}
	if _, ok := gic.mutation.MaxUsers(); !ok {
		return &ValidationError{Name: "max_users", err: errors.New("ent: missing required field \"max_users\"")}
	}
	return nil
}

func (gic *GroupInfoCreate) sqlSave(ctx context.Context) (*GroupInfo, error) {
	_node, _spec := gic.createSpec()
	if err := sqlgraph.CreateNode(ctx, gic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gic *GroupInfoCreate) createSpec() (*GroupInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupInfo{config: gic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: groupinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupinfo.FieldID,
			},
		}
	)
	if value, ok := gic.mutation.Desc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupinfo.FieldDesc,
		})
		_node.Desc = value
	}
	if value, ok := gic.mutation.MaxUsers(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupinfo.FieldMaxUsers,
		})
		_node.MaxUsers = value
	}
	if nodes := gic.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   groupinfo.GroupsTable,
			Columns: []string{groupinfo.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
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

// GroupInfoCreateBulk is the builder for creating many GroupInfo entities in bulk.
type GroupInfoCreateBulk struct {
	config
	builders []*GroupInfoCreate
}

// Save creates the GroupInfo entities in the database.
func (gicb *GroupInfoCreateBulk) Save(ctx context.Context) ([]*GroupInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gicb.builders))
	nodes := make([]*GroupInfo, len(gicb.builders))
	mutators := make([]Mutator, len(gicb.builders))
	for i := range gicb.builders {
		func(i int, root context.Context) {
			builder := gicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, gicb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gicb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gicb *GroupInfoCreateBulk) SaveX(ctx context.Context) []*GroupInfo {
	v, err := gicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
