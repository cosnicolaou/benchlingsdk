# MoleculeBaseRequestChemicalStructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StructureFormat** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** | Chemical structure in SMILES or molfile format. | [optional] 

## Methods

### NewMoleculeBaseRequestChemicalStructure

`func NewMoleculeBaseRequestChemicalStructure() *MoleculeBaseRequestChemicalStructure`

NewMoleculeBaseRequestChemicalStructure instantiates a new MoleculeBaseRequestChemicalStructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoleculeBaseRequestChemicalStructureWithDefaults

`func NewMoleculeBaseRequestChemicalStructureWithDefaults() *MoleculeBaseRequestChemicalStructure`

NewMoleculeBaseRequestChemicalStructureWithDefaults instantiates a new MoleculeBaseRequestChemicalStructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStructureFormat

`func (o *MoleculeBaseRequestChemicalStructure) GetStructureFormat() string`

GetStructureFormat returns the StructureFormat field if non-nil, zero value otherwise.

### GetStructureFormatOk

`func (o *MoleculeBaseRequestChemicalStructure) GetStructureFormatOk() (*string, bool)`

GetStructureFormatOk returns a tuple with the StructureFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureFormat

`func (o *MoleculeBaseRequestChemicalStructure) SetStructureFormat(v string)`

SetStructureFormat sets StructureFormat field to given value.

### HasStructureFormat

`func (o *MoleculeBaseRequestChemicalStructure) HasStructureFormat() bool`

HasStructureFormat returns a boolean if a field has been set.

### GetValue

`func (o *MoleculeBaseRequestChemicalStructure) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MoleculeBaseRequestChemicalStructure) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MoleculeBaseRequestChemicalStructure) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MoleculeBaseRequestChemicalStructure) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


