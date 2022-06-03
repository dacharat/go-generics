package godash_test

import (
	"github.com/dacharat/go-generics/godash"
)

func (s *GodashTestSuite) TestFilter() {
	s.Run("filter int", func() {
		expect := []int{2, 4}
		ints := []int{1, 2, 3, 4, 5}

		result := godash.Filter(ints, func(i int) bool { return i%2 == 0 })

		s.Equal(result, expect)
	})

	s.Run("filter float", func() {
		expect := []float64{3.0, 4.0, 5.0}
		floats := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

		result := godash.Filter(floats, func(i float64) bool { return i > 2.5 })

		s.Equal(result, expect)
	})

}
