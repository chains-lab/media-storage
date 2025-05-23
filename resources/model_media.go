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

// checks if the Media type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Media{}

// Media struct for Media
type Media struct {
	Data MediaData `json:"data"`
}

type _Media Media

// NewMedia instantiates a new Media object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMedia(data MediaData) *Media {
	this := Media{}
	this.Data = data
	return &this
}

// NewMediaWithDefaults instantiates a new Media object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaWithDefaults() *Media {
	this := Media{}
	return &this
}

// GetData returns the Data field value
func (o *Media) GetData() MediaData {
	if o == nil {
		var ret MediaData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Media) GetDataOk() (*MediaData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Media) SetData(v MediaData) {
	o.Data = v
}

func (o Media) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Media) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *Media) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varMedia := _Media{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMedia)

	if err != nil {
		return err
	}

	*o = Media(varMedia)

	return err
}

type NullableMedia struct {
	value *Media
	isSet bool
}

func (v NullableMedia) Get() *Media {
	return v.value
}

func (v *NullableMedia) Set(val *Media) {
	v.value = val
	v.isSet = true
}

func (v NullableMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMedia(val *Media) *NullableMedia {
	return &NullableMedia{value: val, isSet: true}
}

func (v NullableMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


