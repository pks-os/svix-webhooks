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

// checks if the EventIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventIn{}

// EventIn struct for EventIn
type EventIn struct {
	// The event type's name
	EventType NullableString `json:"eventType,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	Payload string `json:"payload"`
}

type _EventIn EventIn

// NewEventIn instantiates a new EventIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventIn(payload string) *EventIn {
	this := EventIn{}
	this.Payload = payload
	return &this
}

// NewEventInWithDefaults instantiates a new EventIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventInWithDefaults() *EventIn {
	this := EventIn{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventIn) GetEventType() string {
	if o == nil || IsNil(o.EventType.Get()) {
		var ret string
		return ret
	}
	return *o.EventType.Get()
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventIn) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventType.Get(), o.EventType.IsSet()
}

// HasEventType returns a boolean if a field has been set.
func (o *EventIn) HasEventType() bool {
	if o != nil && o.EventType.IsSet() {
		return true
	}

	return false
}

// SetEventType gets a reference to the given NullableString and assigns it to the EventType field.
func (o *EventIn) SetEventType(v string) {
	o.EventType.Set(&v)
}
// SetEventTypeNil sets the value for EventType to be an explicit nil
func (o *EventIn) SetEventTypeNil() {
	o.EventType.Set(nil)
}

// UnsetEventType ensures that no value is present for EventType, not even an explicit nil
func (o *EventIn) UnsetEventType() {
	o.EventType.Unset()
}

// GetPayload returns the Payload field value
func (o *EventIn) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *EventIn) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *EventIn) SetPayload(v string) {
	o.Payload = v
}

func (o EventIn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EventType.IsSet() {
		toSerialize["eventType"] = o.EventType.Get()
	}
	toSerialize["payload"] = o.Payload
	return toSerialize, nil
}

func (o *EventIn) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"payload",
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

	varEventIn := _EventIn{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventIn)

	if err != nil {
		return err
	}

	*o = EventIn(varEventIn)

	return err
}

type NullableEventIn struct {
	value *EventIn
	isSet bool
}

func (v NullableEventIn) Get() *EventIn {
	return v.value
}

func (v *NullableEventIn) Set(val *EventIn) {
	v.value = val
	v.isSet = true
}

func (v NullableEventIn) IsSet() bool {
	return v.isSet
}

func (v *NullableEventIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventIn(val *EventIn) *NullableEventIn {
	return &NullableEventIn{value: val, isSet: true}
}

func (v NullableEventIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


