package adapters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUIDGenerator(t *testing.T) {
	t.Parallel()
	t.Run("Test UUIDGenerator", func(t *testing.T) {
		assert.NotNil(t, UUIDGenerator(), "Expected UUIDGenerator to return a UUID")
	})
}

func BenchmarkUUIDGenerator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UUIDGenerator()
	}
}
