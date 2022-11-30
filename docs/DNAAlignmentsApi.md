# \DNAAlignmentsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDnaConsensusAlignment**](DNAAlignmentsApi.md#CreateDnaConsensusAlignment) | **Post** /dna-alignments:create-consensus-alignment | Create a consensus DNA alignment
[**CreateDnaTemplateAlignment**](DNAAlignmentsApi.md#CreateDnaTemplateAlignment) | **Post** /dna-alignments:create-template-alignment | Create a template DNA alignment
[**DeleteDNAAlignment**](DNAAlignmentsApi.md#DeleteDNAAlignment) | **Delete** /dna-alignments/{dna_alignment_id} | Delete a DNA Alignment
[**GetDNAAlignment**](DNAAlignmentsApi.md#GetDNAAlignment) | **Get** /dna-alignments/{dna_alignment_id} | Get a DNA Alignment
[**ListDNAAlignments**](DNAAlignmentsApi.md#ListDNAAlignments) | **Get** /dna-alignments | List DNA Alignments



## CreateDnaConsensusAlignment

> AsyncTaskLink CreateDnaConsensusAlignment(ctx).DnaConsensusAlignmentCreate(dnaConsensusAlignmentCreate).Execute()

Create a consensus DNA alignment



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
    dnaConsensusAlignmentCreate := *openapiclient.NewDnaConsensusAlignmentCreate("Algorithm_example", []openapiclient.DnaAlignmentBaseFilesInner{openapiclient.DnaAlignmentBase_files_inner{DnaAlignmentBaseFilesInnerOneOf: openapiclient.NewDnaAlignmentBaseFilesInnerOneOf()}}) // DnaConsensusAlignmentCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNAAlignmentsApi.CreateDnaConsensusAlignment(context.Background()).DnaConsensusAlignmentCreate(dnaConsensusAlignmentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNAAlignmentsApi.CreateDnaConsensusAlignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDnaConsensusAlignment`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `DNAAlignmentsApi.CreateDnaConsensusAlignment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDnaConsensusAlignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnaConsensusAlignmentCreate** | [**DnaConsensusAlignmentCreate**](DnaConsensusAlignmentCreate.md) |  | 

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


## CreateDnaTemplateAlignment

> AsyncTaskLink CreateDnaTemplateAlignment(ctx).DnaTemplateAlignmentCreate(dnaTemplateAlignmentCreate).Execute()

Create a template DNA alignment



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
    dnaTemplateAlignmentCreate := *openapiclient.NewDnaTemplateAlignmentCreate("Algorithm_example", []openapiclient.DnaAlignmentBaseFilesInner{openapiclient.DnaAlignmentBase_files_inner{DnaAlignmentBaseFilesInnerOneOf: openapiclient.NewDnaAlignmentBaseFilesInnerOneOf()}}, "seq_rXqq0IHU") // DnaTemplateAlignmentCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNAAlignmentsApi.CreateDnaTemplateAlignment(context.Background()).DnaTemplateAlignmentCreate(dnaTemplateAlignmentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNAAlignmentsApi.CreateDnaTemplateAlignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDnaTemplateAlignment`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `DNAAlignmentsApi.CreateDnaTemplateAlignment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDnaTemplateAlignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnaTemplateAlignmentCreate** | [**DnaTemplateAlignmentCreate**](DnaTemplateAlignmentCreate.md) |  | 

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


## DeleteDNAAlignment

> map[string]interface{} DeleteDNAAlignment(ctx, dnaAlignmentId).Execute()

Delete a DNA Alignment



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
    dnaAlignmentId := "dnaAlignmentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNAAlignmentsApi.DeleteDNAAlignment(context.Background(), dnaAlignmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNAAlignmentsApi.DeleteDNAAlignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDNAAlignment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DNAAlignmentsApi.DeleteDNAAlignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnaAlignmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDNAAlignmentRequest struct via the builder pattern


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


## GetDNAAlignment

> DnaAlignment GetDNAAlignment(ctx, dnaAlignmentId).Execute()

Get a DNA Alignment



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
    dnaAlignmentId := "dnaAlignmentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNAAlignmentsApi.GetDNAAlignment(context.Background(), dnaAlignmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNAAlignmentsApi.GetDNAAlignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDNAAlignment`: DnaAlignment
    fmt.Fprintf(os.Stdout, "Response from `DNAAlignmentsApi.GetDNAAlignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dnaAlignmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDNAAlignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DnaAlignment**](DnaAlignment.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDNAAlignments

> DnaAlignmentsPaginatedList ListDNAAlignments(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).Ids(ids).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).SequenceIds(sequenceIds).Execute()

List DNA Alignments



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
    name := "name_example" // string | Name of a DNA Alignment. Restricts results to those with the specified name. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of a DNA Alignment. Restricts results to those with names that include the provided substring. (optional)
    ids := "seqanl_VfVOART1,seqanl_RFhDGaaC" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  (optional)
    sequenceIds := "seq_VfVOART1,seq_RFhDGaaC" // string | Comma-separated list of sequence ids that own one or more DNA Alignments (i.e. ids of sequences used as the template in a Template Alignment or created as the consensus sequence from a Consensus Alignment). Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNAAlignmentsApi.ListDNAAlignments(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).Ids(ids).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).SequenceIds(sequenceIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNAAlignmentsApi.ListDNAAlignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDNAAlignments`: DnaAlignmentsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DNAAlignmentsApi.ListDNAAlignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDNAAlignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100.  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of a DNA Alignment. Restricts results to those with the specified name. | 
 **nameIncludes** | **string** | Name substring of a DNA Alignment. Restricts results to those with names that include the provided substring. | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  | 
 **sequenceIds** | **string** | Comma-separated list of sequence ids that own one or more DNA Alignments (i.e. ids of sequences used as the template in a Template Alignment or created as the consensus sequence from a Consensus Alignment). Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 

### Return type

[**DnaAlignmentsPaginatedList**](DnaAlignmentsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

