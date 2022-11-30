# \FoldersApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveFolders**](FoldersApi.md#ArchiveFolders) | **Post** /folders:archive | Archive folders
[**CreateFolder**](FoldersApi.md#CreateFolder) | **Post** /folders | Create folder
[**GetFolder**](FoldersApi.md#GetFolder) | **Get** /folders/{folder_id} | Get a folder by ID
[**ListFolders**](FoldersApi.md#ListFolders) | **Get** /folders | List folders
[**UnarchiveFolders**](FoldersApi.md#UnarchiveFolders) | **Post** /folders:unarchive | Unarchive folders



## ArchiveFolders

> FoldersArchivalChange ArchiveFolders(ctx).FoldersArchive(foldersArchive).Execute()

Archive folders



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
    foldersArchive := *openapiclient.NewFoldersArchive([]string{"FolderIds_example"}, "Reason_example") // FoldersArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FoldersApi.ArchiveFolders(context.Background()).FoldersArchive(foldersArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.ArchiveFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveFolders`: FoldersArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.ArchiveFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **foldersArchive** | [**FoldersArchive**](FoldersArchive.md) |  | 

### Return type

[**FoldersArchivalChange**](FoldersArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFolder

> Folder CreateFolder(ctx).FolderCreate(folderCreate).Execute()

Create folder



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
    folderCreate := *openapiclient.NewFolderCreate("Name_example", "ParentFolderId_example") // FolderCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FoldersApi.CreateFolder(context.Background()).FolderCreate(folderCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.CreateFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFolder`: Folder
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.CreateFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderCreate** | [**FolderCreate**](FolderCreate.md) |  | 

### Return type

[**Folder**](Folder.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolder

> Folder GetFolder(ctx, folderId).Execute()

Get a folder by ID



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
    folderId := "folderId_example" // string | ID of folder to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FoldersApi.GetFolder(context.Background(), folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.GetFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFolder`: Folder
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.GetFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | ID of folder to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Folder**](Folder.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFolders

> FoldersPaginatedList ListFolders(ctx).NextToken(nextToken).PageSize(pageSize).Sort(sort).ArchiveReason(archiveReason).NameIncludes(nameIncludes).ParentFolderId(parentFolderId).ProjectId(projectId).Ids(ids).Name(name).Section(section).Execute()

List folders



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
    sort := "sort_example" // string | Method by which to order search results. Valid sorts are modifiedAt (modified time, most recent first) and name (folder name, alphabetical). Optionally add :asc or :desc to specify ascending or descending order.  (optional) (default to "name")
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived folders. Use \"ANY_ARCHIVED\" to filter for archived folders regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of a folder. Restricts results to those with names that include the provided substring.  (optional)
    parentFolderId := "parentFolderId_example" // string | ID of a folder. Restricts results to those in the folder. Use \"NO_PARENT\" to filter for root folders. (optional)
    projectId := "projectId_example" // string | ID of a project. Restricts results to those in the project. (optional)
    ids := "lib_qQFY3WQH,lib_QvXryHdi,lib_3eF8mZjn" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    name := "name_example" // string | Name of a folder. Restricts results to those with the specified name. (optional)
    section := "INVENTORY" // string | INVENTORY or NOTEBOOK. Returns folders of inventory type or notebook type. Returns all folders if sections are merged.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FoldersApi.ListFolders(context.Background()).NextToken(nextToken).PageSize(pageSize).Sort(sort).ArchiveReason(archiveReason).NameIncludes(nameIncludes).ParentFolderId(parentFolderId).ProjectId(projectId).Ids(ids).Name(name).Section(section).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.ListFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFolders`: FoldersPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.ListFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **sort** | **string** | Method by which to order search results. Valid sorts are modifiedAt (modified time, most recent first) and name (folder name, alphabetical). Optionally add :asc or :desc to specify ascending or descending order.  | [default to &quot;name&quot;]
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived folders. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived folders regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **nameIncludes** | **string** | Name substring of a folder. Restricts results to those with names that include the provided substring.  | 
 **parentFolderId** | **string** | ID of a folder. Restricts results to those in the folder. Use \&quot;NO_PARENT\&quot; to filter for root folders. | 
 **projectId** | **string** | ID of a project. Restricts results to those in the project. | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **name** | **string** | Name of a folder. Restricts results to those with the specified name. | 
 **section** | **string** | INVENTORY or NOTEBOOK. Returns folders of inventory type or notebook type. Returns all folders if sections are merged.  | 

### Return type

[**FoldersPaginatedList**](FoldersPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveFolders

> FoldersArchivalChange UnarchiveFolders(ctx).FoldersUnarchive(foldersUnarchive).Execute()

Unarchive folders



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
    foldersUnarchive := *openapiclient.NewFoldersUnarchive([]string{"FolderIds_example"}) // FoldersUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FoldersApi.UnarchiveFolders(context.Background()).FoldersUnarchive(foldersUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.UnarchiveFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveFolders`: FoldersArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.UnarchiveFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **foldersUnarchive** | [**FoldersUnarchive**](FoldersUnarchive.md) |  | 

### Return type

[**FoldersArchivalChange**](FoldersArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

