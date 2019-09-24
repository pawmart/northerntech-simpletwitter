package config

import (
	"testing"
	"os"
)

func TestConfig_GetDbConfig(t *testing.T) {

	os.Setenv("NORTHTECH_DB_DATABASE", "northerntech-simpletwitter-test")
	os.Setenv("NORTHTECH_DB_HOST", "host.docker.internal")
	os.Setenv("NORTHTECH_DB_USER", "root")
	os.Setenv("NORTHTECH_DB_PASSWORD", "example")
	os.Setenv("NORTHTECH_DB_AUTH", "admin")

	cnf := NewConfig()
	dbCnf := cnf.GetDbConfig()

	if dbCnf.Auth != "admin" {
		t.Error("Auth not matching")
	}

	if dbCnf.Host != "host.docker.internal" {
		t.Error("Host not matching")
	}

	if dbCnf.User != "root" {
		t.Error("User not matching")
	}

	if dbCnf.Password != "example" {
		t.Error("Password not matching")
	}

	if dbCnf.Database != "northerntech-simpletwitter-test" {
		t.Error("Database not matching")
	}
}
