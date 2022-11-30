# CreateNucleotideTemplateAlignmentAsyncTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **map[string]interface{}** | Present only when status is FAILED for a bulk task. Contains information about the individual errors in the bulk task.  | [optional] 
**Message** | Pointer to **string** | Present only when status is FAILED. Contains information about the error. | [optional] 
**Response** | Pointer to [**NucleotideAlignment**](NucleotideAlignment.md) |  | [optional] 
**Status** | **string** | The current state of the task. | 

## Methods

### NewCreateNucleotideTemplateAlignmentAsyncTask

`func NewCreateNucleotideTemplateAlignmentAsyncTask(status string, ) *CreateNucleotideTemplateAlignmentAsyncTask`

NewCreateNucleotideTemplateAlignmentAsyncTask instantiates a new CreateNucleotideTemplateAlignmentAsyncTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNucleotideTemplateAlignmentAsyncTaskWithDefaults

`func NewCreateNucleotideTemplateAlignmentAsyncTaskWithDefaults() *CreateNucleotideTemplateAlignmentAsyncTask`

NewCreateNucleotideTemplateAlignmentAsyncTaskWithDefaults instantiates a new CreateNucleotideTemplateAlignmentAsyncTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateNucleotideTemplateAlignmentAsyncTask) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateNucleotideTemplateAlignmentAsyncTask) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMessage

`func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateNucleotideTemplateAlignmentAsyncTask) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateNucleotideTemplateAlignmentAsyncTask) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResponse

`func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetResponse() NucleotideAlignment`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetResponseOk() (*NucleotideAlignment, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *CreateNucleotideTemplateAlignmentAsyncTask) SetResponse(v NucleotideAlignment)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *CreateNucleotideTemplateAlignmentAsyncTask) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetStatus

`func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateNucleotideTemplateAlignmentAsyncTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateNucleotideTemplateAlignmentAsyncTask) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


