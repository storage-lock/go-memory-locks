package mariadb_locks

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
	init, err := globalLockFactory.GetOrInit(ctx, "x", func(ctx context.Context) (*storage_lock_factory.StorageLockFactory[any], error) {
		factory, err := NewLockFactory()
		if err != nil {
			return nil, err
		}
		return factory.StorageLockFactory, nil
	})
	if err != nil {
		return nil, err
	}
	return init.CreateLock(lockId)
}

func GetGlobalLockFactory() *storage_lock_factory.StorageLockFactoryBeanFactory[string, any] {
	return globalLockFactory
}
