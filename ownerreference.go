package kubeproto

import (
	"fmt"

	"github.com/VictoriaMetrics/easyproto"
)

// OwnerReference contains enough information to let you identify an owning
// object. An owning object must be in the same namespace as the dependent, or
// be cluster-scoped, so there is no namespace field.
// +structType=atomic
type OwnerReference struct {
	// API version of the referent.
	APIVersion string `json:"apiVersion" protobuf:"bytes,5,opt,name=apiVersion"`
	// Kind of the referent.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind" protobuf:"bytes,1,opt,name=kind"`
	// Name of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names
	Name string `json:"name" protobuf:"bytes,3,opt,name=name"`
	// UID of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids
	UID string `json:"uid" protobuf:"bytes,4,opt,name=uid,casttype=k8s.io/apimachinery/pkg/types.UID"`
	// If true, this reference points to the managing controller.
	// +optional
	Controller *bool `json:"controller,omitempty" protobuf:"varint,6,opt,name=controller"`
	// If true, AND if the owner has the "foregroundDeletion" finalizer, then
	// the owner cannot be deleted from the key-value store until this
	// reference is removed.
	// See https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletion
	// for how the garbage collector interacts with this field and enforces the foreground deletion.
	// Defaults to false.
	// To set this field, a user needs "delete" permission of the owner,
	// otherwise 422 (Unprocessable Entity) will be returned.
	// +optional
	BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty" protobuf:"varint,7,opt,name=blockOwnerDeletion"`
}

// UnmarshalProtobuf unmarshals or from protobuf message at src.
func (or *OwnerReference) UnmarshalProtobuf(src []byte) (err error) {
	// Set default OwnerReference values
	or.APIVersion = ""
	or.Kind = ""
	or.Name = ""
	or.UID = ""
	or.Controller = nil
	or.BlockOwnerDeletion = nil

	// Parse OwnerReference message at src
	var fc easyproto.FieldContext
	var ok bool
	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return fmt.Errorf("cannot read next field in OwnerReference message")
		}
		switch fc.FieldNum {
		case 1:
			or.APIVersion, ok = fc.String()
			if !ok {
				return fmt.Errorf("cannot read OwnerReference name")
			}
		case 2:
			or.Kind, ok = fc.String()
			if !ok {
				return fmt.Errorf("cannot read OwnerReference name")
			}
		case 3:
			or.Name, ok = fc.String()
			if !ok {
				return fmt.Errorf("cannot read OwnerReference name")
			}
		case 4:
			or.UID, ok = fc.String()
			if !ok {
				return fmt.Errorf("cannot read OwnerReference name")
			}
		case 5:
			controller, ok := fc.Bool()
			if !ok {
				return fmt.Errorf("cannot read OwnerReference name")
			}
			or.Controller = &controller
		case 6:
			blockOwnerDeletion, ok := fc.Bool()
			if !ok {
				return fmt.Errorf("cannot read OwnerReference name")
			}
			or.BlockOwnerDeletion = &blockOwnerDeletion
		}
	}
	return
}
