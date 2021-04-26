package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/progmatic-99/gRPC/protos/currency"
)

type Currency struct {
	log hclog.Logger
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Hnadle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
}
