package main

import "fmt"

func main() {
	a := []int{5, 1, 2, 5}
	b := []int{4, 2, 5, 1, 1, 2}

	// The unique values
	uniqueA := getUniqueValues(a)
	uniqueB := getUniqueValues(b)
	fmt.Println("The unique values of slice a: ", uniqueA)
	fmt.Println("The unique values of slice b: ", uniqueB)

	// The same values
	same := getSameValues(a, b)
	fmt.Println("The same values of two slices: ", same)

	// The union unique values
	unionUnique := getUnionUniqueValues(a, b)
	fmt.Println("The union unique values of two slices: ", unionUnique)
}

func getUniqueValues(slice []int) []int {
	tempUnique := make(map[int]bool)
	var unique []int

	for _, value := range slice {
		if _, ok := tempUnique[value]; !ok {
			tempUnique[value] = true
			unique = append(unique, value)
		}
	}
	return unique
}

func getSameValues(a, b []int) []int {
	same := make([]int, 0)
	tempA := make(map[int]bool)
	tempB := make(map[int]bool)

	for _, value := range a {
		tempA[value] = true
	}

	for _, value := range b {
		if tempA[value] && !tempB[value] {
			same = append(same, value)
			tempB[value] = true
		}
	}
	return same
}

func getUnionUniqueValues(a, b []int) []int {
	unionUnique := make([]int, 0)
	temp := make(map[int]bool)

	for _, value := range a {
		if !temp[value] {
			temp[value] = true
			unionUnique = append(unionUnique, value)
		}
	}

	for _, value := range b {
		if !temp[value] {
			temp[value] = true
			unionUnique = append(unionUnique, value)
		}
	}
	return unionUnique
}
