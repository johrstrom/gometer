package samplers

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/johrstrom/gometer/core"
)

const (
	httpMethod string = "HttpSampler.method"
	httpURL    string = "HttpSampler.url"
)

// HTTPSampler the struct for making http samples.
type HTTPSampler struct {
	properties    core.Properties
	client        *http.Client
	payloadBuffer *bytes.Buffer
}

// DefaultHTTPSampler gives the basic most http sampler
func DefaultHTTPSampler() *HTTPSampler {
	sampler := &HTTPSampler{
		properties:    make(core.Properties),
		client:        http.DefaultClient,
		payloadBuffer: bytes.NewBuffer(make([]byte, 100000)),
	}

	return sampler
}

// Properties to implement the TestElement interface
func (sampler *HTTPSampler) Properties() core.Properties {
	return sampler.properties
}

// Sample runs an http request to get a Sample
func (sampler *HTTPSampler) Sample() core.SampleResult {

	request, err := sampler.createRequest()
	if err != nil {
		return failResult(err.Error())
	}

	response, err := sampler.client.Do(request)
	if err != nil {
		return failResult(err.Error())
	}

	result := core.SampleResult{}
	code := response.StatusCode

	if code >= 200 && code < 300 {
		result.Pass = true
	} else {
		result.Pass = false
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return failResult(err.Error())
	}

	result.Response = string(responseBody)

	return result
}

func failResult(message string) core.SampleResult {
	return core.SampleResult{
		Pass:     false,
		Response: message,
	}
}

func (sampler *HTTPSampler) createRequest() (*http.Request, error) {
	if method, ok := sampler.properties[httpMethod].(string); ok {
		if url, ok := sampler.properties[httpURL].(string); ok {
			return http.NewRequest(method, url, sampler.payloadBuffer)
		}
		return nil, errors.New("Invalid URL parameter")
	}

	return nil, errors.New("Invalid http method")
}
