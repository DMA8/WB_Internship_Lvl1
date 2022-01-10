//Поменять местами два числа без создания временной переменной.
package main

import "fmt"

func main(){
	a := 1
	b := 0
	a, b = b, a
	fmt.Println(a, b) // syntaxSuGaR) 
}