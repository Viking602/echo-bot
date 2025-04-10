package bilibili

import (
	"github.com/imroc/req/v3"
	"time"
)

type Client struct {
	client *req.Client
}

func NewClient() *Client {
	client := req.C().
		SetTLSFingerprintChrome().
		ImpersonateChrome().
		SetCommonRetryCount(3).
		SetCommonRetryFixedInterval(5 * time.Second)

	return &Client{
		client: client,
	}
}
