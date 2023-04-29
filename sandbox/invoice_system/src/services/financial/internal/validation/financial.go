package validation

import (
	"project/src/gen"
	validator "project/src/pkg/validation"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func ValidateAddInvoices(req *gen.AddInvoiceRequest) (v []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateRequired(req.GetClientId(), "clientId"); err != nil {
		v = append(v, validator.FieldAndError("ClientId", err))
	}
	if err := validator.ValidateRequired(req.GetDueDate(), "dueDate"); err != nil {
		v = append(v, validator.FieldAndError("DueDate", err))
	}
	if err := validator.ValidateRequired(req.GetValue(), "value"); err != nil {
		v = append(v, validator.FieldAndError("Value", err))
	}
	return v
}
