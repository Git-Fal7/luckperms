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

