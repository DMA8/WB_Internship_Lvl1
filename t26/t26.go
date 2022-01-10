/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CheckAllCharsUnique(inputString string) bool{
	uniqueChars := make(map[rune]int)
	inputString = strings.ToLower(string(inputString)) // Don't check case of the char
	for _, char := range inputString {
		if _, ok := uniqueChars[rune(char)]; ok { // if there's not unique char -> return false
			return false
		} else {
			uniqueChars[rune(char)] = 1
		}
	}
	return true
}

func main(){
	byteString, err := ioutil.ReadAll(os.Stdin) // Read stdin until Ctrl+D
	if err != nil{
		log.Fatal(err)
	}
	inputString := string(byteString) // convetToString
	inputString = strings.TrimRight(inputString, "\n") // Delete \n from the right (bc what if we have 1 \n in string)
	fmt.Println(CheckAllCharsUnique(inputString))
}
