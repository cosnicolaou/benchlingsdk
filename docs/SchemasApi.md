# \SchemasApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBatchSchema**](SchemasApi.md#GetBatchSchema) | **Get** /batch-schemas/{schema_id} | Get a batch schema by ID
[**GetBoxSchema**](SchemasApi.md#GetBoxSchema) | **Get** /box-schemas/{schema_id} | Get a box schema by ID
[**GetContainerSchema**](SchemasApi.md#GetContainerSchema) | **Get** /container-schemas/{schema_id} | Get a container schema by ID
[**GetEntitySchema**](SchemasApi.md#GetEntitySchema) | **Get** /entity-schemas/{schema_id} | Get an entity schema by ID
[**GetEntrySchema**](SchemasApi.md#GetEntrySchema) | **Get** /entry-schemas/{schema_id} | Get an Entry schema by ID
[**GetLocationSchema**](SchemasApi.md#GetLocationSchema) | **Get** /location-schemas/{schema_id} | Get a location schema by ID
[**GetPlateSchema**](SchemasApi.md#GetPlateSchema) | **Get** /plate-schemas/{schema_id} | Get a plate schema by ID
[**GetRequestSchema**](SchemasApi.md#GetRequestSchema) | **Get** /request-schemas/{schema_id} | Get a Request schema by ID
[**GetRequestTaskSchema**](SchemasApi.md#GetRequestTaskSchema) | **Get** /request-task-schemas/{schema_id} | Get a Request Task schema by ID
[**GetResultSchema**](SchemasApi.md#GetResultSchema) | **Get** /assay-result-schemas/{schema_id} | Get a Result schema by ID
[**GetRunSchema**](SchemasApi.md#GetRunSchema) | **Get** /assay-run-schemas/{schema_id} | Get a Run schema by ID
[**GetWorkflowTaskSchema**](SchemasApi.md#GetWorkflowTaskSchema) | **Get** /workflow-task-schemas/{schema_id} | Get a workflow task schema
[**ListAssayResultSchemas**](SchemasApi.md#ListAssayResultSchemas) | **Get** /assay-result-schemas | List assay result schemas
[**ListAssayRunSchemas**](SchemasApi.md#ListAssayRunSchemas) | **Get** /assay-run-schemas | List assay run schemas
[**ListBatchSchemas**](SchemasApi.md#ListBatchSchemas) | **Get** /batch-schemas | List batch schemas
[**ListBoxSchemas**](SchemasApi.md#ListBoxSchemas) | **Get** /box-schemas | List box schemas
[**ListContainerSchemas**](SchemasApi.md#ListContainerSchemas) | **Get** /container-schemas | List container schemas
[**ListEntitySchemas**](SchemasApi.md#ListEntitySchemas) | **Get** /entity-schemas | List entity schemas
[**ListEntrySchemas**](SchemasApi.md#ListEntrySchemas) | **Get** /entry-schemas | List entry schemas
[**ListLocationSchemas**](SchemasApi.md#ListLocationSchemas) | **Get** /location-schemas | List location schemas
[**ListPlateSchemas**](SchemasApi.md#ListPlateSchemas) | **Get** /plate-schemas | List plate schemas
[**ListRequestSchemas**](SchemasApi.md#ListRequestSchemas) | **Get** /request-schemas | List request schemas
[**ListRequestTaskSchemas**](SchemasApi.md#ListRequestTaskSchemas) | **Get** /request-task-schemas | List request task schemas
[**ListWorkflowTaskSchemas**](SchemasApi.md#ListWorkflowTaskSchemas) | **Get** /workflow-task-schemas | List workflow task schemas



## GetBatchSchema

> BatchSchema GetBatchSchema(ctx, schemaId).Execute()

Get a batch schema by ID



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
    schemaId := "schemaId_example" // string | ID of schema to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.GetBatchSchema(context.Background(), schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetBatchSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBatchSchema`: BatchSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetBatchSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** | ID of schema to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BatchSchema**](BatchSchema.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoxSchema

> BoxSchema GetBoxSchema(ctx, schemaId).Execute()

Get a box schema by ID



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
    schemaId := "schemaId_example" // string | ID of schema to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.GetBoxSchema(context.Background(), schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetBoxSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBoxSchema`: BoxSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetBoxSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** | ID of schema to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoxSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BoxSchema**](BoxSchema.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerSchema

> ContainerSchema GetContainerSchema(ctx, schemaId).Execute()

Get a container schema by ID



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
    schemaId := "schemaId_example" // string | ID of schema to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.GetContainerSchema(context.Background(), schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetContainerSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerSchema`: ContainerSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetContainerSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** | ID of schema to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerSchema**](ContainerSchema.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitySchema

> EntitySchema GetEntitySchema(ctx, schemaId).Execute()

Get an entity schema by ID



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
    schemaId := "schemaId_example" // string | ID of schema to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.GetEntitySchema(context.Background(), schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetEntitySchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntitySchema`: EntitySchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetEntitySchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** | ID of schema to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitySchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntitySchema**](EntitySchema.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntrySchema

> EntrySchemaDetailed GetEntrySchema(ctx, schemaId).Execute()

Get an Entry schema by ID



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
    schemaId := "schemaId_example" // string | ID of schema to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.GetEntrySchema(context.Background(), schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetEntrySchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntrySchema`: EntrySchemaDetailed
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetEntrySchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** | ID of schema to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntrySchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntrySchemaDetailed**](EntrySchemaDetailed.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocationSchema

> LocationSchema GetLocationSchema(ctx, schemaId).Execute()

Get a location schema by ID



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
    schemaId := "schemaId_example" // string | ID of schema to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.GetLocationSchema(context.Background(), schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetLocationSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocationSchema`: LocationSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetLocationSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** | ID of schema to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LocationSchema**](LocationSchema.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlateSchema

> PlateSchema GetPlateSchema(ctx, schemaId).Execute()

Get a plate schema by ID



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
    schemaId := "schemaId_example" // string | ID of schema to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.GetPlateSchema(context.Background(), schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetPlateSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlateSchema`: PlateSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetPlateSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** | ID of schema to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlateSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlateSchema**](PlateSchema.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestSchema

> RequestSchema GetRequestSchema(ctx, schemaId).Execute()

Get a Request schema by ID



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
    schemaId := "schemaId_example" // string | ID of schema to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.GetRequestSchema(context.Background(), schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetRequestSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestSchema`: RequestSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetRequestSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** | ID of schema to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestSchema**](RequestSchema.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestTaskSchema

> RequestTaskSchema GetRequestTaskSchema(ctx, schemaId).Execute()

Get a Request Task schema by ID



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
    schemaId := "schemaId_example" // string | ID of schema to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.GetRequestTaskSchema(context.Background(), schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetRequestTaskSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestTaskSchema`: RequestTaskSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetRequestTaskSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** | ID of schema to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestTaskSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestTaskSchema**](RequestTaskSchema.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResultSchema

> AssayResultSchema GetResultSchema(ctx, schemaId).Execute()

Get a Result schema by ID



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
    schemaId := "schemaId_example" // string | ID of schema to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.GetResultSchema(context.Background(), schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetResultSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResultSchema`: AssayResultSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetResultSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** | ID of schema to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResultSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssayResultSchema**](AssayResultSchema.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunSchema

> AssayRunSchema GetRunSchema(ctx, schemaId).Execute()

Get a Run schema by ID



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
    schemaId := "schemaId_example" // string | ID of schema to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.GetRunSchema(context.Background(), schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetRunSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRunSchema`: AssayRunSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetRunSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** | ID of schema to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AssayRunSchema**](AssayRunSchema.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowTaskSchema

> WorkflowTaskSchema GetWorkflowTaskSchema(ctx, schemaId).Execute()

Get a workflow task schema



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
    schemaId := "prstsch_KnR9iVum" // string | The workflow task schema ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.GetWorkflowTaskSchema(context.Background(), schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.GetWorkflowTaskSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowTaskSchema`: WorkflowTaskSchema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.GetWorkflowTaskSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** | The workflow task schema ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowTaskSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowTaskSchema**](WorkflowTaskSchema.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssayResultSchemas

> AssayResultSchemasPaginatedList ListAssayResultSchemas(ctx).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()

List assay result schemas



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
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.ListAssayResultSchemas(context.Background()).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ListAssayResultSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAssayResultSchemas`: AssayResultSchemasPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ListAssayResultSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssayResultSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 

### Return type

[**AssayResultSchemasPaginatedList**](AssayResultSchemasPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssayRunSchemas

> AssayRunSchemasPaginatedList ListAssayRunSchemas(ctx).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()

List assay run schemas



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
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.ListAssayRunSchemas(context.Background()).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ListAssayRunSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAssayRunSchemas`: AssayRunSchemasPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ListAssayRunSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssayRunSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 

### Return type

[**AssayRunSchemasPaginatedList**](AssayRunSchemasPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBatchSchemas

> BatchSchemasPaginatedList ListBatchSchemas(ctx).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()

List batch schemas



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
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.ListBatchSchemas(context.Background()).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ListBatchSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBatchSchemas`: BatchSchemasPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ListBatchSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBatchSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 

### Return type

[**BatchSchemasPaginatedList**](BatchSchemasPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBoxSchemas

> BoxSchemasPaginatedList ListBoxSchemas(ctx).NextToken(nextToken).PageSize(pageSize).Execute()

List box schemas



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
    resp, r, err := apiClient.SchemasApi.ListBoxSchemas(context.Background()).NextToken(nextToken).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ListBoxSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBoxSchemas`: BoxSchemasPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ListBoxSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBoxSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]

### Return type

[**BoxSchemasPaginatedList**](BoxSchemasPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerSchemas

> ContainerSchemasPaginatedList ListContainerSchemas(ctx).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()

List container schemas



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
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.ListContainerSchemas(context.Background()).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ListContainerSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainerSchemas`: ContainerSchemasPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ListContainerSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContainerSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 

### Return type

[**ContainerSchemasPaginatedList**](ContainerSchemasPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitySchemas

> EntitySchemasPaginatedList ListEntitySchemas(ctx).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()

List entity schemas



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
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.ListEntitySchemas(context.Background()).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ListEntitySchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntitySchemas`: EntitySchemasPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ListEntitySchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEntitySchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 

### Return type

[**EntitySchemasPaginatedList**](EntitySchemasPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntrySchemas

> EntrySchemasPaginatedList ListEntrySchemas(ctx).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()

List entry schemas



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
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.ListEntrySchemas(context.Background()).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ListEntrySchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntrySchemas`: EntrySchemasPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ListEntrySchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEntrySchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 

### Return type

[**EntrySchemasPaginatedList**](EntrySchemasPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocationSchemas

> LocationSchemasPaginatedList ListLocationSchemas(ctx).NextToken(nextToken).PageSize(pageSize).Execute()

List location schemas



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
    resp, r, err := apiClient.SchemasApi.ListLocationSchemas(context.Background()).NextToken(nextToken).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ListLocationSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocationSchemas`: LocationSchemasPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ListLocationSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLocationSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]

### Return type

[**LocationSchemasPaginatedList**](LocationSchemasPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlateSchemas

> PlateSchemasPaginatedList ListPlateSchemas(ctx).NextToken(nextToken).PageSize(pageSize).Execute()

List plate schemas



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
    resp, r, err := apiClient.SchemasApi.ListPlateSchemas(context.Background()).NextToken(nextToken).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ListPlateSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPlateSchemas`: PlateSchemasPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ListPlateSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPlateSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]

### Return type

[**PlateSchemasPaginatedList**](PlateSchemasPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRequestSchemas

> RequestSchemasPaginatedList ListRequestSchemas(ctx).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()

List request schemas



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
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.ListRequestSchemas(context.Background()).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ListRequestSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRequestSchemas`: RequestSchemasPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ListRequestSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequestSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 

### Return type

[**RequestSchemasPaginatedList**](RequestSchemasPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRequestTaskSchemas

> RequestTaskSchemasPaginatedList ListRequestTaskSchemas(ctx).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()

List request task schemas



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
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.ListRequestTaskSchemas(context.Background()).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ListRequestTaskSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRequestTaskSchemas`: RequestTaskSchemasPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ListRequestTaskSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequestTaskSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 

### Return type

[**RequestTaskSchemasPaginatedList**](RequestTaskSchemasPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowTaskSchemas

> WorkflowTaskSchemasPaginatedList ListWorkflowTaskSchemas(ctx).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()

List workflow task schemas



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
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.ListWorkflowTaskSchemas(context.Background()).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ListWorkflowTaskSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowTaskSchemas`: WorkflowTaskSchemasPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ListWorkflowTaskSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowTaskSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 

### Return type

[**WorkflowTaskSchemasPaginatedList**](WorkflowTaskSchemasPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

