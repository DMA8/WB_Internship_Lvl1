//Реализовать собственную функцию sleep.
package main

import (
	"fmt"
	"time"
)

func mySleep(msDuration int) {
	end := time.Now().Add(time.Duration(msDuration * int(time.Millisecond))) // when we shoudl awake
	fmt.Println("I am going to sleep")
	for time.Now().Before(end){ // in empty loop while time.Now() less then timestamp end
	}
	fmt.Println("I am awake") // end; Let main works further
}

func main(){
	fmt.Println(time.Now())
	mySleep(3000)
	fmt.Println(time.Now())
}
