/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RetryScheduleInOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetryScheduleInOut{}

// RetryScheduleInOut struct for RetryScheduleInOut
type RetryScheduleInOut struct {
	RetrySchedule []Duration `json:"retrySchedule,omitempty"`
}

// NewRetryScheduleInOut instantiates a new RetryScheduleInOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetryScheduleInOut() *RetryScheduleInOut {
	this := RetryScheduleInOut{}
	return &this
}

// NewRetryScheduleInOutWithDefaults instantiates a new RetryScheduleInOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetryScheduleInOutWithDefaults() *RetryScheduleInOut {
	this := RetryScheduleInOut{}
	return &this
}

// GetRetrySchedule returns the RetrySchedule field value if set, zero value otherwise.
func (o *RetryScheduleInOut) GetRetrySchedule() []Duration {
	if o == nil || IsNil(o.RetrySchedule) {
		var ret []Duration
		return ret
	}
	return o.RetrySchedule
}

// GetRetryScheduleOk returns a tuple with the RetrySchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetryScheduleInOut) GetRetryScheduleOk() ([]Duration, bool) {
	if o == nil || IsNil(o.RetrySchedule) {
		return nil, false
	}
	return o.RetrySchedule, true
}

// HasRetrySchedule returns a boolean if a field has been set.
func (o *RetryScheduleInOut) HasRetrySchedule() bool {
	if o != nil && !IsNil(o.RetrySchedule) {
		return true
	}

	return false
}

// SetRetrySchedule gets a reference to the given []Duration and assigns it to the RetrySchedule field.
func (o *RetryScheduleInOut) SetRetrySchedule(v []Duration) {
	o.RetrySchedule = v
}

func (o RetryScheduleInOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetryScheduleInOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RetrySchedule) {
		toSerialize["retrySchedule"] = o.RetrySchedule
	}
	return toSerialize, nil
}

type NullableRetryScheduleInOut struct {
	value *RetryScheduleInOut
	isSet bool
}

func (v NullableRetryScheduleInOut) Get() *RetryScheduleInOut {
	return v.value
}

func (v *NullableRetryScheduleInOut) Set(val *RetryScheduleInOut) {
	v.value = val
	v.isSet = true
}

func (v NullableRetryScheduleInOut) IsSet() bool {
	return v.isSet
}

func (v *NullableRetryScheduleInOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetryScheduleInOut(val *RetryScheduleInOut) *NullableRetryScheduleInOut {
	return &NullableRetryScheduleInOut{value: val, isSet: true}
}

func (v NullableRetryScheduleInOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetryScheduleInOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


