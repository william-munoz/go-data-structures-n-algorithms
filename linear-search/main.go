package main

import "fmt"

func linearsearch(datalist []int, key int) (bool, int) {
	for pos, item := range datalist {
		if item == key {
			return true, pos
		}
	}
	return false, -1
}

func main() {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	isPresent, pos := linearsearch(items, 58)
	if isPresent {
		fmt.Printf("The value is present in the position %d", pos)
	} else {
		fmt.Println("The value is not present.s")
	}
}
