# \InventoryApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateBarcodes**](InventoryApi.md#ValidateBarcodes) | **Post** /registries/{registry_id}:validate-barcodes | Validate barcodes



## ValidateBarcodes

> BarcodeValidationResults ValidateBarcodes(ctx, registryId).BarcodesList(barcodesList).Execute()

Validate barcodes



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
    registryId := "registryId_example" // string | ID of the registry to validate barcodes in.
    barcodesList := *openapiclient.NewBarcodesList([]string{"Barcodes_example"}) // BarcodesList |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryApi.ValidateBarcodes(context.Background(), registryId).BarcodesList(barcodesList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ValidateBarcodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateBarcodes`: BarcodeValidationResults
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ValidateBarcodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | ID of the registry to validate barcodes in. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateBarcodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **barcodesList** | [**BarcodesList**](BarcodesList.md) |  | 

### Return type

[**BarcodeValidationResults**](BarcodeValidationResults.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

