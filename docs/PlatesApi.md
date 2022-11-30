# \PlatesApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchivePlates**](PlatesApi.md#ArchivePlates) | **Post** /plates:archive | Archive plates
[**BulkGetPlates**](PlatesApi.md#BulkGetPlates) | **Get** /plates:bulk-get | BulkGet plates
[**CreatePlate**](PlatesApi.md#CreatePlate) | **Post** /plates | Create a plate
[**GetPlate**](PlatesApi.md#GetPlate) | **Get** /plates/{plate_id} | Get a plate
[**ListPlates**](PlatesApi.md#ListPlates) | **Get** /plates | List plates
[**UnarchivePlates**](PlatesApi.md#UnarchivePlates) | **Post** /plates:unarchive | Unarchive plates
[**UpdatePlate**](PlatesApi.md#UpdatePlate) | **Patch** /plates/{plate_id} | Update a plate



## ArchivePlates

> PlatesArchivalChange ArchivePlates(ctx).PlatesArchive(platesArchive).Execute()

Archive plates



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
    platesArchive := *openapiclient.NewPlatesArchive([]string{"PlateIds_example"}, "Reason_example", false) // PlatesArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlatesApi.ArchivePlates(context.Background()).PlatesArchive(platesArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatesApi.ArchivePlates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchivePlates`: PlatesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `PlatesApi.ArchivePlates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchivePlatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platesArchive** | [**PlatesArchive**](PlatesArchive.md) |  | 

### Return type

[**PlatesArchivalChange**](PlatesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkGetPlates

> PlatesBulkGet BulkGetPlates(ctx).PlateIds(plateIds).Barcodes(barcodes).Returning(returning).Execute()

BulkGet plates



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
    plateIds := "plateIds_example" // string | Comma-separated list of plate IDs. (optional)
    barcodes := "barcodes_example" // string | Comma-separated list of plate barcodes. (optional)
    returning := "plates.id,plates.wells.*.webURL" // string | Comma-separated list of fields to return. Modifies the output shape. To return all keys at a given level, enumerate them or use the wildcard, '*'. For more information, [click here](https://docs.benchling.com/docs/getting-started-1#returning-query-parameter).  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlatesApi.BulkGetPlates(context.Background()).PlateIds(plateIds).Barcodes(barcodes).Returning(returning).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatesApi.BulkGetPlates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetPlates`: PlatesBulkGet
    fmt.Fprintf(os.Stdout, "Response from `PlatesApi.BulkGetPlates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetPlatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plateIds** | **string** | Comma-separated list of plate IDs. | 
 **barcodes** | **string** | Comma-separated list of plate barcodes. | 
 **returning** | **string** | Comma-separated list of fields to return. Modifies the output shape. To return all keys at a given level, enumerate them or use the wildcard, &#39;*&#39;. For more information, [click here](https://docs.benchling.com/docs/getting-started-1#returning-query-parameter).  | 

### Return type

[**PlatesBulkGet**](PlatesBulkGet.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlate

> Plate CreatePlate(ctx).Returning(returning).PlateCreate(plateCreate).Execute()

Create a plate



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
    returning := "id,webURL" // string | Comma-separated list of fields to return. Modifies the output shape. To return all keys at a given level, enumerate them or use the wildcard, '*'. For more information, [click here](https://docs.benchling.com/docs/getting-started-1#returning-query-parameter).  (optional)
    plateCreate := *openapiclient.NewPlateCreate("SchemaId_example") // PlateCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlatesApi.CreatePlate(context.Background()).Returning(returning).PlateCreate(plateCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatesApi.CreatePlate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlate`: Plate
    fmt.Fprintf(os.Stdout, "Response from `PlatesApi.CreatePlate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returning** | **string** | Comma-separated list of fields to return. Modifies the output shape. To return all keys at a given level, enumerate them or use the wildcard, &#39;*&#39;. For more information, [click here](https://docs.benchling.com/docs/getting-started-1#returning-query-parameter).  | 
 **plateCreate** | [**PlateCreate**](PlateCreate.md) |  | 

### Return type

[**Plate**](Plate.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlate

> Plate GetPlate(ctx, plateId).Returning(returning).Execute()

Get a plate



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
    plateId := "plateId_example" // string | 
    returning := "id,webURL" // string | Comma-separated list of fields to return. Modifies the output shape. To return all keys at a given level, enumerate them or use the wildcard, '*'. For more information, [click here](https://docs.benchling.com/docs/getting-started-1#returning-query-parameter).  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlatesApi.GetPlate(context.Background(), plateId).Returning(returning).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatesApi.GetPlate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlate`: Plate
    fmt.Fprintf(os.Stdout, "Response from `PlatesApi.GetPlate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returning** | **string** | Comma-separated list of fields to return. Modifies the output shape. To return all keys at a given level, enumerate them or use the wildcard, &#39;*&#39;. For more information, [click here](https://docs.benchling.com/docs/getting-started-1#returning-query-parameter).  | 

### Return type

[**Plate**](Plate.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlates

> PlatesPaginatedList ListPlates(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).SchemaId(schemaId).SchemaFields(schemaFields).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).AncestorStorageId(ancestorStorageId).StorageContentsId(storageContentsId).StorageContentsIds(storageContentsIds).ArchiveReason(archiveReason).Ids(ids).Barcodes(barcodes).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).Returning(returning).CreatorIds(creatorIds).Execute()

List plates



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
    sort := "sort_example" // string |  (optional) (default to "modifiedAt:desc")
    schemaId := "schemaId_example" // string | ID of a schema. Restricts results to those of the specified schema.  (optional)
    schemaFields := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Schema field value. For Integer, Float, and Date type fields, supports the >= and <= operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  (optional)
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    name := "name_example" // string | Name of a plate. Restricts results to those with the specified name. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of a plate. Restricts results to those with names that include the provided substring.  (optional)
    ancestorStorageId := "ancestorStorageId_example" // string | ID of a location. Restricts results to those located in the specified storage.  (optional)
    storageContentsId := "storageContentsId_example" // string | ID of a batch, entity, or entity schema. Restricts results to those that hold containers with entities or batches associated with the specified ID.  (optional)
    storageContentsIds := "storageContentsIds_example" // string | Comma-separated list of IDs of batches or entities. Restricts results to those that hold containers with at least one of the specified batches, entities, or batches of the specified entities.  (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived plates. Use \"ANY_ARCHIVED\" to filter for archived plates regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    ids := "plt_xd4hj4eB,plt_xd4hj4C" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    barcodes := "W102477,W102478" // string | Comma-separated list of barcodes. Matches all of the provided barcodes, or returns a 400 error that includes a list of which barcodes are invalid.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  (optional)
    returning := "plates.id,plates.wells.*.webURL" // string | Comma-separated list of fields to return. Modifies the output shape. To return all keys at a given level, enumerate them or use the wildcard, '*'. For more information, [click here](https://docs.benchling.com/docs/getting-started-1#returning-query-parameter).  (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of users IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlatesApi.ListPlates(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).SchemaId(schemaId).SchemaFields(schemaFields).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).AncestorStorageId(ancestorStorageId).StorageContentsId(storageContentsId).StorageContentsIds(storageContentsIds).ArchiveReason(archiveReason).Ids(ids).Barcodes(barcodes).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).Returning(returning).CreatorIds(creatorIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatesApi.ListPlates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPlates`: PlatesPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `PlatesApi.ListPlates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPlatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100.  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema.  | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of a plate. Restricts results to those with the specified name. | 
 **nameIncludes** | **string** | Name substring of a plate. Restricts results to those with names that include the provided substring.  | 
 **ancestorStorageId** | **string** | ID of a location. Restricts results to those located in the specified storage.  | 
 **storageContentsId** | **string** | ID of a batch, entity, or entity schema. Restricts results to those that hold containers with entities or batches associated with the specified ID.  | 
 **storageContentsIds** | **string** | Comma-separated list of IDs of batches or entities. Restricts results to those that hold containers with at least one of the specified batches, entities, or batches of the specified entities.  | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived plates. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived plates regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **barcodes** | **string** | Comma-separated list of barcodes. Matches all of the provided barcodes, or returns a 400 error that includes a list of which barcodes are invalid.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  | 
 **returning** | **string** | Comma-separated list of fields to return. Modifies the output shape. To return all keys at a given level, enumerate them or use the wildcard, &#39;*&#39;. For more information, [click here](https://docs.benchling.com/docs/getting-started-1#returning-query-parameter).  | 
 **creatorIds** | **string** | Comma separated list of users IDs | 

### Return type

[**PlatesPaginatedList**](PlatesPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchivePlates

> PlatesArchivalChange UnarchivePlates(ctx).PlatesUnarchive(platesUnarchive).Execute()

Unarchive plates



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
    platesUnarchive := *openapiclient.NewPlatesUnarchive([]string{"PlateIds_example"}) // PlatesUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlatesApi.UnarchivePlates(context.Background()).PlatesUnarchive(platesUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatesApi.UnarchivePlates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchivePlates`: PlatesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `PlatesApi.UnarchivePlates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchivePlatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platesUnarchive** | [**PlatesUnarchive**](PlatesUnarchive.md) |  | 

### Return type

[**PlatesArchivalChange**](PlatesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlate

> Plate UpdatePlate(ctx, plateId).Returning(returning).PlateUpdate(plateUpdate).Execute()

Update a plate



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
    plateId := "plateId_example" // string | 
    returning := "id,webURL" // string | Comma-separated list of fields to return. Modifies the output shape. To return all keys at a given level, enumerate them or use the wildcard, '*'. For more information, [click here](https://docs.benchling.com/docs/getting-started-1#returning-query-parameter).  (optional)
    plateUpdate := *openapiclient.NewPlateUpdate() // PlateUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlatesApi.UpdatePlate(context.Background(), plateId).Returning(returning).PlateUpdate(plateUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatesApi.UpdatePlate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePlate`: Plate
    fmt.Fprintf(os.Stdout, "Response from `PlatesApi.UpdatePlate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returning** | **string** | Comma-separated list of fields to return. Modifies the output shape. To return all keys at a given level, enumerate them or use the wildcard, &#39;*&#39;. For more information, [click here](https://docs.benchling.com/docs/getting-started-1#returning-query-parameter).  | 
 **plateUpdate** | [**PlateUpdate**](PlateUpdate.md) |  | 

### Return type

[**Plate**](Plate.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

