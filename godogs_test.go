package app_test

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/cucumber/godog"
	app "github.com/elct9620/gopherday2024"
)

var opts = godog.Options{
	Tags:   "~@wip",
	Format: "pretty",
	Paths:  []string{"features"},
}

const SuiteSuccessCode = 0

type HttpFeature struct {
	server   http.Handler
	response *httptest.ResponseRecorder
}

func (feat *HttpFeature) iMakeAGETRequestTo(url string) error {
	req := httptest.NewRequest("GET", url, nil)
	feat.response = httptest.NewRecorder()

	feat.server.ServeHTTP(feat.response, req)
	return nil
}

func (feat *HttpFeature) theResponseStatusCodeShouldBe(statusCode int) error {
	if feat.response.Code != statusCode {
		return fmt.Errorf("expected response code to be %d, but actual is %d", statusCode, feat.response.Code)
	}
	return nil
}

func (feat *HttpFeature) theResponseBodyShouldBe(doc *godog.DocString) error {
	var expected any
	var actual any

	if json.Unmarshal([]byte(doc.Content), &expected) != nil {
		return fmt.Errorf("expected response body to be a valid JSON, but actual is %s", doc.Content)
	}

	actualBody := feat.response.Body.String()
	if json.Unmarshal([]byte(actualBody), &actual) != nil {
		return fmt.Errorf("actual response body is not a valid JSON")
	}

	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected response body to be %s, but actual is %s", doc.Content, actualBody)
	}

	return nil
}

func InitializeScenario(s *godog.ScenarioContext) {
	rest, err := app.Initialize()
	if err != nil {
		panic(err)
	}

	httpFeat := &HttpFeature{
		server: rest,
	}

	s.Step(`^I make a GET request to "([^"]*)"$`, httpFeat.iMakeAGETRequestTo)
	s.Step(`^the response status code should be (\d+)$`, httpFeat.theResponseStatusCodeShouldBe)
	s.Step(`^the response body should be$`, httpFeat.theResponseBodyShouldBe)
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestFeatures(t *testing.T) {
	o := opts
	o.TestingT = t

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options:             &o,
	}

	if suite.Run() != SuiteSuccessCode {
		t.Fatal("Non-zero exit code")
	}
}
