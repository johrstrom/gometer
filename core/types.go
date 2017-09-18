package core

const (
	// TypeProperty the property for type
	TypeProperty string = "core.type"

	// SamplerType the key to be use for samplers
	SamplerType string = "sampler"

	// ThreadGroupType the key to be used for ThreadGroups
	ThreadGroupType string = "thread_group"

	// RootNodeType the key to be used for root nodes
	RootNodeType string = "root"
)

// SampleResult the result of samples
type SampleResult struct {
	Response string
	Pass     bool
}

// TestPlan - the struct that is the script for running tests
type TestPlan struct {
	Name string
	root *TestElementNode
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

// NewRoot creates a new Root Element
func NewRoot() *RootTestElement {
	p := make(Properties)
	p[TypeProperty] = RootNodeType

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
