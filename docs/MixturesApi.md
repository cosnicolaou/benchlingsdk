# \MixturesApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveMixtures**](MixturesApi.md#ArchiveMixtures) | **Post** /mixtures:archive | Archive mixtures
[**BulkCreateMixtures**](MixturesApi.md#BulkCreateMixtures) | **Post** /mixtures:bulk-create | Bulk Create mixtures
[**BulkUpdateMixtures**](MixturesApi.md#BulkUpdateMixtures) | **Post** /mixtures:bulk-update | Bulk Update mixtures
[**CreateMixture**](MixturesApi.md#CreateMixture) | **Post** /mixtures | Create a mixture
[**GetMixture**](MixturesApi.md#GetMixture) | **Get** /mixtures/{mixture_id} | Get a mixture
[**ListMixtures**](MixturesApi.md#ListMixtures) | **Get** /mixtures | List mixtures
[**UnarchiveMixtures**](MixturesApi.md#UnarchiveMixtures) | **Post** /mixtures:unarchive | Unarchive mixtures
[**UpdateMixture**](MixturesApi.md#UpdateMixture) | **Patch** /mixtures/{mixture_id} | Update a mixture



## ArchiveMixtures

> MixturesArchivalChange ArchiveMixtures(ctx).MixturesArchive(mixturesArchive).Execute()

Archive mixtures



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
    mixturesArchive := *openapiclient.NewMixturesArchive([]string{"mxt_djw9aU"}, openapiclient.EntityArchiveReason("Made in error")) // MixturesArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MixturesApi.ArchiveMixtures(context.Background()).MixturesArchive(mixturesArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MixturesApi.ArchiveMixtures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveMixtures`: MixturesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `MixturesApi.ArchiveMixtures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveMixturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mixturesArchive** | [**MixturesArchive**](MixturesArchive.md) |  | 

### Return type

[**MixturesArchivalChange**](MixturesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkCreateMixtures

> AsyncTaskLink BulkCreateMixtures(ctx).MixturesBulkCreateRequest(mixturesBulkCreateRequest).Execute()

Bulk Create mixtures



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
    mixturesBulkCreateRequest := *openapiclient.NewMixturesBulkCreateRequest([]openapiclient.MixtureCreate{*openapiclient.NewMixtureCreate([]openapiclient.IngredientWriteParams{*openapiclient.NewIngredientWriteParams("12", "DION_004", "bfi_37hdg8", "cnt_12345", "bfi_12345", "DION_004-source_001", "Notes_example", "TODO")}, "Name_example", "SchemaId_example", "TODO")}) // MixturesBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MixturesApi.BulkCreateMixtures(context.Background()).MixturesBulkCreateRequest(mixturesBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MixturesApi.BulkCreateMixtures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateMixtures`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `MixturesApi.BulkCreateMixtures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateMixturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mixturesBulkCreateRequest** | [**MixturesBulkCreateRequest**](MixturesBulkCreateRequest.md) |  | 

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


## BulkUpdateMixtures

> AsyncTaskLink BulkUpdateMixtures(ctx).MixturesBulkUpdateRequest(mixturesBulkUpdateRequest).Execute()

Bulk Update mixtures



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
    mixturesBulkUpdateRequest := *openapiclient.NewMixturesBulkUpdateRequest() // MixturesBulkUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MixturesApi.BulkUpdateMixtures(context.Background()).MixturesBulkUpdateRequest(mixturesBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MixturesApi.BulkUpdateMixtures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateMixtures`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `MixturesApi.BulkUpdateMixtures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateMixturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mixturesBulkUpdateRequest** | [**MixturesBulkUpdateRequest**](MixturesBulkUpdateRequest.md) |  | 

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


## CreateMixture

> Mixture CreateMixture(ctx).MixtureCreate(mixtureCreate).Execute()

Create a mixture



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
    mixtureCreate := *openapiclient.NewMixtureCreate([]openapiclient.IngredientWriteParams{*openapiclient.NewIngredientWriteParams("12", "DION_004", "bfi_37hdg8", "cnt_12345", "bfi_12345", "DION_004-source_001", "Notes_example", "TODO")}, "Name_example", "SchemaId_example", "TODO") // MixtureCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MixturesApi.CreateMixture(context.Background()).MixtureCreate(mixtureCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MixturesApi.CreateMixture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMixture`: Mixture
    fmt.Fprintf(os.Stdout, "Response from `MixturesApi.CreateMixture`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMixtureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mixtureCreate** | [**MixtureCreate**](MixtureCreate.md) |  | 

### Return type

[**Mixture**](Mixture.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMixture

> Mixture GetMixture(ctx, mixtureId).Execute()

Get a mixture



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
    mixtureId := "mixtureId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MixturesApi.GetMixture(context.Background(), mixtureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MixturesApi.GetMixture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMixture`: Mixture
    fmt.Fprintf(os.Stdout, "Response from `MixturesApi.GetMixture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mixtureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMixtureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Mixture**](Mixture.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMixtures

> MixturesPaginatedList ListMixtures(ctx).NextToken(nextToken).PageSize(pageSize).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).IngredientComponentEntityIds(ingredientComponentEntityIds).IngredientComponentEntityIdsAnyOf(ingredientComponentEntityIdsAnyOf).AuthorIdsAnyOf(authorIdsAnyOf).Execute()

List mixtures



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
    name := "name_example" // string | Name of a mixture. Restricts results to those with the specified name, alias, or entity registry ID. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of a mixture. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring.  (optional)
    folderId := "folderId_example" // string | ID of a folder. Restricts results to those in the folder. (optional)
    mentionedIn := "mentionedIn_example" // string | Comma-separated list of entry IDs. Restricts results to mixtures mentioned in those entries.  (optional)
    projectId := "projectId_example" // string | ID of a project. Restricts results to those in the project. (optional)
    registryId := "registryId_example" // string | ID of a registry. Restricts results to those registered in this registry. Specifying \"null\" returns unregistered items.  (optional)
    schemaId := "schemaId_example" // string | ID of a schema. Restricts results to those of the specified schema.  (optional)
    schemaFields := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Schema field value. For Integer, Float, and Date type fields, supports the >= and <= operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  (optional)
    archiveReason := "NOT_ARCHIVED" // string | Archive reason. Restricts items to those with the specified archive reason. Use \"NOT_ARCHIVED\" to filter for unarchived mixtures. Use \"ANY_ARCHIVED\" to filter for archived mixtures regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    mentions := "mentions_example" // string | Comma-separated list of resource IDs. Restricts results to those that mention the given items in the description.  (optional)
    ids := "bfi_blhxTUl1,bfi_y5bkDmJp,bfi_xwfILBog" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  (optional)
    entityRegistryIdsAnyOf := "TP001,TP002" // string | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  (optional)
    ingredientComponentEntityIds := "bfi_blhxTUl1,bfi_y5bkDmJp,bfi_xwfILBog" // string | Comma-separated list of ingredient component entity ids. Matches all mixtures that contain ingredients whose component entities match all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    ingredientComponentEntityIdsAnyOf := "bfi_blhxTUl1,bfi_y5bkDmJp,bfi_xwfILBog" // string | Comma-separated list of ingredient component entity ids. Matches all mixtures that contain ingredients whose component entities match any of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    authorIdsAnyOf := "ent_a0SApq3z,ent_b4AApz9b" // string | Comma separated list of user or app IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MixturesApi.ListMixtures(context.Background()).NextToken(nextToken).PageSize(pageSize).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).IngredientComponentEntityIds(ingredientComponentEntityIds).IngredientComponentEntityIdsAnyOf(ingredientComponentEntityIdsAnyOf).AuthorIdsAnyOf(authorIdsAnyOf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MixturesApi.ListMixtures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMixtures`: MixturesPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `MixturesApi.ListMixtures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMixturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 
 **pageSize** | **int32** |  | [default to 50]
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of a mixture. Restricts results to those with the specified name, alias, or entity registry ID. | 
 **nameIncludes** | **string** | Name substring of a mixture. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring.  | 
 **folderId** | **string** | ID of a folder. Restricts results to those in the folder. | 
 **mentionedIn** | **string** | Comma-separated list of entry IDs. Restricts results to mixtures mentioned in those entries.  | 
 **projectId** | **string** | ID of a project. Restricts results to those in the project. | 
 **registryId** | **string** | ID of a registry. Restricts results to those registered in this registry. Specifying \&quot;null\&quot; returns unregistered items.  | 
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema.  | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **archiveReason** | **string** | Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived mixtures. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived mixtures regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **mentions** | **string** | Comma-separated list of resource IDs. Restricts results to those that mention the given items in the description.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive.  | 
 **entityRegistryIdsAnyOf** | **string** | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  | 
 **ingredientComponentEntityIds** | **string** | Comma-separated list of ingredient component entity ids. Matches all mixtures that contain ingredients whose component entities match all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **ingredientComponentEntityIdsAnyOf** | **string** | Comma-separated list of ingredient component entity ids. Matches all mixtures that contain ingredients whose component entities match any of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **authorIdsAnyOf** | **string** | Comma separated list of user or app IDs | 

### Return type

[**MixturesPaginatedList**](MixturesPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveMixtures

> MixturesArchivalChange UnarchiveMixtures(ctx).MixturesUnarchive(mixturesUnarchive).Execute()

Unarchive mixtures



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
    mixturesUnarchive := *openapiclient.NewMixturesUnarchive([]string{"mxt_djw9aU"}) // MixturesUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MixturesApi.UnarchiveMixtures(context.Background()).MixturesUnarchive(mixturesUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MixturesApi.UnarchiveMixtures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveMixtures`: MixturesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `MixturesApi.UnarchiveMixtures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveMixturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mixturesUnarchive** | [**MixturesUnarchive**](MixturesUnarchive.md) |  | 

### Return type

[**MixturesArchivalChange**](MixturesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMixture

> Mixture UpdateMixture(ctx, mixtureId).MixtureUpdate(mixtureUpdate).Execute()

Update a mixture



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
    mixtureId := "mixtureId_example" // string | 
    mixtureUpdate := *openapiclient.NewMixtureUpdate() // MixtureUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MixturesApi.UpdateMixture(context.Background(), mixtureId).MixtureUpdate(mixtureUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MixturesApi.UpdateMixture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMixture`: Mixture
    fmt.Fprintf(os.Stdout, "Response from `MixturesApi.UpdateMixture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mixtureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMixtureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mixtureUpdate** | [**MixtureUpdate**](MixtureUpdate.md) |  | 

### Return type

[**Mixture**](Mixture.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

