package godash_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GodashTestSuite struct {
	suite.Suite
}

func (s *GodashTestSuite) SetupTest() {}

func (s *GodashTestSuite) TearDownSuite() {}

func TestGodashSuite(t *testing.T) {
	suite.Run(t, new(GodashTestSuite))
}
