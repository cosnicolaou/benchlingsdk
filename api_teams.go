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


// TeamsApiService TeamsApi service
type TeamsApiService service

type ApiGetTeamRequest struct {
	ctx context.Context
	ApiService *TeamsApiService
	teamId string
}

func (r ApiGetTeamRequest) Execute() (*Team, *http.Response, error) {
	return r.ApiService.GetTeamExecute(r)
}

/*
GetTeam Get a team by ID

Returns a team by ID if the caller has permission to view. The following roles have view permission:
  - tenant admins
  - members of the team's organization


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId ID of team to get
 @return ApiGetTeamRequest
*/
func (a *TeamsApiService) GetTeam(ctx context.Context, teamId string) ApiGetTeamRequest {
	return ApiGetTeamRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return Team
func (a *TeamsApiService) GetTeamExecute(r ApiGetTeamRequest) (*Team, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Team
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.GetTeam")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{team_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team_id"+"}", url.PathEscape(parameterToString(r.teamId, "")), -1)

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

type ApiListTeamsRequest struct {
	ctx context.Context
	ApiService *TeamsApiService
	ids *string
	name *string
	nameIncludes *string
	namesAnyOf *string
	namesAnyOfCaseSensitive *string
	modifiedAt *string
	mentionedIn *string
	organizationId *string
	hasMembers *string
	hasAdmins *string
	pageSize *int32
	nextToken *string
	sort *string
}

// Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid. 
func (r ApiListTeamsRequest) Ids(ids string) ApiListTeamsRequest {
	r.ids = &ids
	return r
}

// Name of a team. Restricts results to those with the specified name.
func (r ApiListTeamsRequest) Name(name string) ApiListTeamsRequest {
	r.name = &name
	return r
}

// Name substring of a team. Restricts results to those with names that include the provided substring.
func (r ApiListTeamsRequest) NameIncludes(nameIncludes string) ApiListTeamsRequest {
	r.nameIncludes = &nameIncludes
	return r
}

// Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity. 
func (r ApiListTeamsRequest) NamesAnyOf(namesAnyOf string) ApiListTeamsRequest {
	r.namesAnyOf = &namesAnyOf
	return r
}

// Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive. 
func (r ApiListTeamsRequest) NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive string) ApiListTeamsRequest {
	r.namesAnyOfCaseSensitive = &namesAnyOfCaseSensitive
	return r
}

// Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30. 
func (r ApiListTeamsRequest) ModifiedAt(modifiedAt string) ApiListTeamsRequest {
	r.modifiedAt = &modifiedAt
	return r
}

// Comma-separated list of entry IDs. Restricts results to teams mentioned in those entries. 
func (r ApiListTeamsRequest) MentionedIn(mentionedIn string) ApiListTeamsRequest {
	r.mentionedIn = &mentionedIn
	return r
}

// Restricts results to those in the organization.
func (r ApiListTeamsRequest) OrganizationId(organizationId string) ApiListTeamsRequest {
	r.organizationId = &organizationId
	return r
}

// Comma-separated list of user or Benchling app IDs. Restricts results to teams that include all the given users/apps as members.
func (r ApiListTeamsRequest) HasMembers(hasMembers string) ApiListTeamsRequest {
	r.hasMembers = &hasMembers
	return r
}

// Comma-separated list of user or Benchling app IDs. Restricts results to teams that include all the given users/apps as admins.
func (r ApiListTeamsRequest) HasAdmins(hasAdmins string) ApiListTeamsRequest {
	r.hasAdmins = &hasAdmins
	return r
}

func (r ApiListTeamsRequest) PageSize(pageSize int32) ApiListTeamsRequest {
	r.pageSize = &pageSize
	return r
}

// Token for pagination
func (r ApiListTeamsRequest) NextToken(nextToken string) ApiListTeamsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiListTeamsRequest) Sort(sort string) ApiListTeamsRequest {
	r.sort = &sort
	return r
}

func (r ApiListTeamsRequest) Execute() (*TeamsPaginatedList, *http.Response, error) {
	return r.ApiService.ListTeamsExecute(r)
}

/*
ListTeams List teams

Returns all teams that the caller has permission to view. The following roles have view permission:
  - tenant admins
  - members of the team's organization


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListTeamsRequest
*/
func (a *TeamsApiService) ListTeams(ctx context.Context) ApiListTeamsRequest {
	return ApiListTeamsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TeamsPaginatedList
func (a *TeamsApiService) ListTeamsExecute(r ApiListTeamsRequest) (*TeamsPaginatedList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeamsPaginatedList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.ListTeams")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams"

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
	if r.mentionedIn != nil {
		localVarQueryParams.Add("mentionedIn", parameterToString(*r.mentionedIn, ""))
	}
	if r.organizationId != nil {
		localVarQueryParams.Add("organizationId", parameterToString(*r.organizationId, ""))
	}
	if r.hasMembers != nil {
		localVarQueryParams.Add("hasMembers", parameterToString(*r.hasMembers, ""))
	}
	if r.hasAdmins != nil {
		localVarQueryParams.Add("hasAdmins", parameterToString(*r.hasAdmins, ""))
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
