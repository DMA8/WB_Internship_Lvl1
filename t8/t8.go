/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в
1 или 0.
*/
package main

import "fmt"
//https://tproger.ru/articles/awesome-bits/

// ^ - искл. ИЛИ
// &^ - И-НЕ


func SetNBit(num, position, value int64) int64 {
	if value >= 1{
		value = 1
	} else {
		value = 0
	}
	if value == 1{
		mask := int64(1) << position
		//fmt.Printf("Mask when settin N bit to N: %08b\n", mask)
		return num | mask // bitwise OR
	}
	return num &^ (int64(1) << position) // XOR
}
func main(){
	var a int64 = 0
	fmt.Printf("Before transformation: %08b\n", a)
	fmt.Printf("After Transformation: %08b\n", SetNBit(a, 4, 1))
}