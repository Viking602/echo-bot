package douyu

import (
	"echo/pkg/logger"
	"github.com/imroc/req/v3"
	"time"
)

type Client struct {
	client *req.Client
	log    *logger.Logger
}

func NewClient(log *logger.Logger) *Client {
	client := req.C().
		SetTLSFingerprintChrome().
		ImpersonateChrome().
		SetCommonRetryCount(3).
		SetCommonRetryFixedInterval(5 * time.Second)

	return &Client{
		client: client,
		log:    log,
	}
}
