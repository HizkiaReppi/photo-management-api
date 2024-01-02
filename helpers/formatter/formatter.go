package formatter

import (
	"github.com/asaskevich/govalidator"
)

// Response represents the structure of the API response.
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// Meta contains metadata information for the API response.
type Meta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// ApiResponse creates a standardized API response.
func ApiResponse(code int, status string, data interface{}, message string) Response {
	meta := Meta{
		Code:    code,
		Status:  status,
		Message: message,
	}

	return Response{
		Meta: meta,
		Data: data,
	}
}

// FormatValidationError formats validation errors into a slice of strings.
func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(govalidator.Errors).Errors() {
		errors = append(errors, e.Error())
	}

	return errors
}
