package godash_test

import "github.com/dacharat/go-generics/godash"

func (s *GodashTestSuite) TestMin() {
	s.Run("min int", func() {
		nums := []int{5, 3, 56, 1}

		result := godash.Min(nums)

		s.Equal(result, 1)
	})

	s.Run("min empty array", func() {
		nums := []int{}

		result := godash.Min(nums)

		s.Equal(result, 0)
	})
}
