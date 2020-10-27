package utils

import (
	"github.com/google/uuid"
)

// NewUUID provides UUID
func NewUUID() string {
	return uuid.New().String()
}
