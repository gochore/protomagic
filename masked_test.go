package protomagic

import (
	"testing"

	dummyv1 "github.com/gochore/protomagic/testdata/gen/dummy/v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func assertMessageEqual(t *testing.T, expected, actual proto.Message) {
	assert.True(t, proto.Equal(expected, actual))
}

func TestNewMasked(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		m, err := NewMasked(&dummyv1.Dummy{}, "name", "value", "config_a")
		assert.NoError(t, err)
		require.NotNil(t, m)
		assertMessageEqual(t, &dummyv1.Dummy{}, m.msg)
		assertMessageEqual(t, &fieldmaskpb.FieldMask{Paths: []string{"name", "value", "config_a"}}, m.mask)
	})
	t.Run("invalid", func(t *testing.T) {
		m, err := NewMasked(&dummyv1.Dummy{}, "name", "value_not_exist", "config_a")
		assert.Error(t, err)
		assert.Nil(t, m)
	})
	t.Run("nested path", func(t *testing.T) {
		m, err := NewMasked(&dummyv1.Dummy{}, "name", "value", "config_a.name")
		assert.NoError(t, err)
		require.NotNil(t, m)
		assertMessageEqual(t, &dummyv1.Dummy{}, m.msg)
	})
}

func TestMustNewMasked(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		m := MustNewMasked(&dummyv1.Dummy{}, "name", "value", "config_a")
		require.NotNil(t, m)
		assertMessageEqual(t, &dummyv1.Dummy{}, m.msg)
		assertMessageEqual(t, &fieldmaskpb.FieldMask{Paths: []string{"name", "value", "config_a"}}, m.mask)
	})
	t.Run("invalid", func(t *testing.T) {
		m := MustNewMasked(&dummyv1.Dummy{}, "name", "value_not_exist", "config_a")
		require.NotNil(t, m)
		assertMessageEqual(t, &dummyv1.Dummy{}, m.msg)
		assertMessageEqual(t, &fieldmaskpb.FieldMask{Paths: []string{"name", "config_a"}}, m.mask)
	})
}

func TestNewMaskedFromFields(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		m, err := NewMaskedFromFields(&dummyv1.Dummy{}, &dummyv1.DummyConfigA{})
		assert.NoError(t, err)
		require.NotNil(t, m)
		assertMessageEqual(t, &dummyv1.Dummy{}, m.msg)
		assertMessageEqual(t, &fieldmaskpb.FieldMask{Paths: []string{"config_a"}}, m.mask)
	})
	t.Run("not found", func(t *testing.T) {
		m, err := NewMaskedFromFields(&dummyv1.Dummy{}, &dummyv1.DummyConfigA{}, &fieldmaskpb.FieldMask{})
		assert.ErrorContains(t, err, `field "FieldMask" not found in "Dummy"`)
		assert.Nil(t, m)
	})
	t.Run("ambiguous", func(t *testing.T) {
		m, err := NewMaskedFromFields(&dummyv1.Dummy{}, &dummyv1.DummyConfigA{}, &dummyv1.DummyConfigB{})
		assert.ErrorContains(t, err, `field "DummyConfigB" is ambiguous in "Dummy"`)
		assert.Nil(t, m)
	})
}

