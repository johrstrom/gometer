package core

var (
// TestElementTypes are all the test element types
//TestElementTypes map[string]string = AllElementTypes()
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

// SimpleTestElement is the simplest implementation of a test element
type SimpleTestElement struct {
	Properties
}

// Properties for TestElements
type Properties map[string]interface{}

// TestElementNode the backing data structure for test scripts
type TestElementNode struct {
	Type        string
	TestElement TestElement
	SubTree     []*TestElementNode
}

// AllElementTypes is all the different types of test elements
//func AllElementTypes() map[string]string {
//	m := make(map[string]string)
//	m["ThreadGroup"] = "ThreadGroup"

//return m
//}
