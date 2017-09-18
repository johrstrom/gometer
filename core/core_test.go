package core

type TestSampler struct {
	props Properties
}

func NewTestSampler() *TestSampler {
	return &TestSampler{
		props: make(Properties),
	}
}

func (ts *TestSampler) Sample() SampleResult {
	return SampleResult{}
}

func (ts *TestSampler) Properties() Properties {
	return ts.props
}

func NewSimpleTestPlan() *TestPlan {
	tp := NewTestPlan("simple_testplan")
	group := tp.AddThreadGroup()

	fooSampler := NewTestElementNode(NewTestSampler())
	group.AddNodeUnder(fooSampler)

	return tp
}
