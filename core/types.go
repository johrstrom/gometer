package core

type SampleResult struct {
	Response string
	Pass     bool
}

type Sampler interface {
	Sample() SampleResult
}

type Properties map[string]interface{}
