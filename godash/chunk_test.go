package godash_test

import "github.com/dacharat/go-generics/godash"

func (s *GodashTestSuite) TestChunk() {
	s.Run("invalid size", func() {
		result := godash.Chunk([]string{"A"}, 0)

		s.Nil(result)
	})

	s.Run("chunk 2", func() {
		expected := [][]string{
			{"A", "B"},
			{"C", "D"},
		}
		str := []string{"A", "B", "C", "D"}

		result := godash.Chunk(str, 2)

		s.Equal(result, expected)
	})
}
