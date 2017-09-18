package test

import "github.com/johrstrom/gometer/core"

// SimpleSampler is a simple test sampler
type SimpleSampler struct {
	props core.Properties
}

// NewSimpleSampler creates a new SimpleSampler
func NewSimpleSampler() *SimpleSampler {
	return &SimpleSampler{
		props: make(core.Properties),
	}
}

// Sample to implement Sampler
func (sampler *SimpleSampler) Sample() core.SampleResult {
	return core.SampleResult{}
}

// Properties to implement Test element interface
func (sampler *SimpleSampler) Properties() core.Properties {
	return sampler.props
}

// NewSimpleTestPlan is a way to get a new simple test plan
func NewSimpleTestPlan() *core.TestPlan {
	tp := core.NewTestPlan("simple_testplan")
	group := tp.AddThreadGroup()

	fooSampler := core.NewTestElementNode(NewSimpleSampler())
	group.AddNodeUnder(fooSampler)

	return tp
}
