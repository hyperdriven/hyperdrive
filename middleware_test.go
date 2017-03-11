package hyperdrive

import "net/http"

func (suite *HyperdriveTestSuite) TestDefaultMiddlewareChain() {
	suite.Implements((*http.Handler)(nil), suite.TestAPI.DefaultMiddlewareChain(suite.TestHandler), "return an implementation of http.Handler")
}

func (suite *HyperdriveTestSuite) TestLoggingMiddleware() {
	suite.Implements((*http.Handler)(nil), suite.TestAPI.LoggingMiddleware(suite.TestHandler), "return an implementation of http.Handler")
}

func (suite *HyperdriveTestSuite) TestRecoveryMiddleware() {
	suite.Implements((*http.Handler)(nil), suite.TestAPI.RecoveryMiddleware(suite.TestHandler), "return an implementation of http.Handler")
}

func (suite *HyperdriveTestSuite) TestCompressionMiddleware() {
	suite.Implements((*http.Handler)(nil), suite.TestAPI.CompressionMiddleware(suite.TestHandler), "return an implementation of http.Handler")
}

func (suite *HyperdriveTestSuite) TestMethodOverrideMiddleware() {
	suite.Implements((*http.Handler)(nil), suite.TestAPI.MethodOverrideMiddleware(suite.TestHandler), "return an implementation of http.Handler")
}

func (suite *HyperdriveTestSuite) TestCorsMiddleware() {
	suite.Implements((*http.Handler)(nil), suite.TestAPI.CorsMiddleware(suite.TestHandler), "return an implementation of http.Handler")
}

func (suite *HyperdriveTestSuite) TestContentTypeOptionsMiddleware() {
	suite.Implements((*http.Handler)(nil), suite.TestAPI.ContentTypeOptionsMiddleware(suite.TestHandler), "return an implementation of http.Handler")
}

func (suite *HyperdriveTestSuite) TestFrameOptionsMiddleware() {
	suite.Implements((*http.Handler)(nil), suite.TestAPI.FrameOptionsMiddleware(suite.TestHandler), "return an implementation of http.Handler")
}
