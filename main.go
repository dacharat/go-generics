package main

import (
	"fmt"

	"github.com/dacharat/go-generics/godash"
)

func main() {
	run()

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("start: ", nums)
	modifiedMapNum := godash.Map(nums, func(n, _ int) int {
		return n * 2
	})

	fmt.Println("map: ", modifiedMapNum)

	modifiedFilterNum := godash.Filter(nums, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println("filter: ", modifiedFilterNum)

	modifiedReduceNum := godash.Reduce(nums, func(acc, cur int) int {
		return acc + cur
	}, 0)
	fmt.Println("reduce: ", modifiedReduceNum)

	modifiedEveryNum := godash.Every(nums, func(n int) bool { return n < 10 })
	fmt.Println("every: ", modifiedEveryNum)
}
