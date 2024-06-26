// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package database

import (
	"context"
)

const getGroupsOfUser = `-- name: GetGroupsOfUser :many
SELECT name FROM luckperms_groups
WHERE EXISTS (SELECT id, uuid, permission, value, server, world, expiry, contexts FROM luckperms_user_permissions WHERE luckperms_user_permissions.uuid = $1 AND luckperms_groups.name = RIGHT(permission, length(luckperms_user_permissions.permission) - 6))
`

// unused
func (q *Queries) GetGroupsOfUser(ctx context.Context, uuid string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getGroupsOfUser, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPrefixOfUser = `-- name: GetPrefixOfUser :many
SELECT split_part(prefixAlias.permission, '.', 3) as prefix, cast(split_part(prefixAlias.permission, '.', 2) as int) as weight FROM (SELECT luckperms_group_permissions.permission FROM luckperms_group_permissions
WHERE luckperms_group_permissions.value = true AND luckperms_group_permissions.permission LIKE 'prefix._%._%' AND
	EXISTS (SELECT id, uuid, permission, value, server, world, expiry, contexts FROM luckperms_user_permissions
		WHERE luckperms_user_permissions.uuid = $1 AND luckperms_group_permissions.name = RIGHT(permission, length(luckperms_user_permissions.permission) - 6))
UNION
SELECT luckperms_user_permissions.permission FROM luckperms_user_permissions
WHERE luckperms_user_permissions.permission LIKE 'prefix._%._%' AND luckperms_user_permissions.uuid = $1) as prefixAlias order by weight desc
`

type GetPrefixOfUserRow struct {
	Prefix string
	Weight int32
}

func (q *Queries) GetPrefixOfUser(ctx context.Context, uuid string) ([]GetPrefixOfUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getPrefixOfUser, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPrefixOfUserRow
	for rows.Next() {
		var i GetPrefixOfUserRow
		if err := rows.Scan(&i.Prefix, &i.Weight); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSuffixOfUser = `-- name: GetSuffixOfUser :many
SELECT split_part(suffixAlias.permission, '.', 3) as suffix, cast(split_part(suffixAlias.permission, '.', 2) as int) as weight FROM (SELECT luckperms_group_permissions.permission FROM luckperms_group_permissions
WHERE luckperms_group_permissions.value = true AND luckperms_group_permissions.permission LIKE 'suffix._%._%' AND
	EXISTS (SELECT id, uuid, permission, value, server, world, expiry, contexts FROM luckperms_user_permissions
		WHERE luckperms_user_permissions.uuid = $1 AND luckperms_group_permissions.name = RIGHT(permission, length(luckperms_user_permissions.permission) - 6))
UNION
SELECT luckperms_user_permissions.permission FROM luckperms_user_permissions
WHERE luckperms_user_permissions.permission LIKE 'suffix._%._%' AND luckperms_user_permissions.uuid = $1) as suffixAlias order by weight desc
`

type GetSuffixOfUserRow struct {
	Suffix string
	Weight int32
}

func (q *Queries) GetSuffixOfUser(ctx context.Context, uuid string) ([]GetSuffixOfUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getSuffixOfUser, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSuffixOfUserRow
	for rows.Next() {
		var i GetSuffixOfUserRow
		if err := rows.Scan(&i.Suffix, &i.Weight); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const userHasPermission = `-- name: UserHasPermission :one
SELECT (CASE
	WHEN EXISTS (SELECT 1
				 FROM luckperms_user_permissions
				 WHERE luckperms_user_permissions.uuid = $1 AND (luckperms_user_permissions.permission = $2 OR luckperms_user_permissions.permission = '*') AND value = true) THEN 1 
	WHEN EXISTS (SELECT 1
				 FROM luckperms_group_permissions
				 WHERE EXISTS (SELECT id, uuid, permission, value, server, world, expiry, contexts 
							   FROM luckperms_user_permissions 
							   WHERE luckperms_user_permissions.uuid = $1 AND luckperms_group_permissions.name = RIGHT(permission, length(luckperms_user_permissions.permission) - 6)) AND (luckperms_group_permissions.permission = $2 OR luckperms_group_permissions.permission = '*') AND luckperms_group_permissions.value = true) THEN 1
	ELSE 0 END) AS result LIMIT 1
`

type UserHasPermissionParams struct {
	Uuid       string
	Permission string
}

func (q *Queries) UserHasPermission(ctx context.Context, arg UserHasPermissionParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, userHasPermission, arg.Uuid, arg.Permission)
	var result int32
	err := row.Scan(&result)
	return result, err
}
