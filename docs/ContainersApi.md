# \ContainersApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveContainers**](ContainersApi.md#ArchiveContainers) | **Post** /containers:archive | Archive containers
[**BulkCreateContainers**](ContainersApi.md#BulkCreateContainers) | **Post** /containers:bulk-create | Bulk create containers. Limit of 1000 containers per request.
[**BulkGetContainers**](ContainersApi.md#BulkGetContainers) | **Get** /containers:bulk-get | Bulk get a set of containers
[**BulkUpdateContainers**](ContainersApi.md#BulkUpdateContainers) | **Post** /containers:bulk-update | Bulk update containers
[**CheckinContainers**](ContainersApi.md#CheckinContainers) | **Post** /containers:check-in | Check in containers
[**CheckoutContainers**](ContainersApi.md#CheckoutContainers) | **Post** /containers:check-out | Check out containers
[**CreateContainer**](ContainersApi.md#CreateContainer) | **Post** /containers | Create a new container
[**DeleteContainerContent**](ContainersApi.md#DeleteContainerContent) | **Delete** /containers/{container_id}/contents/{containable_id} | Delete a container content
[**GetContainer**](ContainersApi.md#GetContainer) | **Get** /containers/{container_id} | get a container by id
[**GetContainerContent**](ContainersApi.md#GetContainerContent) | **Get** /containers/{container_id}/contents/{containable_id} | Get a container content
[**ListContainerContents**](ContainersApi.md#ListContainerContents) | **Get** /containers/{container_id}/contents | List a container&#39;s contents
[**ListContainers**](ContainersApi.md#ListContainers) | **Get** /containers | List containers
[**PrintLabels**](ContainersApi.md#PrintLabels) | **Post** /containers:print-labels | Print labels
[**ReserveContainers**](ContainersApi.md#ReserveContainers) | **Post** /containers:reserve | Reserve containers
[**TransferIntoContainer**](ContainersApi.md#TransferIntoContainer) | **Post** /containers/{destination_container_id}:transfer | Transfer into container
[**TransferIntoContainers**](ContainersApi.md#TransferIntoContainers) | **Post** /transfers | Transfers into containers
[**UnarchiveContainers**](ContainersApi.md#UnarchiveContainers) | **Post** /containers:unarchive | Unarchive containers
[**UpdateContainer**](ContainersApi.md#UpdateContainer) | **Patch** /containers/{container_id} | update a container
[**UpdateContainerContent**](ContainersApi.md#UpdateContainerContent) | **Patch** /containers/{container_id}/contents/{containable_id} | Update a container content



## ArchiveContainers

> ContainersArchivalChange ArchiveContainers(ctx).ContainersArchive(containersArchive).Execute()

Archive containers



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
    containersArchive := *openapiclient.NewContainersArchive([]string{"ContainerIds_example"}, "Reason_example") // ContainersArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.ArchiveContainers(context.Background()).ContainersArchive(containersArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.ArchiveContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveContainers`: ContainersArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.ArchiveContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containersArchive** | [**ContainersArchive**](ContainersArchive.md) |  | 

### Return type

[**ContainersArchivalChange**](ContainersArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkCreateContainers

> AsyncTaskLink BulkCreateContainers(ctx).ContainersBulkCreateRequest(containersBulkCreateRequest).Execute()

Bulk create containers. Limit of 1000 containers per request.



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
    containersBulkCreateRequest := *openapiclient.NewContainersBulkCreateRequest([]openapiclient.ContainerCreate{*openapiclient.NewContainerCreate("SchemaId_example")}) // ContainersBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.BulkCreateContainers(context.Background()).ContainersBulkCreateRequest(containersBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.BulkCreateContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateContainers`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.BulkCreateContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containersBulkCreateRequest** | [**ContainersBulkCreateRequest**](ContainersBulkCreateRequest.md) |  | 

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


## BulkGetContainers

> ContainersList BulkGetContainers(ctx).ContainerIds(containerIds).Barcodes(barcodes).Execute()

Bulk get a set of containers



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
    containerIds := "containerIds_example" // string | Comma-separated list of container IDs.  (optional)
    barcodes := "W102477,W102478" // string | Comma-separated list of barcodes. Matches all of the provided barcodes, or returns a 400 error that includes a list of which barcodes are invalid.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.BulkGetContainers(context.Background()).ContainerIds(containerIds).Barcodes(barcodes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.BulkGetContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetContainers`: ContainersList
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.BulkGetContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerIds** | **string** | Comma-separated list of container IDs.  | 
 **barcodes** | **string** | Comma-separated list of barcodes. Matches all of the provided barcodes, or returns a 400 error that includes a list of which barcodes are invalid.  | 

### Return type

[**ContainersList**](ContainersList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateContainers

> AsyncTaskLink BulkUpdateContainers(ctx).ContainersBulkUpdateRequest(containersBulkUpdateRequest).Execute()

Bulk update containers



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
    containersBulkUpdateRequest := *openapiclient.NewContainersBulkUpdateRequest([]openapiclient.ContainerBulkUpdateItem{*openapiclient.NewContainerBulkUpdateItem("ContainerId_example")}) // ContainersBulkUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.BulkUpdateContainers(context.Background()).ContainersBulkUpdateRequest(containersBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.BulkUpdateContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateContainers`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.BulkUpdateContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containersBulkUpdateRequest** | [**ContainersBulkUpdateRequest**](ContainersBulkUpdateRequest.md) |  | 

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


## CheckinContainers

> map[string]interface{} CheckinContainers(ctx).ContainersCheckin(containersCheckin).Execute()

Check in containers



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
    containersCheckin := *openapiclient.NewContainersCheckin([]string{"ContainerIds_example"}) // ContainersCheckin |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.CheckinContainers(context.Background()).ContainersCheckin(containersCheckin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.CheckinContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckinContainers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.CheckinContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckinContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containersCheckin** | [**ContainersCheckin**](ContainersCheckin.md) |  | 

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


## CheckoutContainers

> map[string]interface{} CheckoutContainers(ctx).ContainersCheckout(containersCheckout).Execute()

Check out containers



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
    containersCheckout := *openapiclient.NewContainersCheckout("AssigneeId_example", []string{"ContainerIds_example"}) // ContainersCheckout |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.CheckoutContainers(context.Background()).ContainersCheckout(containersCheckout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.CheckoutContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckoutContainers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.CheckoutContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckoutContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containersCheckout** | [**ContainersCheckout**](ContainersCheckout.md) |  | 

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


## CreateContainer

> Container CreateContainer(ctx).ContainerCreate(containerCreate).Execute()

Create a new container



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
    containerCreate := *openapiclient.NewContainerCreate("SchemaId_example") // ContainerCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.CreateContainer(context.Background()).ContainerCreate(containerCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.CreateContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContainer`: Container
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.CreateContainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerCreate** | [**ContainerCreate**](ContainerCreate.md) |  | 

### Return type

[**Container**](Container.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainerContent

> DeleteContainerContent(ctx, containerId, containableId).Execute()

Delete a container content



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
    containerId := "containerId_example" // string | 
    containableId := "containableId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.DeleteContainerContent(context.Background(), containerId, containableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.DeleteContainerContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** |  | 
**containableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainer

> Container GetContainer(ctx, containerId).Execute()

get a container by id



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
    containerId := "containerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.GetContainer(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.GetContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainer`: Container
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.GetContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Container**](Container.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerContent

> ContainerContent GetContainerContent(ctx, containerId, containableId).Execute()

Get a container content



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
    containerId := "containerId_example" // string | 
    containableId := "containableId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.GetContainerContent(context.Background(), containerId, containableId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.GetContainerContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerContent`: ContainerContent
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.GetContainerContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** |  | 
**containableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContainerContent**](ContainerContent.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerContents

> ContainerContentsList ListContainerContents(ctx, containerId).Execute()

List a container's contents



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
    containerId := "containerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.ListContainerContents(context.Background(), containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.ListContainerContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainerContents`: ContainerContentsList
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.ListContainerContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerContentsList**](ContainerContentsList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainers

> ContainersPaginatedList ListContainers(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).SchemaId(schemaId).SchemaFields(schemaFields).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).AncestorStorageId(ancestorStorageId).StorageContentsId(storageContentsId).StorageContentsIds(storageContentsIds).ArchiveReason(archiveReason).CheckoutStatus(checkoutStatus).Ids(ids).Barcodes(barcodes).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).Execute()

List containers



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
    sort := "sort_example" // string | Method by which to order search results. Valid sorts are barcode (bar code, alphabetical) modifiedAt (modified time, most recent first) and name (entity name, alphabetical). Optionally add :asc or :desc to specify ascending or descending order. Default is modifiedAt.  (optional) (default to "modifiedAt")
    schemaId := "schemaId_example" // string | ID of a schema. Restricts results to those of the specified schema. (optional)
    schemaFields := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Schema field value. For Integer, Float, and Date type fields, supports the >= and <= operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  (optional)
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    name := "name_example" // string | Name of a container. Restricts results to those with the specified name. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of a container. Restricts results to those with names that include the provided substring.  (optional)
    ancestorStorageId := "ancestorStorageId_example" // string | ID of a plate, box, or location. Restricts results to those located in the specified storage. (optional)
    storageContentsId := "storageContentsId_example" // string | ID of a batch, entity, or entity schema. Restricts results to those that contain the specified batches, batches of the specified entities, or batches of entities of the specified schema.  (optional)
    storageContentsIds := "storageContentsIds_example" // string | Comma-separated list of IDs of batches or entities. Restricts results to those that hold containers with at least one of the specified batches, entities, or batches of the specified entities.  (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived containers. Use \"ANY_ARCHIVED\" to filter for archived containers regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    checkoutStatus := "checkoutStatus_example" // string | Comma-separated list of check-out statuses. Restricts results to those that match one of the specified statuses. Valid statuses are AVAILABLE, RESERVED, and CHECKED_OUT.  (optional)
    ids := "con_Q6uhNZvw,con_OwmERWGE,con_nzuDFhNvz" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    barcodes := "W103371,W103343,W103366" // string | Comma-separated list of barcodes. Matches all of the provided barcodes, or returns a 400 error that includes a list of which barcodes are invalid.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of users IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.ListContainers(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).SchemaId(schemaId).SchemaFields(schemaFields).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).AncestorStorageId(ancestorStorageId).StorageContentsId(storageContentsId).StorageContentsIds(storageContentsIds).ArchiveReason(archiveReason).CheckoutStatus(checkoutStatus).Ids(ids).Barcodes(barcodes).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.ListContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainers`: ContainersPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.ListContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** | Method by which to order search results. Valid sorts are barcode (bar code, alphabetical) modifiedAt (modified time, most recent first) and name (entity name, alphabetical). Optionally add :asc or :desc to specify ascending or descending order. Default is modifiedAt.  | [default to &quot;modifiedAt&quot;]
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema. | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of a container. Restricts results to those with the specified name. | 
 **nameIncludes** | **string** | Name substring of a container. Restricts results to those with names that include the provided substring.  | 
 **ancestorStorageId** | **string** | ID of a plate, box, or location. Restricts results to those located in the specified storage. | 
 **storageContentsId** | **string** | ID of a batch, entity, or entity schema. Restricts results to those that contain the specified batches, batches of the specified entities, or batches of entities of the specified schema.  | 
 **storageContentsIds** | **string** | Comma-separated list of IDs of batches or entities. Restricts results to those that hold containers with at least one of the specified batches, entities, or batches of the specified entities.  | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived containers. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived containers regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **checkoutStatus** | **string** | Comma-separated list of check-out statuses. Restricts results to those that match one of the specified statuses. Valid statuses are AVAILABLE, RESERVED, and CHECKED_OUT.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **barcodes** | **string** | Comma-separated list of barcodes. Matches all of the provided barcodes, or returns a 400 error that includes a list of which barcodes are invalid.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  | 
 **creatorIds** | **string** | Comma separated list of users IDs | 

### Return type

[**ContainersPaginatedList**](ContainersPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintLabels

> map[string]interface{} PrintLabels(ctx).PrintLabels(printLabels).Execute()

Print labels



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
    printLabels := *openapiclient.NewPrintLabels([]string{"ContainerIds_example"}, "lbltmp_lCpY5Qvh", "print_YJQkilOJ") // PrintLabels |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.PrintLabels(context.Background()).PrintLabels(printLabels).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.PrintLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrintLabels`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.PrintLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **printLabels** | [**PrintLabels**](PrintLabels.md) |  | 

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


## ReserveContainers

> map[string]interface{} ReserveContainers(ctx).ContainersCheckout(containersCheckout).Execute()

Reserve containers



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
    containersCheckout := *openapiclient.NewContainersCheckout("AssigneeId_example", []string{"ContainerIds_example"}) // ContainersCheckout |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.ReserveContainers(context.Background()).ContainersCheckout(containersCheckout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.ReserveContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReserveContainers`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.ReserveContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReserveContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containersCheckout** | [**ContainersCheckout**](ContainersCheckout.md) |  | 

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


## TransferIntoContainer

> map[string]interface{} TransferIntoContainer(ctx, destinationContainerId).ContainerTransfer(containerTransfer).Execute()

Transfer into container



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
    destinationContainerId := "destinationContainerId_example" // string | 
    containerTransfer := *openapiclient.NewContainerTransfer([]openapiclient.ContainerTransferDestinationContentsItem{*openapiclient.NewContainerTransferDestinationContentsItem("EntityId_example")}) // ContainerTransfer |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.TransferIntoContainer(context.Background(), destinationContainerId).ContainerTransfer(containerTransfer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.TransferIntoContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferIntoContainer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.TransferIntoContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationContainerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferIntoContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerTransfer** | [**ContainerTransfer**](ContainerTransfer.md) |  | 

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


## TransferIntoContainers

> AsyncTaskLink TransferIntoContainers(ctx).MultipleContainersTransfersList(multipleContainersTransfersList).Execute()

Transfers into containers



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
    multipleContainersTransfersList := *openapiclient.NewMultipleContainersTransfersList([]openapiclient.MultipleContainersTransfer{*openapiclient.NewMultipleContainersTransfer("DestinationContainerId_example")}) // MultipleContainersTransfersList |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.TransferIntoContainers(context.Background()).MultipleContainersTransfersList(multipleContainersTransfersList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.TransferIntoContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferIntoContainers`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.TransferIntoContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferIntoContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multipleContainersTransfersList** | [**MultipleContainersTransfersList**](MultipleContainersTransfersList.md) |  | 

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


## UnarchiveContainers

> ContainersArchivalChange UnarchiveContainers(ctx).ContainersUnarchive(containersUnarchive).Execute()

Unarchive containers



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
    containersUnarchive := *openapiclient.NewContainersUnarchive([]string{"ContainerIds_example"}) // ContainersUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.UnarchiveContainers(context.Background()).ContainersUnarchive(containersUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.UnarchiveContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveContainers`: ContainersArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.UnarchiveContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containersUnarchive** | [**ContainersUnarchive**](ContainersUnarchive.md) |  | 

### Return type

[**ContainersArchivalChange**](ContainersArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContainer

> Container UpdateContainer(ctx, containerId).ContainerUpdate(containerUpdate).Execute()

update a container



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
    containerId := "containerId_example" // string | 
    containerUpdate := *openapiclient.NewContainerUpdate() // ContainerUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.UpdateContainer(context.Background(), containerId).ContainerUpdate(containerUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.UpdateContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateContainer`: Container
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.UpdateContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **containerUpdate** | [**ContainerUpdate**](ContainerUpdate.md) |  | 

### Return type

[**Container**](Container.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContainerContent

> ContainerContent UpdateContainerContent(ctx, containerId, containableId).ContainerContentUpdate(containerContentUpdate).Execute()

Update a container content



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
    containerId := "containerId_example" // string | 
    containableId := "containableId_example" // string | 
    containerContentUpdate := *openapiclient.NewContainerContentUpdate(*openapiclient.NewMeasurement("Units_example", NullableFloat32(123))) // ContainerContentUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainersApi.UpdateContainerContent(context.Background(), containerId, containableId).ContainerContentUpdate(containerContentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainersApi.UpdateContainerContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateContainerContent`: ContainerContent
    fmt.Fprintf(os.Stdout, "Response from `ContainersApi.UpdateContainerContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerId** | **string** |  | 
**containableId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContainerContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **containerContentUpdate** | [**ContainerContentUpdate**](ContainerContentUpdate.md) |  | 

### Return type

[**ContainerContent**](ContainerContent.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

