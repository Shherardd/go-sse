package events

import (
	"fmt"
	"net/http"
	"sync"
)

type EventMessage struct {
	EventName string
	Data      any
}

type HandlerEvent struct {
	m        sync.Mutex
	clientes map[string]*client
}

func NewHandlerEvent() *HandlerEvent {
	return &HandlerEvent{
		clientes: make(map[string]*client),
	}
}

func (h *HandlerEvent) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	id := r.URL.Query().Get("id")

	c := newClient(id)
	h.AddClient(c)
	fmt.Println("Client connected: ", id)
	c.OnLine(r.Context(), w, w.(http.Flusher))
	fmt.Println("Client disconnected: ", id)
	h.RemoveClient(id)
}

func (h *HandlerEvent) AddClient(c *client) {
	h.m.Lock()
	defer h.m.Unlock()
	h.clientes[c.ID] = c
}

func (h *HandlerEvent) RemoveClient(id string) {
	h.m.Lock()
	defer h.m.Unlock()
	delete(h.clientes, id)
}

func (h *HandlerEvent) BroadCast(m EventMessage) {
	h.m.Lock()
	defer h.m.Unlock()
	for _, c := range h.clientes {
		c.sendMessage <- m
	}
}
