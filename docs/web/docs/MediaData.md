# MediaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | file name | 
**Type** | **string** |  | 
**Attributes** | [**MediaAttributes**](MediaAttributes.md) |  | 

## Methods

### NewMediaData

`func NewMediaData(id string, type_ string, attributes MediaAttributes, ) *MediaData`

NewMediaData instantiates a new MediaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaDataWithDefaults

`func NewMediaDataWithDefaults() *MediaData`

NewMediaDataWithDefaults instantiates a new MediaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MediaData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *MediaData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MediaData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MediaData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *MediaData) GetAttributes() MediaAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MediaData) GetAttributesOk() (*MediaAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MediaData) SetAttributes(v MediaAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


