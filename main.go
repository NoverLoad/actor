package main

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
)

type SetState struct {
	value uint
}
type ResetState struct{}
type Handler struct {
	state uint
}

func newHandler() actor.Receiver {
	return &Handler{}
}

func (h *Handler) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case ResetState:
		h.state = 0
		fmt.Println("handler received reset state", h.state)
	case SetState:
		h.state = msg.value
		fmt.Println("handler received new state", h.state)
	case actor.Initialized:
		h.state = 10
		fmt.Println("handler Initialized, my state", h.state)
	case actor.Started:
		fmt.Println("handler started")
	case actor.Stopped:

	}
}

func main() {
	systemConfig := actor.NewEngineConfig()
	e, _ := actor.NewEngine(systemConfig)
	pid := e.Spawn(newHandler, "handler")
	for i := 0; i < 10; i++ {
		go e.Send(pid, SetState{value: uint(i)})
	}
	e.Send(pid, ResetState{})
}
