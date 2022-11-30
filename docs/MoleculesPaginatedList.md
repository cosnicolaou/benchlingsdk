# MoleculesPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**Molecules** | Pointer to [**[]Molecule**](Molecule.md) |  | [optional] 

## Methods

### NewMoleculesPaginatedList

`func NewMoleculesPaginatedList() *MoleculesPaginatedList`

NewMoleculesPaginatedList instantiates a new MoleculesPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoleculesPaginatedListWithDefaults

`func NewMoleculesPaginatedListWithDefaults() *MoleculesPaginatedList`

NewMoleculesPaginatedListWithDefaults instantiates a new MoleculesPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *MoleculesPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *MoleculesPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *MoleculesPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *MoleculesPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetMolecules

`func (o *MoleculesPaginatedList) GetMolecules() []Molecule`

GetMolecules returns the Molecules field if non-nil, zero value otherwise.

### GetMoleculesOk

`func (o *MoleculesPaginatedList) GetMoleculesOk() (*[]Molecule, bool)`

GetMoleculesOk returns a tuple with the Molecules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMolecules

`func (o *MoleculesPaginatedList) SetMolecules(v []Molecule)`

SetMolecules sets Molecules field to given value.

### HasMolecules

`func (o *MoleculesPaginatedList) HasMolecules() bool`

HasMolecules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


