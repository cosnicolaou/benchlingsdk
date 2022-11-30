# \AssayRunsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveAssayRuns**](AssayRunsApi.md#ArchiveAssayRuns) | **Post** /assay-runs:archive | Archive Assay Runs
[**BulkGetAssayRuns**](AssayRunsApi.md#BulkGetAssayRuns) | **Get** /assay-runs:bulk-get | Bulk get runs by ID
[**CreateAssayRuns**](AssayRunsApi.md#CreateAssayRuns) | **Post** /assay-runs | Create 1 or more runs.
[**GetAssayRun**](AssayRunsApi.md#GetAssayRun) | **Get** /assay-runs/{assay_run_id} | Get a run
[**ListAssayRuns**](AssayRunsApi.md#ListAssayRuns) | **Get** /assay-runs | List runs
[**ListAutomationInputGenerators**](AssayRunsApi.md#ListAutomationInputGenerators) | **Get** /assay-runs/{assay_run_id}/automation-input-generators | list AutomationInputGenerators by Run
[**ListAutomationOutputProcessorsDeprecated**](AssayRunsApi.md#ListAutomationOutputProcessorsDeprecated) | **Get** /assay-runs/{assay_run_id}/automation-output-processors | list AutomationOutputProcessors by Run
[**UnarchiveAssayRuns**](AssayRunsApi.md#UnarchiveAssayRuns) | **Post** /assay-runs:unarchive | Unarchive Assay Runs
[**UpdateAssayRun**](AssayRunsApi.md#UpdateAssayRun) | **Patch** /assay-runs/{assay_run_id} | Update a run



## ArchiveAssayRuns

> AssayRunsArchivalChange ArchiveAssayRuns(ctx).AssayRunsArchive(assayRunsArchive).Execute()

Archive Assay Runs



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
    assayRunsArchive := *openapiclient.NewAssayRunsArchive([]string{"AssayRunIds_example"}, "Reason_example") // AssayRunsArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayRunsApi.ArchiveAssayRuns(context.Background()).AssayRunsArchive(assayRunsArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayRunsApi.ArchiveAssayRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveAssayRuns`: AssayRunsArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `AssayRunsApi.ArchiveAssayRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveAssayRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assayRunsArchive** | [**AssayRunsArchive**](AssayRunsArchive.md) |  | 

### Return type

[**AssayRunsArchivalChange**](AssayRunsArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkGetAssayRuns

> AssayRunsBulkGet BulkGetAssayRuns(ctx).AssayRunIds(assayRunIds).Execute()

Bulk get runs by ID



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
    assayRunIds := "assayRunIds_example" // string | Comma-separated list of assay run IDs

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayRunsApi.BulkGetAssayRuns(context.Background()).AssayRunIds(assayRunIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayRunsApi.BulkGetAssayRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetAssayRuns`: AssayRunsBulkGet
    fmt.Fprintf(os.Stdout, "Response from `AssayRunsApi.BulkGetAssayRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetAssayRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assayRunIds** | **string** | Comma-separated list of assay run IDs | 

### Return type

[**AssayRunsBulkGet**](AssayRunsBulkGet.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssayRuns

> AssayRunsBulkCreateResponse CreateAssayRuns(ctx).AssayRunsBulkCreateRequest(assayRunsBulkCreateRequest).Execute()

Create 1 or more runs.



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
    assayRunsBulkCreateRequest := *openapiclient.NewAssayRunsBulkCreateRequest([]openapiclient.AssayRunCreate{*openapiclient.NewAssayRunCreate(*openapiclient.NewAssayRunCreateFields(), "SchemaId_example")}) // AssayRunsBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayRunsApi.CreateAssayRuns(context.Background()).AssayRunsBulkCreateRequest(assayRunsBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayRunsApi.CreateAssayRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssayRuns`: AssayRunsBulkCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `AssayRunsApi.CreateAssayRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssayRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assayRunsBulkCreateRequest** | [**AssayRunsBulkCreateRequest**](AssayRunsBulkCreateRequest.md) |  | 

### Return type

[**AssayRunsBulkCreateResponse**](AssayRunsBulkCreateResponse.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssayRun

> AssayRun GetAssayRun(ctx, assayRunId).Execute()

Get a run



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
    assayRunId := "assayRunId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayRunsApi.GetAssayRun(context.Background(), assayRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayRunsApi.GetAssayRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssayRun`: AssayRun
    fmt.Fprintf(os.Stdout, "Response from `AssayRunsApi.GetAssayRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assayRunId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssayRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssayRun**](AssayRun.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssayRuns

> AssayRunsPaginatedList ListAssayRuns(ctx).SchemaId(schemaId).MinCreatedTime(minCreatedTime).MaxCreatedTime(maxCreatedTime).NextToken(nextToken).PageSize(pageSize).Ids(ids).Execute()

List runs



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
    schemaId := "schemaId_example" // string | ID of the assay run schema to filter by
    minCreatedTime := int32(56) // int32 | Filter by runs created after this unix timestamp (optional)
    maxCreatedTime := int32(56) // int32 | Filter by runs created before this unix timestamp (optional)
    nextToken := "nextToken_example" // string | Token for pagination (optional)
    pageSize := int32(56) // int32 | Number of results to return. Defaults to 50, maximum of 100. (optional) (default to 50)
    ids := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayRunsApi.ListAssayRuns(context.Background()).SchemaId(schemaId).MinCreatedTime(minCreatedTime).MaxCreatedTime(maxCreatedTime).NextToken(nextToken).PageSize(pageSize).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayRunsApi.ListAssayRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAssayRuns`: AssayRunsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `AssayRunsApi.ListAssayRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssayRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaId** | **string** | ID of the assay run schema to filter by | 
 **minCreatedTime** | **int32** | Filter by runs created after this unix timestamp | 
 **maxCreatedTime** | **int32** | Filter by runs created before this unix timestamp | 
 **nextToken** | **string** | Token for pagination | 
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100. | [default to 50]
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 

### Return type

[**AssayRunsPaginatedList**](AssayRunsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAutomationInputGenerators

> AutomationFileInputsPaginatedList ListAutomationInputGenerators(ctx, assayRunId).NextToken(nextToken).ModifiedAt(modifiedAt).Execute()

list AutomationInputGenerators by Run



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
    assayRunId := "assayRunId_example" // string | 
    nextToken := "nextToken_example" // string | Token for pagination (optional)
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayRunsApi.ListAutomationInputGenerators(context.Background(), assayRunId).NextToken(nextToken).ModifiedAt(modifiedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayRunsApi.ListAutomationInputGenerators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAutomationInputGenerators`: AutomationFileInputsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `AssayRunsApi.ListAutomationInputGenerators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assayRunId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAutomationInputGeneratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextToken** | **string** | Token for pagination | 
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 

### Return type

[**AutomationFileInputsPaginatedList**](AutomationFileInputsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAutomationOutputProcessorsDeprecated

> DeprecatedAutomationOutputProcessorsPaginatedList ListAutomationOutputProcessorsDeprecated(ctx, assayRunId).NextToken(nextToken).Execute()

list AutomationOutputProcessors by Run



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
    assayRunId := "assayRunId_example" // string | 
    nextToken := "nextToken_example" // string | Token for pagination (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayRunsApi.ListAutomationOutputProcessorsDeprecated(context.Background(), assayRunId).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayRunsApi.ListAutomationOutputProcessorsDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAutomationOutputProcessorsDeprecated`: DeprecatedAutomationOutputProcessorsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `AssayRunsApi.ListAutomationOutputProcessorsDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assayRunId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAutomationOutputProcessorsDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextToken** | **string** | Token for pagination | 

### Return type

[**DeprecatedAutomationOutputProcessorsPaginatedList**](DeprecatedAutomationOutputProcessorsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveAssayRuns

> AssayRunsArchivalChange UnarchiveAssayRuns(ctx).AssayRunsUnarchive(assayRunsUnarchive).Execute()

Unarchive Assay Runs



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
    assayRunsUnarchive := *openapiclient.NewAssayRunsUnarchive([]string{"AssayRunIds_example"}) // AssayRunsUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayRunsApi.UnarchiveAssayRuns(context.Background()).AssayRunsUnarchive(assayRunsUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayRunsApi.UnarchiveAssayRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveAssayRuns`: AssayRunsArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `AssayRunsApi.UnarchiveAssayRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveAssayRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assayRunsUnarchive** | [**AssayRunsUnarchive**](AssayRunsUnarchive.md) |  | 

### Return type

[**AssayRunsArchivalChange**](AssayRunsArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssayRun

> AssayRun UpdateAssayRun(ctx, assayRunId).AssayRunUpdate(assayRunUpdate).Execute()

Update a run



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
    assayRunId := "assayRunId_example" // string | ID of the Run to be updated
    assayRunUpdate := *openapiclient.NewAssayRunUpdate() // AssayRunUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssayRunsApi.UpdateAssayRun(context.Background(), assayRunId).AssayRunUpdate(assayRunUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssayRunsApi.UpdateAssayRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssayRun`: AssayRun
    fmt.Fprintf(os.Stdout, "Response from `AssayRunsApi.UpdateAssayRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assayRunId** | **string** | ID of the Run to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssayRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assayRunUpdate** | [**AssayRunUpdate**](AssayRunUpdate.md) |  | 

### Return type

[**AssayRun**](AssayRun.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

