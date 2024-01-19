package v1

import (
	"fmt"

	"github.com/phuslu/kubeproto/easyproto"
)

// Timestamp is a struct that is equivalent to Time, but intended for
// protobuf marshalling/unmarshalling. It is generated into a serialization
// that matches Time. Do not use in Go structs.
type Timestamp struct {
	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `json:"seconds" protobuf:"varint,1,opt,name=seconds"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive. This field may be limited in precision depending on context.
	Nanos int32 `json:"nanos" protobuf:"varint,2,opt,name=nanos"`
}

// UnmarshalProtobuf unmarshals ts from protobuf message at src.
func (ts *Timestamp) UnmarshalProtobuf(src []byte) (err error) {
	// Set default Timestamp values
	ts.Seconds = 0
	ts.Nanos = 0

	// Parse TypeMeta message at src
	var fc easyproto.FieldContext
	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return fmt.Errorf("cannot read next field in TypeMeta message")
		}
		switch fc.FieldNum {
		case 1:
			seconds, ok := fc.Int64()
			if !ok {
				return fmt.Errorf("cannot read TypeMeta name")
			}
			ts.Seconds = seconds
		case 2:
			nanos, ok := fc.Int32()
			if !ok {
				return fmt.Errorf("cannot read TypeMeta name")
			}
			ts.Nanos = nanos
		}
	}
	return
}
