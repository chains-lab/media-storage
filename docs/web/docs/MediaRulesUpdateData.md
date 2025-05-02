# MediaRulesUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | media type | 
**Type** | **string** |  | 
**Attributes** | [**MediaRulesUpdateDataAttributes**](MediaRulesUpdateDataAttributes.md) |  | 

## Methods

### NewMediaRulesUpdateData

`func NewMediaRulesUpdateData(id string, type_ string, attributes MediaRulesUpdateDataAttributes, ) *MediaRulesUpdateData`

NewMediaRulesUpdateData instantiates a new MediaRulesUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaRulesUpdateDataWithDefaults

`func NewMediaRulesUpdateDataWithDefaults() *MediaRulesUpdateData`

NewMediaRulesUpdateDataWithDefaults instantiates a new MediaRulesUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MediaRulesUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaRulesUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaRulesUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *MediaRulesUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MediaRulesUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MediaRulesUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *MediaRulesUpdateData) GetAttributes() MediaRulesUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MediaRulesUpdateData) GetAttributesOk() (*MediaRulesUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MediaRulesUpdateData) SetAttributes(v MediaRulesUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


