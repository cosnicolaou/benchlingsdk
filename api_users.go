/*
Benchling API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package benchlingsdk

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// UsersApiService UsersApi service
type UsersApiService service

type ApiGetUserRequest struct {
	ctx context.Context
	ApiService *UsersApiService
	userId string
}

func (r ApiGetUserRequest) Execute() (*User, *http.Response, error) {
	return r.ApiService.GetUserExecute(r)
}

/*
GetUser Get a user by ID

Returns a user by ID if the caller has permission to view. The following roles have view permission:
  - tenant admins
  - members of any of the user's organizations


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ID of user to get
 @return ApiGetUserRequest
*/
func (a *UsersApiService) GetUser(ctx context.Context, userId string) ApiGetUserRequest {
	return ApiGetUserRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return User
func (a *UsersApiService) GetUserExecute(r ApiGetUserRequest) (*User, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *User
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.GetUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v NotFoundError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListUsersRequest struct {
	ctx context.Context
	ApiService *UsersApiService
	ids *string
	name *string
	nameIncludes *string
	namesAnyOf *string
	namesAnyOfCaseSensitive *string
	modifiedAt *string
	memberOf *string
	adminOf *string
	handles *string
	passwordLastChangedAt *string
	pageSize *int32
	nextToken *string
	sort *string
}

// Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid. 
func (r ApiListUsersRequest) Ids(ids string) ApiListUsersRequest {
	r.ids = &ids
	return r
}

// Name of a user. Restricts results to those with the specified name.
func (r ApiListUsersRequest) Name(name string) ApiListUsersRequest {
	r.name = &name
	return r
}

// Name substring of a user. Restricts results to those with names that include the provided substring.
func (r ApiListUsersRequest) NameIncludes(nameIncludes string) ApiListUsersRequest {
	r.nameIncludes = &nameIncludes
	return r
}

// Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity. 
func (r ApiListUsersRequest) NamesAnyOf(namesAnyOf string) ApiListUsersRequest {
	r.namesAnyOf = &namesAnyOf
	return r
}

// Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive. 
func (r ApiListUsersRequest) NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive string) ApiListUsersRequest {
	r.namesAnyOfCaseSensitive = &namesAnyOfCaseSensitive
	return r
}

// Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30. 
func (r ApiListUsersRequest) ModifiedAt(modifiedAt string) ApiListUsersRequest {
	r.modifiedAt = &modifiedAt
	return r
}

// Comma-separated list of organization and/or team API IDs. Restricts results to users that are members of all given groups.
func (r ApiListUsersRequest) MemberOf(memberOf string) ApiListUsersRequest {
	r.memberOf = &memberOf
	return r
}

// Comma-separated list of organization and/or team API IDs. Restricts results to users that are admins of all given groups.
func (r ApiListUsersRequest) AdminOf(adminOf string) ApiListUsersRequest {
	r.adminOf = &adminOf
	return r
}

// Comma-separated list of handles. Restricts results to the users with the specified handles.
func (r ApiListUsersRequest) Handles(handles string) ApiListUsersRequest {
	r.handles = &handles
	return r
}

// Datetime, in RFC 3339 format. Supports the &gt;, &gt;&#x3D;, &lt;, &lt;&#x3D;, operators. Time zone defaults to UTC. Restricts results to users who have last changed their password in the specified range. e.g. &gt; 2017-04-30. If \&quot;null\&quot; is provided returns users that have no password set (SAML). 
func (r ApiListUsersRequest) PasswordLastChangedAt(passwordLastChangedAt string) ApiListUsersRequest {
	r.passwordLastChangedAt = &passwordLastChangedAt
	return r
}

func (r ApiListUsersRequest) PageSize(pageSize int32) ApiListUsersRequest {
	r.pageSize = &pageSize
	return r
}

// Token for pagination
func (r ApiListUsersRequest) NextToken(nextToken string) ApiListUsersRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiListUsersRequest) Sort(sort string) ApiListUsersRequest {
	r.sort = &sort
	return r
}

func (r ApiListUsersRequest) Execute() (*UsersPaginatedList, *http.Response, error) {
	return r.ApiService.ListUsersExecute(r)
}

/*
ListUsers List users

Returns all users that the caller has permission to view. The following roles have view permission:
  - tenant admins
  - members of the user's organizations


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListUsersRequest
*/
func (a *UsersApiService) ListUsers(ctx context.Context) ApiListUsersRequest {
	return ApiListUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UsersPaginatedList
func (a *UsersApiService) ListUsersExecute(r ApiListUsersRequest) (*UsersPaginatedList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsersPaginatedList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.ListUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ids != nil {
		localVarQueryParams.Add("ids", parameterToString(*r.ids, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.nameIncludes != nil {
		localVarQueryParams.Add("nameIncludes", parameterToString(*r.nameIncludes, ""))
	}
	if r.namesAnyOf != nil {
		localVarQueryParams.Add("names.anyOf", parameterToString(*r.namesAnyOf, ""))
	}
	if r.namesAnyOfCaseSensitive != nil {
		localVarQueryParams.Add("names.anyOf.caseSensitive", parameterToString(*r.namesAnyOfCaseSensitive, ""))
	}
	if r.modifiedAt != nil {
		localVarQueryParams.Add("modifiedAt", parameterToString(*r.modifiedAt, ""))
	}
	if r.memberOf != nil {
		localVarQueryParams.Add("memberOf", parameterToString(*r.memberOf, ""))
	}
	if r.adminOf != nil {
		localVarQueryParams.Add("adminOf", parameterToString(*r.adminOf, ""))
	}
	if r.handles != nil {
		localVarQueryParams.Add("handles", parameterToString(*r.handles, ""))
	}
	if r.passwordLastChangedAt != nil {
		localVarQueryParams.Add("passwordLastChangedAt", parameterToString(*r.passwordLastChangedAt, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.nextToken != nil {
		localVarQueryParams.Add("nextToken", parameterToString(*r.nextToken, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}