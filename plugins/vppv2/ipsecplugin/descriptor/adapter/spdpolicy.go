// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/gogo/protobuf/proto"
	. "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	"github.com/ligato/vpp-agent/plugins/vppv2/model/ipsec"
)

////////// type-safe key-value pair with metadata //////////

type SPDPolicyKVWithMetadata struct {
	Key      string
	Value    *ipsec.SecurityPolicyDatabase_PolicyEntry
	Metadata interface{}
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type SPDPolicyDescriptor struct {
	Name               string
	KeySelector        KeySelector
	ValueTypeName      string
	KeyLabel           func(key string) string
	ValueComparator    func(key string, oldValue, newValue *ipsec.SecurityPolicyDatabase_PolicyEntry) bool
	NBKeyPrefix        string
	WithMetadata       bool
	MetadataMapFactory MetadataMapFactory
	Add                func(key string, value *ipsec.SecurityPolicyDatabase_PolicyEntry) (metadata interface{}, err error)
	Delete             func(key string, value *ipsec.SecurityPolicyDatabase_PolicyEntry, metadata interface{}) error
	Modify             func(key string, oldValue, newValue *ipsec.SecurityPolicyDatabase_PolicyEntry, oldMetadata interface{}) (newMetadata interface{}, err error)
	ModifyWithRecreate func(key string, oldValue, newValue *ipsec.SecurityPolicyDatabase_PolicyEntry, metadata interface{}) bool
	Update             func(key string, value *ipsec.SecurityPolicyDatabase_PolicyEntry, metadata interface{}) error
	IsRetriableFailure func(err error) bool
	Dependencies       func(key string, value *ipsec.SecurityPolicyDatabase_PolicyEntry) []Dependency
	DerivedValues      func(key string, value *ipsec.SecurityPolicyDatabase_PolicyEntry) []KeyValuePair
	Dump               func(correlate []SPDPolicyKVWithMetadata) ([]SPDPolicyKVWithMetadata, error)
	DumpDependencies   []string /* descriptor name */
}

////////// Descriptor adapter //////////

type SPDPolicyDescriptorAdapter struct {
	descriptor *SPDPolicyDescriptor
}

func NewSPDPolicyDescriptor(typedDescriptor *SPDPolicyDescriptor) *KVDescriptor {
	adapter := &SPDPolicyDescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:               typedDescriptor.Name,
		KeySelector:        typedDescriptor.KeySelector,
		ValueTypeName:      typedDescriptor.ValueTypeName,
		KeyLabel:           typedDescriptor.KeyLabel,
		NBKeyPrefix:        typedDescriptor.NBKeyPrefix,
		WithMetadata:       typedDescriptor.WithMetadata,
		MetadataMapFactory: typedDescriptor.MetadataMapFactory,
		IsRetriableFailure: typedDescriptor.IsRetriableFailure,
		DumpDependencies:   typedDescriptor.DumpDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Add != nil {
		descriptor.Add = adapter.Add
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Modify != nil {
		descriptor.Modify = adapter.Modify
	}
	if typedDescriptor.ModifyWithRecreate != nil {
		descriptor.ModifyWithRecreate = adapter.ModifyWithRecreate
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	if typedDescriptor.Dump != nil {
		descriptor.Dump = adapter.Dump
	}
	return descriptor
}

func (da *SPDPolicyDescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castSPDPolicyValue(key, oldValue)
	typedNewValue, err2 := castSPDPolicyValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *SPDPolicyDescriptorAdapter) Add(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castSPDPolicyValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Add(key, typedValue)
}

func (da *SPDPolicyDescriptorAdapter) Modify(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castSPDPolicyValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castSPDPolicyValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castSPDPolicyMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Modify(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *SPDPolicyDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castSPDPolicyValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castSPDPolicyMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *SPDPolicyDescriptorAdapter) ModifyWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castSPDPolicyValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castSPDPolicyValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castSPDPolicyMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.ModifyWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *SPDPolicyDescriptorAdapter) Update(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castSPDPolicyValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castSPDPolicyMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Update(key, typedValue, typedMetadata)
}

func (da *SPDPolicyDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castSPDPolicyValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

func (da *SPDPolicyDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castSPDPolicyValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *SPDPolicyDescriptorAdapter) Dump(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []SPDPolicyKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castSPDPolicyValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castSPDPolicyMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			SPDPolicyKVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedDump, err := da.descriptor.Dump(correlateWithType)
	if err != nil {
		return nil, err
	}
	var dump []KVWithMetadata
	for _, typedKVWithMetadata := range typedDump {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		dump = append(dump, kvWithMetadata)
	}
	return dump, err
}

////////// Helper methods //////////

func castSPDPolicyValue(key string, value proto.Message) (*ipsec.SecurityPolicyDatabase_PolicyEntry, error) {
	typedValue, ok := value.(*ipsec.SecurityPolicyDatabase_PolicyEntry)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castSPDPolicyMetadata(key string, metadata Metadata) (interface{}, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(interface{})
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}