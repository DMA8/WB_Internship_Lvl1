/*К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.
var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}
func main() {
someFunc()
}*/
package main

import (
	"fmt"
	"unsafe"
)

func createHugeString(size int) *[]byte {
	v := make([]byte, size)
	for i := 0; i < size; i++ {
		v[i] = 'a'
	}
	return &v //DOESN'T GIVE ANY PROFIT WHEN WORKING WITH STRINGS, BC ALL STRINGS ARE LOCATED IN HEAP
}

func createHugeStringOrig(size int) string {
	var str string
	for i := 0; i < size; i++ {
		str += "а"
	}
	return str
}

func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Println("Ptr size : ",unsafe.Sizeof(v))
	//justStringLight := string((*v)[:100])
	heavyString := createHugeStringOrig(1 << 10)
	runeSlice := []rune(heavyString)
	niceJustString := string(runeSlice[:99])
	justString := heavyString[:99] // If string is Unicode, you can't be sure how many symbols you sliced 
	fmt.Println(justString)
	fmt.Println(niceJustString)
}

func main() {
	someFunc()
}