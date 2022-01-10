/*Реализовать быструю сортировку массива (quicksort) встроенными методами
языка*/
package main

import (
	"fmt"
)

func Qsort(a []int, left, rigth int) {
	if rigth - left > 0 {
		pivotInd := partition(a, left, rigth)
		Qsort(a, left, pivotInd - 1)
		Qsort(a, pivotInd + 1, rigth)
	}
}

func partition(a []int, left, rigth int) int {
	var i, pivot, wall int
	pivot = rigth // pivot is the last element of the array
	wall = left // Put all elemens less than pivot to the left of the wall index and move the wall to the right 1 step
	for i = left; i < rigth; i++ {
		if a[i] < a[pivot] { // if array elem < pivot => put it behind the wall
			a[i], a[wall] = a[wall], a[i] // swap
			wall++ // move the wall to the rigth
		}
	}
	a[pivot], a[wall] = a[wall], a[pivot] // all elements that are greater than pivot
	// is to the right of the wall => swap to wall element
	return wall
}

func main() {
	a := []int{5,2,7,1,4,9,3,6,8}
	Qsort(a, 0, len(a) - 1)
	fmt.Println(a)
}