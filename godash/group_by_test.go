package godash_test

import "github.com/dacharat/go-generics/godash"

func (s *GodashTestSuite) TestGroupBy() {
	s.Run("group int with mod 3", func() {
		expected := map[int][]int{
			0: {3, 6, 9},
			1: {1, 4, 7},
			2: {2, 5, 8},
		}
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		result := godash.GroupBy(nums, func(n int) int {
			return n % 3
		})

		s.Equal(result, expected)
	})
}
