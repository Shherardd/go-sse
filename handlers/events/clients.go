package events

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type client struct {
	ID          string
	sendMessage chan EventMessage
}

func newClient(id string) *client {
	return &client{
		ID:          id,
		sendMessage: make(chan EventMessage),
	}
}

func (c *client) OnLine(ctx context.Context, w io.Writer, flusher http.Flusher) {
	for {
		select {
		case <-ctx.Done():
			return
		case m := <-c.sendMessage:
			data, err := json.Marshal(m.Data)
			if err != nil {
				log.Println(err)
			}
			const format = "event: %s\ndata: %s\nID:%s\n\n"
			fmt.Fprintf(w, format, m.EventName, data, 5)
			flusher.Flush()
		}
	}
}
