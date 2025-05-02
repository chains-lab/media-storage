# MediaRulesAddDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSize** | **int32** | Maximum file size in bytes | 
**AllowedExits** | **[]string** |  | 
**Folder** | **string** | Folder where the media is stored | 
**Roles** | **[]string** |  | 

## Methods

### NewMediaRulesAddDataAttributes

`func NewMediaRulesAddDataAttributes(maxSize int32, allowedExits []string, folder string, roles []string, ) *MediaRulesAddDataAttributes`

NewMediaRulesAddDataAttributes instantiates a new MediaRulesAddDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaRulesAddDataAttributesWithDefaults

`func NewMediaRulesAddDataAttributesWithDefaults() *MediaRulesAddDataAttributes`

NewMediaRulesAddDataAttributesWithDefaults instantiates a new MediaRulesAddDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSize

`func (o *MediaRulesAddDataAttributes) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *MediaRulesAddDataAttributes) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *MediaRulesAddDataAttributes) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.


### GetAllowedExits

`func (o *MediaRulesAddDataAttributes) GetAllowedExits() []string`

GetAllowedExits returns the AllowedExits field if non-nil, zero value otherwise.

### GetAllowedExitsOk

`func (o *MediaRulesAddDataAttributes) GetAllowedExitsOk() (*[]string, bool)`

GetAllowedExitsOk returns a tuple with the AllowedExits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedExits

`func (o *MediaRulesAddDataAttributes) SetAllowedExits(v []string)`

SetAllowedExits sets AllowedExits field to given value.


### GetFolder

`func (o *MediaRulesAddDataAttributes) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *MediaRulesAddDataAttributes) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *MediaRulesAddDataAttributes) SetFolder(v string)`

SetFolder sets Folder field to given value.


### GetRoles

`func (o *MediaRulesAddDataAttributes) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *MediaRulesAddDataAttributes) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *MediaRulesAddDataAttributes) SetRoles(v []string)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


