package core

// ThreadGroup is a group of threads
type ThreadGroup struct {
	props Properties
}

// Properties to implement TestElement
func (tg *ThreadGroup) Properties() Properties {
	return tg.props
}

// NewThreadGroupNode creates a new thread group node
func newThreadGroupNode() *TestElementNode {
	p := make(Properties)
	p[TypeProperty] = ThreadGroupType

	node := &TestElementNode{
		SubTree: make([]*TestElementNode, 0),
		TestElement: &ThreadGroup{
			props: p,
		},
	}

	return node
}
