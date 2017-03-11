package hyperdrive

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type HyperdriveTestSuite struct {
	suite.Suite
	TestAPI                    API
	TestEndpoint               Endpointer
	TestHandler                http.Handler
	TestRoot                   *RootResource
	TestRootRepresentation     Representation
	TestEndpointRepresentation Representation
}

func (suite *HyperdriveTestSuite) SetupTest() {
	suite.TestAPI = NewAPI("API", "Test API Desc")
	suite.TestEndpoint = NewEndpoint("Test", "Test Endpoint", "/test", "1.0.1")
	suite.TestHandler = NewMethodHandler(suite.TestEndpoint)
	suite.TestRoot = NewRootResource(suite.TestAPI)
	suite.TestEndpointRepresentation = Representation{"name": "Test", "desc": "Test Endpoint", "path": "/test", "methods": []string{"OPTIONS"}}
	suite.TestRootRepresentation = Representation{"resource": "api", "name": "API", "endpoints": []Representation{suite.TestEndpointRepresentation}}
}

func (suite *HyperdriveTestSuite) TestNewAPI() {
	suite.IsType(API{}, suite.TestAPI, "expects an instance of hyperdrive.API")
}

func (suite *HyperdriveTestSuite) TestAPIServer() {
	suite.IsType(&http.Server{}, suite.TestAPI.Server, "expects an instance of *http.Server")
}

func (suite *HyperdriveTestSuite) TestGetMediaType() {
	suite.Equal("application/vnd.api.test.v1.0.1", suite.TestAPI.GetMediaType(suite.TestEndpoint), "returns a media type string")
}

func TestHyperdriveTestSuite(t *testing.T) {
	suite.Run(t, new(HyperdriveTestSuite))
}
