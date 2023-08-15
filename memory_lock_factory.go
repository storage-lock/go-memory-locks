package mariadb_locks

import (
	memory_storage "github.com/storage-lock/go-memory-storage"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

type LockFactory struct {
	*storage_lock_factory.StorageLockFactory[any]
}

// NewLockFactory 从DSN创建锁
func NewLockFactory() (*LockFactory, error) {
	storage := memory_storage.NewMemoryStorage()
	factory := storage_lock_factory.NewStorageLockFactory[any](storage, nil)
	return &LockFactory{
		StorageLockFactory: factory,
	}, nil
}
