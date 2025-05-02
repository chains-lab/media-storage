# UpdateMediaRulesDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSize** | Pointer to **int64** | Maximum file size in bytes | [optional] 
**AllowedExits** | Pointer to **[]string** |  | [optional] 
**Folder** | Pointer to **string** | Folder where the media is stored | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateMediaRulesDataAttributes

`func NewUpdateMediaRulesDataAttributes() *UpdateMediaRulesDataAttributes`

NewUpdateMediaRulesDataAttributes instantiates a new UpdateMediaRulesDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMediaRulesDataAttributesWithDefaults

`func NewUpdateMediaRulesDataAttributesWithDefaults() *UpdateMediaRulesDataAttributes`

NewUpdateMediaRulesDataAttributesWithDefaults instantiates a new UpdateMediaRulesDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSize

`func (o *UpdateMediaRulesDataAttributes) GetMaxSize() int64`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *UpdateMediaRulesDataAttributes) GetMaxSizeOk() (*int64, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *UpdateMediaRulesDataAttributes) SetMaxSize(v int64)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *UpdateMediaRulesDataAttributes) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### GetAllowedExits

`func (o *UpdateMediaRulesDataAttributes) GetAllowedExits() []string`

GetAllowedExits returns the AllowedExits field if non-nil, zero value otherwise.

### GetAllowedExitsOk

`func (o *UpdateMediaRulesDataAttributes) GetAllowedExitsOk() (*[]string, bool)`

GetAllowedExitsOk returns a tuple with the AllowedExits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedExits

`func (o *UpdateMediaRulesDataAttributes) SetAllowedExits(v []string)`

SetAllowedExits sets AllowedExits field to given value.

### HasAllowedExits

`func (o *UpdateMediaRulesDataAttributes) HasAllowedExits() bool`

HasAllowedExits returns a boolean if a field has been set.

### GetFolder

`func (o *UpdateMediaRulesDataAttributes) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *UpdateMediaRulesDataAttributes) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *UpdateMediaRulesDataAttributes) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *UpdateMediaRulesDataAttributes) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetRoles

`func (o *UpdateMediaRulesDataAttributes) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateMediaRulesDataAttributes) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateMediaRulesDataAttributes) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateMediaRulesDataAttributes) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


