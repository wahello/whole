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
	"github.com/realotz/whole/internal/services/users/data/ent/customer"
	"github.com/realotz/whole/internal/services/users/data/ent/predicate"
	"github.com/realotz/whole/internal/services/users/data/ent/wechat"
)

// CustomerUpdate is the builder for updating Customer entities.
type CustomerUpdate struct {
	config
	hooks    []Hook
	mutation *CustomerMutation
}

// Where adds a new predicate for the CustomerUpdate builder.
func (cu *CustomerUpdate) Where(ps ...predicate.Customer) *CustomerUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// SetSex sets the "sex" field.
func (cu *CustomerUpdate) SetSex(c customer.Sex) *CustomerUpdate {
	cu.mutation.SetSex(c)
	return cu
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableSex(c *customer.Sex) *CustomerUpdate {
	if c != nil {
		cu.SetSex(*c)
	}
	return cu
}

// SetUUID sets the "uuid" field.
func (cu *CustomerUpdate) SetUUID(u uuid.UUID) *CustomerUpdate {
	cu.mutation.SetUUID(u)
	return cu
}

// SetAccount sets the "account" field.
func (cu *CustomerUpdate) SetAccount(s string) *CustomerUpdate {
	cu.mutation.SetAccount(s)
	return cu
}

// SetAvatar sets the "avatar" field.
func (cu *CustomerUpdate) SetAvatar(s string) *CustomerUpdate {
	cu.mutation.SetAvatar(s)
	return cu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableAvatar(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetAvatar(*s)
	}
	return cu
}

// ClearAvatar clears the value of the "avatar" field.
func (cu *CustomerUpdate) ClearAvatar() *CustomerUpdate {
	cu.mutation.ClearAvatar()
	return cu
}

// SetName sets the "name" field.
func (cu *CustomerUpdate) SetName(s string) *CustomerUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableName(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// ClearName clears the value of the "name" field.
func (cu *CustomerUpdate) ClearName() *CustomerUpdate {
	cu.mutation.ClearName()
	return cu
}

// SetRole sets the "role" field.
func (cu *CustomerUpdate) SetRole(s string) *CustomerUpdate {
	cu.mutation.SetRole(s)
	return cu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableRole(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetRole(*s)
	}
	return cu
}

// ClearRole clears the value of the "role" field.
func (cu *CustomerUpdate) ClearRole() *CustomerUpdate {
	cu.mutation.ClearRole()
	return cu
}

// SetNickName sets the "nick_name" field.
func (cu *CustomerUpdate) SetNickName(s string) *CustomerUpdate {
	cu.mutation.SetNickName(s)
	return cu
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableNickName(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetNickName(*s)
	}
	return cu
}

// ClearNickName clears the value of the "nick_name" field.
func (cu *CustomerUpdate) ClearNickName() *CustomerUpdate {
	cu.mutation.ClearNickName()
	return cu
}

// SetEmail sets the "email" field.
func (cu *CustomerUpdate) SetEmail(s string) *CustomerUpdate {
	cu.mutation.SetEmail(s)
	return cu
}

// SetMobile sets the "mobile" field.
func (cu *CustomerUpdate) SetMobile(s string) *CustomerUpdate {
	cu.mutation.SetMobile(s)
	return cu
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableMobile(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetMobile(*s)
	}
	return cu
}

// ClearMobile clears the value of the "mobile" field.
func (cu *CustomerUpdate) ClearMobile() *CustomerUpdate {
	cu.mutation.ClearMobile()
	return cu
}

// SetIDCard sets the "id_card" field.
func (cu *CustomerUpdate) SetIDCard(s string) *CustomerUpdate {
	cu.mutation.SetIDCard(s)
	return cu
}

// SetNillableIDCard sets the "id_card" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableIDCard(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetIDCard(*s)
	}
	return cu
}

// ClearIDCard clears the value of the "id_card" field.
func (cu *CustomerUpdate) ClearIDCard() *CustomerUpdate {
	cu.mutation.ClearIDCard()
	return cu
}

