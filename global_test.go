package memory_locks

import (
	"context"
	"errors"
	"testing"

	storage_lock "github.com/storage-lock/go-storage-lock"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
	"github.com/stretchr/testify/assert"
)

func TestNewLock(t *testing.T) {
	lock, err := NewLock(context.Background(), "test-lock-id")
	assert.Nil(t, err)
	assert.NotNil(t, lock)
}

func TestNewLock_InitError(t *testing.T) {
	oldFunc := newLockFactoryFunc
	oldFactory := globalLockFactory
	newLockFactoryFunc = func() (*LockFactory, error) {
		return nil, errors.New("boom")
	}
	globalLockFactory = storage_lock_factory.NewStorageLockFactoryBeanFactory[string, any]()
	defer func() {
		newLockFactoryFunc = oldFunc
		globalLockFactory = oldFactory
	}()

	// GetLockFactory 走 initFunc 错误分支
	_, err := GetLockFactory(context.Background())
	assert.NotNil(t, err)
	// NewLock / NewLockWithOptions 的 GetLockFactory 错误分支
	_, err = NewLock(context.Background(), "x")
	assert.NotNil(t, err)
	_, err = NewLockWithOptions(context.Background(), storage_lock.NewStorageLockOptionsWithLockId("x"))
	assert.NotNil(t, err)
}

func TestNewLockWithOptions(t *testing.T) {
	options := storage_lock.NewStorageLockOptionsWithLockId("test-lock-options")
	lock, err := NewLockWithOptions(context.Background(), options)
	assert.Nil(t, err)
	assert.NotNil(t, lock)
}

func TestNewLockFactory(t *testing.T) {
	factory, err := NewLockFactory()
	assert.Nil(t, err)
	assert.NotNil(t, factory)
}

func TestGetLockFactory_Cached(t *testing.T) {
	f1, err := GetLockFactory(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, f1)
	f2, err := GetLockFactory(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, f1, f2)
}

func TestGetGlobalLockFactory(t *testing.T) {
	assert.NotNil(t, GetGlobalLockFactory())
}
