/*
Реализовать бинарный поиск встроенными методами языка.
*/
package main

import "fmt"

func BinarySearch(inpSlice []int, goal int) (int, bool) {
	// it works only with sorted collections!
	var left, mid, right, prevMid int
	right = len(inpSlice) - 1
	mid = (right + left) / 2
	for inpSlice[mid] != goal{ //while not found
		if inpSlice[mid] > goal{
			right = mid
		} else {
			left = mid
		}
		prevMid = mid
		mid = (right + left) / 2 + 1
		if mid == prevMid{ //if mid did'nt changed => we are stopped and didn't find
			return mid, false
		}
	}
	return mid, true
}


func main(){
	test :=[]int{1, 2, 3, 4, 5, 7, 9}
	fmt.Println(BinarySearch(test, 9))
}