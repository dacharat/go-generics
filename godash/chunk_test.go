package godash_test

import (
	"github.com/dacharat/go-generics/godash"
)

type chunkTestCase[T any] struct {
	name     string
	input    []T
	size     int
	expected [][]T
}

func (s *GodashTestSuite) TestChunkString() {
	testcases := []chunkTestCase[string]{
		{
			name:     "invalid size",
			input:    []string{"A"},
			expected: nil,
		},
		{
			name:  "chunk 2",
			input: []string{"A", "B", "C", "D"},
			size:  2,
			expected: [][]string{
				{"A", "B"},
				{"C", "D"},
			},
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			result := godash.Chunk(tc.input, tc.size)

			s.Equal(result, tc.expected)
		})
	}
}

func (s *GodashTestSuite) TestChunkInt() {
	testcases := []chunkTestCase[int]{
		{
			name:     "invalid size",
			input:    []int{1},
			expected: nil,
		},
		{
			name:  "chunk 2",
			input: []int{1, 2, 3, 4},
			size:  2,
			expected: [][]int{
				{1, 2},
				{3, 4},
			},
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			result := godash.Chunk(tc.input, tc.size)

			s.Equal(result, tc.expected)
		})
	}
}
