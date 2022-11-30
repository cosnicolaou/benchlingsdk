# \DNASequencesApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveDNASequences**](DNASequencesApi.md#ArchiveDNASequences) | **Post** /dna-sequences:archive | Archive DNA sequences
[**AutoAnnotateDnaSequences**](DNASequencesApi.md#AutoAnnotateDnaSequences) | **Post** /dna-sequences:auto-annotate | Auto-annotate DNA sequences with matching features from specified Feature Libraries
[**AutofillDNASequenceParts**](DNASequencesApi.md#AutofillDNASequenceParts) | **Post** /dna-sequences:autofill-parts | Autofill DNA sequence parts
[**AutofillDNASequenceTranslations**](DNASequencesApi.md#AutofillDNASequenceTranslations) | **Post** /dna-sequences:autofill-translations | Autofill DNA sequence translations
[**BulkCreateDNASequences**](DNASequencesApi.md#BulkCreateDNASequences) | **Post** /dna-sequences:bulk-create | Bulk Create DNA sequences
[**BulkGetDNASequences**](DNASequencesApi.md#BulkGetDNASequences) | **Get** /dna-sequences:bulk-get | Bulk get DNA sequences by ID
[**BulkUpdateDNASequences**](DNASequencesApi.md#BulkUpdateDNASequences) | **Post** /dna-sequences:bulk-update | Bulk Update DNA sequences
[**CreateDNASequence**](DNASequencesApi.md#CreateDNASequence) | **Post** /dna-sequences | Create a DNA sequence
[**GetDNASequence**](DNASequencesApi.md#GetDNASequence) | **Get** /dna-sequences/{dna_sequence_id} | Get a DNA sequence
[**ListDNASequences**](DNASequencesApi.md#ListDNASequences) | **Get** /dna-sequences | List DNA sequences
[**UnarchiveDNASequences**](DNASequencesApi.md#UnarchiveDNASequences) | **Post** /dna-sequences:unarchive | Unarchive DNA sequences
[**UpdateDNASequence**](DNASequencesApi.md#UpdateDNASequence) | **Patch** /dna-sequences/{dna_sequence_id} | Update a DNA sequence



## ArchiveDNASequences

> DnaSequencesArchivalChange ArchiveDNASequences(ctx).DnaSequencesArchive(dnaSequencesArchive).Execute()

Archive DNA sequences



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
    dnaSequencesArchive := *openapiclient.NewDnaSequencesArchive([]string{"DnaSequenceIds_example"}, openapiclient.EntityArchiveReason("Made in error")) // DnaSequencesArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNASequencesApi.ArchiveDNASequences(context.Background()).DnaSequencesArchive(dnaSequencesArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNASequencesApi.ArchiveDNASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveDNASequences`: DnaSequencesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `DNASequencesApi.ArchiveDNASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveDNASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnaSequencesArchive** | [**DnaSequencesArchive**](DnaSequencesArchive.md) |  | 

### Return type

[**DnaSequencesArchivalChange**](DnaSequencesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoAnnotateDnaSequences

> AsyncTaskLink AutoAnnotateDnaSequences(ctx).AutoAnnotateDnaSequences(autoAnnotateDnaSequences).Execute()

Auto-annotate DNA sequences with matching features from specified Feature Libraries



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
    autoAnnotateDnaSequences := *openapiclient.NewAutoAnnotateDnaSequences([]string{"DnaSequenceIds_example"}, []string{"FeatureLibraryIds_example"}) // AutoAnnotateDnaSequences |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNASequencesApi.AutoAnnotateDnaSequences(context.Background()).AutoAnnotateDnaSequences(autoAnnotateDnaSequences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNASequencesApi.AutoAnnotateDnaSequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutoAnnotateDnaSequences`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `DNASequencesApi.AutoAnnotateDnaSequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoAnnotateDnaSequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoAnnotateDnaSequences** | [**AutoAnnotateDnaSequences**](AutoAnnotateDnaSequences.md) |  | 

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


## AutofillDNASequenceParts

> AsyncTaskLink AutofillDNASequenceParts(ctx).AutofillSequences(autofillSequences).Execute()

Autofill DNA sequence parts



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
    autofillSequences := *openapiclient.NewAutofillSequences([]string{"DnaSequenceIds_example"}) // AutofillSequences |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNASequencesApi.AutofillDNASequenceParts(context.Background()).AutofillSequences(autofillSequences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNASequencesApi.AutofillDNASequenceParts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutofillDNASequenceParts`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `DNASequencesApi.AutofillDNASequenceParts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutofillDNASequencePartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autofillSequences** | [**AutofillSequences**](AutofillSequences.md) |  | 

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


## AutofillDNASequenceTranslations

> AsyncTaskLink AutofillDNASequenceTranslations(ctx).AutofillSequences(autofillSequences).Execute()

Autofill DNA sequence translations



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
    autofillSequences := *openapiclient.NewAutofillSequences([]string{"DnaSequenceIds_example"}) // AutofillSequences |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNASequencesApi.AutofillDNASequenceTranslations(context.Background()).AutofillSequences(autofillSequences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNASequencesApi.AutofillDNASequenceTranslations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AutofillDNASequenceTranslations`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `DNASequencesApi.AutofillDNASequenceTranslations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutofillDNASequenceTranslationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autofillSequences** | [**AutofillSequences**](AutofillSequences.md) |  | 

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


## BulkCreateDNASequences

> AsyncTaskLink BulkCreateDNASequences(ctx).DnaSequencesBulkCreateRequest(dnaSequencesBulkCreateRequest).Execute()

Bulk Create DNA sequences



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
    dnaSequencesBulkCreateRequest := *openapiclient.NewDnaSequencesBulkCreateRequest() // DnaSequencesBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNASequencesApi.BulkCreateDNASequences(context.Background()).DnaSequencesBulkCreateRequest(dnaSequencesBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNASequencesApi.BulkCreateDNASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateDNASequences`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `DNASequencesApi.BulkCreateDNASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateDNASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnaSequencesBulkCreateRequest** | [**DnaSequencesBulkCreateRequest**](DnaSequencesBulkCreateRequest.md) |  | 

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


## BulkGetDNASequences

> DnaSequencesBulkGet BulkGetDNASequences(ctx).DnaSequenceIds(dnaSequenceIds).Execute()

Bulk get DNA sequences by ID



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
    dnaSequenceIds := "dnaSequenceIds_example" // string | Comma-separated list of IDs of DNA sequences to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNASequencesApi.BulkGetDNASequences(context.Background()).DnaSequenceIds(dnaSequenceIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNASequencesApi.BulkGetDNASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetDNASequences`: DnaSequencesBulkGet
    fmt.Fprintf(os.Stdout, "Response from `DNASequencesApi.BulkGetDNASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetDNASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnaSequenceIds** | **string** | Comma-separated list of IDs of DNA sequences to get.  | 

### Return type

[**DnaSequencesBulkGet**](DnaSequencesBulkGet.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateDNASequences

> AsyncTaskLink BulkUpdateDNASequences(ctx).DnaSequencesBulkUpdateRequest(dnaSequencesBulkUpdateRequest).Execute()

Bulk Update DNA sequences



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
    dnaSequencesBulkUpdateRequest := *openapiclient.NewDnaSequencesBulkUpdateRequest() // DnaSequencesBulkUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNASequencesApi.BulkUpdateDNASequences(context.Background()).DnaSequencesBulkUpdateRequest(dnaSequencesBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNASequencesApi.BulkUpdateDNASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateDNASequences`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `DNASequencesApi.BulkUpdateDNASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateDNASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnaSequencesBulkUpdateRequest** | [**DnaSequencesBulkUpdateRequest**](DnaSequencesBulkUpdateRequest.md) |  | 

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


## CreateDNASequence

> DnaSequence CreateDNASequence(ctx).DnaSequenceCreate(dnaSequenceCreate).Execute()

Create a DNA sequence



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
    dnaSequenceCreate := *openapiclient.NewDnaSequenceCreate("Bases_example", false, "Name_example") // DnaSequenceCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNASequencesApi.CreateDNASequence(context.Background()).DnaSequenceCreate(dnaSequenceCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNASequencesApi.CreateDNASequence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDNASequence`: DnaSequence
    fmt.Fprintf(os.Stdout, "Response from `DNASequencesApi.CreateDNASequence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDNASequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnaSequenceCreate** | [**DnaSequenceCreate**](DnaSequenceCreate.md) |  | 

### Return type

[**DnaSequence**](DnaSequence.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDNASequence

> DnaSequence GetDNASequence(ctx, dnaSequenceId).Execute()

Get a DNA sequence



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
    dnaSequenceId := "dnaSequenceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNASequencesApi.GetDNASequence(context.Background(), dnaSequenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNASequencesApi.GetDNASequence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDNASequence`: DnaSequence
    fmt.Fprintf(os.Stdout, "Response from `DNASequencesApi.GetDNASequence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnaSequenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDNASequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DnaSequence**](DnaSequence.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDNASequences

> DnaSequencesPaginatedList ListDNASequences(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).Bases(bases).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).AuthorIdsAnyOf(authorIdsAnyOf).Execute()

List DNA sequences



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
    name := "name_example" // string | Name of a DNA Sequence. Restricts results to those with the specified name, alias, or entity registry ID. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of a DNA Sequence. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring. (optional)
    bases := "bases_example" // string | Full bases of the DNA sequence. Restricts results to those with the specified bases, case-insensitive, allowing for circular or reverse complement matches. Does not allow partial matching or loose matching via degenerate bases. (optional)
    folderId := "folderId_example" // string | ID of a folder. Restricts results to those in the folder. (optional)
    mentionedIn := "mentionedIn_example" // string | Comma-separated list of entry IDs. Restricts results to DNA sequences mentioned in those entries.  (optional)
    projectId := "projectId_example" // string | ID of a project. Restricts results to those in the project. (optional)
    registryId := "registryId_example" // string | ID of a registry. Restricts results to those registered in this registry. Specifying \"null\" returns unregistered items.  (optional)
    schemaId := "schemaId_example" // string | ID of a schema. Restricts results to those of the specified schema.  (optional)
    schemaFields := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Schema field value. For Integer, Float, and Date type fields, supports the >= and <= operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived DNA sequences. Use \"ANY_ARCHIVED\" to filter for archived DNA sequences regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    mentions := "mentions_example" // string | Comma-separated list of item IDs. Restricts results to those that mention the given items in the description.  (optional)
    ids := "seq_VfVOART1,seq_RFhDGaaC" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    entityRegistryIdsAnyOf := "TP001,TP002" // string | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of users IDs (optional)
    authorIdsAnyOf := "ent_a0SApq3z,ent_b4AApz9b" // string | Comma separated list of user or app IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNASequencesApi.ListDNASequences(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).Bases(bases).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).AuthorIdsAnyOf(authorIdsAnyOf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNASequencesApi.ListDNASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDNASequences`: DnaSequencesPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DNASequencesApi.ListDNASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDNASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100.  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of a DNA Sequence. Restricts results to those with the specified name, alias, or entity registry ID. | 
 **nameIncludes** | **string** | Name substring of a DNA Sequence. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring. | 
 **bases** | **string** | Full bases of the DNA sequence. Restricts results to those with the specified bases, case-insensitive, allowing for circular or reverse complement matches. Does not allow partial matching or loose matching via degenerate bases. | 
 **folderId** | **string** | ID of a folder. Restricts results to those in the folder. | 
 **mentionedIn** | **string** | Comma-separated list of entry IDs. Restricts results to DNA sequences mentioned in those entries.  | 
 **projectId** | **string** | ID of a project. Restricts results to those in the project. | 
 **registryId** | **string** | ID of a registry. Restricts results to those registered in this registry. Specifying \&quot;null\&quot; returns unregistered items.  | 
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema.  | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived DNA sequences. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived DNA sequences regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **mentions** | **string** | Comma-separated list of item IDs. Restricts results to those that mention the given items in the description.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **entityRegistryIdsAnyOf** | **string** | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  | 
 **creatorIds** | **string** | Comma separated list of users IDs | 
 **authorIdsAnyOf** | **string** | Comma separated list of user or app IDs | 

### Return type

[**DnaSequencesPaginatedList**](DnaSequencesPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveDNASequences

> DnaSequencesArchivalChange UnarchiveDNASequences(ctx).DnaSequencesUnarchive(dnaSequencesUnarchive).Execute()

Unarchive DNA sequences



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
    dnaSequencesUnarchive := *openapiclient.NewDnaSequencesUnarchive([]string{"DnaSequenceIds_example"}) // DnaSequencesUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNASequencesApi.UnarchiveDNASequences(context.Background()).DnaSequencesUnarchive(dnaSequencesUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNASequencesApi.UnarchiveDNASequences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveDNASequences`: DnaSequencesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `DNASequencesApi.UnarchiveDNASequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveDNASequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnaSequencesUnarchive** | [**DnaSequencesUnarchive**](DnaSequencesUnarchive.md) |  | 

### Return type

[**DnaSequencesArchivalChange**](DnaSequencesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDNASequence

> DnaSequence UpdateDNASequence(ctx, dnaSequenceId).DnaSequenceUpdate(dnaSequenceUpdate).Execute()

Update a DNA sequence



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
    dnaSequenceId := "dnaSequenceId_example" // string | 
    dnaSequenceUpdate := *openapiclient.NewDnaSequenceUpdate() // DnaSequenceUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNASequencesApi.UpdateDNASequence(context.Background(), dnaSequenceId).DnaSequenceUpdate(dnaSequenceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNASequencesApi.UpdateDNASequence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDNASequence`: DnaSequence
    fmt.Fprintf(os.Stdout, "Response from `DNASequencesApi.UpdateDNASequence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnaSequenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDNASequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnaSequenceUpdate** | [**DnaSequenceUpdate**](DnaSequenceUpdate.md) |  | 

### Return type

[**DnaSequence**](DnaSequence.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

