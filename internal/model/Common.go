// Package this contains Model definitions
package model

// GenericResponse generic response
type GenericResponse struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason,omitempty"`
}

// GenericRes response model
//
// swagger:response GenericRes
type GenericRes struct {
	// in: body
	// required: true
	GenericResponse GenericResponse
}
