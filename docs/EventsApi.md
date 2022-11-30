# \EventsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEvents**](EventsApi.md#ListEvents) | **Get** /events | List Events



## ListEvents

> EventsPaginatedList ListEvents(ctx).PageSize(pageSize).NextToken(nextToken).CreatedAtGte(createdAtGte).StartingAfter(startingAfter).EventTypes(eventTypes).Poll(poll).Execute()

List Events



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
    createdAtGte := "createdAtGte_example" // string | Datetime, in RFC 3339 format. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. \"2020-05-23\".  (optional)
    startingAfter := "startingAfter_example" // string | Event ID after which events will be returned. (optional)
    eventTypes := "eventTypes_example" // string | Comma-separated list of event types to return. (optional)
    poll := true // bool | When True, the API will always return a nextToken to enable polling events indefinitely. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.ListEvents(context.Background()).PageSize(pageSize).NextToken(nextToken).CreatedAtGte(createdAtGte).StartingAfter(startingAfter).EventTypes(eventTypes).Poll(poll).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEvents`: EventsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100.  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **createdAtGte** | **string** | Datetime, in RFC 3339 format. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. \&quot;2020-05-23\&quot;.  | 
 **startingAfter** | **string** | Event ID after which events will be returned. | 
 **eventTypes** | **string** | Comma-separated list of event types to return. | 
 **poll** | **bool** | When True, the API will always return a nextToken to enable polling events indefinitely. | 

### Return type

[**EventsPaginatedList**](EventsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

