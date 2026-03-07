-- name: CreateUser :one
INSERT INTO users (
    id,
    email,
    password_hash,
    tenant_id,
    is_internal
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;


-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;


-- name: ListUserByTenant :many
SELECT *
FROM users
WHERE tenant_id = $1
ORDER BY created_at DESC;


-- name: ListInternalUser :many
SELECT *
FROM users
WHERE is_internal = TRUE
ORDER BY created_at DESC;


-- name: UpdateUserTenant :one
UPDATE users
SET tenant_id = $2
WHERE id = $1
RETURNING *;


-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;