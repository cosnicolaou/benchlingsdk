# RnaAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** | Hex color code used when displaying this feature in the UI. | [optional] 
**CustomFields** | Pointer to [**[]SequenceFeatureCustomField**](SequenceFeatureCustomField.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**End** | Pointer to **int32** |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Strand** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewRnaAnnotation

`func NewRnaAnnotation() *RnaAnnotation`

NewRnaAnnotation instantiates a new RnaAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRnaAnnotationWithDefaults

`func NewRnaAnnotationWithDefaults() *RnaAnnotation`

NewRnaAnnotationWithDefaults instantiates a new RnaAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *RnaAnnotation) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *RnaAnnotation) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *RnaAnnotation) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *RnaAnnotation) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCustomFields

`func (o *RnaAnnotation) GetCustomFields() []SequenceFeatureCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RnaAnnotation) GetCustomFieldsOk() (*[]SequenceFeatureCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RnaAnnotation) SetCustomFields(v []SequenceFeatureCustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RnaAnnotation) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetName

`func (o *RnaAnnotation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RnaAnnotation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RnaAnnotation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RnaAnnotation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *RnaAnnotation) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *RnaAnnotation) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *RnaAnnotation) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *RnaAnnotation) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetEnd

`func (o *RnaAnnotation) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *RnaAnnotation) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *RnaAnnotation) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *RnaAnnotation) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *RnaAnnotation) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *RnaAnnotation) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *RnaAnnotation) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *RnaAnnotation) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStrand

`func (o *RnaAnnotation) GetStrand() int32`

GetStrand returns the Strand field if non-nil, zero value otherwise.

### GetStrandOk

`func (o *RnaAnnotation) GetStrandOk() (*int32, bool)`

GetStrandOk returns a tuple with the Strand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrand

`func (o *RnaAnnotation) SetStrand(v int32)`

SetStrand sets Strand field to given value.

### HasStrand

`func (o *RnaAnnotation) HasStrand() bool`

HasStrand returns a boolean if a field has been set.

### GetType

`func (o *RnaAnnotation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RnaAnnotation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RnaAnnotation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RnaAnnotation) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


