package core

// ThreadGroup is a group of threads
type ThreadGroup struct {
	props Properties
}

// Properties to implement TestElement
func (t *ThreadGroup) Properties() Properties {
	return t.props
}
