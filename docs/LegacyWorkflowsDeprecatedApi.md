# \LegacyWorkflowsDeprecatedApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListStageRunInputSamples**](LegacyWorkflowsDeprecatedApi.md#ListStageRunInputSamples) | **Get** /workflow-stage-runs/{stage_run_id}/input-samples | List stage run input samples
[**ListStageRunOutputSamples**](LegacyWorkflowsDeprecatedApi.md#ListStageRunOutputSamples) | **Get** /workflow-stage-runs/{stage_run_id}/output-samples | List stage run output samples
[**ListStageRunRegisteredSamples**](LegacyWorkflowsDeprecatedApi.md#ListStageRunRegisteredSamples) | **Get** /workflow-stage-runs/{stage_run_id}/registered-samples | List stage run registered samples
[**ListWorkflowStageRuns**](LegacyWorkflowsDeprecatedApi.md#ListWorkflowStageRuns) | **Get** /workflow-stages/{stage_id}/workflow-stage-runs | List workflow stage runs
[**ListWorkflowStages**](LegacyWorkflowsDeprecatedApi.md#ListWorkflowStages) | **Get** /workflows/{workflow_id}/workflow-stages | List workflow stages
[**ListWorkflows**](LegacyWorkflowsDeprecatedApi.md#ListWorkflows) | **Get** /workflows | List workflows
[**UpdateWorkflowMetadata**](LegacyWorkflowsDeprecatedApi.md#UpdateWorkflowMetadata) | **Patch** /workflows/{workflow_id} | Update workflow



## ListStageRunInputSamples

> WorkflowSampleList ListStageRunInputSamples(ctx, stageRunId).Execute()

List stage run input samples



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
    stageRunId := "stageRunId_example" // string | ID of the stage run to list input samples for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyWorkflowsDeprecatedApi.ListStageRunInputSamples(context.Background(), stageRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWorkflowsDeprecatedApi.ListStageRunInputSamples``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStageRunInputSamples`: WorkflowSampleList
    fmt.Fprintf(os.Stdout, "Response from `LegacyWorkflowsDeprecatedApi.ListStageRunInputSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageRunId** | **string** | ID of the stage run to list input samples for | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStageRunInputSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowSampleList**](WorkflowSampleList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStageRunOutputSamples

> WorkflowSampleList ListStageRunOutputSamples(ctx, stageRunId).Execute()

List stage run output samples



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
    stageRunId := "stageRunId_example" // string | ID of the stage run to list output samples for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyWorkflowsDeprecatedApi.ListStageRunOutputSamples(context.Background(), stageRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWorkflowsDeprecatedApi.ListStageRunOutputSamples``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStageRunOutputSamples`: WorkflowSampleList
    fmt.Fprintf(os.Stdout, "Response from `LegacyWorkflowsDeprecatedApi.ListStageRunOutputSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageRunId** | **string** | ID of the stage run to list output samples for | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStageRunOutputSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowSampleList**](WorkflowSampleList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStageRunRegisteredSamples

> WorkflowSampleList ListStageRunRegisteredSamples(ctx, stageRunId).Execute()

List stage run registered samples



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
    stageRunId := "stageRunId_example" // string | ID of the stage run to list registered samples for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyWorkflowsDeprecatedApi.ListStageRunRegisteredSamples(context.Background(), stageRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWorkflowsDeprecatedApi.ListStageRunRegisteredSamples``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStageRunRegisteredSamples`: WorkflowSampleList
    fmt.Fprintf(os.Stdout, "Response from `LegacyWorkflowsDeprecatedApi.ListStageRunRegisteredSamples`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageRunId** | **string** | ID of the stage run to list registered samples for | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStageRunRegisteredSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowSampleList**](WorkflowSampleList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowStageRuns

> WorkflowStageRunList ListWorkflowStageRuns(ctx, stageId).Execute()

List workflow stage runs



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
    stageId := "stageId_example" // string | ID of the workflow stage to list runs for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyWorkflowsDeprecatedApi.ListWorkflowStageRuns(context.Background(), stageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWorkflowsDeprecatedApi.ListWorkflowStageRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowStageRuns`: WorkflowStageRunList
    fmt.Fprintf(os.Stdout, "Response from `LegacyWorkflowsDeprecatedApi.ListWorkflowStageRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageId** | **string** | ID of the workflow stage to list runs for | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowStageRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowStageRunList**](WorkflowStageRunList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowStages

> WorkflowStageList ListWorkflowStages(ctx, workflowId).Execute()

List workflow stages



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
    workflowId := "workflowId_example" // string | ID of the workflow to list stages for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyWorkflowsDeprecatedApi.ListWorkflowStages(context.Background(), workflowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWorkflowsDeprecatedApi.ListWorkflowStages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowStages`: WorkflowStageList
    fmt.Fprintf(os.Stdout, "Response from `LegacyWorkflowsDeprecatedApi.ListWorkflowStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | ID of the workflow to list stages for | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowStageList**](WorkflowStageList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflows

> WorkflowList ListWorkflows(ctx).Execute()

List workflows



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
    resp, r, err := apiClient.LegacyWorkflowsDeprecatedApi.ListWorkflows(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWorkflowsDeprecatedApi.ListWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflows`: WorkflowList
    fmt.Fprintf(os.Stdout, "Response from `LegacyWorkflowsDeprecatedApi.ListWorkflows`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowsRequest struct via the builder pattern


### Return type

[**WorkflowList**](WorkflowList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowMetadata

> LegacyWorkflow UpdateWorkflowMetadata(ctx, workflowId).WorkflowPatch(workflowPatch).Execute()

Update workflow



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
    workflowId := "workflowId_example" // string | ID of the workflow to update
    workflowPatch := *openapiclient.NewWorkflowPatch() // WorkflowPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LegacyWorkflowsDeprecatedApi.UpdateWorkflowMetadata(context.Background(), workflowId).WorkflowPatch(workflowPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LegacyWorkflowsDeprecatedApi.UpdateWorkflowMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflowMetadata`: LegacyWorkflow
    fmt.Fprintf(os.Stdout, "Response from `LegacyWorkflowsDeprecatedApi.UpdateWorkflowMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | ID of the workflow to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowPatch** | [**WorkflowPatch**](WorkflowPatch.md) |  | 

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

