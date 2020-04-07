// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/interservice/local_user/users.proto

package local_user

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on Email with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Email) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetEmail()) < 1 {
		return EmailValidationError{
			field:  "Email",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// EmailValidationError is the validation error returned by Email.Validate if
// the designated constraints aren't met.
type EmailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmailValidationError) ErrorName() string { return "EmailValidationError" }

// Error satisfies the builtin error interface
func (e EmailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmailValidationError{}

// Validate checks the field values on GetUsersReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetUsersReq) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetUsersReqValidationError is the validation error returned by
// GetUsersReq.Validate if the designated constraints aren't met.
type GetUsersReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsersReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsersReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsersReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsersReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsersReqValidationError) ErrorName() string { return "GetUsersReqValidationError" }

// Error satisfies the builtin error interface
func (e GetUsersReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsersReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsersReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsersReqValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *User) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Email

	return nil
}

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on CreateUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreateUserReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	if !_CreateUserReq_Email_Pattern.MatchString(m.GetEmail()) {
		return CreateUserReqValidationError{
			field:  "Email",
			reason: "value does not match regex pattern \"^[[:alnum:]_.+@-]+$\"",
		}
	}

	// no validation rules for Password

	return nil
}

// CreateUserReqValidationError is the validation error returned by
// CreateUserReq.Validate if the designated constraints aren't met.
type CreateUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserReqValidationError) ErrorName() string { return "CreateUserReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserReqValidationError{}

var _CreateUserReq_Email_Pattern = regexp.MustCompile("^[[:alnum:]_.+@-]+$")

// Validate checks the field values on UpdateUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UpdateUserReq) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return UpdateUserReqValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
	}

	if !_UpdateUserReq_Email_Pattern.MatchString(m.GetEmail()) {
		return UpdateUserReqValidationError{
			field:  "Email",
			reason: "value does not match regex pattern \"^[[:alnum:]_.+@-]+$\"",
		}
	}

	// no validation rules for Name

	// no validation rules for Password

	return nil
}

// UpdateUserReqValidationError is the validation error returned by
// UpdateUserReq.Validate if the designated constraints aren't met.
type UpdateUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserReqValidationError) ErrorName() string { return "UpdateUserReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdateUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserReqValidationError{}

var _UpdateUserReq_Email_Pattern = regexp.MustCompile("^[[:alnum:]_.+@-]+$")

// Validate checks the field values on UpdateSelfReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UpdateSelfReq) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return UpdateSelfReqValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
	}

	if !_UpdateSelfReq_Email_Pattern.MatchString(m.GetEmail()) {
		return UpdateSelfReqValidationError{
			field:  "Email",
			reason: "value does not match regex pattern \"^[[:alnum:]_.+@-]+$\"",
		}
	}

	// no validation rules for Name

	// no validation rules for Password

	// no validation rules for PreviousPassword

	return nil
}

// UpdateSelfReqValidationError is the validation error returned by
// UpdateSelfReq.Validate if the designated constraints aren't met.
type UpdateSelfReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSelfReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSelfReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSelfReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSelfReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSelfReqValidationError) ErrorName() string { return "UpdateSelfReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdateSelfReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSelfReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSelfReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSelfReqValidationError{}

var _UpdateSelfReq_Email_Pattern = regexp.MustCompile("^[[:alnum:]_.+@-]+$")

// Validate checks the field values on DeleteUserResp with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DeleteUserResp) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteUserRespValidationError is the validation error returned by
// DeleteUserResp.Validate if the designated constraints aren't met.
type DeleteUserRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserRespValidationError) ErrorName() string { return "DeleteUserRespValidationError" }

// Error satisfies the builtin error interface
func (e DeleteUserRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserRespValidationError{}

// Validate checks the field values on Users with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Users) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Users

	return nil
}

// UsersValidationError is the validation error returned by Users.Validate if
// the designated constraints aren't met.
type UsersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UsersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UsersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UsersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UsersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UsersValidationError) ErrorName() string { return "UsersValidationError" }

// Error satisfies the builtin error interface
func (e UsersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUsers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UsersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UsersValidationError{}
