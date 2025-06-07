# \MediaRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChainsMediaStorageV1MediaRulesResourceDelete**](MediaRulesAPI.md#ChainsMediaStorageV1MediaRulesResourceDelete) | **Delete** /chains/media-storage/v1/media-rules/{resource} | Delete Media Rules
[**ChainsMediaStorageV1MediaRulesResourceGet**](MediaRulesAPI.md#ChainsMediaStorageV1MediaRulesResourceGet) | **Get** /chains/media-storage/v1/media-rules/{resource} | Get Media Rules
[**ChainsMediaStorageV1MediaRulesResourcePatch**](MediaRulesAPI.md#ChainsMediaStorageV1MediaRulesResourcePatch) | **Patch** /chains/media-storage/v1/media-rules/{resource} | Update Media Rules
[**ChainsMediaStorageV1MediaRulesResourcePost**](MediaRulesAPI.md#ChainsMediaStorageV1MediaRulesResourcePost) | **Post** /chains/media-storage/v1/media-rules/{resource} | Create Media Rules



## ChainsMediaStorageV1MediaRulesResourceDelete

> ChainsMediaStorageV1MediaRulesResourceDelete(ctx, resource).Execute()

Delete Media Rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	resource := "resource_example" // string | Идентификатор ресурсной записи

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MediaRulesAPI.ChainsMediaStorageV1MediaRulesResourceDelete(context.Background(), resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaRulesAPI.ChainsMediaStorageV1MediaRulesResourceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string** | Идентификатор ресурсной записи | 

### Other Parameters

Other parameters are passed through a pointer to a apiChainsMediaStorageV1MediaRulesResourceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChainsMediaStorageV1MediaRulesResourceGet

> MediaRules ChainsMediaStorageV1MediaRulesResourceGet(ctx, resource).Execute()

Get Media Rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	resource := "resource_example" // string | Идентификатор ресурсной записи

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaRulesAPI.ChainsMediaStorageV1MediaRulesResourceGet(context.Background(), resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaRulesAPI.ChainsMediaStorageV1MediaRulesResourceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChainsMediaStorageV1MediaRulesResourceGet`: MediaRules
	fmt.Fprintf(os.Stdout, "Response from `MediaRulesAPI.ChainsMediaStorageV1MediaRulesResourceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string** | Идентификатор ресурсной записи | 

### Other Parameters

Other parameters are passed through a pointer to a apiChainsMediaStorageV1MediaRulesResourceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MediaRules**](MediaRules.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChainsMediaStorageV1MediaRulesResourcePatch

> MediaRules ChainsMediaStorageV1MediaRulesResourcePatch(ctx, resource).UpdateMediaRules(updateMediaRules).Execute()

Update Media Rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	resource := "resource_example" // string | Идентификатор ресурсной записи
	updateMediaRules := *openapiclient.NewUpdateMediaRules(*openapiclient.NewUpdateMediaRulesData("Id_example", "Type_example", *openapiclient.NewUpdateMediaRulesDataAttributes())) // UpdateMediaRules | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaRulesAPI.ChainsMediaStorageV1MediaRulesResourcePatch(context.Background(), resource).UpdateMediaRules(updateMediaRules).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaRulesAPI.ChainsMediaStorageV1MediaRulesResourcePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChainsMediaStorageV1MediaRulesResourcePatch`: MediaRules
	fmt.Fprintf(os.Stdout, "Response from `MediaRulesAPI.ChainsMediaStorageV1MediaRulesResourcePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string** | Идентификатор ресурсной записи | 

### Other Parameters

Other parameters are passed through a pointer to a apiChainsMediaStorageV1MediaRulesResourcePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMediaRules** | [**UpdateMediaRules**](UpdateMediaRules.md) |  | 

### Return type

[**MediaRules**](MediaRules.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChainsMediaStorageV1MediaRulesResourcePost

> MediaRules ChainsMediaStorageV1MediaRulesResourcePost(ctx, resource).CreateMediaRules(createMediaRules).Execute()

Create Media Rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	resource := "resource_example" // string | Идентификатор ресурсной записи
	createMediaRules := *openapiclient.NewCreateMediaRules(*openapiclient.NewCreateMediaRulesData("Id_example", "Type_example", *openapiclient.NewCreateMediaRulesDataAttributes([]string{"Extensions_example"}, int64(123), []string{"Roles_example"}))) // CreateMediaRules | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaRulesAPI.ChainsMediaStorageV1MediaRulesResourcePost(context.Background(), resource).CreateMediaRules(createMediaRules).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaRulesAPI.ChainsMediaStorageV1MediaRulesResourcePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChainsMediaStorageV1MediaRulesResourcePost`: MediaRules
	fmt.Fprintf(os.Stdout, "Response from `MediaRulesAPI.ChainsMediaStorageV1MediaRulesResourcePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string** | Идентификатор ресурсной записи | 

### Other Parameters

Other parameters are passed through a pointer to a apiChainsMediaStorageV1MediaRulesResourcePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMediaRules** | [**CreateMediaRules**](CreateMediaRules.md) |  | 

### Return type

[**MediaRules**](MediaRules.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

