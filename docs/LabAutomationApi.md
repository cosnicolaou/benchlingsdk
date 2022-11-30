# \LabAutomationApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveAutomationOutputProcessors**](LabAutomationApi.md#ArchiveAutomationOutputProcessors) | **Post** /automation-output-processors:archive | Archive Automation Output Processors and linked Results
[**CreateAutomationOutputProcessor**](LabAutomationApi.md#CreateAutomationOutputProcessor) | **Post** /automation-output-processors | Create Automation Output Processor
[**GenerateInputWithAutomationInputGenerator**](LabAutomationApi.md#GenerateInputWithAutomationInputGenerator) | **Post** /automation-input-generators/{input_generator_id}:generate-input | Generate Input with an Automation Input Generator
[**GetAutomationInputGenerator**](LabAutomationApi.md#GetAutomationInputGenerator) | **Get** /automation-input-generators/{input_generator_id} | Get an Automation Input Generator
[**GetAutomationOutputProcessor**](LabAutomationApi.md#GetAutomationOutputProcessor) | **Get** /automation-output-processors/{output_processor_id} | Get an Automation Output Processor
[**GetLabAutomationTransform**](LabAutomationApi.md#GetLabAutomationTransform) | **Get** /automation-file-transforms/{transform_id} | Get a Lab Automation Transform step
[**ListAutomationOutputProcessors**](LabAutomationApi.md#ListAutomationOutputProcessors) | **Get** /automation-output-processors | List non-empty Automation Output Processors
[**ProcessOutputWithAutomationOutputProcessor**](LabAutomationApi.md#ProcessOutputWithAutomationOutputProcessor) | **Post** /automation-output-processors/{output_processor_id}:process-output | Process Output with an Automation Output Processor
[**UnarchiveAutomationOutputProcessors**](LabAutomationApi.md#UnarchiveAutomationOutputProcessors) | **Post** /automation-output-processors:unarchive | Unarchive Automation Output Processors and linked Results
[**UpdateAutomationInputGenerator**](LabAutomationApi.md#UpdateAutomationInputGenerator) | **Patch** /automation-input-generators/{input_generator_id} | Update an Automation Input Generator
[**UpdateAutomationOutputProcessor**](LabAutomationApi.md#UpdateAutomationOutputProcessor) | **Patch** /automation-output-processors/{output_processor_id} | Update an Automation Output Processor
[**UpdateLabAutomationTransform**](LabAutomationApi.md#UpdateLabAutomationTransform) | **Patch** /automation-file-transforms/{transform_id} | Update a Lab Automation Transform step



## ArchiveAutomationOutputProcessors

> AutomationOutputProcessorArchivalChange ArchiveAutomationOutputProcessors(ctx).AutomationOutputProcessorsArchive(automationOutputProcessorsArchive).Execute()

Archive Automation Output Processors and linked Results



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
    automationOutputProcessorsArchive := *openapiclient.NewAutomationOutputProcessorsArchive([]string{"AutomationOutputProcessorIds_example"}) // AutomationOutputProcessorsArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabAutomationApi.ArchiveAutomationOutputProcessors(context.Background()).AutomationOutputProcessorsArchive(automationOutputProcessorsArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabAutomationApi.ArchiveAutomationOutputProcessors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveAutomationOutputProcessors`: AutomationOutputProcessorArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `LabAutomationApi.ArchiveAutomationOutputProcessors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveAutomationOutputProcessorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **automationOutputProcessorsArchive** | [**AutomationOutputProcessorsArchive**](AutomationOutputProcessorsArchive.md) |  | 

### Return type

[**AutomationOutputProcessorArchivalChange**](AutomationOutputProcessorArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAutomationOutputProcessor

> AutomationOutputProcessor CreateAutomationOutputProcessor(ctx).AutomationOutputProcessorCreate(automationOutputProcessorCreate).Execute()

Create Automation Output Processor



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
    automationOutputProcessorCreate := *openapiclient.NewAutomationOutputProcessorCreate("AssayRunId_example", "AutomationFileConfigName_example", "cd624536-c6ba-41b9-b802-9461689e2ea3") // AutomationOutputProcessorCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabAutomationApi.CreateAutomationOutputProcessor(context.Background()).AutomationOutputProcessorCreate(automationOutputProcessorCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabAutomationApi.CreateAutomationOutputProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAutomationOutputProcessor`: AutomationOutputProcessor
    fmt.Fprintf(os.Stdout, "Response from `LabAutomationApi.CreateAutomationOutputProcessor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutomationOutputProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **automationOutputProcessorCreate** | [**AutomationOutputProcessorCreate**](AutomationOutputProcessorCreate.md) |  | 

### Return type

[**AutomationOutputProcessor**](AutomationOutputProcessor.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateInputWithAutomationInputGenerator

> AsyncTaskLink GenerateInputWithAutomationInputGenerator(ctx, inputGeneratorId).Execute()

Generate Input with an Automation Input Generator



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
    inputGeneratorId := "inputGeneratorId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabAutomationApi.GenerateInputWithAutomationInputGenerator(context.Background(), inputGeneratorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabAutomationApi.GenerateInputWithAutomationInputGenerator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateInputWithAutomationInputGenerator`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `LabAutomationApi.GenerateInputWithAutomationInputGenerator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inputGeneratorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateInputWithAutomationInputGeneratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncTaskLink**](AsyncTaskLink.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationInputGenerator

> AutomationInputGenerator GetAutomationInputGenerator(ctx, inputGeneratorId).Execute()

Get an Automation Input Generator



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
    inputGeneratorId := "inputGeneratorId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabAutomationApi.GetAutomationInputGenerator(context.Background(), inputGeneratorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabAutomationApi.GetAutomationInputGenerator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationInputGenerator`: AutomationInputGenerator
    fmt.Fprintf(os.Stdout, "Response from `LabAutomationApi.GetAutomationInputGenerator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inputGeneratorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationInputGeneratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutomationInputGenerator**](AutomationInputGenerator.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationOutputProcessor

> AutomationOutputProcessor GetAutomationOutputProcessor(ctx, outputProcessorId).Execute()

Get an Automation Output Processor



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
    outputProcessorId := "outputProcessorId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabAutomationApi.GetAutomationOutputProcessor(context.Background(), outputProcessorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabAutomationApi.GetAutomationOutputProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationOutputProcessor`: AutomationOutputProcessor
    fmt.Fprintf(os.Stdout, "Response from `LabAutomationApi.GetAutomationOutputProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outputProcessorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationOutputProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutomationOutputProcessor**](AutomationOutputProcessor.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabAutomationTransform

> LabAutomationTransform GetLabAutomationTransform(ctx, transformId).Execute()

Get a Lab Automation Transform step



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
    transformId := "transformId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabAutomationApi.GetLabAutomationTransform(context.Background(), transformId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabAutomationApi.GetLabAutomationTransform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLabAutomationTransform`: LabAutomationTransform
    fmt.Fprintf(os.Stdout, "Response from `LabAutomationApi.GetLabAutomationTransform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transformId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabAutomationTransformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LabAutomationTransform**](LabAutomationTransform.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAutomationOutputProcessors

> AutomationOutputProcessorsPaginatedList ListAutomationOutputProcessors(ctx).AssayRunId(assayRunId).AutomationFileConfigName(automationFileConfigName).ArchiveReason(archiveReason).ModifiedAt(modifiedAt).NextToken(nextToken).Execute()

List non-empty Automation Output Processors



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
    assayRunId := "assayRunId_example" // string | Id of the Run (optional)
    automationFileConfigName := "automationFileConfigName_example" // string | Name of the Automation File Config (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived processors. Use \"ANY_ARCHIVED\" to filter for archived processors regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    nextToken := "nextToken_example" // string | Token for pagination (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabAutomationApi.ListAutomationOutputProcessors(context.Background()).AssayRunId(assayRunId).AutomationFileConfigName(automationFileConfigName).ArchiveReason(archiveReason).ModifiedAt(modifiedAt).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabAutomationApi.ListAutomationOutputProcessors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAutomationOutputProcessors`: AutomationOutputProcessorsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `LabAutomationApi.ListAutomationOutputProcessors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAutomationOutputProcessorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assayRunId** | **string** | Id of the Run | 
 **automationFileConfigName** | **string** | Name of the Automation File Config | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived processors. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived processors regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **nextToken** | **string** | Token for pagination | 

### Return type

[**AutomationOutputProcessorsPaginatedList**](AutomationOutputProcessorsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessOutputWithAutomationOutputProcessor

> AsyncTaskLink ProcessOutputWithAutomationOutputProcessor(ctx, outputProcessorId).Execute()

Process Output with an Automation Output Processor



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
    outputProcessorId := "outputProcessorId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabAutomationApi.ProcessOutputWithAutomationOutputProcessor(context.Background(), outputProcessorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabAutomationApi.ProcessOutputWithAutomationOutputProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessOutputWithAutomationOutputProcessor`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `LabAutomationApi.ProcessOutputWithAutomationOutputProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outputProcessorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessOutputWithAutomationOutputProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncTaskLink**](AsyncTaskLink.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveAutomationOutputProcessors

> AutomationOutputProcessorArchivalChange UnarchiveAutomationOutputProcessors(ctx).AutomationOutputProcessorsUnarchive(automationOutputProcessorsUnarchive).Execute()

Unarchive Automation Output Processors and linked Results



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
    automationOutputProcessorsUnarchive := *openapiclient.NewAutomationOutputProcessorsUnarchive([]string{"AutomationOutputProcessorIds_example"}) // AutomationOutputProcessorsUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabAutomationApi.UnarchiveAutomationOutputProcessors(context.Background()).AutomationOutputProcessorsUnarchive(automationOutputProcessorsUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabAutomationApi.UnarchiveAutomationOutputProcessors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveAutomationOutputProcessors`: AutomationOutputProcessorArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `LabAutomationApi.UnarchiveAutomationOutputProcessors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveAutomationOutputProcessorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **automationOutputProcessorsUnarchive** | [**AutomationOutputProcessorsUnarchive**](AutomationOutputProcessorsUnarchive.md) |  | 

### Return type

[**AutomationOutputProcessorArchivalChange**](AutomationOutputProcessorArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutomationInputGenerator

> AutomationInputGenerator UpdateAutomationInputGenerator(ctx, inputGeneratorId).AutomationInputGeneratorUpdate(automationInputGeneratorUpdate).Execute()

Update an Automation Input Generator



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
    inputGeneratorId := "inputGeneratorId_example" // string | 
    automationInputGeneratorUpdate := *openapiclient.NewAutomationInputGeneratorUpdate() // AutomationInputGeneratorUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabAutomationApi.UpdateAutomationInputGenerator(context.Background(), inputGeneratorId).AutomationInputGeneratorUpdate(automationInputGeneratorUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabAutomationApi.UpdateAutomationInputGenerator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAutomationInputGenerator`: AutomationInputGenerator
    fmt.Fprintf(os.Stdout, "Response from `LabAutomationApi.UpdateAutomationInputGenerator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inputGeneratorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutomationInputGeneratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **automationInputGeneratorUpdate** | [**AutomationInputGeneratorUpdate**](AutomationInputGeneratorUpdate.md) |  | 

### Return type

[**AutomationInputGenerator**](AutomationInputGenerator.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutomationOutputProcessor

> AutomationOutputProcessor UpdateAutomationOutputProcessor(ctx, outputProcessorId).AutomationOutputProcessorUpdate(automationOutputProcessorUpdate).Execute()

Update an Automation Output Processor



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
    outputProcessorId := "outputProcessorId_example" // string | 
    automationOutputProcessorUpdate := *openapiclient.NewAutomationOutputProcessorUpdate("cd624536-c6ba-41b9-b802-9461689e2ea3") // AutomationOutputProcessorUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabAutomationApi.UpdateAutomationOutputProcessor(context.Background(), outputProcessorId).AutomationOutputProcessorUpdate(automationOutputProcessorUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabAutomationApi.UpdateAutomationOutputProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAutomationOutputProcessor`: AutomationOutputProcessor
    fmt.Fprintf(os.Stdout, "Response from `LabAutomationApi.UpdateAutomationOutputProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outputProcessorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutomationOutputProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **automationOutputProcessorUpdate** | [**AutomationOutputProcessorUpdate**](AutomationOutputProcessorUpdate.md) |  | 

### Return type

[**AutomationOutputProcessor**](AutomationOutputProcessor.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLabAutomationTransform

> LabAutomationTransform UpdateLabAutomationTransform(ctx, transformId).LabAutomationTransformUpdate(labAutomationTransformUpdate).Execute()

Update a Lab Automation Transform step



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
    transformId := "transformId_example" // string | 
    labAutomationTransformUpdate := *openapiclient.NewLabAutomationTransformUpdate() // LabAutomationTransformUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabAutomationApi.UpdateLabAutomationTransform(context.Background(), transformId).LabAutomationTransformUpdate(labAutomationTransformUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabAutomationApi.UpdateLabAutomationTransform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLabAutomationTransform`: LabAutomationTransform
    fmt.Fprintf(os.Stdout, "Response from `LabAutomationApi.UpdateLabAutomationTransform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transformId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLabAutomationTransformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labAutomationTransformUpdate** | [**LabAutomationTransformUpdate**](LabAutomationTransformUpdate.md) |  | 

### Return type

[**LabAutomationTransform**](LabAutomationTransform.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

