// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/realotz/whole/internal/services/users/data/ent/employee"
	"github.com/realotz/whole/internal/services/users/data/ent/permission"
	"github.com/realotz/whole/internal/services/users/data/ent/predicate"
	"github.com/realotz/whole/internal/services/users/data/ent/role"
)

// RoleUpdate is the builder for updating Role entities.
type RoleUpdate struct {
	config
	hooks    []Hook
	mutation *RoleMutation
}

// Where adds a new predicate for the RoleUpdate builder.
func (ru *RoleUpdate) Where(ps ...predicate.Role) *RoleUpdate {
	ru.mutation.predicates = append(ru.mutation.predicates, ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RoleUpdate) SetName(s string) *RoleUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetDesc sets the "desc" field.
func (ru *RoleUpdate) SetDesc(s string) *RoleUpdate {
	ru.mutation.SetDesc(s)
	return ru
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableDesc(s *string) *RoleUpdate {
	if s != nil {
		ru.SetDesc(*s)
	}
	return ru
}

// ClearDesc clears the value of the "desc" field.
func (ru *RoleUpdate) ClearDesc() *RoleUpdate {
	ru.mutation.ClearDesc()
	return ru
}

// SetStatus sets the "status" field.
func (ru *RoleUpdate) SetStatus(r role.Status) *RoleUpdate {
	ru.mutation.SetStatus(r)
	return ru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableStatus(r *role.Status) *RoleUpdate {
	if r != nil {
		ru.SetStatus(*r)
	}
	return ru
}

// ClearStatus clears the value of the "status" field.
func (ru *RoleUpdate) ClearStatus() *RoleUpdate {
	ru.mutation.ClearStatus()
	return ru
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (ru *RoleUpdate) AddPermissionIDs(ids ...int64) *RoleUpdate {
	ru.mutation.AddPermissionIDs(ids...)
	return ru
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (ru *RoleUpdate) AddPermissions(p ...*Permission) *RoleUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddPermissionIDs(ids...)
}

// AddEmployeeIDs adds the "employees" edge to the Employee entity by IDs.
func (ru *RoleUpdate) AddEmployeeIDs(ids ...int64) *RoleUpdate {
	ru.mutation.AddEmployeeIDs(ids...)
	return ru
}

// AddEmployees adds the "employees" edges to the Employee entity.
func (ru *RoleUpdate) AddEmployees(e ...*Employee) *RoleUpdate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ru.AddEmployeeIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ru *RoleUpdate) Mutation() *RoleMutation {
	return ru.mutation
}

// ClearPermissions clears all "permissions" edges to the Permission entity.
func (ru *RoleUpdate) ClearPermissions() *RoleUpdate {
	ru.mutation.ClearPermissions()
	return ru
}

// RemovePermissionIDs removes the "permissions" edge to Permission entities by IDs.
func (ru *RoleUpdate) RemovePermissionIDs(ids ...int64) *RoleUpdate {
	ru.mutation.RemovePermissionIDs(ids...)
	return ru
}

// RemovePermissions removes "permissions" edges to Permission entities.
func (ru *RoleUpdate) RemovePermissions(p ...*Permission) *RoleUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemovePermissionIDs(ids...)
}

// ClearEmployees clears all "employees" edges to the Employee entity.
func (ru *RoleUpdate) ClearEmployees() *RoleUpdate {
	ru.mutation.ClearEmployees()
	return ru
}

// RemoveEmployeeIDs removes the "employees" edge to Employee entities by IDs.
func (ru *RoleUpdate) RemoveEmployeeIDs(ids ...int64) *RoleUpdate {
	ru.mutation.RemoveEmployeeIDs(ids...)
	return ru
}

// RemoveEmployees removes "employees" edges to Employee entities.
func (ru *RoleUpdate) RemoveEmployees(e ...*Employee) *RoleUpdate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ru.RemoveEmployeeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoleUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoleUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoleUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RoleUpdate) defaults() {
	if _, ok := ru.mutation.UpdateTime(); !ok {
		v := role.UpdateDefaultUpdateTime()
		ru.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RoleUpdate) check() error {
	if v, ok := ru.mutation.Name(); ok {
		if err := role.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := ru.mutation.Status(); ok {
		if err := role.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (ru *RoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   role.Table,
			Columns: role.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: role.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: role.FieldUpdateTime,
		})
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldName,
		})
	}
	if value, ok := ru.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldDesc,
		})
	}
	if ru.mutation.DescCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: role.FieldDesc,
		})
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: role.FieldStatus,
		})
	}
	if ru.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: role.FieldStatus,
		})
	}
	if ru.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: permission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !ru.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: permission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: permission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.EmployeesTable,
			Columns: role.EmployeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: employee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedEmployeesIDs(); len(nodes) > 0 && !ru.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.EmployeesTable,
			Columns: role.EmployeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.EmployeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.EmployeesTable,
			Columns: role.EmployeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RoleUpdateOne is the builder for updating a single Role entity.
