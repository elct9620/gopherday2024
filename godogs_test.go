package app_test

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	app "github.com/elct9620/gopherday2024"
	"github.com/jmespath/go-jmespath"
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

func (feat *HttpFeature) iMakeAPOSTRequestToWithTheBody(url string, doc *godog.DocString) error {
	req := httptest.NewRequest("POST", url, strings.NewReader(doc.Content))
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
		return fmt.Errorf("actual response body is not a valid JSON: %s", actualBody)
	}

	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected response body to be %s, but actual is %s", doc.Content, actualBody)
	}

	return nil
}

func (feat *HttpFeature) theResponseJsonShouldHave(path string) error {
	var actual any

	actualBody := feat.response.Body.String()
	if json.Unmarshal([]byte(actualBody), &actual) != nil {
		return fmt.Errorf("actual response body is not a valid JSON: %s", feat.response.Body.String())
	}

	res, err := jmespath.Search(path, actual)
	if err != nil {
		return fmt.Errorf("failed to search path %s: %s", path, err)
	}

	if res == nil {
		return fmt.Errorf(`expected response body to have JMESPath "%s", but actual is %s`, path, actualBody)
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	rest, err := app.InitializeTest()
	if err != nil {
		panic(err)
	}

	httpFeat := &HttpFeature{
		server: rest,
	}

	ctx.Step(`^I make a GET request to "([^"]*)"$`, httpFeat.iMakeAGETRequestTo)
	ctx.Step(`^I make a POST request to "([^"]*)" with the body$`, httpFeat.iMakeAPOSTRequestToWithTheBody)
	ctx.Step(`^the response status code should be (\d+)$`, httpFeat.theResponseStatusCodeShouldBe)
	ctx.Step(`^the response json should have "([^"]*)"$`, httpFeat.theResponseJsonShouldHave)
	ctx.Step(`^the response body should be$`, httpFeat.theResponseBodyShouldBe)
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
