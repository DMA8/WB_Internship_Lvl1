//Разработать программу, которая в рантайме способна определить тип
//переменной: int, string, bool, channel из переменной типа interface{}.

package main

import "fmt"

func IKnowYourType(a interface{}){ //empty interface allows accept everything as a parameter
	switch a.(type) { //get type of content
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("This is int")
	case bool:
		fmt.Println("bool FOUND")
	default:
			fmt.Println("WHO ARE YOU")
	}
}

func main(){
	IKnowYourType(true)
}