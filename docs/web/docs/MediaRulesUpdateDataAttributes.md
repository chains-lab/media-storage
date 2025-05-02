# MediaRulesUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSize** | Pointer to **int32** | Maximum file size in bytes | [optional] 
**AllowedExits** | Pointer to **[]string** |  | [optional] 
**Folder** | Pointer to **string** | Folder where the media is stored | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMediaRulesUpdateDataAttributes

`func NewMediaRulesUpdateDataAttributes() *MediaRulesUpdateDataAttributes`

NewMediaRulesUpdateDataAttributes instantiates a new MediaRulesUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaRulesUpdateDataAttributesWithDefaults

`func NewMediaRulesUpdateDataAttributesWithDefaults() *MediaRulesUpdateDataAttributes`

NewMediaRulesUpdateDataAttributesWithDefaults instantiates a new MediaRulesUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSize

`func (o *MediaRulesUpdateDataAttributes) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *MediaRulesUpdateDataAttributes) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *MediaRulesUpdateDataAttributes) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *MediaRulesUpdateDataAttributes) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### GetAllowedExits

`func (o *MediaRulesUpdateDataAttributes) GetAllowedExits() []string`

GetAllowedExits returns the AllowedExits field if non-nil, zero value otherwise.

### GetAllowedExitsOk

`func (o *MediaRulesUpdateDataAttributes) GetAllowedExitsOk() (*[]string, bool)`

GetAllowedExitsOk returns a tuple with the AllowedExits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedExits

`func (o *MediaRulesUpdateDataAttributes) SetAllowedExits(v []string)`

SetAllowedExits sets AllowedExits field to given value.

### HasAllowedExits

`func (o *MediaRulesUpdateDataAttributes) HasAllowedExits() bool`

HasAllowedExits returns a boolean if a field has been set.

### GetFolder

`func (o *MediaRulesUpdateDataAttributes) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *MediaRulesUpdateDataAttributes) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *MediaRulesUpdateDataAttributes) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *MediaRulesUpdateDataAttributes) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetRoles

`func (o *MediaRulesUpdateDataAttributes) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *MediaRulesUpdateDataAttributes) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *MediaRulesUpdateDataAttributes) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *MediaRulesUpdateDataAttributes) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


