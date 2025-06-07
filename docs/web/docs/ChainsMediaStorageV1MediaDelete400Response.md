# ChainsMediaStorageV1MediaDelete400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ChainsMediaStorageV1MediaDelete400ResponseErrorsInner**](ChainsMediaStorageV1MediaDelete400ResponseErrorsInner.md) | Non empty array of errors occurred during request processing | 

## Methods

### NewChainsMediaStorageV1MediaDelete400Response

`func NewChainsMediaStorageV1MediaDelete400Response(errors []ChainsMediaStorageV1MediaDelete400ResponseErrorsInner, ) *ChainsMediaStorageV1MediaDelete400Response`

NewChainsMediaStorageV1MediaDelete400Response instantiates a new ChainsMediaStorageV1MediaDelete400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChainsMediaStorageV1MediaDelete400ResponseWithDefaults

`func NewChainsMediaStorageV1MediaDelete400ResponseWithDefaults() *ChainsMediaStorageV1MediaDelete400Response`

NewChainsMediaStorageV1MediaDelete400ResponseWithDefaults instantiates a new ChainsMediaStorageV1MediaDelete400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ChainsMediaStorageV1MediaDelete400Response) GetErrors() []ChainsMediaStorageV1MediaDelete400ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ChainsMediaStorageV1MediaDelete400Response) GetErrorsOk() (*[]ChainsMediaStorageV1MediaDelete400ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ChainsMediaStorageV1MediaDelete400Response) SetErrors(v []ChainsMediaStorageV1MediaDelete400ResponseErrorsInner)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


