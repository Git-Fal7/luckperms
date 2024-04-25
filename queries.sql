-- name: UserHasPermission :one
SELECT (CASE
	WHEN EXISTS (SELECT 1
				 FROM luckperms_user_permissions
				 WHERE luckperms_user_permissions.uuid = $1 AND (luckperms_user_permissions.permission = $2 OR luckperms_user_permissions.permission = '*') AND value = true) THEN 1 
	WHEN EXISTS (SELECT 1
				 FROM luckperms_group_permissions
				 WHERE EXISTS (SELECT * 
							   FROM luckperms_user_permissions 
							   WHERE luckperms_user_permissions.uuid = $1 AND luckperms_group_permissions.name = RIGHT(permission, length(luckperms_user_permissions.permission) - 6)) AND (luckperms_group_permissions.permission = $2 OR luckperms_group_permissions.permission = '*') AND luckperms_group_permissions.value = true) THEN 1
	ELSE 0 END) AS result LIMIT 1;

-- unused
-- name: GetGroupsOfUser :many
SELECT * FROM luckperms_groups
WHERE EXISTS (SELECT * FROM luckperms_user_permissions WHERE luckperms_user_permissions.uuid = $1 AND luckperms_groups.name = RIGHT(permission, length(luckperms_user_permissions.permission) - 6));

-- name: GetPrefixOfUser :many
SELECT split_part(prefixAlias.permission, '.', 3) as prefix, cast(split_part(prefixAlias.permission, '.', 2) as int) as weight FROM (SELECT luckperms_group_permissions.permission FROM luckperms_group_permissions
WHERE luckperms_group_permissions.value = true AND luckperms_group_permissions.permission LIKE 'prefix._%._%' AND
	EXISTS (SELECT * FROM luckperms_user_permissions
		WHERE luckperms_user_permissions.uuid = $1 AND luckperms_group_permissions.name = RIGHT(permission, length(luckperms_user_permissions.permission) - 6))
UNION
SELECT luckperms_user_permissions.permission FROM luckperms_user_permissions
WHERE luckperms_user_permissions.permission LIKE 'prefix._%._%' AND luckperms_user_permissions.uuid = $1) as prefixAlias order by weight desc;

-- name: GetSuffixOfUser :many
SELECT split_part(suffixAlias.permission, '.', 3) as suffix, cast(split_part(suffixAlias.permission, '.', 2) as int) as weight FROM (SELECT luckperms_group_permissions.permission FROM luckperms_group_permissions
WHERE luckperms_group_permissions.value = true AND luckperms_group_permissions.permission LIKE 'suffix._%._%' AND
	EXISTS (SELECT * FROM luckperms_user_permissions
		WHERE luckperms_user_permissions.uuid = $1 AND luckperms_group_permissions.name = RIGHT(permission, length(luckperms_user_permissions.permission) - 6))
UNION
SELECT luckperms_user_permissions.permission FROM luckperms_user_permissions
WHERE luckperms_user_permissions.permission LIKE 'suffix._%._%' AND luckperms_user_permissions.uuid = $1) as suffixAlias order by weight desc;
