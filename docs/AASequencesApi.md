# \AASequencesApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveAASequences**](AASequencesApi.md#ArchiveAASequences) | **Post** /aa-sequences:archive | Archive AA sequences
[**AutoAnnotateAaSequences**](AASequencesApi.md#AutoAnnotateAaSequences) | **Post** /aa-sequences:auto-annotate | Auto-annotate AA sequences with matching features from specified Feature Libraries
[**BulkCreateAASequences**](AASequencesApi.md#BulkCreateAASequences) | **Post** /aa-sequences:bulk-create | Bulk Create AA sequences
[**BulkGetAASequences**](AASequencesApi.md#BulkGetAASequences) | **Get** /aa-sequences:bulk-get | Bulk get AA sequences by ID
[**BulkUpdateAASequences**](AASequencesApi.md#BulkUpdateAASequences) | **Post** /aa-sequences:bulk-update | Bulk Update AA sequences
[**CreateAASequence**](AASequencesApi.md#CreateAASequence) | **Post** /aa-sequences | Create an AA sequence
[**GetAASequence**](AASequencesApi.md#GetAASequence) | **Get** /aa-sequences/{aa_sequence_id} | Get an AA sequence
[**ListAASequences**](AASequencesApi.md#ListAASequences) | **Get** /aa-sequences | List AA sequences
[**UnarchiveAASequences**](AASequencesApi.md#UnarchiveAASequences) | **Post** /aa-sequences:unarchive | Unarchive AA sequences
[**UpdateAASequence**](AASequencesApi.md#UpdateAASequence) | **Patch** /aa-sequences/{aa_sequence_id} | Update an AA sequence



## ArchiveAASequences

> AaSequencesArchivalChange ArchiveAASequences(ctx).AaSequencesArchive(aaSequencesArchive).Execute()

Archive AA sequences



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
    aaSequencesArchive := *openapiclient.NewAaSequencesArchive([]string{"AaSequenceIds_example"}, openapiclient.EntityArchiveReason("Made in error")) // AaSequencesArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AASequencesApi.ArchiveAASequences(context.Background()).AaSequencesArchive(aaSequencesArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AASequencesApi.ArchiveAASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveAASequences`: AaSequencesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `AASequencesApi.ArchiveAASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveAASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aaSequencesArchive** | [**AaSequencesArchive**](AaSequencesArchive.md) |  | 

### Return type

[**AaSequencesArchivalChange**](AaSequencesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoAnnotateAaSequences

> AsyncTaskLink AutoAnnotateAaSequences(ctx).AutoAnnotateAaSequences(autoAnnotateAaSequences).Execute()

Auto-annotate AA sequences with matching features from specified Feature Libraries



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
    autoAnnotateAaSequences := *openapiclient.NewAutoAnnotateAaSequences([]string{"AaSequenceIds_example"}, []string{"FeatureLibraryIds_example"}) // AutoAnnotateAaSequences |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AASequencesApi.AutoAnnotateAaSequences(context.Background()).AutoAnnotateAaSequences(autoAnnotateAaSequences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AASequencesApi.AutoAnnotateAaSequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutoAnnotateAaSequences`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `AASequencesApi.AutoAnnotateAaSequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoAnnotateAaSequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoAnnotateAaSequences** | [**AutoAnnotateAaSequences**](AutoAnnotateAaSequences.md) |  | 

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


## BulkCreateAASequences

> AsyncTaskLink BulkCreateAASequences(ctx).AaSequencesBulkCreateRequest(aaSequencesBulkCreateRequest).Execute()

Bulk Create AA sequences



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
    aaSequencesBulkCreateRequest := *openapiclient.NewAaSequencesBulkCreateRequest() // AaSequencesBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AASequencesApi.BulkCreateAASequences(context.Background()).AaSequencesBulkCreateRequest(aaSequencesBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AASequencesApi.BulkCreateAASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateAASequences`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `AASequencesApi.BulkCreateAASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateAASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aaSequencesBulkCreateRequest** | [**AaSequencesBulkCreateRequest**](AaSequencesBulkCreateRequest.md) |  | 

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


## BulkGetAASequences

> AaSequencesBulkGet BulkGetAASequences(ctx).AaSequenceIds(aaSequenceIds).Execute()

Bulk get AA sequences by ID



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
    aaSequenceIds := "aaSequenceIds_example" // string | Comma-separated list of IDs of AA sequences to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AASequencesApi.BulkGetAASequences(context.Background()).AaSequenceIds(aaSequenceIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AASequencesApi.BulkGetAASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetAASequences`: AaSequencesBulkGet
    fmt.Fprintf(os.Stdout, "Response from `AASequencesApi.BulkGetAASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetAASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aaSequenceIds** | **string** | Comma-separated list of IDs of AA sequences to get.  | 

### Return type

[**AaSequencesBulkGet**](AaSequencesBulkGet.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateAASequences

> AsyncTaskLink BulkUpdateAASequences(ctx).AaSequencesBulkUpdateRequest(aaSequencesBulkUpdateRequest).Execute()

Bulk Update AA sequences



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
    aaSequencesBulkUpdateRequest := *openapiclient.NewAaSequencesBulkUpdateRequest() // AaSequencesBulkUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AASequencesApi.BulkUpdateAASequences(context.Background()).AaSequencesBulkUpdateRequest(aaSequencesBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AASequencesApi.BulkUpdateAASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateAASequences`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `AASequencesApi.BulkUpdateAASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateAASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aaSequencesBulkUpdateRequest** | [**AaSequencesBulkUpdateRequest**](AaSequencesBulkUpdateRequest.md) |  | 

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


## CreateAASequence

> AaSequence CreateAASequence(ctx).AaSequenceCreate(aaSequenceCreate).Execute()

Create an AA sequence



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
    aaSequenceCreate := *openapiclient.NewAaSequenceCreate("AminoAcids_example", "Name_example") // AaSequenceCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AASequencesApi.CreateAASequence(context.Background()).AaSequenceCreate(aaSequenceCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AASequencesApi.CreateAASequence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAASequence`: AaSequence
    fmt.Fprintf(os.Stdout, "Response from `AASequencesApi.CreateAASequence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAASequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aaSequenceCreate** | [**AaSequenceCreate**](AaSequenceCreate.md) |  | 

### Return type

[**AaSequence**](AaSequence.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAASequence

> AaSequence GetAASequence(ctx, aaSequenceId).Execute()

Get an AA sequence



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
    aaSequenceId := "aaSequenceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AASequencesApi.GetAASequence(context.Background(), aaSequenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AASequencesApi.GetAASequence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAASequence`: AaSequence
    fmt.Fprintf(os.Stdout, "Response from `AASequencesApi.GetAASequence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aaSequenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAASequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AaSequence**](AaSequence.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAASequences

> AaSequencesPaginatedList ListAASequences(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).AminoAcids(aminoAcids).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).AuthorIdsAnyOf(authorIdsAnyOf).Execute()

List AA sequences



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
    name := "name_example" // string | Name of an AA Sequence. Restricts results to those with the specified name, alias, or entity registry ID. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of an AA sequence. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring. (optional)
    aminoAcids := "aminoAcids_example" // string | String of amino acids. Restricts results to AA sequences exactly matching these amino acids (case-insensitive). (optional)
    folderId := "folderId_example" // string | ID of a folder. Restricts results to those in the folder. (optional)
    mentionedIn := "mentionedIn_example" // string | Comma-separated list of entry IDs. Restricts results to AA sequences mentioned in those entries.  (optional)
    projectId := "projectId_example" // string | ID of a project. Restricts results to those in the project. (optional)
    registryId := "registryId_example" // string | ID of a registry. Restricts results to those registered in this registry. Specifying \"null\" returns unregistered items.  (optional)
    schemaId := "schemaId_example" // string | ID of a schema. Restricts results to those of the specified schema.  (optional)
    schemaFields := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Schema field value. For Integer, Float, and Date type fields, supports the >= and <= operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived AA sequences. Use \"ANY_ARCHIVED\" to filter for archived AA sequences regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    mentions := "mentions_example" // string | Comma-separated list of item IDs. Restricts results to those that mention the given items in the description.  (optional)
    ids := "prtn_6gxJGfPh,prtn_u7fOvqWg" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    entityRegistryIdsAnyOf := "TP001,TP002" // string | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of users IDs (optional)
    authorIdsAnyOf := "ent_a0SApq3z,ent_b4AApz9b" // string | Comma separated list of user or app IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AASequencesApi.ListAASequences(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).AminoAcids(aminoAcids).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).AuthorIdsAnyOf(authorIdsAnyOf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AASequencesApi.ListAASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAASequences`: AaSequencesPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `AASequencesApi.ListAASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100.  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of an AA Sequence. Restricts results to those with the specified name, alias, or entity registry ID. | 
 **nameIncludes** | **string** | Name substring of an AA sequence. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring. | 
 **aminoAcids** | **string** | String of amino acids. Restricts results to AA sequences exactly matching these amino acids (case-insensitive). | 
 **folderId** | **string** | ID of a folder. Restricts results to those in the folder. | 
 **mentionedIn** | **string** | Comma-separated list of entry IDs. Restricts results to AA sequences mentioned in those entries.  | 
 **projectId** | **string** | ID of a project. Restricts results to those in the project. | 
 **registryId** | **string** | ID of a registry. Restricts results to those registered in this registry. Specifying \&quot;null\&quot; returns unregistered items.  | 
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema.  | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived AA sequences. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived AA sequences regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **mentions** | **string** | Comma-separated list of item IDs. Restricts results to those that mention the given items in the description.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **entityRegistryIdsAnyOf** | **string** | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  | 
 **creatorIds** | **string** | Comma separated list of users IDs | 
 **authorIdsAnyOf** | **string** | Comma separated list of user or app IDs | 

### Return type

[**AaSequencesPaginatedList**](AaSequencesPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveAASequences

> AaSequencesArchivalChange UnarchiveAASequences(ctx).AaSequencesUnarchive(aaSequencesUnarchive).Execute()

Unarchive AA sequences



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
    aaSequencesUnarchive := *openapiclient.NewAaSequencesUnarchive([]string{"AaSequenceIds_example"}) // AaSequencesUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AASequencesApi.UnarchiveAASequences(context.Background()).AaSequencesUnarchive(aaSequencesUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AASequencesApi.UnarchiveAASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveAASequences`: AaSequencesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `AASequencesApi.UnarchiveAASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveAASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aaSequencesUnarchive** | [**AaSequencesUnarchive**](AaSequencesUnarchive.md) |  | 

### Return type

[**AaSequencesArchivalChange**](AaSequencesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAASequence

> AaSequence UpdateAASequence(ctx, aaSequenceId).AaSequenceUpdate(aaSequenceUpdate).Execute()

Update an AA sequence



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
    aaSequenceId := "aaSequenceId_example" // string | 
    aaSequenceUpdate := *openapiclient.NewAaSequenceUpdate() // AaSequenceUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AASequencesApi.UpdateAASequence(context.Background(), aaSequenceId).AaSequenceUpdate(aaSequenceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AASequencesApi.UpdateAASequence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAASequence`: AaSequence
    fmt.Fprintf(os.Stdout, "Response from `AASequencesApi.UpdateAASequence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aaSequenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAASequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aaSequenceUpdate** | [**AaSequenceUpdate**](AaSequenceUpdate.md) |  | 

### Return type

[**AaSequence**](AaSequence.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

