//Реализовать пересечение двух неупорядоченных множеств.
package main

import "fmt"


func innerJoin(a, b []int) []int {
	map1 := make(map[int]int)
	map2 := make(map[int]int)
	//O(n) = len(a) + len(b) + min(len(map1), len(map2)) Худш случай - 3n; Лучший случай - 2n 
	var inner []int
	for _, v := range(a) {
		map1[v] = 1
	}
	for _, v := range(b) {
		map2[v] = 1
	}
	if len(map1) > len(map2){
		for key, _ := range(map2){
			if map2[key] == 1 &&  map1[key] == 1{
				inner = append(inner, key)
			}
		}
	} else {
		for key, _ := range(map1){
			if map1[key] == 1 &&  map2[key] == 1{
				inner = append(inner, key)
			}
		}
	}
	return inner
}

func main(){
	a := []int{5,4,7,1,2,3}
	b := []int{9,8,1,4,2,5,6,3,4,5,7,23,4,3,2}
	c := innerJoin(a, b)
	fmt.Println(c)
}