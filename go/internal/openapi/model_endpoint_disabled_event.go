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

// checks if the EndpointDisabledEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointDisabledEvent{}

// EndpointDisabledEvent Sent when an endpoint has been automatically disabled after continuous failures.
type EndpointDisabledEvent struct {
	Data EndpointDisabledEventData `json:"data"`
	Type string `json:"type"`
}

type _EndpointDisabledEvent EndpointDisabledEvent

// NewEndpointDisabledEvent instantiates a new EndpointDisabledEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointDisabledEvent(data EndpointDisabledEventData, type_ string) *EndpointDisabledEvent {
	this := EndpointDisabledEvent{}
	this.Data = data
	this.Type = type_
	return &this
}

// NewEndpointDisabledEventWithDefaults instantiates a new EndpointDisabledEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointDisabledEventWithDefaults() *EndpointDisabledEvent {
	this := EndpointDisabledEvent{}
	var type_ string = "endpoint.disabled"
	this.Type = type_
	return &this
}

// GetData returns the Data field value
func (o *EndpointDisabledEvent) GetData() EndpointDisabledEventData {
	if o == nil {
		var ret EndpointDisabledEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EndpointDisabledEvent) GetDataOk() (*EndpointDisabledEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *EndpointDisabledEvent) SetData(v EndpointDisabledEventData) {
	o.Data = v
}

// GetType returns the Type field value
func (o *EndpointDisabledEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EndpointDisabledEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EndpointDisabledEvent) SetType(v string) {
	o.Type = v
}

func (o EndpointDisabledEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointDisabledEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *EndpointDisabledEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varEndpointDisabledEvent := _EndpointDisabledEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointDisabledEvent)

	if err != nil {
		return err
	}

	*o = EndpointDisabledEvent(varEndpointDisabledEvent)

	return err
}

type NullableEndpointDisabledEvent struct {
	value *EndpointDisabledEvent
	isSet bool
}

func (v NullableEndpointDisabledEvent) Get() *EndpointDisabledEvent {
	return v.value
}

func (v *NullableEndpointDisabledEvent) Set(val *EndpointDisabledEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointDisabledEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointDisabledEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointDisabledEvent(val *EndpointDisabledEvent) *NullableEndpointDisabledEvent {
	return &NullableEndpointDisabledEvent{value: val, isSet: true}
}

func (v NullableEndpointDisabledEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointDisabledEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


