package godash_test

import (
	"github.com/dacharat/go-generics/godash"
)

func (s *GodashTestSuite) TestReduce() {
	s.Run("reduce int", func() {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

		result := godash.Reduce(nums, func(acc, cur int) int { return acc + cur }, 0)

		s.Equal(result, 36)
	})

	s.Run("reduce array to map[string]value", func() {
		type Reward struct {
			Type     string
			Title    string
			Discount float64
		}

		reward1 := Reward{
			Type:     "FIXED",
			Title:    "Discount 10฿",
			Discount: 10,
		}
		reward2 := Reward{
			Type:     "RATIO",
			Title:    "Discount 10%",
			Discount: 10,
		}
		expect := map[string]Reward{
			"FIXED": reward1,
			"RATIO": reward2,
		}
		rewards := []Reward{
			reward1,
			reward2,
		}

		result := godash.Reduce(rewards,
			func(acc map[string]Reward, r Reward) map[string]Reward {
				acc[r.Type] = r
				return acc
			},
			map[string]Reward{})

		s.Equal(result, expect)
	})
}
