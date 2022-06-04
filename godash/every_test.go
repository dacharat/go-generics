package godash_test

import "github.com/dacharat/go-generics/godash"

type everyTestCase[T any] struct {
	name     string
	input    []T
	fun      func(T) bool
	expected bool
}

func (s *GodashTestSuite) TestEveryInt() {
	testcases := []everyTestCase[int]{
		{
			name:     "every are true",
			input:    []int{2, 4, 6},
			fun:      func(n int) bool { return n%2 == 0 },
			expected: true,
		},
		{
			name:     "something false",
			input:    []int{2, 3, 6},
			fun:      func(n int) bool { return n%2 == 0 },
			expected: false,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			result := godash.Every(tc.input, tc.fun)

			s.Equal(result, tc.expected)
		})
	}
}

func (s *GodashTestSuite) TestEveryString() {
	testcases := []everyTestCase[string]{
		{
			name:     "every are true",
			input:    []string{"123", "45", "6"},
			fun:      func(s string) bool { return len(s) < 4 },
			expected: true,
		},
		{
			name:     "something false",
			input:    []string{"123", "45", "6"},
			fun:      func(s string) bool { return len(s) == 0 },
			expected: false,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.name, func() {
			result := godash.Every(tc.input, tc.fun)

			s.Equal(result, tc.expected)
		})
	}
}
