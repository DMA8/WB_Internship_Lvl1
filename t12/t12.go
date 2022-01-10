// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество
package main

import "fmt"

func getUniqueSet(words []string) []string{
	var uniqueSet []string
	uniqueMap := make(map[string] int)

	for _, v := range words {
		if _, ok := uniqueMap[v]; ok == false {
			uniqueMap[v] = 1
		}
	}
	for key, _ := range uniqueMap{
		uniqueSet = append(uniqueSet, key)
	}
	return uniqueSet
}

func main(){
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(getUniqueSet(words))
}