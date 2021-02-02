package service

import (
	"context"
	"fmt"
)

type ManageService interface {
	Open(ctx context.Context) error
	Close()
}

type manageService struct {
	stopCh <-chan struct{}
}

func NewManageService(stopCh <-chan struct{}) (ManageService, error) {
	return &manageService{
		stopCh: stopCh,
	}, nil
}

func (base *manageService) Open(ctx context.Context) error {
	fmt.Printf("###ManageService Opened!\n")
	return nil
}

func (base *manageService) Close() {
	fmt.Printf("###ManageService Closed!\n")
	return
}
