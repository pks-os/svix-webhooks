/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// StreamSinkPatch - struct for StreamSinkPatch
type StreamSinkPatch struct {
	StreamSinkInOneOf *StreamSinkInOneOf
	StreamSinkInOneOf1 *StreamSinkInOneOf1
	StreamSinkInOneOf2 *StreamSinkInOneOf2
	StreamSinkInOneOf3 *StreamSinkInOneOf3
	StreamSinkInOneOf4 *StreamSinkInOneOf4
	StreamSinkInOneOf5 *StreamSinkInOneOf5
	StreamSinkInOneOf6 *StreamSinkInOneOf6
	StreamSinkInOneOf7 *StreamSinkInOneOf7
}

// StreamSinkInOneOfAsStreamSinkPatch is a convenience function that returns StreamSinkInOneOf wrapped in StreamSinkPatch
func StreamSinkInOneOfAsStreamSinkPatch(v *StreamSinkInOneOf) StreamSinkPatch {
	return StreamSinkPatch{
		StreamSinkInOneOf: v,
	}
}

// StreamSinkInOneOf1AsStreamSinkPatch is a convenience function that returns StreamSinkInOneOf1 wrapped in StreamSinkPatch
func StreamSinkInOneOf1AsStreamSinkPatch(v *StreamSinkInOneOf1) StreamSinkPatch {
	return StreamSinkPatch{
		StreamSinkInOneOf1: v,
	}
}

// StreamSinkInOneOf2AsStreamSinkPatch is a convenience function that returns StreamSinkInOneOf2 wrapped in StreamSinkPatch
func StreamSinkInOneOf2AsStreamSinkPatch(v *StreamSinkInOneOf2) StreamSinkPatch {
	return StreamSinkPatch{
		StreamSinkInOneOf2: v,
	}
}

// StreamSinkInOneOf3AsStreamSinkPatch is a convenience function that returns StreamSinkInOneOf3 wrapped in StreamSinkPatch
func StreamSinkInOneOf3AsStreamSinkPatch(v *StreamSinkInOneOf3) StreamSinkPatch {
	return StreamSinkPatch{
		StreamSinkInOneOf3: v,
	}
}

// StreamSinkInOneOf4AsStreamSinkPatch is a convenience function that returns StreamSinkInOneOf4 wrapped in StreamSinkPatch
func StreamSinkInOneOf4AsStreamSinkPatch(v *StreamSinkInOneOf4) StreamSinkPatch {
	return StreamSinkPatch{
		StreamSinkInOneOf4: v,
	}
}

// StreamSinkInOneOf5AsStreamSinkPatch is a convenience function that returns StreamSinkInOneOf5 wrapped in StreamSinkPatch
func StreamSinkInOneOf5AsStreamSinkPatch(v *StreamSinkInOneOf5) StreamSinkPatch {
	return StreamSinkPatch{
		StreamSinkInOneOf5: v,
	}
}

// StreamSinkInOneOf6AsStreamSinkPatch is a convenience function that returns StreamSinkInOneOf6 wrapped in StreamSinkPatch
func StreamSinkInOneOf6AsStreamSinkPatch(v *StreamSinkInOneOf6) StreamSinkPatch {
	return StreamSinkPatch{
		StreamSinkInOneOf6: v,
	}
}

