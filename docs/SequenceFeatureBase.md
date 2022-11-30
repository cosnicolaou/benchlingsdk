# SequenceFeatureBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** | Hex color code used when displaying this feature in the UI. | [optional] 
**CustomFields** | Pointer to [**[]SequenceFeatureCustomField**](SequenceFeatureCustomField.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 

## Methods

### NewSequenceFeatureBase

`func NewSequenceFeatureBase() *SequenceFeatureBase`

NewSequenceFeatureBase instantiates a new SequenceFeatureBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSequenceFeatureBaseWithDefaults

`func NewSequenceFeatureBaseWithDefaults() *SequenceFeatureBase`

NewSequenceFeatureBaseWithDefaults instantiates a new SequenceFeatureBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *SequenceFeatureBase) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *SequenceFeatureBase) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *SequenceFeatureBase) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *SequenceFeatureBase) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCustomFields

`func (o *SequenceFeatureBase) GetCustomFields() []SequenceFeatureCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SequenceFeatureBase) GetCustomFieldsOk() (*[]SequenceFeatureCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SequenceFeatureBase) SetCustomFields(v []SequenceFeatureCustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SequenceFeatureBase) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetName

`func (o *SequenceFeatureBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SequenceFeatureBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SequenceFeatureBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SequenceFeatureBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *SequenceFeatureBase) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SequenceFeatureBase) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SequenceFeatureBase) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SequenceFeatureBase) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


