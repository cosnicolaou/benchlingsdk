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


// DNAOligosApiService DNAOligosApi service
type DNAOligosApiService service

type ApiArchiveDNAOligosRequest struct {
	ctx context.Context
	ApiService *DNAOligosApiService
	dnaOligosArchive *DnaOligosArchive
}

func (r ApiArchiveDNAOligosRequest) DnaOligosArchive(dnaOligosArchive DnaOligosArchive) ApiArchiveDNAOligosRequest {
	r.dnaOligosArchive = &dnaOligosArchive
	return r
}

func (r ApiArchiveDNAOligosRequest) Execute() (*DnaOligosArchivalChange, *http.Response, error) {
	return r.ApiService.ArchiveDNAOligosExecute(r)
}

/*
ArchiveDNAOligos Archive DNA Oligos

Archive DNA Oligos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiArchiveDNAOligosRequest
*/
func (a *DNAOligosApiService) ArchiveDNAOligos(ctx context.Context) ApiArchiveDNAOligosRequest {
	return ApiArchiveDNAOligosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DnaOligosArchivalChange
func (a *DNAOligosApiService) ArchiveDNAOligosExecute(r ApiArchiveDNAOligosRequest) (*DnaOligosArchivalChange, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DnaOligosArchivalChange
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNAOligosApiService.ArchiveDNAOligos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dna-oligos:archive"

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
	localVarPostBody = r.dnaOligosArchive
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

type ApiBulkCreateDNAOligosRequest struct {
	ctx context.Context
	ApiService *DNAOligosApiService
	dnaOligosBulkCreateRequest *DnaOligosBulkCreateRequest
}

func (r ApiBulkCreateDNAOligosRequest) DnaOligosBulkCreateRequest(dnaOligosBulkCreateRequest DnaOligosBulkCreateRequest) ApiBulkCreateDNAOligosRequest {
	r.dnaOligosBulkCreateRequest = &dnaOligosBulkCreateRequest
	return r
}

func (r ApiBulkCreateDNAOligosRequest) Execute() (*AsyncTaskLink, *http.Response, error) {
	return r.ApiService.BulkCreateDNAOligosExecute(r)
}

/*
BulkCreateDNAOligos Bulk Create DNA Oligos

Bulk Create DNA Oligos. Limit of 1000 DNA Oligos per request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBulkCreateDNAOligosRequest
*/
func (a *DNAOligosApiService) BulkCreateDNAOligos(ctx context.Context) ApiBulkCreateDNAOligosRequest {
	return ApiBulkCreateDNAOligosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncTaskLink
func (a *DNAOligosApiService) BulkCreateDNAOligosExecute(r ApiBulkCreateDNAOligosRequest) (*AsyncTaskLink, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncTaskLink
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNAOligosApiService.BulkCreateDNAOligos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dna-oligos:bulk-create"

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
	localVarPostBody = r.dnaOligosBulkCreateRequest
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

type ApiBulkUpdateDNAOligosRequest struct {
	ctx context.Context
	ApiService *DNAOligosApiService
	dnaOligosBulkUpdateRequest *DnaOligosBulkUpdateRequest
}

func (r ApiBulkUpdateDNAOligosRequest) DnaOligosBulkUpdateRequest(dnaOligosBulkUpdateRequest DnaOligosBulkUpdateRequest) ApiBulkUpdateDNAOligosRequest {
	r.dnaOligosBulkUpdateRequest = &dnaOligosBulkUpdateRequest
	return r
}

func (r ApiBulkUpdateDNAOligosRequest) Execute() (*AsyncTaskLink, *http.Response, error) {
	return r.ApiService.BulkUpdateDNAOligosExecute(r)
}

/*
BulkUpdateDNAOligos Bulk Update DNA Oligos

Bulk Update DNA Oligos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBulkUpdateDNAOligosRequest
*/
func (a *DNAOligosApiService) BulkUpdateDNAOligos(ctx context.Context) ApiBulkUpdateDNAOligosRequest {
	return ApiBulkUpdateDNAOligosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncTaskLink
func (a *DNAOligosApiService) BulkUpdateDNAOligosExecute(r ApiBulkUpdateDNAOligosRequest) (*AsyncTaskLink, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncTaskLink
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNAOligosApiService.BulkUpdateDNAOligos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dna-oligos:bulk-update"

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
	localVarPostBody = r.dnaOligosBulkUpdateRequest
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

type ApiCreateDNAOligoRequest struct {
	ctx context.Context
	ApiService *DNAOligosApiService
	dnaOligoCreate *DnaOligoCreate
}

func (r ApiCreateDNAOligoRequest) DnaOligoCreate(dnaOligoCreate DnaOligoCreate) ApiCreateDNAOligoRequest {
	r.dnaOligoCreate = &dnaOligoCreate
	return r
}

func (r ApiCreateDNAOligoRequest) Execute() (*DnaOligo, *http.Response, error) {
	return r.ApiService.CreateDNAOligoExecute(r)
}

/*
CreateDNAOligo Create a DNA Oligo

Create a DNA Oligo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDNAOligoRequest
*/
func (a *DNAOligosApiService) CreateDNAOligo(ctx context.Context) ApiCreateDNAOligoRequest {
	return ApiCreateDNAOligoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DnaOligo
func (a *DNAOligosApiService) CreateDNAOligoExecute(r ApiCreateDNAOligoRequest) (*DnaOligo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DnaOligo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNAOligosApiService.CreateDNAOligo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dna-oligos"

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
	localVarPostBody = r.dnaOligoCreate
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

type ApiGetDNAOligoRequest struct {
	ctx context.Context
	ApiService *DNAOligosApiService
	oligoId string
}

func (r ApiGetDNAOligoRequest) Execute() (*DnaOligo, *http.Response, error) {
	return r.ApiService.GetDNAOligoExecute(r)
}

/*
GetDNAOligo Get a DNA Oligo

Get a DNA Oligo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param oligoId
 @return ApiGetDNAOligoRequest
*/
func (a *DNAOligosApiService) GetDNAOligo(ctx context.Context, oligoId string) ApiGetDNAOligoRequest {
	return ApiGetDNAOligoRequest{
		ApiService: a,
		ctx: ctx,
		oligoId: oligoId,
	}
}

// Execute executes the request
//  @return DnaOligo
func (a *DNAOligosApiService) GetDNAOligoExecute(r ApiGetDNAOligoRequest) (*DnaOligo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DnaOligo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNAOligosApiService.GetDNAOligo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dna-oligos/{oligo_id}"
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

type ApiListDNAOligosRequest struct {
	ctx context.Context
	ApiService *DNAOligosApiService
	pageSize *int32
	nextToken *string
	sort *string
	modifiedAt *string
	name *string
	nameIncludes *string
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
	authorIdsAnyOf *string
}

// Number of results to return. Defaults to 50, maximum of 100. 
func (r ApiListDNAOligosRequest) PageSize(pageSize int32) ApiListDNAOligosRequest {
	r.pageSize = &pageSize
	return r
}

// Token for pagination
func (r ApiListDNAOligosRequest) NextToken(nextToken string) ApiListDNAOligosRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiListDNAOligosRequest) Sort(sort string) ApiListDNAOligosRequest {
	r.sort = &sort
	return r
}

// Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30. 
func (r ApiListDNAOligosRequest) ModifiedAt(modifiedAt string) ApiListDNAOligosRequest {
	r.modifiedAt = &modifiedAt
	return r
}

// Name of a DNA Oligo. Restricts results to those with the specified name, alias, or entity registry ID.
func (r ApiListDNAOligosRequest) Name(name string) ApiListDNAOligosRequest {
	r.name = &name
	return r
}

// Name substring of a DNA Oligo. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring.
func (r ApiListDNAOligosRequest) NameIncludes(nameIncludes string) ApiListDNAOligosRequest {
	r.nameIncludes = &nameIncludes
	return r
}

// Full bases of the DNA Oligo. Restricts results to those with the specified bases, case-insensitive, allowing for circular or reverse complement matches. Does not allow partial matching or loose matching via degenerate bases. 
func (r ApiListDNAOligosRequest) Bases(bases string) ApiListDNAOligosRequest {
	r.bases = &bases
	return r
}

// ID of a folder. Restricts results to those in the folder.
func (r ApiListDNAOligosRequest) FolderId(folderId string) ApiListDNAOligosRequest {
	r.folderId = &folderId
	return r
}

// Comma-separated list of entry IDs. Restricts results to DNA Oligos mentioned in those entries. 
func (r ApiListDNAOligosRequest) MentionedIn(mentionedIn string) ApiListDNAOligosRequest {
	r.mentionedIn = &mentionedIn
	return r
}

// ID of a project. Restricts results to those in the project.
func (r ApiListDNAOligosRequest) ProjectId(projectId string) ApiListDNAOligosRequest {
	r.projectId = &projectId
	return r
}

// ID of a registry. Restricts results to those registered in this registry. Specifying \&quot;null\&quot; returns unregistered items. 
func (r ApiListDNAOligosRequest) RegistryId(registryId string) ApiListDNAOligosRequest {
	r.registryId = &registryId
	return r
}

// ID of a schema. Restricts results to those of the specified schema. 
func (r ApiListDNAOligosRequest) SchemaId(schemaId string) ApiListDNAOligosRequest {
	r.schemaId = &schemaId
	return r
}

// Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter. 
func (r ApiListDNAOligosRequest) SchemaFields(schemaFields map[string]interface{}) ApiListDNAOligosRequest {
	r.schemaFields = &schemaFields
	return r
}

// Archive reason. Restricts items to those with the specified archive reason. Use \&quot;NOT_ARCHIVED\&quot; to filter for unarchived DNA Oligos. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived DNA Oligos regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived. 
func (r ApiListDNAOligosRequest) ArchiveReason(archiveReason string) ApiListDNAOligosRequest {
	r.archiveReason = &archiveReason
	return r
}

// Comma-separated list of item IDs. Restricts results to those that mention the given items in the description. 
func (r ApiListDNAOligosRequest) Mentions(mentions string) ApiListDNAOligosRequest {
	r.mentions = &mentions
	return r
}

// Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid. 
func (r ApiListDNAOligosRequest) Ids(ids string) ApiListDNAOligosRequest {
	r.ids = &ids
	return r
}

// Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs. 
func (r ApiListDNAOligosRequest) EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf string) ApiListDNAOligosRequest {
	r.entityRegistryIdsAnyOf = &entityRegistryIdsAnyOf
	return r
}

// Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case insensitive.  Warning - this filter can be non-performant due to case insensitivity. 
func (r ApiListDNAOligosRequest) NamesAnyOf(namesAnyOf string) ApiListDNAOligosRequest {
	r.namesAnyOf = &namesAnyOf
	return r
}

// Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs, case sensitive. 
func (r ApiListDNAOligosRequest) NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive string) ApiListDNAOligosRequest {
	r.namesAnyOfCaseSensitive = &namesAnyOfCaseSensitive
	return r
}

// Comma separated list of users IDs
func (r ApiListDNAOligosRequest) CreatorIds(creatorIds string) ApiListDNAOligosRequest {
	r.creatorIds = &creatorIds
	return r
}

// Comma separated list of user or app IDs
func (r ApiListDNAOligosRequest) AuthorIdsAnyOf(authorIdsAnyOf string) ApiListDNAOligosRequest {
	r.authorIdsAnyOf = &authorIdsAnyOf
	return r
}

func (r ApiListDNAOligosRequest) Execute() (*DnaOligosPaginatedList, *http.Response, error) {
	return r.ApiService.ListDNAOligosExecute(r)
}

/*
ListDNAOligos List DNA Oligos

List DNA Oligos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDNAOligosRequest
*/
func (a *DNAOligosApiService) ListDNAOligos(ctx context.Context) ApiListDNAOligosRequest {
	return ApiListDNAOligosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DnaOligosPaginatedList
func (a *DNAOligosApiService) ListDNAOligosExecute(r ApiListDNAOligosRequest) (*DnaOligosPaginatedList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DnaOligosPaginatedList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNAOligosApiService.ListDNAOligos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dna-oligos"

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
	if r.nameIncludes != nil {
		localVarQueryParams.Add("nameIncludes", parameterToString(*r.nameIncludes, ""))
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
	if r.authorIdsAnyOf != nil {
		localVarQueryParams.Add("authorIds.anyOf", parameterToString(*r.authorIdsAnyOf, ""))
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

type ApiUnarchiveDNAOligosRequest struct {
	ctx context.Context
	ApiService *DNAOligosApiService
	dnaOligosUnarchive *DnaOligosUnarchive
}

func (r ApiUnarchiveDNAOligosRequest) DnaOligosUnarchive(dnaOligosUnarchive DnaOligosUnarchive) ApiUnarchiveDNAOligosRequest {
	r.dnaOligosUnarchive = &dnaOligosUnarchive
	return r
}

func (r ApiUnarchiveDNAOligosRequest) Execute() (*DnaOligosArchivalChange, *http.Response, error) {
	return r.ApiService.UnarchiveDNAOligosExecute(r)
}

/*
UnarchiveDNAOligos Unarchive DNA Oligos

Unarchive DNA Oligos

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUnarchiveDNAOligosRequest
*/
func (a *DNAOligosApiService) UnarchiveDNAOligos(ctx context.Context) ApiUnarchiveDNAOligosRequest {
	return ApiUnarchiveDNAOligosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DnaOligosArchivalChange
func (a *DNAOligosApiService) UnarchiveDNAOligosExecute(r ApiUnarchiveDNAOligosRequest) (*DnaOligosArchivalChange, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DnaOligosArchivalChange
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNAOligosApiService.UnarchiveDNAOligos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dna-oligos:unarchive"

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
	localVarPostBody = r.dnaOligosUnarchive
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

type ApiUpdateDNAOligoRequest struct {
	ctx context.Context
	ApiService *DNAOligosApiService
	oligoId string
	dnaOligoUpdate *DnaOligoUpdate
}

func (r ApiUpdateDNAOligoRequest) DnaOligoUpdate(dnaOligoUpdate DnaOligoUpdate) ApiUpdateDNAOligoRequest {
	r.dnaOligoUpdate = &dnaOligoUpdate
	return r
}

func (r ApiUpdateDNAOligoRequest) Execute() (*DnaOligo, *http.Response, error) {
	return r.ApiService.UpdateDNAOligoExecute(r)
}

/*
UpdateDNAOligo Update a DNA Oligo

Update a DNA Oligo

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param oligoId
 @return ApiUpdateDNAOligoRequest
*/
func (a *DNAOligosApiService) UpdateDNAOligo(ctx context.Context, oligoId string) ApiUpdateDNAOligoRequest {
	return ApiUpdateDNAOligoRequest{
		ApiService: a,
		ctx: ctx,
		oligoId: oligoId,
	}
}

// Execute executes the request
//  @return DnaOligo
func (a *DNAOligosApiService) UpdateDNAOligoExecute(r ApiUpdateDNAOligoRequest) (*DnaOligo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DnaOligo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNAOligosApiService.UpdateDNAOligo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dna-oligos/{oligo_id}"
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
	localVarPostBody = r.dnaOligoUpdate
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