# \WorkflowOutputsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveWorkflowOutputs**](WorkflowOutputsApi.md#ArchiveWorkflowOutputs) | **Post** /workflow-outputs:archive | Archive one or more workflow outputs
[**BulkCreateWorkflowOutputs**](WorkflowOutputsApi.md#BulkCreateWorkflowOutputs) | **Post** /workflow-outputs:bulk-create | Bulk create new workflow outputs
[**BulkUpdateWorkflowOutputs**](WorkflowOutputsApi.md#BulkUpdateWorkflowOutputs) | **Post** /workflow-outputs:bulk-update | Bulk update workflow outputs
[**CreateWorkflowOutput**](WorkflowOutputsApi.md#CreateWorkflowOutput) | **Post** /workflow-outputs | Create a new workflow output
[**GetWorkflowOutput**](WorkflowOutputsApi.md#GetWorkflowOutput) | **Get** /workflow-outputs/{workflow_output_id} | Get a workflow output
[**ListWorkflowOutputs**](WorkflowOutputsApi.md#ListWorkflowOutputs) | **Get** /workflow-outputs | List workflow outputs
[**UnarchiveWorkflowOutputs**](WorkflowOutputsApi.md#UnarchiveWorkflowOutputs) | **Post** /workflow-outputs:unarchive | Unarchive one or more workflow outputs
[**UpdateWorkflowOutput**](WorkflowOutputsApi.md#UpdateWorkflowOutput) | **Patch** /workflow-outputs/{workflow_output_id} | Update a workflow output



## ArchiveWorkflowOutputs

> WorkflowOutputsArchivalChange ArchiveWorkflowOutputs(ctx).WorkflowOutputsArchive(workflowOutputsArchive).Execute()

Archive one or more workflow outputs



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
    workflowOutputsArchive := *openapiclient.NewWorkflowOutputsArchive(openapiclient.WorkflowOutputArchiveReason("Made in error"), []string{"wfout_5cJLQKVF"}) // WorkflowOutputsArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowOutputsApi.ArchiveWorkflowOutputs(context.Background()).WorkflowOutputsArchive(workflowOutputsArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowOutputsApi.ArchiveWorkflowOutputs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveWorkflowOutputs`: WorkflowOutputsArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `WorkflowOutputsApi.ArchiveWorkflowOutputs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveWorkflowOutputsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowOutputsArchive** | [**WorkflowOutputsArchive**](WorkflowOutputsArchive.md) |  | 

### Return type

[**WorkflowOutputsArchivalChange**](WorkflowOutputsArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkCreateWorkflowOutputs

> AsyncTaskLink BulkCreateWorkflowOutputs(ctx).WorkflowOutputsBulkCreateRequest(workflowOutputsBulkCreateRequest).Execute()

Bulk create new workflow outputs



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
    workflowOutputsBulkCreateRequest := *openapiclient.NewWorkflowOutputsBulkCreateRequest() // WorkflowOutputsBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowOutputsApi.BulkCreateWorkflowOutputs(context.Background()).WorkflowOutputsBulkCreateRequest(workflowOutputsBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowOutputsApi.BulkCreateWorkflowOutputs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateWorkflowOutputs`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `WorkflowOutputsApi.BulkCreateWorkflowOutputs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateWorkflowOutputsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowOutputsBulkCreateRequest** | [**WorkflowOutputsBulkCreateRequest**](WorkflowOutputsBulkCreateRequest.md) |  | 

### Return type

[**AsyncTaskLink**](AsyncTaskLink.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateWorkflowOutputs

> AsyncTaskLink BulkUpdateWorkflowOutputs(ctx).WorkflowOutputsBulkUpdateRequest(workflowOutputsBulkUpdateRequest).Execute()

Bulk update workflow outputs



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
    workflowOutputsBulkUpdateRequest := *openapiclient.NewWorkflowOutputsBulkUpdateRequest() // WorkflowOutputsBulkUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowOutputsApi.BulkUpdateWorkflowOutputs(context.Background()).WorkflowOutputsBulkUpdateRequest(workflowOutputsBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowOutputsApi.BulkUpdateWorkflowOutputs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateWorkflowOutputs`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `WorkflowOutputsApi.BulkUpdateWorkflowOutputs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateWorkflowOutputsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowOutputsBulkUpdateRequest** | [**WorkflowOutputsBulkUpdateRequest**](WorkflowOutputsBulkUpdateRequest.md) |  | 

### Return type

[**AsyncTaskLink**](AsyncTaskLink.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflowOutput

> WorkflowOutput CreateWorkflowOutput(ctx).WorkflowOutputCreate(workflowOutputCreate).Execute()

Create a new workflow output



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
    workflowOutputCreate := *openapiclient.NewWorkflowOutputCreate("wftask_OnnsW08k") // WorkflowOutputCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowOutputsApi.CreateWorkflowOutput(context.Background()).WorkflowOutputCreate(workflowOutputCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowOutputsApi.CreateWorkflowOutput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowOutput`: WorkflowOutput
    fmt.Fprintf(os.Stdout, "Response from `WorkflowOutputsApi.CreateWorkflowOutput`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowOutputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowOutputCreate** | [**WorkflowOutputCreate**](WorkflowOutputCreate.md) |  | 

### Return type

[**WorkflowOutput**](WorkflowOutput.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowOutput

> WorkflowOutput GetWorkflowOutput(ctx, workflowOutputId).Execute()

Get a workflow output



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
    workflowOutputId := "wfout_5cJLQKVF" // string | The ID of the workflow task output

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowOutputsApi.GetWorkflowOutput(context.Background(), workflowOutputId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowOutputsApi.GetWorkflowOutput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowOutput`: WorkflowOutput
    fmt.Fprintf(os.Stdout, "Response from `WorkflowOutputsApi.GetWorkflowOutput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowOutputId** | **string** | The ID of the workflow task output | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowOutputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowOutput**](WorkflowOutput.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowOutputs

> WorkflowOutputsPaginatedList ListWorkflowOutputs(ctx).Ids(ids).WorkflowTaskGroupIds(workflowTaskGroupIds).WorkflowTaskIds(workflowTaskIds).SchemaId(schemaId).WatcherIds(watcherIds).ResponsibleTeamIds(responsibleTeamIds).CreationOriginIds(creationOriginIds).LinkedItemIdsAnyOf(linkedItemIdsAnyOf).LinkedItemIdsAllOf(linkedItemIdsAllOf).LinkedItemIdsNoneOf(linkedItemIdsNoneOf).SchemaFields(schemaFields).Name(name).NameIncludes(nameIncludes).CreatorIds(creatorIds).ModifiedAt(modifiedAt).NextToken(nextToken).PageSize(pageSize).DisplayIds(displayIds).ArchiveReason(archiveReason).Execute()

List workflow outputs



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
    ids := "wfout_5cJLQKVF,wfout_9jENXd3t" // string | Comma separated list of workflow output IDs (optional)
    workflowTaskGroupIds := "prs_giVNQcTL,prst6m99v1" // string | Comma separated list of workflow IDs (optional)
    workflowTaskIds := "wftask_OnnsW08k,wftask_4ejSW7en" // string | Comma separated list of workflow task IDs (optional)
    schemaId := "prstsch_KnR9iVum" // string | The ID of the workflow task schema of the workflow output (optional)
    watcherIds := "ent_a0SApq3z,ent_asdf72354,null" // string | Comma separated list of user IDs or \"null\" (optional)
    responsibleTeamIds := "team_Thepp2c7,team_QqHMbfqK,null" // string | Comma separated list of team IDs or \"null\" (optional)
    creationOriginIds := "etr_d00c97,etr_30ad79" // string | Comma separated list of entry IDs (optional)
    linkedItemIdsAnyOf := "bfi_ed1ef7,con_1c76c9" // string | Comma separated list of bioentity or storable IDs. Returns workflow outputs where the output's schema fields reference at least one of the provided items.  (optional)
    linkedItemIdsAllOf := "bfi_ed1ef7,con_1c76c9" // string | Comma separated list of bioentity or storable IDs. Returns workflow outputs where the output's schema fields reference all of the provided items.  (optional)
    linkedItemIdsNoneOf := "bfi_ed1ef7,con_1c76c9" // string | Comma separated list of bioentity or storable IDs. Returns workflow outputs where the output's schema fields do not reference any of the provided items.  (optional)
    schemaFields := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Schema field value. For Integer, Float, and Date type fields, supports the >= and <= operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  (optional)
    name := "PR-1" // string | The name of the workflow task (optional)
    nameIncludes := "PR" // string | Part of the name of the workflow task (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of user IDs. (optional)
    modifiedAt := "> 2021-02-04T18:53:36Z" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    nextToken := "Im5ldyB0ZXN0Ig==" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional) (default to 50)
    displayIds := "ANG1-O1,ANG1-O2" // string | Comma-separated list of Workflow Output Display IDs. (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived workflow outputs. Use \"ANY_ARCHIVED\" to filter for archived workflow outputs regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowOutputsApi.ListWorkflowOutputs(context.Background()).Ids(ids).WorkflowTaskGroupIds(workflowTaskGroupIds).WorkflowTaskIds(workflowTaskIds).SchemaId(schemaId).WatcherIds(watcherIds).ResponsibleTeamIds(responsibleTeamIds).CreationOriginIds(creationOriginIds).LinkedItemIdsAnyOf(linkedItemIdsAnyOf).LinkedItemIdsAllOf(linkedItemIdsAllOf).LinkedItemIdsNoneOf(linkedItemIdsNoneOf).SchemaFields(schemaFields).Name(name).NameIncludes(nameIncludes).CreatorIds(creatorIds).ModifiedAt(modifiedAt).NextToken(nextToken).PageSize(pageSize).DisplayIds(displayIds).ArchiveReason(archiveReason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowOutputsApi.ListWorkflowOutputs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowOutputs`: WorkflowOutputsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `WorkflowOutputsApi.ListWorkflowOutputs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowOutputsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Comma separated list of workflow output IDs | 
 **workflowTaskGroupIds** | **string** | Comma separated list of workflow IDs | 
 **workflowTaskIds** | **string** | Comma separated list of workflow task IDs | 
 **schemaId** | **string** | The ID of the workflow task schema of the workflow output | 
 **watcherIds** | **string** | Comma separated list of user IDs or \&quot;null\&quot; | 
 **responsibleTeamIds** | **string** | Comma separated list of team IDs or \&quot;null\&quot; | 
 **creationOriginIds** | **string** | Comma separated list of entry IDs | 
 **linkedItemIdsAnyOf** | **string** | Comma separated list of bioentity or storable IDs. Returns workflow outputs where the output&#39;s schema fields reference at least one of the provided items.  | 
 **linkedItemIdsAllOf** | **string** | Comma separated list of bioentity or storable IDs. Returns workflow outputs where the output&#39;s schema fields reference all of the provided items.  | 
 **linkedItemIdsNoneOf** | **string** | Comma separated list of bioentity or storable IDs. Returns workflow outputs where the output&#39;s schema fields do not reference any of the provided items.  | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **name** | **string** | The name of the workflow task | 
 **nameIncludes** | **string** | Part of the name of the workflow task | 
 **creatorIds** | **string** | Comma separated list of user IDs. | 
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **displayIds** | **string** | Comma-separated list of Workflow Output Display IDs. | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived workflow outputs. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived workflow outputs regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 

### Return type

[**WorkflowOutputsPaginatedList**](WorkflowOutputsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveWorkflowOutputs

> WorkflowOutputsArchivalChange UnarchiveWorkflowOutputs(ctx).WorkflowOutputsUnarchive(workflowOutputsUnarchive).Execute()

Unarchive one or more workflow outputs



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
    workflowOutputsUnarchive := *openapiclient.NewWorkflowOutputsUnarchive([]string{"WorkflowOutputIds_example"}) // WorkflowOutputsUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowOutputsApi.UnarchiveWorkflowOutputs(context.Background()).WorkflowOutputsUnarchive(workflowOutputsUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowOutputsApi.UnarchiveWorkflowOutputs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveWorkflowOutputs`: WorkflowOutputsArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `WorkflowOutputsApi.UnarchiveWorkflowOutputs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveWorkflowOutputsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowOutputsUnarchive** | [**WorkflowOutputsUnarchive**](WorkflowOutputsUnarchive.md) |  | 

### Return type

[**WorkflowOutputsArchivalChange**](WorkflowOutputsArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowOutput

> WorkflowOutput UpdateWorkflowOutput(ctx, workflowOutputId).WorkflowOutputUpdate(workflowOutputUpdate).Execute()

Update a workflow output



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
    workflowOutputId := "wfout_5cJLQKVF" // string | The ID of the workflow output
    workflowOutputUpdate := *openapiclient.NewWorkflowOutputUpdate() // WorkflowOutputUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowOutputsApi.UpdateWorkflowOutput(context.Background(), workflowOutputId).WorkflowOutputUpdate(workflowOutputUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowOutputsApi.UpdateWorkflowOutput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflowOutput`: WorkflowOutput
    fmt.Fprintf(os.Stdout, "Response from `WorkflowOutputsApi.UpdateWorkflowOutput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowOutputId** | **string** | The ID of the workflow output | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowOutputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowOutputUpdate** | [**WorkflowOutputUpdate**](WorkflowOutputUpdate.md) |  | 

### Return type

[**WorkflowOutput**](WorkflowOutput.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

