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

// checks if the ListResponseMessageAttemptOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResponseMessageAttemptOut{}

// ListResponseMessageAttemptOut struct for ListResponseMessageAttemptOut
type ListResponseMessageAttemptOut struct {
	Data []MessageAttemptOut `json:"data"`
	Done bool `json:"done"`
	Iterator NullableString `json:"iterator,omitempty"`
	PrevIterator NullableString `json:"prevIterator,omitempty"`
}

type _ListResponseMessageAttemptOut ListResponseMessageAttemptOut

// NewListResponseMessageAttemptOut instantiates a new ListResponseMessageAttemptOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResponseMessageAttemptOut(data []MessageAttemptOut, done bool) *ListResponseMessageAttemptOut {
	this := ListResponseMessageAttemptOut{}
	this.Data = data
	this.Done = done
	return &this
}

// NewListResponseMessageAttemptOutWithDefaults instantiates a new ListResponseMessageAttemptOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResponseMessageAttemptOutWithDefaults() *ListResponseMessageAttemptOut {
	this := ListResponseMessageAttemptOut{}
	return &this
}

// GetData returns the Data field value
func (o *ListResponseMessageAttemptOut) GetData() []MessageAttemptOut {
	if o == nil {
		var ret []MessageAttemptOut
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListResponseMessageAttemptOut) GetDataOk() ([]MessageAttemptOut, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListResponseMessageAttemptOut) SetData(v []MessageAttemptOut) {
	o.Data = v
}

// GetDone returns the Done field value
func (o *ListResponseMessageAttemptOut) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *ListResponseMessageAttemptOut) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *ListResponseMessageAttemptOut) SetDone(v bool) {
	o.Done = v
}

// GetIterator returns the Iterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListResponseMessageAttemptOut) GetIterator() string {
	if o == nil || IsNil(o.Iterator.Get()) {
		var ret string
		return ret
	}
	return *o.Iterator.Get()
}

// GetIteratorOk returns a tuple with the Iterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseMessageAttemptOut) GetIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Iterator.Get(), o.Iterator.IsSet()
}

// HasIterator returns a boolean if a field has been set.
func (o *ListResponseMessageAttemptOut) HasIterator() bool {
	if o != nil && o.Iterator.IsSet() {
		return true
	}

	return false
}

// SetIterator gets a reference to the given NullableString and assigns it to the Iterator field.
func (o *ListResponseMessageAttemptOut) SetIterator(v string) {
	o.Iterator.Set(&v)
}
// SetIteratorNil sets the value for Iterator to be an explicit nil
func (o *ListResponseMessageAttemptOut) SetIteratorNil() {
	o.Iterator.Set(nil)
}

// UnsetIterator ensures that no value is present for Iterator, not even an explicit nil
func (o *ListResponseMessageAttemptOut) UnsetIterator() {
	o.Iterator.Unset()
}

// GetPrevIterator returns the PrevIterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListResponseMessageAttemptOut) GetPrevIterator() string {
	if o == nil || IsNil(o.PrevIterator.Get()) {
		var ret string
		return ret
	}
	return *o.PrevIterator.Get()
}

// GetPrevIteratorOk returns a tuple with the PrevIterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseMessageAttemptOut) GetPrevIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevIterator.Get(), o.PrevIterator.IsSet()
}

// HasPrevIterator returns a boolean if a field has been set.
func (o *ListResponseMessageAttemptOut) HasPrevIterator() bool {
	if o != nil && o.PrevIterator.IsSet() {
		return true
	}

	return false
}

// SetPrevIterator gets a reference to the given NullableString and assigns it to the PrevIterator field.
func (o *ListResponseMessageAttemptOut) SetPrevIterator(v string) {
	o.PrevIterator.Set(&v)
}
// SetPrevIteratorNil sets the value for PrevIterator to be an explicit nil
func (o *ListResponseMessageAttemptOut) SetPrevIteratorNil() {
	o.PrevIterator.Set(nil)
}

// UnsetPrevIterator ensures that no value is present for PrevIterator, not even an explicit nil
func (o *ListResponseMessageAttemptOut) UnsetPrevIterator() {
	o.PrevIterator.Unset()
}

func (o ListResponseMessageAttemptOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResponseMessageAttemptOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["done"] = o.Done
	if o.Iterator.IsSet() {
		toSerialize["iterator"] = o.Iterator.Get()
	}
	if o.PrevIterator.IsSet() {
		toSerialize["prevIterator"] = o.PrevIterator.Get()
	}
	return toSerialize, nil
}

func (o *ListResponseMessageAttemptOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"done",
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

	varListResponseMessageAttemptOut := _ListResponseMessageAttemptOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListResponseMessageAttemptOut)

	if err != nil {
		return err
	}

	*o = ListResponseMessageAttemptOut(varListResponseMessageAttemptOut)

	return err
}

type NullableListResponseMessageAttemptOut struct {
	value *ListResponseMessageAttemptOut
	isSet bool
}

func (v NullableListResponseMessageAttemptOut) Get() *ListResponseMessageAttemptOut {
	return v.value
}

func (v *NullableListResponseMessageAttemptOut) Set(val *ListResponseMessageAttemptOut) {
	v.value = val
	v.isSet = true
}

func (v NullableListResponseMessageAttemptOut) IsSet() bool {
	return v.isSet
}

func (v *NullableListResponseMessageAttemptOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResponseMessageAttemptOut(val *ListResponseMessageAttemptOut) *NullableListResponseMessageAttemptOut {
	return &NullableListResponseMessageAttemptOut{value: val, isSet: true}
}

func (v NullableListResponseMessageAttemptOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResponseMessageAttemptOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


