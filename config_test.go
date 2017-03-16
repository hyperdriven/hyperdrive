package hyperdrive

import "os"

func (suite *HyperdriveTestSuite) TestNewConfig() {
	c, _ := NewConfig()
	suite.IsType(Config{}, c, "expects an instance of *hyperdrive.Config")
}

func (suite *HyperdriveTestSuite) TestNewConfigError() {
	os.Setenv("PORT", "abc")
	defer os.Unsetenv("PORT")
	_, err := NewConfig()
	suite.Error(err, "will throw an error if type mismatch")
}

func (suite *HyperdriveTestSuite) TestPortConfigFromDefault() {
	c, _ := NewConfig()
	suite.Equal(5000, c.Port, "Port should be equal to default value")
}

func (suite *HyperdriveTestSuite) TestPortConfigFromEnv() {
	os.Setenv("PORT", "5001")
	defer os.Unsetenv("PORT")
	c, _ := NewConfig()
	suite.Equal(5001, c.Port, "Port should be equal to PORT value set via ENV var")
}

func (suite *HyperdriveTestSuite) TestGetPort() {
	c, _ := NewConfig()
	suite.Equal(":5000", c.GetPort(), "c.Port value should be return, prefixed with a colon, e.g. :5000")
}

func (suite *HyperdriveTestSuite) TestEnvConfigFromDefault() {
	c, _ := NewConfig()
	suite.Equal("development", c.Env, "Env should be equal to default value")
}

func (suite *HyperdriveTestSuite) TestEnvConfigFromEnv() {
	os.Setenv("HYPERDRIVE_ENV", "test")
	defer os.Unsetenv("HYPERDRIVE_ENV")
	c, _ := NewConfig()
	suite.Equal("test", c.Env, "Env should be equal to HYPERDRIVE_ENV value set via ENV var")
}

func (suite *HyperdriveTestSuite) TestGzipLevelConfigFromDefault() {
	c, _ := NewConfig()
	suite.Equal(-1, c.GzipLevel, "GzipLevel should be equal to default value")
}

func (suite *HyperdriveTestSuite) TestGzipLevelConfigFromEnv() {
	os.Setenv("GZIP_LEVEL", "9")
	defer os.Unsetenv("GZIP_LEVEL")
	c, _ := NewConfig()
	suite.Equal(9, c.GzipLevel, "GzipLevel should be equal to GZIP_LEVEL value set via ENV var")
}

func (suite *HyperdriveTestSuite) TestCorsEnabledConfigFromDefault() {
	c, _ := NewConfig()
	suite.Equal(true, c.CorsEnabled, "CorsEnabled should be equal to default value")
}

func (suite *HyperdriveTestSuite) TestCorsEnabledConfigFromEnv() {
	os.Setenv("CORS_ENABLED", "false")
	defer os.Unsetenv("CORS_ENABLED")
	c, _ := NewConfig()
	suite.Equal(false, c.CorsEnabled, "CorsEnabled should be equal to CORS_ENABLED value set via ENV var")
}

func (suite *HyperdriveTestSuite) TestCorsOriginsConfigFromDefault() {
	c, _ := NewConfig()
	suite.Equal("*", c.CorsOrigins, "CorsOrigins should be equal to default value")
}

func (suite *HyperdriveTestSuite) TestCorsOriginsConfigFromEnv() {
	os.Setenv("CORS_ORIGINS", "example.com")
	defer os.Unsetenv("CORS_ORIGINS")
	c, _ := NewConfig()
	suite.Equal("example.com", c.CorsOrigins, "CorsOrigins should be equal to CORS_ORIGINS value set via ENV var")
}

func (suite *HyperdriveTestSuite) TestCorsHeadersConfigFromDefault() {
	c, _ := NewConfig()
	suite.Equal("", c.CorsHeaders, "CorsHeaders should be equal to default value")
}

func (suite *HyperdriveTestSuite) TestCorsHeadersConfigFromEnv() {
	os.Setenv("CORS_HEADERS", "example.com")
	defer os.Unsetenv("CORS_HEADERS")
	c, _ := NewConfig()
	suite.Equal("example.com", c.CorsHeaders, "CorsHeaders should be equal to CORS_HEADERS value set via ENV var")
}

func (suite *HyperdriveTestSuite) TestCorsCredentialsConfigFromDefault() {
	c, _ := NewConfig()
	suite.Equal(true, c.CorsCredentials, "CorsCredentials should be equal to default value")
}

func (suite *HyperdriveTestSuite) TestCorsCredentialsConfigFromEnv() {
	os.Setenv("CORS_CREDENTIALS", "false")
	defer os.Unsetenv("CORS_CREDENTIALS")
	c, _ := NewConfig()
	suite.Equal(false, c.CorsCredentials, "CorsCredentials should be equal to CORS_CREDENTIALS value set via ENV var")
}
