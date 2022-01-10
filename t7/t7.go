/*
Реализовать конкурентную запись данных в map
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func touchMap(mux *sync.Mutex, wg *sync.WaitGroup, myMap map[int] int, value *int) {
	defer wg.Done()
	mux.Lock() // panic if don't use mutex with many workers
	myMap[*value] = rand.Intn(100)
	*value += 1
	mux.Unlock()
	time.Sleep(time.Second * 1)
}


func main() {
	var mux sync.Mutex
	var wg sync.WaitGroup
	myMap := make(map[int] int)
	var counter int
	nWorkers := 987987
	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go touchMap(&mux, &wg, myMap, &counter)
	}
	wg.Wait()
	fmt.Println("Everything OK? ", nWorkers == len(myMap)) //simple check
}