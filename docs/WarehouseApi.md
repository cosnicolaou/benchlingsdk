# \WarehouseApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWarehouseCredentials**](WarehouseApi.md#CreateWarehouseCredentials) | **Post** /warehouse-credentials | Create Benchling Warehouse credentials



## CreateWarehouseCredentials

> WarehouseCredentials CreateWarehouseCredentials(ctx).WarehouseCredentialsCreate(warehouseCredentialsCreate).Execute()

Create Benchling Warehouse credentials



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
    warehouseCredentialsCreate := *openapiclient.NewWarehouseCredentialsCreate(int32(123)) // WarehouseCredentialsCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WarehouseApi.CreateWarehouseCredentials(context.Background()).WarehouseCredentialsCreate(warehouseCredentialsCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarehouseApi.CreateWarehouseCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWarehouseCredentials`: WarehouseCredentials
    fmt.Fprintf(os.Stdout, "Response from `WarehouseApi.CreateWarehouseCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWarehouseCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **warehouseCredentialsCreate** | [**WarehouseCredentialsCreate**](WarehouseCredentialsCreate.md) |  | 

### Return type

[**WarehouseCredentials**](WarehouseCredentials.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

