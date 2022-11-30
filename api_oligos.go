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


// OligosApiService OligosApi service
type OligosApiService service

type ApiArchiveOligosRequest struct {
	ctx context.Context
	ApiService *OligosApiService
	oligosArchive *OligosArchive
}

func (r ApiArchiveOligosRequest) OligosArchive(oligosArchive OligosArchive) ApiArchiveOligosRequest {
	r.oligosArchive = &oligosArchive
	return r
}

func (r ApiArchiveOligosRequest) Execute() (*OligosArchivalChange, *http.Response, error) {
	return r.ApiService.ArchiveOligosExecute(r)
}

/*
ArchiveOligos Archive Oligos

Archive Oligos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiArchiveOligosRequest

Deprecated
*/
func (a *OligosApiService) ArchiveOligos(ctx context.Context) ApiArchiveOligosRequest {
	return ApiArchiveOligosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OligosArchivalChange
// Deprecated
func (a *OligosApiService) ArchiveOligosExecute(r ApiArchiveOligosRequest) (*OligosArchivalChange, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OligosArchivalChange
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OligosApiService.ArchiveOligos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oligos:archive"

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
	localVarPostBody = r.oligosArchive
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

type ApiBulkCreateOligosRequest struct {
	ctx context.Context
	ApiService *OligosApiService
	oligosBulkCreateRequest *OligosBulkCreateRequest
}

func (r ApiBulkCreateOligosRequest) OligosBulkCreateRequest(oligosBulkCreateRequest OligosBulkCreateRequest) ApiBulkCreateOligosRequest {
	r.oligosBulkCreateRequest = &oligosBulkCreateRequest
	return r
}

func (r ApiBulkCreateOligosRequest) Execute() (*AsyncTaskLink, *http.Response, error) {
	return r.ApiService.BulkCreateOligosExecute(r)
}

/*
BulkCreateOligos Bulk Create DNA Oligos

Bulk Create DNA Oligos
Please migrate to [Bulk Create DNA Oligos](#/DNA%20Oligos/bulkCreateDNAOligos) so that we can support RNA Oligos.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBulkCreateOligosRequest

Deprecated
*/
func (a *OligosApiService) BulkCreateOligos(ctx context.Context) ApiBulkCreateOligosRequest {
	return ApiBulkCreateOligosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncTaskLink
// Deprecated
func (a *OligosApiService) BulkCreateOligosExecute(r ApiBulkCreateOligosRequest) (*AsyncTaskLink, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncTaskLink
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OligosApiService.BulkCreateOligos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oligos:bulk-create"

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
	localVarPostBody = r.oligosBulkCreateRequest
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

type ApiBulkGetOligosRequest struct {
	ctx context.Context
	ApiService *OligosApiService
	oligoIds *string
}

// Comma-separated list of IDs of Oligos to get. 
func (r ApiBulkGetOligosRequest) OligoIds(oligoIds string) ApiBulkGetOligosRequest {
	r.oligoIds = &oligoIds
	return r
}

func (r ApiBulkGetOligosRequest) Execute() (*OligosBulkGet, *http.Response, error) {
	return r.ApiService.BulkGetOligosExecute(r)
}

/*
BulkGetOligos Bulk get Oligos by ID

Bulk get Oligos by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBulkGetOligosRequest
*/
func (a *OligosApiService) BulkGetOligos(ctx context.Context) ApiBulkGetOligosRequest {
	return ApiBulkGetOligosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OligosBulkGet
func (a *OligosApiService) BulkGetOligosExecute(r ApiBulkGetOligosRequest) (*OligosBulkGet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OligosBulkGet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OligosApiService.BulkGetOligos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oligos:bulk-get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oligoIds == nil {
		return localVarReturnValue, nil, reportError("oligoIds is required and must be specified")
	}

	localVarQueryParams.Add("oligoIds", parameterToString(*r.oligoIds, ""))
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

type ApiCreateOligoRequest struct {
	ctx context.Context
	ApiService *OligosApiService
	oligoCreate *OligoCreate
}

func (r ApiCreateOligoRequest) OligoCreate(oligoCreate OligoCreate) ApiCreateOligoRequest {
	r.oligoCreate = &oligoCreate
	return r
}

func (r ApiCreateOligoRequest) Execute() (*DnaOligo, *http.Response, error) {
	return r.ApiService.CreateOligoExecute(r)
}

/*
CreateOligo Create an Oligo

Create an Oligo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateOligoRequest

Deprecated
*/
func (a *OligosApiService) CreateOligo(ctx context.Context) ApiCreateOligoRequest {
	return ApiCreateOligoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DnaOligo
// Deprecated
func (a *OligosApiService) CreateOligoExecute(r ApiCreateOligoRequest) (*DnaOligo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DnaOligo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OligosApiService.CreateOligo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oligos"

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
	localVarPostBody = r.oligoCreate
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

type ApiGetOligoRequest struct {
	ctx context.Context
	ApiService *OligosApiService
	oligoId string
}

func (r ApiGetOligoRequest) Execute() (*DnaOligo, *http.Response, error) {
	return r.ApiService.GetOligoExecute(r)
}

/*
GetOligo Get an Oligo

Get an Oligo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param oligoId
 @return ApiGetOligoRequest

Deprecated
*/
func (a *OligosApiService) GetOligo(ctx context.Context, oligoId string) ApiGetOligoRequest {
	return ApiGetOligoRequest{
		ApiService: a,
		ctx: ctx,
		oligoId: oligoId,
	}
}

// Execute executes the request
//  @return DnaOligo
// Deprecated
func (a *OligosApiService) GetOligoExecute(r ApiGetOligoRequest) (*DnaOligo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DnaOligo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OligosApiService.GetOligo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oligos/{oligo_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"oligo_id"+"}", url.PathEscape(parameterToString(r.oligoId, "")), -1)

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

type ApiListOligosRequest struct {
	ctx context.Context
	ApiService *OligosApiService
	pageSize *int32
	nextToken *string
	sort *string
	modifiedAt *string
	name *string
	bases *string
	folderId *string
	mentionedIn *string
	projectId *string
	registryId *string
	schemaId *string
	schemaFields *map[string]interface{}
	archiveReason *string
	mentions *string
	ids *string
	entityRegistryIdsAnyOf *string
	namesAnyOf *string
	namesAnyOfCaseSensitive *string
	creatorIds *string
}

// Number of results to return. Defaults to 50, maximum of 100. 
func (r ApiListOligosRequest) PageSize(pageSize int32) ApiListOligosRequest {
	r.pageSize = &pageSize
	return r
}

// Token for pagination
func (r ApiListOligosRequest) NextToken(nextToken string) ApiListOligosRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiListOligosRequest) Sort(sort string) ApiListOligosRequest {
	r.sort = &sort
	return r
}

// Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30. 
func (r ApiListOligosRequest) ModifiedAt(modifiedAt string) ApiListOligosRequest {
	r.modifiedAt = &modifiedAt
	return r
}

// Name of an Oligo. Restricts results to those with the specified name, alias, or entity registry ID.
func (r ApiListOligosRequest) Name(name string) ApiListOligosRequest {
	r.name = &name
	return r
}

// Full bases of the oligo. Restricts results to those with the specified bases, case-insensitive, allowing for circular or reverse complement matches. Does not allow partial matching or loose matching via degenerate bases. 
func (r ApiListOligosRequest) Bases(bases string) ApiListOligosRequest {
	r.bases = &bases
	return r
}

// ID of a folder. Restricts results to those in the folder.
func (r ApiListOligosRequest) FolderId(folderId string) ApiListOligosRequest {
	r.folderId = &folderId
	return r
}

// Comma-separated list of entry IDs. Restricts results to Oligos mentioned in those entries. 
func (r ApiListOligosRequest) MentionedIn(mentionedIn string) ApiListOligosRequest {
	r.mentionedIn = &mentionedIn
	return r
}

// ID of a project. Restricts results to those in the project.
func (r ApiListOligosRequest) ProjectId(projectId string) ApiListOligosRequest {
	r.projectId = &projectId
	return r
}

// ID of a registry. Restricts results to those registered in this registry. Specifying \&quot;null\&quot; returns unregistered items. 
func (r ApiListOligosRequest) RegistryId(registryId string) ApiListOligosRequest {
	r.registryId = &registryId
	return r
}

// ID of a schema. Restricts results to those of the specified schema. 
func (r ApiListOligosRequest) SchemaId(schemaId string) ApiListOligosRequest {
	r.schemaId = &schemaId
	return r
}

// Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter. 
func (r ApiListOligosRequest) SchemaFields(schemaFields map[string]interface{}) ApiListOligosRequest {
	r.schemaFields = &schemaFields
	return r
}

// Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived Oligos. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived Oligos regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived. 
func (r ApiListOligosRequest) ArchiveReason(archiveReason string) ApiListOligosRequest {
	r.archiveReason = &archiveReason
	return r
}

// Comma-separated list of item IDs. Restricts results to those that mention the given items in the description. 
func (r ApiListOligosRequest) Mentions(mentions string) ApiListOligosRequest {
	r.mentions = &mentions
	return r
}

// Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid. 
func (r ApiListOligosRequest) Ids(ids string) ApiListOligosRequest {
	r.ids = &ids
	return r
}

// Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs. 
func (r ApiListOligosRequest) EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf string) ApiListOligosRequest {
	r.entityRegistryIdsAnyOf = &entityRegistryIdsAnyOf
	return r
}

// Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity. 
func (r ApiListOligosRequest) NamesAnyOf(namesAnyOf string) ApiListOligosRequest {
	r.namesAnyOf = &namesAnyOf
	return r
}

// Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive. 
func (r ApiListOligosRequest) NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive string) ApiListOligosRequest {
	r.namesAnyOfCaseSensitive = &namesAnyOfCaseSensitive
	return r
}

// Comma separated list of users IDs
func (r ApiListOligosRequest) CreatorIds(creatorIds string) ApiListOligosRequest {
	r.creatorIds = &creatorIds
	return r
}

func (r ApiListOligosRequest) Execute() (*OligosPaginatedList, *http.Response, error) {
	return r.ApiService.ListOligosExecute(r)
}

/*
ListOligos List Oligos

List Oligos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListOligosRequest

Deprecated
*/
func (a *OligosApiService) ListOligos(ctx context.Context) ApiListOligosRequest {
	return ApiListOligosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OligosPaginatedList
// Deprecated
func (a *OligosApiService) ListOligosExecute(r ApiListOligosRequest) (*OligosPaginatedList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OligosPaginatedList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OligosApiService.ListOligos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oligos"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.nextToken != nil {
		localVarQueryParams.Add("nextToken", parameterToString(*r.nextToken, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.modifiedAt != nil {
		localVarQueryParams.Add("modifiedAt", parameterToString(*r.modifiedAt, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.bases != nil {
		localVarQueryParams.Add("bases", parameterToString(*r.bases, ""))
	}
	if r.folderId != nil {
		localVarQueryParams.Add("folderId", parameterToString(*r.folderId, ""))
	}
	if r.mentionedIn != nil {
		localVarQueryParams.Add("mentionedIn", parameterToString(*r.mentionedIn, ""))
	}
	if r.projectId != nil {
		localVarQueryParams.Add("projectId", parameterToString(*r.projectId, ""))
	}
	if r.registryId != nil {
		localVarQueryParams.Add("registryId", parameterToString(*r.registryId, ""))
	}
	if r.schemaId != nil {
		localVarQueryParams.Add("schemaId", parameterToString(*r.schemaId, ""))
	}
	if r.schemaFields != nil {
		localVarQueryParams.Add("schemaFields", parameterToString(*r.schemaFields, ""))
	}
	if r.archiveReason != nil {
		localVarQueryParams.Add("archiveReason", parameterToString(*r.archiveReason, ""))
	}
	if r.mentions != nil {
		localVarQueryParams.Add("mentions", parameterToString(*r.mentions, ""))
	}
	if r.ids != nil {
		localVarQueryParams.Add("ids", parameterToString(*r.ids, ""))
	}
	if r.entityRegistryIdsAnyOf != nil {
		localVarQueryParams.Add("entityRegistryIds.anyOf", parameterToString(*r.entityRegistryIdsAnyOf, ""))
	}
	if r.namesAnyOf != nil {
		localVarQueryParams.Add("names.anyOf", parameterToString(*r.namesAnyOf, ""))
	}
	if r.namesAnyOfCaseSensitive != nil {
		localVarQueryParams.Add("names.anyOf.caseSensitive", parameterToString(*r.namesAnyOfCaseSensitive, ""))
	}
	if r.creatorIds != nil {
		localVarQueryParams.Add("creatorIds", parameterToString(*r.creatorIds, ""))
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

type ApiUnarchiveOligosRequest struct {
	ctx context.Context
	ApiService *OligosApiService
	oligosUnarchive *OligosUnarchive
}

func (r ApiUnarchiveOligosRequest) OligosUnarchive(oligosUnarchive OligosUnarchive) ApiUnarchiveOligosRequest {
	r.oligosUnarchive = &oligosUnarchive
	return r
}

func (r ApiUnarchiveOligosRequest) Execute() (*OligosArchivalChange, *http.Response, error) {
	return r.ApiService.UnarchiveOligosExecute(r)
}

/*
UnarchiveOligos Unarchive Oligos

Unarchive Oligos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUnarchiveOligosRequest

Deprecated
*/
func (a *OligosApiService) UnarchiveOligos(ctx context.Context) ApiUnarchiveOligosRequest {
	return ApiUnarchiveOligosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OligosArchivalChange
// Deprecated
func (a *OligosApiService) UnarchiveOligosExecute(r ApiUnarchiveOligosRequest) (*OligosArchivalChange, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OligosArchivalChange
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OligosApiService.UnarchiveOligos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oligos:unarchive"

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
	localVarPostBody = r.oligosUnarchive
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

type ApiUpdateOligoRequest struct {
	ctx context.Context
	ApiService *OligosApiService
	oligoId string
	oligoUpdate *OligoUpdate
}

func (r ApiUpdateOligoRequest) OligoUpdate(oligoUpdate OligoUpdate) ApiUpdateOligoRequest {
	r.oligoUpdate = &oligoUpdate
	return r
}

func (r ApiUpdateOligoRequest) Execute() (*DnaOligo, *http.Response, error) {
	return r.ApiService.UpdateOligoExecute(r)
}

/*
UpdateOligo Update an Oligo

Update an Oligo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param oligoId
 @return ApiUpdateOligoRequest

Deprecated
*/
func (a *OligosApiService) UpdateOligo(ctx context.Context, oligoId string) ApiUpdateOligoRequest {
	return ApiUpdateOligoRequest{
		ApiService: a,
		ctx: ctx,
		oligoId: oligoId,
	}
}

// Execute executes the request
//  @return DnaOligo
// Deprecated
func (a *OligosApiService) UpdateOligoExecute(r ApiUpdateOligoRequest) (*DnaOligo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DnaOligo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OligosApiService.UpdateOligo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oligos/{oligo_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"oligo_id"+"}", url.PathEscape(parameterToString(r.oligoId, "")), -1)

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
	localVarPostBody = r.oligoUpdate
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
