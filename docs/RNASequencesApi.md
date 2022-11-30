# \RNASequencesApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveRNASequences**](RNASequencesApi.md#ArchiveRNASequences) | **Post** /rna-sequences:archive | Archive RNA Sequences
[**AutoAnnotateRnaSequences**](RNASequencesApi.md#AutoAnnotateRnaSequences) | **Post** /rna-sequences:auto-annotate | Auto-annotate RNA sequences with matching features from specified Feature Libraries
[**AutofillRNASequenceParts**](RNASequencesApi.md#AutofillRNASequenceParts) | **Post** /rna-sequences:autofill-parts | Autofill RNA sequence parts
[**AutofillRNASequenceTranslations**](RNASequencesApi.md#AutofillRNASequenceTranslations) | **Post** /rna-sequences:autofill-translations | Autofill RNA sequence translations from Amino Acid sequences with matching schemas
[**BulkCreateRNASequences**](RNASequencesApi.md#BulkCreateRNASequences) | **Post** /rna-sequences:bulk-create | Bulk Create RNA sequences
[**BulkGetRNASequences**](RNASequencesApi.md#BulkGetRNASequences) | **Get** /rna-sequences:bulk-get | Bulk get RNA sequences by ID
[**BulkUpdateRNASequences**](RNASequencesApi.md#BulkUpdateRNASequences) | **Post** /rna-sequences:bulk-update | Bulk Update RNA sequences
[**CreateRNASequence**](RNASequencesApi.md#CreateRNASequence) | **Post** /rna-sequences | Create an RNA sequence
[**GetRNASequence**](RNASequencesApi.md#GetRNASequence) | **Get** /rna-sequences/{rna_sequence_id} | Get an RNA sequence
[**ListRNASequences**](RNASequencesApi.md#ListRNASequences) | **Get** /rna-sequences | List RNA sequences
[**UnarchiveRNASequences**](RNASequencesApi.md#UnarchiveRNASequences) | **Post** /rna-sequences:unarchive | Unarchive RNA sequences
[**UpdateRNASequence**](RNASequencesApi.md#UpdateRNASequence) | **Patch** /rna-sequences/{rna_sequence_id} | Update an RNA sequence



## ArchiveRNASequences

> RnaSequencesArchivalChange ArchiveRNASequences(ctx).RnaSequencesArchive(rnaSequencesArchive).Execute()

Archive RNA Sequences



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
    rnaSequencesArchive := *openapiclient.NewRnaSequencesArchive(openapiclient.EntityArchiveReason("Made in error"), []string{"RnaSequenceIds_example"}) // RnaSequencesArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNASequencesApi.ArchiveRNASequences(context.Background()).RnaSequencesArchive(rnaSequencesArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNASequencesApi.ArchiveRNASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveRNASequences`: RnaSequencesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `RNASequencesApi.ArchiveRNASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveRNASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rnaSequencesArchive** | [**RnaSequencesArchive**](RnaSequencesArchive.md) |  | 

### Return type

[**RnaSequencesArchivalChange**](RnaSequencesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoAnnotateRnaSequences

> AsyncTaskLink AutoAnnotateRnaSequences(ctx).AutoAnnotateRnaSequences(autoAnnotateRnaSequences).Execute()

Auto-annotate RNA sequences with matching features from specified Feature Libraries



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
    autoAnnotateRnaSequences := *openapiclient.NewAutoAnnotateRnaSequences([]string{"FeatureLibraryIds_example"}, []string{"RnaSequenceIds_example"}) // AutoAnnotateRnaSequences |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNASequencesApi.AutoAnnotateRnaSequences(context.Background()).AutoAnnotateRnaSequences(autoAnnotateRnaSequences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNASequencesApi.AutoAnnotateRnaSequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutoAnnotateRnaSequences`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `RNASequencesApi.AutoAnnotateRnaSequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoAnnotateRnaSequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoAnnotateRnaSequences** | [**AutoAnnotateRnaSequences**](AutoAnnotateRnaSequences.md) |  | 

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


## AutofillRNASequenceParts

> AsyncTaskLink AutofillRNASequenceParts(ctx).AutofillRnaSequences(autofillRnaSequences).Execute()

Autofill RNA sequence parts



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
    autofillRnaSequences := *openapiclient.NewAutofillRnaSequences([]string{"RnaSequenceIds_example"}) // AutofillRnaSequences |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNASequencesApi.AutofillRNASequenceParts(context.Background()).AutofillRnaSequences(autofillRnaSequences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNASequencesApi.AutofillRNASequenceParts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutofillRNASequenceParts`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `RNASequencesApi.AutofillRNASequenceParts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutofillRNASequencePartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autofillRnaSequences** | [**AutofillRnaSequences**](AutofillRnaSequences.md) |  | 

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


## AutofillRNASequenceTranslations

> AsyncTaskLink AutofillRNASequenceTranslations(ctx).AutofillRnaSequences(autofillRnaSequences).Execute()

Autofill RNA sequence translations from Amino Acid sequences with matching schemas



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
    autofillRnaSequences := *openapiclient.NewAutofillRnaSequences([]string{"RnaSequenceIds_example"}) // AutofillRnaSequences |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNASequencesApi.AutofillRNASequenceTranslations(context.Background()).AutofillRnaSequences(autofillRnaSequences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNASequencesApi.AutofillRNASequenceTranslations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutofillRNASequenceTranslations`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `RNASequencesApi.AutofillRNASequenceTranslations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutofillRNASequenceTranslationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autofillRnaSequences** | [**AutofillRnaSequences**](AutofillRnaSequences.md) |  | 

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


## BulkCreateRNASequences

> AsyncTaskLink BulkCreateRNASequences(ctx).RnaSequencesBulkCreateRequest(rnaSequencesBulkCreateRequest).Execute()

Bulk Create RNA sequences



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
    rnaSequencesBulkCreateRequest := *openapiclient.NewRnaSequencesBulkCreateRequest() // RnaSequencesBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNASequencesApi.BulkCreateRNASequences(context.Background()).RnaSequencesBulkCreateRequest(rnaSequencesBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNASequencesApi.BulkCreateRNASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateRNASequences`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `RNASequencesApi.BulkCreateRNASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateRNASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rnaSequencesBulkCreateRequest** | [**RnaSequencesBulkCreateRequest**](RnaSequencesBulkCreateRequest.md) |  | 

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


## BulkGetRNASequences

> RnaSequencesBulkGet BulkGetRNASequences(ctx).RnaSequenceIds(rnaSequenceIds).Execute()

Bulk get RNA sequences by ID



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
    rnaSequenceIds := "rnaSequenceIds_example" // string | Comma-separated list of IDs of RNA sequences to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNASequencesApi.BulkGetRNASequences(context.Background()).RnaSequenceIds(rnaSequenceIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNASequencesApi.BulkGetRNASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetRNASequences`: RnaSequencesBulkGet
    fmt.Fprintf(os.Stdout, "Response from `RNASequencesApi.BulkGetRNASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetRNASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rnaSequenceIds** | **string** | Comma-separated list of IDs of RNA sequences to get.  | 

### Return type

[**RnaSequencesBulkGet**](RnaSequencesBulkGet.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateRNASequences

> AsyncTaskLink BulkUpdateRNASequences(ctx).RnaSequencesBulkUpdateRequest(rnaSequencesBulkUpdateRequest).Execute()

Bulk Update RNA sequences



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
    rnaSequencesBulkUpdateRequest := *openapiclient.NewRnaSequencesBulkUpdateRequest() // RnaSequencesBulkUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNASequencesApi.BulkUpdateRNASequences(context.Background()).RnaSequencesBulkUpdateRequest(rnaSequencesBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNASequencesApi.BulkUpdateRNASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateRNASequences`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `RNASequencesApi.BulkUpdateRNASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateRNASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rnaSequencesBulkUpdateRequest** | [**RnaSequencesBulkUpdateRequest**](RnaSequencesBulkUpdateRequest.md) |  | 

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


## CreateRNASequence

> RnaSequence CreateRNASequence(ctx).RnaSequenceCreate(rnaSequenceCreate).Execute()

Create an RNA sequence



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
    rnaSequenceCreate := *openapiclient.NewRnaSequenceCreate("Bases_example", "Name_example") // RnaSequenceCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNASequencesApi.CreateRNASequence(context.Background()).RnaSequenceCreate(rnaSequenceCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNASequencesApi.CreateRNASequence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRNASequence`: RnaSequence
    fmt.Fprintf(os.Stdout, "Response from `RNASequencesApi.CreateRNASequence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRNASequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rnaSequenceCreate** | [**RnaSequenceCreate**](RnaSequenceCreate.md) |  | 

### Return type

[**RnaSequence**](RnaSequence.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRNASequence

> RnaSequence GetRNASequence(ctx, rnaSequenceId).Execute()

Get an RNA sequence



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
    rnaSequenceId := "rnaSequenceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNASequencesApi.GetRNASequence(context.Background(), rnaSequenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNASequencesApi.GetRNASequence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRNASequence`: RnaSequence
    fmt.Fprintf(os.Stdout, "Response from `RNASequencesApi.GetRNASequence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rnaSequenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRNASequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RnaSequence**](RnaSequence.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRNASequences

> RnaSequencesPaginatedList ListRNASequences(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).Bases(bases).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).AuthorIdsAnyOf(authorIdsAnyOf).Execute()

List RNA sequences



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
    name := "name_example" // string | Name of an RNA Sequence. Restricts results to those with the specified name, alias, or entity registry ID. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of an RNA Sequence. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring. (optional)
    bases := "bases_example" // string | Full bases of the RNA sequence. Restricts results to those with the specified bases, case-insensitive, allowing for circular or reverse complement matches. Does not allow partial matching or loose matching via degenerate bases. (optional)
    folderId := "folderId_example" // string | ID of a folder. Restricts results to those in the folder. (optional)
    mentionedIn := "mentionedIn_example" // string | Comma-separated list of entry IDs. Restricts results to RNA sequences mentioned in those entries.  (optional)
    projectId := "projectId_example" // string | ID of a project. Restricts results to those in the project. (optional)
    registryId := "registryId_example" // string | ID of a registry. Restricts results to those registered in this registry. Specifying \"null\" returns unregistered items.  (optional)
    schemaId := "schemaId_example" // string | ID of a schema. Restricts results to those of the specified schema.  (optional)
    schemaFields := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Schema field value. For Integer, Float, and Date type fields, supports the >= and <= operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived RNA sequences. Use \"ANY_ARCHIVED\" to filter for archived RNA sequences regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    mentions := "mentions_example" // string | Comma-separated list of item IDs. Restricts results to those that mention the given items in the description.  (optional)
    ids := "seq_VzVOART1,seq_RahDGaaC" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    entityRegistryIdsAnyOf := "TP001,TP002" // string | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of users IDs (optional)
    authorIdsAnyOf := "ent_a0SApq3z,ent_b4AApz9b" // string | Comma separated list of user or app IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNASequencesApi.ListRNASequences(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).Bases(bases).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).AuthorIdsAnyOf(authorIdsAnyOf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNASequencesApi.ListRNASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRNASequences`: RnaSequencesPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `RNASequencesApi.ListRNASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRNASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100.  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of an RNA Sequence. Restricts results to those with the specified name, alias, or entity registry ID. | 
 **nameIncludes** | **string** | Name substring of an RNA Sequence. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring. | 
 **bases** | **string** | Full bases of the RNA sequence. Restricts results to those with the specified bases, case-insensitive, allowing for circular or reverse complement matches. Does not allow partial matching or loose matching via degenerate bases. | 
 **folderId** | **string** | ID of a folder. Restricts results to those in the folder. | 
 **mentionedIn** | **string** | Comma-separated list of entry IDs. Restricts results to RNA sequences mentioned in those entries.  | 
 **projectId** | **string** | ID of a project. Restricts results to those in the project. | 
 **registryId** | **string** | ID of a registry. Restricts results to those registered in this registry. Specifying \&quot;null\&quot; returns unregistered items.  | 
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema.  | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived RNA sequences. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived RNA sequences regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **mentions** | **string** | Comma-separated list of item IDs. Restricts results to those that mention the given items in the description.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **entityRegistryIdsAnyOf** | **string** | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  | 
 **creatorIds** | **string** | Comma separated list of users IDs | 
 **authorIdsAnyOf** | **string** | Comma separated list of user or app IDs | 

### Return type

[**RnaSequencesPaginatedList**](RnaSequencesPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveRNASequences

> RnaSequencesArchivalChange UnarchiveRNASequences(ctx).RnaSequencesUnarchive(rnaSequencesUnarchive).Execute()

Unarchive RNA sequences



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
    rnaSequencesUnarchive := *openapiclient.NewRnaSequencesUnarchive([]string{"RnaSequenceIds_example"}) // RnaSequencesUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNASequencesApi.UnarchiveRNASequences(context.Background()).RnaSequencesUnarchive(rnaSequencesUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNASequencesApi.UnarchiveRNASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveRNASequences`: RnaSequencesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `RNASequencesApi.UnarchiveRNASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveRNASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rnaSequencesUnarchive** | [**RnaSequencesUnarchive**](RnaSequencesUnarchive.md) |  | 

### Return type

[**RnaSequencesArchivalChange**](RnaSequencesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRNASequence

> RnaSequence UpdateRNASequence(ctx, rnaSequenceId).RnaSequenceUpdate(rnaSequenceUpdate).Execute()

Update an RNA sequence



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
    rnaSequenceId := "rnaSequenceId_example" // string | 
    rnaSequenceUpdate := *openapiclient.NewRnaSequenceUpdate() // RnaSequenceUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RNASequencesApi.UpdateRNASequence(context.Background(), rnaSequenceId).RnaSequenceUpdate(rnaSequenceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RNASequencesApi.UpdateRNASequence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRNASequence`: RnaSequence
    fmt.Fprintf(os.Stdout, "Response from `RNASequencesApi.UpdateRNASequence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rnaSequenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRNASequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rnaSequenceUpdate** | [**RnaSequenceUpdate**](RnaSequenceUpdate.md) |  | 

### Return type

[**RnaSequence**](RnaSequence.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

