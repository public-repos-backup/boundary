// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package iam

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/boundary/globals"
	"github.com/hashicorp/boundary/internal/db"
	"github.com/hashicorp/boundary/internal/db/timestamp"
	"github.com/hashicorp/boundary/internal/errors"
	"github.com/hashicorp/boundary/internal/util"
	"github.com/hashicorp/go-dbw"
)

// CreateRole will create a role in the repository and return the written
// role.  No options are currently supported.
func (r *Repository) CreateRole(ctx context.Context, role *Role, opt ...Option) (*Role, []*PrincipalRole, []*RoleGrant, []*RoleGrantScope, error) {
	const op = "iam.(Repository).CreateRole"

	switch {
	case role == nil:
		return nil, nil, nil, nil, errors.New(ctx, errors.InvalidParameter, op, "missing role")
	case role.Role == nil:
		return nil, nil, nil, nil, errors.New(ctx, errors.InvalidParameter, op, "missing role store")
	case role.PublicId != "":
		return nil, nil, nil, nil, errors.New(ctx, errors.InvalidParameter, op, "public id not empty")
	case role.ScopeId == "":
		return nil, nil, nil, nil, errors.New(ctx, errors.InvalidParameter, op, "missing scope id")
	}

	id, err := newRoleId(ctx)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(ctx, err, op)
	}
	c := role.Clone().(*Role)
	c.PublicId = id

	var resource Resource
	var pr []*PrincipalRole
	var rg []*RoleGrant
	var grantScopes []*RoleGrantScope
	_, err = r.writer.DoTx(
		ctx,
		db.StdRetryCnt,
		db.ExpBackoff{},
		func(reader db.Reader, writer db.Writer) error {
			resource, err = r.create(ctx, c, WithReaderWriter(reader, writer))
			if err != nil {
				return errors.Wrap(ctx, err, op, errors.WithMsg("while creating role"))
			}
			_, _, err = r.SetRoleGrantScopes(ctx, id, resource.(*Role).Version, []string{globals.GrantScopeThis}, WithReaderWriter(reader, writer))
			if err != nil {
				return errors.Wrap(ctx, err, op, errors.WithMsg("while setting grant scopes"))
			}
			// Do a fresh lookup to get all return values
			resource, pr, rg, grantScopes, err = r.LookupRole(ctx, resource.(*Role).PublicId, WithReaderWriter(reader, writer))
			if err != nil {
				return errors.Wrap(ctx, err, op)
			}
			return nil
		})
	if err != nil {
		if errors.IsUniqueError(err) {
			return nil, nil, nil, nil, errors.New(ctx, errors.NotUnique, op, fmt.Sprintf("role %s already exists in scope %s", role.Name, role.ScopeId))
		}
		return nil, nil, nil, nil, errors.Wrap(ctx, err, op, errors.WithMsg(fmt.Sprintf("for %s", c.PublicId)))
	}
	return resource.(*Role), pr, rg, grantScopes, nil
}

