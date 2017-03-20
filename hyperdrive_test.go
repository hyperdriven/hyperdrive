package hyperdrive

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TaggedStruct struct {
	*Endpoint
	TestParam         string `param:"test_param;a=GET,PUT;r=PUT"`
	TestParamDefault  string `param:"test_param_default"`
	TestParamEmpty    string `param:""`
	TestParamRequired string `param:"test_param_required;a=GET;r=PUT"`
}

type HyperdriveTestSuite struct {
	suite.Suite
	TestAPI                 API
	TestEndpoint            Endpointer
	TestHandler             http.Handler
	TestRoot                *RootResource
	TestEndpointResource    EndpointResource
	TestGetRequest          *http.Request
	TestPostRequest         *http.Request
	TestParsedParam         parsedParam
	TestParsedParamDefault  parsedParam
	TestParsedParamEmpty    parsedParam
	TestParsedParamRequired parsedParam
	TestTaggedStruct        TaggedStruct
	TestParsedParamMap      map[string]parsedParam
}

func (suite *HyperdriveTestSuite) SetupTest() {
	suite.TestAPI = NewAPI("API", "Test API Desc")
	suite.TestEndpoint = NewEndpoint("Test", "Test Endpoint", "/test", "1.0.1-beta")
	suite.TestHandler = NewMethodHandler(suite.TestEndpoint)
	suite.TestRoot = NewRootResource(suite.TestAPI)
	suite.TestEndpointResource = NewEndpointResource(suite.TestEndpoint)
	suite.TestGetRequest = httptest.NewRequest("GET", "/test/2?id=1&a=b", nil)
	suite.TestPostRequest = httptest.NewRequest("POST", "/test/2?id=1&a=b", strings.NewReader(`{"id":3}`))
	suite.TestParsedParam = parsedParam{"TestParam", "string", "test_param", []string{"GET", "PUT"}, []string{"PUT"}}
	suite.TestParsedParamDefault = parsedParam{"TestParamDefault", "string", "test_param_default", []string{"GET", "PATCH", "POST", "PUT"}, []string{}}
	suite.TestParsedParamEmpty = parsedParam{"TestParamEmpty", "string", "TestParamEmpty", []string{"GET", "PATCH", "POST", "PUT"}, []string{}}
	suite.TestParsedParamRequired = parsedParam{"TestParamRequired", "string", "test_param_required", []string{"GET", "PUT"}, []string{"PUT"}}
	suite.TestParsedParamMap = map[string]parsedParam{"test_param": suite.TestParsedParam, "test_param_default": suite.TestParsedParamDefault, "TestParamEmpty": suite.TestParsedParamEmpty, "test_param_required": suite.TestParsedParamRequired}
	suite.TestTaggedStruct = TaggedStruct{}
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
