# Memory Locks

# 一、这是什么

Memory Locks基本没有实际作用，更多的是用于在框架内部作为测试用，补充流程完整使流程能够走下去。

# 二、安装

```bash
go get -u github.com/storage-lock/go-memory-locks
```

# 三、示例

```go
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
```