// UpdateRole will update a role in the repository and return the written role.
// fieldMaskPaths provides field_mask.proto paths for fields that should be
// updated.  Fields will be set to NULL if the field is a zero value and
// included in fieldMask. Name, Description, and GrantScopeId are the only
// updatable fields, If no updatable fields are included in the fieldMaskPaths,
// then an error is returned.
func (r *Repository) UpdateRole(ctx context.Context, role *Role, version uint32, fieldMaskPaths []string, opt ...Option) (*Role, []*PrincipalRole, []*RoleGrant, []*RoleGrantScope, int, error) {
	const op = "iam.(Repository).UpdateRole"
	if role == nil {
		return nil, nil, nil, nil, db.NoRowsAffected, errors.New(ctx, errors.InvalidParameter, op, "missing role")
	}
	if role.Role == nil {
		return nil, nil, nil, nil, db.NoRowsAffected, errors.New(ctx, errors.InvalidParameter, op, "missing role store")
	}
	if role.PublicId == "" {
		return nil, nil, nil, nil, db.NoRowsAffected, errors.New(ctx, errors.InvalidParameter, op, "missing public id")
	}

	for _, f := range fieldMaskPaths {
		switch {
		case strings.EqualFold("name", f):
		case strings.EqualFold("description", f):
		default:
			return nil, nil, nil, nil, db.NoRowsAffected, errors.New(ctx, errors.InvalidFieldMask, op, fmt.Sprintf("invalid field mask: %s", f))
		}
	}

	var dbMask, nullFields []string
	dbMask, nullFields = dbw.BuildUpdatePaths(
		map[string]any{
			"name":        role.Name,
			"description": role.Description,
		},
		fieldMaskPaths,
		nil,
	)

	if len(dbMask) == 0 && len(nullFields) == 0 {
		return nil, nil, nil, nil, db.NoRowsAffected, errors.E(ctx, errors.WithCode(errors.EmptyFieldMask), errors.WithOp(op))
	}

	var resource Resource
	var rowsUpdated int
	var pr []*PrincipalRole
	var rg []*RoleGrant
	var grantScopes []*RoleGrantScope
	_, err := r.writer.DoTx(
		ctx,
		db.StdRetryCnt,
		db.ExpBackoff{},
		func(read db.Reader, w db.Writer) error {
			var err error
			c := role.Clone().(*Role)
			resource = c // If we don't have dbMask or nullFields, we'll return this
			if len(dbMask) > 0 || len(nullFields) > 0 {
				resource, rowsUpdated, err = r.update(ctx, c, version, dbMask, nullFields, WithReaderWriter(read, w))
				if err != nil {
					return errors.Wrap(ctx, err, op)
				}
				version = resource.(*Role).Version
			}

			// Do a fresh lookup since version may have gone up by 1 or 2 based
			// on grant scope id
			resource, pr, rg, grantScopes, err = r.LookupRole(ctx, role.PublicId, WithReaderWriter(read, w))
			if err != nil {
				return errors.Wrap(ctx, err, op)
			}
			return nil
		},
	)
	if err != nil {
		if errors.IsUniqueError(err) {
			return nil, nil, nil, nil, db.NoRowsAffected, errors.New(ctx, errors.NotUnique, op, fmt.Sprintf("role %s already exists in org %s", role.Name, role.ScopeId))
		}
		return nil, nil, nil, nil, db.NoRowsAffected, errors.Wrap(ctx, err, op, errors.WithMsg(fmt.Sprintf("for %s", role.PublicId)))
	}
	return resource.(*Role), pr, rg, grantScopes, rowsUpdated, nil
}

// LookupRole will look up a role in the repository.  If the role is not
// found, it will return nil, nil.
//
// Supported options: WithReaderWriter
func (r *Repository) LookupRole(ctx context.Context, withPublicId string, opt ...Option) (*Role, []*PrincipalRole, []*RoleGrant, []*RoleGrantScope, error) {
	const op = "iam.(Repository).LookupRole"
	if withPublicId == "" {
		return nil, nil, nil, nil, errors.New(ctx, errors.InvalidParameter, op, "missing public id")
	}
	opts := getOpts(opt...)
	role := allocRole()
	role.PublicId = withPublicId
	var pr []*PrincipalRole
	var rg []*RoleGrant
	var rgs []*RoleGrantScope

	lookupFunc := func(read db.Reader, w db.Writer) error {
		if err := read.LookupByPublicId(ctx, &role); err != nil {
			return errors.Wrap(ctx, err, op)
		}
		repo, err := NewRepository(ctx, read, w, r.kms)
		if err != nil {
			return errors.Wrap(ctx, err, op)
		}
		pr, err = repo.ListPrincipalRoles(ctx, withPublicId)
		if err != nil {
			return errors.Wrap(ctx, err, op)
		}
		rg, err = repo.ListRoleGrants(ctx, withPublicId)
		if err != nil {
			return errors.Wrap(ctx, err, op)
		}
		rgs, err = repo.ListRoleGrantScopes(ctx, []string{withPublicId})
		if err != nil {
			return errors.Wrap(ctx, err, op)
		}
		return nil
	}

	var err error
	if !util.IsNil(opts.withReader) && !util.IsNil(opts.withWriter) {
		if !opts.withWriter.IsTx(ctx) {
			return nil, nil, nil, nil, errors.New(ctx, errors.Internal, op, "writer is not in transaction")
		}
		err = lookupFunc(opts.withReader, opts.withWriter)
	} else {
		_, err = r.writer.DoTx(
			ctx,
			db.StdRetryCnt,
			db.ExpBackoff{},
			lookupFunc,
		)
	}
	if err != nil {
		if errors.IsNotFoundError(err) {
			return nil, nil, nil, nil, nil
		}
		return nil, nil, nil, nil, errors.Wrap(ctx, err, op, errors.WithMsg(fmt.Sprintf("for %s", withPublicId)))
	}
	return &role, pr, rg, rgs, nil
}

