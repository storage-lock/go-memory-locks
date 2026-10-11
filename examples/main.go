package main

import (
	"context"
	"fmt"
	memory_locks "github.com/storage-lock/go-memory-locks"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	lockId := "test-lock"
	lock, err := memory_locks.NewLock(context.Background(), lockId)
	must(err)

	ownerId := "CC11001100"
	must(lock.Lock(context.Background(), ownerId))

	fmt.Println("safe here")

	must(lock.UnLock(context.Background(), ownerId))

}
