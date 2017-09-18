package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicUse(test *testing.T) {
	tp := NewTestPlan("testName")
	root := tp.GetRootNode()
	assert.NotNil(test, root, "Get Root node cannot be nil")

	tp = &TestPlan{Name: "testName"} //test root not null even when user creates the struct
	root = tp.GetRootNode()
	assert.NotNil(test, root, "Get Root node cannot be nil")

	tp.AddThreadGroup()
	assert.NotEmpty(test, root.SubTree, "SubTree should not be empty")

}

func TestAddNodesUnder(test *testing.T) {
	tp := NewTestPlan("testName")
	group := tp.AddThreadGroup()

	fooSampler := NewTestElementNode(NewTestSampler())
	fooProps := fooSampler.TestElement.Properties()
	fooProps["foo"] = "bar"

	group.AddNodeUnder(fooSampler)
	assert.NotNil(test, group.SubTree, "ThreadGroup's SubTree should not be nil")
	assert.NotEmpty(test, group.SubTree, "ThreadGroup's SubTree should not be empty")

	assert.Equal(test, fooSampler, group.SubTree[0], "First element of ThreadGroup's Tree should be a generic sampler")

}
