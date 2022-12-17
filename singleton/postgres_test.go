package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPostgresClient(t *testing.T) {
	t.Run("Should return a pointer to postgres client", func(t *testing.T) {
		postgresClient := NewPostgresClient()
		assert.IsType(t, &PostgresClient{}, postgresClient)
	})
}
