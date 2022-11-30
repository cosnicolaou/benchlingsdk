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


// WorkflowTaskGroupsApiService WorkflowTaskGroupsApi service
type WorkflowTaskGroupsApiService service

type ApiArchiveWorkflowTaskGroupsRequest struct {
	ctx context.Context
	ApiService *WorkflowTaskGroupsApiService
	workflowTaskGroupsArchive *WorkflowTaskGroupsArchive
}

func (r ApiArchiveWorkflowTaskGroupsRequest) WorkflowTaskGroupsArchive(workflowTaskGroupsArchive WorkflowTaskGroupsArchive) ApiArchiveWorkflowTaskGroupsRequest {
	r.workflowTaskGroupsArchive = &workflowTaskGroupsArchive
	return r
}

func (r ApiArchiveWorkflowTaskGroupsRequest) Execute() (*WorkflowTaskGroupsArchivalChange, *http.Response, error) {
	return r.ApiService.ArchiveWorkflowTaskGroupsExecute(r)
}

/*
ArchiveWorkflowTaskGroups Archive one or more workflows

Archive one or more workflows

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiArchiveWorkflowTaskGroupsRequest
*/
func (a *WorkflowTaskGroupsApiService) ArchiveWorkflowTaskGroups(ctx context.Context) ApiArchiveWorkflowTaskGroupsRequest {
	return ApiArchiveWorkflowTaskGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WorkflowTaskGroupsArchivalChange
func (a *WorkflowTaskGroupsApiService) ArchiveWorkflowTaskGroupsExecute(r ApiArchiveWorkflowTaskGroupsRequest) (*WorkflowTaskGroupsArchivalChange, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorkflowTaskGroupsArchivalChange
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowTaskGroupsApiService.ArchiveWorkflowTaskGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workflow-task-groups:archive"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.workflowTaskGroupsArchive
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

type ApiCreateWorkflowTaskGroupRequest struct {
	ctx context.Context
	ApiService *WorkflowTaskGroupsApiService
	workflowTaskGroupCreate *WorkflowTaskGroupCreate
}

func (r ApiCreateWorkflowTaskGroupRequest) WorkflowTaskGroupCreate(workflowTaskGroupCreate WorkflowTaskGroupCreate) ApiCreateWorkflowTaskGroupRequest {
	r.workflowTaskGroupCreate = &workflowTaskGroupCreate
	return r
}

func (r ApiCreateWorkflowTaskGroupRequest) Execute() (*WorkflowTaskGroup, *http.Response, error) {
	return r.ApiService.CreateWorkflowTaskGroupExecute(r)
}

/*
CreateWorkflowTaskGroup Create a new workflow task group

Create a new workflow task group. If no name is specified, uses the workflow schema name and a unique incrementor separated by a single whitespace.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateWorkflowTaskGroupRequest
*/
func (a *WorkflowTaskGroupsApiService) CreateWorkflowTaskGroup(ctx context.Context) ApiCreateWorkflowTaskGroupRequest {
	return ApiCreateWorkflowTaskGroupRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WorkflowTaskGroup
func (a *WorkflowTaskGroupsApiService) CreateWorkflowTaskGroupExecute(r ApiCreateWorkflowTaskGroupRequest) (*WorkflowTaskGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorkflowTaskGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowTaskGroupsApiService.CreateWorkflowTaskGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workflow-task-groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.workflowTaskGroupCreate
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

type ApiGetWorkflowTaskGroupRequest struct {
	ctx context.Context
	ApiService *WorkflowTaskGroupsApiService
	workflowTaskGroupId string
}

func (r ApiGetWorkflowTaskGroupRequest) Execute() (*WorkflowTaskGroup, *http.Response, error) {
	return r.ApiService.GetWorkflowTaskGroupExecute(r)
}

/*
GetWorkflowTaskGroup Get a workflow task group

Get a workflow task group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workflowTaskGroupId The ID of the workflow task group
 @return ApiGetWorkflowTaskGroupRequest
*/
func (a *WorkflowTaskGroupsApiService) GetWorkflowTaskGroup(ctx context.Context, workflowTaskGroupId string) ApiGetWorkflowTaskGroupRequest {
	return ApiGetWorkflowTaskGroupRequest{
		ApiService: a,
		ctx: ctx,
		workflowTaskGroupId: workflowTaskGroupId,
	}
}

// Execute executes the request
//  @return WorkflowTaskGroup
func (a *WorkflowTaskGroupsApiService) GetWorkflowTaskGroupExecute(r ApiGetWorkflowTaskGroupRequest) (*WorkflowTaskGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorkflowTaskGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowTaskGroupsApiService.GetWorkflowTaskGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workflow-task-groups/{workflow_task_group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"workflow_task_group_id"+"}", url.PathEscape(parameterToString(r.workflowTaskGroupId, "")), -1)

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

type ApiListWorkflowTaskGroupsRequest struct {
	ctx context.Context
	ApiService *WorkflowTaskGroupsApiService
	ids *string
	schemaId *string
	folderId *string
	projectId *string
	mentionedIn *string
	watcherIds *string
	executionTypes *string
	responsibleTeamIds *string
	statusIdsAnyOf *string
	statusIdsNoneOf *string
	statusIdsOnly *string
	name *string
	nameIncludes *string
	creatorIds *string
	modifiedAt *string
	nextToken *string
	pageSize *int32
	displayIds *string
	archiveReason *string
}

// Comma separated list of workflow task group IDs
func (r ApiListWorkflowTaskGroupsRequest) Ids(ids string) ApiListWorkflowTaskGroupsRequest {
	r.ids = &ids
	return r
}

// The workflow task schema ID of tasks in this task group
func (r ApiListWorkflowTaskGroupsRequest) SchemaId(schemaId string) ApiListWorkflowTaskGroupsRequest {
	r.schemaId = &schemaId
	return r
}

// A folder ID
func (r ApiListWorkflowTaskGroupsRequest) FolderId(folderId string) ApiListWorkflowTaskGroupsRequest {
	r.folderId = &folderId
	return r
}

// A project ID
func (r ApiListWorkflowTaskGroupsRequest) ProjectId(projectId string) ApiListWorkflowTaskGroupsRequest {
	r.projectId = &projectId
	return r
}

// A comma separated list entry IDs
func (r ApiListWorkflowTaskGroupsRequest) MentionedIn(mentionedIn string) ApiListWorkflowTaskGroupsRequest {
	r.mentionedIn = &mentionedIn
	return r
}

// Comma separated list of user IDs or \&quot;null\&quot;
func (r ApiListWorkflowTaskGroupsRequest) WatcherIds(watcherIds string) ApiListWorkflowTaskGroupsRequest {
	r.watcherIds = &watcherIds
	return r
}

// Comma separated list of workflow execution types. Acceptable execution types are \&quot;DIRECT\&quot; and \&quot;ENTRY\&quot; 
func (r ApiListWorkflowTaskGroupsRequest) ExecutionTypes(executionTypes string) ApiListWorkflowTaskGroupsRequest {
	r.executionTypes = &executionTypes
	return r
}

// Comma separated list of team IDs or \&quot;null\&quot;
func (r ApiListWorkflowTaskGroupsRequest) ResponsibleTeamIds(responsibleTeamIds string) ApiListWorkflowTaskGroupsRequest {
	r.responsibleTeamIds = &responsibleTeamIds
	return r
}

// Commas separated list of Status ids. Returns workflows where at least one task is of one of the provided statuses.
func (r ApiListWorkflowTaskGroupsRequest) StatusIdsAnyOf(statusIdsAnyOf string) ApiListWorkflowTaskGroupsRequest {
	r.statusIdsAnyOf = &statusIdsAnyOf
	return r
}

// Commas separated list of Status ids. Returns workflows where none of the tasks are of any of the provided statuses.
func (r ApiListWorkflowTaskGroupsRequest) StatusIdsNoneOf(statusIdsNoneOf string) ApiListWorkflowTaskGroupsRequest {
	r.statusIdsNoneOf = &statusIdsNoneOf
	return r
}

// Commas separated list of Status ids. Returns workflows where all of the tasks are of one of the provided statuses.
func (r ApiListWorkflowTaskGroupsRequest) StatusIdsOnly(statusIdsOnly string) ApiListWorkflowTaskGroupsRequest {
	r.statusIdsOnly = &statusIdsOnly
	return r
}

// The name of the workflow task group
func (r ApiListWorkflowTaskGroupsRequest) Name(name string) ApiListWorkflowTaskGroupsRequest {
	r.name = &name
	return r
}

// Part of the name of the workflow task group
func (r ApiListWorkflowTaskGroupsRequest) NameIncludes(nameIncludes string) ApiListWorkflowTaskGroupsRequest {
	r.nameIncludes = &nameIncludes
	return r
}

// Comma separated list of user IDs.
func (r ApiListWorkflowTaskGroupsRequest) CreatorIds(creatorIds string) ApiListWorkflowTaskGroupsRequest {
	r.creatorIds = &creatorIds
	return r
}

// Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30. 
func (r ApiListWorkflowTaskGroupsRequest) ModifiedAt(modifiedAt string) ApiListWorkflowTaskGroupsRequest {
	r.modifiedAt = &modifiedAt
	return r
}

func (r ApiListWorkflowTaskGroupsRequest) NextToken(nextToken string) ApiListWorkflowTaskGroupsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiListWorkflowTaskGroupsRequest) PageSize(pageSize int32) ApiListWorkflowTaskGroupsRequest {
	r.pageSize = &pageSize
	return r
}

// Comma-separated list of Workflow Display IDs.
func (r ApiListWorkflowTaskGroupsRequest) DisplayIds(displayIds string) ApiListWorkflowTaskGroupsRequest {
	r.displayIds = &displayIds
	return r
}

// Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived workflow task groups. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived workflow task groups regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived. 
func (r ApiListWorkflowTaskGroupsRequest) ArchiveReason(archiveReason string) ApiListWorkflowTaskGroupsRequest {
	r.archiveReason = &archiveReason
	return r
}

func (r ApiListWorkflowTaskGroupsRequest) Execute() (*WorkflowTaskGroupsPaginatedList, *http.Response, error) {
	return r.ApiService.ListWorkflowTaskGroupsExecute(r)
}

/*
ListWorkflowTaskGroups List workflow task groups

List workflow task groups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListWorkflowTaskGroupsRequest
*/
func (a *WorkflowTaskGroupsApiService) ListWorkflowTaskGroups(ctx context.Context) ApiListWorkflowTaskGroupsRequest {
	return ApiListWorkflowTaskGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WorkflowTaskGroupsPaginatedList
func (a *WorkflowTaskGroupsApiService) ListWorkflowTaskGroupsExecute(r ApiListWorkflowTaskGroupsRequest) (*WorkflowTaskGroupsPaginatedList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorkflowTaskGroupsPaginatedList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowTaskGroupsApiService.ListWorkflowTaskGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workflow-task-groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ids != nil {
		localVarQueryParams.Add("ids", parameterToString(*r.ids, ""))
	}
	if r.schemaId != nil {
		localVarQueryParams.Add("schemaId", parameterToString(*r.schemaId, ""))
	}
	if r.folderId != nil {
		localVarQueryParams.Add("folderId", parameterToString(*r.folderId, ""))
	}
	if r.projectId != nil {
		localVarQueryParams.Add("projectId", parameterToString(*r.projectId, ""))
	}
	if r.mentionedIn != nil {
		localVarQueryParams.Add("mentionedIn", parameterToString(*r.mentionedIn, ""))
	}
	if r.watcherIds != nil {
		localVarQueryParams.Add("watcherIds", parameterToString(*r.watcherIds, ""))
	}
	if r.executionTypes != nil {
		localVarQueryParams.Add("executionTypes", parameterToString(*r.executionTypes, ""))
	}
	if r.responsibleTeamIds != nil {
		localVarQueryParams.Add("responsibleTeamIds", parameterToString(*r.responsibleTeamIds, ""))
	}
	if r.statusIdsAnyOf != nil {
		localVarQueryParams.Add("statusIds.anyOf", parameterToString(*r.statusIdsAnyOf, ""))
	}
	if r.statusIdsNoneOf != nil {
		localVarQueryParams.Add("statusIds.noneOf", parameterToString(*r.statusIdsNoneOf, ""))
	}
	if r.statusIdsOnly != nil {
		localVarQueryParams.Add("statusIds.only", parameterToString(*r.statusIdsOnly, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.nameIncludes != nil {
		localVarQueryParams.Add("nameIncludes", parameterToString(*r.nameIncludes, ""))
	}
	if r.creatorIds != nil {
		localVarQueryParams.Add("creatorIds", parameterToString(*r.creatorIds, ""))
	}
	if r.modifiedAt != nil {
		localVarQueryParams.Add("modifiedAt", parameterToString(*r.modifiedAt, ""))
	}
	if r.nextToken != nil {
		localVarQueryParams.Add("nextToken", parameterToString(*r.nextToken, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.displayIds != nil {
		localVarQueryParams.Add("displayIds", parameterToString(*r.displayIds, ""))
	}
	if r.archiveReason != nil {
		localVarQueryParams.Add("archiveReason", parameterToString(*r.archiveReason, ""))
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

type ApiUnarchiveWorkflowTaskGroupsRequest struct {
	ctx context.Context
	ApiService *WorkflowTaskGroupsApiService
	workflowTaskGroupsUnarchive *WorkflowTaskGroupsUnarchive
}

func (r ApiUnarchiveWorkflowTaskGroupsRequest) WorkflowTaskGroupsUnarchive(workflowTaskGroupsUnarchive WorkflowTaskGroupsUnarchive) ApiUnarchiveWorkflowTaskGroupsRequest {
	r.workflowTaskGroupsUnarchive = &workflowTaskGroupsUnarchive
	return r
}

func (r ApiUnarchiveWorkflowTaskGroupsRequest) Execute() (*WorkflowTaskGroupsArchivalChange, *http.Response, error) {
	return r.ApiService.UnarchiveWorkflowTaskGroupsExecute(r)
}

/*
UnarchiveWorkflowTaskGroups Unarchive one or more workflows

Unarchive one or more workflows

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUnarchiveWorkflowTaskGroupsRequest
*/
func (a *WorkflowTaskGroupsApiService) UnarchiveWorkflowTaskGroups(ctx context.Context) ApiUnarchiveWorkflowTaskGroupsRequest {
	return ApiUnarchiveWorkflowTaskGroupsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WorkflowTaskGroupsArchivalChange
func (a *WorkflowTaskGroupsApiService) UnarchiveWorkflowTaskGroupsExecute(r ApiUnarchiveWorkflowTaskGroupsRequest) (*WorkflowTaskGroupsArchivalChange, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorkflowTaskGroupsArchivalChange
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowTaskGroupsApiService.UnarchiveWorkflowTaskGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workflow-task-groups:unarchive"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.workflowTaskGroupsUnarchive
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

type ApiUpdateWorkflowTaskGroupRequest struct {
	ctx context.Context
	ApiService *WorkflowTaskGroupsApiService
	workflowTaskGroupId string
	workflowTaskGroupUpdate *WorkflowTaskGroupUpdate
}

func (r ApiUpdateWorkflowTaskGroupRequest) WorkflowTaskGroupUpdate(workflowTaskGroupUpdate WorkflowTaskGroupUpdate) ApiUpdateWorkflowTaskGroupRequest {
	r.workflowTaskGroupUpdate = &workflowTaskGroupUpdate
	return r
}

func (r ApiUpdateWorkflowTaskGroupRequest) Execute() (*WorkflowTaskGroup, *http.Response, error) {
	return r.ApiService.UpdateWorkflowTaskGroupExecute(r)
}

/*
UpdateWorkflowTaskGroup Update a workflow task group

Update a workflow task group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workflowTaskGroupId The ID of the workflow task group
 @return ApiUpdateWorkflowTaskGroupRequest
*/
func (a *WorkflowTaskGroupsApiService) UpdateWorkflowTaskGroup(ctx context.Context, workflowTaskGroupId string) ApiUpdateWorkflowTaskGroupRequest {
	return ApiUpdateWorkflowTaskGroupRequest{
		ApiService: a,
		ctx: ctx,
		workflowTaskGroupId: workflowTaskGroupId,
	}
}

// Execute executes the request
//  @return WorkflowTaskGroup
func (a *WorkflowTaskGroupsApiService) UpdateWorkflowTaskGroupExecute(r ApiUpdateWorkflowTaskGroupRequest) (*WorkflowTaskGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorkflowTaskGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkflowTaskGroupsApiService.UpdateWorkflowTaskGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workflow-task-groups/{workflow_task_group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"workflow_task_group_id"+"}", url.PathEscape(parameterToString(r.workflowTaskGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.workflowTaskGroupUpdate
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
			return localVarReturnValue, localVarHTTPResponse, newErr
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