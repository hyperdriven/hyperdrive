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

func (suite *HyperdriveTestSuite) TestAllowedList() {
	suite.Equal("GET,PUT", suite.TestParsedParamMap["test_param"].AllowedList(), "expects it to return the correct methods")
}

func (suite *HyperdriveTestSuite) TestRequiredList() {
	suite.Equal("PUT", suite.TestParsedParamMap["test_param"].RequiredList(), "expects it to return the correct methods")
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

func (suite *HyperdriveTestSuite) TestParsedParamsAllowed() {
	suite.IsType([]string{}, suite.TestParsedParamMap.Allowed("GET"), "expects it to return the correct slice of strings")
}

func (suite *HyperdriveTestSuite) TestParsedParamsNotAllowed() {
	suite.IsType([]string{}, suite.TestParsedParamMap.Allowed("POST"), "expects it to return the correct slice of strings")
}

func (suite *HyperdriveTestSuite) TestParsedParamsRequired() {
	suite.IsType([]string{}, suite.TestParsedParamMap.Required("PUT"), "expects it to return the correct slice of strings")
}

func (suite *HyperdriveTestSuite) TestParsedParamsNotRequired() {
	suite.IsType([]string{}, suite.TestParsedParamMap.Required("GET"), "expects it to return the correct slice of strings")
}

func (suite *HyperdriveTestSuite) TestParsedParamsRequiredEmpty() {
	suite.IsType([]string{}, suite.TestParsedParamMap.Required("POST"), "expects it to return a map of parsedParams")
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

func (suite *HyperdriveTestSuite) TestParameterName() {
	suite.Equal("ID", parseEndpoint(suite.TestCustomEndpoint)["id"].Name, "expects it to return the correct Name")
}

func (suite *HyperdriveTestSuite) TestParameterDesc() {
	suite.Equal("The unique identifer for this resource.", parseEndpoint(suite.TestCustomEndpoint)["id"].Desc, "expects it to return the correct Desc")
}

func (suite *HyperdriveTestSuite) TestParseTestParamCustom() {
	suite.Equal(suite.TestParsedParamCustom, parseEndpoint(suite.TestCustomEndpoint)["id"], "expects it to return the correct parsedParam")
}
