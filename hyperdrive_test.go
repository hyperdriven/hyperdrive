package hypderdrive

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HyperdriveTestSuite struct {
	suite.Suite
}

func (suite *HyperdriveTestSuite) TestNewAPI() {
	suite.IsType(API{}, NewAPI(), "expects an instance of hyperdrive.API")
}

func TestHyperdriveTestSuite(t *testing.T) {
	suite.Run(t, new(HyperdriveTestSuite))
}
