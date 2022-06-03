package godash_test

import "github.com/dacharat/go-generics/godash"

func (s *GodashTestSuite) TestMax() {
	s.Run("max int", func() {
		nums := []int{5, 3, 56, 1}

		result := godash.Max(nums)

		s.Equal(result, 56)
	})

	s.Run("min empty array", func() {
		nums := []int{}

		result := godash.Max(nums)

		s.Equal(result, 0)
	})
}
