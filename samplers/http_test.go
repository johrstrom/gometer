package samplers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	getGoodSampleURL = "/get/good/sample"
)

type TestHTTPHandler struct {
	funcMap map[string]func(http.ResponseWriter, *http.Request)
}

var testServer = httptest.NewServer(createTestHandler())

func TestGetSamplerWithGoodParams(test *testing.T) {
	sampler := DefaultHTTPSampler()

	sampler.Properties[httpMethod] = http.MethodGet
	sampler.Properties[httpURL] = testServer.URL + getGoodSampleURL

	result := sampler.Sample()
	assert.Equal(test, true, result.Pass)
	assert.Equal(test, "{ \"foo\": \"bar\"}", result.Response)

}

func TestGetSampler404(test *testing.T) {
	sampler := DefaultHTTPSampler()

	sampler.Properties[httpMethod] = http.MethodGet
	sampler.Properties[httpURL] = testServer.URL + "/doesnt/exist"

	result := sampler.Sample()
	assert.Equal(test, false, result.Pass)
	assert.Equal(test, "Not Found", result.Response)

}

func TestGetSamplerCantConnect(test *testing.T) {
	sampler := DefaultHTTPSampler()

	sampler.Properties[httpMethod] = http.MethodGet
	sampler.Properties[httpURL] = "http://localhost:23"

	result := sampler.Sample()
	assert.Equal(test, false, result.Pass)
	assert.Equal(test, "Get http://localhost:23: dial tcp [::1]:23: getsockopt: connection refused", result.Response)
}

func TestGetSamplerBadParameters(test *testing.T) {
	sampler := DefaultHTTPSampler()

	result := sampler.Sample()
	assert.Equal(test, false, result.Pass)
	assert.Equal(test, "Invalid http method", result.Response)

	sampler.Properties[httpMethod] = http.MethodGet
	sampler.Properties[httpURL] = 1000
	result = sampler.Sample()

	assert.Equal(test, false, result.Pass)
	assert.Equal(test, "Invalid URL parameter", result.Response)
}

func createTestHandler() *TestHTTPHandler {
	handler := &TestHTTPHandler{
		funcMap: make(map[string]func(writer http.ResponseWriter, request *http.Request)),
	}

	handler.funcMap[getGoodSampleURL] = getGoodSampleFunc

	return handler
}

func (handler *TestHTTPHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path

	if handlerFunc, ok := handler.funcMap[path]; ok {
		handlerFunc(writer, request)
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Not Found"))
	}
}

func getGoodSampleFunc(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, "{ \"foo\": \"bar\"}")
}
