# \MediaRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChainsMediaStorageV1MediaRulesDelete**](MediaRulesAPI.md#ChainsMediaStorageV1MediaRulesDelete) | **Delete** /chains/media-storage/v1/media-rules/ | Delete Media Rules
[**ChainsMediaStorageV1MediaRulesGet**](MediaRulesAPI.md#ChainsMediaStorageV1MediaRulesGet) | **Get** /chains/media-storage/v1/media-rules/ | Get Media Rules
[**ChainsMediaStorageV1MediaRulesPatch**](MediaRulesAPI.md#ChainsMediaStorageV1MediaRulesPatch) | **Patch** /chains/media-storage/v1/media-rules/ | Update Media Rules
[**ChainsMediaStorageV1MediaRulesPost**](MediaRulesAPI.md#ChainsMediaStorageV1MediaRulesPost) | **Post** /chains/media-storage/v1/media-rules/ | Create Media Rules



## ChainsMediaStorageV1MediaRulesDelete

> ChainsMediaStorageV1MediaRulesDelete(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MediaRulesAPI.ChainsMediaStorageV1MediaRulesDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaRulesAPI.ChainsMediaStorageV1MediaRulesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChainsMediaStorageV1MediaRulesDeleteRequest struct via the builder pattern


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


## ChainsMediaStorageV1MediaRulesGet

> MediaRules ChainsMediaStorageV1MediaRulesGet(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaRulesAPI.ChainsMediaStorageV1MediaRulesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaRulesAPI.ChainsMediaStorageV1MediaRulesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChainsMediaStorageV1MediaRulesGet`: MediaRules
	fmt.Fprintf(os.Stdout, "Response from `MediaRulesAPI.ChainsMediaStorageV1MediaRulesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChainsMediaStorageV1MediaRulesGetRequest struct via the builder pattern


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


## ChainsMediaStorageV1MediaRulesPatch

> MediaRules ChainsMediaStorageV1MediaRulesPatch(ctx).UpdateMediaRules(updateMediaRules).Execute()

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
	updateMediaRules := *openapiclient.NewUpdateMediaRules(*openapiclient.NewUpdateMediaRulesData("Id_example", "Type_example", *openapiclient.NewUpdateMediaRulesDataAttributes())) // UpdateMediaRules | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaRulesAPI.ChainsMediaStorageV1MediaRulesPatch(context.Background()).UpdateMediaRules(updateMediaRules).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaRulesAPI.ChainsMediaStorageV1MediaRulesPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChainsMediaStorageV1MediaRulesPatch`: MediaRules
	fmt.Fprintf(os.Stdout, "Response from `MediaRulesAPI.ChainsMediaStorageV1MediaRulesPatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChainsMediaStorageV1MediaRulesPatchRequest struct via the builder pattern


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


## ChainsMediaStorageV1MediaRulesPost

> MediaRules ChainsMediaStorageV1MediaRulesPost(ctx).CreateMediaRules(createMediaRules).Execute()

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
	createMediaRules := *openapiclient.NewCreateMediaRules(*openapiclient.NewCreateMediaRulesData("Id_example", "Type_example", *openapiclient.NewCreateMediaRulesDataAttributes([]string{"Extensions_example"}, int64(123), []string{"Roles_example"}))) // CreateMediaRules | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaRulesAPI.ChainsMediaStorageV1MediaRulesPost(context.Background()).CreateMediaRules(createMediaRules).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaRulesAPI.ChainsMediaStorageV1MediaRulesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChainsMediaStorageV1MediaRulesPost`: MediaRules
	fmt.Fprintf(os.Stdout, "Response from `MediaRulesAPI.ChainsMediaStorageV1MediaRulesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChainsMediaStorageV1MediaRulesPostRequest struct via the builder pattern


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

