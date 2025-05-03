# MediaRulesAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExitSize** | [**[]ExitSizeInner**](ExitSizeInner.md) |  | 
**Roles** | **[]string** |  | 
**UpdatedAt** | **time.Time** | Date and time when the media rules were last updated | 
**CreatedAt** | **time.Time** | Date and time when the media rules were created | 

## Methods

### NewMediaRulesAttributes

`func NewMediaRulesAttributes(exitSize []ExitSizeInner, roles []string, updatedAt time.Time, createdAt time.Time, ) *MediaRulesAttributes`

NewMediaRulesAttributes instantiates a new MediaRulesAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaRulesAttributesWithDefaults

`func NewMediaRulesAttributesWithDefaults() *MediaRulesAttributes`

NewMediaRulesAttributesWithDefaults instantiates a new MediaRulesAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExitSize

`func (o *MediaRulesAttributes) GetExitSize() []ExitSizeInner`

GetExitSize returns the ExitSize field if non-nil, zero value otherwise.

### GetExitSizeOk

`func (o *MediaRulesAttributes) GetExitSizeOk() (*[]ExitSizeInner, bool)`

GetExitSizeOk returns a tuple with the ExitSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitSize

`func (o *MediaRulesAttributes) SetExitSize(v []ExitSizeInner)`

SetExitSize sets ExitSize field to given value.


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


### GetUpdatedAt

`func (o *MediaRulesAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MediaRulesAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MediaRulesAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *MediaRulesAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MediaRulesAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MediaRulesAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


