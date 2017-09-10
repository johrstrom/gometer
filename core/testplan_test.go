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