// DeleteRole will delete a role from the repository.
func (r *Repository) DeleteRole(ctx context.Context, withPublicId string, _ ...Option) (int, error) {
	const op = "iam.(Repository).DeleteRole"
	if withPublicId == "" {
		return db.NoRowsAffected, errors.New(ctx, errors.InvalidParameter, op, "missing public id")
	}
	role := allocRole()
	role.PublicId = withPublicId
	if err := r.reader.LookupByPublicId(ctx, &role); err != nil {
		return db.NoRowsAffected, errors.Wrap(ctx, err, op, errors.WithMsg(fmt.Sprintf("failed for %s", withPublicId)))
	}
	rowsDeleted, err := r.delete(ctx, &role)
	if err != nil {
		return db.NoRowsAffected, errors.Wrap(ctx, err, op, errors.WithMsg(fmt.Sprintf("failed for %s", withPublicId)))
	}
	return rowsDeleted, nil
}

// listRoles lists roles in the given scopes and supports WithLimit option.
func (r *Repository) listRoles(ctx context.Context, withScopeIds []string, opt ...Option) ([]*Role, time.Time, error) {
	const op = "iam.(Repository).listRoles"
	if len(withScopeIds) == 0 {
		return nil, time.Time{}, errors.New(ctx, errors.InvalidParameter, op, "missing scope id")
	}
	opts := getOpts(opt...)

	limit := r.defaultLimit
	switch {
	case opts.withLimit > 0:
		// non-zero signals an override of the default limit for the repo.
		limit = opts.withLimit
	case opts.withLimit < 0:
		return nil, time.Time{}, errors.New(ctx, errors.InvalidParameter, op, "limit must be non-negative")
	}

	var args []any
	whereClause := "scope_id in @scope_ids"
	args = append(args, sql.Named("scope_ids", withScopeIds))

	if opts.withStartPageAfterItem != nil {
		whereClause = fmt.Sprintf("(create_time, public_id) < (@last_item_create_time, @last_item_id) and %s", whereClause)
		args = append(args,
			sql.Named("last_item_create_time", opts.withStartPageAfterItem.GetCreateTime()),
			sql.Named("last_item_id", opts.withStartPageAfterItem.GetPublicId()),
		)
	}
	dbOpts := []db.Option{db.WithLimit(limit), db.WithOrder("create_time desc, public_id desc")}
	return r.queryRoles(ctx, whereClause, args, dbOpts...)
}

// listRolesRefresh lists roles in the given scopes and supports the
// WithLimit and WithStartPageAfterItem options.
func (r *Repository) listRolesRefresh(ctx context.Context, updatedAfter time.Time, withScopeIds []string, opt ...Option) ([]*Role, time.Time, error) {
	const op = "iam.(Repository).listRolesRefresh"

	switch {
	case updatedAfter.IsZero():
		return nil, time.Time{}, errors.New(ctx, errors.InvalidParameter, op, "missing updated after time")

	case len(withScopeIds) == 0:
		return nil, time.Time{}, errors.New(ctx, errors.InvalidParameter, op, "missing scope id")
	}

	opts := getOpts(opt...)

	limit := r.defaultLimit
	switch {
	case opts.withLimit > 0:
		// non-zero signals an override of the default limit for the repo.
		limit = opts.withLimit
	case opts.withLimit < 0:
		return nil, time.Time{}, errors.New(ctx, errors.InvalidParameter, op, "limit must be non-negative")
	}

	var args []any
	whereClause := "update_time > @updated_after_time and scope_id in @scope_ids"
	args = append(args,
		sql.Named("updated_after_time", timestamp.New(updatedAfter)),
		sql.Named("scope_ids", withScopeIds),
	)
	if opts.withStartPageAfterItem != nil {
		whereClause = fmt.Sprintf("(update_time, public_id) < (@last_item_update_time, @last_item_id) and %s", whereClause)
		args = append(args,
			sql.Named("last_item_update_time", opts.withStartPageAfterItem.GetUpdateTime()),
			sql.Named("last_item_id", opts.withStartPageAfterItem.GetPublicId()),
		)
	}

	dbOpts := []db.Option{db.WithLimit(limit), db.WithOrder("update_time desc, public_id desc")}
	return r.queryRoles(ctx, whereClause, args, dbOpts...)
}

