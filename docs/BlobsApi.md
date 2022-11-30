# \BlobsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortMultipartBlob**](BlobsApi.md#AbortMultipartBlob) | **Post** /blobs/{blob_id}:abort-upload | Abort multi-part blob upload
[**BulkGetBlobs**](BlobsApi.md#BulkGetBlobs) | **Get** /blobs:bulk-get | Bulk get Blobs by UUID
[**CompleteMultipartBlob**](BlobsApi.md#CompleteMultipartBlob) | **Post** /blobs/{blob_id}:complete-upload | Complete multi-part blob upload
[**CreateBlob**](BlobsApi.md#CreateBlob) | **Post** /blobs | Upload single-part blob
[**CreateBlobPart**](BlobsApi.md#CreateBlobPart) | **Post** /blobs/{blob_id}/parts | Upload a part of a multi-part blob
[**CreateMultipartBlob**](BlobsApi.md#CreateMultipartBlob) | **Post** /blobs:start-multipart-upload | Initiate multi-part blob upload
[**GetBlob**](BlobsApi.md#GetBlob) | **Get** /blobs/{blob_id} | Get a Blob
[**GetBlobUrl**](BlobsApi.md#GetBlobUrl) | **Get** /blobs/{blob_id}/download-url | Get a Blob&#39;s download url



## AbortMultipartBlob

> map[string]interface{} AbortMultipartBlob(ctx, blobId).Execute()

Abort multi-part blob upload



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    blobId := "blobId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlobsApi.AbortMultipartBlob(context.Background(), blobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobsApi.AbortMultipartBlob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AbortMultipartBlob`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BlobsApi.AbortMultipartBlob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAbortMultipartBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkGetBlobs

> BlobsBulkGet BulkGetBlobs(ctx).BlobIds(blobIds).Execute()

Bulk get Blobs by UUID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    blobIds := "blobIds_example" // string | Comma-separated list of blob IDs. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlobsApi.BulkGetBlobs(context.Background()).BlobIds(blobIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobsApi.BulkGetBlobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetBlobs`: BlobsBulkGet
    fmt.Fprintf(os.Stdout, "Response from `BlobsApi.BulkGetBlobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetBlobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blobIds** | **string** | Comma-separated list of blob IDs. | 

### Return type

[**BlobsBulkGet**](BlobsBulkGet.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteMultipartBlob

> Blob CompleteMultipartBlob(ctx, blobId).BlobComplete(blobComplete).Execute()

Complete multi-part blob upload



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    blobId := "blobId_example" // string | 
    blobComplete := *openapiclient.NewBlobComplete() // BlobComplete |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlobsApi.CompleteMultipartBlob(context.Background(), blobId).BlobComplete(blobComplete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobsApi.CompleteMultipartBlob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompleteMultipartBlob`: Blob
    fmt.Fprintf(os.Stdout, "Response from `BlobsApi.CompleteMultipartBlob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteMultipartBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blobComplete** | [**BlobComplete**](BlobComplete.md) |  | 

### Return type

[**Blob**](Blob.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBlob

> Blob CreateBlob(ctx).BlobCreate(blobCreate).Execute()

Upload single-part blob



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    blobCreate := *openapiclient.NewBlobCreate(string(123), "Md5_example", "Name_example", "Type_example") // BlobCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlobsApi.CreateBlob(context.Background()).BlobCreate(blobCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobsApi.CreateBlob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBlob`: Blob
    fmt.Fprintf(os.Stdout, "Response from `BlobsApi.CreateBlob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blobCreate** | [**BlobCreate**](BlobCreate.md) |  | 

### Return type

[**Blob**](Blob.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBlobPart

> BlobPart CreateBlobPart(ctx, blobId).BlobPartCreate(blobPartCreate).Execute()

Upload a part of a multi-part blob



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    blobId := "blobId_example" // string | 
    blobPartCreate := *openapiclient.NewBlobPartCreate("Data64_example", "Md5_example", int32(123)) // BlobPartCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlobsApi.CreateBlobPart(context.Background(), blobId).BlobPartCreate(blobPartCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobsApi.CreateBlobPart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBlobPart`: BlobPart
    fmt.Fprintf(os.Stdout, "Response from `BlobsApi.CreateBlobPart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlobPartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blobPartCreate** | [**BlobPartCreate**](BlobPartCreate.md) |  | 

### Return type

[**BlobPart**](BlobPart.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMultipartBlob

> Blob CreateMultipartBlob(ctx).BlobMultipartCreate(blobMultipartCreate).Execute()

Initiate multi-part blob upload



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    blobMultipartCreate := *openapiclient.NewBlobMultipartCreate("Name_example", "Type_example") // BlobMultipartCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlobsApi.CreateMultipartBlob(context.Background()).BlobMultipartCreate(blobMultipartCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobsApi.CreateMultipartBlob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMultipartBlob`: Blob
    fmt.Fprintf(os.Stdout, "Response from `BlobsApi.CreateMultipartBlob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMultipartBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blobMultipartCreate** | [**BlobMultipartCreate**](BlobMultipartCreate.md) |  | 

### Return type

[**Blob**](Blob.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlob

> Blob GetBlob(ctx, blobId).Execute()

Get a Blob



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    blobId := "blobId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlobsApi.GetBlob(context.Background(), blobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobsApi.GetBlob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlob`: Blob
    fmt.Fprintf(os.Stdout, "Response from `BlobsApi.GetBlob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Blob**](Blob.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlobUrl

> BlobUrl GetBlobUrl(ctx, blobId).Execute()

Get a Blob's download url



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    blobId := "blobId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlobsApi.GetBlobUrl(context.Background(), blobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobsApi.GetBlobUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlobUrl`: BlobUrl
    fmt.Fprintf(os.Stdout, "Response from `BlobsApi.GetBlobUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlobUrl**](BlobUrl.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

