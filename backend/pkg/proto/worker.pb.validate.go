// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/worker.proto

package proto

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetPolynomialCalculationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetPolynomialCalculationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPolynomialCalculationRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetPolynomialCalculationRequestMultiError, or nil if none found.
func (m *GetPolynomialCalculationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPolynomialCalculationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Matrix

	if len(errors) > 0 {
		return GetPolynomialCalculationRequestMultiError(errors)
	}

	return nil
}

// GetPolynomialCalculationRequestMultiError is an error wrapping multiple
// validation errors returned by GetPolynomialCalculationRequest.ValidateAll()
// if the designated constraints aren't met.
type GetPolynomialCalculationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPolynomialCalculationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPolynomialCalculationRequestMultiError) AllErrors() []error { return m }

// GetPolynomialCalculationRequestValidationError is the validation error
// returned by GetPolynomialCalculationRequest.Validate if the designated
// constraints aren't met.
type GetPolynomialCalculationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPolynomialCalculationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPolynomialCalculationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPolynomialCalculationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPolynomialCalculationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPolynomialCalculationRequestValidationError) ErrorName() string {
	return "GetPolynomialCalculationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPolynomialCalculationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPolynomialCalculationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPolynomialCalculationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPolynomialCalculationRequestValidationError{}

// Validate checks the field values on GetPolynomialCalculationResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetPolynomialCalculationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPolynomialCalculationResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetPolynomialCalculationResponseMultiError, or nil if none found.
func (m *GetPolynomialCalculationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPolynomialCalculationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Operation

	// no validation rules for ResultMatrix

	// no validation rules for Time

	// no validation rules for OperationType

	if len(errors) > 0 {
		return GetPolynomialCalculationResponseMultiError(errors)
	}

	return nil
}

// GetPolynomialCalculationResponseMultiError is an error wrapping multiple
// validation errors returned by
// GetPolynomialCalculationResponse.ValidateAll() if the designated
// constraints aren't met.
type GetPolynomialCalculationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPolynomialCalculationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPolynomialCalculationResponseMultiError) AllErrors() []error { return m }

// GetPolynomialCalculationResponseValidationError is the validation error
// returned by GetPolynomialCalculationResponse.Validate if the designated
// constraints aren't met.
type GetPolynomialCalculationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPolynomialCalculationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPolynomialCalculationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPolynomialCalculationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPolynomialCalculationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPolynomialCalculationResponseValidationError) ErrorName() string {
	return "GetPolynomialCalculationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetPolynomialCalculationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPolynomialCalculationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPolynomialCalculationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPolynomialCalculationResponseValidationError{}

// Validate checks the field values on GetParallelPolynomialCalculationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *GetParallelPolynomialCalculationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// GetParallelPolynomialCalculationRequest with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// GetParallelPolynomialCalculationRequestMultiError, or nil if none found.
func (m *GetParallelPolynomialCalculationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetParallelPolynomialCalculationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Matrix

	if len(errors) > 0 {
		return GetParallelPolynomialCalculationRequestMultiError(errors)
	}

	return nil
}

// GetParallelPolynomialCalculationRequestMultiError is an error wrapping
// multiple validation errors returned by
// GetParallelPolynomialCalculationRequest.ValidateAll() if the designated
// constraints aren't met.
type GetParallelPolynomialCalculationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetParallelPolynomialCalculationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetParallelPolynomialCalculationRequestMultiError) AllErrors() []error { return m }

// GetParallelPolynomialCalculationRequestValidationError is the validation
// error returned by GetParallelPolynomialCalculationRequest.Validate if the
// designated constraints aren't met.
type GetParallelPolynomialCalculationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetParallelPolynomialCalculationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetParallelPolynomialCalculationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetParallelPolynomialCalculationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetParallelPolynomialCalculationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetParallelPolynomialCalculationRequestValidationError) ErrorName() string {
	return "GetParallelPolynomialCalculationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetParallelPolynomialCalculationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetParallelPolynomialCalculationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetParallelPolynomialCalculationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetParallelPolynomialCalculationRequestValidationError{}

// Validate checks the field values on GetParallelPolynomialCalculationResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *GetParallelPolynomialCalculationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// GetParallelPolynomialCalculationResponse with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// GetParallelPolynomialCalculationResponseMultiError, or nil if none found.
func (m *GetParallelPolynomialCalculationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetParallelPolynomialCalculationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Operation

	// no validation rules for ResultMatrix

	// no validation rules for Time

	// no validation rules for OperationType

	if len(errors) > 0 {
		return GetParallelPolynomialCalculationResponseMultiError(errors)
	}

	return nil
}

// GetParallelPolynomialCalculationResponseMultiError is an error wrapping
// multiple validation errors returned by
// GetParallelPolynomialCalculationResponse.ValidateAll() if the designated
// constraints aren't met.
type GetParallelPolynomialCalculationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetParallelPolynomialCalculationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetParallelPolynomialCalculationResponseMultiError) AllErrors() []error { return m }

