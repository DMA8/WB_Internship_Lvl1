/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(22+32+42….) с использованием конкурентных вычислений.
*/
package main

import (
	"fmt"
	"sync"
)

func AddSquare(mux *sync.Mutex, wg *sync.WaitGroup, r *int, value int){
	defer  wg.Done()
	mux.Lock() 
	//RaceCondition Protection
	*r += value * value
	mux.Unlock()
	
}

func main(){
	var wg sync.WaitGroup
	var mux sync.Mutex 
	a := []int{2, 4, 6, 8, 10}
	result := 0
	wg.Add(len(a))
	for _, i := range a { 
		// we run len(a) goRoutines here
		// and each of it wants to change result simultaneously
		go AddSquare(&mux, &wg, &result, i)
	}
	wg.Wait()
	fmt.Println("Result is : ",result)
}
