// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/realotz/whole/internal/services/users/data/ent/customer"
	"github.com/realotz/whole/internal/services/users/data/ent/predicate"
	"github.com/realotz/whole/internal/services/users/data/ent/wechat"
)

// WechatUpdate is the builder for updating Wechat entities.
type WechatUpdate struct {
	config
	hooks    []Hook
	mutation *WechatMutation
}

// Where adds a new predicate for the WechatUpdate builder.
func (wu *WechatUpdate) Where(ps ...predicate.Wechat) *WechatUpdate {
	wu.mutation.predicates = append(wu.mutation.predicates, ps...)
	return wu
}

// SetOpenid sets the "openid" field.
func (wu *WechatUpdate) SetOpenid(s string) *WechatUpdate {
	wu.mutation.SetOpenid(s)
	return wu
}

// SetUnionId sets the "unionId" field.
func (wu *WechatUpdate) SetUnionId(s string) *WechatUpdate {
	wu.mutation.SetUnionId(s)
	return wu
}

// SetAppType sets the "app_type" field.
func (wu *WechatUpdate) SetAppType(wt wechat.AppType) *WechatUpdate {
	wu.mutation.SetAppType(wt)
	return wu
}

// SetMetaData sets the "meta_data" field.
func (wu *WechatUpdate) SetMetaData(b []byte) *WechatUpdate {
	wu.mutation.SetMetaData(b)
	return wu
}

// ClearMetaData clears the value of the "meta_data" field.
func (wu *WechatUpdate) ClearMetaData() *WechatUpdate {
	wu.mutation.ClearMetaData()
	return wu
}

// SetCustomersID sets the "customers" edge to the Customer entity by ID.
func (wu *WechatUpdate) SetCustomersID(id int64) *WechatUpdate {
	wu.mutation.SetCustomersID(id)
	return wu
}

// SetNillableCustomersID sets the "customers" edge to the Customer entity by ID if the given value is not nil.
func (wu *WechatUpdate) SetNillableCustomersID(id *int64) *WechatUpdate {
	if id != nil {
		wu = wu.SetCustomersID(*id)
	}
	return wu
}

// SetCustomers sets the "customers" edge to the Customer entity.
func (wu *WechatUpdate) SetCustomers(c *Customer) *WechatUpdate {
	return wu.SetCustomersID(c.ID)
}

// Mutation returns the WechatMutation object of the builder.
func (wu *WechatUpdate) Mutation() *WechatMutation {
	return wu.mutation
}

