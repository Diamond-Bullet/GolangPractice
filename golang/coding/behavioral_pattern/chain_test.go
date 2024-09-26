package behavioral_pattern

import (
	"context"
)

/*
Chain of responsibility
usually comes with Composite pattern. In this case, the parent component is the successor of the component.

Here are many ways to implement it. For a request, you can flexibly register multiple handlers.
1. Based on List: for instance, `before_update` and `after_update` in GORM, `handlers` and `c.Next()` in Gin framework.
2. Each object store the reference of the successor.
*/

/*
Implementation 1: List
*/

type ChainOnArray struct {
	ctx      context.Context
	index    int
	handlers []func(context.Context)
}

func (h *ChainOnArray) Register(f func(context.Context)) {
	h.handlers = append(h.handlers, f)
}

func (h *ChainOnArray) Next() {
	if h.index < len(h.handlers)-1 {
		h.index++
		h.handlers[h.index](h.ctx)
	}
}

/*
Implementation 2: Reference
*/

type ChainOnLinkedList interface {
	Handle(c context.Context)
}

type Handler struct {
	Successor ChainOnLinkedList
}

func (h *Handler) Handle(c context.Context) {
	if c.Value("Handler can handle").(string) == "true" {
		return
	}
	h.Successor.Handle(c)
}
