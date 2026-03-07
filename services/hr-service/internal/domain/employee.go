package domain

import "github.com/google/uuid"

type Employee struct {
	ID           uuid.UUID
	TenantID     uuid.UUID
	FirstName    string
	LastName     string
	DepartmentID uuid.UUID
}