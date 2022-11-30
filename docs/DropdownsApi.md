# \DropdownsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveDropdownOptions**](DropdownsApi.md#ArchiveDropdownOptions) | **Post** /dropdowns/{dropdown_id}/options:archive | Archive dropdown options
[**CreateDropdown**](DropdownsApi.md#CreateDropdown) | **Post** /dropdowns | Create a dropdown
[**GetDropdown**](DropdownsApi.md#GetDropdown) | **Get** /dropdowns/{dropdown_id} | Get a dropdown
[**ListDropdowns**](DropdownsApi.md#ListDropdowns) | **Get** /dropdowns | List dropdowns
[**UnarchiveDropdownOptions**](DropdownsApi.md#UnarchiveDropdownOptions) | **Post** /dropdowns/{dropdown_id}/options:unarchive | Unarchive dropdown options
[**UpdateDropdown**](DropdownsApi.md#UpdateDropdown) | **Patch** /dropdowns/{dropdown_id} | Update a dropdown



## ArchiveDropdownOptions

> DropdownOptionsArchivalChange ArchiveDropdownOptions(ctx, dropdownId).DropdownOptionsArchive(dropdownOptionsArchive).Execute()

Archive dropdown options



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
    dropdownId := "dropdownId_example" // string | ID of the dropdown to archive options for.
    dropdownOptionsArchive := *openapiclient.NewDropdownOptionsArchive() // DropdownOptionsArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DropdownsApi.ArchiveDropdownOptions(context.Background(), dropdownId).DropdownOptionsArchive(dropdownOptionsArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DropdownsApi.ArchiveDropdownOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveDropdownOptions`: DropdownOptionsArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `DropdownsApi.ArchiveDropdownOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dropdownId** | **string** | ID of the dropdown to archive options for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveDropdownOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dropdownOptionsArchive** | [**DropdownOptionsArchive**](DropdownOptionsArchive.md) |  | 

### Return type

[**DropdownOptionsArchivalChange**](DropdownOptionsArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDropdown

> Dropdown CreateDropdown(ctx).DropdownCreate(dropdownCreate).Execute()

Create a dropdown



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
    dropdownCreate := *openapiclient.NewDropdownCreate("Name_example", []openapiclient.DropdownOptionCreate{*openapiclient.NewDropdownOptionCreate("Name_example")}) // DropdownCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DropdownsApi.CreateDropdown(context.Background()).DropdownCreate(dropdownCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DropdownsApi.CreateDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDropdown`: Dropdown
    fmt.Fprintf(os.Stdout, "Response from `DropdownsApi.CreateDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dropdownCreate** | [**DropdownCreate**](DropdownCreate.md) |  | 

### Return type

[**Dropdown**](Dropdown.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDropdown

> Dropdown GetDropdown(ctx, dropdownId).Execute()

Get a dropdown



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
    dropdownId := "dropdownId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DropdownsApi.GetDropdown(context.Background(), dropdownId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DropdownsApi.GetDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDropdown`: Dropdown
    fmt.Fprintf(os.Stdout, "Response from `DropdownsApi.GetDropdown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dropdownId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dropdown**](Dropdown.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDropdowns

> DropdownSummariesPaginatedList ListDropdowns(ctx).NextToken(nextToken).PageSize(pageSize).Execute()

List dropdowns



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
    nextToken := "nextToken_example" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DropdownsApi.ListDropdowns(context.Background()).NextToken(nextToken).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DropdownsApi.ListDropdowns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDropdowns`: DropdownSummariesPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `DropdownsApi.ListDropdowns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDropdownsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]

### Return type

[**DropdownSummariesPaginatedList**](DropdownSummariesPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveDropdownOptions

> DropdownOptionsArchivalChange UnarchiveDropdownOptions(ctx, dropdownId).DropdownOptionsUnarchive(dropdownOptionsUnarchive).Execute()

Unarchive dropdown options



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
    dropdownId := "dropdownId_example" // string | ID of the dropdown to archive options for.
    dropdownOptionsUnarchive := *openapiclient.NewDropdownOptionsUnarchive() // DropdownOptionsUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DropdownsApi.UnarchiveDropdownOptions(context.Background(), dropdownId).DropdownOptionsUnarchive(dropdownOptionsUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DropdownsApi.UnarchiveDropdownOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveDropdownOptions`: DropdownOptionsArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `DropdownsApi.UnarchiveDropdownOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dropdownId** | **string** | ID of the dropdown to archive options for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveDropdownOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dropdownOptionsUnarchive** | [**DropdownOptionsUnarchive**](DropdownOptionsUnarchive.md) |  | 

### Return type

[**DropdownOptionsArchivalChange**](DropdownOptionsArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDropdown

> Dropdown UpdateDropdown(ctx, dropdownId).DropdownUpdate(dropdownUpdate).Execute()

Update a dropdown



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
    dropdownId := "dropdownId_example" // string | 
    dropdownUpdate := *openapiclient.NewDropdownUpdate([]openapiclient.DropdownOptionUpdate{*openapiclient.NewDropdownOptionUpdate("Name_example")}) // DropdownUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DropdownsApi.UpdateDropdown(context.Background(), dropdownId).DropdownUpdate(dropdownUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DropdownsApi.UpdateDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDropdown`: Dropdown
    fmt.Fprintf(os.Stdout, "Response from `DropdownsApi.UpdateDropdown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dropdownId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dropdownUpdate** | [**DropdownUpdate**](DropdownUpdate.md) |  | 

### Return type

[**Dropdown**](Dropdown.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

