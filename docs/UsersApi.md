# \UsersApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUser**](UsersApi.md#GetUser) | **Get** /users/{user_id} | Get a user by ID
[**ListUsers**](UsersApi.md#ListUsers) | **Get** /users | List users



## GetUser

> User GetUser(ctx, userId).Execute()

Get a user by ID



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
    userId := "userId_example" // string | ID of user to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of user to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> UsersPaginatedList ListUsers(ctx).Ids(ids).Name(name).NameIncludes(nameIncludes).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).ModifiedAt(modifiedAt).MemberOf(memberOf).AdminOf(adminOf).Handles(handles).PasswordLastChangedAt(passwordLastChangedAt).PageSize(pageSize).NextToken(nextToken).Sort(sort).Execute()

List users



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
    name := "name_example" // string | Name of a user. Restricts results to those with the specified name. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of a user. Restricts results to those with names that include the provided substring. (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  (optional)
    namesAnyOfCaseSensitive := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  (optional)
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    memberOf := "memberOf_example" // string | Comma-separated list of organization and/or team API IDs. Restricts results to users that are members of all given groups. (optional)
    adminOf := "adminOf_example" // string | Comma-separated list of organization and/or team API IDs. Restricts results to users that are admins of all given groups. (optional)
    handles := "handles_example" // string | Comma-separated list of handles. Restricts results to the users with the specified handles. (optional)
    passwordLastChangedAt := "passwordLastChangedAt_example" // string | Datetime, in RFC 3339 format. Supports the >, >=, <, <=, operators. Time zone defaults to UTC. Restricts results to users who have last changed their password in the specified range. e.g. > 2017-04-30. If \"null\" is provided returns users that have no password set (SAML).  (optional)
    pageSize := int32(56) // int32 |  (optional) (default to 50)
    nextToken := "nextToken_example" // string | Token for pagination (optional)
    sort := "sort_example" // string |  (optional) (default to "modifiedAt:desc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.ListUsers(context.Background()).Ids(ids).Name(name).NameIncludes(nameIncludes).NamesAnyOf(namesAnyOf).NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive).ModifiedAt(modifiedAt).MemberOf(memberOf).AdminOf(adminOf).Handles(handles).PasswordLastChangedAt(passwordLastChangedAt).PageSize(pageSize).NextToken(nextToken).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: UsersPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **name** | **string** | Name of a user. Restricts results to those with the specified name. | 
 **nameIncludes** | **string** | Name substring of a user. Restricts results to those with names that include the provided substring. | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity.  | 
 **namesAnyOfCaseSensitive** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive.  | 
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **memberOf** | **string** | Comma-separated list of organization and/or team API IDs. Restricts results to users that are members of all given groups. | 
 **adminOf** | **string** | Comma-separated list of organization and/or team API IDs. Restricts results to users that are admins of all given groups. | 
 **handles** | **string** | Comma-separated list of handles. Restricts results to the users with the specified handles. | 
 **passwordLastChangedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt;, &gt;&#x3D;, &lt;, &lt;&#x3D;, operators. Time zone defaults to UTC. Restricts results to users who have last changed their password in the specified range. e.g. &gt; 2017-04-30. If \&quot;null\&quot; is provided returns users that have no password set (SAML).  | 
 **pageSize** | **int32** |  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]

### Return type

[**UsersPaginatedList**](UsersPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

