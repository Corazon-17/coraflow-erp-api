-- name: CreateDepartment :one
INSERT INTO departments (
    id,
    tenant_id,
    name,
    parent_id
)
VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: GetDepartment :one
SELECT *
FROM departments
WHERE tenant_id = $1
AND id = $2;


-- name: ListDepartment :many
SELECT *
FROM departments
WHERE tenant_id = $1
ORDER BY created_at DESC;


-- name: ListRootDepartment :many
SELECT *
FROM departments
WHERE tenant_id = $1
AND parent_id IS NULL
ORDER BY created_at DESC;


-- name: ListChildDepartment :many
SELECT *
FROM departments
WHERE tenant_id = $1
AND parent_id = $2
ORDER BY created_at DESC;


-- name: UpdateDepartment :one
UPDATE departments
SET name = $3,
    parent_id = $4
WHERE tenant_id = $1
AND id = $2
RETURNING *;


-- name: DeleteDepartment :exec
DELETE FROM departments
WHERE tenant_id = $1
AND id = $2;