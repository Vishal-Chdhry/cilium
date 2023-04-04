// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/common/dynamic_forward_proxy/v3/dns_cache.proto

package dynamic_forward_proxyv3

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

	v3 "github.com/cilium/proxy/go/envoy/config/cluster/v3"
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

	_ = v3.Cluster_DnsLookupFamily(0)
)

// Validate checks the field values on DnsCacheCircuitBreakers with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DnsCacheCircuitBreakers) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DnsCacheCircuitBreakers with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DnsCacheCircuitBreakersMultiError, or nil if none found.
func (m *DnsCacheCircuitBreakers) ValidateAll() error {
	return m.validate(true)
}

func (m *DnsCacheCircuitBreakers) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMaxPendingRequests()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DnsCacheCircuitBreakersValidationError{
					field:  "MaxPendingRequests",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DnsCacheCircuitBreakersValidationError{
					field:  "MaxPendingRequests",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMaxPendingRequests()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsCacheCircuitBreakersValidationError{
				field:  "MaxPendingRequests",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DnsCacheCircuitBreakersMultiError(errors)
	}
	return nil
}

// DnsCacheCircuitBreakersMultiError is an error wrapping multiple validation
// errors returned by DnsCacheCircuitBreakers.ValidateAll() if the designated
// constraints aren't met.
type DnsCacheCircuitBreakersMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DnsCacheCircuitBreakersMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DnsCacheCircuitBreakersMultiError) AllErrors() []error { return m }

// DnsCacheCircuitBreakersValidationError is the validation error returned by
// DnsCacheCircuitBreakers.Validate if the designated constraints aren't met.
type DnsCacheCircuitBreakersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsCacheCircuitBreakersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsCacheCircuitBreakersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsCacheCircuitBreakersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsCacheCircuitBreakersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsCacheCircuitBreakersValidationError) ErrorName() string {
	return "DnsCacheCircuitBreakersValidationError"
}

// Error satisfies the builtin error interface
func (e DnsCacheCircuitBreakersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsCacheCircuitBreakers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsCacheCircuitBreakersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsCacheCircuitBreakersValidationError{}

// Validate checks the field values on DnsCacheConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DnsCacheConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DnsCacheConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DnsCacheConfigMultiError,
// or nil if none found.
func (m *DnsCacheConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *DnsCacheConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := DnsCacheConfigValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := v3.Cluster_DnsLookupFamily_name[int32(m.GetDnsLookupFamily())]; !ok {
		err := DnsCacheConfigValidationError{
			field:  "DnsLookupFamily",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if d := m.GetDnsRefreshRate(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = DnsCacheConfigValidationError{
				field:  "DnsRefreshRate",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gte := time.Duration(0*time.Second + 1000000*time.Nanosecond)

			if dur < gte {
				err := DnsCacheConfigValidationError{
					field:  "DnsRefreshRate",
					reason: "value must be greater than or equal to 1ms",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if d := m.GetDnsMinRefreshRate(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = DnsCacheConfigValidationError{
				field:  "DnsMinRefreshRate",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gte := time.Duration(5*time.Second + 0*time.Nanosecond)

			if dur < gte {
				err := DnsCacheConfigValidationError{
					field:  "DnsMinRefreshRate",
					reason: "value must be greater than or equal to 5s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if d := m.GetHostTtl(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = DnsCacheConfigValidationError{
				field:  "HostTtl",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := DnsCacheConfigValidationError{
					field:  "HostTtl",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if wrapper := m.GetMaxHosts(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			err := DnsCacheConfigValidationError{
				field:  "MaxHosts",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if all {
		switch v := interface{}(m.GetDnsFailureRefreshRate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DnsCacheConfigValidationError{
					field:  "DnsFailureRefreshRate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DnsCacheConfigValidationError{
					field:  "DnsFailureRefreshRate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDnsFailureRefreshRate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsCacheConfigValidationError{
				field:  "DnsFailureRefreshRate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDnsCacheCircuitBreaker()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DnsCacheConfigValidationError{
					field:  "DnsCacheCircuitBreaker",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DnsCacheConfigValidationError{
					field:  "DnsCacheCircuitBreaker",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDnsCacheCircuitBreaker()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsCacheConfigValidationError{
				field:  "DnsCacheCircuitBreaker",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UseTcpForDnsLookups

	if all {
		switch v := interface{}(m.GetDnsResolutionConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DnsCacheConfigValidationError{
					field:  "DnsResolutionConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DnsCacheConfigValidationError{
					field:  "DnsResolutionConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDnsResolutionConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsCacheConfigValidationError{
				field:  "DnsResolutionConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTypedDnsResolverConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DnsCacheConfigValidationError{
					field:  "TypedDnsResolverConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DnsCacheConfigValidationError{
					field:  "TypedDnsResolverConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTypedDnsResolverConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsCacheConfigValidationError{
				field:  "TypedDnsResolverConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetPreresolveHostnames() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DnsCacheConfigValidationError{
						field:  fmt.Sprintf("PreresolveHostnames[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DnsCacheConfigValidationError{
						field:  fmt.Sprintf("PreresolveHostnames[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsCacheConfigValidationError{
					field:  fmt.Sprintf("PreresolveHostnames[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if d := m.GetDnsQueryTimeout(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = DnsCacheConfigValidationError{
				field:  "DnsQueryTimeout",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := DnsCacheConfigValidationError{
					field:  "DnsQueryTimeout",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if all {
		switch v := interface{}(m.GetKeyValueConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DnsCacheConfigValidationError{
					field:  "KeyValueConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DnsCacheConfigValidationError{
					field:  "KeyValueConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetKeyValueConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsCacheConfigValidationError{
				field:  "KeyValueConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DnsCacheConfigMultiError(errors)
	}
	return nil
}

// DnsCacheConfigMultiError is an error wrapping multiple validation errors
// returned by DnsCacheConfig.ValidateAll() if the designated constraints
// aren't met.
type DnsCacheConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DnsCacheConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DnsCacheConfigMultiError) AllErrors() []error { return m }

// DnsCacheConfigValidationError is the validation error returned by
// DnsCacheConfig.Validate if the designated constraints aren't met.
type DnsCacheConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsCacheConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsCacheConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsCacheConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsCacheConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsCacheConfigValidationError) ErrorName() string { return "DnsCacheConfigValidationError" }

// Error satisfies the builtin error interface
func (e DnsCacheConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsCacheConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsCacheConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsCacheConfigValidationError{}
