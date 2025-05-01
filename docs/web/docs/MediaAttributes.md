# MediaAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** | file format | 
**Size** | **int64** | file size in bytes | 
**Url** | **string** | media URL | 
**CreatedAt** | **time.Time** | media creation date and time | 

## Methods

### NewMediaAttributes

`func NewMediaAttributes(format string, size int64, url string, createdAt time.Time, ) *MediaAttributes`

NewMediaAttributes instantiates a new MediaAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaAttributesWithDefaults

`func NewMediaAttributesWithDefaults() *MediaAttributes`

NewMediaAttributesWithDefaults instantiates a new MediaAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *MediaAttributes) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MediaAttributes) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MediaAttributes) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetSize

`func (o *MediaAttributes) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MediaAttributes) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MediaAttributes) SetSize(v int64)`

SetSize sets Size field to given value.


### GetUrl

`func (o *MediaAttributes) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MediaAttributes) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MediaAttributes) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCreatedAt

`func (o *MediaAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MediaAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MediaAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


