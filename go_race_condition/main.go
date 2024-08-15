package main

import (
	"fmt"
	"runtime"
	"sync"
)

type User struct {
	ID      int
	balance int64
}

var mapOfUser map[int]User

func init() {
	mapOfUser = make(map[int]User)
	for i := 1; i < 100; i++ {
		mapOfUser[i] = User{
			ID:      i,
			balance: 10000000,
		}
	}
}

func withdraw(id int, amount int64, mtx *sync.Mutex) {
	mtx.Lock()
	if foundUser, exist := mapOfUser[id]; exist {
		foundUser.balance = foundUser.balance - amount
		mapOfUser[id] = foundUser

	}
	mtx.Unlock()
}

func main() {
	runtime.GOMAXPROCS(2)
	var mtx sync.Mutex

    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)

        go func() {
			withdraw(2, 100, &mtx)
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println(mapOfUser[2])
}