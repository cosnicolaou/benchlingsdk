# \WorkflowTasksApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveWorkflowTasks**](WorkflowTasksApi.md#ArchiveWorkflowTasks) | **Post** /workflow-tasks:archive | Archive one or more workflow tasks
[**BulkCopyWorkflowTasks**](WorkflowTasksApi.md#BulkCopyWorkflowTasks) | **Post** /workflow-tasks:bulk-copy | Bulk creates new workflow tasks where each new task has the same fields and assignee as one of the provided tasks and creates a relationship between the provided task and its copy 
[**BulkCreateWorkflowTasks**](WorkflowTasksApi.md#BulkCreateWorkflowTasks) | **Post** /workflow-tasks:bulk-create | Create one or more workflow tasks
[**BulkUpdateWorkflowTasks**](WorkflowTasksApi.md#BulkUpdateWorkflowTasks) | **Post** /workflow-tasks:bulk-update | Update one or more workflow task
[**CopyWorkflowTask**](WorkflowTasksApi.md#CopyWorkflowTask) | **Post** /workflow-tasks/{workflow_task_id}:copy | Creates a new workflow task with the same fields and assignee as the provided task and creates a relationship between the two tasks 
[**CreateWorkflowTask**](WorkflowTasksApi.md#CreateWorkflowTask) | **Post** /workflow-tasks | Create a new workflow task
[**GetWorkflowTask**](WorkflowTasksApi.md#GetWorkflowTask) | **Get** /workflow-tasks/{workflow_task_id} | Get a workflow task
[**ListWorkflowTasks**](WorkflowTasksApi.md#ListWorkflowTasks) | **Get** /workflow-tasks | List workflow tasks
[**UnarchiveWorkflowTasks**](WorkflowTasksApi.md#UnarchiveWorkflowTasks) | **Post** /workflow-tasks:unarchive | Unarchive one or more workflow tasks
[**UpdateWorkflowTask**](WorkflowTasksApi.md#UpdateWorkflowTask) | **Patch** /workflow-tasks/{workflow_task_id} | Update a workflow task



## ArchiveWorkflowTasks

> WorkflowTasksArchivalChange ArchiveWorkflowTasks(ctx).WorkflowTasksArchive(workflowTasksArchive).Execute()

Archive one or more workflow tasks



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
    workflowTasksArchive := *openapiclient.NewWorkflowTasksArchive(openapiclient.WorkflowTaskArchiveReason("Made in error"), []string{"wftask_OnnsW08k"}) // WorkflowTasksArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTasksApi.ArchiveWorkflowTasks(context.Background()).WorkflowTasksArchive(workflowTasksArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTasksApi.ArchiveWorkflowTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveWorkflowTasks`: WorkflowTasksArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTasksApi.ArchiveWorkflowTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveWorkflowTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowTasksArchive** | [**WorkflowTasksArchive**](WorkflowTasksArchive.md) |  | 

### Return type

[**WorkflowTasksArchivalChange**](WorkflowTasksArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkCopyWorkflowTasks

> AsyncTaskLink BulkCopyWorkflowTasks(ctx).WorkflowTasksBulkCopyRequest(workflowTasksBulkCopyRequest).Execute()

Bulk creates new workflow tasks where each new task has the same fields and assignee as one of the provided tasks and creates a relationship between the provided task and its copy 



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
    workflowTasksBulkCopyRequest := *openapiclient.NewWorkflowTasksBulkCopyRequest() // WorkflowTasksBulkCopyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTasksApi.BulkCopyWorkflowTasks(context.Background()).WorkflowTasksBulkCopyRequest(workflowTasksBulkCopyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTasksApi.BulkCopyWorkflowTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCopyWorkflowTasks`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTasksApi.BulkCopyWorkflowTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCopyWorkflowTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowTasksBulkCopyRequest** | [**WorkflowTasksBulkCopyRequest**](WorkflowTasksBulkCopyRequest.md) |  | 

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


## BulkCreateWorkflowTasks

> AsyncTaskLink BulkCreateWorkflowTasks(ctx).WorkflowTasksBulkCreateRequest(workflowTasksBulkCreateRequest).Execute()

Create one or more workflow tasks



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
    workflowTasksBulkCreateRequest := *openapiclient.NewWorkflowTasksBulkCreateRequest() // WorkflowTasksBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTasksApi.BulkCreateWorkflowTasks(context.Background()).WorkflowTasksBulkCreateRequest(workflowTasksBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTasksApi.BulkCreateWorkflowTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateWorkflowTasks`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTasksApi.BulkCreateWorkflowTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateWorkflowTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowTasksBulkCreateRequest** | [**WorkflowTasksBulkCreateRequest**](WorkflowTasksBulkCreateRequest.md) |  | 

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


## BulkUpdateWorkflowTasks

> AsyncTaskLink BulkUpdateWorkflowTasks(ctx).WorkflowTasksBulkUpdateRequest(workflowTasksBulkUpdateRequest).Execute()

Update one or more workflow task



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
    workflowTasksBulkUpdateRequest := *openapiclient.NewWorkflowTasksBulkUpdateRequest() // WorkflowTasksBulkUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTasksApi.BulkUpdateWorkflowTasks(context.Background()).WorkflowTasksBulkUpdateRequest(workflowTasksBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTasksApi.BulkUpdateWorkflowTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateWorkflowTasks`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTasksApi.BulkUpdateWorkflowTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateWorkflowTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowTasksBulkUpdateRequest** | [**WorkflowTasksBulkUpdateRequest**](WorkflowTasksBulkUpdateRequest.md) |  | 

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


## CopyWorkflowTask

> WorkflowTask CopyWorkflowTask(ctx, workflowTaskId).Execute()

Creates a new workflow task with the same fields and assignee as the provided task and creates a relationship between the two tasks 



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
    workflowTaskId := "wftask_OnnsW08k" // string | The ID of the workflow task

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTasksApi.CopyWorkflowTask(context.Background(), workflowTaskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTasksApi.CopyWorkflowTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyWorkflowTask`: WorkflowTask
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTasksApi.CopyWorkflowTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowTaskId** | **string** | The ID of the workflow task | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyWorkflowTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowTask**](WorkflowTask.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflowTask

> WorkflowTask CreateWorkflowTask(ctx).WorkflowTaskCreate(workflowTaskCreate).Execute()

Create a new workflow task



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
    workflowTaskCreate := *openapiclient.NewWorkflowTaskCreate("prs_giVNQcTL") // WorkflowTaskCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTasksApi.CreateWorkflowTask(context.Background()).WorkflowTaskCreate(workflowTaskCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTasksApi.CreateWorkflowTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowTask`: WorkflowTask
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTasksApi.CreateWorkflowTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowTaskCreate** | [**WorkflowTaskCreate**](WorkflowTaskCreate.md) |  | 

### Return type

[**WorkflowTask**](WorkflowTask.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowTask

> WorkflowTask GetWorkflowTask(ctx, workflowTaskId).Execute()

Get a workflow task



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
    workflowTaskId := "wftask_OnnsW08k" // string | The ID of the workflow task

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTasksApi.GetWorkflowTask(context.Background(), workflowTaskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTasksApi.GetWorkflowTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowTask`: WorkflowTask
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTasksApi.GetWorkflowTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowTaskId** | **string** | The ID of the workflow task | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowTask**](WorkflowTask.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowTasks

> WorkflowTasksPaginatedList ListWorkflowTasks(ctx).Ids(ids).WorkflowTaskGroupIds(workflowTaskGroupIds).SchemaId(schemaId).StatusIds(statusIds).AssigneeIds(assigneeIds).WatcherIds(watcherIds).ResponsibleTeamIds(responsibleTeamIds).ExecutionOriginIds(executionOriginIds).ExecutionTypes(executionTypes).LinkedItemIdsAnyOf(linkedItemIdsAnyOf).LinkedItemIdsAllOf(linkedItemIdsAllOf).LinkedItemIdsNoneOf(linkedItemIdsNoneOf).SchemaFields(schemaFields).Name(name).NameIncludes(nameIncludes).CreatorIds(creatorIds).ScheduledOn(scheduledOn).ScheduledOnLt(scheduledOnLt).ScheduledOnLte(scheduledOnLte).ScheduledOnGte(scheduledOnGte).ScheduledOnGt(scheduledOnGt).ModifiedAt(modifiedAt).NextToken(nextToken).PageSize(pageSize).DisplayIds(displayIds).ArchiveReason(archiveReason).Execute()

List workflow tasks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    ids := "wftask_OnnsW08k,wftask_4ejSW7en" // string | Comma separated list of workflow task IDs (optional)
    workflowTaskGroupIds := "prs_giVNQcTL,prs_t6m99v1" // string | Comma separated list of workflow IDs (optional)
    schemaId := "prstsch_KnR9iVum" // string | The ID of the workflow task schema of the workflow task (optional)
    statusIds := "wfts_wQzUCsW0,wfts_VFvwv7JV" // string | Comma separated list of workflow task status ids (optional)
    assigneeIds := "ent_a0SApq3z,null" // string | Comma separated list of user ids or \"null\" (optional)
    watcherIds := "ent_a0SApq3z,ent_asdf72354,null" // string | Comma separated list of user IDs or \"null\" (optional)
    responsibleTeamIds := "team_Thepp2c7,team_QqHMbfqK,null" // string | Comma separated list of team IDs or \"null\" (optional)
    executionOriginIds := "etr_d00c97,etr_30ad79" // string | Comma separated list of entry IDs (optional)
    executionTypes := "ENTRY,DIRECT" // string | Comma separated list of workflow execution types. Acceptable execution types are \"DIRECT\" and \"ENTRY\"  (optional)
    linkedItemIdsAnyOf := "bfi_ed1ef7,con_1c76c9" // string | Comma separated list of bioentity or storable IDs. Returns workflow tasks where the task's schema fields reference at least one of the provided items.  (optional)
    linkedItemIdsAllOf := "bfi_ed1ef7,con_1c76c9" // string | Comma separated list of bioentity or storable IDs. Returns workflow tasks where the task's schema fields reference all of the provided items.  (optional)
    linkedItemIdsNoneOf := "bfi_ed1ef7,con_1c76c9" // string | Comma separated list of bioentity or storable IDs. Returns workflow tasks where the task's schema fields do not reference any of the provided items.  (optional)
    schemaFields := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Schema field value. For Integer, Float, and Date type fields, supports the >= and <= operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  (optional)
    name := "PR-1" // string | The name of the workflow task (optional)
    nameIncludes := "PR" // string | Part of the name of the workflow task (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of user IDs. (optional)
    scheduledOn := *openapiclient.NewListWorkflowTasksScheduledOnParameter() // ListWorkflowTasksScheduledOnParameter | The date on which the task was scheduled to be executed. Returns tasks which are scheduled on the provided date. If \"null\" is provided returns tasks which are unshceduled.  (optional)
    scheduledOnLt := time.Now() // string | The date on which the task was scheduled to be executed. Returns tasks which are scheduled before the provided date.  (optional)
    scheduledOnLte := time.Now() // string | The date on which the task was scheduled to be executed. Returns tasks which are scheduled before or on the provided date.  (optional)
    scheduledOnGte := time.Now() // string | The date on which the task was scheduled to be executed. Returns tasks which are scheduled on or after the provided date.  (optional)
    scheduledOnGt := time.Now() // string | The date on which the task was scheduled to be executed. Returns tasks which are scheduled after the provided date.  (optional)
    modifiedAt := "> 2021-02-04T18:53:36Z" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    nextToken := "Im5ldyB0ZXN0Ig==" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional) (default to 50)
    displayIds := "ANG1-T1,ANG1-T2" // string | Comma-separated list of Workflow Task Display IDs. (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived workflow tasks. Use \"ANY_ARCHIVED\" to filter for archived workflow tasks regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTasksApi.ListWorkflowTasks(context.Background()).Ids(ids).WorkflowTaskGroupIds(workflowTaskGroupIds).SchemaId(schemaId).StatusIds(statusIds).AssigneeIds(assigneeIds).WatcherIds(watcherIds).ResponsibleTeamIds(responsibleTeamIds).ExecutionOriginIds(executionOriginIds).ExecutionTypes(executionTypes).LinkedItemIdsAnyOf(linkedItemIdsAnyOf).LinkedItemIdsAllOf(linkedItemIdsAllOf).LinkedItemIdsNoneOf(linkedItemIdsNoneOf).SchemaFields(schemaFields).Name(name).NameIncludes(nameIncludes).CreatorIds(creatorIds).ScheduledOn(scheduledOn).ScheduledOnLt(scheduledOnLt).ScheduledOnLte(scheduledOnLte).ScheduledOnGte(scheduledOnGte).ScheduledOnGt(scheduledOnGt).ModifiedAt(modifiedAt).NextToken(nextToken).PageSize(pageSize).DisplayIds(displayIds).ArchiveReason(archiveReason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTasksApi.ListWorkflowTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowTasks`: WorkflowTasksPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTasksApi.ListWorkflowTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Comma separated list of workflow task IDs | 
 **workflowTaskGroupIds** | **string** | Comma separated list of workflow IDs | 
 **schemaId** | **string** | The ID of the workflow task schema of the workflow task | 
 **statusIds** | **string** | Comma separated list of workflow task status ids | 
 **assigneeIds** | **string** | Comma separated list of user ids or \&quot;null\&quot; | 
 **watcherIds** | **string** | Comma separated list of user IDs or \&quot;null\&quot; | 
 **responsibleTeamIds** | **string** | Comma separated list of team IDs or \&quot;null\&quot; | 
 **executionOriginIds** | **string** | Comma separated list of entry IDs | 
 **executionTypes** | **string** | Comma separated list of workflow execution types. Acceptable execution types are \&quot;DIRECT\&quot; and \&quot;ENTRY\&quot;  | 
 **linkedItemIdsAnyOf** | **string** | Comma separated list of bioentity or storable IDs. Returns workflow tasks where the task&#39;s schema fields reference at least one of the provided items.  | 
 **linkedItemIdsAllOf** | **string** | Comma separated list of bioentity or storable IDs. Returns workflow tasks where the task&#39;s schema fields reference all of the provided items.  | 
 **linkedItemIdsNoneOf** | **string** | Comma separated list of bioentity or storable IDs. Returns workflow tasks where the task&#39;s schema fields do not reference any of the provided items.  | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **name** | **string** | The name of the workflow task | 
 **nameIncludes** | **string** | Part of the name of the workflow task | 
 **creatorIds** | **string** | Comma separated list of user IDs. | 
 **scheduledOn** | [**ListWorkflowTasksScheduledOnParameter**](ListWorkflowTasksScheduledOnParameter.md) | The date on which the task was scheduled to be executed. Returns tasks which are scheduled on the provided date. If \&quot;null\&quot; is provided returns tasks which are unshceduled.  | 
 **scheduledOnLt** | **string** | The date on which the task was scheduled to be executed. Returns tasks which are scheduled before the provided date.  | 
 **scheduledOnLte** | **string** | The date on which the task was scheduled to be executed. Returns tasks which are scheduled before or on the provided date.  | 
 **scheduledOnGte** | **string** | The date on which the task was scheduled to be executed. Returns tasks which are scheduled on or after the provided date.  | 
 **scheduledOnGt** | **string** | The date on which the task was scheduled to be executed. Returns tasks which are scheduled after the provided date.  | 
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **displayIds** | **string** | Comma-separated list of Workflow Task Display IDs. | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived workflow tasks. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived workflow tasks regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 

### Return type

[**WorkflowTasksPaginatedList**](WorkflowTasksPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveWorkflowTasks

> WorkflowTasksArchivalChange UnarchiveWorkflowTasks(ctx).WorkflowTasksUnarchive(workflowTasksUnarchive).Execute()

Unarchive one or more workflow tasks



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
    workflowTasksUnarchive := *openapiclient.NewWorkflowTasksUnarchive([]string{"wftask_OnnsW08k"}) // WorkflowTasksUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTasksApi.UnarchiveWorkflowTasks(context.Background()).WorkflowTasksUnarchive(workflowTasksUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTasksApi.UnarchiveWorkflowTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveWorkflowTasks`: WorkflowTasksArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTasksApi.UnarchiveWorkflowTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveWorkflowTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowTasksUnarchive** | [**WorkflowTasksUnarchive**](WorkflowTasksUnarchive.md) |  | 

### Return type

[**WorkflowTasksArchivalChange**](WorkflowTasksArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowTask

> WorkflowTask UpdateWorkflowTask(ctx, workflowTaskId).WorkflowTaskUpdate(workflowTaskUpdate).Execute()

Update a workflow task



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
    workflowTaskId := "wftask_OnnsW08k" // string | The ID of the workflow task
    workflowTaskUpdate := *openapiclient.NewWorkflowTaskUpdate() // WorkflowTaskUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTasksApi.UpdateWorkflowTask(context.Background(), workflowTaskId).WorkflowTaskUpdate(workflowTaskUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTasksApi.UpdateWorkflowTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflowTask`: WorkflowTask
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTasksApi.UpdateWorkflowTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowTaskId** | **string** | The ID of the workflow task | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowTaskUpdate** | [**WorkflowTaskUpdate**](WorkflowTaskUpdate.md) |  | 

### Return type

[**WorkflowTask**](WorkflowTask.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

