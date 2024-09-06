package fieldmaskmagic

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

func TestFromFields(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		mask, err := FromFields(&dummyv1.Dummy{}, &dummyv1.DummyConfigA{})
		assert.NoError(t, err)
		require.NotNil(t, mask)
		assertMessageEqual(t, &fieldmaskpb.FieldMask{Paths: []string{"config_a"}}, mask)
	})
	t.Run("not found", func(t *testing.T) {
		mask, err := FromFields(&dummyv1.Dummy{}, &dummyv1.DummyConfigA{}, &fieldmaskpb.FieldMask{})
		assert.ErrorContains(t, err, `field "FieldMask" not found in "Dummy"`)
		assert.Nil(t, mask)
	})
	t.Run("ambiguous", func(t *testing.T) {
		mask, err := FromFields(&dummyv1.Dummy{}, &dummyv1.DummyConfigA{}, &dummyv1.DummyConfigB{})
		assert.ErrorContains(t, err, `field "DummyConfigB" is ambiguous in "Dummy"`)
		assert.Nil(t, mask)
	})
}

func TestPrune(t *testing.T) {
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

		mask, err := fieldmaskpb.New(dummy, "name", "value", "config_a")
		require.NoError(t, err)
		require.NotNil(t, mask)

		msg, err := Prune(mask, dummy)
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

		mask, err := fieldmaskpb.New(dummy)
		require.NoError(t, err)
		require.NotNil(t, mask)

		msg, err := Prune(mask, dummy)
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

		mask, err := fieldmaskpb.New(dummy, "name", "value", "config_a.name")
		require.NoError(t, err)
		require.NotNil(t, mask)

		msg, err := Prune(mask, dummy)
		assert.ErrorContains(t, err, `nested field "config_a.name" is not supported`)
		assert.Nil(t, msg)
	})
}

func TestPatch(t *testing.T) {
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

		mask, err := fieldmaskpb.New(dummy, "name", "value", "config_a")
		require.NoError(t, err)
		require.NotNil(t, mask)

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

		msg, err := Patch(mask, dummy, patch)
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

		mask, err := fieldmaskpb.New(dummy)
		require.NoError(t, err)
		require.NotNil(t, mask)

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

		msg, err := Patch(mask, dummy, patch)
		require.NoError(t, err)
		assertMessageEqual(t, dummy, msg)
	})

	t.Run("empty message", func(t *testing.T) {
		var dummy *dummyv1.Dummy

		mask, err := fieldmaskpb.New(dummy, "name", "value", "config_a")
		require.NoError(t, err)
		require.NotNil(t, mask)

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

		msg, err := Patch(mask, dummy, patch)
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

		mask, err := fieldmaskpb.New(dummy, "name", "value", "config_a")
		require.NoError(t, err)
		require.NotNil(t, mask)

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

		msg, err := Patch(mask, dummy, patch)
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

		mask, err := fieldmaskpb.New(dummy, "name", "value", "config_a.name")
		require.NoError(t, err)
		require.NotNil(t, mask)

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

		msg, err := Patch(mask, dummy, patch)
		assert.ErrorContains(t, err, `nested field "config_a.name" is not supported`)
		assert.Nil(t, msg)
	})
}
