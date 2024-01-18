package v1

import (
	"time"
)

// Time is a wrapper around time.Time which supports correct
// marshaling to YAML and JSON.  Wrappers are provided for many
// of the factory methods that the time package offers.
//
// +protobuf.options.marshal=false
// +protobuf.as=Timestamp
// +protobuf.options.(gogoproto.goproto_stringer)=false
type Time struct {
	time.Time `protobuf:"-"`
}

// UnmarshalProtobuf unmarshals ts from protobuf message at src.
func (t *Time) UnmarshalProtobuf(src []byte) (err error) {
	// Set default Timestamp values
	var ts Timestamp

	err = ts.UnmarshalProtobuf(src)
	if err != nil {
		return
	}

	t.Time = time.Unix(ts.Seconds, int64(ts.Nanos))
	return
}
