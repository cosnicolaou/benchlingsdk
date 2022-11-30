# \RegistryApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkGetRegisteredEntities**](RegistryApi.md#BulkGetRegisteredEntities) | **Get** /registries/{registry_id}/registered-entities:bulk-get | Bulk get registered entities
[**GetRegistry**](RegistryApi.md#GetRegistry) | **Get** /registries/{registry_id} | Get registry
[**ListBatchSchemasByRegistry**](RegistryApi.md#ListBatchSchemasByRegistry) | **Get** /registries/{registry_id}/batch-schemas | List batch schemas by registry
[**ListBoxSchemasByRegistry**](RegistryApi.md#ListBoxSchemasByRegistry) | **Get** /registries/{registry_id}/box-schemas | List box schemas by registry
[**ListContainerSchemasByRegistry**](RegistryApi.md#ListContainerSchemasByRegistry) | **Get** /registries/{registry_id}/container-schemas | List container schemas by registry
[**ListDropdownsByRegistry**](RegistryApi.md#ListDropdownsByRegistry) | **Get** /registries/{registry_id}/dropdowns | List dropdowns for a given registry
[**ListEntitySchemasByRegistry**](RegistryApi.md#ListEntitySchemasByRegistry) | **Get** /registries/{registry_id}/entity-schemas | List entity schemas by registry
[**ListLocationSchemasByRegistry**](RegistryApi.md#ListLocationSchemasByRegistry) | **Get** /registries/{registry_id}/location-schemas | List location schemas by registry
[**ListPlateSchemasByRegistry**](RegistryApi.md#ListPlateSchemasByRegistry) | **Get** /registries/{registry_id}/plate-schemas | List plate schemas by registry
[**ListRegistries**](RegistryApi.md#ListRegistries) | **Get** /registries | List registries
[**RegisterEntities**](RegistryApi.md#RegisterEntities) | **Post** /registries/{registry_id}:bulk-register-entities | Register entities
[**UnregisterEntities**](RegistryApi.md#UnregisterEntities) | **Post** /registries/{registry_id}:unregister-entities | Unregister entities



## BulkGetRegisteredEntities

> RegisteredEntitiesList BulkGetRegisteredEntities(ctx, registryId).EntityRegistryIds(entityRegistryIds).Execute()

Bulk get registered entities



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
    registryId := "registryId_example" // string | ID for the registry
    entityRegistryIds := "entityRegistryIds_example" // string | Comma separated list of entity Registry IDs

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.BulkGetRegisteredEntities(context.Background(), registryId).EntityRegistryIds(entityRegistryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.BulkGetRegisteredEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetRegisteredEntities`: RegisteredEntitiesList
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.BulkGetRegisteredEntities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | ID for the registry | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetRegisteredEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityRegistryIds** | **string** | Comma separated list of entity Registry IDs | 

### Return type

[**RegisteredEntitiesList**](RegisteredEntitiesList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistry

> Registry GetRegistry(ctx, registryId).Execute()

Get registry



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
    registryId := "registryId_example" // string | ID for the registry

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.GetRegistry(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.GetRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistry`: Registry
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.GetRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | ID for the registry | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Registry**](Registry.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBatchSchemasByRegistry

> BatchSchemasList ListBatchSchemasByRegistry(ctx, registryId).Execute()

List batch schemas by registry



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
    registryId := "registryId_example" // string | ID for the registry

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.ListBatchSchemasByRegistry(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ListBatchSchemasByRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBatchSchemasByRegistry`: BatchSchemasList
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.ListBatchSchemasByRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | ID for the registry | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBatchSchemasByRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BatchSchemasList**](BatchSchemasList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBoxSchemasByRegistry

> BoxSchemasList ListBoxSchemasByRegistry(ctx, registryId).Execute()

List box schemas by registry



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
    registryId := "registryId_example" // string | ID for the registry

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.ListBoxSchemasByRegistry(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ListBoxSchemasByRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBoxSchemasByRegistry`: BoxSchemasList
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.ListBoxSchemasByRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | ID for the registry | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBoxSchemasByRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BoxSchemasList**](BoxSchemasList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerSchemasByRegistry

> ContainerSchemasList ListContainerSchemasByRegistry(ctx, registryId).Execute()

List container schemas by registry



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
    registryId := "registryId_example" // string | ID for the registry

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.ListContainerSchemasByRegistry(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ListContainerSchemasByRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainerSchemasByRegistry`: ContainerSchemasList
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.ListContainerSchemasByRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | ID for the registry | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerSchemasByRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerSchemasList**](ContainerSchemasList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDropdownsByRegistry

> DropdownsRegistryList ListDropdownsByRegistry(ctx, registryId).Execute()

List dropdowns for a given registry



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
    registryId := "registryId_example" // string | ID of the registry to list the dropdowns for.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.ListDropdownsByRegistry(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ListDropdownsByRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDropdownsByRegistry`: DropdownsRegistryList
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.ListDropdownsByRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | ID of the registry to list the dropdowns for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDropdownsByRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DropdownsRegistryList**](DropdownsRegistryList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitySchemasByRegistry

> DeprecatedEntitySchemasList ListEntitySchemasByRegistry(ctx, registryId).Execute()

List entity schemas by registry



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
    registryId := "registryId_example" // string | ID for the registry

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.ListEntitySchemasByRegistry(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ListEntitySchemasByRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntitySchemasByRegistry`: DeprecatedEntitySchemasList
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.ListEntitySchemasByRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | ID for the registry | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntitySchemasByRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeprecatedEntitySchemasList**](DeprecatedEntitySchemasList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocationSchemasByRegistry

> LocationSchemasList ListLocationSchemasByRegistry(ctx, registryId).Execute()

List location schemas by registry



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
    registryId := "registryId_example" // string | ID for the registry

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.ListLocationSchemasByRegistry(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ListLocationSchemasByRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocationSchemasByRegistry`: LocationSchemasList
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.ListLocationSchemasByRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | ID for the registry | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLocationSchemasByRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LocationSchemasList**](LocationSchemasList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlateSchemasByRegistry

> PlateSchemasList ListPlateSchemasByRegistry(ctx, registryId).Execute()

List plate schemas by registry



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
    registryId := "registryId_example" // string | ID for the registry

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.ListPlateSchemasByRegistry(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ListPlateSchemasByRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPlateSchemasByRegistry`: PlateSchemasList
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.ListPlateSchemasByRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | ID for the registry | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPlateSchemasByRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlateSchemasList**](PlateSchemasList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRegistries

> RegistriesList ListRegistries(ctx).Name(name).ModifiedAt(modifiedAt).Execute()

List registries



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
    name := "name_example" // string | Name of a registry. Restricts results to those with the specified name. (optional)
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.ListRegistries(context.Background()).Name(name).ModifiedAt(modifiedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ListRegistries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRegistries`: RegistriesList
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.ListRegistries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRegistriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of a registry. Restricts results to those with the specified name. | 
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 

### Return type

[**RegistriesList**](RegistriesList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterEntities

> AsyncTaskLink RegisterEntities(ctx, registryId).RegisterEntities(registerEntities).Execute()

Register entities



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
    registryId := "registryId_example" // string | ID for the registry
    registerEntities := *openapiclient.NewRegisterEntities([]string{"EntityIds_example"}, openapiclient.NamingStrategy("NEW_IDS")) // RegisterEntities |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.RegisterEntities(context.Background(), registryId).RegisterEntities(registerEntities).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.RegisterEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterEntities`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.RegisterEntities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | ID for the registry | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registerEntities** | [**RegisterEntities**](RegisterEntities.md) |  | 

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


## UnregisterEntities

> map[string]interface{} UnregisterEntities(ctx, registryId).UnregisterEntities(unregisterEntities).Execute()

Unregister entities



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
    registryId := "registryId_example" // string | ID of the registry
    unregisterEntities := *openapiclient.NewUnregisterEntities([]string{"EntityIds_example"}, "FolderId_example") // UnregisterEntities |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.UnregisterEntities(context.Background(), registryId).UnregisterEntities(unregisterEntities).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.UnregisterEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnregisterEntities`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.UnregisterEntities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | ID of the registry | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unregisterEntities** | [**UnregisterEntities**](UnregisterEntities.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

