package godash_test

import (
	"github.com/dacharat/go-generics/godash"
)

func (s *GodashTestSuite) TestUniq() {
	s.Run("unique int", func() {
		expected := []int{1, 3, 5, 4, 6}
		nums := []int{1, 3, 5, 3, 4, 5, 6}

		result := godash.Uniq(nums)

		s.Equal(result, expected)
	})
}
