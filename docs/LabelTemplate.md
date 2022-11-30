# LabelTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the label template. | [optional] 
**Name** | Pointer to **string** | Name of the label template. | [optional] 
**ZplTemplate** | Pointer to **string** | The ZPL template that will be filled in and sent to a printer. | [optional] 

## Methods

### NewLabelTemplate

`func NewLabelTemplate() *LabelTemplate`

NewLabelTemplate instantiates a new LabelTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelTemplateWithDefaults

`func NewLabelTemplateWithDefaults() *LabelTemplate`

NewLabelTemplateWithDefaults instantiates a new LabelTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LabelTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LabelTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LabelTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LabelTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LabelTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LabelTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LabelTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LabelTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetZplTemplate

`func (o *LabelTemplate) GetZplTemplate() string`

GetZplTemplate returns the ZplTemplate field if non-nil, zero value otherwise.

### GetZplTemplateOk

`func (o *LabelTemplate) GetZplTemplateOk() (*string, bool)`

GetZplTemplateOk returns a tuple with the ZplTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZplTemplate

`func (o *LabelTemplate) SetZplTemplate(v string)`

SetZplTemplate sets ZplTemplate field to given value.

### HasZplTemplate

`func (o *LabelTemplate) HasZplTemplate() bool`

HasZplTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