// SetBirthday sets the "birthday" field.
func (cu *CustomerUpdate) SetBirthday(t time.Time) *CustomerUpdate {
	cu.mutation.SetBirthday(t)
	return cu
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableBirthday(t *time.Time) *CustomerUpdate {
	if t != nil {
		cu.SetBirthday(*t)
	}
	return cu
}

// ClearBirthday clears the value of the "birthday" field.
func (cu *CustomerUpdate) ClearBirthday() *CustomerUpdate {
	cu.mutation.ClearBirthday()
	return cu
}

// SetPassword sets the "password" field.
func (cu *CustomerUpdate) SetPassword(s string) *CustomerUpdate {
	cu.mutation.SetPassword(s)
	return cu
}

// SetSalt sets the "salt" field.
func (cu *CustomerUpdate) SetSalt(s string) *CustomerUpdate {
	cu.mutation.SetSalt(s)
	return cu
}

// SetLastIP sets the "last_ip" field.
func (cu *CustomerUpdate) SetLastIP(s string) *CustomerUpdate {
	cu.mutation.SetLastIP(s)
	return cu
}

// SetNillableLastIP sets the "last_ip" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableLastIP(s *string) *CustomerUpdate {
	if s != nil {
		cu.SetLastIP(*s)
	}
	return cu
}

// ClearLastIP clears the value of the "last_ip" field.
func (cu *CustomerUpdate) ClearLastIP() *CustomerUpdate {
	cu.mutation.ClearLastIP()
	return cu
}

// SetLastTime sets the "last_time" field.
func (cu *CustomerUpdate) SetLastTime(t time.Time) *CustomerUpdate {
	cu.mutation.SetLastTime(t)
	return cu
}

// SetNillableLastTime sets the "last_time" field if the given value is not nil.
func (cu *CustomerUpdate) SetNillableLastTime(t *time.Time) *CustomerUpdate {
	if t != nil {
		cu.SetLastTime(*t)
	}
	return cu
}

// ClearLastTime clears the value of the "last_time" field.
func (cu *CustomerUpdate) ClearLastTime() *CustomerUpdate {
	cu.mutation.ClearLastTime()
	return cu
}

// AddWechatIDs adds the "wechats" edge to the Wechat entity by IDs.
func (cu *CustomerUpdate) AddWechatIDs(ids ...int) *CustomerUpdate {
	cu.mutation.AddWechatIDs(ids...)
	return cu
}

// AddWechats adds the "wechats" edges to the Wechat entity.
func (cu *CustomerUpdate) AddWechats(w ...*Wechat) *CustomerUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cu.AddWechatIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cu *CustomerUpdate) Mutation() *CustomerMutation {
	return cu.mutation
}

// ClearWechats clears all "wechats" edges to the Wechat entity.
func (cu *CustomerUpdate) ClearWechats() *CustomerUpdate {
	cu.mutation.ClearWechats()
	return cu
}

// RemoveWechatIDs removes the "wechats" edge to Wechat entities by IDs.
func (cu *CustomerUpdate) RemoveWechatIDs(ids ...int) *CustomerUpdate {
	cu.mutation.RemoveWechatIDs(ids...)
	return cu
}

