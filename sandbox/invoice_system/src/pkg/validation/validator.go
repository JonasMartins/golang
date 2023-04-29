package validation

import (
	"fmt"
	"net/mail"
	"regexp"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func FieldAndError(fieldName string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       fieldName,
		Description: err.Error(),
	}
}

func InvalidRequestEror(v []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{FieldViolations: v}
	statusInvalid := status.New(codes.InvalidArgument, "invalid params")
	sd, err := statusInvalid.WithDetails(badRequest)
	if err != nil {
		return statusInvalid.Err()
	}

	return sd.Err()
}

func ValidateRequired(arg any, fieldName string) error {
	if arg == nil {
		return fmt.Errorf("field %s is required", fieldName)
	}
	return nil
}

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lowercase letters, digits, or underscore")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidFullName(value) {
		return fmt.Errorf("must contain only letters or spaces")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}
