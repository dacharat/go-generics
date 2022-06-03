package godash_test

import (
	"math"

	"github.com/dacharat/go-generics/godash"
)

func (s *GodashTestSuite) TestForEach() {
	s.Run("for each int", func() {
		expected := []float64{1, 4, 9}
		nums := []int{1, 2, 3}

		result := make([]float64, len(nums))
		godash.ForEach(nums, func(n, index int) {
			result[index] = math.Pow(float64(n), 2)
		})

		s.Equal(result, expected)
	})
}
