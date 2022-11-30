# EntitySchemaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraint** | Pointer to [**NullableEntitySchemaAllOfConstraint**](EntitySchemaAllOfConstraint.md) |  | [optional] 
**ContainableType** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the Entity Schema was last modified | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewEntitySchemaAllOf

`func NewEntitySchemaAllOf() *EntitySchemaAllOf`

NewEntitySchemaAllOf instantiates a new EntitySchemaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySchemaAllOfWithDefaults

`func NewEntitySchemaAllOfWithDefaults() *EntitySchemaAllOf`

NewEntitySchemaAllOfWithDefaults instantiates a new EntitySchemaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraint

`func (o *EntitySchemaAllOf) GetConstraint() EntitySchemaAllOfConstraint`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *EntitySchemaAllOf) GetConstraintOk() (*EntitySchemaAllOfConstraint, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *EntitySchemaAllOf) SetConstraint(v EntitySchemaAllOfConstraint)`

SetConstraint sets Constraint field to given value.

### HasConstraint

`func (o *EntitySchemaAllOf) HasConstraint() bool`

HasConstraint returns a boolean if a field has been set.

### SetConstraintNil

`func (o *EntitySchemaAllOf) SetConstraintNil(b bool)`

 SetConstraintNil sets the value for Constraint to be an explicit nil

### UnsetConstraint
`func (o *EntitySchemaAllOf) UnsetConstraint()`

UnsetConstraint ensures that no value is present for Constraint, not even an explicit nil
### GetContainableType

`func (o *EntitySchemaAllOf) GetContainableType() string`

GetContainableType returns the ContainableType field if non-nil, zero value otherwise.

### GetContainableTypeOk

`func (o *EntitySchemaAllOf) GetContainableTypeOk() (*string, bool)`

GetContainableTypeOk returns a tuple with the ContainableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainableType

`func (o *EntitySchemaAllOf) SetContainableType(v string)`

SetContainableType sets ContainableType field to given value.

### HasContainableType

`func (o *EntitySchemaAllOf) HasContainableType() bool`

HasContainableType returns a boolean if a field has been set.

### GetModifiedAt

`func (o *EntitySchemaAllOf) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *EntitySchemaAllOf) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *EntitySchemaAllOf) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *EntitySchemaAllOf) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetType

`func (o *EntitySchemaAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitySchemaAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitySchemaAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntitySchemaAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


