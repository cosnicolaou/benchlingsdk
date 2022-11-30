# \LegacyWorkflowsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLegacyWorkflowStageRunInputSamples**](LegacyWorkflowsApi.md#ListLegacyWorkflowStageRunInputSamples) | **Get** /legacy-workflow-stage-runs/{stage_run_id}/input-samples | List legacy workflow stage run input samples
[**ListLegacyWorkflowStageRunOutputSamples**](LegacyWorkflowsApi.md#ListLegacyWorkflowStageRunOutputSamples) | **Get** /legacy-workflow-stage-runs/{stage_run_id}/output-samples | List legacy workflow stage run output samples
[**ListLegacyWorkflowStageRunRegisteredSamples**](LegacyWorkflowsApi.md#ListLegacyWorkflowStageRunRegisteredSamples) | **Get** /legacy-workflow-stage-runs/{stage_run_id}/registered-samples | List legacy workflow stage run registered samples
[**ListLegacyWorkflowStageRuns**](LegacyWorkflowsApi.md#ListLegacyWorkflowStageRuns) | **Get** /legacy-workflow-stages/{stage_id}/workflow-stage-runs | List legacy workflow stage runs
[**ListLegacyWorkflowStages**](LegacyWorkflowsApi.md#ListLegacyWorkflowStages) | **Get** /legacy-workflows/{legacy_workflow_id}/workflow-stages | List legacy workflow stages
[**ListLegacyWorkflows**](LegacyWorkflowsApi.md#ListLegacyWorkflows) | **Get** /legacy-workflows | List legacy workflows
[**UpdateLegacyWorkflowMetadata**](LegacyWorkflowsApi.md#UpdateLegacyWorkflowMetadata) | **Patch** /legacy-workflows/{legacy_workflow_id} | Update legacy workflow



## ListLegacyWorkflowStageRunInputSamples

> LegacyWorkflowSampleList ListLegacyWorkflowStageRunInputSamples(ctx, stageRunId).Execute()

List legacy workflow stage run input samples



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
    stageRunId := "stageRunId_example" // string | ID of the legacy workflow stage run to list input samples for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyWorkflowsApi.ListLegacyWorkflowStageRunInputSamples(context.Background(), stageRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWorkflowsApi.ListLegacyWorkflowStageRunInputSamples``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLegacyWorkflowStageRunInputSamples`: LegacyWorkflowSampleList
    fmt.Fprintf(os.Stdout, "Response from `LegacyWorkflowsApi.ListLegacyWorkflowStageRunInputSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageRunId** | **string** | ID of the legacy workflow stage run to list input samples for | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLegacyWorkflowStageRunInputSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LegacyWorkflowSampleList**](LegacyWorkflowSampleList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLegacyWorkflowStageRunOutputSamples

> LegacyWorkflowSampleList ListLegacyWorkflowStageRunOutputSamples(ctx, stageRunId).Execute()

List legacy workflow stage run output samples



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
    stageRunId := "stageRunId_example" // string | ID of the legacy workflow stage run to list output samples for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyWorkflowsApi.ListLegacyWorkflowStageRunOutputSamples(context.Background(), stageRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWorkflowsApi.ListLegacyWorkflowStageRunOutputSamples``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLegacyWorkflowStageRunOutputSamples`: LegacyWorkflowSampleList
    fmt.Fprintf(os.Stdout, "Response from `LegacyWorkflowsApi.ListLegacyWorkflowStageRunOutputSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageRunId** | **string** | ID of the legacy workflow stage run to list output samples for | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLegacyWorkflowStageRunOutputSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LegacyWorkflowSampleList**](LegacyWorkflowSampleList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLegacyWorkflowStageRunRegisteredSamples

> LegacyWorkflowSampleList ListLegacyWorkflowStageRunRegisteredSamples(ctx, stageRunId).Execute()

List legacy workflow stage run registered samples



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
    stageRunId := "stageRunId_example" // string | ID of the legacy workflow stage run to list registered samples for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyWorkflowsApi.ListLegacyWorkflowStageRunRegisteredSamples(context.Background(), stageRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWorkflowsApi.ListLegacyWorkflowStageRunRegisteredSamples``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLegacyWorkflowStageRunRegisteredSamples`: LegacyWorkflowSampleList
    fmt.Fprintf(os.Stdout, "Response from `LegacyWorkflowsApi.ListLegacyWorkflowStageRunRegisteredSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageRunId** | **string** | ID of the legacy workflow stage run to list registered samples for | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLegacyWorkflowStageRunRegisteredSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LegacyWorkflowSampleList**](LegacyWorkflowSampleList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLegacyWorkflowStageRuns

> LegacyWorkflowStageRunList ListLegacyWorkflowStageRuns(ctx, stageId).Execute()

List legacy workflow stage runs



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
    stageId := "stageId_example" // string | ID of the legacy workflow stage to list runs for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyWorkflowsApi.ListLegacyWorkflowStageRuns(context.Background(), stageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWorkflowsApi.ListLegacyWorkflowStageRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLegacyWorkflowStageRuns`: LegacyWorkflowStageRunList
    fmt.Fprintf(os.Stdout, "Response from `LegacyWorkflowsApi.ListLegacyWorkflowStageRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageId** | **string** | ID of the legacy workflow stage to list runs for | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLegacyWorkflowStageRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LegacyWorkflowStageRunList**](LegacyWorkflowStageRunList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLegacyWorkflowStages

> LegacyWorkflowStageList ListLegacyWorkflowStages(ctx, legacyWorkflowId).Execute()

List legacy workflow stages



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
    legacyWorkflowId := "legacyWorkflowId_example" // string | ID of the legacy workflow to list stages for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyWorkflowsApi.ListLegacyWorkflowStages(context.Background(), legacyWorkflowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWorkflowsApi.ListLegacyWorkflowStages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLegacyWorkflowStages`: LegacyWorkflowStageList
    fmt.Fprintf(os.Stdout, "Response from `LegacyWorkflowsApi.ListLegacyWorkflowStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**legacyWorkflowId** | **string** | ID of the legacy workflow to list stages for | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLegacyWorkflowStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LegacyWorkflowStageList**](LegacyWorkflowStageList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLegacyWorkflows

> LegacyWorkflowList ListLegacyWorkflows(ctx).Execute()

List legacy workflows



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
    resp, r, err := apiClient.LegacyWorkflowsApi.ListLegacyWorkflows(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWorkflowsApi.ListLegacyWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLegacyWorkflows`: LegacyWorkflowList
    fmt.Fprintf(os.Stdout, "Response from `LegacyWorkflowsApi.ListLegacyWorkflows`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLegacyWorkflowsRequest struct via the builder pattern


### Return type

[**LegacyWorkflowList**](LegacyWorkflowList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLegacyWorkflowMetadata

> LegacyWorkflow UpdateLegacyWorkflowMetadata(ctx, legacyWorkflowId).LegacyWorkflowPatch(legacyWorkflowPatch).Execute()

Update legacy workflow



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
    legacyWorkflowId := "legacyWorkflowId_example" // string | ID of the legacy workflow to update
    legacyWorkflowPatch := *openapiclient.NewLegacyWorkflowPatch() // LegacyWorkflowPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyWorkflowsApi.UpdateLegacyWorkflowMetadata(context.Background(), legacyWorkflowId).LegacyWorkflowPatch(legacyWorkflowPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWorkflowsApi.UpdateLegacyWorkflowMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLegacyWorkflowMetadata`: LegacyWorkflow
    fmt.Fprintf(os.Stdout, "Response from `LegacyWorkflowsApi.UpdateLegacyWorkflowMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**legacyWorkflowId** | **string** | ID of the legacy workflow to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLegacyWorkflowMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **legacyWorkflowPatch** | [**LegacyWorkflowPatch**](LegacyWorkflowPatch.md) |  | 

### Return type

[**LegacyWorkflow**](LegacyWorkflow.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

