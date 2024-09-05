package protomagic

import (
	"testing"

	dummyv1 "github.com/gochore/protomagic/testdata/gen/dummy/v1"

	"github.com/stretchr/testify/assert"
)

func TestEnum_Wrap(t *testing.T) {
	v := dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO
	assert.Equal(t, v, WrapEnum(v).Unwrap())
}

func TestEnum_ShortName(t *testing.T) {
	t.Run("TestEnumType", func(t *testing.T) {
		cases := []struct {
			enum dummyv1.TestEnumType
			want string
		}{
			{
				enum: dummyv1.TestEnumType_TEST_ENUM_TYPE_UNSPECIFIED,
				want: "unspecified",
			},
			{
				enum: dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO,
				want: "foo",
			},
			{
				enum: dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR,
				want: "bar",
			},
			{
				enum: dummyv1.TestEnumType_TEST_ENUM_TYPE_HELLO_WORLD,
				want: "hello_world",
			},
			{
				enum: dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO2,
				want: "foo2",
			},
			{
				enum: dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO_3,
				want: "foo_3",
			},
			{
				enum: dummyv1.TestEnumType_TEST_ENUM_TYPE_HELLO,
				want: "hello",
			},
		}
		for _, tt := range cases {
			t.Run(tt.enum.String(), func(t *testing.T) {
				assert.Equal(t, tt.want, WrapEnum(tt.enum).ShortName())
			})
		}
	})
	t.Run("TestEnum2Type", func(t *testing.T) {
		cases := []struct {
			enum dummyv1.TestEnum2Type
			want string
		}{
			{
				enum: dummyv1.TestEnum2Type_TEST_ENUM_2_TYPE_UNSPECIFIED,
				want: "unspecified",
			},
			{
				enum: dummyv1.TestEnum2Type_TEST_ENUM_2_TYPE_FOO,
				want: "foo",
			},
			{
				enum: dummyv1.TestEnum2Type_TEST_ENUM_2_TYPE_BAR,
				want: "bar",
			},
		}
		for _, tt := range cases {
			t.Run(tt.enum.String(), func(t *testing.T) {
				assert.Equal(t, tt.want, WrapEnum(tt.enum).ShortName())
			})
		}
	})
}

func TestEnum_AllValues(t *testing.T) {
	t.Run("TestEnumType", func(t *testing.T) {
		values := WrapEnum(dummyv1.TestEnumType_TEST_ENUM_TYPE_UNSPECIFIED).AllValues()
		assert.Equal(t, []dummyv1.TestEnumType{
			dummyv1.TestEnumType_TEST_ENUM_TYPE_UNSPECIFIED,
			dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO,
			dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR,
			dummyv1.TestEnumType_TEST_ENUM_TYPE_HELLO_WORLD,
			dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO2,
			dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO_3,
			dummyv1.TestEnumType_TEST_ENUM_TYPE_HELLO,
		}, values)
	})
	t.Run("TestEnum2Type", func(t *testing.T) {
		values := WrapEnum(dummyv1.TestEnum2Type_TEST_ENUM_2_TYPE_UNSPECIFIED).AllValues()
		assert.Equal(t, []dummyv1.TestEnum2Type{
			dummyv1.TestEnum2Type_TEST_ENUM_2_TYPE_UNSPECIFIED,
			dummyv1.TestEnum2Type_TEST_ENUM_2_TYPE_FOO,
			dummyv1.TestEnum2Type_TEST_ENUM_2_TYPE_BAR,
		}, values)
	})
}
