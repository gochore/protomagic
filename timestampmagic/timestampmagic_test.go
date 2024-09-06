package timestampmagic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNew(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		tm := time.Unix(123456, 789)
		ts := New(tm)
		assert.Equal(t, ts.Seconds, int64(123456))
		assert.Equal(t, ts.Nanos, int32(789))
	})
	t.Run("zero", func(t *testing.T) {
		zero := time.Time{}
		ts := New(zero)
		assert.Equal(t, ts.Seconds, int64(0))
		assert.Equal(t, ts.Nanos, int32(0))
	})
}

func TestAsTime(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		ts := &timestamppb.Timestamp{
			Seconds: 123456,
			Nanos:   789,
		}
		tm := AsTime(ts)
		assert.Equal(t, tm.Unix(), int64(123456))
		assert.Equal(t, tm.Nanosecond(), 789)
	})
	t.Run("zero", func(t *testing.T) {
		ts := &timestamppb.Timestamp{}
		tm := AsTime(ts)
		assert.True(t, tm.IsZero())
	})
}

func TestIsZero(t *testing.T) {
	t.Run("not zero", func(t *testing.T) {
		ts := &timestamppb.Timestamp{
			Seconds: 123456,
			Nanos:   789,
		}
		assert.False(t, IsZero(ts))
	})
	t.Run("zero time", func(t *testing.T) {
		zero := time.Time{}
		ts := &timestamppb.Timestamp{
			Seconds: zero.Unix(),
			Nanos:   int32(zero.Nanosecond()),
		}
		assert.True(t, IsZero(ts))
	})
	t.Run("zero timestamp", func(t *testing.T) {
		ts := &timestamppb.Timestamp{}
		assert.True(t, IsZero(ts))
	})
}
