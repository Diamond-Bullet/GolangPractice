package design_pattern

import (
	"context"
)

/*
职责链模式，Chain of responsibility
常与composite一起使用，此时父构件作为对象的后继

实现方式和应用较多，对于一个请求，可以灵活地注册若干个处理器
1. 基于列表的方式，例如gorm，before_update， after_update 实现 或者 gin，handlers && c.Next 实现
2. 每个对象保持对链上下个对象的引用
*/

/*
实现1，列表的形式
*/

type HandlerChain struct {
	ctx      context.Context
	index    int
	handlers []func(context.Context)
}

func (h *HandlerChain) Register(f func(context.Context)) {
	h.handlers = append(h.handlers, f)
}

func (h *HandlerChain) Next() {
	if h.index < len(h.handlers)-1 {
		h.index++
		h.handlers[h.index](h.ctx)
	}
}

/*
实现2，保持引用
*/

type HandlerChain2 interface {
	Handle(c context.Context)
}

type H1 struct {
	Subsequence HandlerChain2
}

func (h *H1) Handle(c context.Context) {
	if c.Value("H1 can handle").(string) == "true" {
		return
	}
	h.Subsequence.Handle(c)
}
