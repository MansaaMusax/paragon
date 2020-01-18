// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/kcarretto/paragon/ent/credential"
	"github.com/kcarretto/paragon/ent/predicate"
)

// CredentialUpdate is the builder for updating Credential entities.
type CredentialUpdate struct {
	config
	principal  *string
	secret     *string
	kind       *credential.Kind
	fails      *int
	addfails   *int
	predicates []predicate.Credential
}

// Where adds a new predicate for the builder.
func (cu *CredentialUpdate) Where(ps ...predicate.Credential) *CredentialUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetPrincipal sets the principal field.
func (cu *CredentialUpdate) SetPrincipal(s string) *CredentialUpdate {
	cu.principal = &s
	return cu
}

// SetSecret sets the secret field.
func (cu *CredentialUpdate) SetSecret(s string) *CredentialUpdate {
	cu.secret = &s
	return cu
}

// SetKind sets the kind field.
func (cu *CredentialUpdate) SetKind(c credential.Kind) *CredentialUpdate {
	cu.kind = &c
	return cu
}

// SetFails sets the fails field.
func (cu *CredentialUpdate) SetFails(i int) *CredentialUpdate {
	cu.fails = &i
	cu.addfails = nil
	return cu
}

// SetNillableFails sets the fails field if the given value is not nil.
func (cu *CredentialUpdate) SetNillableFails(i *int) *CredentialUpdate {
	if i != nil {
		cu.SetFails(*i)
	}
	return cu
}

// AddFails adds i to fails.
func (cu *CredentialUpdate) AddFails(i int) *CredentialUpdate {
	if cu.addfails == nil {
		cu.addfails = &i
	} else {
		*cu.addfails += i
	}
	return cu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CredentialUpdate) Save(ctx context.Context) (int, error) {
	if cu.kind != nil {
		if err := credential.KindValidator(*cu.kind); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"kind\": %v", err)
		}
	}
	if cu.fails != nil {
		if err := credential.FailsValidator(*cu.fails); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"fails\": %v", err)
		}
	}
	return cu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CredentialUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CredentialUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CredentialUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CredentialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(credential.FieldID).From(sql.Table(credential.Table))
	for _, p := range cu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = cu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := cu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		builder = sql.Update(credential.Table).Where(sql.InInts(credential.FieldID, ids...))
	)
	if value := cu.principal; value != nil {
		builder.Set(credential.FieldPrincipal, *value)
	}
	if value := cu.secret; value != nil {
		builder.Set(credential.FieldSecret, *value)
	}
	if value := cu.kind; value != nil {
		builder.Set(credential.FieldKind, *value)
	}
	if value := cu.fails; value != nil {
		builder.Set(credential.FieldFails, *value)
	}
	if value := cu.addfails; value != nil {
		builder.Add(credential.FieldFails, *value)
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// CredentialUpdateOne is the builder for updating a single Credential entity.
type CredentialUpdateOne struct {
	config
	id        int
	principal *string
	secret    *string
	kind      *credential.Kind
	fails     *int
	addfails  *int
}

// SetPrincipal sets the principal field.
func (cuo *CredentialUpdateOne) SetPrincipal(s string) *CredentialUpdateOne {
	cuo.principal = &s
	return cuo
}

// SetSecret sets the secret field.
func (cuo *CredentialUpdateOne) SetSecret(s string) *CredentialUpdateOne {
	cuo.secret = &s
	return cuo
}

// SetKind sets the kind field.
func (cuo *CredentialUpdateOne) SetKind(c credential.Kind) *CredentialUpdateOne {
	cuo.kind = &c
	return cuo
}

// SetFails sets the fails field.
func (cuo *CredentialUpdateOne) SetFails(i int) *CredentialUpdateOne {
	cuo.fails = &i
	cuo.addfails = nil
	return cuo
}

// SetNillableFails sets the fails field if the given value is not nil.
func (cuo *CredentialUpdateOne) SetNillableFails(i *int) *CredentialUpdateOne {
	if i != nil {
		cuo.SetFails(*i)
	}
	return cuo
}

// AddFails adds i to fails.
func (cuo *CredentialUpdateOne) AddFails(i int) *CredentialUpdateOne {
	if cuo.addfails == nil {
		cuo.addfails = &i
	} else {
		*cuo.addfails += i
	}
	return cuo
}

// Save executes the query and returns the updated entity.
func (cuo *CredentialUpdateOne) Save(ctx context.Context) (*Credential, error) {
	if cuo.kind != nil {
		if err := credential.KindValidator(*cuo.kind); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"kind\": %v", err)
		}
	}
	if cuo.fails != nil {
		if err := credential.FailsValidator(*cuo.fails); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"fails\": %v", err)
		}
	}
	return cuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CredentialUpdateOne) SaveX(ctx context.Context) *Credential {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CredentialUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CredentialUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CredentialUpdateOne) sqlSave(ctx context.Context) (c *Credential, err error) {
	selector := sql.Select(credential.Columns...).From(sql.Table(credential.Table))
	credential.ID(cuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = cuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		c = &Credential{config: cuo.config}
		if err := c.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Credential: %v", err)
		}
		id = c.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("Credential with id: %v", cuo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Credential with the same id: %v", cuo.id)
	}

	tx, err := cuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		builder = sql.Update(credential.Table).Where(sql.InInts(credential.FieldID, ids...))
	)
	if value := cuo.principal; value != nil {
		builder.Set(credential.FieldPrincipal, *value)
		c.Principal = *value
	}
	if value := cuo.secret; value != nil {
		builder.Set(credential.FieldSecret, *value)
		c.Secret = *value
	}
	if value := cuo.kind; value != nil {
		builder.Set(credential.FieldKind, *value)
		c.Kind = *value
	}
	if value := cuo.fails; value != nil {
		builder.Set(credential.FieldFails, *value)
		c.Fails = *value
	}
	if value := cuo.addfails; value != nil {
		builder.Add(credential.FieldFails, *value)
		c.Fails += *value
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return c, nil
}
