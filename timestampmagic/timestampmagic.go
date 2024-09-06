package timestampmagic

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// New returns a new timestamppb.Timestamp.
// It works like timestamppb.New,
// but it treats zero time as an empty timestamp.
func New(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return &timestamppb.Timestamp{}
	}
	return timestamppb.New(t)
}

// AsTime returns the time.Time value of the timestamp.
// It works like timestamppb.Timestamp.AsTime,
// but it treats an empty timestamp as a zero time.
// Please note that it doesn't support "1970-01-01 00:00:00" and returns time.Time{} instead.
func AsTime(ts *timestamppb.Timestamp) time.Time {
	if IsZeroTimestamp(ts) {
		return time.Time{}
	}
	return ts.AsTime()
}

// IsZero returns true if the timestamp is zero or the time is zero.
func IsZero(ts *timestamppb.Timestamp) bool {
	return IsZeroTimestamp(ts) || IsZeroTime(ts)
}

// IsZeroTime returns true if the time is zero.
// Use IsZero instead if you don't have a specific reason.
func IsZeroTime(ts *timestamppb.Timestamp) bool {
	return ts.AsTime().IsZero()
}

// IsZeroTimestamp returns true if the timestamp is zero.
// Use IsZero instead if you don't have a specific reason.
func IsZeroTimestamp(ts *timestamppb.Timestamp) bool {
	return ts.GetSeconds() == 0 && ts.GetNanos() == 0
}
