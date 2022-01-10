/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/
package main

import "fmt"

func ReverseWord(inpStr string) string{
	out := make([]rune, len(inpStr)) // To avoid memory relocation
	// using runes bc of unicode is in task
	bytes := []rune(inpStr)
	justCounter := 0
	for i := len(bytes) - 1; i >= 0; i-- {
		out[justCounter] = bytes[i]
		justCounter ++
	}
	return string(out)
}

func main(){
	test := "абвGD"
	out := ReverseWord(test)
	fmt.Println(out)
}