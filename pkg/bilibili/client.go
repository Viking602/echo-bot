package bilibili

import (
	"echo/pkg/logger"
	"github.com/imroc/req/v3"
	"time"
)

type Client struct {
	client *req.Client
	logger *logger.Logger
}

func NewClient(logger *logger.Logger) *Client {
	client := req.C().
		SetLogger(logger).
		SetTLSFingerprintChrome().
		ImpersonateChrome().
		SetCommonRetryCount(3).
		SetCommonRetryFixedInterval(5 * time.Second)

	return &Client{
		client: client,
		logger: logger,
	}
}
