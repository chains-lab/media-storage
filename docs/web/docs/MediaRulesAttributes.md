# MediaRulesAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSize** | **int64** | Maximum file size in bytes | 
**AllowedExits** | **[]string** |  | 
**Folder** | **string** | Folder where the media is stored | 
**Roles** | **[]string** |  | 

## Methods

### NewMediaRulesAttributes

`func NewMediaRulesAttributes(maxSize int64, allowedExits []string, folder string, roles []string, ) *MediaRulesAttributes`

NewMediaRulesAttributes instantiates a new MediaRulesAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaRulesAttributesWithDefaults

`func NewMediaRulesAttributesWithDefaults() *MediaRulesAttributes`

NewMediaRulesAttributesWithDefaults instantiates a new MediaRulesAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSize

`func (o *MediaRulesAttributes) GetMaxSize() int64`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *MediaRulesAttributes) GetMaxSizeOk() (*int64, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *MediaRulesAttributes) SetMaxSize(v int64)`

SetMaxSize sets MaxSize field to given value.


### GetAllowedExits

`func (o *MediaRulesAttributes) GetAllowedExits() []string`

GetAllowedExits returns the AllowedExits field if non-nil, zero value otherwise.

### GetAllowedExitsOk

`func (o *MediaRulesAttributes) GetAllowedExitsOk() (*[]string, bool)`

GetAllowedExitsOk returns a tuple with the AllowedExits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedExits

`func (o *MediaRulesAttributes) SetAllowedExits(v []string)`

SetAllowedExits sets AllowedExits field to given value.


### GetFolder

`func (o *MediaRulesAttributes) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *MediaRulesAttributes) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *MediaRulesAttributes) SetFolder(v string)`

SetFolder sets Folder field to given value.


### GetRoles

`func (o *MediaRulesAttributes) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *MediaRulesAttributes) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *MediaRulesAttributes) SetRoles(v []string)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


