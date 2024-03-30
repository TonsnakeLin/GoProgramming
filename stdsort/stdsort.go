// test code for std lib sort
package stdsort

import (
	"fmt"
	"sort"
)

func Test_sort_returns() {
	input := []int{1, 5, 10, 15, 20, 25}
	i := sort.Search(len(input), func(i int) bool {
		return input[i] > 10
	})
	fmt.Println("The pos of > 10: ", i)

	i = sort.Search(len(input), func(i int) bool {
		return input[i] > 30
	})
	fmt.Println("The pos of > 30: ", i)

	i = sort.Search(len(input), func(i int) bool {
		return input[i] > 0
	})
	fmt.Println("The pos of > 0: ", i)
}
