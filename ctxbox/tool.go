package ctxbox

import (
	"context"
)

func Opened(_ctx context.Context) bool {
	select {
	case <-_ctx.Done():
		return false
	default:
		return true
	}
}
