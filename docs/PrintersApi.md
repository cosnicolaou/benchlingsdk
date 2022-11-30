# \PrintersApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPrinters**](PrintersApi.md#ListPrinters) | **Get** /registries/{registry_id}/label-printers | List printers



## ListPrinters

> PrintersList ListPrinters(ctx, registryId).Name(name).Execute()

List printers



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
    registryId := "registryId_example" // string | ID of the registry to list the printers for.
    name := "name_example" // string | Name of a printer. Restricts results to those with the specified name. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrintersApi.ListPrinters(context.Background(), registryId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrintersApi.ListPrinters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrinters`: PrintersList
    fmt.Fprintf(os.Stdout, "Response from `PrintersApi.ListPrinters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | ID of the registry to list the printers for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPrintersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of a printer. Restricts results to those with the specified name. | 

### Return type

[**PrintersList**](PrintersList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

