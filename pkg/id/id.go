package id

import "github.com/google/uuid"

type UUID = uuid.UUID

func NewUUID() UUID {
	return uuid.New()
}
