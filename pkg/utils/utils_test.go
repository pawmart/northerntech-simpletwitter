package utils

import (
	"testing"
	"regexp"
)

func TestGenerateUUIDString(t *testing.T) {
	uuid := GenerateUUIDString()
	if !isValidUUID(uuid) {
		t.Errorf("invalid uuid string expected for %v, is not valid uuid", uuid)
	}
}

func isValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}