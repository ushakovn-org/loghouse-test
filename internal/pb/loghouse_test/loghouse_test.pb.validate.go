// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/loghouse_test/loghouse_test.proto

package loghouse_test

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

// Validate checks the field values on GetDummyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetDummyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDummyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDummyRequestMultiError, or nil if none found.
func (m *GetDummyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDummyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetDummyRequestMultiError(errors)
	}

	return nil
}

// GetDummyRequestMultiError is an error wrapping multiple validation errors
// returned by GetDummyRequest.ValidateAll() if the designated constraints
// aren't met.
type GetDummyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDummyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDummyRequestMultiError) AllErrors() []error { return m }

// GetDummyRequestValidationError is the validation error returned by
// GetDummyRequest.Validate if the designated constraints aren't met.
type GetDummyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDummyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDummyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDummyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDummyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDummyRequestValidationError) ErrorName() string { return "GetDummyRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetDummyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDummyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDummyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDummyRequestValidationError{}

// Validate checks the field values on GetDummyResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetDummyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDummyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDummyResponseMultiError, or nil if none found.
func (m *GetDummyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDummyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetDummy() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetDummyResponseValidationError{
						field:  fmt.Sprintf("Dummy[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetDummyResponseValidationError{
						field:  fmt.Sprintf("Dummy[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetDummyResponseValidationError{
					field:  fmt.Sprintf("Dummy[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetDummyResponseMultiError(errors)
	}

	return nil
}

// GetDummyResponseMultiError is an error wrapping multiple validation errors
// returned by GetDummyResponse.ValidateAll() if the designated constraints
// aren't met.
type GetDummyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDummyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDummyResponseMultiError) AllErrors() []error { return m }

// GetDummyResponseValidationError is the validation error returned by
// GetDummyResponse.Validate if the designated constraints aren't met.
type GetDummyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDummyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDummyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDummyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDummyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDummyResponseValidationError) ErrorName() string { return "GetDummyResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetDummyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDummyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDummyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDummyResponseValidationError{}

// Validate checks the field values on Dummy with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Dummy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Dummy with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DummyMultiError, or nil if none found.
func (m *Dummy) ValidateAll() error {
	return m.validate(true)
}

func (m *Dummy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Field

	if len(errors) > 0 {
		return DummyMultiError(errors)
	}

	return nil
}

// DummyMultiError is an error wrapping multiple validation errors returned by
// Dummy.ValidateAll() if the designated constraints aren't met.
type DummyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DummyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DummyMultiError) AllErrors() []error { return m }

// DummyValidationError is the validation error returned by Dummy.Validate if
// the designated constraints aren't met.
type DummyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DummyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DummyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DummyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DummyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DummyValidationError) ErrorName() string { return "DummyValidationError" }

// Error satisfies the builtin error interface
func (e DummyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDummy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DummyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DummyValidationError{}
