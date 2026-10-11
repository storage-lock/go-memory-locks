package memory_locks

import (
	"context"
	storage_lock "github.com/storage-lock/go-storage-lock"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

var (
	globalLockFactory = storage_lock_factory.NewStorageLockFactoryBeanFactory[string, any]()

	// newLockFactoryFunc 指向 NewLockFactory，包级变量以便单测注入构造失败，
	// 覆盖 GetOrInit initFunc 与 NewLock 的错误分支（默认实现永不失败，不注入则分支不可达）。
	newLockFactoryFunc = NewLockFactory
)

// NewLock 从DSN创建锁
func NewLock(ctx context.Context, lockId string) (*storage_lock.StorageLock, error) {
	init, err := GetLockFactory(ctx)
	if err != nil {
		return nil, err
	}
	return init.CreateLock(lockId)
}

func NewLockWithOptions(ctx context.Context, options *storage_lock.StorageLockOptions) (*storage_lock.StorageLock, error) {
	init, err := GetLockFactory(ctx)
	if err != nil {
		return nil, err
	}
	return init.CreateLockWithOptions(options)
}

func GetLockFactory(ctx context.Context) (*storage_lock_factory.StorageLockFactory[any], error) {
	return globalLockFactory.GetOrInit(ctx, "x", func(ctx context.Context) (*storage_lock_factory.StorageLockFactory[any], error) {
		factory, err := newLockFactoryFunc()
		if err != nil {
			return nil, err
		}
		return factory.StorageLockFactory, nil
	})
}

func GetGlobalLockFactory() *storage_lock_factory.StorageLockFactoryBeanFactory[string, any] {
	return globalLockFactory
}
