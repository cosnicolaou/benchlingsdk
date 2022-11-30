# Translation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** | Hex color code used when displaying this feature in the UI. | [optional] 
**CustomFields** | Pointer to [**[]SequenceFeatureCustomField**](SequenceFeatureCustomField.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**AminoAcids** | Pointer to **string** |  | [optional] [readonly] 
**End** | Pointer to **int32** |  | [optional] 
**GeneticCode** | Pointer to **string** | The genetic code to use when translating the nucleotide sequence into amino acids. | [optional] 
**Regions** | Pointer to [**[]TranslationAllOfRegions**](TranslationAllOfRegions.md) |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Strand** | Pointer to **int32** |  | [optional] 

## Methods

### NewTranslation

`func NewTranslation() *Translation`

NewTranslation instantiates a new Translation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranslationWithDefaults

`func NewTranslationWithDefaults() *Translation`

NewTranslationWithDefaults instantiates a new Translation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *Translation) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Translation) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Translation) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Translation) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCustomFields

`func (o *Translation) GetCustomFields() []SequenceFeatureCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Translation) GetCustomFieldsOk() (*[]SequenceFeatureCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Translation) SetCustomFields(v []SequenceFeatureCustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Translation) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetName

`func (o *Translation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Translation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Translation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Translation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *Translation) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Translation) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Translation) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Translation) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAminoAcids

`func (o *Translation) GetAminoAcids() string`

GetAminoAcids returns the AminoAcids field if non-nil, zero value otherwise.

### GetAminoAcidsOk

`func (o *Translation) GetAminoAcidsOk() (*string, bool)`

GetAminoAcidsOk returns a tuple with the AminoAcids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcids

`func (o *Translation) SetAminoAcids(v string)`

SetAminoAcids sets AminoAcids field to given value.

### HasAminoAcids

`func (o *Translation) HasAminoAcids() bool`

HasAminoAcids returns a boolean if a field has been set.

### GetEnd

`func (o *Translation) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Translation) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Translation) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Translation) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetGeneticCode

`func (o *Translation) GetGeneticCode() string`

GetGeneticCode returns the GeneticCode field if non-nil, zero value otherwise.

### GetGeneticCodeOk

`func (o *Translation) GetGeneticCodeOk() (*string, bool)`

GetGeneticCodeOk returns a tuple with the GeneticCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneticCode

`func (o *Translation) SetGeneticCode(v string)`

SetGeneticCode sets GeneticCode field to given value.

### HasGeneticCode

`func (o *Translation) HasGeneticCode() bool`

HasGeneticCode returns a boolean if a field has been set.

### GetRegions

`func (o *Translation) GetRegions() []TranslationAllOfRegions`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Translation) GetRegionsOk() (*[]TranslationAllOfRegions, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Translation) SetRegions(v []TranslationAllOfRegions)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Translation) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetStart

`func (o *Translation) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Translation) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Translation) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *Translation) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStrand

`func (o *Translation) GetStrand() int32`

GetStrand returns the Strand field if non-nil, zero value otherwise.

### GetStrandOk

`func (o *Translation) GetStrandOk() (*int32, bool)`

GetStrandOk returns a tuple with the Strand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrand

`func (o *Translation) SetStrand(v int32)`

SetStrand sets Strand field to given value.

### HasStrand

`func (o *Translation) HasStrand() bool`

HasStrand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


