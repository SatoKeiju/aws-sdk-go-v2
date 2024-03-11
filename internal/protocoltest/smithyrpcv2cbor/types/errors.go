// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// This error is thrown when a request is invalid.
type ComplexError struct {
	Message *string

	ErrorCodeOverride *string

	TopLevel *string
	Nested   *ComplexNestedErrorData

	noSmithyDocumentSerde
}

func (e *ComplexError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ComplexError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ComplexError) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ComplexError"
	}
	return *e.ErrorCodeOverride
}
func (e *ComplexError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This error is thrown when an invalid greeting value is provided.
type InvalidGreeting struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidGreeting) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidGreeting) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidGreeting) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidGreeting"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidGreeting) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A standard error for input validation failures. This should be thrown by
// services when a member of the input structure falls outside of the modeled or
// documented constraints.
type ValidationException struct {
	Message *string

	ErrorCodeOverride *string

	FieldList []ValidationExceptionField

	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ValidationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
