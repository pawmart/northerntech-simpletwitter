package utils

import (
	"log"

	"github.com/gofrs/uuid"
)

// GenerateUUIDString function.
func GenerateUUIDString() string {
	uuID, err := uuid.NewV4()
	if err != nil {
		log.Print("could not generate UUID", err.Error())
	}
	return uuID.String()
}
