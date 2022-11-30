# \BoxesApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveBoxes**](BoxesApi.md#ArchiveBoxes) | **Post** /boxes:archive | Archive boxes
[**BulkGetBoxes**](BoxesApi.md#BulkGetBoxes) | **Get** /boxes:bulk-get | BulkGet boxes
[**CreateBox**](BoxesApi.md#CreateBox) | **Post** /boxes | Create a box
[**GetBox**](BoxesApi.md#GetBox) | **Get** /boxes/{box_id} | Get a box
[**ListBoxContents**](BoxesApi.md#ListBoxContents) | **Get** /boxes/{box_id}/contents | List a box&#39;s contents
[**ListBoxes**](BoxesApi.md#ListBoxes) | **Get** /boxes | List boxes
[**UnarchiveBoxes**](BoxesApi.md#UnarchiveBoxes) | **Post** /boxes:unarchive | Unarchive boxes
[**UpdateBox**](BoxesApi.md#UpdateBox) | **Patch** /boxes/{box_id} | Update a box



## ArchiveBoxes

> BoxesArchivalChange ArchiveBoxes(ctx).BoxesArchive(boxesArchive).Execute()

Archive boxes



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
    boxesArchive := *openapiclient.NewBoxesArchive([]string{"BoxIds_example"}, "Reason_example") // BoxesArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BoxesApi.ArchiveBoxes(context.Background()).BoxesArchive(boxesArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoxesApi.ArchiveBoxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveBoxes`: BoxesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `BoxesApi.ArchiveBoxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveBoxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **boxesArchive** | [**BoxesArchive**](BoxesArchive.md) |  | 

### Return type

[**BoxesArchivalChange**](BoxesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkGetBoxes

> BoxesBulkGet BulkGetBoxes(ctx).BoxIds(boxIds).Barcodes(barcodes).Execute()

BulkGet boxes



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
    boxIds := "boxIds_example" // string |  (optional)
    barcodes := "barcodes_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BoxesApi.BulkGetBoxes(context.Background()).BoxIds(boxIds).Barcodes(barcodes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoxesApi.BulkGetBoxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetBoxes`: BoxesBulkGet
    fmt.Fprintf(os.Stdout, "Response from `BoxesApi.BulkGetBoxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetBoxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **boxIds** | **string** |  | 
 **barcodes** | **string** |  | 

### Return type

[**BoxesBulkGet**](BoxesBulkGet.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBox

> Box CreateBox(ctx).BoxCreate(boxCreate).Execute()

Create a box



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
    boxCreate := *openapiclient.NewBoxCreate("SchemaId_example") // BoxCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BoxesApi.CreateBox(context.Background()).BoxCreate(boxCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoxesApi.CreateBox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBox`: Box
    fmt.Fprintf(os.Stdout, "Response from `BoxesApi.CreateBox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBoxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **boxCreate** | [**BoxCreate**](BoxCreate.md) |  | 

### Return type

[**Box**](Box.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBox

> Box GetBox(ctx, boxId).Execute()

Get a box



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
    boxId := "boxId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BoxesApi.GetBox(context.Background(), boxId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoxesApi.GetBox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBox`: Box
    fmt.Fprintf(os.Stdout, "Response from `BoxesApi.GetBox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Box**](Box.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBoxContents

> BoxContentsPaginatedList ListBoxContents(ctx, boxId).PageSize(pageSize).NextToken(nextToken).Execute()

List a box's contents



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
    boxId := "boxId_example" // string | 
    pageSize := int32(56) // int32 | Number of results to return. (optional) (default to 50)
    nextToken := "nextToken_example" // string | Token for pagination (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BoxesApi.ListBoxContents(context.Background(), boxId).PageSize(pageSize).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoxesApi.ListBoxContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBoxContents`: BoxContentsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `BoxesApi.ListBoxContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBoxContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | Number of results to return. | [default to 50]
 **nextToken** | **string** | Token for pagination | 

### Return type

[**BoxContentsPaginatedList**](BoxContentsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBoxes

> BoxesPaginatedList ListBoxes(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).SchemaId(schemaId).SchemaFields(schemaFields).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).EmptyPositions(emptyPositions).EmptyPositionsGte(emptyPositionsGte).EmptyPositionsGt(emptyPositionsGt).EmptyPositionsLte(emptyPositionsLte).EmptyPositionsLt(emptyPositionsLt).EmptyContainers(emptyContainers).EmptyContainersGte(emptyContainersGte).EmptyContainersGt(emptyContainersGt).EmptyContainersLte(emptyContainersLte).EmptyContainersLt(emptyContainersLt).AncestorStorageId(ancestorStorageId).StorageContentsId(storageContentsId).StorageContentsIds(storageContentsIds).ArchiveReason(archiveReason).Ids(ids).Barcodes(barcodes).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).Execute()

List boxes



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
    name := "name_example" // string | Name of a box. Restricts results to those with the specified name. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of a box. Restricts results to those with names that include the provided substring.  (optional)
    emptyPositions := int32(56) // int32 | Only return boxes that have the specified number of empty positions  (optional)
    emptyPositionsGte := int32(56) // int32 | Only return boxes that have greater-than or equal-to the specified number of empty positions.  (optional)
    emptyPositionsGt := int32(56) // int32 | Only return boxes that have greater-than the specified number of empty positions.  (optional)
    emptyPositionsLte := int32(56) // int32 | Only return boxes that have less-than or equal-to the specified number of empty positions.  (optional)
    emptyPositionsLt := int32(56) // int32 | Only return boxes that have less-than the specified number of empty positions.  (optional)
    emptyContainers := int32(56) // int32 | Only return boxes that have the specified number of empty containers (containers without contents).  (optional)
    emptyContainersGte := int32(56) // int32 | Only return boxes that have greater-than or equal-to the specified number of empty containers (containers without contents).  (optional)
    emptyContainersGt := int32(56) // int32 | Only return boxes that have greater-than the specified number of empty containers (containers without contents).  (optional)
    emptyContainersLte := int32(56) // int32 | Only return boxes that have less-than or equal-to the specified number of empty containers (containers without contents).  (optional)
    emptyContainersLt := int32(56) // int32 | Only return boxes that have less-than the specified number of empty containers (containers without contents).  (optional)
    ancestorStorageId := "ancestorStorageId_example" // string | ID of a location. Restricts results to those located in the specified storage.  (optional)
    storageContentsId := "storageContentsId_example" // string | ID of a batch, entity, or entity schema. Restricts results to those that hold containers with entities or batches associated with the specified ID.  (optional)
    storageContentsIds := "storageContentsIds_example" // string | Comma-separated list of IDs of batches or entities. Restricts results to those that hold containers with at least one of the specified batches, entities, or batches of the specified entities.  (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived boxes. Use \"ANY_ARCHIVED\" to filter for archived boxes regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    ids := "box_Jx8Zsphf,box_s9Zv7Jto,box_mFDuLwA6" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    barcodes := "10x10-BOX105,10x10-BOX115" // string | Comma-separated list of barcodes. Matches all of the provided barcodes, or returns a 400 error that includes a list of which barcodes are invalid.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of users IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BoxesApi.ListBoxes(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).SchemaId(schemaId).SchemaFields(schemaFields).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).EmptyPositions(emptyPositions).EmptyPositionsGte(emptyPositionsGte).EmptyPositionsGt(emptyPositionsGt).EmptyPositionsLte(emptyPositionsLte).EmptyPositionsLt(emptyPositionsLt).EmptyContainers(emptyContainers).EmptyContainersGte(emptyContainersGte).EmptyContainersGt(emptyContainersGt).EmptyContainersLte(emptyContainersLte).EmptyContainersLt(emptyContainersLt).AncestorStorageId(ancestorStorageId).StorageContentsId(storageContentsId).StorageContentsIds(storageContentsIds).ArchiveReason(archiveReason).Ids(ids).Barcodes(barcodes).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoxesApi.ListBoxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBoxes`: BoxesPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `BoxesApi.ListBoxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBoxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100.  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema.  | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of a box. Restricts results to those with the specified name. | 
 **nameIncludes** | **string** | Name substring of a box. Restricts results to those with names that include the provided substring.  | 
 **emptyPositions** | **int32** | Only return boxes that have the specified number of empty positions  | 
 **emptyPositionsGte** | **int32** | Only return boxes that have greater-than or equal-to the specified number of empty positions.  | 
 **emptyPositionsGt** | **int32** | Only return boxes that have greater-than the specified number of empty positions.  | 
 **emptyPositionsLte** | **int32** | Only return boxes that have less-than or equal-to the specified number of empty positions.  | 
 **emptyPositionsLt** | **int32** | Only return boxes that have less-than the specified number of empty positions.  | 
 **emptyContainers** | **int32** | Only return boxes that have the specified number of empty containers (containers without contents).  | 
 **emptyContainersGte** | **int32** | Only return boxes that have greater-than or equal-to the specified number of empty containers (containers without contents).  | 
 **emptyContainersGt** | **int32** | Only return boxes that have greater-than the specified number of empty containers (containers without contents).  | 
 **emptyContainersLte** | **int32** | Only return boxes that have less-than or equal-to the specified number of empty containers (containers without contents).  | 
 **emptyContainersLt** | **int32** | Only return boxes that have less-than the specified number of empty containers (containers without contents).  | 
 **ancestorStorageId** | **string** | ID of a location. Restricts results to those located in the specified storage.  | 
 **storageContentsId** | **string** | ID of a batch, entity, or entity schema. Restricts results to those that hold containers with entities or batches associated with the specified ID.  | 
 **storageContentsIds** | **string** | Comma-separated list of IDs of batches or entities. Restricts results to those that hold containers with at least one of the specified batches, entities, or batches of the specified entities.  | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived boxes. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived boxes regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **barcodes** | **string** | Comma-separated list of barcodes. Matches all of the provided barcodes, or returns a 400 error that includes a list of which barcodes are invalid.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  | 
 **creatorIds** | **string** | Comma separated list of users IDs | 

### Return type

[**BoxesPaginatedList**](BoxesPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveBoxes

> BoxesArchivalChange UnarchiveBoxes(ctx).BoxesUnarchive(boxesUnarchive).Execute()

Unarchive boxes



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
    boxesUnarchive := *openapiclient.NewBoxesUnarchive([]string{"BoxIds_example"}) // BoxesUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BoxesApi.UnarchiveBoxes(context.Background()).BoxesUnarchive(boxesUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoxesApi.UnarchiveBoxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveBoxes`: BoxesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `BoxesApi.UnarchiveBoxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveBoxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **boxesUnarchive** | [**BoxesUnarchive**](BoxesUnarchive.md) |  | 

### Return type

[**BoxesArchivalChange**](BoxesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBox

> Box UpdateBox(ctx, boxId).BoxUpdate(boxUpdate).Execute()

Update a box



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
    boxId := "boxId_example" // string | 
    boxUpdate := *openapiclient.NewBoxUpdate() // BoxUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BoxesApi.UpdateBox(context.Background(), boxId).BoxUpdate(boxUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoxesApi.UpdateBox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBox`: Box
    fmt.Fprintf(os.Stdout, "Response from `BoxesApi.UpdateBox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBoxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **boxUpdate** | [**BoxUpdate**](BoxUpdate.md) |  | 

### Return type

[**Box**](Box.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

