package messagemagic

import (
	"encoding/json"
	"testing"

	dummyv1 "github.com/gochore/protomagic/testdata/gen/dummy/v1"

	"github.com/gochore/pt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func assertMessageEqual(t *testing.T, expected, actual proto.Message) {
	if !assert.True(t, proto.Equal(expected, actual)) {
		expectedJson, _ := json.Marshal(expected)
		actualJson, _ := json.Marshal(actual)
		assert.JSONEq(t, string(expectedJson), string(actualJson)) // to show the diff
	}
}

func TestPatch(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		msg := &dummyv1.DummyA{
			Name:      "name",
			Value:     1,
			Values:    []string{"a", "b"},
			TestType:  dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO,
			ConfigA:   &dummyv1.DummyConfigA{Name: "name"},
			OName:     pt.P("o_name"),
			OValue:    pt.P[int32](2),
			OTestType: pt.P(dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR),
			OConfigA:  &dummyv1.DummyConfigA{Name: "name"},
		}
		patch := &dummyv1.DummyA{
			Name:      "name2",
			Value:     2,
			Values:    []string{"c", "d"},
			TestType:  dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR,
			ConfigA:   &dummyv1.DummyConfigA{Name: "name2"},
			OName:     pt.P("o_name2"),
			OValue:    pt.P[int32](3),
			OTestType: pt.P(dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO),
			OConfigA:  &dummyv1.DummyConfigA{Name: "name2"},
		}
		want := &dummyv1.DummyA{
			Name:      "name2",
			Value:     2,
			Values:    []string{"c", "d"},
			TestType:  dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR,
			ConfigA:   &dummyv1.DummyConfigA{Name: "name2"},
			OName:     pt.P("o_name2"),
			OValue:    pt.P[int32](3),
			OTestType: pt.P(dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO),
			OConfigA:  &dummyv1.DummyConfigA{Name: "name2"},
		}
		got := Patch(msg, patch)
		assertMessageEqual(t, want, got)
	})

	t.Run("empty patch", func(t *testing.T) {
		msg := &dummyv1.DummyA{
			Name:      "name",
			Value:     1,
			Values:    []string{"a", "b"},
			TestType:  dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO,
			ConfigA:   &dummyv1.DummyConfigA{Name: "name"},
			OName:     pt.P("o_name"),
			OValue:    pt.P[int32](2),
			OTestType: pt.P(dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR),
			OConfigA:  &dummyv1.DummyConfigA{Name: "name"},
		}
		patch := &dummyv1.DummyA{
			Name:      "",
			Value:     0,
			Values:    nil,
			TestType:  0,
			ConfigA:   nil,
			OName:     nil,
			OValue:    nil,
			OTestType: nil,
			OConfigA:  nil,
		}
		want := &dummyv1.DummyA{
			Name:      "",
			Value:     0,
			Values:    nil,
			TestType:  0,
			ConfigA:   nil,
			OName:     pt.P("o_name"),
			OValue:    pt.P[int32](2),
			OTestType: pt.P(dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR),
			OConfigA:  &dummyv1.DummyConfigA{Name: "name"},
		}
		got := Patch(msg, patch)
		assertMessageEqual(t, want, got)
	})

	t.Run("empty msg", func(t *testing.T) {
		msg := &dummyv1.DummyA{
			Name:      "",
			Value:     0,
			Values:    nil,
			TestType:  0,
			ConfigA:   nil,
			OName:     nil,
			OValue:    nil,
			OTestType: nil,
			OConfigA:  nil,
		}
		patch := &dummyv1.DummyA{
			Name:      "name2",
			Value:     2,
			Values:    []string{"c", "d"},
			TestType:  dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR,
			ConfigA:   &dummyv1.DummyConfigA{Name: "name2"},
			OName:     pt.P("o_name2"),
			OValue:    pt.P[int32](3),
			OTestType: pt.P(dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO),
			OConfigA:  &dummyv1.DummyConfigA{Name: "name2"},
		}
		want := &dummyv1.DummyA{
			Name:      "name2",
			Value:     2,
			Values:    []string{"c", "d"},
			TestType:  dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR,
			ConfigA:   &dummyv1.DummyConfigA{Name: "name2"},
			OName:     pt.P("o_name2"),
			OValue:    pt.P[int32](3),
			OTestType: pt.P(dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO),
			OConfigA:  &dummyv1.DummyConfigA{Name: "name2"},
		}
		got := Patch(msg, patch)
		assertMessageEqual(t, want, got)
	})

	t.Run("zero patch", func(t *testing.T) {
		msg := &dummyv1.DummyA{
			Name:      "name",
			Value:     1,
			Values:    []string{"a", "b"},
			TestType:  dummyv1.TestEnumType_TEST_ENUM_TYPE_FOO,
			ConfigA:   &dummyv1.DummyConfigA{Name: "name"},
			OName:     pt.P("o_name"),
			OValue:    pt.P[int32](2),
			OTestType: pt.P(dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR),
			OConfigA:  &dummyv1.DummyConfigA{Name: "name"},
		}
		patch := &dummyv1.DummyA{
			Name:      "name2",
			Value:     2,
			Values:    []string{"c", "d"},
			TestType:  dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR,
			ConfigA:   &dummyv1.DummyConfigA{Name: "name2"},
			OName:     pt.P(""),
			OValue:    pt.P[int32](0),
			OTestType: pt.P(dummyv1.TestEnumType_TEST_ENUM_TYPE_UNSPECIFIED),
			OConfigA:  &dummyv1.DummyConfigA{},
		}
		want := &dummyv1.DummyA{
			Name:      "name2",
			Value:     2,
			Values:    []string{"c", "d"},
			TestType:  dummyv1.TestEnumType_TEST_ENUM_TYPE_BAR,
			ConfigA:   &dummyv1.DummyConfigA{Name: "name2"},
			OName:     pt.P(""),
			OValue:    pt.P[int32](0),
			OTestType: pt.P(dummyv1.TestEnumType_TEST_ENUM_TYPE_UNSPECIFIED),
			OConfigA:  &dummyv1.DummyConfigA{},
		}
		got := Patch(msg, patch)
		assertMessageEqual(t, want, got)
	})
}
