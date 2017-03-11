package hyperdrive

import (
	"errors"
	"net/http"
)

func (suite *HyperdriveTestSuite) TestGetErrorTextProduction() {
	conf.Env = "production"
	suite.Equal(http.StatusText(406), GetErrorText(406, errors.New("Test Error")), "returns 406 Status Text")
}

func (suite *HyperdriveTestSuite) TestGetErrorText() {
	suite.Equal("Test Error", GetErrorText(406, errors.New("Test Error")), "returns Error Text")
}
