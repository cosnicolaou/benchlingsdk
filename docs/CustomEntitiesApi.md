# \CustomEntitiesApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveCustomEntities**](CustomEntitiesApi.md#ArchiveCustomEntities) | **Post** /custom-entities:archive | Archive custom entities
[**BulkCreateCustomEntities**](CustomEntitiesApi.md#BulkCreateCustomEntities) | **Post** /custom-entities:bulk-create | Bulk Create custom entities
[**BulkGetCustomEntities**](CustomEntitiesApi.md#BulkGetCustomEntities) | **Get** /custom-entities:bulk-get | Bulk get custom entities by ID
[**BulkUpdateCustomEntities**](CustomEntitiesApi.md#BulkUpdateCustomEntities) | **Post** /custom-entities:bulk-update | Bulk Update custom entities
[**CreateCustomEntity**](CustomEntitiesApi.md#CreateCustomEntity) | **Post** /custom-entities | Create a custom entity
[**GetCustomEntity**](CustomEntitiesApi.md#GetCustomEntity) | **Get** /custom-entities/{custom_entity_id} | Get a custom entity
[**ListCustomEntities**](CustomEntitiesApi.md#ListCustomEntities) | **Get** /custom-entities | List custom entities
[**UnarchiveCustomEntities**](CustomEntitiesApi.md#UnarchiveCustomEntities) | **Post** /custom-entities:unarchive | Unarchive custom entities
[**UpdateCustomEntity**](CustomEntitiesApi.md#UpdateCustomEntity) | **Patch** /custom-entities/{custom_entity_id} | Update a custom entity



## ArchiveCustomEntities

> CustomEntitiesArchivalChange ArchiveCustomEntities(ctx).CustomEntitiesArchive(customEntitiesArchive).Execute()

Archive custom entities



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
    customEntitiesArchive := *openapiclient.NewCustomEntitiesArchive([]string{"CustomEntityIds_example"}, openapiclient.EntityArchiveReason("Made in error")) // CustomEntitiesArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEntitiesApi.ArchiveCustomEntities(context.Background()).CustomEntitiesArchive(customEntitiesArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEntitiesApi.ArchiveCustomEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveCustomEntities`: CustomEntitiesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `CustomEntitiesApi.ArchiveCustomEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveCustomEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customEntitiesArchive** | [**CustomEntitiesArchive**](CustomEntitiesArchive.md) |  | 

### Return type

[**CustomEntitiesArchivalChange**](CustomEntitiesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkCreateCustomEntities

> AsyncTaskLink BulkCreateCustomEntities(ctx).CustomEntitiesBulkCreateRequest(customEntitiesBulkCreateRequest).Execute()

Bulk Create custom entities



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
    customEntitiesBulkCreateRequest := *openapiclient.NewCustomEntitiesBulkCreateRequest([]openapiclient.CustomEntityBulkCreate{*openapiclient.NewCustomEntityBulkCreate("Name_example", "SchemaId_example")}) // CustomEntitiesBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEntitiesApi.BulkCreateCustomEntities(context.Background()).CustomEntitiesBulkCreateRequest(customEntitiesBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEntitiesApi.BulkCreateCustomEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateCustomEntities`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `CustomEntitiesApi.BulkCreateCustomEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateCustomEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customEntitiesBulkCreateRequest** | [**CustomEntitiesBulkCreateRequest**](CustomEntitiesBulkCreateRequest.md) |  | 

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


## BulkGetCustomEntities

> CustomEntitiesList BulkGetCustomEntities(ctx).CustomEntityIds(customEntityIds).Execute()

Bulk get custom entities by ID



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
    customEntityIds := "customEntityIds_example" // string | Comma-separated list of IDs of custom entities to get. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEntitiesApi.BulkGetCustomEntities(context.Background()).CustomEntityIds(customEntityIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEntitiesApi.BulkGetCustomEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkGetCustomEntities`: CustomEntitiesList
    fmt.Fprintf(os.Stdout, "Response from `CustomEntitiesApi.BulkGetCustomEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetCustomEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customEntityIds** | **string** | Comma-separated list of IDs of custom entities to get.  | 

### Return type

[**CustomEntitiesList**](CustomEntitiesList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateCustomEntities

> AsyncTaskLink BulkUpdateCustomEntities(ctx).CustomEntitiesBulkUpdateRequest(customEntitiesBulkUpdateRequest).Execute()

Bulk Update custom entities



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
    customEntitiesBulkUpdateRequest := *openapiclient.NewCustomEntitiesBulkUpdateRequest([]openapiclient.CustomEntityBulkUpdate{*openapiclient.NewCustomEntityBulkUpdate("Id_example")}) // CustomEntitiesBulkUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEntitiesApi.BulkUpdateCustomEntities(context.Background()).CustomEntitiesBulkUpdateRequest(customEntitiesBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEntitiesApi.BulkUpdateCustomEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateCustomEntities`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `CustomEntitiesApi.BulkUpdateCustomEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateCustomEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customEntitiesBulkUpdateRequest** | [**CustomEntitiesBulkUpdateRequest**](CustomEntitiesBulkUpdateRequest.md) |  | 

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


## CreateCustomEntity

> CustomEntity CreateCustomEntity(ctx).CustomEntityCreate(customEntityCreate).Execute()

Create a custom entity



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
    customEntityCreate := *openapiclient.NewCustomEntityCreate("Name_example", "SchemaId_example") // CustomEntityCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEntitiesApi.CreateCustomEntity(context.Background()).CustomEntityCreate(customEntityCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEntitiesApi.CreateCustomEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomEntity`: CustomEntity
    fmt.Fprintf(os.Stdout, "Response from `CustomEntitiesApi.CreateCustomEntity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customEntityCreate** | [**CustomEntityCreate**](CustomEntityCreate.md) |  | 

### Return type

[**CustomEntity**](CustomEntity.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomEntity

> CustomEntity GetCustomEntity(ctx, customEntityId).Execute()

Get a custom entity



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
    customEntityId := "customEntityId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEntitiesApi.GetCustomEntity(context.Background(), customEntityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEntitiesApi.GetCustomEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomEntity`: CustomEntity
    fmt.Fprintf(os.Stdout, "Response from `CustomEntitiesApi.GetCustomEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customEntityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomEntity**](CustomEntity.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomEntities

> CustomEntitiesPaginatedList ListCustomEntities(ctx).NextToken(nextToken).PageSize(pageSize).Sort(sort).ModifiedAt(modifiedAt).Name(name).Returning(returning).NameIncludes(nameIncludes).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).CreatorIds(creatorIds).AuthorIdsAnyOf(authorIdsAnyOf).Execute()

List custom entities



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
    sort := "sort_example" // string |  (optional) (default to "modifiedAt:desc")
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    name := "name_example" // string | Name of a custom entity. Restricts results to those with the specified name, alias, or entity registry ID. (optional)
    returning := "customEntities.id,customEntities.modifiedAt" // string | Comma-separated list of fields to return. Modifies the output shape. To return all keys at a given level, enumerate them or use the wildcard, '*'. For more information, [click here](https://docs.benchling.com/docs/getting-started-1#returning-query-parameter). (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of a custom entity. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring.  (optional)
    folderId := "folderId_example" // string | ID of a folder. Restricts results to those in the folder. (optional)
    mentionedIn := "mentionedIn_example" // string | Comma-separated list of entry IDs. Restricts results to custom entities mentioned in those entries.  (optional)
    projectId := "projectId_example" // string | ID of a project. Restricts results to those in the project. (optional)
    registryId := "registryId_example" // string | ID of a registry. Restricts results to those registered in this registry. Specifying \"null\" returns unregistered items.  (optional)
    schemaId := "schemaId_example" // string | ID of a schema. Restricts results to those of the specified schema.  (optional)
    schemaFields := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Schema field value. For Integer, Float, and Date type fields, supports the >= and <= operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived custom entities. Use \"ANY_ARCHIVED\" to filter for archived custom entities regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    mentions := "mentions_example" // string | Comma-separated list of resource IDs. Restricts results to those that mention the given items in the description.  (optional)
    ids := "bfi_blhxTUl1,bfi_y5bkDmJp,bfi_xwfILBog" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  (optional)
    entityRegistryIdsAnyOf := "TP001,TP002" // string | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs  (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of users IDs (optional)
    authorIdsAnyOf := "ent_a0SApq3z,ent_b4AApz9b" // string | Comma separated list of user or app IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEntitiesApi.ListCustomEntities(context.Background()).NextToken(nextToken).PageSize(pageSize).Sort(sort).ModifiedAt(modifiedAt).Name(name).Returning(returning).NameIncludes(nameIncludes).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).CreatorIds(creatorIds).AuthorIdsAnyOf(authorIdsAnyOf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEntitiesApi.ListCustomEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCustomEntities`: CustomEntitiesPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `CustomEntitiesApi.ListCustomEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of a custom entity. Restricts results to those with the specified name, alias, or entity registry ID. | 
 **returning** | **string** | Comma-separated list of fields to return. Modifies the output shape. To return all keys at a given level, enumerate them or use the wildcard, &#39;*&#39;. For more information, [click here](https://docs.benchling.com/docs/getting-started-1#returning-query-parameter). | 
 **nameIncludes** | **string** | Name substring of a custom entity. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring.  | 
 **folderId** | **string** | ID of a folder. Restricts results to those in the folder. | 
 **mentionedIn** | **string** | Comma-separated list of entry IDs. Restricts results to custom entities mentioned in those entries.  | 
 **projectId** | **string** | ID of a project. Restricts results to those in the project. | 
 **registryId** | **string** | ID of a registry. Restricts results to those registered in this registry. Specifying \&quot;null\&quot; returns unregistered items.  | 
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema.  | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived custom entities. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived custom entities regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **mentions** | **string** | Comma-separated list of resource IDs. Restricts results to those that mention the given items in the description.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  | 
 **entityRegistryIdsAnyOf** | **string** | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs  | 
 **creatorIds** | **string** | Comma separated list of users IDs | 
 **authorIdsAnyOf** | **string** | Comma separated list of user or app IDs | 

### Return type

[**CustomEntitiesPaginatedList**](CustomEntitiesPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveCustomEntities

> CustomEntitiesArchivalChange UnarchiveCustomEntities(ctx).CustomEntitiesUnarchive(customEntitiesUnarchive).Execute()

Unarchive custom entities



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
    customEntitiesUnarchive := *openapiclient.NewCustomEntitiesUnarchive([]string{"CustomEntityIds_example"}) // CustomEntitiesUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEntitiesApi.UnarchiveCustomEntities(context.Background()).CustomEntitiesUnarchive(customEntitiesUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEntitiesApi.UnarchiveCustomEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveCustomEntities`: CustomEntitiesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `CustomEntitiesApi.UnarchiveCustomEntities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveCustomEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customEntitiesUnarchive** | [**CustomEntitiesUnarchive**](CustomEntitiesUnarchive.md) |  | 

### Return type

[**CustomEntitiesArchivalChange**](CustomEntitiesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomEntity

> CustomEntity UpdateCustomEntity(ctx, customEntityId).CustomEntityUpdate(customEntityUpdate).Execute()

Update a custom entity



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
    customEntityId := "customEntityId_example" // string | 
    customEntityUpdate := *openapiclient.NewCustomEntityUpdate() // CustomEntityUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEntitiesApi.UpdateCustomEntity(context.Background(), customEntityId).CustomEntityUpdate(customEntityUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEntitiesApi.UpdateCustomEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomEntity`: CustomEntity
    fmt.Fprintf(os.Stdout, "Response from `CustomEntitiesApi.UpdateCustomEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customEntityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customEntityUpdate** | [**CustomEntityUpdate**](CustomEntityUpdate.md) |  | 

### Return type

[**CustomEntity**](CustomEntity.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

