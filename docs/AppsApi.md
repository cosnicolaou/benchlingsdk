# \AppsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveBenchlingApps**](AppsApi.md#ArchiveBenchlingApps) | **Post** /apps:archive | Archive apps
[**BulkCreateAppConfigurationItems**](AppsApi.md#BulkCreateAppConfigurationItems) | **Post** /app-configuration-items:bulk-create | Bulk Create app configuration items. Limit of 1000 App Config Items per request.
[**BulkUpdateAppConfigurationItems**](AppsApi.md#BulkUpdateAppConfigurationItems) | **Post** /app-configuration-items:bulk-update | Bulk Update app configuration items. Limit of 1000 App Config Items per request.
[**CreateAppConfigurationItem**](AppsApi.md#CreateAppConfigurationItem) | **Post** /app-configuration-items | Create app configuration item
[**CreateBenchlingApp**](AppsApi.md#CreateBenchlingApp) | **Post** /apps | Create an app
[**GetAppConfigurationItemById**](AppsApi.md#GetAppConfigurationItemById) | **Get** /app-configuration-items/{item_id} | Get app configuration item
[**GetBenchlingAppByID**](AppsApi.md#GetBenchlingAppByID) | **Get** /apps/{app_id} | Get an app by ID
[**ListAppConfigurationItems**](AppsApi.md#ListAppConfigurationItems) | **Get** /app-configuration-items | Get app configuration items
[**ListBenchlingApps**](AppsApi.md#ListBenchlingApps) | **Get** /apps | List apps
[**PatchBenchlingApp**](AppsApi.md#PatchBenchlingApp) | **Patch** /apps/{app_id} | Update an app
[**UnarchiveBenchlingApps**](AppsApi.md#UnarchiveBenchlingApps) | **Post** /apps:unarchive | Unarchive apps
[**UpdateAppConfigurationItem**](AppsApi.md#UpdateAppConfigurationItem) | **Patch** /app-configuration-items/{item_id} | Update app configuration item



## ArchiveBenchlingApps

> BenchlingAppsArchivalChange ArchiveBenchlingApps(ctx).BenchlingAppsArchive(benchlingAppsArchive).Execute()

Archive apps



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
    benchlingAppsArchive := *openapiclient.NewBenchlingAppsArchive([]string{"AppIds_example"}, "Reason_example") // BenchlingAppsArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.ArchiveBenchlingApps(context.Background()).BenchlingAppsArchive(benchlingAppsArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.ArchiveBenchlingApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveBenchlingApps`: BenchlingAppsArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.ArchiveBenchlingApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveBenchlingAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **benchlingAppsArchive** | [**BenchlingAppsArchive**](BenchlingAppsArchive.md) |  | 

### Return type

[**BenchlingAppsArchivalChange**](BenchlingAppsArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkCreateAppConfigurationItems

> AsyncTaskLink BulkCreateAppConfigurationItems(ctx).AppConfigItemsBulkCreateRequest(appConfigItemsBulkCreateRequest).Execute()

Bulk Create app configuration items. Limit of 1000 App Config Items per request.



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
    appConfigItemsBulkCreateRequest := *openapiclient.NewAppConfigItemsBulkCreateRequest([]openapiclient.AppConfigItemCreate{openapiclient.AppConfigItemCreate{AppConfigItemBooleanCreate: openapiclient.NewAppConfigItemBooleanCreate("app_J39na03L1nsLS34o", []string{"Path_example"}, "boolean", false)}}) // AppConfigItemsBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.BulkCreateAppConfigurationItems(context.Background()).AppConfigItemsBulkCreateRequest(appConfigItemsBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.BulkCreateAppConfigurationItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateAppConfigurationItems`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.BulkCreateAppConfigurationItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateAppConfigurationItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appConfigItemsBulkCreateRequest** | [**AppConfigItemsBulkCreateRequest**](AppConfigItemsBulkCreateRequest.md) |  | 

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


## BulkUpdateAppConfigurationItems

> AsyncTaskLink BulkUpdateAppConfigurationItems(ctx).AppConfigItemsBulkUpdateRequest(appConfigItemsBulkUpdateRequest).Execute()

Bulk Update app configuration items. Limit of 1000 App Config Items per request.



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
    appConfigItemsBulkUpdateRequest := *openapiclient.NewAppConfigItemsBulkUpdateRequest([]openapiclient.AppConfigItemBulkUpdate{openapiclient.AppConfigItemBulkUpdate{AppConfigItemBooleanBulkUpdate: openapiclient.NewAppConfigItemBooleanBulkUpdate("boolean", false, "aci_ae92kBv9aNSl593z")}}) // AppConfigItemsBulkUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.BulkUpdateAppConfigurationItems(context.Background()).AppConfigItemsBulkUpdateRequest(appConfigItemsBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.BulkUpdateAppConfigurationItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateAppConfigurationItems`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.BulkUpdateAppConfigurationItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateAppConfigurationItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appConfigItemsBulkUpdateRequest** | [**AppConfigItemsBulkUpdateRequest**](AppConfigItemsBulkUpdateRequest.md) |  | 

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


## CreateAppConfigurationItem

> AppConfigItem CreateAppConfigurationItem(ctx).AppConfigItemCreate(appConfigItemCreate).Execute()

Create app configuration item



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
    appConfigItemCreate := openapiclient.AppConfigItemCreate{AppConfigItemBooleanCreate: openapiclient.NewAppConfigItemBooleanCreate("app_J39na03L1nsLS34o", []string{"Path_example"}, "boolean", false)} // AppConfigItemCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.CreateAppConfigurationItem(context.Background()).AppConfigItemCreate(appConfigItemCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.CreateAppConfigurationItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAppConfigurationItem`: AppConfigItem
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.CreateAppConfigurationItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppConfigurationItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appConfigItemCreate** | [**AppConfigItemCreate**](AppConfigItemCreate.md) |  | 

### Return type

[**AppConfigItem**](AppConfigItem.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBenchlingApp

> BenchlingApp CreateBenchlingApp(ctx).BenchlingAppCreate(benchlingAppCreate).Execute()

Create an app



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
    benchlingAppCreate := *openapiclient.NewBenchlingAppCreate("My First App") // BenchlingAppCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.CreateBenchlingApp(context.Background()).BenchlingAppCreate(benchlingAppCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.CreateBenchlingApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBenchlingApp`: BenchlingApp
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.CreateBenchlingApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBenchlingAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **benchlingAppCreate** | [**BenchlingAppCreate**](BenchlingAppCreate.md) |  | 

### Return type

[**BenchlingApp**](BenchlingApp.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppConfigurationItemById

> AppConfigItem GetAppConfigurationItemById(ctx, itemId).Execute()

Get app configuration item



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
    itemId := "aci_e59sjL23Pqn30xHg" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.GetAppConfigurationItemById(context.Background(), itemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.GetAppConfigurationItemById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppConfigurationItemById`: AppConfigItem
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.GetAppConfigurationItemById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppConfigurationItemByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppConfigItem**](AppConfigItem.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBenchlingAppByID

> BenchlingApp GetBenchlingAppByID(ctx, appId).Execute()

Get an app by ID



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
    appId := "app_J39na03L1nsLS34o" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.GetBenchlingAppByID(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.GetBenchlingAppByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBenchlingAppByID`: BenchlingApp
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.GetBenchlingAppByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBenchlingAppByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BenchlingApp**](BenchlingApp.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppConfigurationItems

> AppConfigurationPaginatedList ListAppConfigurationItems(ctx).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).AppId(appId).Ids(ids).Sort(sort).Execute()

Get app configuration items



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
    nextToken := "nextToken_example" // string | Token for pagination (optional)
    pageSize := int32(56) // int32 | Number of results to return. Defaults to 50, maximum of 100. (optional) (default to 50)
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    appId := "app_e59sjL23Pqn30xHg" // string | App id to which the configuration belongs. (optional)
    ids := "aci_VfVOART1,aci_RFhDGaaC" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    sort := "sort_example" // string |  (optional) (default to "modifiedAt:desc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.ListAppConfigurationItems(context.Background()).NextToken(nextToken).PageSize(pageSize).ModifiedAt(modifiedAt).AppId(appId).Ids(ids).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.ListAppConfigurationItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAppConfigurationItems`: AppConfigurationPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.ListAppConfigurationItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppConfigurationItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** | Token for pagination | 
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100. | [default to 50]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **appId** | **string** | App id to which the configuration belongs. | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]

### Return type

[**AppConfigurationPaginatedList**](AppConfigurationPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBenchlingApps

> BenchlingAppsPaginatedList ListBenchlingApps(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).Ids(ids).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).MemberOf(memberOf).AdminOf(adminOf).Execute()

List apps



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
    sort := "sort_example" // string | Method by which to order search results. Valid sorts are modifiedAt (modified time, most recent first) and name (app name, alphabetical). Optionally add :asc or :desc to specify ascending or descending order. Default is modifiedAt.  (optional) (default to "modifiedAt")
    ids := "app_J39na03L1nsLS34o,app_ae92kBv9aNSl593z,app_e59sjL23Pqn30xHg" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    name := "name_example" // string | Name of an app. Restricts results to those with the specified name. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of an app. Restricts results to those with names that include the provided substring. (optional)
    namesAnyOf := "MyName1, MyName2" // string | Comma-separated list of app names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of app names. Restricts results to those that match any of the specified names, case sensitive.  (optional)
    creatorIds := "ent_a0SApq3z" // string | Comma separated list of users IDs (optional)
    memberOf := "memberOf_example" // string | Comma-separated list of organization and/or team API IDs. Restricts results to apps that are members of all given groups. (optional)
    adminOf := "adminOf_example" // string | Comma-separated list of organization and/or team API IDs. Restricts results to apps that are admins of all given groups. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.ListBenchlingApps(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).Ids(ids).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).CreatorIds(creatorIds).MemberOf(memberOf).AdminOf(adminOf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.ListBenchlingApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBenchlingApps`: BenchlingAppsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.ListBenchlingApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBenchlingAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** | Method by which to order search results. Valid sorts are modifiedAt (modified time, most recent first) and name (app name, alphabetical). Optionally add :asc or :desc to specify ascending or descending order. Default is modifiedAt.  | [default to &quot;modifiedAt&quot;]
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of an app. Restricts results to those with the specified name. | 
 **nameIncludes** | **string** | Name substring of an app. Restricts results to those with names that include the provided substring. | 
 **namesAnyOf** | **string** | Comma-separated list of app names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of app names. Restricts results to those that match any of the specified names, case sensitive.  | 
 **creatorIds** | **string** | Comma separated list of users IDs | 
 **memberOf** | **string** | Comma-separated list of organization and/or team API IDs. Restricts results to apps that are members of all given groups. | 
 **adminOf** | **string** | Comma-separated list of organization and/or team API IDs. Restricts results to apps that are admins of all given groups. | 

### Return type

[**BenchlingAppsPaginatedList**](BenchlingAppsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchBenchlingApp

> BenchlingApp PatchBenchlingApp(ctx, appId).BenchlingAppUpdate(benchlingAppUpdate).Execute()

Update an app



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
    appId := "app_e59sjL23Pqn30xHg" // string | 
    benchlingAppUpdate := *openapiclient.NewBenchlingAppUpdate() // BenchlingAppUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.PatchBenchlingApp(context.Background(), appId).BenchlingAppUpdate(benchlingAppUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.PatchBenchlingApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchBenchlingApp`: BenchlingApp
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.PatchBenchlingApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchBenchlingAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **benchlingAppUpdate** | [**BenchlingAppUpdate**](BenchlingAppUpdate.md) |  | 

### Return type

[**BenchlingApp**](BenchlingApp.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveBenchlingApps

> BenchlingAppsArchivalChange UnarchiveBenchlingApps(ctx).BenchlingAppsUnarchive(benchlingAppsUnarchive).Execute()

Unarchive apps



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
    benchlingAppsUnarchive := *openapiclient.NewBenchlingAppsUnarchive([]string{"AppIds_example"}) // BenchlingAppsUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.UnarchiveBenchlingApps(context.Background()).BenchlingAppsUnarchive(benchlingAppsUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.UnarchiveBenchlingApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveBenchlingApps`: BenchlingAppsArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.UnarchiveBenchlingApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveBenchlingAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **benchlingAppsUnarchive** | [**BenchlingAppsUnarchive**](BenchlingAppsUnarchive.md) |  | 

### Return type

[**BenchlingAppsArchivalChange**](BenchlingAppsArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppConfigurationItem

> AppConfigItem UpdateAppConfigurationItem(ctx, itemId).AppConfigItemUpdate(appConfigItemUpdate).Execute()

Update app configuration item



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
    itemId := "aci_e59sjL23Pqn30xHg" // string | 
    appConfigItemUpdate := openapiclient.AppConfigItemUpdate{AppConfigItemBooleanUpdate: openapiclient.NewAppConfigItemBooleanUpdate("boolean", false)} // AppConfigItemUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.UpdateAppConfigurationItem(context.Background(), itemId).AppConfigItemUpdate(appConfigItemUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.UpdateAppConfigurationItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppConfigurationItem`: AppConfigItem
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.UpdateAppConfigurationItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppConfigurationItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appConfigItemUpdate** | [**AppConfigItemUpdate**](AppConfigItemUpdate.md) |  | 

### Return type

[**AppConfigItem**](AppConfigItem.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

