package context

import (
	"context"
	"github.com/networm6/gopherBox/shutdown"
)

func Opened(_ctx context.Context) bool {
	select {
	case <-_ctx.Done():
		return false
	default:
		return true
	}
}

func Wait() {
	shutdown.New().Listen()
}
