package core

const (
	typeProperty string = "core.type"
)

// SampleResult the result of samples
type SampleResult struct {
	Response string
	Pass     bool
}

// TestPlan - the struct that is the script for running tests
type TestPlan struct {
	Name string
	Tree []*TestElementNode
}

// Sampler the interface all Samplers need to implement
type Sampler interface {
	Sample() SampleResult
	TestElement
}

// TestElement the interface for all test elements
type TestElement interface {
	Properties() Properties
}

// RootTestElement is the simplest implementation of a test element
type RootTestElement struct {
	props Properties
}

func (root *RootTestElement) Properties() Properties {
	return root.props
}

func NewRoot() *RootTestElement {
	p := make(Properties)
	p[typeProperty] = "root"

	return &RootTestElement{
		props: p,
	}
}

// Properties for TestElements
type Properties map[string]interface{}

// TestElementNode the backing data structure for test scripts
type TestElementNode struct {
	TestElement
	SubTree []*TestElementNode
}