func (r *Repository) queryRoles(ctx context.Context, whereClause string, args []any, opt ...db.Option) ([]*Role, time.Time, error) {
	const op = "iam.(Repository).queryRoles"

	var transactionTimestamp time.Time
	var retRoles []*Role
	var retRoleGrantScopes []*RoleGrantScope
	if _, err := r.writer.DoTx(ctx, db.StdRetryCnt, db.ExpBackoff{}, func(rd db.Reader, w db.Writer) error {
		var inRet []*Role
		if err := rd.SearchWhere(ctx, &inRet, whereClause, args, opt...); err != nil {
			return errors.Wrap(ctx, err, op, errors.WithMsg("failed to query roles"))
		}
		retRoles = inRet
		var err error
		if len(retRoles) > 0 {
			roleIds := make([]string, 0, len(retRoles))
			for _, retRole := range retRoles {
				roleIds = append(roleIds, retRole.PublicId)
			}
			retRoleGrantScopes, err = r.ListRoleGrantScopes(ctx, roleIds, WithReaderWriter(rd, w))
			if err != nil {
				return errors.Wrap(ctx, err, op, errors.WithMsg("failed to query role grant scopes"))
			}
		}
		transactionTimestamp, err = rd.Now(ctx)
		return err
	}); err != nil {
		return nil, time.Time{}, err
	}
	roleGrantScopesMap := make(map[string][]*RoleGrantScope)
	for _, rgs := range retRoleGrantScopes {
		roleGrantScopesMap[rgs.RoleId] = append(roleGrantScopesMap[rgs.RoleId], rgs)
	}
	for _, role := range retRoles {
		role.GrantScopes = roleGrantScopesMap[role.PublicId]
	}
	return retRoles, transactionTimestamp, nil
}

// listRoleDeletedIds lists the public IDs of any roles deleted since the timestamp provided.
func (r *Repository) listRoleDeletedIds(ctx context.Context, since time.Time) ([]string, time.Time, error) {
	const op = "iam.(Repository).listRoleDeletedIds"
	var deletedResources []*deletedRole
	var transactionTimestamp time.Time
	if _, err := r.writer.DoTx(ctx, db.StdRetryCnt, db.ExpBackoff{}, func(r db.Reader, _ db.Writer) error {
		if err := r.SearchWhere(ctx, &deletedResources, "delete_time >= ?", []any{since}); err != nil {
			return errors.Wrap(ctx, err, op, errors.WithMsg("failed to query deleted roles"))
		}
		var err error
		transactionTimestamp, err = r.Now(ctx)
		if err != nil {
			return errors.Wrap(ctx, err, op, errors.WithMsg("failed to get transaction timestamp"))
		}
		return nil
	}); err != nil {
		return nil, time.Time{}, err
	}
	var dIds []string
	for _, res := range deletedResources {
		dIds = append(dIds, res.PublicId)
	}
	return dIds, transactionTimestamp, nil
}

// estimatedRoleCount returns an estimate of the total number of items in the iam_role table.
func (r *Repository) estimatedRoleCount(ctx context.Context) (int, error) {
	const op = "iam.(Repository).estimatedRoleCount"
	rows, err := r.reader.Query(ctx, estimateCountRoles, nil)
	if err != nil {
		return 0, errors.Wrap(ctx, err, op, errors.WithMsg("failed to query total roles"))
	}
	var count int
	for rows.Next() {
		if err := r.reader.ScanRows(ctx, rows, &count); err != nil {
			return 0, errors.Wrap(ctx, err, op, errors.WithMsg("failed to query total roles"))
		}
	}
	if err := rows.Err(); err != nil {
		return 0, errors.Wrap(ctx, err, op, errors.WithMsg("failed to query total roles"))
	}
	return count, nil
}
