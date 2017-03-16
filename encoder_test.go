package hyperdrive

import (
	"encoding/json"
	"encoding/xml"
	"net/http/httptest"
)

func (suite *HyperdriveTestSuite) TestNullEncoder() {
	suite.Implements((*ContentEncoder)(nil), NullEncoder{}, "return an implementation of ContentEncoder")
}

func (suite *HyperdriveTestSuite) TestNullEncoderEncode() {
	suite.Error(NullEncoder{}.Encode(suite.TestEndpointResource), "return an error")
}

func (suite *HyperdriveTestSuite) TestJSONEncoder() {
	suite.Implements((*ContentEncoder)(nil), JSONEncoder{}, "return an implementation of ContentEncoder")
}

func (suite *HyperdriveTestSuite) TestJSONEncoderEncodeNoError() {
	rw := httptest.NewRecorder()
	enc := JSONEncoder{Encoder: json.NewEncoder(rw)}
	suite.Nil(enc.Encode(suite.TestEndpointResource), "returns nil")
}

func (suite *HyperdriveTestSuite) TestJSONEncoderEncode() {
	rw := httptest.NewRecorder()
	enc := JSONEncoder{Encoder: json.NewEncoder(rw)}
	enc.Encode(suite.TestEndpointResource)
	json := `{"resource":"endpoint","name":"Test","path":"/test","methods":["OPTIONS"],"media-types":["application/vnd.api.test.v1.0.1-beta.json","application/vnd.api.test.v1.0.1-beta.xml"],"description":"Test Endpoint"}` + "\n"
	suite.Equal(json, rw.Body.String(), "returns nil")
}

func (suite *HyperdriveTestSuite) TestXMLEncoder() {
	suite.Implements((*ContentEncoder)(nil), XMLEncoder{}, "return an implementation of ContentEncoder")
}

func (suite *HyperdriveTestSuite) TestXMLEncoderEncodeNoError() {
	rw := httptest.NewRecorder()
	enc := XMLEncoder{Encoder: xml.NewEncoder(rw)}
	suite.Nil(enc.Encode(suite.TestEndpointResource), "returns nil")
}

func (suite *HyperdriveTestSuite) TestXMLEncoderEncode() {
	rw := httptest.NewRecorder()
	enc := XMLEncoder{Encoder: xml.NewEncoder(rw)}
	enc.Encode(suite.TestEndpointResource)
	xml := `<endpoint name="Test" path="/test" methods="OPTIONS" media-types="application/vnd.api.test.v1.0.1-beta.json,application/vnd.api.test.v1.0.1-beta.xml"><description>Test Endpoint</description></endpoint>`
	suite.Equal(xml, rw.Body.String(), "returns nil")
}

func (suite *HyperdriveTestSuite) TestGetEncoder() {
	enc, _ := GetEncoder(httptest.NewRecorder(), "text/plain")
	suite.Implements((*ContentEncoder)(nil), enc, "return an implementation of ContentEncoder")
}

func (suite *HyperdriveTestSuite) TestGetEncoderXML() {
	enc, _ := GetEncoder(httptest.NewRecorder(), "application/xml")
	suite.IsType(XMLEncoder{}, enc, "return an XMLEncoder")
}

func (suite *HyperdriveTestSuite) TestGetEncoderJSON() {
	enc, _ := GetEncoder(httptest.NewRecorder(), "application/json")
	suite.IsType(JSONEncoder{}, enc, "return a JSONEncoder")
}

func (suite *HyperdriveTestSuite) TestGetEncoderNULL() {
	enc, _ := GetEncoder(httptest.NewRecorder(), "text/plain")
	suite.IsType(NullEncoder{}, enc, "return a NullEncoder")
}
