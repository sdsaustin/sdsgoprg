package main

import "fmt"

func main() {

	sliceA := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	sliceB := sliceA[:6]

	fmt.Printf("%v\n%v\n", sliceA, sliceB)

	sliceB[2] = 22

	fmt.Printf("%v\n%v\n", sliceA, sliceB)
}
