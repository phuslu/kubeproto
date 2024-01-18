package v1

import (
	"fmt"

	"github.com/VictoriaMetrics/easyproto"
)

// TypeMeta describes an individual object in an API response or request
// with strings representing the type of the object and its API schema version.
// Structures that are versioned or persisted should inline TypeMeta.
//
// +k8s:deepcopy-gen=false
type TypeMeta struct {
	// Kind is a string value representing the REST resource this object represents.
	// Servers may infer this from the endpoint the client submits requests to.
	// Cannot be updated.
	// In CamelCase.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	// +optional
	Kind string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`

	// APIVersion defines the versioned schema of this representation of an object.
	// Servers should convert recognized schemas to the latest internal value, and
	// may reject unrecognized values.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	// +optional
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,2,opt,name=apiVersion"`
}

// UnmarshalProtobuf unmarshals tm from protobuf message at src.
func (tm *TypeMeta) UnmarshalProtobuf(src []byte) (err error) {
	// Set default TypeMeta values
	tm.Kind = ""
	tm.APIVersion = ""

	// Parse TypeMeta message at src
	var fc easyproto.FieldContext
	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return fmt.Errorf("cannot read next field in TypeMeta message")
		}
		switch fc.FieldNum {
		case 1:
			kind, ok := fc.String()
			if !ok {
				return fmt.Errorf("cannot read TypeMeta name")
			}
			tm.Kind = kind
		case 2:
			apiversion, ok := fc.String()
			if !ok {
				return fmt.Errorf("cannot read TypeMeta name")
			}
			tm.APIVersion = apiversion
		}
	}
	return nil
}
