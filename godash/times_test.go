package godash_test

import (
	"fmt"

	"github.com/dacharat/go-generics/godash"
)

func (s *GodashTestSuite) TestTimes() {
	s.Run("times string", func() {
		expected := []string{"1", "2", "3"}

		result := godash.Times(3, func(index int) string {
			return fmt.Sprintf("%d", index+1)
		})

		s.Equal(result, expected)
	})

	s.Run("times equal 0", func() {
		result := godash.Times(0, func(index int) string {
			return fmt.Sprintf("%d", index)
		})

		s.Nil(result)
	})
}
