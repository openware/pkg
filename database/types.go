package database

import "time"

// Timestamps adding time at the end of models
type Timestamps struct {
	CreatedAt time.Time `yaml:"created_at"`
	UpdatedAt time.Time `yaml:"updated_at"`
}
