package help

import uuid "github.com/satori/go.uuid"

func GenerateUuid() string {
	return uuid.NewV4().String()
}
