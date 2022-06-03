package godash_test

import (
	"fmt"

	"github.com/dacharat/go-generics/godash"
)

func (s *GodashTestSuite) TestMap() {
	s.Run("map int", func() {
		expect := []int{2, 4, 6, 8, 10}
		ints := []int{1, 2, 3, 4, 5}

		result := godash.Map(ints, func(i, _ int) int { return i * 2 })

		s.Equal(result, expect)
	})

	s.Run("map float", func() {
		expect := []float64{2.0, 4.0, 6.0, 8.0, 10.0}
		floats := []float64{1, 2, 3, 4, 5}

		result := godash.Map(floats, func(i float64, _ int) float64 { return i * 2 })

		s.Equal(result, expect)
	})

	s.Run("map int to string", func() {
		expect := []string{"0_1", "1_2", "2_3", "3_4", "4_5"}
		str := []int{1, 2, 3, 4, 5}

		result := godash.Map(str, func(i int, index int) string { return fmt.Sprintf("%d_%d", index, i) })

		s.Equal(result, expect)
	})
}