// RemoveWechats removes "wechats" edges to Wechat entities.
func (cu *CustomerUpdate) RemoveWechats(w ...*Wechat) *CustomerUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cu.RemoveWechatIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CustomerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CustomerUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CustomerUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CustomerUpdate) defaults() {
	if _, ok := cu.mutation.UpdateTime(); !ok {
		v := customer.UpdateDefaultUpdateTime()
		cu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CustomerUpdate) check() error {
	if v, ok := cu.mutation.Sex(); ok {
		if err := customer.SexValidator(v); err != nil {
			return &ValidationError{Name: "sex", err: fmt.Errorf("ent: validator failed for field \"sex\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Account(); ok {
		if err := customer.AccountValidator(v); err != nil {
			return &ValidationError{Name: "account", err: fmt.Errorf("ent: validator failed for field \"account\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Password(); ok {
		if err := customer.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Salt(); ok {
		if err := customer.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf("ent: validator failed for field \"salt\": %w", err)}
		}
	}
	return nil
}

func (cu *CustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: customer.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldUpdateTime,
		})
	}
	if value, ok := cu.mutation.Sex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: customer.FieldSex,
		})
	}
	if value, ok := cu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: customer.FieldUUID,
		})
	}
	if value, ok := cu.mutation.Account(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAccount,
		})
	}
	if value, ok := cu.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAvatar,
		})
	}
	if cu.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldAvatar,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldName,
		})
	}
	if cu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldName,
		})
	}
	if value, ok := cu.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldRole,
		})
	}
	if cu.mutation.RoleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldRole,
		})
	}
	if value, ok := cu.mutation.NickName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldNickName,
		})
	}
	if cu.mutation.NickNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldNickName,
		})
	}
	if value, ok := cu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldEmail,
		})
	}
	if value, ok := cu.mutation.Mobile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldMobile,
		})
	}
	if cu.mutation.MobileCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldMobile,
		})
	}
	if value, ok := cu.mutation.IDCard(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldIDCard,
		})
	}
	if cu.mutation.IDCardCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldIDCard,
		})
	}
	if value, ok := cu.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldBirthday,
		})
	}
	if cu.mutation.BirthdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: customer.FieldBirthday,
		})
	}
	if value, ok := cu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPassword,
		})
	}
	if value, ok := cu.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldSalt,
		})
	}
	if value, ok := cu.mutation.LastIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldLastIP,
		})
	}
	if cu.mutation.LastIPCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldLastIP,
		})
	}
	if value, ok := cu.mutation.LastTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldLastTime,
		})
	}
	if cu.mutation.LastTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: customer.FieldLastTime,
		})
	}
	if cu.mutation.WechatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.WechatsTable,
			Columns: []string{customer.WechatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: wechat.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedWechatsIDs(); len(nodes) > 0 && !cu.mutation.WechatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.WechatsTable,
			Columns: []string{customer.WechatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: wechat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.WechatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.WechatsTable,
			Columns: []string{customer.WechatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: wechat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CustomerUpdateOne is the builder for updating a single Customer entity.
type CustomerUpdateOne struct {
	config
	hooks    []Hook
	mutation *CustomerMutation
}

// SetSex sets the "sex" field.
func (cuo *CustomerUpdateOne) SetSex(c customer.Sex) *CustomerUpdateOne {
	cuo.mutation.SetSex(c)
	return cuo
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableSex(c *customer.Sex) *CustomerUpdateOne {
	if c != nil {
		cuo.SetSex(*c)
	}
	return cuo
}

// SetUUID sets the "uuid" field.
func (cuo *CustomerUpdateOne) SetUUID(u uuid.UUID) *CustomerUpdateOne {
	cuo.mutation.SetUUID(u)
	return cuo
}

// SetAccount sets the "account" field.
func (cuo *CustomerUpdateOne) SetAccount(s string) *CustomerUpdateOne {
	cuo.mutation.SetAccount(s)
	return cuo
}

// SetAvatar sets the "avatar" field.
func (cuo *CustomerUpdateOne) SetAvatar(s string) *CustomerUpdateOne {
	cuo.mutation.SetAvatar(s)
	return cuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableAvatar(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetAvatar(*s)
	}
	return cuo
}

// ClearAvatar clears the value of the "avatar" field.
func (cuo *CustomerUpdateOne) ClearAvatar() *CustomerUpdateOne {
	cuo.mutation.ClearAvatar()
	return cuo
}

// SetName sets the "name" field.
func (cuo *CustomerUpdateOne) SetName(s string) *CustomerUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableName(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// ClearName clears the value of the "name" field.
func (cuo *CustomerUpdateOne) ClearName() *CustomerUpdateOne {
	cuo.mutation.ClearName()
	return cuo
}

// SetRole sets the "role" field.
func (cuo *CustomerUpdateOne) SetRole(s string) *CustomerUpdateOne {
	cuo.mutation.SetRole(s)
	return cuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableRole(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetRole(*s)
	}
	return cuo
}

// ClearRole clears the value of the "role" field.
func (cuo *CustomerUpdateOne) ClearRole() *CustomerUpdateOne {
	cuo.mutation.ClearRole()
	return cuo
}

// SetNickName sets the "nick_name" field.
func (cuo *CustomerUpdateOne) SetNickName(s string) *CustomerUpdateOne {
	cuo.mutation.SetNickName(s)
	return cuo
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableNickName(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetNickName(*s)
	}
	return cuo
}

// ClearNickName clears the value of the "nick_name" field.
func (cuo *CustomerUpdateOne) ClearNickName() *CustomerUpdateOne {
	cuo.mutation.ClearNickName()
	return cuo
}

// SetEmail sets the "email" field.
func (cuo *CustomerUpdateOne) SetEmail(s string) *CustomerUpdateOne {
	cuo.mutation.SetEmail(s)
	return cuo
}

// SetMobile sets the "mobile" field.
func (cuo *CustomerUpdateOne) SetMobile(s string) *CustomerUpdateOne {
	cuo.mutation.SetMobile(s)
	return cuo
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableMobile(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetMobile(*s)
	}
	return cuo
}

// ClearMobile clears the value of the "mobile" field.
func (cuo *CustomerUpdateOne) ClearMobile() *CustomerUpdateOne {
	cuo.mutation.ClearMobile()
	return cuo
}

// SetIDCard sets the "id_card" field.
func (cuo *CustomerUpdateOne) SetIDCard(s string) *CustomerUpdateOne {
	cuo.mutation.SetIDCard(s)
	return cuo
}

// SetNillableIDCard sets the "id_card" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableIDCard(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetIDCard(*s)
	}
	return cuo
}

// ClearIDCard clears the value of the "id_card" field.
func (cuo *CustomerUpdateOne) ClearIDCard() *CustomerUpdateOne {
	cuo.mutation.ClearIDCard()
	return cuo
}

// SetBirthday sets the "birthday" field.
func (cuo *CustomerUpdateOne) SetBirthday(t time.Time) *CustomerUpdateOne {
	cuo.mutation.SetBirthday(t)
	return cuo
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableBirthday(t *time.Time) *CustomerUpdateOne {
	if t != nil {
		cuo.SetBirthday(*t)
	}
	return cuo
}

// ClearBirthday clears the value of the "birthday" field.
func (cuo *CustomerUpdateOne) ClearBirthday() *CustomerUpdateOne {
	cuo.mutation.ClearBirthday()
	return cuo
}

// SetPassword sets the "password" field.
func (cuo *CustomerUpdateOne) SetPassword(s string) *CustomerUpdateOne {
	cuo.mutation.SetPassword(s)
	return cuo
}

// SetSalt sets the "salt" field.
func (cuo *CustomerUpdateOne) SetSalt(s string) *CustomerUpdateOne {
	cuo.mutation.SetSalt(s)
	return cuo
}

// SetLastIP sets the "last_ip" field.
func (cuo *CustomerUpdateOne) SetLastIP(s string) *CustomerUpdateOne {
	cuo.mutation.SetLastIP(s)
	return cuo
}

// SetNillableLastIP sets the "last_ip" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableLastIP(s *string) *CustomerUpdateOne {
	if s != nil {
		cuo.SetLastIP(*s)
	}
	return cuo
}

// ClearLastIP clears the value of the "last_ip" field.
func (cuo *CustomerUpdateOne) ClearLastIP() *CustomerUpdateOne {
	cuo.mutation.ClearLastIP()
	return cuo
}

// SetLastTime sets the "last_time" field.
func (cuo *CustomerUpdateOne) SetLastTime(t time.Time) *CustomerUpdateOne {
	cuo.mutation.SetLastTime(t)
	return cuo
}

// SetNillableLastTime sets the "last_time" field if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableLastTime(t *time.Time) *CustomerUpdateOne {
	if t != nil {
		cuo.SetLastTime(*t)
	}
	return cuo
}

// ClearLastTime clears the value of the "last_time" field.
func (cuo *CustomerUpdateOne) ClearLastTime() *CustomerUpdateOne {
	cuo.mutation.ClearLastTime()
	return cuo
}

// AddWechatIDs adds the "wechats" edge to the Wechat entity by IDs.
func (cuo *CustomerUpdateOne) AddWechatIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.AddWechatIDs(ids...)
	return cuo
}

// AddWechats adds the "wechats" edges to the Wechat entity.
func (cuo *CustomerUpdateOne) AddWechats(w ...*Wechat) *CustomerUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cuo.AddWechatIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cuo *CustomerUpdateOne) Mutation() *CustomerMutation {
	return cuo.mutation
}

