# \RequestsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateRequestTasks**](RequestsApi.md#BulkCreateRequestTasks) | **Post** /requests/{request_id}/tasks:bulk-create | Create tasks for a request
[**BulkGetRequests**](RequestsApi.md#BulkGetRequests) | **Get** /requests:bulk-get | Bulk get requests
[**BulkUpdateRequestTasks**](RequestsApi.md#BulkUpdateRequestTasks) | **Post** /requests/{request_id}/tasks:bulk-update | Bulk update tasks for a request
[**CreateRequest**](RequestsApi.md#CreateRequest) | **Post** /requests | Create a request
[**ExecuteRequestsSampleGroups**](RequestsApi.md#ExecuteRequestsSampleGroups) | **Post** /requests/{request_id}:execute-sample-groups | Update the status of sample groups in a request
[**GetRequest**](RequestsApi.md#GetRequest) | **Get** /requests/{request_id} | Get a request by ID
[**GetRequestFulfillment**](RequestsApi.md#GetRequestFulfillment) | **Get** /request-fulfillments/{request_fulfillment_id} | Get a request&#39;s fulfillment
[**GetRequestResponse**](RequestsApi.md#GetRequestResponse) | **Get** /requests/{request_id}/response | Get a request&#39;s response
[**ListRequestFulfillments**](RequestsApi.md#ListRequestFulfillments) | **Get** /request-fulfillments | List Request Fulfillments
[**ListRequests**](RequestsApi.md#ListRequests) | **Get** /requests | List requests
[**PatchRequest**](RequestsApi.md#PatchRequest) | **Patch** /requests/{request_id} | Update a request



## BulkCreateRequestTasks

> RequestTasksBulkCreateResponse BulkCreateRequestTasks(ctx, requestId).RequestTasksBulkCreateRequest(requestTasksBulkCreateRequest).Execute()

Create tasks for a request



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
    requestId := "req_JekfeyVS" // string | 
    requestTasksBulkCreateRequest := *openapiclient.NewRequestTasksBulkCreateRequest([]openapiclient.RequestTasksBulkCreate{*openapiclient.NewRequestTasksBulkCreate("reqtsksch_XHu79QRw")}) // RequestTasksBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.BulkCreateRequestTasks(context.Background(), requestId).RequestTasksBulkCreateRequest(requestTasksBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.BulkCreateRequestTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateRequestTasks`: RequestTasksBulkCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.BulkCreateRequestTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateRequestTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestTasksBulkCreateRequest** | [**RequestTasksBulkCreateRequest**](RequestTasksBulkCreateRequest.md) |  | 

### Return type

[**RequestTasksBulkCreateResponse**](RequestTasksBulkCreateResponse.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkGetRequests

> RequestsBulkGet BulkGetRequests(ctx).RequestIds(requestIds).DisplayIds(displayIds).Execute()

Bulk get requests



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
    requestIds := "req_xJk20sla,req_lQJ3nMs5" // string | Comma-separated list of request IDs. Exactly one of requestIds or displayIds must be specified. (optional)
    displayIds := "VPR001,VPR002" // string | Comma-separated list of display IDs. Exactly one of requestIds or displayIds must be specified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.BulkGetRequests(context.Background()).RequestIds(requestIds).DisplayIds(displayIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.BulkGetRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetRequests`: RequestsBulkGet
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.BulkGetRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestIds** | **string** | Comma-separated list of request IDs. Exactly one of requestIds or displayIds must be specified. | 
 **displayIds** | **string** | Comma-separated list of display IDs. Exactly one of requestIds or displayIds must be specified. | 

### Return type

[**RequestsBulkGet**](RequestsBulkGet.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateRequestTasks

> RequestTasksBulkUpdateResponse BulkUpdateRequestTasks(ctx, requestId).RequestTasksBulkUpdateRequest(requestTasksBulkUpdateRequest).Execute()

Bulk update tasks for a request



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
    requestId := "req_JekfeyVS" // string | 
    requestTasksBulkUpdateRequest := *openapiclient.NewRequestTasksBulkUpdateRequest([]openapiclient.RequestTaskBase{*openapiclient.NewRequestTaskBase("reqtsk_PFHQ8rmb")}) // RequestTasksBulkUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.BulkUpdateRequestTasks(context.Background(), requestId).RequestTasksBulkUpdateRequest(requestTasksBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.BulkUpdateRequestTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateRequestTasks`: RequestTasksBulkUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.BulkUpdateRequestTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateRequestTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestTasksBulkUpdateRequest** | [**RequestTasksBulkUpdateRequest**](RequestTasksBulkUpdateRequest.md) |  | 

### Return type

[**RequestTasksBulkUpdateResponse**](RequestTasksBulkUpdateResponse.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRequest

> Request CreateRequest(ctx).RequestCreate(requestCreate).Execute()

Create a request



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
    requestCreate := *openapiclient.NewRequestCreate("TODO", "src_29pt8Ida", []openapiclient.RequestSampleGroupCreate{*openapiclient.NewRequestSampleGroupCreate(map[string]RequestSampleGroupSamplesValue{"key": openapiclient.RequestSampleGroupSamples_value{RequestSampleWithBatch: openapiclient.NewRequestSampleWithBatch("bat_XK0UnLyk")}})}, "assaysch_3IF58QOf") // RequestCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.CreateRequest(context.Background()).RequestCreate(requestCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.CreateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRequest`: Request
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.CreateRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestCreate** | [**RequestCreate**](RequestCreate.md) |  | 

### Return type

[**Request**](Request.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteRequestsSampleGroups

> map[string]interface{} ExecuteRequestsSampleGroups(ctx, requestId).SampleGroupsStatusUpdate(sampleGroupsStatusUpdate).Execute()

Update the status of sample groups in a request



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
    requestId := "req_JekfeyVS" // string | 
    sampleGroupsStatusUpdate := *openapiclient.NewSampleGroupsStatusUpdate([]openapiclient.SampleGroupStatusUpdate{*openapiclient.NewSampleGroupStatusUpdate("sgrp_2GPCXk6", openapiclient.SampleGroupStatus("IN_PROGRESS"))}) // SampleGroupsStatusUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.ExecuteRequestsSampleGroups(context.Background(), requestId).SampleGroupsStatusUpdate(sampleGroupsStatusUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.ExecuteRequestsSampleGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteRequestsSampleGroups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.ExecuteRequestsSampleGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteRequestsSampleGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sampleGroupsStatusUpdate** | [**SampleGroupsStatusUpdate**](SampleGroupsStatusUpdate.md) |  | 

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


## GetRequest

> Request GetRequest(ctx, requestId).Execute()

Get a request by ID



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
    requestId := "req_JekfeyVS" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.GetRequest(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.GetRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequest`: Request
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.GetRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Request**](Request.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestFulfillment

> RequestFulfillment GetRequestFulfillment(ctx, requestFulfillmentId).Execute()

Get a request's fulfillment



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
    requestFulfillmentId := "reqffm_8Hm71Usw" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.GetRequestFulfillment(context.Background(), requestFulfillmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.GetRequestFulfillment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestFulfillment`: RequestFulfillment
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.GetRequestFulfillment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestFulfillmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestFulfillmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestFulfillment**](RequestFulfillment.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestResponse

> RequestResponse GetRequestResponse(ctx, requestId).Execute()

Get a request's response



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
    requestId := "req_JekfeyVS" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.GetRequestResponse(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.GetRequestResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestResponse`: RequestResponse
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.GetRequestResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestResponse**](RequestResponse.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRequestFulfillments

> RequestFulfillmentsPaginatedList ListRequestFulfillments(ctx).EntryId(entryId).ModifiedAt(modifiedAt).NextToken(nextToken).PageSize(pageSize).Execute()

List Request Fulfillments



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
    entryId := "etr_IKwdYx31" // string | 
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    nextToken := "Im5ldyB0ZXN0Ig==" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.ListRequestFulfillments(context.Background()).EntryId(entryId).ModifiedAt(modifiedAt).NextToken(nextToken).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.ListRequestFulfillments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRequestFulfillments`: RequestFulfillmentsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.ListRequestFulfillments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequestFulfillmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entryId** | **string** |  | 
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]

### Return type

[**RequestFulfillmentsPaginatedList**](RequestFulfillmentsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRequests

> RequestsPaginatedList ListRequests(ctx).SchemaId(schemaId).RequestStatus(requestStatus).MinCreatedTime(minCreatedTime).MaxCreatedTime(maxCreatedTime).NextToken(nextToken).PageSize(pageSize).Execute()

List requests



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
    schemaId := "schemaId_example" // string | 
    requestStatus := openapiclient.RequestStatus("REQUESTED") // RequestStatus |  (optional)
    minCreatedTime := int32(56) // int32 | minimum create time (unix seconds) (optional)
    maxCreatedTime := int32(56) // int32 | maximum create time (unix seconds) (optional)
    nextToken := "Im5ldyB0ZXN0Ig==" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.ListRequests(context.Background()).SchemaId(schemaId).RequestStatus(requestStatus).MinCreatedTime(minCreatedTime).MaxCreatedTime(maxCreatedTime).NextToken(nextToken).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.ListRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRequests`: RequestsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.ListRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaId** | **string** |  | 
 **requestStatus** | [**RequestStatus**](RequestStatus.md) |  | 
 **minCreatedTime** | **int32** | minimum create time (unix seconds) | 
 **maxCreatedTime** | **int32** | maximum create time (unix seconds) | 
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]

### Return type

[**RequestsPaginatedList**](RequestsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRequest

> Request PatchRequest(ctx, requestId).RequestUpdate(requestUpdate).Execute()

Update a request



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
    requestId := "requestId_example" // string | 
    requestUpdate := *openapiclient.NewRequestUpdate() // RequestUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestsApi.PatchRequest(context.Background(), requestId).RequestUpdate(requestUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestsApi.PatchRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchRequest`: Request
    fmt.Fprintf(os.Stdout, "Response from `RequestsApi.PatchRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestUpdate** | [**RequestUpdate**](RequestUpdate.md) |  | 

### Return type

[**Request**](Request.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

