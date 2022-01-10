// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func main(){
	wg := new(sync.WaitGroup) // Creating queue of goroutines in order to wait its finish until main completes
	nums := [5]int{2, 4, 6, 8, 10}
	for _, i := range nums{
		wg.Add(1) // add 1 to counter of Goroutines
		go func(a int){
			defer wg.Done() // extract 1 from counter of Goroutines
			fmt.Println(a * a) // payload
		}(i) // calling anonymous func with "i" as argument
	}
	wg.Wait() // wait until wg.Counter is not 0
}
