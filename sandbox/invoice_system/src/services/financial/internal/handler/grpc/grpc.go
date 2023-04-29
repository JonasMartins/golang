package grpc

import (
	"context"
	"project/src/gen"
	validator "project/src/pkg/validation"
	controller "project/src/services/financial/internal/controller/financial"
	financialValidator "project/src/services/financial/internal/validation"
)

type Handler struct {
	gen.UnimplementedFinancialServiceServer
	ctrl *controller.Controller
}

func New(ctrl *controller.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) AddInvoice(ctx context.Context, req *gen.AddInvoiceRequest) (*gen.BasicResponse, error) {
	v := financialValidator.ValidateAddInvoices(req)
	if v != nil {
		return nil, validator.InvalidRequestEror(v)
	}
	i, err := h.ctrl.AddInvoice(ctx, req)
	if err != nil {
		return nil, err
	}
	return i, nil
}
