package godash_test

import "github.com/dacharat/go-generics/godash"

func (s *GodashTestSuite) TestIncludes() {
	s.Run("includes string", func() {
		str := []string{"A", "B", "C"}

		result := godash.Includes(str, "B")

		s.True(result)
	})

	s.Run("not include string", func() {
		str := []string{"A", "B", "C"}

		result := godash.Includes(str, "D")

		s.False(result)
	})

	s.Run("include custom struct", func() {
		type Person struct {
			Name string
			Age  int
		}

		person := Person{
			Name: "Jack",
			Age:  10,
		}

		people := []Person{
			person,
		}

		result := godash.Includes(people, person)

		s.True(result)
	})
}

func (s *GodashTestSuite) TestIncludesBy() {
	s.Run("includes odd number", func() {
		nums := []int{2, 3, 4}

		result := godash.IncludesBy(nums, func(n int) bool { return n%2 != 0 })

		s.True(result)
	})

	s.Run("not include odd number", func() {
		nums := []int{2, 4, 6}

		result := godash.IncludesBy(nums, func(n int) bool { return n%2 != 0 })

		s.False(result)
	})
}
