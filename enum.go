package protomagic

import (
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// Enum is a wrapper around an enum value.
// It provides some useful methods to work with enums.
// There are two ways to create an Enum[T]:
//  1. WrapEnum[T](T) wraps an existing enum value.
//  2. NewEnum[T]() creates a new Enum[T] with the zero value of T.
type Enum[T protoreflect.Enum] struct {
	inner T
}

// WrapEnum wraps an enum value with Enum[T].
func WrapEnum[T protoreflect.Enum](inner T) *Enum[T] {
	return &Enum[T]{
		inner: inner,
	}
}

// NewEnum creates a new Enum[T] with the zero value of T.
func NewEnum[T protoreflect.Enum]() *Enum[T] {
	var zero T
	return WrapEnum[T](zero)
}

// Unwrap returns the inner enum value.
func (e *Enum[T]) Unwrap() T {
	return e.inner
}

// ShortName returns the short name in the lower case of the enum value.
// For example, "TEST_ENUM_TYPE_FOO" -> "foo".
func (e *Enum[T]) ShortName() string {
	prefix := string(e.inner.Type().Descriptor().Name()) // like "TestEnumType"
	prefix = toSnake(prefix) + "_"                       // like "TEST_ENUM_TYPE_"

	ret := protoimpl.X.EnumStringOf(e.inner.Descriptor(), e.inner.Number()) // like "TEST_ENUM_TYPE_FOO", follow the implementation of String() method
	ret = strings.TrimPrefix(ret, prefix)                                   // like "FOO"
	ret = strings.ToLower(ret)                                              // like "foo"
	return ret
}

// AllValues returns all declared values of the enum type.
// It includes the zero (_UNSPECIFIED) value.
func (e *Enum[T]) AllValues() []T {
	typ := e.inner.Type()
	values := typ.Descriptor().Values()
	ret := make([]T, 0, values.Len())
	for i := range values.Len() {
		ret = append(ret, typ.New(values.Get(i).Number()).(T))
	}
	return ret
}

// SpecifiedValues returns all specified values of the enum type.
// It excludes the zero (_UNSPECIFIED) value.
func (e *Enum[T]) SpecifiedValues() []T {
	typ := e.inner.Type()
	values := typ.Descriptor().Values()
	ret := make([]T, 0, values.Len()-1)
	for i := range values.Len() {
		if values.Get(i).Number() != 0 {
			ret = append(ret, typ.New(values.Get(i).Number()).(T))
		}
	}
	return ret
}

// toSnake converts a string to SNACK_CASE.
// It assumes the input string is in CamelCase and contains only A-Z, a-z and 0-9.
// It treats numbers as lower case letters.
func toSnake(s string) string {
	ret := &strings.Builder{}
	toUpper := func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r - 'a' + 'A'
		}
		return r
	}
	for i, v := range s {
		isCap := v >= 'A' && v <= 'Z'
		if !isCap {
			ret.WriteRune(toUpper(v))
			continue
		}
		if i > 0 {
			ret.WriteRune('_')
		}
		ret.WriteRune(v)
	}
	return ret.String()
}
