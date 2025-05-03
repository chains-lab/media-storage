# CreateMediaRulesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | media resource type | 
**Type** | **string** |  | 
**Attributes** | [**CreateMediaRulesDataAttributes**](CreateMediaRulesDataAttributes.md) |  | 

## Methods

### NewCreateMediaRulesData

`func NewCreateMediaRulesData(id string, type_ string, attributes CreateMediaRulesDataAttributes, ) *CreateMediaRulesData`

NewCreateMediaRulesData instantiates a new CreateMediaRulesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMediaRulesDataWithDefaults

`func NewCreateMediaRulesDataWithDefaults() *CreateMediaRulesData`

NewCreateMediaRulesDataWithDefaults instantiates a new CreateMediaRulesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateMediaRulesData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateMediaRulesData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateMediaRulesData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CreateMediaRulesData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateMediaRulesData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateMediaRulesData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *CreateMediaRulesData) GetAttributes() CreateMediaRulesDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreateMediaRulesData) GetAttributesOk() (*CreateMediaRulesDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreateMediaRulesData) SetAttributes(v CreateMediaRulesDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


