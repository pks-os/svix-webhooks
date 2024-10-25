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

// checks if the EventTypeImportOpenApiOutData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventTypeImportOpenApiOutData{}

// EventTypeImportOpenApiOutData struct for EventTypeImportOpenApiOutData
type EventTypeImportOpenApiOutData struct {
	Modified []string `json:"modified"`
	ToModify []EventTypeFromOpenApi `json:"to_modify,omitempty"`
}

type _EventTypeImportOpenApiOutData EventTypeImportOpenApiOutData

// NewEventTypeImportOpenApiOutData instantiates a new EventTypeImportOpenApiOutData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventTypeImportOpenApiOutData(modified []string) *EventTypeImportOpenApiOutData {
	this := EventTypeImportOpenApiOutData{}
	this.Modified = modified
	return &this
}

// NewEventTypeImportOpenApiOutDataWithDefaults instantiates a new EventTypeImportOpenApiOutData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventTypeImportOpenApiOutDataWithDefaults() *EventTypeImportOpenApiOutData {
	this := EventTypeImportOpenApiOutData{}
	return &this
}

// GetModified returns the Modified field value
func (o *EventTypeImportOpenApiOutData) GetModified() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *EventTypeImportOpenApiOutData) GetModifiedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified, true
}

// SetModified sets field value
func (o *EventTypeImportOpenApiOutData) SetModified(v []string) {
	o.Modified = v
}

// GetToModify returns the ToModify field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTypeImportOpenApiOutData) GetToModify() []EventTypeFromOpenApi {
	if o == nil {
		var ret []EventTypeFromOpenApi
		return ret
	}
	return o.ToModify
}

// GetToModifyOk returns a tuple with the ToModify field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventTypeImportOpenApiOutData) GetToModifyOk() ([]EventTypeFromOpenApi, bool) {
	if o == nil || IsNil(o.ToModify) {
		return nil, false
	}
	return o.ToModify, true
}

// HasToModify returns a boolean if a field has been set.
func (o *EventTypeImportOpenApiOutData) HasToModify() bool {
	if o != nil && !IsNil(o.ToModify) {
		return true
	}

	return false
}

// SetToModify gets a reference to the given []EventTypeFromOpenApi and assigns it to the ToModify field.
func (o *EventTypeImportOpenApiOutData) SetToModify(v []EventTypeFromOpenApi) {
	o.ToModify = v
}

func (o EventTypeImportOpenApiOutData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventTypeImportOpenApiOutData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["modified"] = o.Modified
	if o.ToModify != nil {
		toSerialize["to_modify"] = o.ToModify
	}
	return toSerialize, nil
}

func (o *EventTypeImportOpenApiOutData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"modified",
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

	varEventTypeImportOpenApiOutData := _EventTypeImportOpenApiOutData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventTypeImportOpenApiOutData)

	if err != nil {
		return err
	}

	*o = EventTypeImportOpenApiOutData(varEventTypeImportOpenApiOutData)

	return err
}

type NullableEventTypeImportOpenApiOutData struct {
	value *EventTypeImportOpenApiOutData
	isSet bool
}

func (v NullableEventTypeImportOpenApiOutData) Get() *EventTypeImportOpenApiOutData {
	return v.value
}

func (v *NullableEventTypeImportOpenApiOutData) Set(val *EventTypeImportOpenApiOutData) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeImportOpenApiOutData) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeImportOpenApiOutData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeImportOpenApiOutData(val *EventTypeImportOpenApiOutData) *NullableEventTypeImportOpenApiOutData {
	return &NullableEventTypeImportOpenApiOutData{value: val, isSet: true}
}

func (v NullableEventTypeImportOpenApiOutData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeImportOpenApiOutData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


