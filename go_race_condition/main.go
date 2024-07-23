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

func withdraw(id int, amount int64) {
	if foundUser, exist := mapOfUser[id]; exist {
		foundUser.balance = foundUser.balance - amount
		mapOfUser[id] = foundUser
		fmt.Printf("Updated Balance ID %d Balance %d \n\n", mapOfUser[id].ID, mapOfUser[id].balance)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

    var wg sync.WaitGroup
	var mtx sync.Mutex

    for i := 0; i < 10; i++ {
        wg.Add(1)

        go func() {
            for j := 0; j < 10; j++ {
				mtx.Lock()
                withdraw(2, 100)
				mtx.Unlock()
            }

            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println(mapOfUser[2])
}