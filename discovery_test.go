package hyperdrive

import "net/http"

func (suite *HyperdriveTestSuite) TestNewRootResource() {
	suite.IsType(&RootResource{}, suite.TestRoot, "expects an instance of RootResource")
}

func (suite *HyperdriveTestSuite) TestRootResourceEndpointsEmpty() {
	suite.Equal(0, len(suite.TestRoot.Endpoints), "expects 0 Endpoints")
}

func (suite *HyperdriveTestSuite) TestAddEndpointer() {
	suite.TestRoot.AddEndpointer(suite.TestEndpoint)
	suite.Equal(1, len(suite.TestRoot.Endpoints), "expects 1 Endpoints")
}

func (suite *HyperdriveTestSuite) TestRootResourceServeHTTP() {
	suite.Implements((*http.Handler)(nil), suite.TestRoot, "return an implementation of http.Handler")
}

func (suite *HyperdriveTestSuite) TestPresentRepresentation() {
	suite.IsType(Representation{}, suite.TestRoot.Present(), "return a Representation")
}

func (suite *HyperdriveTestSuite) TestPresent() {
	suite.TestRoot.AddEndpointer(suite.TestEndpoint)
	suite.Equal(suite.TestRootRepresentation, suite.TestRoot.Present(), "return the correct Representation of RootResource")
}

func (suite *HyperdriveTestSuite) TestPresentEndpoint() {
	suite.Equal(suite.TestEndpointRepresentation, PresentEndpoint(suite.TestEndpoint), "return the correct Representation of RootResource")
}
