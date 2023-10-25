package syntax

import (
	"fmt"
)

// add two slice together
func AppendTwoIntSlice(slice1 []int, slice2 []int) (returnslice []int) {
	returnslice = append(slice1, slice2...)
	for i := 0; i < len(returnslice); i++ {
		fmt.Printf("index %d from slice returnslice is %v \n", i, returnslice[i])
	}
	fmt.Printf("length of new slice is %v , and capacity of it is %v \n", len(returnslice), cap(returnslice))
	return returnslice
}

func OptimizeSlice( a int[]{...}){
	targetSlice:=make([]int ,0)

}

func AddSomElemanToSlice(slice []int, a int) []int {
	slice2 := append(slice, a)
	fmt.Printf("%d added to slice \n", a)
	return slice2
}

// Copy Slice


// for 

