package config

import (
	goconfig "github.com/micro/go-config"
	"github.com/micro/go-config/source/env"
)

// DbConfig holding db mongo configuration.
type DbConfig struct {
	Database   string `json:"database,omitempty"`
	Host       string `json:"host,omitempty"`
	User       string `json:"user,omitempty"`
	Password   string `json:"password,omitempty"`
	Auth       string `json:"auth,omitempty"`
	Ssl        string `json:"ssl,omitempty"`
	Replicaset string `json:"replicaset,omitempty"`
}

func NewConfig() *Config  {
	c:=  new(Config)
	return c.loadConfiguration()
}

// Config holding app configuration.
type Config struct {
	Db *DbConfig `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty"`
}

// GetDbConfig function.
func (c *Config) GetDbConfig() *DbConfig {
	return c.Db
}

// LoadConfiguration for the app.
func (*Config) loadConfiguration() *Config {

	// TODO: Fetch config from Vault or any other secure place etc.
	c := new(Config)

	src := env.NewSource(
		env.WithStrippedPrefix("NORTHTECH"),
	)

	conf := goconfig.NewConfig()
	conf.Load(src)
	conf.Scan(c)

	return c
}
