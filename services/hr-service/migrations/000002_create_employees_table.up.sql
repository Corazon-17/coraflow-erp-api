CREATE TABLE employees (
    id UUID PRIMARY KEY,
    tenant_id UUID NOT NULL,
    user_id UUID NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    department_id UUID NULL,
    manager_id UUID NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_employees_tenant
ON employees (tenant_id);

CREATE INDEX idx_employees_user
ON employees (user_id);

CREATE INDEX idx_employees_department
ON employees (department_id);