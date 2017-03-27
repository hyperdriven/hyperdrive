package hyperdrive

import (
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TaggedStruct struct {
	Endpoint
	TestParam         string `param:"test_param;a=GET,PUT;r=PUT"`
	TestParamDefault  string `param:"test_param_default"`
	TestParamEmpty    string `param:""`
	TestParamRequired string `param:"test_param_required;a=GET;r=PUT"`
}

type TaggedEndpoint struct {
	Endpoint
	ID string `param:"id;r=GET"`
}

type CustomEndpoint struct {
	Endpoint
	ID ID `param:"id;r=GET"`
}

type ID int

func (id *ID) GetName() string {
	return "ID"
}

func (id *ID) GetDesc() string {
	return "The unique identifer for this resource."
}

type HyperdriveTestSuite struct {
	suite.Suite
	TestAPI                    API
	TestEndpoint               Endpointer
	TestHandler                http.Handler
	TestRoot                   *RootResource
	TestEndpointResource       EndpointResource
	TestEndpointResourceCustom EndpointResource
	TestEndpointResourceParam  EndpointResourceParam
	TestGetRequest             *http.Request
	TestGetRequestNoParams     *http.Request
	TestPostRequest            *http.Request
	TestParsedParam            parsedParam
	TestParsedParamDefault     parsedParam
	TestParsedParamEmpty       parsedParam
	TestParsedParamRequired    parsedParam
	TestParsedParamCustom      parsedParam
	TestTaggedStruct           *TaggedStruct
	TestParsedParamMap         parsedParams
	TestTaggedEndpoint         *TaggedEndpoint
	TestCustomEndpoint         *CustomEndpoint
}

func (suite *HyperdriveTestSuite) SetupTest() {
	suite.TestAPI = NewAPI("API", "Test API Desc")
	suite.TestEndpoint = NewEndpoint("Test", "Test Endpoint", "/test", "1.0.1-beta")
	suite.TestHandler = NewMethodHandler(suite.TestEndpoint)
	suite.TestRoot = NewRootResource(suite.TestAPI)
	suite.TestEndpointResource = NewEndpointResource(suite.TestEndpoint)
	suite.TestEndpointResourceParam = EndpointResourceParam{
		XMLName:      xml.Name{Space: "", Local: ""},
		Name:         "ID",
		Desc:         "The unique identifer for this resource.",
		Allowed:      []string{"GET", "PATCH", "POST", "PUT"},
		AllowedList:  "GET,PATCH,POST,PUT",
		Required:     []string{"GET"},
		RequiredList: "GET",
	}
	suite.TestEndpointResourceCustom = EndpointResource{
		XMLName:        xml.Name{Space: "", Local: ""},
		Resource:       "endpoint",
		Name:           "Test",
		Path:           "/test",
		MethodsList:    "OPTIONS",
		Methods:        []string{"OPTIONS"},
		MediaTypesList: "application/vnd.api.test.v1.0.1-beta.json,application/vnd.api.test.v1.0.1-beta.xml",
		MediaTypes:     []string{"application/vnd.api.test.v1.0.1-beta.json", "application/vnd.api.test.v1.0.1-beta.xml"},
		Desc:           "Test Endpoint",
		Params:         []EndpointResourceParam{suite.TestEndpointResourceParam},
	}
	suite.TestGetRequest = httptest.NewRequest("GET", "/test/2?id=1&a=b", nil)
	suite.TestGetRequestNoParams = httptest.NewRequest("GET", "/test", nil)
	suite.TestPostRequest = httptest.NewRequest("POST", "/test/2?id=1&a=b", strings.NewReader(`{"id":3}`))
	suite.TestParsedParam = parsedParam{"TestParam", "...", "TestParam", "string", "test_param", []string{"GET", "PUT"}, []string{"PUT"}}
	suite.TestParsedParamDefault = parsedParam{"TestParamDefault", "...", "TestParamDefault", "string", "test_param_default", []string{"GET", "PATCH", "POST", "PUT"}, []string{}}
	suite.TestParsedParamEmpty = parsedParam{"TestParamEmpty", "...", "TestParamEmpty", "string", "TestParamEmpty", []string{"GET", "PATCH", "POST", "PUT"}, []string{}}
	suite.TestParsedParamRequired = parsedParam{"TestParamRequired", "...", "TestParamRequired", "string", "test_param_required", []string{"GET", "PUT"}, []string{"PUT"}}
	suite.TestParsedParamCustom = parsedParam{"ID", "The unique identifer for this resource.", "ID", "ID", "id", []string{"GET", "PATCH", "POST", "PUT"}, []string{"GET"}}
	suite.TestParsedParamMap = parsedParams{"test_param": suite.TestParsedParam, "test_param_default": suite.TestParsedParamDefault, "TestParamEmpty": suite.TestParsedParamEmpty, "test_param_required": suite.TestParsedParamRequired}
	suite.TestTaggedStruct = &TaggedStruct{Endpoint: *NewEndpoint("Test", "Test Endpoint", "/test", "1.0.1-beta")}
	suite.TestTaggedEndpoint = &TaggedEndpoint{Endpoint: *NewEndpoint("Test", "Test Endpoint", "/test", "1.0.1-beta")}
	suite.TestCustomEndpoint = &CustomEndpoint{Endpoint: *NewEndpoint("Test", "Test Endpoint", "/test", "1.0.1-beta")}
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
