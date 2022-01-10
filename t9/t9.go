/*Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var Wg sync.WaitGroup
	A := [10]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	firstChan := make(chan int, len(A))
	secondChan := make(chan int, len(A))
	// firstChan := make(chan int)
	// secondChan := make(chan int)
	time1 := time.Now()
	Wg.Add(3)
	go func(){
		for _, v := range A {
			firstChan <- v
			time.Sleep(time.Millisecond * 30)
		}
		Wg.Done()
		close(firstChan)
		fmt.Println("FIRST DONE")
	}()
	go func(){
		for valueFromFirstChan := range firstChan {
			time.Sleep(time.Millisecond * 200)
			secondChan <- valueFromFirstChan * valueFromFirstChan
		}
		Wg.Done()
		close(secondChan)
		fmt.Println("Second Done")
	}()
	go func(){
		for valueFromSecondChan := range secondChan {
			time.Sleep(time.Millisecond * 1000)

			fmt.Println("This is what I got in secondChan ", valueFromSecondChan)
		}		
		Wg.Done()
		fmt.Println("Third Done")
	}()
	Wg.Wait()
	dur := time.Since(time1)
	fmt.Println(dur)
}