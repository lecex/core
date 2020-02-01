package env

import (
	"testing"

	"github.com/lecex/core/env"
)

func TestEnv(t *testing.T) {
	val := env.Getenv("DB_DRIVER", "123456")
	if val != "123456" {
		t.Errorf("Database connection failed, %v!", val)
	}
}
