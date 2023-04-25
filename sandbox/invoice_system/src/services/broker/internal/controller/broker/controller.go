package broker

import "context"

type brokerGateway interface {
	AddInvoice(ctx context.Context) error
	PayInvoice(ctx context.Context) error
}

type Controller struct {
	brokerGateway brokerGateway
}

func New(bg brokerGateway) *Controller {
	return &Controller{bg}
}

func (c *Controller) AddInvoice(ctx context.Context) error {
	err := c.brokerGateway.AddInvoice(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) PayInvoice(ctx context.Context) error {
	err := c.brokerGateway.PayInvoice(ctx)
	if err != nil {
		return err
	}
	return nil
}
