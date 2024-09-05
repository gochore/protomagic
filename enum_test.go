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

func TestNewEnum(t *testing.T) {
	assert.Equal(t, dummyv1.TestEnumType_TEST_ENUM_TYPE_UNSPECIFIED, NewEnum[dummyv1.TestEnumType]().Unwrap())
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

	t.Run("with numbers", func(t *testing.T) {
		cases := []struct {
			enum int32
			want string
		}{
			{enum: 0, want: "unspecified"},
			{enum: 1, want: "foo"},
			{enum: 2, want: "bar"},
		}
		t.Run("TestEnumType2", func(t *testing.T) {
			for _, tt := range cases {
				enum := dummyv1.TestEnumType2(tt.enum)
				t.Run(enum.String(), func(t *testing.T) {
					assert.Equal(t, tt.want, WrapEnum(enum).ShortName())
				})
			}
		})
		t.Run("TestEnum3Type", func(t *testing.T) {
			for _, tt := range cases {
				enum := dummyv1.TestEnum3Type(tt.enum)
				t.Run(enum.String(), func(t *testing.T) {
					assert.Equal(t, tt.want, WrapEnum(enum).ShortName())
				})
			}
		})
		t.Run("TestEnum4ThType", func(t *testing.T) {
			for _, tt := range cases {
				enum := dummyv1.TestEnum4ThType(tt.enum)
				t.Run(enum.String(), func(t *testing.T) {
					assert.Equal(t, tt.want, WrapEnum(enum).ShortName())
				})
			}
		})
		t.Run("TestEnum05Type", func(t *testing.T) {
			for _, tt := range cases {
				enum := dummyv1.TestEnum05ThType(tt.enum)
				t.Run(enum.String(), func(t *testing.T) {
					assert.Equal(t, tt.want, WrapEnum(enum).ShortName())
				})
			}
		})
	})
}

func TestEnum_AllValues(t *testing.T) {
	t.Run("TestEnumType", func(t *testing.T) {
		values := NewEnum[dummyv1.TestEnumType]().AllValues()
		assert.Len(t, values, 7)
	})
	t.Run("TestEnum2Type", func(t *testing.T) {
		values := NewEnum[dummyv1.TestEnumType2]().AllValues()
		assert.Len(t, values, 3)
	})
}

func TestEnum_SpecifiedValues(t *testing.T) {
	t.Run("TestEnumType", func(t *testing.T) {
		values := NewEnum[dummyv1.TestEnumType]().SpecifiedValues()
		assert.Len(t, values, 6)
	})
	t.Run("TestEnum2Type", func(t *testing.T) {
		values := NewEnum[dummyv1.TestEnumType2]().SpecifiedValues()
		assert.Len(t, values, 2)
	})
}
