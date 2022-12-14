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


// DNAAlignmentsApiService DNAAlignmentsApi service
type DNAAlignmentsApiService service

type ApiCreateDnaConsensusAlignmentRequest struct {
	ctx context.Context
	ApiService *DNAAlignmentsApiService
	dnaConsensusAlignmentCreate *DnaConsensusAlignmentCreate
}

func (r ApiCreateDnaConsensusAlignmentRequest) DnaConsensusAlignmentCreate(dnaConsensusAlignmentCreate DnaConsensusAlignmentCreate) ApiCreateDnaConsensusAlignmentRequest {
	r.dnaConsensusAlignmentCreate = &dnaConsensusAlignmentCreate
	return r
}

func (r ApiCreateDnaConsensusAlignmentRequest) Execute() (*AsyncTaskLink, *http.Response, error) {
	return r.ApiService.CreateDnaConsensusAlignmentExecute(r)
}

/*
CreateDnaConsensusAlignment Create a consensus DNA alignment

Create a consensus DNA alignment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDnaConsensusAlignmentRequest

Deprecated
*/
func (a *DNAAlignmentsApiService) CreateDnaConsensusAlignment(ctx context.Context) ApiCreateDnaConsensusAlignmentRequest {
	return ApiCreateDnaConsensusAlignmentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncTaskLink
// Deprecated
func (a *DNAAlignmentsApiService) CreateDnaConsensusAlignmentExecute(r ApiCreateDnaConsensusAlignmentRequest) (*AsyncTaskLink, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncTaskLink
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNAAlignmentsApiService.CreateDnaConsensusAlignment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dna-alignments:create-consensus-alignment"

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
	localVarPostBody = r.dnaConsensusAlignmentCreate
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

type ApiCreateDnaTemplateAlignmentRequest struct {
	ctx context.Context
	ApiService *DNAAlignmentsApiService
	dnaTemplateAlignmentCreate *DnaTemplateAlignmentCreate
}

func (r ApiCreateDnaTemplateAlignmentRequest) DnaTemplateAlignmentCreate(dnaTemplateAlignmentCreate DnaTemplateAlignmentCreate) ApiCreateDnaTemplateAlignmentRequest {
	r.dnaTemplateAlignmentCreate = &dnaTemplateAlignmentCreate
	return r
}

func (r ApiCreateDnaTemplateAlignmentRequest) Execute() (*AsyncTaskLink, *http.Response, error) {
	return r.ApiService.CreateDnaTemplateAlignmentExecute(r)
}

/*
CreateDnaTemplateAlignment Create a template DNA alignment

Create a template DNA alignment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDnaTemplateAlignmentRequest

Deprecated
*/
func (a *DNAAlignmentsApiService) CreateDnaTemplateAlignment(ctx context.Context) ApiCreateDnaTemplateAlignmentRequest {
	return ApiCreateDnaTemplateAlignmentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncTaskLink
// Deprecated
func (a *DNAAlignmentsApiService) CreateDnaTemplateAlignmentExecute(r ApiCreateDnaTemplateAlignmentRequest) (*AsyncTaskLink, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncTaskLink
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNAAlignmentsApiService.CreateDnaTemplateAlignment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dna-alignments:create-template-alignment"

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
	localVarPostBody = r.dnaTemplateAlignmentCreate
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

type ApiDeleteDNAAlignmentRequest struct {
	ctx context.Context
	ApiService *DNAAlignmentsApiService
	dnaAlignmentId string
}

func (r ApiDeleteDNAAlignmentRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DeleteDNAAlignmentExecute(r)
}

/*
DeleteDNAAlignment Delete a DNA Alignment

Delete a DNA Alignment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dnaAlignmentId
 @return ApiDeleteDNAAlignmentRequest

Deprecated
*/
func (a *DNAAlignmentsApiService) DeleteDNAAlignment(ctx context.Context, dnaAlignmentId string) ApiDeleteDNAAlignmentRequest {
	return ApiDeleteDNAAlignmentRequest{
		ApiService: a,
		ctx: ctx,
		dnaAlignmentId: dnaAlignmentId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
// Deprecated
func (a *DNAAlignmentsApiService) DeleteDNAAlignmentExecute(r ApiDeleteDNAAlignmentRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNAAlignmentsApiService.DeleteDNAAlignment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dna-alignments/{dna_alignment_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dna_alignment_id"+"}", url.PathEscape(parameterToString(r.dnaAlignmentId, "")), -1)

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

type ApiGetDNAAlignmentRequest struct {
	ctx context.Context
	ApiService *DNAAlignmentsApiService
	dnaAlignmentId string
}

func (r ApiGetDNAAlignmentRequest) Execute() (*DnaAlignment, *http.Response, error) {
	return r.ApiService.GetDNAAlignmentExecute(r)
}

/*
GetDNAAlignment Get a DNA Alignment

Get a DNA Alignment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dnaAlignmentId
 @return ApiGetDNAAlignmentRequest

Deprecated
*/
func (a *DNAAlignmentsApiService) GetDNAAlignment(ctx context.Context, dnaAlignmentId string) ApiGetDNAAlignmentRequest {
	return ApiGetDNAAlignmentRequest{
		ApiService: a,
		ctx: ctx,
		dnaAlignmentId: dnaAlignmentId,
	}
}

// Execute executes the request
//  @return DnaAlignment
// Deprecated
func (a *DNAAlignmentsApiService) GetDNAAlignmentExecute(r ApiGetDNAAlignmentRequest) (*DnaAlignment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DnaAlignment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNAAlignmentsApiService.GetDNAAlignment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dna-alignments/{dna_alignment_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dna_alignment_id"+"}", url.PathEscape(parameterToString(r.dnaAlignmentId, "")), -1)

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

type ApiListDNAAlignmentsRequest struct {
	ctx context.Context
	ApiService *DNAAlignmentsApiService
	pageSize *int32
	nextToken *string
	sort *string
	modifiedAt *string
	name *string
	nameIncludes *string
	ids *string
	namesAnyOf *string
	namesAnyOfCaseSensitive *string
	sequenceIds *string
}

// Number of results to return. Defaults to 50, maximum of 100. 
func (r ApiListDNAAlignmentsRequest) PageSize(pageSize int32) ApiListDNAAlignmentsRequest {
	r.pageSize = &pageSize
	return r
}

// Token for pagination
func (r ApiListDNAAlignmentsRequest) NextToken(nextToken string) ApiListDNAAlignmentsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiListDNAAlignmentsRequest) Sort(sort string) ApiListDNAAlignmentsRequest {
	r.sort = &sort
	return r
}

// Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30. 
func (r ApiListDNAAlignmentsRequest) ModifiedAt(modifiedAt string) ApiListDNAAlignmentsRequest {
	r.modifiedAt = &modifiedAt
	return r
}

// Name of a DNA Alignment. Restricts results to those with the specified name.
func (r ApiListDNAAlignmentsRequest) Name(name string) ApiListDNAAlignmentsRequest {
	r.name = &name
	return r
}

// Name substring of a DNA Alignment. Restricts results to those with names that include the provided substring.
func (r ApiListDNAAlignmentsRequest) NameIncludes(nameIncludes string) ApiListDNAAlignmentsRequest {
	r.nameIncludes = &nameIncludes
	return r
}

// Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid. 
func (r ApiListDNAAlignmentsRequest) Ids(ids string) ApiListDNAAlignmentsRequest {
	r.ids = &ids
	return r
}

// Comma-separated list of names. Restricts results to those that match any of the specified names, case insensitive.  Warning - this filter can be non-performant due to case insensitivity. 
func (r ApiListDNAAlignmentsRequest) NamesAnyOf(namesAnyOf string) ApiListDNAAlignmentsRequest {
	r.namesAnyOf = &namesAnyOf
	return r
}

// Comma-separated list of names. Restricts results to those that match any of the specified names, case sensitive. 
func (r ApiListDNAAlignmentsRequest) NamesAnyOfCaseSensitive(namesAnyOfCaseSensitive string) ApiListDNAAlignmentsRequest {
	r.namesAnyOfCaseSensitive = &namesAnyOfCaseSensitive
	return r
}

// Comma-separated list of sequence ids that own one or more DNA Alignments (i.e. ids of sequences used as the template in a Template Alignment or created as the consensus sequence from a Consensus Alignment). Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid. 
func (r ApiListDNAAlignmentsRequest) SequenceIds(sequenceIds string) ApiListDNAAlignmentsRequest {
	r.sequenceIds = &sequenceIds
	return r
}

func (r ApiListDNAAlignmentsRequest) Execute() (*DnaAlignmentsPaginatedList, *http.Response, error) {
	return r.ApiService.ListDNAAlignmentsExecute(r)
}

/*
ListDNAAlignments List DNA Alignments

List DNA Alignments

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDNAAlignmentsRequest

Deprecated
*/
func (a *DNAAlignmentsApiService) ListDNAAlignments(ctx context.Context) ApiListDNAAlignmentsRequest {
	return ApiListDNAAlignmentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DnaAlignmentsPaginatedList
// Deprecated
func (a *DNAAlignmentsApiService) ListDNAAlignmentsExecute(r ApiListDNAAlignmentsRequest) (*DnaAlignmentsPaginatedList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DnaAlignmentsPaginatedList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNAAlignmentsApiService.ListDNAAlignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dna-alignments"

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
	if r.ids != nil {
		localVarQueryParams.Add("ids", parameterToString(*r.ids, ""))
	}
	if r.namesAnyOf != nil {
		localVarQueryParams.Add("names.anyOf", parameterToString(*r.namesAnyOf, ""))
	}
	if r.namesAnyOfCaseSensitive != nil {
		localVarQueryParams.Add("names.anyOf.caseSensitive", parameterToString(*r.namesAnyOfCaseSensitive, ""))
	}
	if r.sequenceIds != nil {
		localVarQueryParams.Add("sequenceIds", parameterToString(*r.sequenceIds, ""))
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
