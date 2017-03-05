package hyperdrive

import (
	"net/http"
	"os"
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

func (suite *HyperdriveTestSuite) TestAPIServer() {
	suite.IsType(&http.Server{}, NewAPI().Server, "expects an instance of *http.Server")
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

func (suite *HyperdriveTestSuite) TestNewConfig() {
	suite.IsType(Config{}, NewConfig(), "expects an instance of *hyperdrive.config")
}

func (suite *HyperdriveTestSuite) TestPortConfigFromDefault() {
	c := NewConfig()
	suite.Equal(5000, c.Port, "Port should be equal to default value")
}

func (suite *HyperdriveTestSuite) TestPortConfigFromEnv() {
	os.Setenv("PORT", "5001")
	c := NewConfig()
	suite.Equal(5001, c.Port, "Port should be equal to PORT value set via ENV var")
}

func (suite *HyperdriveTestSuite) TestGetPort() {
	c := NewConfig()
	suite.Equal(":5000", c.GetPort(), "c.Port value should be return, prefixed with a colon, e.g. :5000")
}

func (suite *HyperdriveTestSuite) TestEnvConfigFromDefault() {
	c := NewConfig()
	suite.Equal("development", c.Env, "Env should be equal to default value")
}

func (suite *HyperdriveTestSuite) TestEnvConfigFromEnv() {
	os.Setenv("HYPERDRIVE_ENVIRONMENT", "test")
	c := NewConfig()
	suite.Equal("test", c.Env, "Env should be equal to HYPERDRIVE_ENVIRONMENT value set via ENV var")
}

func TestHyperdriveTestSuite(t *testing.T) {
	suite.Run(t, new(HyperdriveTestSuite))
}
