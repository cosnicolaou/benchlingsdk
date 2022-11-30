# \OrganizationsApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganization**](OrganizationsApi.md#GetOrganization) | **Get** /organizations/{organization_id} | Get an organization by ID
[**ListOrganizations**](OrganizationsApi.md#ListOrganizations) | **Get** /organizations | List organizations



## GetOrganization

> Organization GetOrganization(ctx, organizationId).Execute()

Get an organization by ID



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
    organizationId := "organizationId_example" // string | ID of organization to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganization(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganization`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | ID of organization to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Organization**](Organization.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrganizations

> OrganizationsPaginatedList ListOrganizations(ctx).Ids(ids).Name(name).NameIncludes(nameIncludes).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).ModifiedAt(modifiedAt).HasMembers(hasMembers).HasAdmins(hasAdmins).PageSize(pageSize).NextToken(nextToken).Sort(sort).Execute()

List organizations



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
    ids := "ent_ZJy8RTbo,ent_8GVbVkPj,ent_qREJ33rn" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    name := "name_example" // string | Name of an organization. Restricts results to those with the specified name. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of an organization. Restricts results to those with names that include the provided substring. (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  (optional)
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    hasMembers := "hasMembers_example" // string | Comma-separated list of user or Benchling app IDs. Restricts results to organizations that include all the given users/apps as members. (optional)
    hasAdmins := "hasAdmins_example" // string | Comma-separated list of user or Benchling app IDs. Restricts results to organizations that include all the given users/apps as admins. (optional)
    pageSize := int32(56) // int32 |  (optional) (default to 50)
    nextToken := "nextToken_example" // string | Token for pagination (optional)
    sort := "sort_example" // string |  (optional) (default to "modifiedAt:desc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.ListOrganizations(context.Background()).Ids(ids).Name(name).NameIncludes(nameIncludes).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).ModifiedAt(modifiedAt).HasMembers(hasMembers).HasAdmins(hasAdmins).PageSize(pageSize).NextToken(nextToken).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.ListOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrganizations`: OrganizationsPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.ListOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **name** | **string** | Name of an organization. Restricts results to those with the specified name. | 
 **nameIncludes** | **string** | Name substring of an organization. Restricts results to those with names that include the provided substring. | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  | 
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **hasMembers** | **string** | Comma-separated list of user or Benchling app IDs. Restricts results to organizations that include all the given users/apps as members. | 
 **hasAdmins** | **string** | Comma-separated list of user or Benchling app IDs. Restricts results to organizations that include all the given users/apps as admins. | 
 **pageSize** | **int32** |  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]

### Return type

[**OrganizationsPaginatedList**](OrganizationsPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

