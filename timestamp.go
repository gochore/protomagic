package protomagic

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Timestamp is a wrapper around timestamppb.Timestamp.
// It mainly optimizes the processing logic for zero values.
type Timestamp timestamppb.Timestamp

func WrapTimestamp(inner *timestamppb.Timestamp) *Timestamp {
	return (*Timestamp)(inner)
}

func WrapTime(t time.Time) *Timestamp {
	if t.IsZero() {
		return WrapTimestamp(&timestamppb.Timestamp{})
	}
	return WrapTimestamp(timestamppb.New(t))
}

func (t *Timestamp) Unwrap() *timestamppb.Timestamp {
	return (*timestamppb.Timestamp)(t)
}

func (t *Timestamp) AsTime() time.Time {
	if t.IsZeroTimestamp() {
		return time.Time{}
	}
	return t.Unwrap().AsTime()
}

func (t *Timestamp) IsZero() bool {
	return t.IsZeroTimestamp() || t.IsZeroTime()
}

func (t *Timestamp) IsZeroTimestamp() bool {
	ts := t.Unwrap()
	return ts.Seconds == 0 && ts.Nanos == 0
}

func (t *Timestamp) IsZeroTime() bool {
	return t.Unwrap().AsTime().IsZero()
}
