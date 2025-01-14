/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SinkInOneOf4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SinkInOneOf4{}

// SinkInOneOf4 struct for SinkInOneOf4
type SinkInOneOf4 struct {
	Type string `json:"type"`
}

type _SinkInOneOf4 SinkInOneOf4

// NewSinkInOneOf4 instantiates a new SinkInOneOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSinkInOneOf4(type_ string) *SinkInOneOf4 {
	this := SinkInOneOf4{}
	this.Type = type_
	return &this
}

// NewSinkInOneOf4WithDefaults instantiates a new SinkInOneOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSinkInOneOf4WithDefaults() *SinkInOneOf4 {
	this := SinkInOneOf4{}
	return &this
}

// GetType returns the Type field value
func (o *SinkInOneOf4) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SinkInOneOf4) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SinkInOneOf4) SetType(v string) {
	o.Type = v
}

func (o SinkInOneOf4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SinkInOneOf4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *SinkInOneOf4) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSinkInOneOf4 := _SinkInOneOf4{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSinkInOneOf4)

	if err != nil {
		return err
	}

	*o = SinkInOneOf4(varSinkInOneOf4)

	return err
}

type NullableSinkInOneOf4 struct {
	value *SinkInOneOf4
	isSet bool
}

func (v NullableSinkInOneOf4) Get() *SinkInOneOf4 {
	return v.value
}

func (v *NullableSinkInOneOf4) Set(val *SinkInOneOf4) {
	v.value = val
	v.isSet = true
}

func (v NullableSinkInOneOf4) IsSet() bool {
	return v.isSet
}

func (v *NullableSinkInOneOf4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSinkInOneOf4(val *SinkInOneOf4) *NullableSinkInOneOf4 {
	return &NullableSinkInOneOf4{value: val, isSet: true}
}

func (v NullableSinkInOneOf4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSinkInOneOf4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


