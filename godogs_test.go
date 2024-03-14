package app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
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

type apiServerCtxKey struct{}

func getApiServer(ctx context.Context) (http.Handler, error) {
	v, ok := ctx.Value(apiServerCtxKey{}).(http.Handler)
	if !ok {
		return nil, fmt.Errorf("api server not found in context")
	}

	return v, nil
}

type apiResponseCtxKey struct{}

func getApiResponse(ctx context.Context) (*httptest.ResponseRecorder, error) {
	v, ok := ctx.Value(apiResponseCtxKey{}).(*httptest.ResponseRecorder)
	if !ok {
		return nil, fmt.Errorf("api response not found in context")
	}

	return v, nil
}

func getApiResponseJson(ctx context.Context) (any, error) {
	res, err := getApiResponse(ctx)
	if err != nil {
		return nil, err
	}

	var body io.Reader = bytes.NewBufferString(res.Body.String())
	var actual any
	if json.NewDecoder(body).Decode(&actual) != nil {
		return nil, fmt.Errorf("actual response body is not a valid JSON: %s", res.Body.String())
	}

	return actual, nil
}

func iMakeAGETRequestTo(ctx context.Context, url string) (context.Context, error) {
	server, err := getApiServer(ctx)
	if err != nil {
		return ctx, err
	}

	req := httptest.NewRequest("GET", url, nil)
	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)

	return context.WithValue(ctx, apiResponseCtxKey{}, res), nil
}

func iMakeARequestToWithTheBody(ctx context.Context, method, url string, doc *godog.DocString) (context.Context, error) {
	server, err := getApiServer(ctx)
	if err != nil {
		return ctx, err
	}

	req := httptest.NewRequest(method, url, strings.NewReader(doc.Content))
	req.Header.Add("Content-Type", "application/json")
	res := httptest.NewRecorder()
	server.ServeHTTP(res, req)

	return context.WithValue(ctx, apiResponseCtxKey{}, res), nil
}

func theResponseStatusCodeShouldBe(ctx context.Context, statusCode int) (context.Context, error) {
	res, err := getApiResponse(ctx)
	if err != nil {
		return ctx, err
	}

	if res.Code != statusCode {
		return ctx, fmt.Errorf("expected response code to be %d, but actual is %d", statusCode, res.Code)
	}

	return ctx, nil
}

func theResponseBodyShouldBe(ctx context.Context, doc *godog.DocString) error {
	res, err := getApiResponse(ctx)
	if err != nil {
		return err
	}

	var expected any
	if json.Unmarshal([]byte(doc.Content), &expected) != nil {
		return fmt.Errorf("expected response body to be a valid JSON, but actual is %s", doc.Content)
	}

	actual, err := getApiResponseJson(ctx)
	if err != nil {
		return err
	}

	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected response body to be %s, but actual is %s", doc.Content, res.Body.String())
	}

	return nil
}

func theResponseJsonShouldHave(ctx context.Context, path string) error {
	res, err := getApiResponse(ctx)
	if err != nil {
		return err
	}

	actual, err := getApiResponseJson(ctx)
	if err != nil {
		return err
	}

	found, err := jmespath.Search(path, actual)
	if err != nil {
		return fmt.Errorf("failed to search path %s: %s", path, err)
	}

	if found == nil {
		return fmt.Errorf(`expected response body to have JMESPath "%s", but actual is %s`, path, res.Body.String())
	}

	return nil
}

func theResponseJsonShouldHaveWithValue(ctx context.Context, path, value string) error {
	res, err := getApiResponse(ctx)
	if err != nil {
		return err
	}

	actual, err := getApiResponseJson(ctx)
	if err != nil {
		return err
	}

	found, err := jmespath.Search(path, actual)
	if err != nil {
		return fmt.Errorf("failed to search path %s: %s", path, err)
	}

	if found == nil {
		return fmt.Errorf(`expected response body to have JMESPath "%s", but actual is %s`, path, res.Body.String())
	}

	if fmt.Sprintf("%v", found) != value {
		return fmt.Errorf(`expected response body to have JMESPath "%s" with value "%s", but actual is %v`, path, value, res)
	}

	return nil
}

func setupHttpServer(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	rest, err := app.InitializeTest()
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, apiServerCtxKey{}, rest), nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(setupHttpServer)

	ctx.Step(`^I make a GET request to "([^"]*)"$`, iMakeAGETRequestTo)
	ctx.Step(`^I make a (POST|PUT) request to "([^"]*)" with the body$`, iMakeARequestToWithTheBody)
	ctx.Step(`^the response status code should be (\d+)$`, theResponseStatusCodeShouldBe)
	ctx.Step(`^the response body should be$`, theResponseBodyShouldBe)
	ctx.Step(`^the response json should have "([^"]*)"$`, theResponseJsonShouldHave)
	ctx.Step(`^the response json should have "([^"]*)" with value "([^"]*)"$`, theResponseJsonShouldHaveWithValue)
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
