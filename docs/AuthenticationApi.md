# \AuthenticationApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateToken**](AuthenticationApi.md#GenerateToken) | **Post** /token | Generate a token for usage with authenticating via OAuth2 in subsequent API calls.



## GenerateToken

> TokenResponse GenerateToken(ctx).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Execute()

Generate a token for usage with authenticating via OAuth2 in subsequent API calls.



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
    grantType := "grantType_example" // string | 
    clientId := "clientId_example" // string | ID of client to request token for. Leave off if client ID and secret are being supplied through Authorization header.  (optional)
    clientSecret := "clientSecret_example" // string | Leave off if client ID and secret are being supplied through Authorization header.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.GenerateToken(context.Background()).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.GenerateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateToken`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.GenerateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** |  | 
 **clientId** | **string** | ID of client to request token for. Leave off if client ID and secret are being supplied through Authorization header.  | 
 **clientSecret** | **string** | Leave off if client ID and secret are being supplied through Authorization header.  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[basicClientIdSecretAuth](../README.md#basicClientIdSecretAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

