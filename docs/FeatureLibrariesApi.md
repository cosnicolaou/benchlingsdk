# \FeatureLibrariesApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateFeatures**](FeatureLibrariesApi.md#BulkCreateFeatures) | **Post** /features:bulk-create | Bulk create Features
[**CreateFeature**](FeatureLibrariesApi.md#CreateFeature) | **Post** /features | Create a Feature
[**CreateFeatureLibrary**](FeatureLibrariesApi.md#CreateFeatureLibrary) | **Post** /feature-libraries | Create a Feature Library
[**GetFeature**](FeatureLibrariesApi.md#GetFeature) | **Get** /features/{feature_id} | Get a feature by ID
[**GetFeatureLibrary**](FeatureLibrariesApi.md#GetFeatureLibrary) | **Get** /feature-libraries/{feature_library_id} | Get a feature library by ID
[**ListFeatureLibraries**](FeatureLibrariesApi.md#ListFeatureLibraries) | **Get** /feature-libraries | List Feature Libraries
[**ListFeatures**](FeatureLibrariesApi.md#ListFeatures) | **Get** /features | List Features
[**UpdateFeature**](FeatureLibrariesApi.md#UpdateFeature) | **Patch** /features/{feature_id} | Update a feature
[**UpdateFeatureLibrary**](FeatureLibrariesApi.md#UpdateFeatureLibrary) | **Patch** /feature-libraries/{feature_library_id} | Update a feature library



## BulkCreateFeatures

> AsyncTaskLink BulkCreateFeatures(ctx).FeaturesBulkCreateRequest(featuresBulkCreateRequest).Execute()

Bulk create Features



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
    featuresBulkCreateRequest := *openapiclient.NewFeaturesBulkCreateRequest() // FeaturesBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureLibrariesApi.BulkCreateFeatures(context.Background()).FeaturesBulkCreateRequest(featuresBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureLibrariesApi.BulkCreateFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateFeatures`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `FeatureLibrariesApi.BulkCreateFeatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featuresBulkCreateRequest** | [**FeaturesBulkCreateRequest**](FeaturesBulkCreateRequest.md) |  | 

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


## CreateFeature

> Feature CreateFeature(ctx).FeatureCreate(featureCreate).Execute()

Create a Feature



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
    featureCreate := *openapiclient.NewFeatureCreate("featlib_19kd9aDq", "FeatureType_example", "Name_example", "Pattern_example", "MatchType_example") // FeatureCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureLibrariesApi.CreateFeature(context.Background()).FeatureCreate(featureCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureLibrariesApi.CreateFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFeature`: Feature
    fmt.Fprintf(os.Stdout, "Response from `FeatureLibrariesApi.CreateFeature`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featureCreate** | [**FeatureCreate**](FeatureCreate.md) |  | 

### Return type

[**Feature**](Feature.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFeatureLibrary

> FeatureLibrary CreateFeatureLibrary(ctx).FeatureLibraryCreate(featureLibraryCreate).Execute()

Create a Feature Library



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
    featureLibraryCreate := *openapiclient.NewFeatureLibraryCreate("Description_example", "Name_example") // FeatureLibraryCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureLibrariesApi.CreateFeatureLibrary(context.Background()).FeatureLibraryCreate(featureLibraryCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureLibrariesApi.CreateFeatureLibrary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFeatureLibrary`: FeatureLibrary
    fmt.Fprintf(os.Stdout, "Response from `FeatureLibrariesApi.CreateFeatureLibrary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeatureLibraryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featureLibraryCreate** | [**FeatureLibraryCreate**](FeatureLibraryCreate.md) |  | 

### Return type

[**FeatureLibrary**](FeatureLibrary.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeature

> Feature GetFeature(ctx, featureId).Execute()

Get a feature by ID



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
    featureId := "featureId_example" // string | ID of feature to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureLibrariesApi.GetFeature(context.Background(), featureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureLibrariesApi.GetFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeature`: Feature
    fmt.Fprintf(os.Stdout, "Response from `FeatureLibrariesApi.GetFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | ID of feature to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Feature**](Feature.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureLibrary

> FeatureLibrary GetFeatureLibrary(ctx, featureLibraryId).Execute()

Get a feature library by ID



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
    featureLibraryId := "featureLibraryId_example" // string | ID of feature library to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureLibrariesApi.GetFeatureLibrary(context.Background(), featureLibraryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureLibrariesApi.GetFeatureLibrary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureLibrary`: FeatureLibrary
    fmt.Fprintf(os.Stdout, "Response from `FeatureLibrariesApi.GetFeatureLibrary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureLibraryId** | **string** | ID of feature library to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureLibraryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FeatureLibrary**](FeatureLibrary.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeatureLibraries

> FeatureLibrariesPaginatedList ListFeatureLibraries(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).Ids(ids).NamesAnyOf(namesAnyOf).Execute()

List Feature Libraries



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
    name := "name_example" // string | Name of a Feature Library. Restricts results to those with the specified name. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of a Feature Library. Restricts results to those with names that include the provided substring. (optional)
    ids := "featlib_VfVOART1,featlib_RFhDGaaC" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureLibrariesApi.ListFeatureLibraries(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).Ids(ids).NamesAnyOf(namesAnyOf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureLibrariesApi.ListFeatureLibraries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFeatureLibraries`: FeatureLibrariesPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `FeatureLibrariesApi.ListFeatureLibraries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFeatureLibrariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100.  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of a Feature Library. Restricts results to those with the specified name. | 
 **nameIncludes** | **string** | Name substring of a Feature Library. Restricts results to those with names that include the provided substring. | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names.  | 

### Return type

[**FeatureLibrariesPaginatedList**](FeatureLibrariesPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeatures

> FeaturesPaginatedList ListFeatures(ctx).PageSize(pageSize).NextToken(nextToken).Name(name).Ids(ids).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).FeatureLibraryId(featureLibraryId).FeatureType(featureType).MatchType(matchType).Execute()

List Features



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
    name := "name_example" // string | Name of a Feature. Restricts results to those with the specified name. (optional)
    ids := "feat_VfVOART1,feat_RFhDGaaC" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names.  (optional)
    featureLibraryId := "featlib_D0v2x9Q7" // string | Id of a Feature Library. Restricts results to those associated with the provided feature library (optional)
    featureType := "terminator" // string | The type of feature, like gene, promoter, etc. Note: This is an arbitrary string, not an enum  (optional)
    matchType := "nucleotide" // string | The match type of the feature used to determine how auto-annotate matches are made. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureLibrariesApi.ListFeatures(context.Background()).PageSize(pageSize).NextToken(nextToken).Name(name).Ids(ids).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).FeatureLibraryId(featureLibraryId).FeatureType(featureType).MatchType(matchType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureLibrariesApi.ListFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFeatures`: FeaturesPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `FeatureLibrariesApi.ListFeatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100.  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **name** | **string** | Name of a Feature. Restricts results to those with the specified name. | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names.  | 
 **featureLibraryId** | **string** | Id of a Feature Library. Restricts results to those associated with the provided feature library | 
 **featureType** | **string** | The type of feature, like gene, promoter, etc. Note: This is an arbitrary string, not an enum  | 
 **matchType** | **string** | The match type of the feature used to determine how auto-annotate matches are made. | 

### Return type

[**FeaturesPaginatedList**](FeaturesPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeature

> Feature UpdateFeature(ctx, featureId).FeatureUpdate(featureUpdate).Execute()

Update a feature



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
    featureId := "featureId_example" // string | 
    featureUpdate := *openapiclient.NewFeatureUpdate() // FeatureUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureLibrariesApi.UpdateFeature(context.Background(), featureId).FeatureUpdate(featureUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureLibrariesApi.UpdateFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeature`: Feature
    fmt.Fprintf(os.Stdout, "Response from `FeatureLibrariesApi.UpdateFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **featureUpdate** | [**FeatureUpdate**](FeatureUpdate.md) |  | 

### Return type

[**Feature**](Feature.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeatureLibrary

> FeatureLibrary UpdateFeatureLibrary(ctx, featureLibraryId).FeatureLibraryUpdate(featureLibraryUpdate).Execute()

Update a feature library



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
    featureLibraryId := "featureLibraryId_example" // string | 
    featureLibraryUpdate := *openapiclient.NewFeatureLibraryUpdate() // FeatureLibraryUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureLibrariesApi.UpdateFeatureLibrary(context.Background(), featureLibraryId).FeatureLibraryUpdate(featureLibraryUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureLibrariesApi.UpdateFeatureLibrary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeatureLibrary`: FeatureLibrary
    fmt.Fprintf(os.Stdout, "Response from `FeatureLibrariesApi.UpdateFeatureLibrary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureLibraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureLibraryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **featureLibraryUpdate** | [**FeatureLibraryUpdate**](FeatureLibraryUpdate.md) |  | 

### Return type

[**FeatureLibrary**](FeatureLibrary.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

