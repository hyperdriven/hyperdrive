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
