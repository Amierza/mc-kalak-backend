package external

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type (
	IExternalGateway interface {
	}

	externalGateway struct {
		baseURL string
		client  *http.Client
		logger  *zap.Logger
	}
)

func NewExternalGateway(baseURL string, logger *zap.Logger) *externalGateway {
	return &externalGateway{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 10 * time.Second},
		logger:  logger,
	}
}
