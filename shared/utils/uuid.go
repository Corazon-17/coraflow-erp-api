package utils

import "github.com/google/uuid"

func NewUUID() uuid.UUID {
	return uuid.Must(uuid.NewV7())
}

func ToUUID(strId string) (uuid.UUID, error) {
	id, err := uuid.Parse(strId)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
