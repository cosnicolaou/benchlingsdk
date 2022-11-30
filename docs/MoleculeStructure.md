# MoleculeStructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StructureFormat** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** | Chemical structure in SMILES or molfile format. | [optional] 

## Methods

### NewMoleculeStructure

`func NewMoleculeStructure() *MoleculeStructure`

NewMoleculeStructure instantiates a new MoleculeStructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoleculeStructureWithDefaults

`func NewMoleculeStructureWithDefaults() *MoleculeStructure`

NewMoleculeStructureWithDefaults instantiates a new MoleculeStructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStructureFormat

`func (o *MoleculeStructure) GetStructureFormat() string`

GetStructureFormat returns the StructureFormat field if non-nil, zero value otherwise.

### GetStructureFormatOk

`func (o *MoleculeStructure) GetStructureFormatOk() (*string, bool)`

GetStructureFormatOk returns a tuple with the StructureFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureFormat

`func (o *MoleculeStructure) SetStructureFormat(v string)`

SetStructureFormat sets StructureFormat field to given value.

### HasStructureFormat

`func (o *MoleculeStructure) HasStructureFormat() bool`

HasStructureFormat returns a boolean if a field has been set.

### GetValue

`func (o *MoleculeStructure) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MoleculeStructure) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MoleculeStructure) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MoleculeStructure) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


