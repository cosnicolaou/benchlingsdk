# \NucleotideAlignmentsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConsensusNucleotideAlignment**](NucleotideAlignmentsApi.md#CreateConsensusNucleotideAlignment) | **Post** /nucleotide-alignments:create-consensus-alignment | Create a consensus Nucleotide Alignment
[**CreateTemplateNucleotideAlignment**](NucleotideAlignmentsApi.md#CreateTemplateNucleotideAlignment) | **Post** /nucleotide-alignments:create-template-alignment | Create a template Nucleotide Alignment
[**DeleteNucleotideAlignment**](NucleotideAlignmentsApi.md#DeleteNucleotideAlignment) | **Delete** /nucleotide-alignments/{alignment_id} | Delete a Nucleotide Alignment
[**GetNucleotideAlignment**](NucleotideAlignmentsApi.md#GetNucleotideAlignment) | **Get** /nucleotide-alignments/{alignment_id} | Get a Nucleotide Alignment
[**ListNucleotideAlignments**](NucleotideAlignmentsApi.md#ListNucleotideAlignments) | **Get** /nucleotide-alignments | List Nucleotide Alignments



## CreateConsensusNucleotideAlignment

> AsyncTaskLink CreateConsensusNucleotideAlignment(ctx).NucleotideConsensusAlignmentCreate(nucleotideConsensusAlignmentCreate).Execute()

Create a consensus Nucleotide Alignment



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
    nucleotideConsensusAlignmentCreate := *openapiclient.NewNucleotideConsensusAlignmentCreate("Algorithm_example", []openapiclient.NucleotideAlignmentBaseFilesInner{openapiclient.NucleotideAlignmentBase_files_inner{DnaAlignmentBaseFilesInnerOneOf: openapiclient.NewDnaAlignmentBaseFilesInnerOneOf()}}) // NucleotideConsensusAlignmentCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NucleotideAlignmentsApi.CreateConsensusNucleotideAlignment(context.Background()).NucleotideConsensusAlignmentCreate(nucleotideConsensusAlignmentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NucleotideAlignmentsApi.CreateConsensusNucleotideAlignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConsensusNucleotideAlignment`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `NucleotideAlignmentsApi.CreateConsensusNucleotideAlignment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConsensusNucleotideAlignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nucleotideConsensusAlignmentCreate** | [**NucleotideConsensusAlignmentCreate**](NucleotideConsensusAlignmentCreate.md) |  | 

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


## CreateTemplateNucleotideAlignment

> AsyncTaskLink CreateTemplateNucleotideAlignment(ctx).NucleotideTemplateAlignmentCreate(nucleotideTemplateAlignmentCreate).Execute()

Create a template Nucleotide Alignment



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
    nucleotideTemplateAlignmentCreate := *openapiclient.NewNucleotideTemplateAlignmentCreate("Algorithm_example", []openapiclient.NucleotideAlignmentBaseFilesInner{openapiclient.NucleotideAlignmentBase_files_inner{DnaAlignmentBaseFilesInnerOneOf: openapiclient.NewDnaAlignmentBaseFilesInnerOneOf()}}, "seq_rXqq0IHU") // NucleotideTemplateAlignmentCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NucleotideAlignmentsApi.CreateTemplateNucleotideAlignment(context.Background()).NucleotideTemplateAlignmentCreate(nucleotideTemplateAlignmentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NucleotideAlignmentsApi.CreateTemplateNucleotideAlignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTemplateNucleotideAlignment`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `NucleotideAlignmentsApi.CreateTemplateNucleotideAlignment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateNucleotideAlignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nucleotideTemplateAlignmentCreate** | [**NucleotideTemplateAlignmentCreate**](NucleotideTemplateAlignmentCreate.md) |  | 

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


## DeleteNucleotideAlignment

> map[string]interface{} DeleteNucleotideAlignment(ctx, alignmentId).Execute()

Delete a Nucleotide Alignment



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
    alignmentId := "alignmentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NucleotideAlignmentsApi.DeleteNucleotideAlignment(context.Background(), alignmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NucleotideAlignmentsApi.DeleteNucleotideAlignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNucleotideAlignment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `NucleotideAlignmentsApi.DeleteNucleotideAlignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alignmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNucleotideAlignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNucleotideAlignment

> NucleotideAlignment GetNucleotideAlignment(ctx, alignmentId).Execute()

Get a Nucleotide Alignment



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
    alignmentId := "alignmentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NucleotideAlignmentsApi.GetNucleotideAlignment(context.Background(), alignmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NucleotideAlignmentsApi.GetNucleotideAlignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNucleotideAlignment`: NucleotideAlignment
    fmt.Fprintf(os.Stdout, "Response from `NucleotideAlignmentsApi.GetNucleotideAlignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alignmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNucleotideAlignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NucleotideAlignment**](NucleotideAlignment.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNucleotideAlignments

> NucleotideAlignmentsPaginatedList ListNucleotideAlignments(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).Ids(ids).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).SequenceIds(sequenceIds).Execute()

List Nucleotide Alignments



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
    name := "name_example" // string | Name of a Nucleotide Alignment. Restricts results to those with the specified name. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of a Nucleotide Alignment. Restricts results to those with names that include the provided substring. (optional)
    ids := "seqanl_VfVOART1,seqanl_RFhDGaaC" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  (optional)
    sequenceIds := "seq_VfVOART1,seq_RFhDGaaC" // string | Comma-separated list of sequence ids that own one or more Nucleotide Alignments (i.e. ids of sequences used as the template in a Template Alignment or created as the consensus sequence from a Consensus Alignment). Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NucleotideAlignmentsApi.ListNucleotideAlignments(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).Ids(ids).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).SequenceIds(sequenceIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NucleotideAlignmentsApi.ListNucleotideAlignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNucleotideAlignments`: NucleotideAlignmentsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `NucleotideAlignmentsApi.ListNucleotideAlignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNucleotideAlignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100.  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of a Nucleotide Alignment. Restricts results to those with the specified name. | 
 **nameIncludes** | **string** | Name substring of a Nucleotide Alignment. Restricts results to those with names that include the provided substring. | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  | 
 **sequenceIds** | **string** | Comma-separated list of sequence ids that own one or more Nucleotide Alignments (i.e. ids of sequences used as the template in a Template Alignment or created as the consensus sequence from a Consensus Alignment). Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 

### Return type

[**NucleotideAlignmentsPaginatedList**](NucleotideAlignmentsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

