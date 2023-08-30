package memory_locks

import (
	"context"
	storage_lock "github.com/storage-lock/go-storage-lock"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

var (
	globalLockFactory = storage_lock_factory.NewStorageLockFactoryBeanFactory[string, any]()
)

// NewLock 从DSN创建锁
func NewLock(ctx context.Context, lockId string) (*storage_lock.StorageLock, error) {
	init, err := GetLockFactory(ctx)
	if err != nil {
		return nil, err
	}
	return init.CreateLock(lockId)
}

func NewLockWithOptions(ctx context.Context, lockId string, options *storage_lock.StorageLockOptions) (*storage_lock.StorageLock, error) {
	init, err := GetLockFactory(ctx)
	if err != nil {
		return nil, err
	}
	return init.CreateLockWithOptions(lockId, options)
}

func GetLockFactory(ctx context.Context) (*storage_lock_factory.StorageLockFactory[any], error) {
	return globalLockFactory.GetOrInit(ctx, "x", func(ctx context.Context) (*storage_lock_factory.StorageLockFactory[any], error) {
		factory, err := NewLockFactory()
		if err != nil {
			return nil, err
		}
		return factory.StorageLockFactory, nil
	})
}

func GetGlobalLockFactory() *storage_lock_factory.StorageLockFactoryBeanFactory[string, any] {
	return globalLockFactory
}
