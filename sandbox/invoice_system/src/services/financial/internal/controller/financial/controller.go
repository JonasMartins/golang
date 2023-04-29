package financial

import (
	"context"
	"project/src/gen"
)

type financialRepository interface {
	AddInvoice(ctx context.Context, params *gen.AddInvoiceRequest) (*gen.BasicResponse, error)
}

type Controller struct {
	financialRepository financialRepository
}

func New(financialRepository financialRepository) *Controller {
	return &Controller{financialRepository}
}

func (c *Controller) AddInvoice(ctx context.Context, params *gen.AddInvoiceRequest) (*gen.BasicResponse, error) {
	response, err := c.financialRepository.AddInvoice(ctx, params)
	if err != nil {
		return nil, err
	}
	return response, nil
}
