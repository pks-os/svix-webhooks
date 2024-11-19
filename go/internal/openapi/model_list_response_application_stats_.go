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

// checks if the ListResponseApplicationStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListResponseApplicationStats{}

// ListResponseApplicationStats struct for ListResponseApplicationStats
type ListResponseApplicationStats struct {
	Data []ApplicationStats `json:"data"`
	Done bool `json:"done"`
	Iterator NullableString `json:"iterator"`
	PrevIterator NullableString `json:"prevIterator,omitempty"`
}

type _ListResponseApplicationStats ListResponseApplicationStats

// NewListResponseApplicationStats instantiates a new ListResponseApplicationStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResponseApplicationStats(data []ApplicationStats, done bool, iterator NullableString) *ListResponseApplicationStats {
	this := ListResponseApplicationStats{}
	this.Data = data
	this.Done = done
	this.Iterator = iterator
	return &this
}

// NewListResponseApplicationStatsWithDefaults instantiates a new ListResponseApplicationStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResponseApplicationStatsWithDefaults() *ListResponseApplicationStats {
	this := ListResponseApplicationStats{}
	return &this
}

// GetData returns the Data field value
func (o *ListResponseApplicationStats) GetData() []ApplicationStats {
	if o == nil {
		var ret []ApplicationStats
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListResponseApplicationStats) GetDataOk() ([]ApplicationStats, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListResponseApplicationStats) SetData(v []ApplicationStats) {
	o.Data = v
}

// GetDone returns the Done field value
func (o *ListResponseApplicationStats) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *ListResponseApplicationStats) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *ListResponseApplicationStats) SetDone(v bool) {
	o.Done = v
}

// GetIterator returns the Iterator field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListResponseApplicationStats) GetIterator() string {
	if o == nil || o.Iterator.Get() == nil {
		var ret string
		return ret
	}

	return *o.Iterator.Get()
}

// GetIteratorOk returns a tuple with the Iterator field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseApplicationStats) GetIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Iterator.Get(), o.Iterator.IsSet()
}

// SetIterator sets field value
func (o *ListResponseApplicationStats) SetIterator(v string) {
	o.Iterator.Set(&v)
}

// GetPrevIterator returns the PrevIterator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListResponseApplicationStats) GetPrevIterator() string {
	if o == nil || IsNil(o.PrevIterator.Get()) {
		var ret string
		return ret
	}
	return *o.PrevIterator.Get()
}

// GetPrevIteratorOk returns a tuple with the PrevIterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListResponseApplicationStats) GetPrevIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevIterator.Get(), o.PrevIterator.IsSet()
}

// HasPrevIterator returns a boolean if a field has been set.
func (o *ListResponseApplicationStats) HasPrevIterator() bool {
	if o != nil && o.PrevIterator.IsSet() {
		return true
	}

	return false
}

// SetPrevIterator gets a reference to the given NullableString and assigns it to the PrevIterator field.
func (o *ListResponseApplicationStats) SetPrevIterator(v string) {
	o.PrevIterator.Set(&v)
}
// SetPrevIteratorNil sets the value for PrevIterator to be an explicit nil
func (o *ListResponseApplicationStats) SetPrevIteratorNil() {
	o.PrevIterator.Set(nil)
}

// UnsetPrevIterator ensures that no value is present for PrevIterator, not even an explicit nil
func (o *ListResponseApplicationStats) UnsetPrevIterator() {
	o.PrevIterator.Unset()
}

func (o ListResponseApplicationStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListResponseApplicationStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["done"] = o.Done
	toSerialize["iterator"] = o.Iterator.Get()
	if o.PrevIterator.IsSet() {
		toSerialize["prevIterator"] = o.PrevIterator.Get()
	}
	return toSerialize, nil
}

func (o *ListResponseApplicationStats) UnmarshalJSON(data []byte) (err error) {
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

	varListResponseApplicationStats := _ListResponseApplicationStats{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListResponseApplicationStats)

	if err != nil {
		return err
	}

	*o = ListResponseApplicationStats(varListResponseApplicationStats)

	return err
}

type NullableListResponseApplicationStats struct {
	value *ListResponseApplicationStats
	isSet bool
}

func (v NullableListResponseApplicationStats) Get() *ListResponseApplicationStats {
	return v.value
}

func (v *NullableListResponseApplicationStats) Set(val *ListResponseApplicationStats) {
	v.value = val
	v.isSet = true
}

func (v NullableListResponseApplicationStats) IsSet() bool {
	return v.isSet
}

func (v *NullableListResponseApplicationStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResponseApplicationStats(val *ListResponseApplicationStats) *NullableListResponseApplicationStats {
	return &NullableListResponseApplicationStats{value: val, isSet: true}
}

func (v NullableListResponseApplicationStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResponseApplicationStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


