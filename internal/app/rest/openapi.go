package rest

import (
	"bytes"
	_ "embed"
	"io"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/google/wire"
)

//go:embed openapi.yaml
var openapiData []byte

var OpenApiSet = wire.NewSet(
	NewOpenApi,
	NewOpenApiRouter,
	ProvideOpenApiMiddleware,
)

type OpenApiMiddleware func(next http.Handler) http.Handler

func NewOpenApi() (*openapi3.T, error) {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(openapiData)
	if err != nil {
		return nil, err
	}

	if err := doc.Validate(loader.Context); err != nil {
		return nil, err
	}

	return doc, nil
}

func NewOpenApiRouter(doc *openapi3.T) (routers.Router, error) {
	return gorillamux.NewRouter(doc)
}

type OpenApiResponseWriter struct {
	http.ResponseWriter
	statusCode   int
	responseBody []byte
}

func (w *OpenApiResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *OpenApiResponseWriter) Write(b []byte) (int, error) {
	w.responseBody = b
	return w.ResponseWriter.Write(b)
}

func ProvideOpenApiMiddleware(router routers.Router) OpenApiMiddleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			route, pathParams, err := router.FindRoute(r)

			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			requestValidationInput := &openapi3filter.RequestValidationInput{
				Request:    r,
				PathParams: pathParams,
				Route:      route,
			}
			err = openapi3filter.ValidateRequest(r.Context(), requestValidationInput)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			responseWriter := &OpenApiResponseWriter{ResponseWriter: w}
			next.ServeHTTP(responseWriter, r)

			responseValidationInput := &openapi3filter.ResponseValidationInput{
				RequestValidationInput: requestValidationInput,
				Status:                 responseWriter.statusCode,
				Header:                 responseWriter.Header(),
				Body:                   io.NopCloser(bytes.NewReader(responseWriter.responseBody)),
			}
			err = openapi3filter.ValidateResponse(r.Context(), responseValidationInput)
			if err != nil {
				panic(err)
				return
			}
		})
	}
}
