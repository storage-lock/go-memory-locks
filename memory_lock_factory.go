package mariadb_locks

import (
	memory_storage "github.com/storage-lock/go-memory-storage"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

type MemoryLockFactory struct {
	*storage_lock_factory.StorageLockFactory[string]
}

func NewMemoryLockFactory() (*MemoryLockFactory, error) {
	storage := memory_storage.NewMemoryStorage()
	factory := storage_lock_factory.NewStorageLockFactory[string](storage, nil)

	return &MemoryLockFactory{
		StorageLockFactory: factory,
	}, nil
}
