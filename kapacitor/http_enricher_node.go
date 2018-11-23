package kapacitor

import (
	"github.com/influxdb/kapacitor/pipeline"
)

type MyCustomNode struct {
	// Include the generic node implementation
	node
	// Keep a reference to the pipeline node
	h *pipeline.MyCustomNode
}

func newMyCustomNode(et *ExecutingTask, n *pipeline.MyCustomNode) (*MyCustomNode, error) {
	h := &MyCustomNode{
		// pass in necessary fields to the 'node' struct
		node: node{Node: n, et: et},
		// Keep a reference to the pipeline.MyCustomNode
		h: n,
	}
	// Set the function to be called when running the node
	h.node.runF = h.runOut
	return nil, h
}

func (h *MyCustomNode) runOut() error {
	switch h.Wants() {
	case pipeline.StreamEdge:
		// Read stream data and send to HouseDB
		for p, ok := h.ins[0].NextPoint(); ok; p, ok = h.ins[0].NextPoint() {
			// Turn the point into a batch with just one point.
			batch := models.Batch{
				Name:   p.Name,
				Group:  p.Group,
				Tags:   p.Tags,
				Points: []models.TimeFields{{Time: p.Time, Fields: p.Fields}},
			}
			// Write the batch
			err := h.write(batch)
			if err != nil {
				return err
			}
		}
	case pipeline.BatchEdge:
		// Read batch data and send to HouseDB
		for b, ok := h.ins[0].NextBatch(); ok; b, ok = h.ins[0].NextBatch() {
			// Write the batch
			err := h.write(b)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Write a batch of data to HouseDB
func (h *MyCustomNode) write(batch models.Batch) error {
	// Implement writing to HouseDB here...
	return nil
}
