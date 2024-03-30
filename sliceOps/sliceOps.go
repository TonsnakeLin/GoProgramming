// This is for regular operation of slice
package sliceops

import (
	"fmt"
	"sort"
)

func Test_slice_ops1() {
	fmt.Println("**********Test_slice_ops1**********")
	input := []int{20, 50, 40, 30, 30, 10, 60}
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})
	fmt.Println(input)
	sortItems := make([]int, 0, 10)
	for i, v := range input {
		sortItems = append(sortItems, v)
		if i == len(input)-1 || v != input[i+1] {
			tmp := append(sortItems[:0:0], sortItems...)
			fmt.Println("tmp: ", tmp, "sortItems: ", sortItems)
		}
	}
}
