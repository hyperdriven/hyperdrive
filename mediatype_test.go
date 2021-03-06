package hyperdrive

func (suite *HyperdriveTestSuite) TestGetMediaType() {
	suite.Equal("application/vnd.api.test.v1.0.1-beta", GetMediaType(suite.TestAPI, suite.TestEndpoint), "returns a media type string")
}

func (suite *HyperdriveTestSuite) TestGetContentTypeJSON() {
	suite.Equal("application/vnd.api.test.v1.0.1-beta.json", GetContentTypeJSON(suite.TestAPI, suite.TestEndpoint), "returns a media type string")
}

func (suite *HyperdriveTestSuite) TestGetContentTypeXML() {
	suite.Equal("application/vnd.api.test.v1.0.1-beta.xml", GetContentTypeXML(suite.TestAPI, suite.TestEndpoint), "returns a media type string")
}

func (suite *HyperdriveTestSuite) TestGetContentTypes() {
	suite.Equal([]string{"application/vnd.api.test.v1.0.1-beta.json", "application/vnd.api.test.v1.0.1-beta.xml"}, GetContentTypes(suite.TestAPI, suite.TestEndpoint), "returns a media type string")
}

func (suite *HyperdriveTestSuite) TestGetContentTypesList() {
	suite.Equal("application/vnd.api.test.v1.0.1-beta.json,application/vnd.api.test.v1.0.1-beta.xml", GetContentTypesList(suite.TestAPI, suite.TestEndpoint), "returns a media type string")
}
