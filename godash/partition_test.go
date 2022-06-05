package godash_test

import "github.com/dacharat/go-generics/godash"

type partitionTestCase[T any] struct {
	name      string
	input     []T
	fun       func(T) bool
	expected1 []T
	expected2 []T
}

func (s *GodashTestSuite) TestPartitionInt() {
	testcases := []partitionTestCase[int]{
		{
			name:      "partition odd and even number",
			input:     []int{1, 2, 3, 4, 5, 6, 7, 8},
			fun:       func(i int) bool { return i%2 == 0 },
			expected1: []int{2, 4, 6, 8},
			expected2: []int{1, 3, 5, 7},
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			result1, result2 := godash.Partition(tc.input, tc.fun)

			s.Equal(result1, tc.expected1)
			s.Equal(result2, tc.expected2)
		})
	}
}

func (s *GodashTestSuite) TestPartitionPerson() {
	testcases := []partitionTestCase[int]{
		{
			name:      "partition odd and even number",
			input:     []int{1, 2, 3, 4, 5, 6, 7, 8},
			fun:       func(i int) bool { return i%2 == 0 },
			expected1: []int{2, 4, 6, 8},
			expected2: []int{1, 3, 5, 7},
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			result1, result2 := godash.Partition(tc.input, tc.fun)

			s.Equal(result1, tc.expected1)
			s.Equal(result2, tc.expected2)
		})
	}
}
