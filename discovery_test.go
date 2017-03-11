package hyperdrive

import "net/http"

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

func (suite *HyperdriveTestSuite) TestRootResourceServeHTTP() {
	suite.Implements((*http.Handler)(nil), suite.TestRoot, "return an implementation of http.Handler")
}
