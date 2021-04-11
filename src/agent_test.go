package go_learning

import "context"

type DemoCollector struct {
	evtReceiver EventReceiver
	agtCtx context.Context
	stopChan chan struct{}
	name string
	content string
}

func NewCollect(name string, content string) *DemoCollector {
	return &DemoCollector {
		stopChan: make(chan struct{}),
		name: name,
		content: content,
	}
}


