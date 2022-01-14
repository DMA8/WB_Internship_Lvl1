/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(22+32+42….) с использованием конкурентных вычислений.
*/
package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	allNumbersSum := 0
	a := []int{2, 4, 6, 8, 10}
	//tubeWithSquares := make(chan int, len(a)) 
	//When using buffered chan, goroutine doesn't block until buffChan is full
	tubeWithSquares := make(chan int) 
	//when there's something in channel, goroutine is blocked until some routine gets data from channel

	wg.Add(1) // wgCounter++
	go func(){
		for _, v := range a {
			tubeWithSquares <- v * v // Putting square of V in chan
		}
		close(tubeWithSquares) // When channel is no more supposed to be filled - close it
		wg.Done() //wgCounter--
	}()

	wg.Add(1)
	go func(){
		for value := range tubeWithSquares {
			allNumbersSum += value // No need to save it with mutex cause there's only 1 goroutine that works with AllNumberSum
		}
		wg.Done()
	}()

	wg.Wait() // waits until wgCounter == 0
	fmt.Println(allNumbersSum)
}
