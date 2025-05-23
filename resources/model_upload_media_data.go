/*
Title

Title

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UploadMediaData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadMediaData{}

// UploadMediaData struct for UploadMediaData
type UploadMediaData struct {
	Type string `json:"type"`
	Attributes UploadMediaDataAttributes `json:"attributes"`
}

type _UploadMediaData UploadMediaData

// NewUploadMediaData instantiates a new UploadMediaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadMediaData(type_ string, attributes UploadMediaDataAttributes) *UploadMediaData {
	this := UploadMediaData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewUploadMediaDataWithDefaults instantiates a new UploadMediaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadMediaDataWithDefaults() *UploadMediaData {
	this := UploadMediaData{}
	return &this
}

// GetType returns the Type field value
func (o *UploadMediaData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UploadMediaData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UploadMediaData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *UploadMediaData) GetAttributes() UploadMediaDataAttributes {
	if o == nil {
		var ret UploadMediaDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *UploadMediaData) GetAttributesOk() (*UploadMediaDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *UploadMediaData) SetAttributes(v UploadMediaDataAttributes) {
	o.Attributes = v
}

func (o UploadMediaData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadMediaData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *UploadMediaData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
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

	varUploadMediaData := _UploadMediaData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUploadMediaData)

	if err != nil {
		return err
	}

	*o = UploadMediaData(varUploadMediaData)

	return err
}

type NullableUploadMediaData struct {
	value *UploadMediaData
	isSet bool
}

func (v NullableUploadMediaData) Get() *UploadMediaData {
	return v.value
}

func (v *NullableUploadMediaData) Set(val *UploadMediaData) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadMediaData) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadMediaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadMediaData(val *UploadMediaData) *NullableUploadMediaData {
	return &NullableUploadMediaData{value: val, isSet: true}
}

func (v NullableUploadMediaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadMediaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


