-- name: CreateEmployee :one
INSERT INTO employees (
    id,
    tenant_id,
    user_id,
    first_name,
    last_name,
    department_id,
    manager_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;


-- name: GetEmployee :one
SELECT *
FROM employees
WHERE tenant_id = $1
AND id = $2;


-- name: GetEmployeeByUser :one
SELECT *
FROM employees
WHERE tenant_id = $1
AND user_id = $2;


-- name: ListEmployee :many
SELECT *
FROM employees
WHERE tenant_id = $1
ORDER BY created_at DESC;


-- name: ListEmployeeByDepartment :many
SELECT *
FROM employees
WHERE tenant_id = $1
AND department_id = $2
ORDER BY created_at DESC;


-- name: ListEmployeeByManager :many
SELECT *
FROM employees
WHERE tenant_id = $1
AND manager_id = $2
ORDER BY created_at DESC;


-- name: UpdateEmployee :one
UPDATE employees
SET first_name = $3,
    last_name = $4,
    department_id = $5,
    manager_id = $6
WHERE tenant_id = $1
AND id = $2
RETURNING *;


-- name: UpdateEmployeeManager :one
UPDATE employees
SET manager_id = $3
WHERE tenant_id = $1
AND id = $2
RETURNING *;


-- name: DeleteEmployee :exec
DELETE FROM employees
WHERE tenant_id = $1
AND id = $2;