// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/realotz/whole/internal/services/users/data/ent/permission"
	"github.com/realotz/whole/internal/services/users/data/ent/role"
)

// PermissionCreate is the builder for creating a Permission entity.
type PermissionCreate struct {
	config
	mutation *PermissionMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *PermissionCreate) SetCreateTime(t time.Time) *PermissionCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableCreateTime(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *PermissionCreate) SetUpdateTime(t time.Time) *PermissionCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableUpdateTime(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PermissionCreate) SetName(s string) *PermissionCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetService sets the "service" field.
func (pc *PermissionCreate) SetService(s string) *PermissionCreate {
	pc.mutation.SetService(s)
	return pc
}

// SetPath sets the "path" field.
func (pc *PermissionCreate) SetPath(s string) *PermissionCreate {
	pc.mutation.SetPath(s)
	return pc
}

// SetAction sets the "action" field.
func (pc *PermissionCreate) SetAction(s string) *PermissionCreate {
	pc.mutation.SetAction(s)
	return pc
}

// SetDesc sets the "desc" field.
func (pc *PermissionCreate) SetDesc(s string) *PermissionCreate {
	pc.mutation.SetDesc(s)
	return pc
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableDesc(s *string) *PermissionCreate {
	if s != nil {
		pc.SetDesc(*s)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PermissionCreate) SetID(i int64) *PermissionCreate {
	pc.mutation.SetID(i)
	return pc
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (pc *PermissionCreate) AddRoleIDs(ids ...int64) *PermissionCreate {
	pc.mutation.AddRoleIDs(ids...)
	return pc
}

// AddRoles adds the "roles" edges to the Role entity.
func (pc *PermissionCreate) AddRoles(r ...*Role) *PermissionCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pc.AddRoleIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pc *PermissionCreate) Mutation() *PermissionMutation {
	return pc.mutation
}

// Save creates the Permission in the database.
func (pc *PermissionCreate) Save(ctx context.Context) (*Permission, error) {
	var (
		err  error
		node *Permission
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PermissionCreate) SaveX(ctx context.Context) *Permission {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pc *PermissionCreate) defaults() {
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := permission.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := permission.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PermissionCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := permission.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Service(); !ok {
		return &ValidationError{Name: "service", err: errors.New("ent: missing required field \"service\"")}
	}
	if v, ok := pc.mutation.Service(); ok {
		if err := permission.ServiceValidator(v); err != nil {
			return &ValidationError{Name: "service", err: fmt.Errorf("ent: validator failed for field \"service\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New("ent: missing required field \"path\"")}
	}
	if v, ok := pc.mutation.Path(); ok {
		if err := permission.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf("ent: validator failed for field \"path\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Action(); !ok {
		return &ValidationError{Name: "action", err: errors.New("ent: missing required field \"action\"")}
	}
	if v, ok := pc.mutation.Action(); ok {
		if err := permission.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf("ent: validator failed for field \"action\": %w", err)}
		}
	}
	return nil
}

func (pc *PermissionCreate) sqlSave(ctx context.Context) (*Permission, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (pc *PermissionCreate) createSpec() (*Permission, *sqlgraph.CreateSpec) {
	var (
		_node = &Permission{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: permission.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: permission.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.Service(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldService,
		})
		_node.Service = value
	}
	if value, ok := pc.mutation.Path(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldPath,
		})
		_node.Path = value
	}
	if value, ok := pc.mutation.Action(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldAction,
		})
		_node.Action = value
	}
	if value, ok := pc.mutation.Desc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldDesc,
		})
		_node.Desc = value
	}
	if nodes := pc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: role.FieldID,
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

// PermissionCreateBulk is the builder for creating many Permission entities in bulk.
type PermissionCreateBulk struct {
	config
	builders []*PermissionCreate
}

// Save creates the Permission entities in the database.
func (pcb *PermissionCreateBulk) Save(ctx context.Context) ([]*Permission, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Permission, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PermissionMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PermissionCreateBulk) SaveX(ctx context.Context) []*Permission {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
