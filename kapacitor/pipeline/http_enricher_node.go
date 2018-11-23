package pipeline

type MyCustomNode struct {
    // Include pipeline.chainnode so we have all the chaining methods available
    // to our new node
    chainnode

    // Mark this field as ignored for docs
    // Since it is set via the Names method below
    // tick:ignore
    NameList []string `tick:"Names"`

}

func newMyCustomNode(e EdgeType, n Node) *MyCustomNode {
    m := &MyCustomNode{
        chainnode: newBasicChainNode("mycustom", e, e),
    }
    n.linkChild(m)
    return m
}

// Set the NameList field on the node via this method.
//
// Example:
//    node.names('name0', 'name1')
//
// Use the tickdoc comment 'tick:property' to mark this method
// as a 'property method'
// tick:property
func (m *MyCustomNode) Names(name ...string) *MyCustomNode {
    m.NameList = name
    return m
}