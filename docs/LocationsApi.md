# \LocationsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveLocations**](LocationsApi.md#ArchiveLocations) | **Post** /locations:archive | Archive locations
[**BulkGetLocations**](LocationsApi.md#BulkGetLocations) | **Get** /locations:bulk-get | BulkGet locations
[**CreateLocation**](LocationsApi.md#CreateLocation) | **Post** /locations | Create a location
[**GetLocation**](LocationsApi.md#GetLocation) | **Get** /locations/{location_id} | Get a location by ID
[**ListLocations**](LocationsApi.md#ListLocations) | **Get** /locations | List locations
[**UnarchiveLocations**](LocationsApi.md#UnarchiveLocations) | **Post** /locations:unarchive | Unarchive locations
[**UpdateLocation**](LocationsApi.md#UpdateLocation) | **Patch** /locations/{location_id} | Update a location



## ArchiveLocations

> LocationsArchivalChange ArchiveLocations(ctx).LocationsArchive(locationsArchive).Execute()

Archive locations



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
    locationsArchive := *openapiclient.NewLocationsArchive([]string{"LocationIds_example"}, "Reason_example") // LocationsArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.ArchiveLocations(context.Background()).LocationsArchive(locationsArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.ArchiveLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveLocations`: LocationsArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.ArchiveLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locationsArchive** | [**LocationsArchive**](LocationsArchive.md) |  | 

### Return type

[**LocationsArchivalChange**](LocationsArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkGetLocations

> LocationsBulkGet BulkGetLocations(ctx).LocationIds(locationIds).Barcodes(barcodes).Execute()

BulkGet locations



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
    locationIds := "locationIds_example" // string | Comma-separated list of location IDs. (optional)
    barcodes := "barcodes_example" // string | Comma-separated list of location barcodes. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.BulkGetLocations(context.Background()).LocationIds(locationIds).Barcodes(barcodes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.BulkGetLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetLocations`: LocationsBulkGet
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.BulkGetLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locationIds** | **string** | Comma-separated list of location IDs. | 
 **barcodes** | **string** | Comma-separated list of location barcodes. | 

### Return type

[**LocationsBulkGet**](LocationsBulkGet.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLocation

> Location CreateLocation(ctx).LocationCreate(locationCreate).Execute()

Create a location



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
    locationCreate := *openapiclient.NewLocationCreate("Name_example", "SchemaId_example") // LocationCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.CreateLocation(context.Background()).LocationCreate(locationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.CreateLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocation`: Location
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.CreateLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locationCreate** | [**LocationCreate**](LocationCreate.md) |  | 

### Return type

[**Location**](Location.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocation

> Location GetLocation(ctx, locationId).Execute()

Get a location by ID



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
    locationId := "locationId_example" // string | ID of location to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.GetLocation(context.Background(), locationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.GetLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocation`: Location
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.GetLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | ID of location to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Location**](Location.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocations

> LocationsPaginatedList ListLocations(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).SchemaId(schemaId).SchemaFields(schemaFields).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).AncestorStorageId(ancestorStorageId).ArchiveReason(archiveReason).Ids(ids).Barcodes(barcodes).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).Execute()

List locations



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
    pageSize := int32(56) // int32 | Number of results to return. (optional) (default to 50)
    nextToken := "nextToken_example" // string | Token for pagination (optional)
    sort := "sort_example" // string | Method by which to order search results. Valid sorts are barcode (barcode, alphabetical) modifiedAt (modified time, most recent first) and name (location name, alphabetical). Optionally add :asc or :desc to specify ascending or descending order. Default is modifiedAt.  (optional) (default to "modifiedAt")
    schemaId := "schemaId_example" // string | ID of a schema. Restricts results to those of the specified schema. (optional)
    schemaFields := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Schema field value. For Integer, Float, and Date type fields, supports the >= and <= operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  (optional)
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    name := "name_example" // string | Name of a location. Restricts results to those with the specified name. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of a location. Restricts results to those with names that include the provided substring.  (optional)
    ancestorStorageId := "ancestorStorageId_example" // string | ID of a plate, box, or location. Restricts results to those located in the specified storage. (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived locations. Use \"ANY_ARCHIVED\" to filter for archived locations regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    ids := "loc_9fxPzGDy,loc_fALwBTI7,loc_GyxUeUIi" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    barcodes := "b001, b002, f001" // string | Comma-separated list of barcodes. Matches all of the provided barcodes, or returns a 400 error that includes a list of which barcodes are invalid.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of users IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.ListLocations(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).SchemaId(schemaId).SchemaFields(schemaFields).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).AncestorStorageId(ancestorStorageId).ArchiveReason(archiveReason).Ids(ids).Barcodes(barcodes).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.ListLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocations`: LocationsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.ListLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** | Method by which to order search results. Valid sorts are barcode (barcode, alphabetical) modifiedAt (modified time, most recent first) and name (location name, alphabetical). Optionally add :asc or :desc to specify ascending or descending order. Default is modifiedAt.  | [default to &quot;modifiedAt&quot;]
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema. | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of a location. Restricts results to those with the specified name. | 
 **nameIncludes** | **string** | Name substring of a location. Restricts results to those with names that include the provided substring.  | 
 **ancestorStorageId** | **string** | ID of a plate, box, or location. Restricts results to those located in the specified storage. | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived locations. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived locations regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **barcodes** | **string** | Comma-separated list of barcodes. Matches all of the provided barcodes, or returns a 400 error that includes a list of which barcodes are invalid.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  | 
 **creatorIds** | **string** | Comma separated list of users IDs | 

### Return type

[**LocationsPaginatedList**](LocationsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveLocations

> LocationsArchivalChange UnarchiveLocations(ctx).LocationsUnarchive(locationsUnarchive).Execute()

Unarchive locations



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
    locationsUnarchive := *openapiclient.NewLocationsUnarchive([]string{"LocationIds_example"}) // LocationsUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.UnarchiveLocations(context.Background()).LocationsUnarchive(locationsUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.UnarchiveLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveLocations`: LocationsArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.UnarchiveLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locationsUnarchive** | [**LocationsUnarchive**](LocationsUnarchive.md) |  | 

### Return type

[**LocationsArchivalChange**](LocationsArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocation

> Location UpdateLocation(ctx, locationId).LocationUpdate(locationUpdate).Execute()

Update a location



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
    locationId := "locationId_example" // string | ID of the location to update
    locationUpdate := *openapiclient.NewLocationUpdate() // LocationUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.UpdateLocation(context.Background(), locationId).LocationUpdate(locationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.UpdateLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLocation`: Location
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.UpdateLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | ID of the location to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationUpdate** | [**LocationUpdate**](LocationUpdate.md) |  | 

### Return type

[**Location**](Location.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

