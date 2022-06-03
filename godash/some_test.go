package godash_test

import "github.com/dacharat/go-generics/godash"

func (s *GodashTestSuite) TestSome() {
	s.Run("every are true", func() {
		nums := []int{2, 3, 5}

		result := godash.Some(nums, func(n int) bool { return n%2 == 0 })

		s.True(result)
	})

	s.Run("all false", func() {
		nums := []int{3, 5, 7}

		result := godash.Some(nums, func(n int) bool { return n%2 == 0 })

		s.False(result)
	})
}
