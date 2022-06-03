package godash_test

import (
	"github.com/dacharat/go-generics/godash"
)

func (s *GodashTestSuite) TestFind() {
	s.Run("find and found", func() {
		nums := []int{1, 2, 3, 4, 5}

		result, found := godash.Find(nums, func(n int) bool { return n > 2 })

		s.Equal(result, 3)
		s.Equal(found, true)
	})

	s.Run("find and not found", func() {
		type Person struct {
			Name string
			Age  int
		}
		person := []Person{
			{
				Name: "Jack",
				Age:  18,
			},
			{
				Name: "Tuu",
				Age:  44,
			},
		}

		emptyPerson := Person{}
		result, found := godash.Find(person, func(p Person) bool { return p.Age > 112 })

		s.Equal(result, emptyPerson)
		s.Equal(found, false)
	})
}

func (s *GodashTestSuite) TestFindIndex() {
	s.Run("find and found", func() {
		nums := []int{1, 2, 3, 4, 5}

		result := godash.FindIndex(nums, func(n int) bool { return n > 2 })

		s.Equal(result, 2)
	})

	s.Run("find and not found", func() {
		type Person struct {
			Name string
			Age  int
		}
		person := []Person{
			{
				Name: "Jack",
				Age:  18,
			},
			{
				Name: "Tuu",
				Age:  44,
			},
		}

		result := godash.FindIndex(person, func(p Person) bool { return p.Age > 112 })

		s.Equal(result, -1)
	})
}
