package v1

import (
	"fmt"

	"github.com/phuslu/kubeproto/internal/easyproto"
)

// ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource
// that the fieldset applies to.
type ManagedFieldsEntry struct {
	// Manager is an identifier of the workflow managing these fields.
	Manager string `json:"manager,omitempty" protobuf:"bytes,1,opt,name=manager"`
	// Operation is the type of operation which lead to this ManagedFieldsEntry being created.
	// The only valid values for this field are 'Apply' and 'Update'.
	Operation string `json:"operation,omitempty" protobuf:"bytes,2,opt,name=operation,casttype=ManagedFieldsOperationType"`
	// APIVersion defines the version of this resource that this field set
	// applies to. The format is "group/version" just like the top-level
	// APIVersion field. It is necessary to track the version of a field
	// set because it cannot be automatically converted.
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,3,opt,name=apiVersion"`
	// Time is the timestamp of when the ManagedFields entry was added. The
	// timestamp will also be updated if a field is added, the manager
	// changes any of the owned fields value or removes a field. The
	// timestamp does not update when a field is removed from the entry
	// because another manager took it over.
	// +optional
	Time *Time `json:"time,omitempty" protobuf:"bytes,4,opt,name=time"`

	// Fields is tombstoned to show why 5 is a reserved protobuf tag.
	//Fields *Fields `json:"fields,omitempty" protobuf:"bytes,5,opt,name=fields,casttype=Fields"`

	// FieldsType is the discriminator for the different fields format and version.
	// There is currently only one possible value: "FieldsV1"
	FieldsType string `json:"fieldsType,omitempty" protobuf:"bytes,6,opt,name=fieldsType"`
	// FieldsV1 holds the first JSON version format as described in the "FieldsV1" type.
	// +optional
	FieldsV1 *FieldsV1 `json:"fieldsV1,omitempty" protobuf:"bytes,7,opt,name=fieldsV1"`

	// Subresource is the name of the subresource used to update that object, or
	// empty string if the object was updated through the main resource. The
	// value of this field is used to distinguish between managers, even if they
	// share the same name. For example, a status update will be distinct from a
	// regular update using the same manager name.
	// Note that the APIVersion field is not related to the Subresource field and
	// it always corresponds to the version of the main resource.
	Subresource string `json:"subresource,omitempty" protobuf:"bytes,8,opt,name=subresource"`
}

// UnmarshalProtobuf unmarshals or from protobuf message at src.
func (mfe *ManagedFieldsEntry) UnmarshalProtobuf(src []byte) (err error) {
	// Set default ManagedFieldsEntry values
	mfe.Manager = ""
	mfe.Operation = ""
	mfe.APIVersion = ""
	if mfe.Time != nil {
		*mfe.Time = Time{}
	}
	mfe.FieldsType = ""
	if mfe.FieldsV1 != nil {
		mfe.FieldsV1.Raw = mfe.FieldsV1.Raw[:0]
	}
	mfe.Subresource = ""

	// Parse ManagedFieldsEntry message at src
	var fc easyproto.FieldContext
	var ok bool
	for len(src) > 0 {
		src, err = fc.NextField(src)
		if err != nil {
			return fmt.Errorf("cannot read next field in ManagedFieldsEntry message")
		}
		switch fc.FieldNum {
		case 1:
			mfe.Manager, ok = fc.String()
			if !ok {
				return fmt.Errorf("cannot read ManagedFieldsEntry name")
			}
		case 2:
			mfe.Operation, ok = fc.String()
			if !ok {
				return fmt.Errorf("cannot read ManagedFieldsEntry name")
			}
		case 3:
			mfe.APIVersion, ok = fc.String()
			if !ok {
				return fmt.Errorf("cannot read ManagedFieldsEntry name")
			}
		case 4:
			data, ok := fc.MessageData()
			if !ok {
				return fmt.Errorf("cannot read ManagedFieldsEntry name")
			}
			if mfe.Time == nil {
				mfe.Time = new(Time)
			}
			if err := mfe.Time.UnmarshalProtobuf(data); err != nil {
				return fmt.Errorf("cannot unmarshal ManagedFieldsEntry: %w", err)
			}
		case 5:
			mfe.FieldsType, ok = fc.String()
			if !ok {
				return fmt.Errorf("cannot read ManagedFieldsEntry name")
			}
		case 6:
			data, ok := fc.MessageData()
			if !ok {
				return fmt.Errorf("cannot read ManagedFieldsEntry name")
			}
			if mfe.FieldsV1 == nil {
				mfe.FieldsV1 = new(FieldsV1)
			}
			if err := mfe.FieldsV1.UnmarshalProtobuf(data); err != nil {
				return fmt.Errorf("cannot unmarshal ManagedFieldsEntry: %w", err)
			}
		case 7:
			mfe.Subresource, ok = fc.String()
			if !ok {
				return fmt.Errorf("cannot read ManagedFieldsEntry name")
			}
		}
	}
	return
}
