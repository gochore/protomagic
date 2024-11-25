package messagemagic

import (
	"google.golang.org/protobuf/proto"
)

// Patch patches the message with the patch.
// It copies all not optional fields from the patch to the message.
// For optional fields, it copies the value only if the patch has the field.
// Otherwise, it keeps the value in the message unchanged.
func Patch[T proto.Message](msg, patch T) T {
	dst := proto.Clone(msg).(T)

	dstRefl, srcRefl := dst.ProtoReflect(), patch.ProtoReflect()
	for i := range srcRefl.Descriptor().Fields().Len() {
		fd := srcRefl.Descriptor().Fields().Get(i)
		if fd.HasOptionalKeyword() {
			if srcRefl.Has(fd) {
				dstRefl.Set(fd, srcRefl.Get(fd))
			}
		} else {
			if srcRefl.Has(fd) {
				dstRefl.Set(fd, srcRefl.Get(fd))
			} else {
				dstRefl.Clear(fd)
			}
		}
	}

	return dst
}