// StreamSinkInOneOf7AsStreamSinkPatch is a convenience function that returns StreamSinkInOneOf7 wrapped in StreamSinkPatch
func StreamSinkInOneOf7AsStreamSinkPatch(v *StreamSinkInOneOf7) StreamSinkPatch {
	return StreamSinkPatch{
		StreamSinkInOneOf7: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *StreamSinkPatch) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into StreamSinkInOneOf
	err = newStrictDecoder(data).Decode(&dst.StreamSinkInOneOf)
	if err == nil {
		jsonStreamSinkInOneOf, _ := json.Marshal(dst.StreamSinkInOneOf)
		if string(jsonStreamSinkInOneOf) == "{}" { // empty struct
			dst.StreamSinkInOneOf = nil
		} else {
			if err = validator.Validate(dst.StreamSinkInOneOf); err != nil {
				dst.StreamSinkInOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamSinkInOneOf = nil
	}

	// try to unmarshal data into StreamSinkInOneOf1
	err = newStrictDecoder(data).Decode(&dst.StreamSinkInOneOf1)
	if err == nil {
		jsonStreamSinkInOneOf1, _ := json.Marshal(dst.StreamSinkInOneOf1)
		if string(jsonStreamSinkInOneOf1) == "{}" { // empty struct
			dst.StreamSinkInOneOf1 = nil
		} else {
			if err = validator.Validate(dst.StreamSinkInOneOf1); err != nil {
				dst.StreamSinkInOneOf1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamSinkInOneOf1 = nil
	}

	// try to unmarshal data into StreamSinkInOneOf2
	err = newStrictDecoder(data).Decode(&dst.StreamSinkInOneOf2)
	if err == nil {
		jsonStreamSinkInOneOf2, _ := json.Marshal(dst.StreamSinkInOneOf2)
		if string(jsonStreamSinkInOneOf2) == "{}" { // empty struct
			dst.StreamSinkInOneOf2 = nil
		} else {
			if err = validator.Validate(dst.StreamSinkInOneOf2); err != nil {
				dst.StreamSinkInOneOf2 = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamSinkInOneOf2 = nil
	}

	// try to unmarshal data into StreamSinkInOneOf3
	err = newStrictDecoder(data).Decode(&dst.StreamSinkInOneOf3)
	if err == nil {
		jsonStreamSinkInOneOf3, _ := json.Marshal(dst.StreamSinkInOneOf3)
		if string(jsonStreamSinkInOneOf3) == "{}" { // empty struct
			dst.StreamSinkInOneOf3 = nil
		} else {
			if err = validator.Validate(dst.StreamSinkInOneOf3); err != nil {
				dst.StreamSinkInOneOf3 = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamSinkInOneOf3 = nil
	}

	// try to unmarshal data into StreamSinkInOneOf4
	err = newStrictDecoder(data).Decode(&dst.StreamSinkInOneOf4)
	if err == nil {
		jsonStreamSinkInOneOf4, _ := json.Marshal(dst.StreamSinkInOneOf4)
		if string(jsonStreamSinkInOneOf4) == "{}" { // empty struct
			dst.StreamSinkInOneOf4 = nil
		} else {
			if err = validator.Validate(dst.StreamSinkInOneOf4); err != nil {
				dst.StreamSinkInOneOf4 = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamSinkInOneOf4 = nil
	}

	// try to unmarshal data into StreamSinkInOneOf5
	err = newStrictDecoder(data).Decode(&dst.StreamSinkInOneOf5)
	if err == nil {
		jsonStreamSinkInOneOf5, _ := json.Marshal(dst.StreamSinkInOneOf5)
		if string(jsonStreamSinkInOneOf5) == "{}" { // empty struct
			dst.StreamSinkInOneOf5 = nil
		} else {
			if err = validator.Validate(dst.StreamSinkInOneOf5); err != nil {
				dst.StreamSinkInOneOf5 = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamSinkInOneOf5 = nil
	}

	// try to unmarshal data into StreamSinkInOneOf6
	err = newStrictDecoder(data).Decode(&dst.StreamSinkInOneOf6)
	if err == nil {
		jsonStreamSinkInOneOf6, _ := json.Marshal(dst.StreamSinkInOneOf6)
		if string(jsonStreamSinkInOneOf6) == "{}" { // empty struct
			dst.StreamSinkInOneOf6 = nil
		} else {
			if err = validator.Validate(dst.StreamSinkInOneOf6); err != nil {
				dst.StreamSinkInOneOf6 = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamSinkInOneOf6 = nil
	}

	// try to unmarshal data into StreamSinkInOneOf7
	err = newStrictDecoder(data).Decode(&dst.StreamSinkInOneOf7)
	if err == nil {
		jsonStreamSinkInOneOf7, _ := json.Marshal(dst.StreamSinkInOneOf7)
		if string(jsonStreamSinkInOneOf7) == "{}" { // empty struct
			dst.StreamSinkInOneOf7 = nil
		} else {
			if err = validator.Validate(dst.StreamSinkInOneOf7); err != nil {
				dst.StreamSinkInOneOf7 = nil
			} else {
				match++
			}
		}
	} else {
		dst.StreamSinkInOneOf7 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.StreamSinkInOneOf = nil
		dst.StreamSinkInOneOf1 = nil
		dst.StreamSinkInOneOf2 = nil
		dst.StreamSinkInOneOf3 = nil
		dst.StreamSinkInOneOf4 = nil
		dst.StreamSinkInOneOf5 = nil
		dst.StreamSinkInOneOf6 = nil
		dst.StreamSinkInOneOf7 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(StreamSinkPatch)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(StreamSinkPatch)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StreamSinkPatch) MarshalJSON() ([]byte, error) {
	if src.StreamSinkInOneOf != nil {
		return json.Marshal(&src.StreamSinkInOneOf)
	}

	if src.StreamSinkInOneOf1 != nil {
		return json.Marshal(&src.StreamSinkInOneOf1)
	}

	if src.StreamSinkInOneOf2 != nil {
		return json.Marshal(&src.StreamSinkInOneOf2)
	}

	if src.StreamSinkInOneOf3 != nil {
		return json.Marshal(&src.StreamSinkInOneOf3)
	}

	if src.StreamSinkInOneOf4 != nil {
		return json.Marshal(&src.StreamSinkInOneOf4)
	}

	if src.StreamSinkInOneOf5 != nil {
		return json.Marshal(&src.StreamSinkInOneOf5)
	}

	if src.StreamSinkInOneOf6 != nil {
		return json.Marshal(&src.StreamSinkInOneOf6)
	}

	if src.StreamSinkInOneOf7 != nil {
		return json.Marshal(&src.StreamSinkInOneOf7)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StreamSinkPatch) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.StreamSinkInOneOf != nil {
		return obj.StreamSinkInOneOf
	}

	if obj.StreamSinkInOneOf1 != nil {
		return obj.StreamSinkInOneOf1
	}

	if obj.StreamSinkInOneOf2 != nil {
		return obj.StreamSinkInOneOf2
	}

	if obj.StreamSinkInOneOf3 != nil {
		return obj.StreamSinkInOneOf3
	}

	if obj.StreamSinkInOneOf4 != nil {
		return obj.StreamSinkInOneOf4
	}

	if obj.StreamSinkInOneOf5 != nil {
		return obj.StreamSinkInOneOf5
	}

	if obj.StreamSinkInOneOf6 != nil {
		return obj.StreamSinkInOneOf6
	}

	if obj.StreamSinkInOneOf7 != nil {
		return obj.StreamSinkInOneOf7
	}

	// all schemas are nil
	return nil
}

type NullableStreamSinkPatch struct {
	value *StreamSinkPatch
	isSet bool
}

func (v NullableStreamSinkPatch) Get() *StreamSinkPatch {
	return v.value
}

func (v *NullableStreamSinkPatch) Set(val *StreamSinkPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamSinkPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamSinkPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamSinkPatch(val *StreamSinkPatch) *NullableStreamSinkPatch {
	return &NullableStreamSinkPatch{value: val, isSet: true}
}

func (v NullableStreamSinkPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamSinkPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


