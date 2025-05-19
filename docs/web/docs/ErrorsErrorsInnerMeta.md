# ErrorsErrorsInnerMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Request ID is a unique identifier for the request, used for debugging and tracing | 
**Timestamp** | **time.Time** | Timestamp is the time when the error occurred, in ISO 8601 format | 

## Methods

### NewErrorsErrorsInnerMeta

`func NewErrorsErrorsInnerMeta(requestId string, timestamp time.Time, ) *ErrorsErrorsInnerMeta`

NewErrorsErrorsInnerMeta instantiates a new ErrorsErrorsInnerMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsErrorsInnerMetaWithDefaults

`func NewErrorsErrorsInnerMetaWithDefaults() *ErrorsErrorsInnerMeta`

NewErrorsErrorsInnerMetaWithDefaults instantiates a new ErrorsErrorsInnerMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsErrorsInnerMeta) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsErrorsInnerMeta) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsErrorsInnerMeta) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetTimestamp

`func (o *ErrorsErrorsInnerMeta) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ErrorsErrorsInnerMeta) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ErrorsErrorsInnerMeta) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


