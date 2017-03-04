package hyperdrive

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HyperdriveTestSuite struct {
	suite.Suite
	Endpoint Endpointer
}

func (suite *HyperdriveTestSuite) SetupTest() {
	suite.Endpoint = NewEndpoint("Test", "Test Endpoint", "/test")
}

func (suite *HyperdriveTestSuite) TestNewAPI() {
	suite.IsType(API{}, NewAPI(), "expects an instance of hyperdrive.API")
}

func (suite *HyperdriveTestSuite) TestNewEndpoint() {
	suite.IsType(&Endpoint{}, suite.Endpoint, "expects an instance of hyperdrive.Endpoint")
}

func (suite *HyperdriveTestSuite) TestGetName() {
	suite.Equal("Test", suite.Endpoint.GetName(), "expects GetName() to return Name")
}

func (suite *HyperdriveTestSuite) TestGetDesc() {
	suite.Equal("Test Endpoint", suite.Endpoint.GetDesc(), "expects GetDesc() to return Desc")
}

func (suite *HyperdriveTestSuite) TestGetPath() {
	suite.Equal("/test", suite.Endpoint.GetPath(), "expects GetPath() to return Path")
}

func (suite *HyperdriveTestSuite) TestEndpointer() {
	suite.Implements((*Endpointer)(nil), suite.Endpoint, "expects an implementation of hyperdrive.Endpointer interface")
}

func TestHyperdriveTestSuite(t *testing.T) {
	suite.Run(t, new(HyperdriveTestSuite))
}
