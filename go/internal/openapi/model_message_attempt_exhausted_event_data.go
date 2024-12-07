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

// checks if the MessageAttemptExhaustedEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageAttemptExhaustedEventData{}

// MessageAttemptExhaustedEventData Sent when a message delivery has failed (all of the retry attempts have been exhausted) as a \"message.attempt.exhausted\" type or after it's failed four times as a \"message.attempt.failing\" event.
type MessageAttemptExhaustedEventData struct {
	// The app's ID
	AppId string `json:"appId"`
	// The app's UID
	AppUid *string `json:"appUid,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	// The ep's ID
	EndpointId string `json:"endpointId"`
	LastAttempt MessageAttemptFailedData `json:"lastAttempt"`
	// The msg's UID
	MsgEventId *string `json:"msgEventId,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	// The msg's ID
	MsgId string `json:"msgId"`
}

type _MessageAttemptExhaustedEventData MessageAttemptExhaustedEventData

// NewMessageAttemptExhaustedEventData instantiates a new MessageAttemptExhaustedEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageAttemptExhaustedEventData(appId string, endpointId string, lastAttempt MessageAttemptFailedData, msgId string) *MessageAttemptExhaustedEventData {
	this := MessageAttemptExhaustedEventData{}
	this.AppId = appId
	this.EndpointId = endpointId
	this.LastAttempt = lastAttempt
	this.MsgId = msgId
	return &this
}

// NewMessageAttemptExhaustedEventDataWithDefaults instantiates a new MessageAttemptExhaustedEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageAttemptExhaustedEventDataWithDefaults() *MessageAttemptExhaustedEventData {
	this := MessageAttemptExhaustedEventData{}
	return &this
}

// GetAppId returns the AppId field value
func (o *MessageAttemptExhaustedEventData) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptExhaustedEventData) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *MessageAttemptExhaustedEventData) SetAppId(v string) {
	o.AppId = v
}

// GetAppUid returns the AppUid field value if set, zero value otherwise.
func (o *MessageAttemptExhaustedEventData) GetAppUid() string {
	if o == nil || IsNil(o.AppUid) {
		var ret string
		return ret
	}
	return *o.AppUid
}

// GetAppUidOk returns a tuple with the AppUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageAttemptExhaustedEventData) GetAppUidOk() (*string, bool) {
	if o == nil || IsNil(o.AppUid) {
		return nil, false
	}
	return o.AppUid, true
}

// HasAppUid returns a boolean if a field has been set.
func (o *MessageAttemptExhaustedEventData) HasAppUid() bool {
	if o != nil && !IsNil(o.AppUid) {
		return true
	}

	return false
}

// SetAppUid gets a reference to the given string and assigns it to the AppUid field.
func (o *MessageAttemptExhaustedEventData) SetAppUid(v string) {
	o.AppUid = &v
}

// GetEndpointId returns the EndpointId field value
func (o *MessageAttemptExhaustedEventData) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptExhaustedEventData) GetEndpointIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *MessageAttemptExhaustedEventData) SetEndpointId(v string) {
	o.EndpointId = v
}

// GetLastAttempt returns the LastAttempt field value
func (o *MessageAttemptExhaustedEventData) GetLastAttempt() MessageAttemptFailedData {
	if o == nil {
		var ret MessageAttemptFailedData
		return ret
	}

	return o.LastAttempt
}

// GetLastAttemptOk returns a tuple with the LastAttempt field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptExhaustedEventData) GetLastAttemptOk() (*MessageAttemptFailedData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastAttempt, true
}

// SetLastAttempt sets field value
func (o *MessageAttemptExhaustedEventData) SetLastAttempt(v MessageAttemptFailedData) {
	o.LastAttempt = v
}

// GetMsgEventId returns the MsgEventId field value if set, zero value otherwise.
func (o *MessageAttemptExhaustedEventData) GetMsgEventId() string {
	if o == nil || IsNil(o.MsgEventId) {
		var ret string
		return ret
	}
	return *o.MsgEventId
}

// GetMsgEventIdOk returns a tuple with the MsgEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageAttemptExhaustedEventData) GetMsgEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.MsgEventId) {
		return nil, false
	}
	return o.MsgEventId, true
}

// HasMsgEventId returns a boolean if a field has been set.
func (o *MessageAttemptExhaustedEventData) HasMsgEventId() bool {
	if o != nil && !IsNil(o.MsgEventId) {
		return true
	}

	return false
}

// SetMsgEventId gets a reference to the given string and assigns it to the MsgEventId field.
func (o *MessageAttemptExhaustedEventData) SetMsgEventId(v string) {
	o.MsgEventId = &v
}

// GetMsgId returns the MsgId field value
func (o *MessageAttemptExhaustedEventData) GetMsgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MsgId
}

// GetMsgIdOk returns a tuple with the MsgId field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptExhaustedEventData) GetMsgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MsgId, true
}

// SetMsgId sets field value
func (o *MessageAttemptExhaustedEventData) SetMsgId(v string) {
	o.MsgId = v
}

func (o MessageAttemptExhaustedEventData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageAttemptExhaustedEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appId"] = o.AppId
	if !IsNil(o.AppUid) {
		toSerialize["appUid"] = o.AppUid
	}
	toSerialize["endpointId"] = o.EndpointId
	toSerialize["lastAttempt"] = o.LastAttempt
	if !IsNil(o.MsgEventId) {
		toSerialize["msgEventId"] = o.MsgEventId
	}
	toSerialize["msgId"] = o.MsgId
	return toSerialize, nil
}

func (o *MessageAttemptExhaustedEventData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appId",
		"endpointId",
		"lastAttempt",
		"msgId",
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

	varMessageAttemptExhaustedEventData := _MessageAttemptExhaustedEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageAttemptExhaustedEventData)

	if err != nil {
		return err
	}

	*o = MessageAttemptExhaustedEventData(varMessageAttemptExhaustedEventData)

	return err
}

type NullableMessageAttemptExhaustedEventData struct {
	value *MessageAttemptExhaustedEventData
	isSet bool
}

func (v NullableMessageAttemptExhaustedEventData) Get() *MessageAttemptExhaustedEventData {
	return v.value
}

func (v *NullableMessageAttemptExhaustedEventData) Set(val *MessageAttemptExhaustedEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageAttemptExhaustedEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageAttemptExhaustedEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageAttemptExhaustedEventData(val *MessageAttemptExhaustedEventData) *NullableMessageAttemptExhaustedEventData {
	return &NullableMessageAttemptExhaustedEventData{value: val, isSet: true}
}

func (v NullableMessageAttemptExhaustedEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageAttemptExhaustedEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