func TestMasked_Prune(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		dummy := &dummyv1.Dummy{
			Name:     "name",
			Value:    100,
			Values:   []string{"a", "b", "c"},
			TestType: dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO,
			ConfigA: &dummyv1.DummyConfigA{
				Name: "config_a",
			},
			ConfigB: &dummyv1.DummyConfigB{
				Name: "config_b",
			},
			ConfigB2: &dummyv1.DummyConfigB{
				Name: "config_b2",
			},
		}

		m, err := NewMasked(dummy, "name", "value", "config_a")
		require.NoError(t, err)
		require.NotNil(t, m)

		msg, err := m.Prune()
		require.NoError(t, err)
		require.NotNil(t, msg)
		assertMessageEqual(t, &dummyv1.Dummy{
			Name:  "name",
			Value: 100,
			ConfigA: &dummyv1.DummyConfigA{
				Name: "config_a",
			},
		}, msg)
	})

	t.Run("empty mask", func(t *testing.T) {
		dummy := &dummyv1.Dummy{
			Name:     "name",
			Value:    100,
			Values:   []string{"a", "b", "c"},
			TestType: dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO,
			ConfigA: &dummyv1.DummyConfigA{
				Name: "config_a",
			},
			ConfigB: &dummyv1.DummyConfigB{
				Name: "config_b",
			},
			ConfigB2: &dummyv1.DummyConfigB{
				Name: "config_b2",
			},
		}

		m, err := NewMasked(dummy)
		require.NoError(t, err)
		require.NotNil(t, m)

		msg, err := m.Prune()
		require.NoError(t, err)
		require.NotNil(t, msg)
		assertMessageEqual(t, dummy, msg)
	})

	t.Run("nested path", func(t *testing.T) {
		dummy := &dummyv1.Dummy{
			Name:     "name",
			Value:    100,
			Values:   []string{"a", "b", "c"},
			TestType: dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO,
			ConfigA: &dummyv1.DummyConfigA{
				Name: "config_a",
			},
			ConfigB: &dummyv1.DummyConfigB{
				Name: "config_b",
			},
			ConfigB2: &dummyv1.DummyConfigB{
				Name: "config_b2",
			},
		}

		m, err := NewMasked(dummy, "name", "value", "config_a.name")
		require.NoError(t, err)
		require.NotNil(t, m)

		msg, err := m.Prune()
		assert.ErrorContains(t, err, `nested field "config_a.name" is not supported`)
		assert.Nil(t, msg)
	})
}

