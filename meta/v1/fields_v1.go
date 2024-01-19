package v1

import (
	"fmt"

	"github.com/phuslu/kubeproto/easyproto"
)

// FieldsV1 stores a set of fields in a data structure like a Trie, in JSON format.
//
// Each key is either a '.' representing the field itself, and will always map to an empty set,
// or a string representing a sub-field or item. The string will follow one of these four formats:
// 'f:<name>', where <name> is the name of a field in a struct, or key in a map
// 'v:<value>', where <value> is the exact json formatted value of a list item
// 'i:<index>', where <index> is position of a item in a list
// 'k:<keys>', where <keys> is a map of  a list item's key fields to their unique values
// If a key maps to an empty Fields value, the field that key represents is part of the set.
//
// The exact format is defined in sigs.k8s.io/structured-merge-diff
// +protobuf.options.(gogoproto.goproto_stringer)=false
type FieldsV1 struct {
	// Raw is the underlying serialization of this object.
	Raw []byte `json:"-" protobuf:"bytes,1,opt,name=Raw"`
}

// UnmarshalProtobuf unmarshals f from protobuf message at src.
func (f *FieldsV1) UnmarshalProtobuf(src []byte) (err error) {
	// Set default FieldsV1 values
	f.Raw = nil

	// Parse FieldsV1 message at src
	var fc easyproto.FieldContext
	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return fmt.Errorf("cannot read next field in FieldsV1 message")
		}
		switch fc.FieldNum {
		case 1:
			raw, ok := fc.Bytes()
			if !ok {
				return fmt.Errorf("cannot read FieldsV1 name")
			}
			f.Raw = raw
		}
	}
	return
}
