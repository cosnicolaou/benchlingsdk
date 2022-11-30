# \LabelTemplatesApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLabelTemplates**](LabelTemplatesApi.md#ListLabelTemplates) | **Get** /registries/{registry_id}/label-templates | List label templates



## ListLabelTemplates

> LabelTemplatesList ListLabelTemplates(ctx, registryId).Name(name).Execute()

List label templates



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
    registryId := "registryId_example" // string | ID of the registry to list the label templates of.
    name := "name_example" // string | Name of a label template. Restricts results to those with the specified name. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelTemplatesApi.ListLabelTemplates(context.Background(), registryId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelTemplatesApi.ListLabelTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLabelTemplates`: LabelTemplatesList
    fmt.Fprintf(os.Stdout, "Response from `LabelTemplatesApi.ListLabelTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | ID of the registry to list the label templates of. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLabelTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of a label template. Restricts results to those with the specified name. | 

### Return type

[**LabelTemplatesList**](LabelTemplatesList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

