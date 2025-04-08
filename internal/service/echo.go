package service

import (
	"echo/internal/model"
	"fmt"
)

type EchoService struct{}

func NewEchoService() *EchoService {
	return &EchoService{}
}

func (s *EchoService) Echo(msg *model.OneBotMessage) string {
	return fmt.Sprintf("Echo: %s", "hello world")
}
