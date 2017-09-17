package core

// NewTestPlan creates a new test plan with a given name
func NewTestPlan(name string) *TestPlan {
	plan := &TestPlan{
		Name: name,
		Tree: make([]*TestElementNode, 0),
	}

	plan.initRoot()

	return plan
}

// GetRootNode gets the root node of this test plan.
func (plan *TestPlan) GetRootNode() *TestElementNode {
	if plan.Tree == nil || len(plan.Tree) == 0 {
		plan.initRoot()
	}

	return plan.Tree[0]
}

// AddThreadGroup adds a thread group to the test plan
func (plan *TestPlan) AddThreadGroup() *TestElementNode {
	root := plan.GetRootNode()
	group := newThreadGroupNode()

	root.AddNodeUnder(newThreadGroupNode())
	return group
}

func (plan *TestPlan) initRoot() {
	root := &TestElementNode{
		Type: "root",
	}
	plan.Tree = append(plan.Tree, root)
}

// AddNodeUnder Adds a node under the node given as the argument
// i.e., within it's scope
func (parent *TestElementNode) AddNodeUnder(child *TestElementNode) {
	arr := []*TestElementNode{child}
	parent.SubTree = append(arr, parent.SubTree...)
}

// NewThreadGroupNode creates a new thread group node
func newThreadGroupNode() *TestElementNode {

	node := &TestElementNode{
		Type:        "thread_group",
		SubTree:     make([]*TestElementNode, 0),
		TestElement: &ThreadGroup{},
	}

	return node
}

// NewSamplerNode creates a new node in the tree
func NewSamplerNode(sampler Sampler) *TestElementNode {
	return &TestElementNode{
		Type:        "sampler",
		SubTree:     make([]*TestElementNode, 0),
		TestElement: sampler,
	}
}
