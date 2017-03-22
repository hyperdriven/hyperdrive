package hyperdrive

func (suite *HyperdriveTestSuite) TestTagName() {
	suite.Equal("param", tagName, "expects tagName to be correct")
}

func (suite *HyperdriveTestSuite) TestIsAllowedTrue() {
	suite.Equal(true, suite.TestParsedParam.IsAllowed("GET"), "expects it to return true")
}

func (suite *HyperdriveTestSuite) TestIsAllowedFalse() {
	suite.Equal(false, suite.TestParsedParam.IsAllowed("POST"), "expects it to return false")
}

func (suite *HyperdriveTestSuite) TestIsRequiredTrue() {
	suite.Equal(true, suite.TestParsedParam.IsRequired("PUT"), "expects it to return true")
}

func (suite *HyperdriveTestSuite) TestIsRequiredFalse() {
	suite.Equal(false, suite.TestParsedParam.IsRequired("POST"), "expects it to return false")
}

func (suite *HyperdriveTestSuite) TestContainsTrue() {
	suite.Equal(true, contains([]string{"GET"}, "GET"), "expects it to return true")
}

func (suite *HyperdriveTestSuite) TestContainsFalse() {
	suite.Equal(false, contains([]string{"GET"}, "POST"), "expects it to return false")
}

func (suite *HyperdriveTestSuite) TestParse() {
	suite.IsType(parsedParams{}, parseEndpoint(suite.TestTaggedStruct), "expects it to return a map of parsedParams")
}

func (suite *HyperdriveTestSuite) TestParseTestParam() {
	suite.Equal(suite.TestParsedParamMap["test_param"], parseEndpoint(suite.TestTaggedStruct)["test_param"], "expects it to return the correct parsedParam")
}

func (suite *HyperdriveTestSuite) TestParseTestParamDefault() {
	suite.Equal(suite.TestParsedParamMap["test_param_default"], parseEndpoint(suite.TestTaggedStruct)["test_param_default"], "expects it to return the correct parsedParam")
}

func (suite *HyperdriveTestSuite) TestParseTestParamEmpty() {
	suite.Equal(suite.TestParsedParamMap["TestParamEmpty"], parseEndpoint(suite.TestTaggedStruct)["TestParamEmpty"], "expects it to return the correct parsedParam")
}

func (suite *HyperdriveTestSuite) TestParseTestParamRequired() {
	suite.Equal(suite.TestParsedParamMap["test_param_required"], parseEndpoint(suite.TestTaggedStruct)["test_param_required"], "expects it to return the correct parsedParam")
}
