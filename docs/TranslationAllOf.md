# TranslationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AminoAcids** | Pointer to **string** |  | [optional] [readonly] 
**End** | Pointer to **int32** |  | [optional] 
**GeneticCode** | Pointer to **string** | The genetic code to use when translating the nucleotide sequence into amino acids. | [optional] 
**Regions** | Pointer to [**[]TranslationAllOfRegions**](TranslationAllOfRegions.md) |  | [optional] 
**Start** | Pointer to **int32** |  | [optional] 
**Strand** | Pointer to **int32** |  | [optional] 

## Methods

### NewTranslationAllOf

`func NewTranslationAllOf() *TranslationAllOf`

NewTranslationAllOf instantiates a new TranslationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranslationAllOfWithDefaults

`func NewTranslationAllOfWithDefaults() *TranslationAllOf`

NewTranslationAllOfWithDefaults instantiates a new TranslationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAminoAcids

`func (o *TranslationAllOf) GetAminoAcids() string`

GetAminoAcids returns the AminoAcids field if non-nil, zero value otherwise.

### GetAminoAcidsOk

`func (o *TranslationAllOf) GetAminoAcidsOk() (*string, bool)`

GetAminoAcidsOk returns a tuple with the AminoAcids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcids

`func (o *TranslationAllOf) SetAminoAcids(v string)`

SetAminoAcids sets AminoAcids field to given value.

### HasAminoAcids

`func (o *TranslationAllOf) HasAminoAcids() bool`

HasAminoAcids returns a boolean if a field has been set.

### GetEnd

`func (o *TranslationAllOf) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TranslationAllOf) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TranslationAllOf) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *TranslationAllOf) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetGeneticCode

`func (o *TranslationAllOf) GetGeneticCode() string`

GetGeneticCode returns the GeneticCode field if non-nil, zero value otherwise.

### GetGeneticCodeOk

`func (o *TranslationAllOf) GetGeneticCodeOk() (*string, bool)`

GetGeneticCodeOk returns a tuple with the GeneticCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneticCode

`func (o *TranslationAllOf) SetGeneticCode(v string)`

SetGeneticCode sets GeneticCode field to given value.

### HasGeneticCode

`func (o *TranslationAllOf) HasGeneticCode() bool`

HasGeneticCode returns a boolean if a field has been set.

### GetRegions

`func (o *TranslationAllOf) GetRegions() []TranslationAllOfRegions`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *TranslationAllOf) GetRegionsOk() (*[]TranslationAllOfRegions, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *TranslationAllOf) SetRegions(v []TranslationAllOfRegions)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *TranslationAllOf) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetStart

`func (o *TranslationAllOf) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TranslationAllOf) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TranslationAllOf) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *TranslationAllOf) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStrand

`func (o *TranslationAllOf) GetStrand() int32`

GetStrand returns the Strand field if non-nil, zero value otherwise.

### GetStrandOk

`func (o *TranslationAllOf) GetStrandOk() (*int32, bool)`

GetStrandOk returns a tuple with the Strand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrand

`func (o *TranslationAllOf) SetStrand(v int32)`

SetStrand sets Strand field to given value.

### HasStrand

`func (o *TranslationAllOf) HasStrand() bool`

HasStrand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


