package hyperdrive

import (
	"net/http"
	"net/http/httptest"
)

func (suite *HyperdriveTestSuite) TestNewRootResource() {
	suite.IsType(&RootResource{}, suite.TestRoot, "expects an instance of RootResource")
}

func (suite *HyperdriveTestSuite) TestRootResourceEndpointsEmpty() {
	suite.Equal(0, len(suite.TestRoot.Endpoints), "expects 0 Endpoints")
}

func (suite *HyperdriveTestSuite) TestAddEndpointer() {
	suite.TestRoot.AddEndpoint(suite.TestEndpoint)
	suite.Equal(1, len(suite.TestRoot.Endpoints), "expects 1 Endpoints")
}

func (suite *HyperdriveTestSuite) TestNewEndpointResource() {
	suite.Equal(suite.TestEndpointResourceCustom, NewEndpointResource(suite.TestCustomEndpoint), "expects the correct EndpointResource")
}

func (suite *HyperdriveTestSuite) TestNewEndpointResourceParam() {
	suite.Equal(suite.TestEndpointResourceParam, NewEndpointResourceParam(suite.TestParsedParamCustom), "expects the correct EndpointResource")
}

func (suite *HyperdriveTestSuite) TestRootResourceServeHTTP() {
	suite.Implements((*http.Handler)(nil), suite.TestRoot, "return an implementation of http.Handler")
}

func (suite *HyperdriveTestSuite) TestRootResourceRequest() {
	rw := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Add("Accept", "application/json")
	suite.TestRoot.ServeHTTP(rw, r)
	suite.Equal(200, rw.Result().StatusCode, "expects request to discovery url to be successful")
}
