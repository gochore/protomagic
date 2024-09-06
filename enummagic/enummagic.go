package enummagic

import (
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// ShortName returns the short name in the lower case of the enum value.
// For example, "TEST_ENUM_TYPE_FOO" -> "foo".
func ShortName[T protoreflect.Enum](enum T) string {
	prefix := string(enum.Type().Descriptor().Name()) // like "TestEnumType"
	prefix = toSnake(prefix) + "_"                    // like "TEST_ENUM_TYPE_"

	ret := protoimpl.X.EnumStringOf(enum.Descriptor(), enum.Number()) // like "TEST_ENUM_TYPE_FOO", follow the implementation of String() method
	ret = strings.TrimPrefix(ret, prefix)                             // like "FOO"
	ret = strings.ToLower(ret)                                        // like "foo"
	return ret
}

// AllDefined returns all defined values of the enum type.
// It includes the zero (_UNSPECIFIED) value.
func AllDefined[T protoreflect.Enum](_ ...T) []T {
	var enum T
	typ := enum.Type()
	values := typ.Descriptor().Values()
	ret := make([]T, 0, values.Len())
	for i := range values.Len() {
		ret = append(ret, typ.New(values.Get(i).Number()).(T))
	}
	return ret
}

// AllSpecified returns all specified values of the enum type.
// It excludes the zero (_UNSPECIFIED) value.
func AllSpecified[T protoreflect.Enum](_ ...T) []T {
	var enum T
	typ := enum.Type()
	values := typ.Descriptor().Values()
	ret := make([]T, 0, values.Len()-1)
	for i := range values.Len() {
		if values.Get(i).Number() != 0 {
			ret = append(ret, typ.New(values.Get(i).Number()).(T))
		}
	}
	return ret
}

// IsDefined returns true if the enum value is defined in the enum type.
func IsDefined[T protoreflect.Enum](enum T) bool {
	return enum.Descriptor().Values().ByNumber(enum.Number()) != nil
}

// IsSpecified returns true if the enum value is defined and not the zero value.
func IsSpecified[T protoreflect.Enum](enum T) bool {
	return IsDefined(enum) && enum.Number() != 0
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
