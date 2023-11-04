package lifecycle

import "context"

type LifeInterface interface {
	Start()
	Destroy()
}

type LifeClass struct {
	LifeInterface
	_ctx    context.Context
	_cancel context.CancelFunc
}
