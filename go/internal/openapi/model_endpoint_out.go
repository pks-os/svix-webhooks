/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the EndpointOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointOut{}

// EndpointOut struct for EndpointOut
type EndpointOut struct {
	// List of message channels this endpoint listens to (omit for all).
	Channels []string `json:"channels,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	// An example endpoint name.
	Description string `json:"description"`
	Disabled *bool `json:"disabled,omitempty"`
	FilterTypes []string `json:"filterTypes,omitempty"`
	// The ep's ID
	Id string `json:"id"`
	Metadata map[string]string `json:"metadata"`
	RateLimit *int32 `json:"rateLimit,omitempty"`
	// Optional unique identifier for the endpoint.
	Uid *string `json:"uid,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	UpdatedAt time.Time `json:"updatedAt"`
	Url string `json:"url"`
	// Deprecated
	Version int32 `json:"version"`
}

type _EndpointOut EndpointOut

// NewEndpointOut instantiates a new EndpointOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointOut(createdAt time.Time, description string, id string, metadata map[string]string, updatedAt time.Time, url string, version int32) *EndpointOut {
	this := EndpointOut{}
	this.CreatedAt = createdAt
	this.Description = description
	var disabled bool = false
	this.Disabled = &disabled
	this.Id = id
	this.Metadata = metadata
	this.UpdatedAt = updatedAt
	this.Url = url
	this.Version = version
	return &this
}

// NewEndpointOutWithDefaults instantiates a new EndpointOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointOutWithDefaults() *EndpointOut {
	this := EndpointOut{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// GetChannels returns the Channels field value if set, zero value otherwise.
func (o *EndpointOut) GetChannels() []string {
	if o == nil || IsNil(o.Channels) {
		var ret []string
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointOut) GetChannelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *EndpointOut) HasChannels() bool {
	if o != nil && !IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []string and assigns it to the Channels field.
func (o *EndpointOut) SetChannels(v []string) {
	o.Channels = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EndpointOut) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EndpointOut) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EndpointOut) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value
func (o *EndpointOut) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EndpointOut) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *EndpointOut) SetDescription(v string) {
	o.Description = v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *EndpointOut) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointOut) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *EndpointOut) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *EndpointOut) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetFilterTypes returns the FilterTypes field value if set, zero value otherwise.
func (o *EndpointOut) GetFilterTypes() []string {
	if o == nil || IsNil(o.FilterTypes) {
		var ret []string
		return ret
	}
	return o.FilterTypes
}

// GetFilterTypesOk returns a tuple with the FilterTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointOut) GetFilterTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.FilterTypes) {
		return nil, false
	}
	return o.FilterTypes, true
}

// HasFilterTypes returns a boolean if a field has been set.
func (o *EndpointOut) HasFilterTypes() bool {
	if o != nil && !IsNil(o.FilterTypes) {
		return true
	}

	return false
}

// SetFilterTypes gets a reference to the given []string and assigns it to the FilterTypes field.
func (o *EndpointOut) SetFilterTypes(v []string) {
	o.FilterTypes = v
}

// GetId returns the Id field value
func (o *EndpointOut) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EndpointOut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EndpointOut) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value
func (o *EndpointOut) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *EndpointOut) GetMetadataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *EndpointOut) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise.
func (o *EndpointOut) GetRateLimit() int32 {
	if o == nil || IsNil(o.RateLimit) {
		var ret int32
		return ret
	}
	return *o.RateLimit
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointOut) GetRateLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.RateLimit) {
		return nil, false
	}
	return o.RateLimit, true
}

// HasRateLimit returns a boolean if a field has been set.
func (o *EndpointOut) HasRateLimit() bool {
	if o != nil && !IsNil(o.RateLimit) {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given int32 and assigns it to the RateLimit field.
func (o *EndpointOut) SetRateLimit(v int32) {
	o.RateLimit = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *EndpointOut) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointOut) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *EndpointOut) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *EndpointOut) SetUid(v string) {
	o.Uid = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *EndpointOut) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EndpointOut) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *EndpointOut) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUrl returns the Url field value
func (o *EndpointOut) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *EndpointOut) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *EndpointOut) SetUrl(v string) {
	o.Url = v
}

// GetVersion returns the Version field value
// Deprecated
func (o *EndpointOut) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *EndpointOut) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
// Deprecated
func (o *EndpointOut) SetVersion(v int32) {
	o.Version = v
}

func (o EndpointOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channels) {
		toSerialize["channels"] = o.Channels
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["description"] = o.Description
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.FilterTypes) {
		toSerialize["filterTypes"] = o.FilterTypes
	}
	toSerialize["id"] = o.Id
	toSerialize["metadata"] = o.Metadata
	if !IsNil(o.RateLimit) {
		toSerialize["rateLimit"] = o.RateLimit
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["url"] = o.Url
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

func (o *EndpointOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"description",
		"id",
		"metadata",
		"updatedAt",
		"url",
		"version",
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

	varEndpointOut := _EndpointOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointOut)

	if err != nil {
		return err
	}

	*o = EndpointOut(varEndpointOut)

	return err
}

type NullableEndpointOut struct {
	value *EndpointOut
	isSet bool
}

func (v NullableEndpointOut) Get() *EndpointOut {
	return v.value
}

func (v *NullableEndpointOut) Set(val *EndpointOut) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointOut) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointOut(val *EndpointOut) *NullableEndpointOut {
	return &NullableEndpointOut{value: val, isSet: true}
}

func (v NullableEndpointOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