// ClearWechats clears all "wechats" edges to the Wechat entity.
func (cuo *CustomerUpdateOne) ClearWechats() *CustomerUpdateOne {
	cuo.mutation.ClearWechats()
	return cuo
}

// RemoveWechatIDs removes the "wechats" edge to Wechat entities by IDs.
func (cuo *CustomerUpdateOne) RemoveWechatIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.RemoveWechatIDs(ids...)
	return cuo
}

// RemoveWechats removes "wechats" edges to Wechat entities.
func (cuo *CustomerUpdateOne) RemoveWechats(w ...*Wechat) *CustomerUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cuo.RemoveWechatIDs(ids...)
}

// Save executes the query and returns the updated Customer entity.
func (cuo *CustomerUpdateOne) Save(ctx context.Context) (*Customer, error) {
	var (
		err  error
		node *Customer
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CustomerUpdateOne) SaveX(ctx context.Context) *Customer {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CustomerUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CustomerUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdateTime(); !ok {
		v := customer.UpdateDefaultUpdateTime()
		cuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CustomerUpdateOne) check() error {
	if v, ok := cuo.mutation.Sex(); ok {
		if err := customer.SexValidator(v); err != nil {
			return &ValidationError{Name: "sex", err: fmt.Errorf("ent: validator failed for field \"sex\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Account(); ok {
		if err := customer.AccountValidator(v); err != nil {
			return &ValidationError{Name: "account", err: fmt.Errorf("ent: validator failed for field \"account\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Password(); ok {
		if err := customer.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Salt(); ok {
		if err := customer.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf("ent: validator failed for field \"salt\": %w", err)}
		}
	}
	return nil
}

func (cuo *CustomerUpdateOne) sqlSave(ctx context.Context) (_node *Customer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: customer.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Customer.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldUpdateTime,
		})
	}
	if value, ok := cuo.mutation.Sex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: customer.FieldSex,
		})
	}
	if value, ok := cuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: customer.FieldUUID,
		})
	}
	if value, ok := cuo.mutation.Account(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAccount,
		})
	}
	if value, ok := cuo.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAvatar,
		})
	}
	if cuo.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldAvatar,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldName,
		})
	}
	if cuo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldName,
		})
	}
	if value, ok := cuo.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldRole,
		})
	}
	if cuo.mutation.RoleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldRole,
		})
	}
	if value, ok := cuo.mutation.NickName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldNickName,
		})
	}
	if cuo.mutation.NickNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldNickName,
		})
	}
	if value, ok := cuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldEmail,
		})
	}
	if value, ok := cuo.mutation.Mobile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldMobile,
		})
	}
	if cuo.mutation.MobileCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldMobile,
		})
	}
	if value, ok := cuo.mutation.IDCard(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldIDCard,
		})
	}
	if cuo.mutation.IDCardCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldIDCard,
		})
	}
	if value, ok := cuo.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldBirthday,
		})
	}
	if cuo.mutation.BirthdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: customer.FieldBirthday,
		})
	}
	if value, ok := cuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPassword,
		})
	}
	if value, ok := cuo.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldSalt,
		})
	}
	if value, ok := cuo.mutation.LastIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldLastIP,
		})
	}
	if cuo.mutation.LastIPCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: customer.FieldLastIP,
		})
	}
	if value, ok := cuo.mutation.LastTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldLastTime,
		})
	}
	if cuo.mutation.LastTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: customer.FieldLastTime,
		})
	}
	if cuo.mutation.WechatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.WechatsTable,
			Columns: []string{customer.WechatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: wechat.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedWechatsIDs(); len(nodes) > 0 && !cuo.mutation.WechatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.WechatsTable,
			Columns: []string{customer.WechatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: wechat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.WechatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.WechatsTable,
			Columns: []string{customer.WechatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: wechat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Customer{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
