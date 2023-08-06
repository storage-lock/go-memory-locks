package mariadb_locks

import (
	storage_lock "github.com/storage-lock/go-storage-lock"
	"sync"
)

var GlobalMemoryLockFactory *MemoryLockFactory
var globalMemoryLockFactoryOnce sync.Once
var globalMemoryLockFactoryErr error

func InitGlobalMemoryLockFactory() error {
	factory, err := NewMemoryLockFactory()
	if err != nil {
		return err
	}
	GlobalMemoryLockFactory = factory
	return nil
}

func NewMemoryLock(lockId string) (*storage_lock.StorageLock, error) {
	globalMemoryLockFactoryOnce.Do(func() {
		globalMemoryLockFactoryErr = InitGlobalMemoryLockFactory()
	})
	if globalMemoryLockFactoryErr != nil {
		return nil, globalMemoryLockFactoryErr
	}
	return GlobalMemoryLockFactory.CreateLock(lockId)
}
