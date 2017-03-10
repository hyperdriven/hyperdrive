package hyperdrive

import "os"

func (suite *HyperdriveTestSuite) TestNewConfig() {
	suite.IsType(Config{}, NewConfig(), "expects an instance of *hyperdrive.Config")
}

func (suite *HyperdriveTestSuite) TestPortConfigFromDefault() {
	c := NewConfig()
	suite.Equal(5000, c.Port, "Port should be equal to default value")
}

func (suite *HyperdriveTestSuite) TestPortConfigFromEnv() {
	os.Setenv("PORT", "5001")
	c := NewConfig()
	suite.Equal(5001, c.Port, "Port should be equal to PORT value set via ENV var")
}

func (suite *HyperdriveTestSuite) TestGetPort() {
	c := NewConfig()
	suite.Equal(":5000", c.GetPort(), "c.Port value should be return, prefixed with a colon, e.g. :5000")
}

func (suite *HyperdriveTestSuite) TestEnvConfigFromDefault() {
	c := NewConfig()
	suite.Equal("development", c.Env, "Env should be equal to default value")
}

func (suite *HyperdriveTestSuite) TestEnvConfigFromEnv() {
	os.Setenv("HYPERDRIVE_ENV", "test")
	c := NewConfig()
	suite.Equal("test", c.Env, "Env should be equal to HYPERDRIVE_ENV value set via ENV var")
}

func (suite *HyperdriveTestSuite) TestGzipLevelConfigFromDefault() {
	c := NewConfig()
	suite.Equal(-1, c.GzipLevel, "GzipLevel should be equal to default value")
}

func (suite *HyperdriveTestSuite) TestGzipLevelConfigFromEnv() {
	os.Setenv("GZIP_LEVEL", "9")
	c := NewConfig()
	suite.Equal(9, c.GzipLevel, "GzipLevel should be equal to GZIP_LEVEL value set via ENV var")
}