// ClearCustomers clears the "customers" edge to the Customer entity.
func (wu *WechatUpdate) ClearCustomers() *WechatUpdate {
	wu.mutation.ClearCustomers()
	return wu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WechatUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	wu.defaults()
	if len(wu.hooks) == 0 {
		if err = wu.check(); err != nil {
			return 0, err
		}
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WechatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wu.check(); err != nil {
				return 0, err
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WechatUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WechatUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WechatUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wu *WechatUpdate) defaults() {
	if _, ok := wu.mutation.UpdateTime(); !ok {
		v := wechat.UpdateDefaultUpdateTime()
		wu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WechatUpdate) check() error {
	if v, ok := wu.mutation.Openid(); ok {
		if err := wechat.OpenidValidator(v); err != nil {
			return &ValidationError{Name: "openid", err: fmt.Errorf("ent: validator failed for field \"openid\": %w", err)}
		}
	}
	if v, ok := wu.mutation.UnionId(); ok {
		if err := wechat.UnionIdValidator(v); err != nil {
			return &ValidationError{Name: "unionId", err: fmt.Errorf("ent: validator failed for field \"unionId\": %w", err)}
		}
	}
	if v, ok := wu.mutation.AppType(); ok {
		if err := wechat.AppTypeValidator(v); err != nil {
			return &ValidationError{Name: "app_type", err: fmt.Errorf("ent: validator failed for field \"app_type\": %w", err)}
		}
	}
	return nil
}

func (wu *WechatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wechat.Table,
			Columns: wechat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: wechat.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: wechat.FieldUpdateTime,
		})
	}
	if value, ok := wu.mutation.Openid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wechat.FieldOpenid,
		})
	}
	if value, ok := wu.mutation.UnionId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wechat.FieldUnionId,
		})
	}
	if value, ok := wu.mutation.AppType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: wechat.FieldAppType,
		})
	}
	if value, ok := wu.mutation.MetaData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: wechat.FieldMetaData,
		})
	}
	if wu.mutation.MetaDataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: wechat.FieldMetaData,
		})
	}
	if wu.mutation.CustomersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wechat.CustomersTable,
			Columns: []string{wechat.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.CustomersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wechat.CustomersTable,
			Columns: []string{wechat.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wechat.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// WechatUpdateOne is the builder for updating a single Wechat entity.
type WechatUpdateOne struct {
	config
	hooks    []Hook
	mutation *WechatMutation
}

// SetOpenid sets the "openid" field.
func (wuo *WechatUpdateOne) SetOpenid(s string) *WechatUpdateOne {
	wuo.mutation.SetOpenid(s)
	return wuo
}

// SetUnionId sets the "unionId" field.
func (wuo *WechatUpdateOne) SetUnionId(s string) *WechatUpdateOne {
	wuo.mutation.SetUnionId(s)
	return wuo
}

// SetAppType sets the "app_type" field.
func (wuo *WechatUpdateOne) SetAppType(wt wechat.AppType) *WechatUpdateOne {
	wuo.mutation.SetAppType(wt)
	return wuo
}

// SetMetaData sets the "meta_data" field.
func (wuo *WechatUpdateOne) SetMetaData(b []byte) *WechatUpdateOne {
	wuo.mutation.SetMetaData(b)
	return wuo
}

// ClearMetaData clears the value of the "meta_data" field.
func (wuo *WechatUpdateOne) ClearMetaData() *WechatUpdateOne {
	wuo.mutation.ClearMetaData()
	return wuo
}

// SetCustomersID sets the "customers" edge to the Customer entity by ID.
func (wuo *WechatUpdateOne) SetCustomersID(id int64) *WechatUpdateOne {
	wuo.mutation.SetCustomersID(id)
	return wuo
}

// SetNillableCustomersID sets the "customers" edge to the Customer entity by ID if the given value is not nil.
func (wuo *WechatUpdateOne) SetNillableCustomersID(id *int64) *WechatUpdateOne {
	if id != nil {
		wuo = wuo.SetCustomersID(*id)
	}
	return wuo
}

// SetCustomers sets the "customers" edge to the Customer entity.
func (wuo *WechatUpdateOne) SetCustomers(c *Customer) *WechatUpdateOne {
	return wuo.SetCustomersID(c.ID)
}

// Mutation returns the WechatMutation object of the builder.
func (wuo *WechatUpdateOne) Mutation() *WechatMutation {
	return wuo.mutation
}

// ClearCustomers clears the "customers" edge to the Customer entity.
func (wuo *WechatUpdateOne) ClearCustomers() *WechatUpdateOne {
	wuo.mutation.ClearCustomers()
	return wuo
}

// Save executes the query and returns the updated Wechat entity.
func (wuo *WechatUpdateOne) Save(ctx context.Context) (*Wechat, error) {
	var (
		err  error
		node *Wechat
	)
	wuo.defaults()
	if len(wuo.hooks) == 0 {
		if err = wuo.check(); err != nil {
			return nil, err
		}
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WechatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wuo.check(); err != nil {
				return nil, err
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WechatUpdateOne) SaveX(ctx context.Context) *Wechat {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WechatUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WechatUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuo *WechatUpdateOne) defaults() {
	if _, ok := wuo.mutation.UpdateTime(); !ok {
		v := wechat.UpdateDefaultUpdateTime()
		wuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WechatUpdateOne) check() error {
	if v, ok := wuo.mutation.Openid(); ok {
		if err := wechat.OpenidValidator(v); err != nil {
			return &ValidationError{Name: "openid", err: fmt.Errorf("ent: validator failed for field \"openid\": %w", err)}
		}
	}
	if v, ok := wuo.mutation.UnionId(); ok {
		if err := wechat.UnionIdValidator(v); err != nil {
			return &ValidationError{Name: "unionId", err: fmt.Errorf("ent: validator failed for field \"unionId\": %w", err)}
		}
	}
	if v, ok := wuo.mutation.AppType(); ok {
		if err := wechat.AppTypeValidator(v); err != nil {
			return &ValidationError{Name: "app_type", err: fmt.Errorf("ent: validator failed for field \"app_type\": %w", err)}
		}
	}
	return nil
}

func (wuo *WechatUpdateOne) sqlSave(ctx context.Context) (_node *Wechat, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wechat.Table,
			Columns: wechat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: wechat.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Wechat.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: wechat.FieldUpdateTime,
		})
	}
	if value, ok := wuo.mutation.Openid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wechat.FieldOpenid,
		})
	}
	if value, ok := wuo.mutation.UnionId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wechat.FieldUnionId,
		})
	}
	if value, ok := wuo.mutation.AppType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: wechat.FieldAppType,
		})
	}
	if value, ok := wuo.mutation.MetaData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: wechat.FieldMetaData,
		})
	}
	if wuo.mutation.MetaDataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: wechat.FieldMetaData,
		})
	}
	if wuo.mutation.CustomersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wechat.CustomersTable,
			Columns: []string{wechat.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.CustomersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wechat.CustomersTable,
			Columns: []string{wechat.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Wechat{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wechat.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}