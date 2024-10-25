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

// checks if the SinkHttpConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SinkHttpConfig{}

// SinkHttpConfig struct for SinkHttpConfig
type SinkHttpConfig struct {
	Headers *map[string]string `json:"headers,omitempty"`
	Url string `json:"url"`
}

type _SinkHttpConfig SinkHttpConfig

// NewSinkHttpConfig instantiates a new SinkHttpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSinkHttpConfig(url string) *SinkHttpConfig {
	this := SinkHttpConfig{}
	this.Url = url
	return &this
}

// NewSinkHttpConfigWithDefaults instantiates a new SinkHttpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSinkHttpConfigWithDefaults() *SinkHttpConfig {
	this := SinkHttpConfig{}
	return &this
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *SinkHttpConfig) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SinkHttpConfig) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *SinkHttpConfig) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *SinkHttpConfig) SetHeaders(v map[string]string) {
	o.Headers = &v
}

// GetUrl returns the Url field value
func (o *SinkHttpConfig) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SinkHttpConfig) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SinkHttpConfig) SetUrl(v string) {
	o.Url = v
}

func (o SinkHttpConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SinkHttpConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *SinkHttpConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varSinkHttpConfig := _SinkHttpConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSinkHttpConfig)

	if err != nil {
		return err
	}

	*o = SinkHttpConfig(varSinkHttpConfig)

	return err
}

type NullableSinkHttpConfig struct {
	value *SinkHttpConfig
	isSet bool
}

func (v NullableSinkHttpConfig) Get() *SinkHttpConfig {
	return v.value
}

func (v *NullableSinkHttpConfig) Set(val *SinkHttpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSinkHttpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSinkHttpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSinkHttpConfig(val *SinkHttpConfig) *NullableSinkHttpConfig {
	return &NullableSinkHttpConfig{value: val, isSet: true}
}

func (v NullableSinkHttpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSinkHttpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


