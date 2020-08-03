package errors

import (
	"go.uber.org/zap/zapcore"
)

type Operation string

type ErrorType string

const (
	NotFoundError     ErrorType = "NOT_FOUND"
	UnAuthorizedError ErrorType = "UNAUTHORIZED"
	Unexpected        ErrorType = "UNEXPECTED"
)

type Error struct {
	operations []Operation // Gives Sequence of Operations Performed
	errorType  ErrorType
	error      string
	severity   zapcore.Level
	requestID  string
}

//NewError - Generate New Instance of Error
func NewError(operation Operation, errorType ErrorType, err string, severity zapcore.Level, requestID string) *Error {
	return &Error{
		operations: []Operation{operation},
		errorType:  errorType,
		error:      err,
		severity:   severity,
		requestID:  requestID,
	}
}

//WithOperation - Passing the error to the upper level
func (e *Error) WithOperation(operation Operation) *Error {
	e.operations = append(e.operations, operation)
	return e
}

func (e *Error) Operations() []Operation {
	return e.operations
}

func (e *Error) ErrorType() ErrorType {
	return e.errorType
}

func (e *Error) Error() string {
	return e.error
}

func (e *Error) Severity() zapcore.Level {
	return e.severity
}

func (e *Error) RequestID() string {
	return e.requestID
}
