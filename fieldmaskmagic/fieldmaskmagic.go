package fieldmaskmagic

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// FromFields returns a new FieldMask with given fields.
// It determines the paths by finding the fields with the same type in the message.
// So it works only when:
//  1. The specified fields are all message types.
//  2. The msg does not contain more than one field with the same type.
func FromFields(msg proto.Message, fields ...proto.Message) (*fieldmaskpb.FieldMask, error) {
	typeToFields := map[protoreflect.FullName][]protoreflect.FieldDescriptor{}
	msgDesc := msg.ProtoReflect().Descriptor()
	msgFields := msgDesc.Fields()
	for i := 0; i < msgFields.Len(); i++ {
		fd := msgFields.Get(i)
		if fd.Kind() != protoreflect.MessageKind {
			continue
		}
		fdm := fd.Message().FullName()
		typeToFields[fdm] = append(typeToFields[fdm], fd)
	}

	paths := make([]string, 0, len(fields))
	for _, field := range fields {
		fd := field.ProtoReflect().Descriptor()
		fds := typeToFields[fd.FullName()]
		if len(fds) == 0 {
			return nil, fmt.Errorf("field %q not found in %q", fd.Name(), msgDesc.Name())
		}
		if len(fds) > 1 {
			return nil, fmt.Errorf("field %q is ambiguous in %q", fd.Name(), msgDesc.Name())
		}
		paths = append(paths, string(fds[0].Name()))
	}

	return fieldmaskpb.New(msg, paths...)
}

// Prune returns a copy of msg with all fields not in mask removed.
// If the mask is empty, it returns a copy of msg.
// Please note that it does check recursively.
func Prune[T proto.Message](mask *fieldmaskpb.FieldMask, msg T) (T, error) {
	ret := proto.Clone(msg)

	if len(mask.GetPaths()) == 0 {
		return ret.(T), nil
	}

	// FIXME: nested field is not supported
	for _, p := range mask.GetPaths() {
		if strings.Contains(p, ".") {
			var nilT T
			return nilT, fmt.Errorf("nested field %q is not supported", p)
		}
	}

	refl := ret.ProtoReflect()

	paths := map[string]struct{}{}
	for _, p := range mask.GetPaths() {
		paths[p] = struct{}{}
	}
	refl.Range(func(fd protoreflect.FieldDescriptor, _ protoreflect.Value) bool {
		if _, ok := paths[string(fd.Name())]; !ok {
			refl.Clear(fd)
		}
		return true
	})

	return ret.(T), nil
}

// Patch returns a copy of origin with all fields in patch applied.
// If the mask is empty, it returns a copy of the origin message.
// Please note that it does check recursively.
func Patch[T proto.Message](mask *fieldmaskpb.FieldMask, msg, patch T) (T, error) {
	patchRefl := patch.ProtoReflect()

	ret := proto.Clone(msg)

	if len(mask.GetPaths()) == 0 {
		return ret.(T), nil
	}

	// FIXME: nested field is not supported
	for _, p := range mask.GetPaths() {
		if strings.Contains(p, ".") {
			var nilT T
			return nilT, fmt.Errorf("nested field %q is not supported", p)
		}
	}

	if !msg.ProtoReflect().IsValid() {
		return ret.(T), nil
	}

	paths := map[string]struct{}{}
	for _, p := range mask.GetPaths() {
		paths[p] = struct{}{}
	}

	refl := ret.ProtoReflect()
	fields := refl.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ { // do not use Range, it will ignore unset fields
		fd := fields.Get(i)
		if _, ok := paths[string(fd.Name())]; ok {
			if !patchRefl.Has(fd) {
				refl.Clear(fd)
			} else {
				refl.Set(fd, patchRefl.Get(fd))
			}
		}
	}

	return ret.(T), nil
}
