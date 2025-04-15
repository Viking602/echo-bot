package moyu

import (
	"github.com/imroc/req/v3"
)

type Client struct {
	client *req.Client
}

func NewClient() *Client {
	return &Client{
		client: req.C(),
	}
}
