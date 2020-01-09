package mux

import "net/http"

type Options struct {
	Port                 int
	PathRouterHandlers   MuxHandlers
	PrefixRouterHandlers MuxHandlers
}

type MuxHandlers map[string]http.Handler

func (m1 MuxHandlers) Merge(handlers MuxHandlers) {

	for p, h := range handlers {
		m1.Append(p, h)
	}
}

func (m1 MuxHandlers) Append(routePath string, handler http.Handler) {
	m1[routePath] = handler
}
