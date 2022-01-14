/*Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

package main

import (
	"fmt"
)

type Human struct {
	Name		string
	Age			int
	LikesBeer	bool
}

//Check if Human probably consumes beer
func (h Human)CheckBeerConsumer() bool{ 
	if h.Age >= 18 && h.LikesBeer {
		return true
	}
	return false
}

//Struct for demonstration of embedding 
type Action struct {
	Region	string
	Human
}

func main(){
	test := Action{"Brighton", Human{"TestName", 19, true}}
	fmt.Println(test.CheckBeerConsumer())
}
