/*Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
package main

import (
	"fmt"
	"strings"
)
func ReverseWordsString(inputString string) string{
	spltInp := strings.Split(inputString, " ")
	var outString string
	for i := len(spltInp) - 1; i >= 0; i--{
		outString += spltInp[i] + " "
	}
	return outString
}

func main(){
	test := "first second third четвертый"
	reversed := ReverseWordsString(test)
	fmt.Println(reversed)
}
