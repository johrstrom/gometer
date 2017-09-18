package test

import (
	"github.com/johrstrom/gometer/core"
	"github.com/johrstrom/gometer/samplers"
)

// NewSimpleTestPlan is a way to get a new simple test plan
func NewSimpleTestPlan() *core.TestPlan {
	tp := core.NewTestPlan("simple_testplan")
	group := tp.AddThreadGroup()

	httpSampler := core.NewTestElementNode(samplers.DefaultHTTPSampler())
	group.AddNodeUnder(httpSampler)

	return tp
}
