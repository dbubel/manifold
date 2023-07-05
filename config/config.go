package config

import (
	"encoding/json"
	"fmt"
)

type ENV string

const (
	ENV_LOCAL       ENV = "local"
	ENV_TEST        ENV = "test"
	ENV_DEVELOPMENT ENV = "development"
	ENV_STAGING     ENV = "staging"
	ENV_PROD        ENV = "production"
)

type Postgres struct {
	DatabaseUsername   string `default:"cpm_user" envconfig:"DB_USERNAME"`
	DatabasePassword   string `default:"mypassword" envconfig:"DB_PASSWORD"`
	DatabaseName       string `default:"mlopscontentplatformdb" envconfig:"DB_NAME"`
	DatabasePort       string `default:"5432" envconfig:"DB_PORT"`
	DatabaseWriterHost string `default:"localhost" envconfig:"DB_WRITER_HOST"`
	DatabaseReaderHost string `default:"localhost" envconfig:"DB_READER_HOST"`
}

type Config struct {
	Postgres    Postgres
	Environment ENV    `default:"local" envconfig:"ENVIRONMENT"`
	Port        int    `default:"50051" envconfig:"PORT"`
	AWSRegion   string `default:"us-east-1" envconfig:"AWS_REGION"`
	BuildTag    string
	BuildDate   string
	GitHash     string
}

func (c Config) ReaderDSN() string {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", c.Postgres.DatabaseReaderHost, c.Postgres.DatabasePort, c.Postgres.DatabaseUsername, c.Postgres.DatabasePassword, c.Postgres.DatabaseName)
	if c.Environment == ENV_LOCAL || c.Environment == ENV_TEST {
		connStr += " sslmode=disable"
	}
	return connStr
}

func (c Config) WriterDSN() string {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", c.Postgres.DatabaseWriterHost, c.Postgres.DatabasePort, c.Postgres.DatabaseUsername, c.Postgres.DatabasePassword, c.Postgres.DatabaseName)
	if c.Environment == ENV_LOCAL || c.Environment == ENV_TEST {
		connStr += " sslmode=disable"
	}
	return connStr
}

func (c Config) Dump() string {
	c.Postgres.DatabasePassword = c.Postgres.DatabasePassword[:5]
	s, _ := json.MarshalIndent(c, "", " ")
	return string(s)
}
