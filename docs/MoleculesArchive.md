# MoleculesArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoleculeIds** | **[]string** |  | 
**Reason** | **string** | The reason for archiving the provided Molecules. Accepted reasons may differ based on tenant configuration.  | 

## Methods

### NewMoleculesArchive

`func NewMoleculesArchive(moleculeIds []string, reason string, ) *MoleculesArchive`

NewMoleculesArchive instantiates a new MoleculesArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoleculesArchiveWithDefaults

`func NewMoleculesArchiveWithDefaults() *MoleculesArchive`

NewMoleculesArchiveWithDefaults instantiates a new MoleculesArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoleculeIds

`func (o *MoleculesArchive) GetMoleculeIds() []string`

GetMoleculeIds returns the MoleculeIds field if non-nil, zero value otherwise.

### GetMoleculeIdsOk

`func (o *MoleculesArchive) GetMoleculeIdsOk() (*[]string, bool)`

GetMoleculeIdsOk returns a tuple with the MoleculeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoleculeIds

`func (o *MoleculesArchive) SetMoleculeIds(v []string)`

SetMoleculeIds sets MoleculeIds field to given value.


### GetReason

`func (o *MoleculesArchive) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MoleculesArchive) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MoleculesArchive) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


