# PrintLabels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerIds** | **[]string** | List of IDs of containers that will have labels printed (one label will be printed per container).  | 
**LabelTemplateId** | **string** | ID of label template to use (same template will be used for all labels printed).  | 
**PrinterId** | **string** | ID of printer to use (same printer will be used for all labels printed).  | 

## Methods

### NewPrintLabels

`func NewPrintLabels(containerIds []string, labelTemplateId string, printerId string, ) *PrintLabels`

NewPrintLabels instantiates a new PrintLabels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrintLabelsWithDefaults

`func NewPrintLabelsWithDefaults() *PrintLabels`

NewPrintLabelsWithDefaults instantiates a new PrintLabels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerIds

`func (o *PrintLabels) GetContainerIds() []string`

GetContainerIds returns the ContainerIds field if non-nil, zero value otherwise.

### GetContainerIdsOk

`func (o *PrintLabels) GetContainerIdsOk() (*[]string, bool)`

GetContainerIdsOk returns a tuple with the ContainerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerIds

`func (o *PrintLabels) SetContainerIds(v []string)`

SetContainerIds sets ContainerIds field to given value.


### GetLabelTemplateId

`func (o *PrintLabels) GetLabelTemplateId() string`

GetLabelTemplateId returns the LabelTemplateId field if non-nil, zero value otherwise.

### GetLabelTemplateIdOk

`func (o *PrintLabels) GetLabelTemplateIdOk() (*string, bool)`

GetLabelTemplateIdOk returns a tuple with the LabelTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelTemplateId

`func (o *PrintLabels) SetLabelTemplateId(v string)`

SetLabelTemplateId sets LabelTemplateId field to given value.


### GetPrinterId

`func (o *PrintLabels) GetPrinterId() string`

GetPrinterId returns the PrinterId field if non-nil, zero value otherwise.

### GetPrinterIdOk

`func (o *PrintLabels) GetPrinterIdOk() (*string, bool)`

GetPrinterIdOk returns a tuple with the PrinterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinterId

`func (o *PrintLabels) SetPrinterId(v string)`

SetPrinterId sets PrinterId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


