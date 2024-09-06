package enummagic

import (
	"testing"

	dummyv1 "github.com/gochore/protomagic/testdata/gen/dummy/v1"

	"github.com/stretchr/testify/assert"
)

func TestShortName(t *testing.T) {
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
				assert.Equal(t, tt.want, ShortName(tt.enum))
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
					assert.Equal(t, tt.want, ShortName(enum))
				})
			}
		})
		t.Run("TestEnum3Type", func(t *testing.T) {
			for _, tt := range cases {
				enum := dummyv1.TestEnum3Type(tt.enum)
				t.Run(enum.String(), func(t *testing.T) {
					assert.Equal(t, tt.want, ShortName(enum))
				})
			}
		})
		t.Run("TestEnum4ThType", func(t *testing.T) {
			for _, tt := range cases {
				enum := dummyv1.TestEnum4ThType(tt.enum)
				t.Run(enum.String(), func(t *testing.T) {
					assert.Equal(t, tt.want, ShortName(enum))
				})
			}
		})
		t.Run("TestEnum05Type", func(t *testing.T) {
			for _, tt := range cases {
				enum := dummyv1.TestEnum05ThType(tt.enum)
				t.Run(enum.String(), func(t *testing.T) {
					assert.Equal(t, tt.want, ShortName(enum))
				})
			}
		})
	})
}

func TestAllDefined(t *testing.T) {
	t.Run("TestEnumType", func(t *testing.T) {
		values := AllDefined[dummyv1.TestEnumType]()
		assert.Len(t, values, 7)
	})
	t.Run("TestEnum2Type", func(t *testing.T) {
		values := AllDefined[dummyv1.TestEnumType2]()
		assert.Len(t, values, 3)
	})
	t.Run("TestUnorderedEnumType", func(t *testing.T) {
		values := AllDefined[dummyv1.TestUnorderedEnumType]()
		assert.Equal(t, []dummyv1.TestUnorderedEnumType{
			dummyv1.TestUnorderedEnumType_TEST_UNORDERED_ENUM_TYPE_UNSPECIFIED,
			dummyv1.TestUnorderedEnumType_TEST_UNORDERED_ENUM_TYPE_BAR,
			dummyv1.TestUnorderedEnumType_TEST_UNORDERED_ENUM_TYPE_FOO,
		}, values)
	})
	t.Run("TestEnumType_TEST_ENUM_TYPE_FOO", func(t *testing.T) {
		values := AllDefined(dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO)
		assert.Len(t, values, 7)
	})
}

func TestAllSpecified(t *testing.T) {
	t.Run("TestEnumType", func(t *testing.T) {
		values := AllSpecified[dummyv1.TestEnumType]()
		assert.Len(t, values, 6)
	})
	t.Run("TestEnum2Type", func(t *testing.T) {
		values := AllSpecified[dummyv1.TestEnumType2]()
		assert.Len(t, values, 2)
	})
	t.Run("TestUnorderedEnumType", func(t *testing.T) {
		values := AllSpecified[dummyv1.TestUnorderedEnumType]()
		assert.Equal(t, []dummyv1.TestUnorderedEnumType{
			dummyv1.TestUnorderedEnumType_TEST_UNORDERED_ENUM_TYPE_BAR,
			dummyv1.TestUnorderedEnumType_TEST_UNORDERED_ENUM_TYPE_FOO,
		}, values)
	})
	t.Run("TestEnumType_TEST_ENUM_TYPE_FOO", func(t *testing.T) {
		values := AllSpecified(dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO)
		assert.Len(t, values, 6)
	})
}

func TestIsDefined(t *testing.T) {
	assert.True(t, IsDefined(dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO))
	assert.True(t, IsDefined(dummyv1.TestEnumType_TEST_ENUM_TYPE_UNSPECIFIED))
	assert.False(t, IsDefined(dummyv1.TestEnumType(1000)))
}

func TestIsSpecified(t *testing.T) {
	assert.True(t, IsSpecified(dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO))
	assert.False(t, IsSpecified(dummyv1.TestEnumType_TEST_ENUM_TYPE_UNSPECIFIED))
	assert.False(t, IsSpecified(dummyv1.TestEnumType(1000)))
}
