# CreateMediaRulesDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSize** | **int64** | Maximum file size in bytes | 
**AllowedExits** | **[]string** |  | 
**Folder** | **string** | Folder where the media is stored | 
**Roles** | **[]string** |  | 

## Methods

### NewCreateMediaRulesDataAttributes

`func NewCreateMediaRulesDataAttributes(maxSize int64, allowedExits []string, folder string, roles []string, ) *CreateMediaRulesDataAttributes`

NewCreateMediaRulesDataAttributes instantiates a new CreateMediaRulesDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMediaRulesDataAttributesWithDefaults

`func NewCreateMediaRulesDataAttributesWithDefaults() *CreateMediaRulesDataAttributes`

NewCreateMediaRulesDataAttributesWithDefaults instantiates a new CreateMediaRulesDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSize

`func (o *CreateMediaRulesDataAttributes) GetMaxSize() int64`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *CreateMediaRulesDataAttributes) GetMaxSizeOk() (*int64, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *CreateMediaRulesDataAttributes) SetMaxSize(v int64)`

SetMaxSize sets MaxSize field to given value.


### GetAllowedExits

`func (o *CreateMediaRulesDataAttributes) GetAllowedExits() []string`

GetAllowedExits returns the AllowedExits field if non-nil, zero value otherwise.

### GetAllowedExitsOk

`func (o *CreateMediaRulesDataAttributes) GetAllowedExitsOk() (*[]string, bool)`

GetAllowedExitsOk returns a tuple with the AllowedExits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedExits

`func (o *CreateMediaRulesDataAttributes) SetAllowedExits(v []string)`

SetAllowedExits sets AllowedExits field to given value.


### GetFolder

`func (o *CreateMediaRulesDataAttributes) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *CreateMediaRulesDataAttributes) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *CreateMediaRulesDataAttributes) SetFolder(v string)`

SetFolder sets Folder field to given value.


### GetRoles

`func (o *CreateMediaRulesDataAttributes) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateMediaRulesDataAttributes) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateMediaRulesDataAttributes) SetRoles(v []string)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


