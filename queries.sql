-- name: GetPermission :one
SELECT * FROM luckperms_user_permissions
WHERE uuid = $1 AND (permission = $2 OR permission = '*')
LIMIT 1;