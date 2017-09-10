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
func (plan *TestPlan) AddThreadGroup() {
	root := plan.GetRootNode()
	root.AddNodeUnder(NewThreadGroupNode())
}

func (plan *TestPlan) initRoot() {
	root := &TestElementNode{}
	plan.Tree = append(plan.Tree, root)
}

// AddNodeUnder Adds a node under the node given as the argument
// i.e., within it's scope
func (parent *TestElementNode) AddNodeUnder(child *TestElementNode) {
	arr := []*TestElementNode{child}
	parent.SubTree = append(arr, parent.SubTree...)
}

// NewThreadGroupNode creates a new thread group node
func NewThreadGroupNode() *TestElementNode {

	node := &TestElementNode{
		Type:        "thread_group",
		SubTree:     make([]*TestElementNode, 0),
		TestElement: &ThreadGroup{},
	}

	return node
}
