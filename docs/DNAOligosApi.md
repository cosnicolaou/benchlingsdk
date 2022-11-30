# \DNAOligosApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveDNAOligos**](DNAOligosApi.md#ArchiveDNAOligos) | **Post** /dna-oligos:archive | Archive DNA Oligos
[**BulkCreateDNAOligos**](DNAOligosApi.md#BulkCreateDNAOligos) | **Post** /dna-oligos:bulk-create | Bulk Create DNA Oligos
[**BulkUpdateDNAOligos**](DNAOligosApi.md#BulkUpdateDNAOligos) | **Post** /dna-oligos:bulk-update | Bulk Update DNA Oligos
[**CreateDNAOligo**](DNAOligosApi.md#CreateDNAOligo) | **Post** /dna-oligos | Create a DNA Oligo
[**GetDNAOligo**](DNAOligosApi.md#GetDNAOligo) | **Get** /dna-oligos/{oligo_id} | Get a DNA Oligo
[**ListDNAOligos**](DNAOligosApi.md#ListDNAOligos) | **Get** /dna-oligos | List DNA Oligos
[**UnarchiveDNAOligos**](DNAOligosApi.md#UnarchiveDNAOligos) | **Post** /dna-oligos:unarchive | Unarchive DNA Oligos
[**UpdateDNAOligo**](DNAOligosApi.md#UpdateDNAOligo) | **Patch** /dna-oligos/{oligo_id} | Update a DNA Oligo



## ArchiveDNAOligos

> DnaOligosArchivalChange ArchiveDNAOligos(ctx).DnaOligosArchive(dnaOligosArchive).Execute()

Archive DNA Oligos



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
    dnaOligosArchive := *openapiclient.NewDnaOligosArchive([]string{"DnaOligoIds_example"}, openapiclient.EntityArchiveReason("Made in error")) // DnaOligosArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNAOligosApi.ArchiveDNAOligos(context.Background()).DnaOligosArchive(dnaOligosArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNAOligosApi.ArchiveDNAOligos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveDNAOligos`: DnaOligosArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `DNAOligosApi.ArchiveDNAOligos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveDNAOligosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnaOligosArchive** | [**DnaOligosArchive**](DnaOligosArchive.md) |  | 

### Return type

[**DnaOligosArchivalChange**](DnaOligosArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkCreateDNAOligos

> AsyncTaskLink BulkCreateDNAOligos(ctx).DnaOligosBulkCreateRequest(dnaOligosBulkCreateRequest).Execute()

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
    dnaOligosBulkCreateRequest := *openapiclient.NewDnaOligosBulkCreateRequest() // DnaOligosBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNAOligosApi.BulkCreateDNAOligos(context.Background()).DnaOligosBulkCreateRequest(dnaOligosBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNAOligosApi.BulkCreateDNAOligos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateDNAOligos`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `DNAOligosApi.BulkCreateDNAOligos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateDNAOligosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnaOligosBulkCreateRequest** | [**DnaOligosBulkCreateRequest**](DnaOligosBulkCreateRequest.md) |  | 

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


## BulkUpdateDNAOligos

> AsyncTaskLink BulkUpdateDNAOligos(ctx).DnaOligosBulkUpdateRequest(dnaOligosBulkUpdateRequest).Execute()

Bulk Update DNA Oligos



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
    dnaOligosBulkUpdateRequest := *openapiclient.NewDnaOligosBulkUpdateRequest() // DnaOligosBulkUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNAOligosApi.BulkUpdateDNAOligos(context.Background()).DnaOligosBulkUpdateRequest(dnaOligosBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNAOligosApi.BulkUpdateDNAOligos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateDNAOligos`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `DNAOligosApi.BulkUpdateDNAOligos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateDNAOligosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnaOligosBulkUpdateRequest** | [**DnaOligosBulkUpdateRequest**](DnaOligosBulkUpdateRequest.md) |  | 

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


## CreateDNAOligo

> DnaOligo CreateDNAOligo(ctx).DnaOligoCreate(dnaOligoCreate).Execute()

Create a DNA Oligo



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
    dnaOligoCreate := *openapiclient.NewDnaOligoCreate("Bases_example", "Name_example") // DnaOligoCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNAOligosApi.CreateDNAOligo(context.Background()).DnaOligoCreate(dnaOligoCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNAOligosApi.CreateDNAOligo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDNAOligo`: DnaOligo
    fmt.Fprintf(os.Stdout, "Response from `DNAOligosApi.CreateDNAOligo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDNAOligoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnaOligoCreate** | [**DnaOligoCreate**](DnaOligoCreate.md) |  | 

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


## GetDNAOligo

> DnaOligo GetDNAOligo(ctx, oligoId).Execute()

Get a DNA Oligo



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
    resp, r, err := apiClient.DNAOligosApi.GetDNAOligo(context.Background(), oligoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNAOligosApi.GetDNAOligo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDNAOligo`: DnaOligo
    fmt.Fprintf(os.Stdout, "Response from `DNAOligosApi.GetDNAOligo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oligoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDNAOligoRequest struct via the builder pattern


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


## ListDNAOligos

> DnaOligosPaginatedList ListDNAOligos(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).Bases(bases).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).AuthorIdsAnyOf(authorIdsAnyOf).Execute()

List DNA Oligos



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
    name := "name_example" // string | Name of a DNA Oligo. Restricts results to those with the specified name, alias, or entity registry ID. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of a DNA Oligo. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring. (optional)
    bases := "bases_example" // string | Full bases of the DNA Oligo. Restricts results to those with the specified bases, case-insensitive, allowing for circular or reverse complement matches. Does not allow partial matching or loose matching via degenerate bases.  (optional)
    folderId := "folderId_example" // string | ID of a folder. Restricts results to those in the folder. (optional)
    mentionedIn := "mentionedIn_example" // string | Comma-separated list of entry IDs. Restricts results to DNA Oligos mentioned in those entries.  (optional)
    projectId := "projectId_example" // string | ID of a project. Restricts results to those in the project. (optional)
    registryId := "registryId_example" // string | ID of a registry. Restricts results to those registered in this registry. Specifying \"null\" returns unregistered items.  (optional)
    schemaId := "schemaId_example" // string | ID of a schema. Restricts results to those of the specified schema.  (optional)
    schemaFields := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Schema field value. For Integer, Float, and Date type fields, supports the >= and <= operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived DNA Oligos. Use \"ANY_ARCHIVED\" to filter for archived DNA Oligos regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    mentions := "mentions_example" // string | Comma-separated list of item IDs. Restricts results to those that mention the given items in the description.  (optional)
    ids := "seq_yWs5X7lv,seq_RhYGVnHF" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    entityRegistryIdsAnyOf := "TP001,TP002" // string | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of users IDs (optional)
    authorIdsAnyOf := "ent_a0SApq3z,ent_b4AApz9b" // string | Comma separated list of user or app IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNAOligosApi.ListDNAOligos(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).Bases(bases).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).AuthorIdsAnyOf(authorIdsAnyOf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNAOligosApi.ListDNAOligos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDNAOligos`: DnaOligosPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DNAOligosApi.ListDNAOligos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDNAOligosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100.  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of a DNA Oligo. Restricts results to those with the specified name, alias, or entity registry ID. | 
 **nameIncludes** | **string** | Name substring of a DNA Oligo. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring. | 
 **bases** | **string** | Full bases of the DNA Oligo. Restricts results to those with the specified bases, case-insensitive, allowing for circular or reverse complement matches. Does not allow partial matching or loose matching via degenerate bases.  | 
 **folderId** | **string** | ID of a folder. Restricts results to those in the folder. | 
 **mentionedIn** | **string** | Comma-separated list of entry IDs. Restricts results to DNA Oligos mentioned in those entries.  | 
 **projectId** | **string** | ID of a project. Restricts results to those in the project. | 
 **registryId** | **string** | ID of a registry. Restricts results to those registered in this registry. Specifying \&quot;null\&quot; returns unregistered items.  | 
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema.  | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived DNA Oligos. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived DNA Oligos regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **mentions** | **string** | Comma-separated list of item IDs. Restricts results to those that mention the given items in the description.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **entityRegistryIdsAnyOf** | **string** | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  | 
 **creatorIds** | **string** | Comma separated list of users IDs | 
 **authorIdsAnyOf** | **string** | Comma separated list of user or app IDs | 

### Return type

[**DnaOligosPaginatedList**](DnaOligosPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveDNAOligos

> DnaOligosArchivalChange UnarchiveDNAOligos(ctx).DnaOligosUnarchive(dnaOligosUnarchive).Execute()

Unarchive DNA Oligos



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
    dnaOligosUnarchive := *openapiclient.NewDnaOligosUnarchive([]string{"DnaOligoIds_example"}) // DnaOligosUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNAOligosApi.UnarchiveDNAOligos(context.Background()).DnaOligosUnarchive(dnaOligosUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNAOligosApi.UnarchiveDNAOligos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveDNAOligos`: DnaOligosArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `DNAOligosApi.UnarchiveDNAOligos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveDNAOligosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnaOligosUnarchive** | [**DnaOligosUnarchive**](DnaOligosUnarchive.md) |  | 

### Return type

[**DnaOligosArchivalChange**](DnaOligosArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDNAOligo

> DnaOligo UpdateDNAOligo(ctx, oligoId).DnaOligoUpdate(dnaOligoUpdate).Execute()

Update a DNA Oligo



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
    dnaOligoUpdate := *openapiclient.NewDnaOligoUpdate() // DnaOligoUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNAOligosApi.UpdateDNAOligo(context.Background(), oligoId).DnaOligoUpdate(dnaOligoUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNAOligosApi.UpdateDNAOligo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDNAOligo`: DnaOligo
    fmt.Fprintf(os.Stdout, "Response from `DNAOligosApi.UpdateDNAOligo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oligoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDNAOligoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnaOligoUpdate** | [**DnaOligoUpdate**](DnaOligoUpdate.md) |  | 

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

