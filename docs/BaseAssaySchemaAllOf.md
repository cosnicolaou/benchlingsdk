# BaseAssaySchemaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DerivedFrom** | Pointer to **NullableString** |  | [optional] 
**Organization** | Pointer to [**BaseAssaySchemaAllOfOrganization**](BaseAssaySchemaAllOfOrganization.md) |  | [optional] 
**SystemName** | Pointer to **string** |  | [optional] 

## Methods

### NewBaseAssaySchemaAllOf

`func NewBaseAssaySchemaAllOf() *BaseAssaySchemaAllOf`

NewBaseAssaySchemaAllOf instantiates a new BaseAssaySchemaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseAssaySchemaAllOfWithDefaults

`func NewBaseAssaySchemaAllOfWithDefaults() *BaseAssaySchemaAllOf`

NewBaseAssaySchemaAllOfWithDefaults instantiates a new BaseAssaySchemaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDerivedFrom

`func (o *BaseAssaySchemaAllOf) GetDerivedFrom() string`

GetDerivedFrom returns the DerivedFrom field if non-nil, zero value otherwise.

### GetDerivedFromOk

`func (o *BaseAssaySchemaAllOf) GetDerivedFromOk() (*string, bool)`

GetDerivedFromOk returns a tuple with the DerivedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedFrom

`func (o *BaseAssaySchemaAllOf) SetDerivedFrom(v string)`

SetDerivedFrom sets DerivedFrom field to given value.

### HasDerivedFrom

`func (o *BaseAssaySchemaAllOf) HasDerivedFrom() bool`

HasDerivedFrom returns a boolean if a field has been set.

### SetDerivedFromNil

`func (o *BaseAssaySchemaAllOf) SetDerivedFromNil(b bool)`

 SetDerivedFromNil sets the value for DerivedFrom to be an explicit nil

### UnsetDerivedFrom
`func (o *BaseAssaySchemaAllOf) UnsetDerivedFrom()`

UnsetDerivedFrom ensures that no value is present for DerivedFrom, not even an explicit nil
### GetOrganization

`func (o *BaseAssaySchemaAllOf) GetOrganization() BaseAssaySchemaAllOfOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BaseAssaySchemaAllOf) GetOrganizationOk() (*BaseAssaySchemaAllOfOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BaseAssaySchemaAllOf) SetOrganization(v BaseAssaySchemaAllOfOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BaseAssaySchemaAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSystemName

`func (o *BaseAssaySchemaAllOf) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *BaseAssaySchemaAllOf) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *BaseAssaySchemaAllOf) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *BaseAssaySchemaAllOf) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


