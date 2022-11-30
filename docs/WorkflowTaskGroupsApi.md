# \WorkflowTaskGroupsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveWorkflowTaskGroups**](WorkflowTaskGroupsApi.md#ArchiveWorkflowTaskGroups) | **Post** /workflow-task-groups:archive | Archive one or more workflows
[**CreateWorkflowTaskGroup**](WorkflowTaskGroupsApi.md#CreateWorkflowTaskGroup) | **Post** /workflow-task-groups | Create a new workflow task group
[**GetWorkflowTaskGroup**](WorkflowTaskGroupsApi.md#GetWorkflowTaskGroup) | **Get** /workflow-task-groups/{workflow_task_group_id} | Get a workflow task group
[**ListWorkflowTaskGroups**](WorkflowTaskGroupsApi.md#ListWorkflowTaskGroups) | **Get** /workflow-task-groups | List workflow task groups
[**UnarchiveWorkflowTaskGroups**](WorkflowTaskGroupsApi.md#UnarchiveWorkflowTaskGroups) | **Post** /workflow-task-groups:unarchive | Unarchive one or more workflows
[**UpdateWorkflowTaskGroup**](WorkflowTaskGroupsApi.md#UpdateWorkflowTaskGroup) | **Patch** /workflow-task-groups/{workflow_task_group_id} | Update a workflow task group



## ArchiveWorkflowTaskGroups

> WorkflowTaskGroupsArchivalChange ArchiveWorkflowTaskGroups(ctx).WorkflowTaskGroupsArchive(workflowTaskGroupsArchive).Execute()

Archive one or more workflows



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
    workflowTaskGroupsArchive := *openapiclient.NewWorkflowTaskGroupsArchive(openapiclient.WorkflowTaskGroupArchiveReason("Made in error"), []string{"prs_giVNQcTL"}) // WorkflowTaskGroupsArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTaskGroupsApi.ArchiveWorkflowTaskGroups(context.Background()).WorkflowTaskGroupsArchive(workflowTaskGroupsArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTaskGroupsApi.ArchiveWorkflowTaskGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveWorkflowTaskGroups`: WorkflowTaskGroupsArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTaskGroupsApi.ArchiveWorkflowTaskGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveWorkflowTaskGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowTaskGroupsArchive** | [**WorkflowTaskGroupsArchive**](WorkflowTaskGroupsArchive.md) |  | 

### Return type

[**WorkflowTaskGroupsArchivalChange**](WorkflowTaskGroupsArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflowTaskGroup

> WorkflowTaskGroup CreateWorkflowTaskGroup(ctx).WorkflowTaskGroupCreate(workflowTaskGroupCreate).Execute()

Create a new workflow task group



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
    workflowTaskGroupCreate := *openapiclient.NewWorkflowTaskGroupCreate("FolderId_example", "prstsch_KnR9iVum") // WorkflowTaskGroupCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTaskGroupsApi.CreateWorkflowTaskGroup(context.Background()).WorkflowTaskGroupCreate(workflowTaskGroupCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTaskGroupsApi.CreateWorkflowTaskGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowTaskGroup`: WorkflowTaskGroup
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTaskGroupsApi.CreateWorkflowTaskGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowTaskGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowTaskGroupCreate** | [**WorkflowTaskGroupCreate**](WorkflowTaskGroupCreate.md) |  | 

### Return type

[**WorkflowTaskGroup**](WorkflowTaskGroup.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowTaskGroup

> WorkflowTaskGroup GetWorkflowTaskGroup(ctx, workflowTaskGroupId).Execute()

Get a workflow task group



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
    workflowTaskGroupId := "prs_giVNQcTL" // string | The ID of the workflow task group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTaskGroupsApi.GetWorkflowTaskGroup(context.Background(), workflowTaskGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTaskGroupsApi.GetWorkflowTaskGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowTaskGroup`: WorkflowTaskGroup
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTaskGroupsApi.GetWorkflowTaskGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowTaskGroupId** | **string** | The ID of the workflow task group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowTaskGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowTaskGroup**](WorkflowTaskGroup.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowTaskGroups

> WorkflowTaskGroupsPaginatedList ListWorkflowTaskGroups(ctx).Ids(ids).SchemaId(schemaId).FolderId(folderId).ProjectId(projectId).MentionedIn(mentionedIn).WatcherIds(watcherIds).ExecutionTypes(executionTypes).ResponsibleTeamIds(responsibleTeamIds).StatusIdsAnyOf(statusIdsAnyOf).StatusIdsNoneOf(statusIdsNoneOf).StatusIdsOnly(statusIdsOnly).Name(name).NameIncludes(nameIncludes).CreatorIds(creatorIds).ModifiedAt(modifiedAt).NextToken(nextToken).PageSize(pageSize).DisplayIds(displayIds).ArchiveReason(archiveReason).Execute()

List workflow task groups



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
    ids := "prs_giVNQcTL,prs_t6m99v1" // string | Comma separated list of workflow task group IDs (optional)
    schemaId := "prstsch_KnR9iVum,prstsch_nJ34lw9y" // string | The workflow task schema ID of tasks in this task group (optional)
    folderId := "lib_bf0636" // string | A folder ID (optional)
    projectId := "src_NetYd96a" // string | A project ID (optional)
    mentionedIn := "etr_30ad79,etr_d00c97" // string | A comma separated list entry IDs (optional)
    watcherIds := "ent_a0SApq3z,ent_asdf72354,null" // string | Comma separated list of user IDs or \"null\" (optional)
    executionTypes := "ENTRY,DIRECT" // string | Comma separated list of workflow execution types. Acceptable execution types are \"DIRECT\" and \"ENTRY\"  (optional)
    responsibleTeamIds := "team_Thepp2c7,team_QqHMbfqK,null" // string | Comma separated list of team IDs or \"null\" (optional)
    statusIdsAnyOf := "wfts_VFvwv7JV,wfts_wQzUCsW0" // string | Commas separated list of Status ids. Returns workflows where at least one task is of one of the provided statuses. (optional)
    statusIdsNoneOf := "wfts_VFvwv7JV,wfts_wQzUCsW0" // string | Commas separated list of Status ids. Returns workflows where none of the tasks are of any of the provided statuses. (optional)
    statusIdsOnly := "wfts_VFvwv7JV,wfts_wQzUCsW0" // string | Commas separated list of Status ids. Returns workflows where all of the tasks are of one of the provided statuses. (optional)
    name := "Plasmid Transformation" // string | The name of the workflow task group (optional)
    nameIncludes := "Plasmid" // string | Part of the name of the workflow task group (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of user IDs. (optional)
    modifiedAt := "> 2021-02-04T18:53:36Z" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    nextToken := "Im5ldyB0ZXN0Ig==" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional) (default to 50)
    displayIds := "VPR001,VPR002" // string | Comma-separated list of Workflow Display IDs. (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived workflow task groups. Use \"ANY_ARCHIVED\" to filter for archived workflow task groups regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTaskGroupsApi.ListWorkflowTaskGroups(context.Background()).Ids(ids).SchemaId(schemaId).FolderId(folderId).ProjectId(projectId).MentionedIn(mentionedIn).WatcherIds(watcherIds).ExecutionTypes(executionTypes).ResponsibleTeamIds(responsibleTeamIds).StatusIdsAnyOf(statusIdsAnyOf).StatusIdsNoneOf(statusIdsNoneOf).StatusIdsOnly(statusIdsOnly).Name(name).NameIncludes(nameIncludes).CreatorIds(creatorIds).ModifiedAt(modifiedAt).NextToken(nextToken).PageSize(pageSize).DisplayIds(displayIds).ArchiveReason(archiveReason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTaskGroupsApi.ListWorkflowTaskGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowTaskGroups`: WorkflowTaskGroupsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTaskGroupsApi.ListWorkflowTaskGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowTaskGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Comma separated list of workflow task group IDs | 
 **schemaId** | **string** | The workflow task schema ID of tasks in this task group | 
 **folderId** | **string** | A folder ID | 
 **projectId** | **string** | A project ID | 
 **mentionedIn** | **string** | A comma separated list entry IDs | 
 **watcherIds** | **string** | Comma separated list of user IDs or \&quot;null\&quot; | 
 **executionTypes** | **string** | Comma separated list of workflow execution types. Acceptable execution types are \&quot;DIRECT\&quot; and \&quot;ENTRY\&quot;  | 
 **responsibleTeamIds** | **string** | Comma separated list of team IDs or \&quot;null\&quot; | 
 **statusIdsAnyOf** | **string** | Commas separated list of Status ids. Returns workflows where at least one task is of one of the provided statuses. | 
 **statusIdsNoneOf** | **string** | Commas separated list of Status ids. Returns workflows where none of the tasks are of any of the provided statuses. | 
 **statusIdsOnly** | **string** | Commas separated list of Status ids. Returns workflows where all of the tasks are of one of the provided statuses. | 
 **name** | **string** | The name of the workflow task group | 
 **nameIncludes** | **string** | Part of the name of the workflow task group | 
 **creatorIds** | **string** | Comma separated list of user IDs. | 
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **displayIds** | **string** | Comma-separated list of Workflow Display IDs. | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived workflow task groups. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived workflow task groups regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 

### Return type

[**WorkflowTaskGroupsPaginatedList**](WorkflowTaskGroupsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveWorkflowTaskGroups

> WorkflowTaskGroupsArchivalChange UnarchiveWorkflowTaskGroups(ctx).WorkflowTaskGroupsUnarchive(workflowTaskGroupsUnarchive).Execute()

Unarchive one or more workflows



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
    workflowTaskGroupsUnarchive := *openapiclient.NewWorkflowTaskGroupsUnarchive([]string{"prs_giVNQcTL"}) // WorkflowTaskGroupsUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTaskGroupsApi.UnarchiveWorkflowTaskGroups(context.Background()).WorkflowTaskGroupsUnarchive(workflowTaskGroupsUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTaskGroupsApi.UnarchiveWorkflowTaskGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveWorkflowTaskGroups`: WorkflowTaskGroupsArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTaskGroupsApi.UnarchiveWorkflowTaskGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveWorkflowTaskGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowTaskGroupsUnarchive** | [**WorkflowTaskGroupsUnarchive**](WorkflowTaskGroupsUnarchive.md) |  | 

### Return type

[**WorkflowTaskGroupsArchivalChange**](WorkflowTaskGroupsArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowTaskGroup

> WorkflowTaskGroup UpdateWorkflowTaskGroup(ctx, workflowTaskGroupId).WorkflowTaskGroupUpdate(workflowTaskGroupUpdate).Execute()

Update a workflow task group



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
    workflowTaskGroupId := "prs_giVNQcTL" // string | The ID of the workflow task group
    workflowTaskGroupUpdate := *openapiclient.NewWorkflowTaskGroupUpdate() // WorkflowTaskGroupUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTaskGroupsApi.UpdateWorkflowTaskGroup(context.Background(), workflowTaskGroupId).WorkflowTaskGroupUpdate(workflowTaskGroupUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTaskGroupsApi.UpdateWorkflowTaskGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflowTaskGroup`: WorkflowTaskGroup
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTaskGroupsApi.UpdateWorkflowTaskGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowTaskGroupId** | **string** | The ID of the workflow task group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowTaskGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowTaskGroupUpdate** | [**WorkflowTaskGroupUpdate**](WorkflowTaskGroupUpdate.md) |  | 

### Return type

[**WorkflowTaskGroup**](WorkflowTaskGroup.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

