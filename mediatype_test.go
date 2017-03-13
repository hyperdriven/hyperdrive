package hyperdrive

func (suite *HyperdriveTestSuite) TestGetMediaType() {
	suite.Equal("application/vnd.api.test.v1.0.1", GetMediaType(suite.TestAPI, suite.TestEndpoint), "returns a media type string")
}

func (suite *HyperdriveTestSuite) TestGetContentTypeJSON() {
	suite.Equal("application/vnd.api.test.v1.0.1.json", GetContentTypeJSON(suite.TestAPI, suite.TestEndpoint), "returns a media type string")
}

func (suite *HyperdriveTestSuite) TestGetContentTypeXML() {
	suite.Equal("application/vnd.api.test.v1.0.1.xml", GetContentTypeXML(suite.TestAPI, suite.TestEndpoint), "returns a media type string")
}

func (suite *HyperdriveTestSuite) TestGetContentTypes() {
	suite.Equal([]string{"application/vnd.api.test.v1.0.1.json", "application/vnd.api.test.v1.0.1.xml"}, GetContentTypes(suite.TestAPI, suite.TestEndpoint), "returns a media type string")
}

func (suite *HyperdriveTestSuite) TestGetContentTypesList() {
	suite.Equal("application/vnd.api.test.v1.0.1.json,application/vnd.api.test.v1.0.1.xml", GetContentTypesList(suite.TestAPI, suite.TestEndpoint), "returns a media type string")
}
