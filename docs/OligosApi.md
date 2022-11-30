# \OligosApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveOligos**](OligosApi.md#ArchiveOligos) | **Post** /oligos:archive | Archive Oligos
[**BulkCreateOligos**](OligosApi.md#BulkCreateOligos) | **Post** /oligos:bulk-create | Bulk Create DNA Oligos
[**BulkGetOligos**](OligosApi.md#BulkGetOligos) | **Get** /oligos:bulk-get | Bulk get Oligos by ID
[**CreateOligo**](OligosApi.md#CreateOligo) | **Post** /oligos | Create an Oligo
[**GetOligo**](OligosApi.md#GetOligo) | **Get** /oligos/{oligo_id} | Get an Oligo
[**ListOligos**](OligosApi.md#ListOligos) | **Get** /oligos | List Oligos
[**UnarchiveOligos**](OligosApi.md#UnarchiveOligos) | **Post** /oligos:unarchive | Unarchive Oligos
[**UpdateOligo**](OligosApi.md#UpdateOligo) | **Patch** /oligos/{oligo_id} | Update an Oligo



## ArchiveOligos

> OligosArchivalChange ArchiveOligos(ctx).OligosArchive(oligosArchive).Execute()

Archive Oligos



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
    oligosArchive := *openapiclient.NewOligosArchive([]string{"OligoIds_example"}, openapiclient.EntityArchiveReason("Made in error")) // OligosArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OligosApi.ArchiveOligos(context.Background()).OligosArchive(oligosArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OligosApi.ArchiveOligos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveOligos`: OligosArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `OligosApi.ArchiveOligos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveOligosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oligosArchive** | [**OligosArchive**](OligosArchive.md) |  | 

### Return type

[**OligosArchivalChange**](OligosArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkCreateOligos

> AsyncTaskLink BulkCreateOligos(ctx).OligosBulkCreateRequest(oligosBulkCreateRequest).Execute()

Bulk Create DNA Oligos



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
    oligosBulkCreateRequest := *openapiclient.NewOligosBulkCreateRequest() // OligosBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OligosApi.BulkCreateOligos(context.Background()).OligosBulkCreateRequest(oligosBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OligosApi.BulkCreateOligos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateOligos`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `OligosApi.BulkCreateOligos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateOligosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oligosBulkCreateRequest** | [**OligosBulkCreateRequest**](OligosBulkCreateRequest.md) |  | 

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


## BulkGetOligos

> OligosBulkGet BulkGetOligos(ctx).OligoIds(oligoIds).Execute()

Bulk get Oligos by ID



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
    oligoIds := "oligoIds_example" // string | Comma-separated list of IDs of Oligos to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OligosApi.BulkGetOligos(context.Background()).OligoIds(oligoIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OligosApi.BulkGetOligos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetOligos`: OligosBulkGet
    fmt.Fprintf(os.Stdout, "Response from `OligosApi.BulkGetOligos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetOligosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oligoIds** | **string** | Comma-separated list of IDs of Oligos to get.  | 

### Return type

[**OligosBulkGet**](OligosBulkGet.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOligo

> DnaOligo CreateOligo(ctx).OligoCreate(oligoCreate).Execute()

Create an Oligo



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
    oligoCreate := *openapiclient.NewOligoCreate("Bases_example", "Name_example") // OligoCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OligosApi.CreateOligo(context.Background()).OligoCreate(oligoCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OligosApi.CreateOligo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOligo`: DnaOligo
    fmt.Fprintf(os.Stdout, "Response from `OligosApi.CreateOligo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOligoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oligoCreate** | [**OligoCreate**](OligoCreate.md) |  | 

### Return type

[**DnaOligo**](DnaOligo.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOligo

> DnaOligo GetOligo(ctx, oligoId).Execute()

Get an Oligo



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
    resp, r, err := apiClient.OligosApi.GetOligo(context.Background(), oligoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OligosApi.GetOligo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOligo`: DnaOligo
    fmt.Fprintf(os.Stdout, "Response from `OligosApi.GetOligo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oligoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOligoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DnaOligo**](DnaOligo.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOligos

> OligosPaginatedList ListOligos(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).Bases(bases).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).Execute()

List Oligos



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
    name := "name_example" // string | Name of an Oligo. Restricts results to those with the specified name, alias, or entity registry ID. (optional)
    bases := "bases_example" // string | Full bases of the oligo. Restricts results to those with the specified bases, case-insensitive, allowing for circular or reverse complement matches. Does not allow partial matching or loose matching via degenerate bases.  (optional)
    folderId := "folderId_example" // string | ID of a folder. Restricts results to those in the folder. (optional)
    mentionedIn := "mentionedIn_example" // string | Comma-separated list of entry IDs. Restricts results to Oligos mentioned in those entries.  (optional)
    projectId := "projectId_example" // string | ID of a project. Restricts results to those in the project. (optional)
    registryId := "registryId_example" // string | ID of a registry. Restricts results to those registered in this registry. Specifying \"null\" returns unregistered items.  (optional)
    schemaId := "schemaId_example" // string | ID of a schema. Restricts results to those of the specified schema.  (optional)
    schemaFields := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Schema field value. For Integer, Float, and Date type fields, supports the >= and <= operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived Oligos. Use \"ANY_ARCHIVED\" to filter for archived Oligos regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    mentions := "mentions_example" // string | Comma-separated list of item IDs. Restricts results to those that mention the given items in the description.  (optional)
    ids := "seq_yWs5X7lv,seq_RhYGVnHF" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    entityRegistryIdsAnyOf := "TP001,TP002" // string | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of users IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OligosApi.ListOligos(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).Bases(bases).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OligosApi.ListOligos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOligos`: OligosPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `OligosApi.ListOligos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOligosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100.  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of an Oligo. Restricts results to those with the specified name, alias, or entity registry ID. | 
 **bases** | **string** | Full bases of the oligo. Restricts results to those with the specified bases, case-insensitive, allowing for circular or reverse complement matches. Does not allow partial matching or loose matching via degenerate bases.  | 
 **folderId** | **string** | ID of a folder. Restricts results to those in the folder. | 
 **mentionedIn** | **string** | Comma-separated list of entry IDs. Restricts results to Oligos mentioned in those entries.  | 
 **projectId** | **string** | ID of a project. Restricts results to those in the project. | 
 **registryId** | **string** | ID of a registry. Restricts results to those registered in this registry. Specifying \&quot;null\&quot; returns unregistered items.  | 
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema.  | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived Oligos. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived Oligos regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **mentions** | **string** | Comma-separated list of item IDs. Restricts results to those that mention the given items in the description.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **entityRegistryIdsAnyOf** | **string** | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  | 
 **creatorIds** | **string** | Comma separated list of users IDs | 

### Return type

[**OligosPaginatedList**](OligosPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveOligos

> OligosArchivalChange UnarchiveOligos(ctx).OligosUnarchive(oligosUnarchive).Execute()

Unarchive Oligos



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
    oligosUnarchive := *openapiclient.NewOligosUnarchive([]string{"OligoIds_example"}) // OligosUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OligosApi.UnarchiveOligos(context.Background()).OligosUnarchive(oligosUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OligosApi.UnarchiveOligos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveOligos`: OligosArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `OligosApi.UnarchiveOligos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveOligosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oligosUnarchive** | [**OligosUnarchive**](OligosUnarchive.md) |  | 

### Return type

[**OligosArchivalChange**](OligosArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOligo

> DnaOligo UpdateOligo(ctx, oligoId).OligoUpdate(oligoUpdate).Execute()

Update an Oligo



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
    oligoUpdate := *openapiclient.NewOligoUpdate() // OligoUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OligosApi.UpdateOligo(context.Background(), oligoId).OligoUpdate(oligoUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OligosApi.UpdateOligo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOligo`: DnaOligo
    fmt.Fprintf(os.Stdout, "Response from `OligosApi.UpdateOligo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oligoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOligoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oligoUpdate** | [**OligoUpdate**](OligoUpdate.md) |  | 

### Return type

[**DnaOligo**](DnaOligo.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

