package v1

import (
	"fmt"

	"github.com/phuslu/kubeproto/easyproto"
	metav1 "github.com/phuslu/kubeproto/meta/v1"
)

type Pod struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of the pod.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	// Spec PodSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Most recently observed status of the pod.
	// This data may not be up to date.
	// Populated by the system.
	// Read-only.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	// Status PodStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// UnmarshalProtobuf unmarshals ts from protobuf message at src.
func (pod *Pod) UnmarshalProtobuf(src []byte) (err error) {
	// Set default pod values
	pod.TypeMeta.APIVersion = "v1"
	pod.TypeMeta.Kind = "Pod"
	pod.ObjectMeta = metav1.ObjectMeta{}

	// Parse TypeMeta message at src
	var fc easyproto.FieldContext
	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return fmt.Errorf("cannot read next field in TypeMeta message")
		}
		switch fc.FieldNum {
		case 1:
			data, ok := fc.MessageData()
			if !ok {
				return fmt.Errorf("cannot read Pod ObjectMeta name")
			}
			if err := pod.ObjectMeta.UnmarshalProtobuf(data); err != nil {
				return fmt.Errorf("cannot unmarshal Pod ObjectMeta: %w", err)
			}
		}
	}
	return

}
