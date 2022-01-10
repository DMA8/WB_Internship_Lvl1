// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.
package main

import (
	"fmt"
	"math/big"
)

func main(){
	// Big int becomes useful with numbers > 2^62
	a := new(big.Int).Mul(big.NewInt(1<<62), big.NewInt(1<<62))
	b := new(big.Int).Mul(big.NewInt(1<<61), big.NewInt(1<<62))
	c := new(big.Int).Add(a, b)
	fmt.Printf("%T %d + %d = %d\n", c, a, b, c)
	c = new(big.Int).Mul(a, b)
	fmt.Println("Multiplication ", c)
	c = new(big.Int).Div(a, b)
	fmt.Println("Division ", c)
	c = new(big.Int).Sub(a, b)
	fmt.Println("Submition ", c)
}