// GetParallelPolynomialCalculationResponseValidationError is the validation
// error returned by GetParallelPolynomialCalculationResponse.Validate if the
// designated constraints aren't met.
type GetParallelPolynomialCalculationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetParallelPolynomialCalculationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetParallelPolynomialCalculationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetParallelPolynomialCalculationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetParallelPolynomialCalculationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetParallelPolynomialCalculationResponseValidationError) ErrorName() string {
	return "GetParallelPolynomialCalculationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetParallelPolynomialCalculationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetParallelPolynomialCalculationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetParallelPolynomialCalculationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetParallelPolynomialCalculationResponseValidationError{}

// Validate checks the field values on GetParallelLinearFormCalculationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *GetParallelLinearFormCalculationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// GetParallelLinearFormCalculationRequest with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// GetParallelLinearFormCalculationRequestMultiError, or nil if none found.
func (m *GetParallelLinearFormCalculationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetParallelLinearFormCalculationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Matrix

	if len(errors) > 0 {
		return GetParallelLinearFormCalculationRequestMultiError(errors)
	}

	return nil
}

// GetParallelLinearFormCalculationRequestMultiError is an error wrapping
// multiple validation errors returned by
// GetParallelLinearFormCalculationRequest.ValidateAll() if the designated
// constraints aren't met.
type GetParallelLinearFormCalculationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetParallelLinearFormCalculationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetParallelLinearFormCalculationRequestMultiError) AllErrors() []error { return m }

// GetParallelLinearFormCalculationRequestValidationError is the validation
// error returned by GetParallelLinearFormCalculationRequest.Validate if the
// designated constraints aren't met.
type GetParallelLinearFormCalculationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetParallelLinearFormCalculationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetParallelLinearFormCalculationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetParallelLinearFormCalculationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetParallelLinearFormCalculationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetParallelLinearFormCalculationRequestValidationError) ErrorName() string {
	return "GetParallelLinearFormCalculationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetParallelLinearFormCalculationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetParallelLinearFormCalculationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetParallelLinearFormCalculationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetParallelLinearFormCalculationRequestValidationError{}

// Validate checks the field values on GetParallelLinearFormCalculationResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *GetParallelLinearFormCalculationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// GetParallelLinearFormCalculationResponse with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// GetParallelLinearFormCalculationResponseMultiError, or nil if none found.
func (m *GetParallelLinearFormCalculationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetParallelLinearFormCalculationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Operation

	// no validation rules for ResultMatrix

	// no validation rules for Time

	// no validation rules for OperationType

	if len(errors) > 0 {
		return GetParallelLinearFormCalculationResponseMultiError(errors)
	}

	return nil
}

// GetParallelLinearFormCalculationResponseMultiError is an error wrapping
// multiple validation errors returned by
// GetParallelLinearFormCalculationResponse.ValidateAll() if the designated
// constraints aren't met.
type GetParallelLinearFormCalculationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetParallelLinearFormCalculationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetParallelLinearFormCalculationResponseMultiError) AllErrors() []error { return m }

