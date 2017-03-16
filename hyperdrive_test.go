package hyperdrive

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type HyperdriveTestSuite struct {
	suite.Suite
	TestAPI              API
	TestEndpoint         Endpointer
	TestHandler          http.Handler
	TestRoot             *RootResource
	TestEndpointResource EndpointResource
	TestGetRequest       *http.Request
}

func (suite *HyperdriveTestSuite) SetupTest() {
	suite.TestAPI = NewAPI("API", "Test API Desc")
	suite.TestEndpoint = NewEndpoint("Test", "Test Endpoint", "/test", "1.0.1-beta")
	suite.TestHandler = NewMethodHandler(suite.TestEndpoint)
	suite.TestRoot = NewRootResource(suite.TestAPI)
	suite.TestEndpointResource = NewEndpointResource(suite.TestEndpoint)
	suite.TestGetRequest = httptest.NewRequest("GET", "/test/2?id=1&a=b", nil)
}

func (suite *HyperdriveTestSuite) TestNewAPI() {
	suite.IsType(API{}, suite.TestAPI, "expects an instance of hyperdrive.API")
}

func (suite *HyperdriveTestSuite) TestAPIServer() {
	suite.IsType(&http.Server{}, suite.TestAPI.Server, "expects an instance of *http.Server")
}

func TestHyperdriveTestSuite(t *testing.T) {
	suite.Run(t, new(HyperdriveTestSuite))
}
