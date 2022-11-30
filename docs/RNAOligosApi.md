# \RNAOligosApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveRNAOligos**](RNAOligosApi.md#ArchiveRNAOligos) | **Post** /rna-oligos:archive | Archive RNA Oligos
[**BulkCreateRNAOligos**](RNAOligosApi.md#BulkCreateRNAOligos) | **Post** /rna-oligos:bulk-create | Bulk Create RNA Oligos
[**BulkUpdateRNAOligos**](RNAOligosApi.md#BulkUpdateRNAOligos) | **Post** /rna-oligos:bulk-update | Bulk Update RNA Oligos
[**CreateRNAOligo**](RNAOligosApi.md#CreateRNAOligo) | **Post** /rna-oligos | Create an RNA Oligo
[**GetRNAOligo**](RNAOligosApi.md#GetRNAOligo) | **Get** /rna-oligos/{oligo_id} | Get an RNA Oligo
[**ListRNAOligos**](RNAOligosApi.md#ListRNAOligos) | **Get** /rna-oligos | List RNA Oligos
[**UnarchiveRNAOligos**](RNAOligosApi.md#UnarchiveRNAOligos) | **Post** /rna-oligos:unarchive | Unarchive RNA Oligos
[**UpdateRNAOligo**](RNAOligosApi.md#UpdateRNAOligo) | **Patch** /rna-oligos/{oligo_id} | Update an RNA Oligo



## ArchiveRNAOligos

> RnaOligosArchivalChange ArchiveRNAOligos(ctx).RnaOligosArchive(rnaOligosArchive).Execute()

Archive RNA Oligos



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
    rnaOligosArchive := *openapiclient.NewRnaOligosArchive(openapiclient.EntityArchiveReason("Made in error"), []string{"RnaOligoIds_example"}) // RnaOligosArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNAOligosApi.ArchiveRNAOligos(context.Background()).RnaOligosArchive(rnaOligosArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNAOligosApi.ArchiveRNAOligos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveRNAOligos`: RnaOligosArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `RNAOligosApi.ArchiveRNAOligos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveRNAOligosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rnaOligosArchive** | [**RnaOligosArchive**](RnaOligosArchive.md) |  | 

### Return type

[**RnaOligosArchivalChange**](RnaOligosArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkCreateRNAOligos

> AsyncTaskLink BulkCreateRNAOligos(ctx).RnaOligosBulkCreateRequest(rnaOligosBulkCreateRequest).Execute()

Bulk Create RNA Oligos



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
    rnaOligosBulkCreateRequest := *openapiclient.NewRnaOligosBulkCreateRequest() // RnaOligosBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNAOligosApi.BulkCreateRNAOligos(context.Background()).RnaOligosBulkCreateRequest(rnaOligosBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNAOligosApi.BulkCreateRNAOligos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateRNAOligos`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `RNAOligosApi.BulkCreateRNAOligos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateRNAOligosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rnaOligosBulkCreateRequest** | [**RnaOligosBulkCreateRequest**](RnaOligosBulkCreateRequest.md) |  | 

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


## BulkUpdateRNAOligos

> AsyncTaskLink BulkUpdateRNAOligos(ctx).RnaOligosBulkUpdateRequest(rnaOligosBulkUpdateRequest).Execute()

Bulk Update RNA Oligos



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
    rnaOligosBulkUpdateRequest := *openapiclient.NewRnaOligosBulkUpdateRequest() // RnaOligosBulkUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNAOligosApi.BulkUpdateRNAOligos(context.Background()).RnaOligosBulkUpdateRequest(rnaOligosBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNAOligosApi.BulkUpdateRNAOligos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateRNAOligos`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `RNAOligosApi.BulkUpdateRNAOligos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateRNAOligosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rnaOligosBulkUpdateRequest** | [**RnaOligosBulkUpdateRequest**](RnaOligosBulkUpdateRequest.md) |  | 

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


## CreateRNAOligo

> RnaOligo CreateRNAOligo(ctx).RnaOligoCreate(rnaOligoCreate).Execute()

Create an RNA Oligo



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
    rnaOligoCreate := *openapiclient.NewRnaOligoCreate("Bases_example", "Name_example") // RnaOligoCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNAOligosApi.CreateRNAOligo(context.Background()).RnaOligoCreate(rnaOligoCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNAOligosApi.CreateRNAOligo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRNAOligo`: RnaOligo
    fmt.Fprintf(os.Stdout, "Response from `RNAOligosApi.CreateRNAOligo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRNAOligoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rnaOligoCreate** | [**RnaOligoCreate**](RnaOligoCreate.md) |  | 

### Return type

[**RnaOligo**](RnaOligo.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRNAOligo

> RnaOligo GetRNAOligo(ctx, oligoId).Execute()

Get an RNA Oligo



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
    oligoId := "oligoId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNAOligosApi.GetRNAOligo(context.Background(), oligoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNAOligosApi.GetRNAOligo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRNAOligo`: RnaOligo
    fmt.Fprintf(os.Stdout, "Response from `RNAOligosApi.GetRNAOligo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oligoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRNAOligoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RnaOligo**](RnaOligo.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRNAOligos

> RnaOligosPaginatedList ListRNAOligos(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).Bases(bases).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).AuthorIdsAnyOf(authorIdsAnyOf).Execute()

List RNA Oligos



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
    pageSize := int32(56) // int32 | Number of results to return. Defaults to 50, maximum of 100.  (optional) (default to 50)
    nextToken := "nextToken_example" // string | Token for pagination (optional)
    sort := "sort_example" // string |  (optional) (default to "modifiedAt:desc")
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    name := "name_example" // string | Name of an RNA Oligo. Restricts results to those with the specified name, alias, or entity registry ID. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of an RNA Oligo. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring. (optional)
    bases := "bases_example" // string | Full bases of the RNA Oligo. Restricts results to those with the specified bases, case-insensitive, allowing for circular or reverse complement matches. Does not allow partial matching or loose matching via degenerate bases.  (optional)
    folderId := "folderId_example" // string | ID of a folder. Restricts results to those in the folder. (optional)
    mentionedIn := "mentionedIn_example" // string | Comma-separated list of entry IDs. Restricts results to RNA Oligos mentioned in those entries.  (optional)
    projectId := "projectId_example" // string | ID of a project. Restricts results to those in the project. (optional)
    registryId := "registryId_example" // string | ID of a registry. Restricts results to those registered in this registry. Specifying \"null\" returns unregistered items.  (optional)
    schemaId := "schemaId_example" // string | ID of a schema. Restricts results to those of the specified schema.  (optional)
    schemaFields := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Schema field value. For Integer, Float, and Date type fields, supports the >= and <= operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived RNA Oligos. Use \"ANY_ARCHIVED\" to filter for archived RNA Oligos regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    mentions := "mentions_example" // string | Comma-separated list of item IDs. Restricts results to those that mention the given items in the description.  (optional)
    ids := "seq_yWs5X7lv,seq_RhYGVnHF" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    entityRegistryIdsAnyOf := "TP001,TP002" // string | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of users IDs (optional)
    authorIdsAnyOf := "ent_a0SApq3z,ent_b4AApz9b" // string | Comma separated list of user or app IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNAOligosApi.ListRNAOligos(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).Bases(bases).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).AuthorIdsAnyOf(authorIdsAnyOf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNAOligosApi.ListRNAOligos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRNAOligos`: RnaOligosPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `RNAOligosApi.ListRNAOligos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRNAOligosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100.  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of an RNA Oligo. Restricts results to those with the specified name, alias, or entity registry ID. | 
 **nameIncludes** | **string** | Name substring of an RNA Oligo. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring. | 
 **bases** | **string** | Full bases of the RNA Oligo. Restricts results to those with the specified bases, case-insensitive, allowing for circular or reverse complement matches. Does not allow partial matching or loose matching via degenerate bases.  | 
 **folderId** | **string** | ID of a folder. Restricts results to those in the folder. | 
 **mentionedIn** | **string** | Comma-separated list of entry IDs. Restricts results to RNA Oligos mentioned in those entries.  | 
 **projectId** | **string** | ID of a project. Restricts results to those in the project. | 
 **registryId** | **string** | ID of a registry. Restricts results to those registered in this registry. Specifying \&quot;null\&quot; returns unregistered items.  | 
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema.  | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived RNA Oligos. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived RNA Oligos regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **mentions** | **string** | Comma-separated list of item IDs. Restricts results to those that mention the given items in the description.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **entityRegistryIdsAnyOf** | **string** | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  | 
 **creatorIds** | **string** | Comma separated list of users IDs | 
 **authorIdsAnyOf** | **string** | Comma separated list of user or app IDs | 

### Return type

[**RnaOligosPaginatedList**](RnaOligosPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveRNAOligos

> RnaOligosArchivalChange UnarchiveRNAOligos(ctx).RnaOligosUnarchive(rnaOligosUnarchive).Execute()

Unarchive RNA Oligos



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
    rnaOligosUnarchive := *openapiclient.NewRnaOligosUnarchive([]string{"RnaOligoIds_example"}) // RnaOligosUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNAOligosApi.UnarchiveRNAOligos(context.Background()).RnaOligosUnarchive(rnaOligosUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNAOligosApi.UnarchiveRNAOligos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveRNAOligos`: RnaOligosArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `RNAOligosApi.UnarchiveRNAOligos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveRNAOligosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rnaOligosUnarchive** | [**RnaOligosUnarchive**](RnaOligosUnarchive.md) |  | 

### Return type

[**RnaOligosArchivalChange**](RnaOligosArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRNAOligo

> RnaOligo UpdateRNAOligo(ctx, oligoId).RnaOligoUpdate(rnaOligoUpdate).Execute()

Update an RNA Oligo



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
    oligoId := "oligoId_example" // string | 
    rnaOligoUpdate := *openapiclient.NewRnaOligoUpdate() // RnaOligoUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNAOligosApi.UpdateRNAOligo(context.Background(), oligoId).RnaOligoUpdate(rnaOligoUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNAOligosApi.UpdateRNAOligo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRNAOligo`: RnaOligo
    fmt.Fprintf(os.Stdout, "Response from `RNAOligosApi.UpdateRNAOligo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oligoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRNAOligoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rnaOligoUpdate** | [**RnaOligoUpdate**](RnaOligoUpdate.md) |  | 

### Return type

[**RnaOligo**](RnaOligo.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