// GetParallelLinearFormCalculationResponseValidationError is the validation
// error returned by GetParallelLinearFormCalculationResponse.Validate if the
// designated constraints aren't met.
type GetParallelLinearFormCalculationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetParallelLinearFormCalculationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetParallelLinearFormCalculationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetParallelLinearFormCalculationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetParallelLinearFormCalculationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetParallelLinearFormCalculationResponseValidationError) ErrorName() string {
	return "GetParallelLinearFormCalculationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetParallelLinearFormCalculationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetParallelLinearFormCalculationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetParallelLinearFormCalculationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetParallelLinearFormCalculationResponseValidationError{}

// Validate checks the field values on GetLinearFormCalculationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetLinearFormCalculationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLinearFormCalculationRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetLinearFormCalculationRequestMultiError, or nil if none found.
func (m *GetLinearFormCalculationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLinearFormCalculationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Matrix

	if len(errors) > 0 {
		return GetLinearFormCalculationRequestMultiError(errors)
	}

	return nil
}

// GetLinearFormCalculationRequestMultiError is an error wrapping multiple
// validation errors returned by GetLinearFormCalculationRequest.ValidateAll()
// if the designated constraints aren't met.
type GetLinearFormCalculationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLinearFormCalculationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLinearFormCalculationRequestMultiError) AllErrors() []error { return m }

// GetLinearFormCalculationRequestValidationError is the validation error
// returned by GetLinearFormCalculationRequest.Validate if the designated
// constraints aren't met.
type GetLinearFormCalculationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLinearFormCalculationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLinearFormCalculationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLinearFormCalculationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLinearFormCalculationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLinearFormCalculationRequestValidationError) ErrorName() string {
	return "GetLinearFormCalculationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetLinearFormCalculationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLinearFormCalculationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLinearFormCalculationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLinearFormCalculationRequestValidationError{}

// Validate checks the field values on GetLinearFormCalculationResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetLinearFormCalculationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLinearFormCalculationResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetLinearFormCalculationResponseMultiError, or nil if none found.
func (m *GetLinearFormCalculationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLinearFormCalculationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Operation

	// no validation rules for ResultMatrix

	// no validation rules for Time

	// no validation rules for OperationType

	if len(errors) > 0 {
		return GetLinearFormCalculationResponseMultiError(errors)
	}

	return nil
}

// GetLinearFormCalculationResponseMultiError is an error wrapping multiple
// validation errors returned by
// GetLinearFormCalculationResponse.ValidateAll() if the designated
// constraints aren't met.
type GetLinearFormCalculationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLinearFormCalculationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLinearFormCalculationResponseMultiError) AllErrors() []error { return m }

// GetLinearFormCalculationResponseValidationError is the validation error
// returned by GetLinearFormCalculationResponse.Validate if the designated
// constraints aren't met.
type GetLinearFormCalculationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLinearFormCalculationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLinearFormCalculationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLinearFormCalculationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLinearFormCalculationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLinearFormCalculationResponseValidationError) ErrorName() string {
	return "GetLinearFormCalculationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetLinearFormCalculationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLinearFormCalculationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLinearFormCalculationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLinearFormCalculationResponseValidationError{}

// Validate checks the field values on GetStatusRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetStatusRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStatusRequestMultiError, or nil if none found.
func (m *GetStatusRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStatusRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WorkerId

	if len(errors) > 0 {
		return GetStatusRequestMultiError(errors)
	}

	return nil
}

// GetStatusRequestMultiError is an error wrapping multiple validation errors
// returned by GetStatusRequest.ValidateAll() if the designated constraints
// aren't met.
type GetStatusRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStatusRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStatusRequestMultiError) AllErrors() []error { return m }

// GetStatusRequestValidationError is the validation error returned by
// GetStatusRequest.Validate if the designated constraints aren't met.
type GetStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStatusRequestValidationError) ErrorName() string { return "GetStatusRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStatusRequestValidationError{}

// Validate checks the field values on GetStatusResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetStatusResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetStatusResponseMultiError, or nil if none found.
func (m *GetStatusResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStatusResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WorkerId

	// no validation rules for Status

	if len(errors) > 0 {
		return GetStatusResponseMultiError(errors)
	}

	return nil
}

// GetStatusResponseMultiError is an error wrapping multiple validation errors
// returned by GetStatusResponse.ValidateAll() if the designated constraints
// aren't met.
type GetStatusResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStatusResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStatusResponseMultiError) AllErrors() []error { return m }

// GetStatusResponseValidationError is the validation error returned by
// GetStatusResponse.Validate if the designated constraints aren't met.
type GetStatusResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStatusResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStatusResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStatusResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStatusResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStatusResponseValidationError) ErrorName() string {
	return "GetStatusResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetStatusResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStatusResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStatusResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStatusResponseValidationError{}
