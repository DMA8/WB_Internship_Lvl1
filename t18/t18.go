/*Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.*/
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	order int
}

func incrementCounter(ctx context.Context, mux *sync.Mutex, wg *sync.WaitGroup,c *Counter) {
	for {
		select{
		case <- ctx.Done():
			defer wg.Done()
			return
		default:
			mux.Lock()
			c.order += 1
			mux.Unlock()
			time.Sleep(time.Millisecond*42)
		}
	}
}


func main(){
	var wg sync.WaitGroup
	var mux sync.Mutex
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 10)

	a := Counter{}
	for i := 0; i < 100; i++ {
		wg.Add(1) //incr wg Counter
		go incrementCounter(ctx, &mux, &wg, &a) // struct is given as a pointer
	}
	wg.Wait() // blocks until counter is not 0
	fmt.Println("Counted up to :", a.order)
}
