package main

import (
	"context"
	"fmt"
	memory_locks "github.com/storage-lock/go-memory-locks"
)

func main() {

	lockId := "test-lock"
	lock, err := memory_locks.NewLock(context.Background(), lockId)
	if err != nil {
		panic(err)
	}

	ownerId := "CC11001100"
	err = lock.Lock(context.Background(), ownerId)
	if err != nil {
		panic(err)
	}

	fmt.Println("safe here")

	err = lock.UnLock(context.Background(), ownerId)
	if err != nil {
		panic(err)
	}

}
