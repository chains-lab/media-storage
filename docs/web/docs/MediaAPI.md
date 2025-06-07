# \MediaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChainsMediaStorageV1MediaDelete**](MediaAPI.md#ChainsMediaStorageV1MediaDelete) | **Delete** /chains/media-storage/v1/media/ | Delete Media
[**ChainsMediaStorageV1MediaPost**](MediaAPI.md#ChainsMediaStorageV1MediaPost) | **Post** /chains/media-storage/v1/media/ | Upload Media



## ChainsMediaStorageV1MediaDelete

> ChainsMediaStorageV1MediaDelete(ctx, mediaId).Execute()

Delete Media



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
	mediaId := "mediaId_example" // string | The ID of the media to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MediaAPI.ChainsMediaStorageV1MediaDelete(context.Background(), mediaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.ChainsMediaStorageV1MediaDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaId** | **string** | The ID of the media to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiChainsMediaStorageV1MediaDeleteRequest struct via the builder pattern


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


## ChainsMediaStorageV1MediaPost

> Media ChainsMediaStorageV1MediaPost(ctx).UploadMedia(uploadMedia).Execute()

Upload Media



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
	uploadMedia := *openapiclient.NewUploadMedia(*openapiclient.NewUploadMediaData("Type_example", *openapiclient.NewUploadMediaDataAttributes("Resource_example", "ResourceId_example", "Category_example", "OwnerId_example"))) // UploadMedia | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.ChainsMediaStorageV1MediaPost(context.Background()).UploadMedia(uploadMedia).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.ChainsMediaStorageV1MediaPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChainsMediaStorageV1MediaPost`: Media
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.ChainsMediaStorageV1MediaPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChainsMediaStorageV1MediaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadMedia** | [**UploadMedia**](UploadMedia.md) |  | 

### Return type

[**Media**](Media.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

