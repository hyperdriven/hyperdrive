package hyperdrive

func (suite *HyperdriveTestSuite) TestNewEndpoint() {
	suite.IsType(&Endpoint{}, suite.TestEndpoint, "expects an instance of hyperdrive.Endpoint")
}

func (suite *HyperdriveTestSuite) TestGetName() {
	suite.Equal("Test", suite.TestEndpoint.GetName(), "expects GetName() to return Name")
}

func (suite *HyperdriveTestSuite) TestGetDesc() {
	suite.Equal("Test Endpoint", suite.TestEndpoint.GetDesc(), "expects GetDesc() to return Desc")
}

func (suite *HyperdriveTestSuite) TestGetPath() {
	suite.Equal("/test", suite.TestEndpoint.GetPath(), "expects GetPath() to return Path")
}

func (suite *HyperdriveTestSuite) TestEndpointer() {
	suite.Implements((*Endpointer)(nil), suite.TestEndpoint, "expects an implementation of hyperdrive.Endpointer interface")
}

func (suite *HyperdriveTestSuite) TestGetMethods() {
	suite.Equal([]string{"OPTIONS"}, GetMethods(suite.TestEndpoint), "expects a slice of supported method strings")
}

func (suite *HyperdriveTestSuite) TestGetMethodsList() {
	suite.Equal("OPTIONS", GetMethodsList(suite.TestEndpoint), "expects a list of supported method strings")
}