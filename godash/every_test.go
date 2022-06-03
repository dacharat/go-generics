package godash_test

import "github.com/dacharat/go-generics/godash"

func (s *GodashTestSuite) TestEvery() {
	s.Run("every are true", func() {
		nums := []int{2, 4, 6}

		result := godash.Every(nums, func(n int) bool { return n%2 == 0 })

		s.Equal(result, true)
	})

	s.Run("something false", func() {
		nums := []int{2, 3, 6}

		result := godash.Every(nums, func(n int) bool { return n%2 == 0 })

		s.Equal(result, false)
	})
}
