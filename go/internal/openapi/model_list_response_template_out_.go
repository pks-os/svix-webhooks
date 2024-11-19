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

// checks if the ListResponseTemplateOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResponseTemplateOut{}

// ListResponseTemplateOut struct for ListResponseTemplateOut
type ListResponseTemplateOut struct {
	Data []TemplateOut `json:"data"`
	Done bool `json:"done"`
	Iterator NullableString `json:"iterator"`
	PrevIterator NullableString `json:"prevIterator,omitempty"`
}

type _ListResponseTemplateOut ListResponseTemplateOut

// NewListResponseTemplateOut instantiates a new ListResponseTemplateOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResponseTemplateOut(data []TemplateOut, done bool, iterator NullableString) *ListResponseTemplateOut {
	this := ListResponseTemplateOut{}
	this.Data = data
	this.Done = done
	this.Iterator = iterator
	return &this
}

// NewListResponseTemplateOutWithDefaults instantiates a new ListResponseTemplateOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResponseTemplateOutWithDefaults() *ListResponseTemplateOut {
	this := ListResponseTemplateOut{}
	return &this
}

// GetData returns the Data field value
func (o *ListResponseTemplateOut) GetData() []TemplateOut {
	if o == nil {
		var ret []TemplateOut
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListResponseTemplateOut) GetDataOk() ([]TemplateOut, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListResponseTemplateOut) SetData(v []TemplateOut) {
	o.Data = v
}

// GetDone returns the Done field value
func (o *ListResponseTemplateOut) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *ListResponseTemplateOut) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *ListResponseTemplateOut) SetDone(v bool) {
	o.Done = v
}

// GetIterator returns the Iterator field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListResponseTemplateOut) GetIterator() string {
	if o == nil || o.Iterator.Get() == nil {
		var ret string
		return ret
	}

	return *o.Iterator.Get()
}

// GetIteratorOk returns a tuple with the Iterator field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseTemplateOut) GetIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Iterator.Get(), o.Iterator.IsSet()
}

// SetIterator sets field value
func (o *ListResponseTemplateOut) SetIterator(v string) {
	o.Iterator.Set(&v)
}

// GetPrevIterator returns the PrevIterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListResponseTemplateOut) GetPrevIterator() string {
	if o == nil || IsNil(o.PrevIterator.Get()) {
		var ret string
		return ret
	}
	return *o.PrevIterator.Get()
}

// GetPrevIteratorOk returns a tuple with the PrevIterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseTemplateOut) GetPrevIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevIterator.Get(), o.PrevIterator.IsSet()
}

// HasPrevIterator returns a boolean if a field has been set.
func (o *ListResponseTemplateOut) HasPrevIterator() bool {
	if o != nil && o.PrevIterator.IsSet() {
		return true
	}

	return false
}

// SetPrevIterator gets a reference to the given NullableString and assigns it to the PrevIterator field.
func (o *ListResponseTemplateOut) SetPrevIterator(v string) {
	o.PrevIterator.Set(&v)
}
// SetPrevIteratorNil sets the value for PrevIterator to be an explicit nil
func (o *ListResponseTemplateOut) SetPrevIteratorNil() {
	o.PrevIterator.Set(nil)
}

// UnsetPrevIterator ensures that no value is present for PrevIterator, not even an explicit nil
func (o *ListResponseTemplateOut) UnsetPrevIterator() {
	o.PrevIterator.Unset()
}

func (o ListResponseTemplateOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResponseTemplateOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["done"] = o.Done
	toSerialize["iterator"] = o.Iterator.Get()
	if o.PrevIterator.IsSet() {
		toSerialize["prevIterator"] = o.PrevIterator.Get()
	}
	return toSerialize, nil
}

func (o *ListResponseTemplateOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"done",
		"iterator",
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

	varListResponseTemplateOut := _ListResponseTemplateOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListResponseTemplateOut)

	if err != nil {
		return err
	}

	*o = ListResponseTemplateOut(varListResponseTemplateOut)

	return err
}

type NullableListResponseTemplateOut struct {
	value *ListResponseTemplateOut
	isSet bool
}

func (v NullableListResponseTemplateOut) Get() *ListResponseTemplateOut {
	return v.value
}

func (v *NullableListResponseTemplateOut) Set(val *ListResponseTemplateOut) {
	v.value = val
	v.isSet = true
}

func (v NullableListResponseTemplateOut) IsSet() bool {
	return v.isSet
}

func (v *NullableListResponseTemplateOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResponseTemplateOut(val *ListResponseTemplateOut) *NullableListResponseTemplateOut {
	return &NullableListResponseTemplateOut{value: val, isSet: true}
}

func (v NullableListResponseTemplateOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResponseTemplateOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