type RoleUpdateOne struct {
	config
	hooks    []Hook
	mutation *RoleMutation
}

// SetName sets the "name" field.
func (ruo *RoleUpdateOne) SetName(s string) *RoleUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetDesc sets the "desc" field.
func (ruo *RoleUpdateOne) SetDesc(s string) *RoleUpdateOne {
	ruo.mutation.SetDesc(s)
	return ruo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableDesc(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetDesc(*s)
	}
	return ruo
}

// ClearDesc clears the value of the "desc" field.
func (ruo *RoleUpdateOne) ClearDesc() *RoleUpdateOne {
	ruo.mutation.ClearDesc()
	return ruo
}

// SetStatus sets the "status" field.
func (ruo *RoleUpdateOne) SetStatus(r role.Status) *RoleUpdateOne {
	ruo.mutation.SetStatus(r)
	return ruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableStatus(r *role.Status) *RoleUpdateOne {
	if r != nil {
		ruo.SetStatus(*r)
	}
	return ruo
}

// ClearStatus clears the value of the "status" field.
func (ruo *RoleUpdateOne) ClearStatus() *RoleUpdateOne {
	ruo.mutation.ClearStatus()
	return ruo
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (ruo *RoleUpdateOne) AddPermissionIDs(ids ...int64) *RoleUpdateOne {
	ruo.mutation.AddPermissionIDs(ids...)
	return ruo
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (ruo *RoleUpdateOne) AddPermissions(p ...*Permission) *RoleUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddPermissionIDs(ids...)
}

// AddEmployeeIDs adds the "employees" edge to the Employee entity by IDs.
func (ruo *RoleUpdateOne) AddEmployeeIDs(ids ...int64) *RoleUpdateOne {
	ruo.mutation.AddEmployeeIDs(ids...)
	return ruo
}

// AddEmployees adds the "employees" edges to the Employee entity.
func (ruo *RoleUpdateOne) AddEmployees(e ...*Employee) *RoleUpdateOne {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ruo.AddEmployeeIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ruo *RoleUpdateOne) Mutation() *RoleMutation {
	return ruo.mutation
}

// ClearPermissions clears all "permissions" edges to the Permission entity.
func (ruo *RoleUpdateOne) ClearPermissions() *RoleUpdateOne {
	ruo.mutation.ClearPermissions()
	return ruo
}

// RemovePermissionIDs removes the "permissions" edge to Permission entities by IDs.
func (ruo *RoleUpdateOne) RemovePermissionIDs(ids ...int64) *RoleUpdateOne {
	ruo.mutation.RemovePermissionIDs(ids...)
	return ruo
}

// RemovePermissions removes "permissions" edges to Permission entities.
func (ruo *RoleUpdateOne) RemovePermissions(p ...*Permission) *RoleUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemovePermissionIDs(ids...)
}

// ClearEmployees clears all "employees" edges to the Employee entity.
func (ruo *RoleUpdateOne) ClearEmployees() *RoleUpdateOne {
	ruo.mutation.ClearEmployees()
	return ruo
}

// RemoveEmployeeIDs removes the "employees" edge to Employee entities by IDs.
func (ruo *RoleUpdateOne) RemoveEmployeeIDs(ids ...int64) *RoleUpdateOne {
	ruo.mutation.RemoveEmployeeIDs(ids...)
	return ruo
}

// RemoveEmployees removes "employees" edges to Employee entities.
func (ruo *RoleUpdateOne) RemoveEmployees(e ...*Employee) *RoleUpdateOne {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ruo.RemoveEmployeeIDs(ids...)
}

// Save executes the query and returns the updated Role entity.
func (ruo *RoleUpdateOne) Save(ctx context.Context) (*Role, error) {
	var (
		err  error
		node *Role
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoleUpdateOne) SaveX(ctx context.Context) *Role {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoleUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoleUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RoleUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdateTime(); !ok {
		v := role.UpdateDefaultUpdateTime()
		ruo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RoleUpdateOne) check() error {
	if v, ok := ruo.mutation.Name(); ok {
		if err := role.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.Status(); ok {
		if err := role.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (ruo *RoleUpdateOne) sqlSave(ctx context.Context) (_node *Role, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   role.Table,
			Columns: role.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: role.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Role.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: role.FieldUpdateTime,
		})
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldName,
		})
	}
	if value, ok := ruo.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldDesc,
		})
	}
	if ruo.mutation.DescCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: role.FieldDesc,
		})
	}
	if value, ok := ruo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: role.FieldStatus,
		})
	}
	if ruo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: role.FieldStatus,
		})
	}
	if ruo.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: permission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !ruo.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: permission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: permission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.EmployeesTable,
			Columns: role.EmployeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: employee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedEmployeesIDs(); len(nodes) > 0 && !ruo.mutation.EmployeesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.EmployeesTable,
			Columns: role.EmployeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.EmployeesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.EmployeesTable,
			Columns: role.EmployeesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Role{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}