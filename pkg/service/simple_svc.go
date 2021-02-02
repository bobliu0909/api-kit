package service

import (
	"context"
)

type SimpleService interface {
	Register(ctx context.Context, node string) error
	UnRegister(ctx context.Context, node string) error
	Release(node string)
}

type simpleService struct {
	stopCh <-chan struct{}
}

func NewSimpleService(stopCh <-chan struct{}) (SimpleService, error) {
	return &simpleService{
		stopCh: stopCh,
	}, nil
}

func (simple *simpleService) Register(ctx context.Context, node string) error {
	return nil
}

func (simple *simpleService) UnRegister(ctx context.Context, node string) error {
	return nil
}

func (simple *simpleService) Release(node string) {
	return
}