package utils

import "github.com/google/uuid"

func NewID() uuid.UUID {
	return uuid.Must(uuid.NewV7())
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}