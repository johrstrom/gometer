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

	root.AddNodeUnder(group)
	return group
}

func (plan *TestPlan) initRoot() {
	root := &TestElementNode{
		TestElement: NewRoot(),
	}
	plan.Tree = append(plan.Tree, root)
}

// AddNodeUnder Adds a node under the node given as the argument
// i.e., within it's scope
func (parent *TestElementNode) AddNodeUnder(child *TestElementNode) {
	parent.SubTree = append(parent.SubTree, child)
}

// NewTestElementNode creates a new node in the tree
func NewTestElementNode(ele TestElement) *TestElementNode {
	return &TestElementNode{
		SubTree:     make([]*TestElementNode, 0),
		TestElement: ele,
	}
}
