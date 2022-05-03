package database

import "time"

// Config for database connection
// TODO Set all default values
type Config struct {
	Driver string `yaml:"driver" env:"DATABASE_DRIVER" env-description:"Database driver"`
	Host   string `yaml:"host" env:"DATABASE_HOST" env-description:"Database host"`
	Port   string `yaml:"port" env:"DATABASE_PORT" env-description:"Database port"`
	Name   string `yaml:"name" env:"DATABASE_NAME" env-description:"Database name"`
	User   string `yaml:"user" env:"DATABASE_USER" env-description:"Database user"`
	Pass   string `env:"DATABASE_PASS" env-description:"Database user password"`
	Pool   int    `yaml:"pool" env:"DATABASE_POOL" env-description:"Database pool size"`
	Schema string `yaml:"schema" env:"DATABASE_SCHEMA" env-description:"Database schema name"`
}

// Timestamps adding time at the end of models
type Timestamps struct {
	CreatedAt time.Time `yaml:"created_at"`
	UpdatedAt time.Time `yaml:"updated_at"`
}
