package hyperdrive

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

func (suite *HyperdriveTestSuite) TestQueryParamsGet() {
	suite.IsType(url.Values{}, QueryParams(suite.TestGetRequest), "expects an instance of url.Values")
}

func (suite *HyperdriveTestSuite) TestQueryParamsGetValues() {
	suite.Equal(url.Values{"id": []string{"1"}, "a": []string{"b"}}, QueryParams(suite.TestGetRequest), "returns populated url.Values")
}

func (suite *HyperdriveTestSuite) TestQueryParamsNonGetValues() {
	suite.Equal(url.Values{}, QueryParams(suite.TestPostRequest), "returns populated url.Values")
}

func (suite *HyperdriveTestSuite) TestBodyParamsGet() {
	suite.IsType(url.Values{}, BodyParams(suite.TestGetRequest), "expects an instance of url.Values")
}

func (suite *HyperdriveTestSuite) TestBodyParamsGetValues() {
	suite.Equal(url.Values{}, BodyParams(suite.TestGetRequest), "returns populated url.Values")
}

func (suite *HyperdriveTestSuite) TestBodyParamsNonGetValues() {
	suite.Equal(url.Values{}, BodyParams(suite.TestPostRequest), "returns populated url.Values")
}

func (suite *HyperdriveTestSuite) TestPathParamsGet() {
	suite.TestAPI.Router.Handle("/test/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		suite.IsType(url.Values{}, PathParams(suite.TestGetRequest), "expects an instance of url.Values")
	}))
	suite.TestAPI.Router.ServeHTTP(httptest.NewRecorder(), suite.TestGetRequest)
}

func (suite *HyperdriveTestSuite) TestPathParamsGetValues() {
	suite.TestAPI.Router.Handle("/test/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(url.Values{"id": []string{"2"}}, PathParams(r), "returns populated url.Values")
	}))
	suite.TestAPI.Router.ServeHTTP(httptest.NewRecorder(), suite.TestGetRequest)
}

func (suite *HyperdriveTestSuite) TestParamsGet() {
	suite.TestAPI.Router.Handle("/test/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		suite.IsType(url.Values{}, Params(suite.TestGetRequest), "expects an instance of url.Values")
	}))
	suite.TestAPI.Router.ServeHTTP(httptest.NewRecorder(), suite.TestGetRequest)

}

func (suite *HyperdriveTestSuite) TestParamsGetValues() {
	suite.TestAPI.Router.Handle("/test/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(url.Values{"id": []string{"1"}, "a": []string{"b"}}, Params(suite.TestGetRequest), "returns populated url.Values")
	}))
	suite.TestAPI.Router.ServeHTTP(httptest.NewRecorder(), suite.TestGetRequest)
}

func (suite *HyperdriveTestSuite) TestParamsNonGetValues() {
	suite.TestAPI.Router.Handle("/test/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		suite.Equal(url.Values{}, Params(suite.TestPostRequest), "returns populated url.Values")
		//suite.Equal(url.Values{"id": []string{"3"}}, Params(suite.TestPostRequest), "returns populated url.Values")
	}))
	suite.TestAPI.Router.ServeHTTP(httptest.NewRecorder(), suite.TestPostRequest)
}

func (suite *HyperdriveTestSuite) TestGetParams() {
	params, err := GetParams(suite.TestTaggedEndpoint, suite.TestGetRequest)
	suite.Equal(url.Values{"id": []string{"1"}}, params, "returns populated url.Values")
	suite.Nil(err, "returns populated url.Values")
}

func (suite *HyperdriveTestSuite) TestGetParamsError() {
	params, err := GetParams(suite.TestTaggedEndpoint, suite.TestGetRequestNoParams)
	suite.Equal(url.Values{}, params, "returns populated url.Values")
	suite.Error(err, "returns populated url.Values")
}

func (suite *HyperdriveTestSuite) TestParameter() {
	suite.Implements((*Parameter)(nil), new(ID), "is an implementation of Parameter")
}
