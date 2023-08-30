package memory_locks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLock(t *testing.T) {
	lock, err := NewLock(context.Background(), "test-lock-id")
	assert.Nil(t, err)
	assert.NotNil(t, lock)
}
