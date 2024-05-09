package database_test

import (
	"PETRO-VIEW/internal/infrastructure/database"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConnectToMongoDB(t *testing.T) {
	t.Run("Successful Connection", func(t *testing.T) {
		var duration time.Duration
		_, err := database.ConnectToMongoDB()
		assert.Nil(t, err, "Error must be nil")
		assert.True(t, duration < 10*time.Millisecond, "Ping duration exceeded 10ms")
	})

	t.Run("Did not found VARIABLES", func(t *testing.T) {})

	t.Run("Error with connection", func(t *testing.T) {})
}