func TestMasked_Patch(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		dummy := &dummyv1.Dummy{
			Name:     "name",
			Value:    100,
			Values:   []string{"a", "b", "c"},
			TestType: dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO,
			ConfigA: &dummyv1.DummyConfigA{
				Name: "config_a",
			},
			ConfigB: &dummyv1.DummyConfigB{
				Name: "config_b",
			},
			ConfigB2: &dummyv1.DummyConfigB{
				Name: "config_b2",
			},
		}

		m, err := NewMasked(dummy, "name", "value", "config_a")
		require.NoError(t, err)
		require.NotNil(t, m)

		patch := &dummyv1.Dummy{
			Name:     "name name",
			Value:    100_100,
			Values:   []string{"a", "b", "c", "a", "b", "c"},
			TestType: dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR,
			ConfigA: &dummyv1.DummyConfigA{
				Name: "config_a config_a",
			},
			ConfigB: &dummyv1.DummyConfigB{
				Name: "config_b config_b",
			},
			ConfigB2: &dummyv1.DummyConfigB{
				Name: "config_b2 config_b2",
			},
		}

		msg, err := m.Patch(patch)
		require.NoError(t, err)
		assertMessageEqual(t, &dummyv1.Dummy{
			Name:     "name name",
			Value:    100_100,
			Values:   []string{"a", "b", "c"},
			TestType: dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO,
			ConfigA: &dummyv1.DummyConfigA{
				Name: "config_a config_a",
			},
			ConfigB: &dummyv1.DummyConfigB{
				Name: "config_b",
			},
			ConfigB2: &dummyv1.DummyConfigB{
				Name: "config_b2",
			},
		}, msg)
	})

	t.Run("empty mask", func(t *testing.T) {
		dummy := &dummyv1.Dummy{
			Name:     "name",
			Value:    100,
			Values:   []string{"a", "b", "c"},
			TestType: dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO,
			ConfigA: &dummyv1.DummyConfigA{
				Name: "config_a",
			},
			ConfigB: &dummyv1.DummyConfigB{
				Name: "config_b",
			},
			ConfigB2: &dummyv1.DummyConfigB{
				Name: "config_b2",
			},
		}

		m, err := NewMasked(dummy)
		require.NoError(t, err)
		require.NotNil(t, m)

		patch := &dummyv1.Dummy{
			Name:     "name name",
			Value:    100_100,
			Values:   []string{"a", "b", "c", "a", "b", "c"},
			TestType: dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR,
			ConfigA: &dummyv1.DummyConfigA{
				Name: "config_a config_a",
			},
			ConfigB: &dummyv1.DummyConfigB{
				Name: "config_b config_b",
			},
			ConfigB2: &dummyv1.DummyConfigB{
				Name: "config_b2 config_b2",
			},
		}

		msg, err := m.Patch(patch)
		require.NoError(t, err)
		assertMessageEqual(t, dummy, msg)
	})

	t.Run("empty message", func(t *testing.T) {
		var dummy *dummyv1.Dummy

		m, err := NewMasked(dummy, "name", "value", "config_a")
		require.NoError(t, err)
		require.NotNil(t, m)

		patch := &dummyv1.Dummy{
			Name:     "name name",
			Value:    100_100,
			Values:   []string{"a", "b", "c", "a", "b", "c"},
			TestType: dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR,
			ConfigA: &dummyv1.DummyConfigA{
				Name: "config_a config_a",
			},
			ConfigB: &dummyv1.DummyConfigB{
				Name: "config_b config_b",
			},
			ConfigB2: &dummyv1.DummyConfigB{
				Name: "config_b2 config_b2",
			},
		}

		msg, err := m.Patch(patch)
		require.NoError(t, err)
		assert.Nil(t, msg)
	})

	t.Run("clear", func(t *testing.T) {
		dummy := &dummyv1.Dummy{
			Name:     "name",
			Value:    100,
			Values:   []string{"a", "b", "c"},
			TestType: dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO,
			ConfigA: &dummyv1.DummyConfigA{
				Name: "config_a",
			},
			ConfigB: &dummyv1.DummyConfigB{
				Name: "config_b",
			},
			ConfigB2: &dummyv1.DummyConfigB{
				Name: "config_b2",
			},
		}

		m, err := NewMasked(dummy, "name", "value", "config_a")
		require.NoError(t, err)
		require.NotNil(t, m)

		patch := &dummyv1.Dummy{
			Values:   []string{"a", "b", "c", "a", "b", "c"},
			TestType: dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR,
			ConfigB: &dummyv1.DummyConfigB{
				Name: "config_b config_b",
			},
			ConfigB2: &dummyv1.DummyConfigB{
				Name: "config_b2 config_b2",
			},
		}

		msg, err := m.Patch(patch)
		require.NoError(t, err)
		assertMessageEqual(t, &dummyv1.Dummy{
			Name:     "",
			Value:    0,
			Values:   []string{"a", "b", "c"},
			TestType: dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO,
			ConfigA:  nil,
			ConfigB: &dummyv1.DummyConfigB{
				Name: "config_b",
			},
			ConfigB2: &dummyv1.DummyConfigB{
				Name: "config_b2",
			},
		}, msg)
	})

	t.Run("nested path", func(t *testing.T) {
		dummy := &dummyv1.Dummy{
			Name:     "name",
			Value:    100,
			Values:   []string{"a", "b", "c"},
			TestType: dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO,
			ConfigA: &dummyv1.DummyConfigA{
				Name: "config_a",
			},
			ConfigB: &dummyv1.DummyConfigB{
				Name: "config_b",
			},
			ConfigB2: &dummyv1.DummyConfigB{
				Name: "config_b2",
			},
		}

		m, err := NewMasked(dummy, "name", "value", "config_a.name")
		require.NoError(t, err)
		require.NotNil(t, m)

		patch := &dummyv1.Dummy{
			Name:     "name name",
			Value:    100_100,
			Values:   []string{"a", "b", "c", "a", "b", "c"},
			TestType: dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR,
			ConfigA: &dummyv1.DummyConfigA{
				Name: "config_a config_a",
			},
			ConfigB: &dummyv1.DummyConfigB{
				Name: "config_b config_b",
			},
			ConfigB2: &dummyv1.DummyConfigB{
				Name: "config_b2 config_b2",
			},
		}

		msg, err := m.Patch(patch)
		assert.ErrorContains(t, err, `nested field "config_a.name" is not supported`)
		assert.Nil(t, msg)
	})
}
