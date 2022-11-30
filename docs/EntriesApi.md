# \EntriesApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveEntries**](EntriesApi.md#ArchiveEntries) | **Post** /entries:archive | Archive notebook entries
[**BulkGetEntries**](EntriesApi.md#BulkGetEntries) | **Get** /entries:bulk-get | Get notebook entries using entry IDs or display IDs
[**CreateEntry**](EntriesApi.md#CreateEntry) | **Post** /entries | Create a notebook entry
[**GetEntry**](EntriesApi.md#GetEntry) | **Get** /entries/{entry_id} | Get a notebook entry by ID
[**GetEntryTemplate**](EntriesApi.md#GetEntryTemplate) | **Get** /entry-templates/{entry_template_id} | Get a notebook template entry by ID
[**GetExternalFileMetadata**](EntriesApi.md#GetExternalFileMetadata) | **Get** /entries/{entry_id}/external-files/{external_file_id} | Retrieves the metadata for an external file. Use the &#39;downloadURL&#39; to download the actual file. 
[**ListEntries**](EntriesApi.md#ListEntries) | **Get** /entries | List entries
[**ListEntryTemplates**](EntriesApi.md#ListEntryTemplates) | **Get** /entry-templates | List entry templates
[**UnarchiveEntries**](EntriesApi.md#UnarchiveEntries) | **Post** /entries:unarchive | Unarchive notebook entries
[**UpdateEntry**](EntriesApi.md#UpdateEntry) | **Patch** /entries/{entry_id} | Update a notebook entry&#39;s metadata



## ArchiveEntries

> EntriesArchivalChange ArchiveEntries(ctx).EntriesArchive(entriesArchive).Execute()

Archive notebook entries



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
    entriesArchive := *openapiclient.NewEntriesArchive([]string{"EntryIds_example"}, "Reason_example") // EntriesArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntriesApi.ArchiveEntries(context.Background()).EntriesArchive(entriesArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntriesApi.ArchiveEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveEntries`: EntriesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `EntriesApi.ArchiveEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entriesArchive** | [**EntriesArchive**](EntriesArchive.md) |  | 

### Return type

[**EntriesArchivalChange**](EntriesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkGetEntries

> Entries BulkGetEntries(ctx).EntryIds(entryIds).DisplayIds(displayIds).Execute()

Get notebook entries using entry IDs or display IDs



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
    entryIds := "entryIds_example" // string | Comma-separated list of Entry IDs. (optional)
    displayIds := "displayIds_example" // string | Comma-separated list of Entry Display IDs. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntriesApi.BulkGetEntries(context.Background()).EntryIds(entryIds).DisplayIds(displayIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntriesApi.BulkGetEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetEntries`: Entries
    fmt.Fprintf(os.Stdout, "Response from `EntriesApi.BulkGetEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entryIds** | **string** | Comma-separated list of Entry IDs. | 
 **displayIds** | **string** | Comma-separated list of Entry Display IDs. | 

### Return type

[**Entries**](Entries.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEntry

> Entry CreateEntry(ctx).EntryCreate(entryCreate).Execute()

Create a notebook entry



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
    entryCreate := *openapiclient.NewEntryCreate("FolderId_example", "Name_example") // EntryCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntriesApi.CreateEntry(context.Background()).EntryCreate(entryCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntriesApi.CreateEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEntry`: Entry
    fmt.Fprintf(os.Stdout, "Response from `EntriesApi.CreateEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entryCreate** | [**EntryCreate**](EntryCreate.md) |  | 

### Return type

[**Entry**](Entry.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntry

> EntryById GetEntry(ctx, entryId).Execute()

Get a notebook entry by ID



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
    entryId := "entryId_example" // string | ID of the entry

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntriesApi.GetEntry(context.Background(), entryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntriesApi.GetEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntry`: EntryById
    fmt.Fprintf(os.Stdout, "Response from `EntriesApi.GetEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | ID of the entry | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntryById**](EntryById.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntryTemplate

> EntryTemplate GetEntryTemplate(ctx, entryTemplateId).Execute()

Get a notebook template entry by ID



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
    entryTemplateId := "entryTemplateId_example" // string | ID of the entry template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntriesApi.GetEntryTemplate(context.Background(), entryTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntriesApi.GetEntryTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntryTemplate`: EntryTemplate
    fmt.Fprintf(os.Stdout, "Response from `EntriesApi.GetEntryTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryTemplateId** | **string** | ID of the entry template | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntryTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntryTemplate**](EntryTemplate.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalFileMetadata

> EntryExternalFileById GetExternalFileMetadata(ctx, entryId, externalFileId).Execute()

Retrieves the metadata for an external file. Use the 'downloadURL' to download the actual file. 



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
    entryId := "entryId_example" // string | ID of the entry the file was uploaded to
    externalFileId := "externalFileId_example" // string | ID of the external file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntriesApi.GetExternalFileMetadata(context.Background(), entryId, externalFileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntriesApi.GetExternalFileMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalFileMetadata`: EntryExternalFileById
    fmt.Fprintf(os.Stdout, "Response from `EntriesApi.GetExternalFileMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | ID of the entry the file was uploaded to | 
**externalFileId** | **string** | ID of the external file | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalFileMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EntryExternalFileById**](EntryExternalFileById.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntries

> EntriesPaginatedList ListEntries(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).ProjectId(projectId).ArchiveReason(archiveReason).ReviewStatus(reviewStatus).MentionedIn(mentionedIn).Mentions(mentions).Ids(ids).SchemaId(schemaId).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).AssignedReviewerIdsAnyOf(assignedReviewerIdsAnyOf).CreatorIds(creatorIds).AuthorIdsAnyOf(authorIdsAnyOf).DisplayIds(displayIds).Execute()

List entries



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
    pageSize := int32(56) // int32 | Number of results to return. Defaults to 50, maximum of 100. (optional) (default to 50)
    nextToken := "nextToken_example" // string | Token for pagination (optional)
    sort := "sort_example" // string | Method by which to order search results. Valid sorts are modifiedAt (modified time, most recent first) and name (entity name, alphabetical). Optionally add :asc or :desc to specify ascending or descending order.  (optional) (default to "modifiedAt:desc")
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > operator. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    name := "name_example" // string | Name of an Entry. Restricts results to those with the specified name. (optional)
    projectId := "projectId_example" // string | ID of a project. Restricts results to those in the project. (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived entries. Use \"ANY_ARCHIVED\" to filter for archived entries regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    reviewStatus := "reviewStatus_example" // string | Restrict results to those with the given review status. Supported statuses: IN_PROGRESS, ACCEPTED, REJECTED, NEEDS_REVIEW, RETRACTED  (optional)
    mentionedIn := "mentionedIn_example" // string | Comma-separated list of entry IDs. Restricts results to those mentioned within the entries in this list.  (optional)
    mentions := "mentions_example" // string | Comma-separated list of resource IDs. Restricts results to entries that mention the given items.  (optional)
    ids := "ids_example" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    schemaId := "schemaId_example" // string | ID of a schema. Restricts results to those of the specified schema.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  (optional)
    assignedReviewerIdsAnyOf := "ent_a0SApq3z,ent_SdUvia1v" // string | Comma-separated list of ids. Restricts results to entries that have assignees of any of the specified ids. (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of users IDs (optional)
    authorIdsAnyOf := "ent_a0SApq3z,ent_b4AApz9b" // string | Comma separated list of user or app IDs (optional)
    displayIds := "VPR001,VPR002" // string | Comma-separated list of Entry Display IDs. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntriesApi.ListEntries(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).ProjectId(projectId).ArchiveReason(archiveReason).ReviewStatus(reviewStatus).MentionedIn(mentionedIn).Mentions(mentions).Ids(ids).SchemaId(schemaId).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).AssignedReviewerIdsAnyOf(assignedReviewerIdsAnyOf).CreatorIds(creatorIds).AuthorIdsAnyOf(authorIdsAnyOf).DisplayIds(displayIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntriesApi.ListEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntries`: EntriesPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `EntriesApi.ListEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100. | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** | Method by which to order search results. Valid sorts are modifiedAt (modified time, most recent first) and name (entity name, alphabetical). Optionally add :asc or :desc to specify ascending or descending order.  | [default to &quot;modifiedAt:desc&quot;]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; operator. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of an Entry. Restricts results to those with the specified name. | 
 **projectId** | **string** | ID of a project. Restricts results to those in the project. | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived entries. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived entries regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **reviewStatus** | **string** | Restrict results to those with the given review status. Supported statuses: IN_PROGRESS, ACCEPTED, REJECTED, NEEDS_REVIEW, RETRACTED  | 
 **mentionedIn** | **string** | Comma-separated list of entry IDs. Restricts results to those mentioned within the entries in this list.  | 
 **mentions** | **string** | Comma-separated list of resource IDs. Restricts results to entries that mention the given items.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  | 
 **assignedReviewerIdsAnyOf** | **string** | Comma-separated list of ids. Restricts results to entries that have assignees of any of the specified ids. | 
 **creatorIds** | **string** | Comma separated list of users IDs | 
 **authorIdsAnyOf** | **string** | Comma separated list of user or app IDs | 
 **displayIds** | **string** | Comma-separated list of Entry Display IDs. | 

### Return type

[**EntriesPaginatedList**](EntriesPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntryTemplates

> EntryTemplatesPaginatedList ListEntryTemplates(ctx).PageSize(pageSize).NextToken(nextToken).ModifiedAt(modifiedAt).Name(name).TemplateCollectionId(templateCollectionId).Ids(ids).SchemaId(schemaId).Execute()

List entry templates



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
    pageSize := int32(56) // int32 | Number of results to return. Defaults to 50, maximum of 100. (optional) (default to 50)
    nextToken := "nextToken_example" // string | Token for pagination (optional)
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    name := "name_example" // string | Name of an Entry Template. Restricts results to those with the specified name. (optional)
    templateCollectionId := "tmplcol_jC7rOniv" // string | ID of a template collection. Resticts results to those in the template collection.  (optional)
    ids := "ids_example" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    schemaId := "schemaId_example" // string | ID of a schema. Restricts results to those of the specified schema.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntriesApi.ListEntryTemplates(context.Background()).PageSize(pageSize).NextToken(nextToken).ModifiedAt(modifiedAt).Name(name).TemplateCollectionId(templateCollectionId).Ids(ids).SchemaId(schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntriesApi.ListEntryTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntryTemplates`: EntryTemplatesPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `EntriesApi.ListEntryTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEntryTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100. | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of an Entry Template. Restricts results to those with the specified name. | 
 **templateCollectionId** | **string** | ID of a template collection. Resticts results to those in the template collection.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema.  | 

### Return type

[**EntryTemplatesPaginatedList**](EntryTemplatesPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveEntries

> EntriesArchivalChange UnarchiveEntries(ctx).EntriesUnarchive(entriesUnarchive).Execute()

Unarchive notebook entries



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
    entriesUnarchive := *openapiclient.NewEntriesUnarchive([]string{"EntryIds_example"}) // EntriesUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntriesApi.UnarchiveEntries(context.Background()).EntriesUnarchive(entriesUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntriesApi.UnarchiveEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveEntries`: EntriesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `EntriesApi.UnarchiveEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entriesUnarchive** | [**EntriesUnarchive**](EntriesUnarchive.md) |  | 

### Return type

[**EntriesArchivalChange**](EntriesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntry

> Entry UpdateEntry(ctx, entryId).EntryUpdate(entryUpdate).Execute()

Update a notebook entry's metadata



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
    entryId := "entryId_example" // string | ID of the entry
    entryUpdate := *openapiclient.NewEntryUpdate() // EntryUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntriesApi.UpdateEntry(context.Background(), entryId).EntryUpdate(entryUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntriesApi.UpdateEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntry`: Entry
    fmt.Fprintf(os.Stdout, "Response from `EntriesApi.UpdateEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | ID of the entry | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entryUpdate** | [**EntryUpdate**](EntryUpdate.md) |  | 

### Return type

[**Entry**](Entry.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

