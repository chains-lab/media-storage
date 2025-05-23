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

// checks if the CreateMediaRulesDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMediaRulesDataAttributes{}

// CreateMediaRulesDataAttributes struct for CreateMediaRulesDataAttributes
type CreateMediaRulesDataAttributes struct {
	Extensions []string `json:"extensions"`
	// Maximum size of the media in bytes
	MaxSize int64 `json:"max_size"`
	Roles []string `json:"roles"`
}

type _CreateMediaRulesDataAttributes CreateMediaRulesDataAttributes

// NewCreateMediaRulesDataAttributes instantiates a new CreateMediaRulesDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMediaRulesDataAttributes(extensions []string, maxSize int64, roles []string) *CreateMediaRulesDataAttributes {
	this := CreateMediaRulesDataAttributes{}
	this.Extensions = extensions
	this.MaxSize = maxSize
	this.Roles = roles
	return &this
}

// NewCreateMediaRulesDataAttributesWithDefaults instantiates a new CreateMediaRulesDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMediaRulesDataAttributesWithDefaults() *CreateMediaRulesDataAttributes {
	this := CreateMediaRulesDataAttributes{}
	return &this
}

// GetExtensions returns the Extensions field value
func (o *CreateMediaRulesDataAttributes) GetExtensions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value
// and a boolean to check if the value has been set.
func (o *CreateMediaRulesDataAttributes) GetExtensionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Extensions, true
}

// SetExtensions sets field value
func (o *CreateMediaRulesDataAttributes) SetExtensions(v []string) {
	o.Extensions = v
}

// GetMaxSize returns the MaxSize field value
func (o *CreateMediaRulesDataAttributes) GetMaxSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value
// and a boolean to check if the value has been set.
func (o *CreateMediaRulesDataAttributes) GetMaxSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxSize, true
}

// SetMaxSize sets field value
func (o *CreateMediaRulesDataAttributes) SetMaxSize(v int64) {
	o.MaxSize = v
}

// GetRoles returns the Roles field value
func (o *CreateMediaRulesDataAttributes) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *CreateMediaRulesDataAttributes) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *CreateMediaRulesDataAttributes) SetRoles(v []string) {
	o.Roles = v
}

func (o CreateMediaRulesDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMediaRulesDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["extensions"] = o.Extensions
	toSerialize["max_size"] = o.MaxSize
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

func (o *CreateMediaRulesDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"extensions",
		"max_size",
		"roles",
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

	varCreateMediaRulesDataAttributes := _CreateMediaRulesDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateMediaRulesDataAttributes)

	if err != nil {
		return err
	}

	*o = CreateMediaRulesDataAttributes(varCreateMediaRulesDataAttributes)

	return err
}

type NullableCreateMediaRulesDataAttributes struct {
	value *CreateMediaRulesDataAttributes
	isSet bool
}

func (v NullableCreateMediaRulesDataAttributes) Get() *CreateMediaRulesDataAttributes {
	return v.value
}

func (v *NullableCreateMediaRulesDataAttributes) Set(val *CreateMediaRulesDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMediaRulesDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMediaRulesDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMediaRulesDataAttributes(val *CreateMediaRulesDataAttributes) *NullableCreateMediaRulesDataAttributes {
	return &NullableCreateMediaRulesDataAttributes{value: val, isSet: true}
}

func (v NullableCreateMediaRulesDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMediaRulesDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


