package env

import (
	"testing"
)

func TestEnv(t *testing.T) {
	val := Getenv("DB_DRIVER", "123456")
	if val != "123456" {
		t.Errorf("Database connection failed, %v!", val)
	}
}
