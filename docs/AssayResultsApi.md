# \AssayResultsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortAssayResultsTransaction**](AssayResultsApi.md#AbortAssayResultsTransaction) | **Post** /result-transactions/{transaction_id}:abort | Abort a transaction
[**ArchiveAssayResults**](AssayResultsApi.md#ArchiveAssayResults) | **Post** /assay-results:archive | Archive 1 or more results.
[**BulkGetAssayResults**](AssayResultsApi.md#BulkGetAssayResults) | **Get** /assay-results:bulk-get | Gets multiple results specified by a list of IDs.
[**CommitAssayResultsTransaction**](AssayResultsApi.md#CommitAssayResultsTransaction) | **Post** /result-transactions/{transaction_id}:commit | Commit a transaction
[**CreateAssayResults**](AssayResultsApi.md#CreateAssayResults) | **Post** /assay-results | Create 1 or more results.
[**CreateAssayResultsInTransaction**](AssayResultsApi.md#CreateAssayResultsInTransaction) | **Post** /result-transactions/{transaction_id}/results | Create results in a transaction
[**CreateAssayResultsTransaction**](AssayResultsApi.md#CreateAssayResultsTransaction) | **Post** /result-transactions | Create a transaction
[**GetAssayResult**](AssayResultsApi.md#GetAssayResult) | **Get** /assay-results/{assay_result_id} | Get a result
[**ListAssayResults**](AssayResultsApi.md#ListAssayResults) | **Get** /assay-results | List results
[**UnarchiveAssayResults**](AssayResultsApi.md#UnarchiveAssayResults) | **Post** /assay-results:unarchive | Unarchive 1 or more results.



## AbortAssayResultsTransaction

> AssayResultTransactionCreateResponse AbortAssayResultsTransaction(ctx, transactionId).Execute()

Abort a transaction



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
    transactionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayResultsApi.AbortAssayResultsTransaction(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayResultsApi.AbortAssayResultsTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AbortAssayResultsTransaction`: AssayResultTransactionCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `AssayResultsApi.AbortAssayResultsTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAbortAssayResultsTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssayResultTransactionCreateResponse**](AssayResultTransactionCreateResponse.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchiveAssayResults

> AssayResultIdsResponse ArchiveAssayResults(ctx).AssayResultsArchive(assayResultsArchive).Execute()

Archive 1 or more results.



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
    assayResultsArchive := *openapiclient.NewAssayResultsArchive([]string{"AssayResultIds_example"}) // AssayResultsArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayResultsApi.ArchiveAssayResults(context.Background()).AssayResultsArchive(assayResultsArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayResultsApi.ArchiveAssayResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveAssayResults`: AssayResultIdsResponse
    fmt.Fprintf(os.Stdout, "Response from `AssayResultsApi.ArchiveAssayResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveAssayResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assayResultsArchive** | [**AssayResultsArchive**](AssayResultsArchive.md) |  | 

### Return type

[**AssayResultIdsResponse**](AssayResultIdsResponse.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkGetAssayResults

> AssayResultsBulkGet BulkGetAssayResults(ctx).AssayResultIds(assayResultIds).Execute()

Gets multiple results specified by a list of IDs.



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
    assayResultIds := "assayResultIds_example" // string | Comma-separated list of assay result IDs.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayResultsApi.BulkGetAssayResults(context.Background()).AssayResultIds(assayResultIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayResultsApi.BulkGetAssayResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetAssayResults`: AssayResultsBulkGet
    fmt.Fprintf(os.Stdout, "Response from `AssayResultsApi.BulkGetAssayResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetAssayResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assayResultIds** | **string** | Comma-separated list of assay result IDs. | 

### Return type

[**AssayResultsBulkGet**](AssayResultsBulkGet.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommitAssayResultsTransaction

> AssayResultTransactionCreateResponse CommitAssayResultsTransaction(ctx, transactionId).Execute()

Commit a transaction



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
    transactionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayResultsApi.CommitAssayResultsTransaction(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayResultsApi.CommitAssayResultsTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommitAssayResultsTransaction`: AssayResultTransactionCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `AssayResultsApi.CommitAssayResultsTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommitAssayResultsTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssayResultTransactionCreateResponse**](AssayResultTransactionCreateResponse.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssayResults

> AssayResultsCreateResponse CreateAssayResults(ctx).AssayResultsBulkCreateRequest(assayResultsBulkCreateRequest).Execute()

Create 1 or more results.



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
    assayResultsBulkCreateRequest := *openapiclient.NewAssayResultsBulkCreateRequest([]openapiclient.AssayResultCreate{*openapiclient.NewAssayResultCreate(*openapiclient.NewAssayResultCreateFields(), "SchemaId_example")}) // AssayResultsBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayResultsApi.CreateAssayResults(context.Background()).AssayResultsBulkCreateRequest(assayResultsBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayResultsApi.CreateAssayResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssayResults`: AssayResultsCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `AssayResultsApi.CreateAssayResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssayResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assayResultsBulkCreateRequest** | [**AssayResultsBulkCreateRequest**](AssayResultsBulkCreateRequest.md) |  | 

### Return type

[**AssayResultsCreateResponse**](AssayResultsCreateResponse.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssayResultsInTransaction

> AssayResultsCreateResponse CreateAssayResultsInTransaction(ctx, transactionId).AssayResultsBulkCreateRequest(assayResultsBulkCreateRequest).Execute()

Create results in a transaction



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
    transactionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    assayResultsBulkCreateRequest := *openapiclient.NewAssayResultsBulkCreateRequest([]openapiclient.AssayResultCreate{*openapiclient.NewAssayResultCreate(*openapiclient.NewAssayResultCreateFields(), "SchemaId_example")}) // AssayResultsBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayResultsApi.CreateAssayResultsInTransaction(context.Background(), transactionId).AssayResultsBulkCreateRequest(assayResultsBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayResultsApi.CreateAssayResultsInTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssayResultsInTransaction`: AssayResultsCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `AssayResultsApi.CreateAssayResultsInTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssayResultsInTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assayResultsBulkCreateRequest** | [**AssayResultsBulkCreateRequest**](AssayResultsBulkCreateRequest.md) |  | 

### Return type

[**AssayResultsCreateResponse**](AssayResultsCreateResponse.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssayResultsTransaction

> AssayResultTransactionCreateResponse CreateAssayResultsTransaction(ctx).Execute()

Create a transaction



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayResultsApi.CreateAssayResultsTransaction(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayResultsApi.CreateAssayResultsTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssayResultsTransaction`: AssayResultTransactionCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `AssayResultsApi.CreateAssayResultsTransaction`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssayResultsTransactionRequest struct via the builder pattern


### Return type

[**AssayResultTransactionCreateResponse**](AssayResultTransactionCreateResponse.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssayResult

> AssayResult GetAssayResult(ctx, assayResultId).Execute()

Get a result



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
    assayResultId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayResultsApi.GetAssayResult(context.Background(), assayResultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayResultsApi.GetAssayResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssayResult`: AssayResult
    fmt.Fprintf(os.Stdout, "Response from `AssayResultsApi.GetAssayResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assayResultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssayResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssayResult**](AssayResult.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssayResults

> AssayResultsPaginatedList ListAssayResults(ctx).SchemaId(schemaId).CreatedAtLt(createdAtLt).CreatedAtGt(createdAtGt).CreatedAtLte(createdAtLte).CreatedAtGte(createdAtGte).MinCreatedTime(minCreatedTime).MaxCreatedTime(maxCreatedTime).Sort(sort).NextToken(nextToken).PageSize(pageSize).EntityIds(entityIds).StorageIds(storageIds).AssayRunIds(assayRunIds).Ids(ids).Execute()

List results



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
    schemaId := "schemaId_example" // string | ID of the assay result schema to filter by
    createdAtLt := "createdAtLt_example" // string | Datetime, in RFC 3339 format. Time zone defaults to UTC. Restricts results to those created before the specified time. e.g. < 2017-04-30.  (optional)
    createdAtGt := "createdAtGt_example" // string | Datetime, in RFC 3339 format. Time zone defaults to UTC. Restricts results to those created after the specified time. e.g. > 2017-04-30.  (optional)
    createdAtLte := "createdAtLte_example" // string | Datetime, in RFC 3339 format. Time zone defaults to UTC. Restricts results to those created on or before the specified time. e.g. <= 2017-04-30.  (optional)
    createdAtGte := "createdAtGte_example" // string | Datetime, in RFC 3339 format. Time zone defaults to UTC. Restricts results to those created on or after the specified time. e.g. >= 2017-04-30.  (optional)
    minCreatedTime := int32(56) // int32 | Filter by results created after this unix timestamp (optional)
    maxCreatedTime := int32(56) // int32 | Filter by results created before this unix timestamp (optional)
    sort := "sort_example" // string | Method by which to order search results. Valid sorts are createdAt (created time, oldest first). Use :asc or :desc to specify ascending or descending order. Default is createdAt:asc.  (optional) (default to "createdAt:asc")
    nextToken := "nextToken_example" // string | Token for pagination (optional)
    pageSize := int32(56) // int32 | Number of results to return. Defaults to 50, maximum of 100. (optional) (default to 50)
    entityIds := "entityIds_example" // string | Filter by comma-separated list of related Entity IDs, maximum of 20. (optional)
    storageIds := "storageIds_example" // string | Filter by comma-separated list of related storage (container, box, plate, or location) IDs, maximum of 20.  (optional)
    assayRunIds := "assayRunIds_example" // string | Filter by comma-separated list of associated AssayRun IDs. (optional)
    ids := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayResultsApi.ListAssayResults(context.Background()).SchemaId(schemaId).CreatedAtLt(createdAtLt).CreatedAtGt(createdAtGt).CreatedAtLte(createdAtLte).CreatedAtGte(createdAtGte).MinCreatedTime(minCreatedTime).MaxCreatedTime(maxCreatedTime).Sort(sort).NextToken(nextToken).PageSize(pageSize).EntityIds(entityIds).StorageIds(storageIds).AssayRunIds(assayRunIds).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayResultsApi.ListAssayResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAssayResults`: AssayResultsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `AssayResultsApi.ListAssayResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssayResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaId** | **string** | ID of the assay result schema to filter by | 
 **createdAtLt** | **string** | Datetime, in RFC 3339 format. Time zone defaults to UTC. Restricts results to those created before the specified time. e.g. &lt; 2017-04-30.  | 
 **createdAtGt** | **string** | Datetime, in RFC 3339 format. Time zone defaults to UTC. Restricts results to those created after the specified time. e.g. &gt; 2017-04-30.  | 
 **createdAtLte** | **string** | Datetime, in RFC 3339 format. Time zone defaults to UTC. Restricts results to those created on or before the specified time. e.g. &lt;&#x3D; 2017-04-30.  | 
 **createdAtGte** | **string** | Datetime, in RFC 3339 format. Time zone defaults to UTC. Restricts results to those created on or after the specified time. e.g. &gt;&#x3D; 2017-04-30.  | 
 **minCreatedTime** | **int32** | Filter by results created after this unix timestamp | 
 **maxCreatedTime** | **int32** | Filter by results created before this unix timestamp | 
 **sort** | **string** | Method by which to order search results. Valid sorts are createdAt (created time, oldest first). Use :asc or :desc to specify ascending or descending order. Default is createdAt:asc.  | [default to &quot;createdAt:asc&quot;]
 **nextToken** | **string** | Token for pagination | 
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100. | [default to 50]
 **entityIds** | **string** | Filter by comma-separated list of related Entity IDs, maximum of 20. | 
 **storageIds** | **string** | Filter by comma-separated list of related storage (container, box, plate, or location) IDs, maximum of 20.  | 
 **assayRunIds** | **string** | Filter by comma-separated list of associated AssayRun IDs. | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 

### Return type

[**AssayResultsPaginatedList**](AssayResultsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveAssayResults

> AssayResultIdsResponse UnarchiveAssayResults(ctx).AssayResultIdsRequest(assayResultIdsRequest).Execute()

Unarchive 1 or more results.



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
    assayResultIdsRequest := *openapiclient.NewAssayResultIdsRequest([]string{"AssayResultIds_example"}) // AssayResultIdsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayResultsApi.UnarchiveAssayResults(context.Background()).AssayResultIdsRequest(assayResultIdsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayResultsApi.UnarchiveAssayResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveAssayResults`: AssayResultIdsResponse
    fmt.Fprintf(os.Stdout, "Response from `AssayResultsApi.UnarchiveAssayResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveAssayResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assayResultIdsRequest** | [**AssayResultIdsRequest**](AssayResultIdsRequest.md) |  | 

### Return type

[**AssayResultIdsResponse**](AssayResultIdsResponse.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

