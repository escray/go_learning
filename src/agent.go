package go_learning

import (
	"context"
	"errors"
	"strings"
)

const(
	Waiting = iota
	Running
)

var ErrorWrongState = errors.New("can not take the operation in the current state")

type CollectorsError struct {
	CollectorErrors []error
}

func (ce CollectorsError) Error() string {
	var strs []string
	for _, err := range ce.CollectorErrors {
		strs = append(strs, err.Error())
	}
	return strings.Join(strs, ";")
}

type Event struct {
	Source string
	Content string
}

type EventReceiver interface {
	OnEvent(evt Event)
}

type Collector interface {
	Init(evtReceiver EventReceiver) error
	Start(agtCtx context.Context) error
	Stor() error
	Destroy() error
}

type Agent struct {
	collectors map[string]Collector
	evtBuf chan Event
	cancel context.CancelFunc
	ctx context.Context
	state int
}


