# \ProjectsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveProjects**](ProjectsApi.md#ArchiveProjects) | **Post** /projects:archive | Archive projects
[**GetProject**](ProjectsApi.md#GetProject) | **Get** /projects/{project_id} | Get a project by ID
[**ListProjects**](ProjectsApi.md#ListProjects) | **Get** /projects | List projects
[**UnarchiveProjects**](ProjectsApi.md#UnarchiveProjects) | **Post** /projects:unarchive | Unarchive projects



## ArchiveProjects

> ProjectsArchivalChange ArchiveProjects(ctx).ProjectsArchive(projectsArchive).Execute()

Archive projects



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
    projectsArchive := *openapiclient.NewProjectsArchive([]string{"ProjectIds_example"}, "Reason_example") // ProjectsArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ArchiveProjects(context.Background()).ProjectsArchive(projectsArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ArchiveProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveProjects`: ProjectsArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ArchiveProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsArchive** | [**ProjectsArchive**](ProjectsArchive.md) |  | 

### Return type

[**ProjectsArchivalChange**](ProjectsArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProject

> Project GetProject(ctx, projectId).Execute()

Get a project by ID



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
    projectId := "projectId_example" // string | ID of project to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProject`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of project to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjects

> ProjectsPaginatedList ListProjects(ctx).NextToken(nextToken).PageSize(pageSize).Sort(sort).ArchiveReason(archiveReason).Ids(ids).Name(name).Execute()

List projects



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
    nextToken := "nextToken_example" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional) (default to 50)
    sort := "sort_example" // string | Method by which to order search results. Valid sorts are modifiedAt (modified time, most recent first) and name (project name, alphabetical). Optionally add :asc or :desc to specify ascending or descending order. Default is modifiedAt.  (optional) (default to "modifiedAt")
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived projects. Use \"ANY_ARCHIVED\" to filter for archived projects regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    ids := "src_ZJy8RTbo,src_8GVbVkPj,src_qREJ33rn" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    name := "name_example" // string | Name of a project. Restricts results to those with the specified name. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ListProjects(context.Background()).NextToken(nextToken).PageSize(pageSize).Sort(sort).ArchiveReason(archiveReason).Ids(ids).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ListProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProjects`: ProjectsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ListProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **sort** | **string** | Method by which to order search results. Valid sorts are modifiedAt (modified time, most recent first) and name (project name, alphabetical). Optionally add :asc or :desc to specify ascending or descending order. Default is modifiedAt.  | [default to &quot;modifiedAt&quot;]
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived projects. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived projects regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **name** | **string** | Name of a project. Restricts results to those with the specified name. | 

### Return type

[**ProjectsPaginatedList**](ProjectsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveProjects

> ProjectsArchivalChange UnarchiveProjects(ctx).ProjectsUnarchive(projectsUnarchive).Execute()

Unarchive projects



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
    projectsUnarchive := *openapiclient.NewProjectsUnarchive([]string{"ProjectIds_example"}) // ProjectsUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.UnarchiveProjects(context.Background()).ProjectsUnarchive(projectsUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.UnarchiveProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveProjects`: ProjectsArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.UnarchiveProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsUnarchive** | [**ProjectsUnarchive**](ProjectsUnarchive.md) |  | 

### Return type

[**ProjectsArchivalChange**](ProjectsArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

