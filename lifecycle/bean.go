package lifecycle

import "context"

type LifeInterface interface {
	Start()
	Destroy()
}

type LifeClass[T interface{}] struct {
	LifeInterface
	_ctx    context.Context
	_cancel context.CancelFunc
}
