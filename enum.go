package protomagic

import (
	"strings"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type Enum[T protoreflect.Enum] struct {
	inner T
}

func WrapEnum[T protoreflect.Enum](inner T) *Enum[T] {
	return &Enum[T]{inner: inner}
}

func (e *Enum[T]) Unwrap() T {
	return e.inner
}

// ShortName returns the short name in the lower case of the enum value.
// For example, "TEST_ENUM_TYPE_FOO" -> "foo".
func (e *Enum[T]) ShortName() string {
	prefix := string(e.inner.Type().Descriptor().Name()) // like "TestEnumType"
	prefix = strcase.ToScreamingSnake(prefix) + "_"      // like "TEST_ENUM_TYPE_"

	ret := protoimpl.X.EnumStringOf(e.inner.Descriptor(), e.inner.Number()) // like "TEST_ENUM_TYPE_FOO", follow the implementation of String() method
	ret = strings.TrimPrefix(ret, prefix)                                   // like "FOO"
	ret = strings.ToLower(ret)                                              // like "foo"
	return ret
}
