# AssayRunsBulkCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssayRuns** | [**[]AssayRunCreate**](AssayRunCreate.md) |  | 

## Methods

### NewAssayRunsBulkCreateRequest

`func NewAssayRunsBulkCreateRequest(assayRuns []AssayRunCreate, ) *AssayRunsBulkCreateRequest`

NewAssayRunsBulkCreateRequest instantiates a new AssayRunsBulkCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayRunsBulkCreateRequestWithDefaults

`func NewAssayRunsBulkCreateRequestWithDefaults() *AssayRunsBulkCreateRequest`

NewAssayRunsBulkCreateRequestWithDefaults instantiates a new AssayRunsBulkCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssayRuns

`func (o *AssayRunsBulkCreateRequest) GetAssayRuns() []AssayRunCreate`

GetAssayRuns returns the AssayRuns field if non-nil, zero value otherwise.

### GetAssayRunsOk

`func (o *AssayRunsBulkCreateRequest) GetAssayRunsOk() (*[]AssayRunCreate, bool)`

GetAssayRunsOk returns a tuple with the AssayRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssayRuns

`func (o *AssayRunsBulkCreateRequest) SetAssayRuns(v []AssayRunCreate)`

SetAssayRuns sets AssayRuns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


