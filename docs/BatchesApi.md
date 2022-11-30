# \BatchesApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveBatches**](BatchesApi.md#ArchiveBatches) | **Post** /batches:archive | Archive Batches
[**BulkGetBatches**](BatchesApi.md#BulkGetBatches) | **Get** /batches:bulk-get | Bulk get batches
[**CreateBatch**](BatchesApi.md#CreateBatch) | **Post** /batches | Create a batch
[**GetBatch**](BatchesApi.md#GetBatch) | **Get** /batches/{batch_id} | Get a batch
[**GetEnitityBatches**](BatchesApi.md#GetEnitityBatches) | **Get** /entities/{entity_id}/batches | Get an entity&#39;s batches
[**ListBatches**](BatchesApi.md#ListBatches) | **Get** /batches | List batches
[**UnarchiveBatches**](BatchesApi.md#UnarchiveBatches) | **Post** /batches:unarchive | Unarchive Batches
[**UpdateBatch**](BatchesApi.md#UpdateBatch) | **Patch** /batches/{batch_id} | Update a batch



## ArchiveBatches

> BatchesArchivalChange ArchiveBatches(ctx).BatchesArchive(batchesArchive).Execute()

Archive Batches



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
    batchesArchive := *openapiclient.NewBatchesArchive([]string{"BatchIds_example"}, "Reason_example") // BatchesArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchesApi.ArchiveBatches(context.Background()).BatchesArchive(batchesArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchesApi.ArchiveBatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveBatches`: BatchesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `BatchesApi.ArchiveBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchesArchive** | [**BatchesArchive**](BatchesArchive.md) |  | 

### Return type

[**BatchesArchivalChange**](BatchesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkGetBatches

> BatchesBulkGet BulkGetBatches(ctx).BatchIds(batchIds).BatchNames(batchNames).RegistryId(registryId).Execute()

Bulk get batches



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
    batchIds := "batchIds_example" // string | Comma-separated list of batch IDs. (optional)
    batchNames := "batchNames_example" // string | Comma-separated list of batch names. Batch names have the format {file name}-{creation date as YYYYMMDD}-{optional inventory number}. Must specify registryId with batchNames.  (optional)
    registryId := "registryId_example" // string | ID of the registry that batches are registered in. Required if querying by batchNames. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchesApi.BulkGetBatches(context.Background()).BatchIds(batchIds).BatchNames(batchNames).RegistryId(registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchesApi.BulkGetBatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetBatches`: BatchesBulkGet
    fmt.Fprintf(os.Stdout, "Response from `BatchesApi.BulkGetBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchIds** | **string** | Comma-separated list of batch IDs. | 
 **batchNames** | **string** | Comma-separated list of batch names. Batch names have the format {file name}-{creation date as YYYYMMDD}-{optional inventory number}. Must specify registryId with batchNames.  | 
 **registryId** | **string** | ID of the registry that batches are registered in. Required if querying by batchNames. | 

### Return type

[**BatchesBulkGet**](BatchesBulkGet.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBatch

> Batch CreateBatch(ctx).BatchCreate(batchCreate).Execute()

Create a batch



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
    batchCreate := *openapiclient.NewBatchCreate() // BatchCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchesApi.CreateBatch(context.Background()).BatchCreate(batchCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchesApi.CreateBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBatch`: Batch
    fmt.Fprintf(os.Stdout, "Response from `BatchesApi.CreateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchCreate** | [**BatchCreate**](BatchCreate.md) |  | 

### Return type

[**Batch**](Batch.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatch

> Batch GetBatch(ctx, batchId).Execute()

Get a batch



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
    batchId := "batchId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchesApi.GetBatch(context.Background(), batchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchesApi.GetBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBatch`: Batch
    fmt.Fprintf(os.Stdout, "Response from `BatchesApi.GetBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Batch**](Batch.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnitityBatches

> BatchesPaginatedList GetEnitityBatches(ctx, entityId).Execute()

Get an entity's batches



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
    entityId := "entityId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchesApi.GetEnitityBatches(context.Background(), entityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchesApi.GetEnitityBatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnitityBatches`: BatchesPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `BatchesApi.GetEnitityBatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnitityBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BatchesPaginatedList**](BatchesPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBatches

> BatchesPaginatedList ListBatches(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Ids(ids).CreatorIds(creatorIds).Execute()

List batches



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
    pageSize := int32(56) // int32 | Number of results to return. Defaults to 50, maximum of 100.  (optional) (default to 50)
    nextToken := "nextToken_example" // string | Token for pagination (optional)
    sort := "sort_example" // string |  (optional) (default to "modifiedAt:desc")
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    schemaId := "schemaId_example" // string | ID of a schema. Restricts results to those of the specified schema.  (optional)
    schemaFields := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Schema field value. For Integer, Float, and Date type fields, supports the >= and <= operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived batches. Use \"ANY_ARCHIVED\" to filter for archived batches regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    ids := "bat_VfVOART1,bat_RFhDGaaC" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of users IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchesApi.ListBatches(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Ids(ids).CreatorIds(creatorIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchesApi.ListBatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBatches`: BatchesPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `BatchesApi.ListBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100.  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema.  | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived batches. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived batches regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **creatorIds** | **string** | Comma separated list of users IDs | 

### Return type

[**BatchesPaginatedList**](BatchesPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveBatches

> BatchesArchivalChange UnarchiveBatches(ctx).BatchesUnarchive(batchesUnarchive).Execute()

Unarchive Batches



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
    batchesUnarchive := *openapiclient.NewBatchesUnarchive([]string{"BatchIds_example"}) // BatchesUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchesApi.UnarchiveBatches(context.Background()).BatchesUnarchive(batchesUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchesApi.UnarchiveBatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveBatches`: BatchesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `BatchesApi.UnarchiveBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchesUnarchive** | [**BatchesUnarchive**](BatchesUnarchive.md) |  | 

### Return type

[**BatchesArchivalChange**](BatchesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBatch

> Batch UpdateBatch(ctx, batchId).BatchUpdate(batchUpdate).Execute()

Update a batch



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
    batchId := "batchId_example" // string | 
    batchUpdate := *openapiclient.NewBatchUpdate() // BatchUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchesApi.UpdateBatch(context.Background(), batchId).BatchUpdate(batchUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchesApi.UpdateBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBatch`: Batch
    fmt.Fprintf(os.Stdout, "Response from `BatchesApi.UpdateBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchUpdate** | [**BatchUpdate**](BatchUpdate.md) |  | 

### Return type

[**Batch**](Batch.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

