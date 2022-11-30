# IngredientWriteParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **NullableString** | The amount value of this ingredient in its mixture, in string format (to preserve full precision). Pair with &#x60;units&#x60;. Supports scientific notation (1.23e4). One ingredient on this mixture can have an amount value of &#x60;\&quot;QS\&quot;&#x60;.  | 
**CatalogIdentifier** | **NullableString** |  | 
**ComponentEntityId** | **string** | The entity that uniquely identifies this component. | 
**ComponentLotContainerId** | **NullableString** | The container representing the component lot for this ingredient. This is only writable if the mixture schema supports component lots in \&quot;storage\&quot; format. | 
**ComponentLotEntityId** | **NullableString** | The entity representing the component lot for this ingredient. This is only writable if the mixture schema supports component lots in \&quot;storage\&quot; format. | 
**ComponentLotText** | **NullableString** | Text representing the component lot for this ingredient. This is only writable if the mixture schema supports component lots in \&quot;text\&quot; format. | 
**Notes** | **NullableString** |  | 
**Units** | [**NullableIngredientMeasurementUnits**](IngredientMeasurementUnits.md) |  | 

## Methods

### NewIngredientWriteParams

`func NewIngredientWriteParams(amount NullableString, catalogIdentifier NullableString, componentEntityId string, componentLotContainerId NullableString, componentLotEntityId NullableString, componentLotText NullableString, notes NullableString, units NullableIngredientMeasurementUnits, ) *IngredientWriteParams`

NewIngredientWriteParams instantiates a new IngredientWriteParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngredientWriteParamsWithDefaults

`func NewIngredientWriteParamsWithDefaults() *IngredientWriteParams`

NewIngredientWriteParamsWithDefaults instantiates a new IngredientWriteParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *IngredientWriteParams) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *IngredientWriteParams) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *IngredientWriteParams) SetAmount(v string)`

SetAmount sets Amount field to given value.


### SetAmountNil

`func (o *IngredientWriteParams) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *IngredientWriteParams) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCatalogIdentifier

`func (o *IngredientWriteParams) GetCatalogIdentifier() string`

GetCatalogIdentifier returns the CatalogIdentifier field if non-nil, zero value otherwise.

### GetCatalogIdentifierOk

`func (o *IngredientWriteParams) GetCatalogIdentifierOk() (*string, bool)`

GetCatalogIdentifierOk returns a tuple with the CatalogIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogIdentifier

`func (o *IngredientWriteParams) SetCatalogIdentifier(v string)`

SetCatalogIdentifier sets CatalogIdentifier field to given value.


### SetCatalogIdentifierNil

`func (o *IngredientWriteParams) SetCatalogIdentifierNil(b bool)`

 SetCatalogIdentifierNil sets the value for CatalogIdentifier to be an explicit nil

### UnsetCatalogIdentifier
`func (o *IngredientWriteParams) UnsetCatalogIdentifier()`

UnsetCatalogIdentifier ensures that no value is present for CatalogIdentifier, not even an explicit nil
### GetComponentEntityId

`func (o *IngredientWriteParams) GetComponentEntityId() string`

GetComponentEntityId returns the ComponentEntityId field if non-nil, zero value otherwise.

### GetComponentEntityIdOk

`func (o *IngredientWriteParams) GetComponentEntityIdOk() (*string, bool)`

GetComponentEntityIdOk returns a tuple with the ComponentEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentEntityId

`func (o *IngredientWriteParams) SetComponentEntityId(v string)`

SetComponentEntityId sets ComponentEntityId field to given value.


### GetComponentLotContainerId

`func (o *IngredientWriteParams) GetComponentLotContainerId() string`

GetComponentLotContainerId returns the ComponentLotContainerId field if non-nil, zero value otherwise.

### GetComponentLotContainerIdOk

`func (o *IngredientWriteParams) GetComponentLotContainerIdOk() (*string, bool)`

GetComponentLotContainerIdOk returns a tuple with the ComponentLotContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLotContainerId

`func (o *IngredientWriteParams) SetComponentLotContainerId(v string)`

SetComponentLotContainerId sets ComponentLotContainerId field to given value.


### SetComponentLotContainerIdNil

`func (o *IngredientWriteParams) SetComponentLotContainerIdNil(b bool)`

 SetComponentLotContainerIdNil sets the value for ComponentLotContainerId to be an explicit nil

### UnsetComponentLotContainerId
`func (o *IngredientWriteParams) UnsetComponentLotContainerId()`

UnsetComponentLotContainerId ensures that no value is present for ComponentLotContainerId, not even an explicit nil
### GetComponentLotEntityId

`func (o *IngredientWriteParams) GetComponentLotEntityId() string`

GetComponentLotEntityId returns the ComponentLotEntityId field if non-nil, zero value otherwise.

### GetComponentLotEntityIdOk

`func (o *IngredientWriteParams) GetComponentLotEntityIdOk() (*string, bool)`

GetComponentLotEntityIdOk returns a tuple with the ComponentLotEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLotEntityId

`func (o *IngredientWriteParams) SetComponentLotEntityId(v string)`

SetComponentLotEntityId sets ComponentLotEntityId field to given value.


### SetComponentLotEntityIdNil

`func (o *IngredientWriteParams) SetComponentLotEntityIdNil(b bool)`

 SetComponentLotEntityIdNil sets the value for ComponentLotEntityId to be an explicit nil

### UnsetComponentLotEntityId
`func (o *IngredientWriteParams) UnsetComponentLotEntityId()`

UnsetComponentLotEntityId ensures that no value is present for ComponentLotEntityId, not even an explicit nil
### GetComponentLotText

`func (o *IngredientWriteParams) GetComponentLotText() string`

GetComponentLotText returns the ComponentLotText field if non-nil, zero value otherwise.

### GetComponentLotTextOk

`func (o *IngredientWriteParams) GetComponentLotTextOk() (*string, bool)`

GetComponentLotTextOk returns a tuple with the ComponentLotText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLotText

`func (o *IngredientWriteParams) SetComponentLotText(v string)`

SetComponentLotText sets ComponentLotText field to given value.


### SetComponentLotTextNil

`func (o *IngredientWriteParams) SetComponentLotTextNil(b bool)`

 SetComponentLotTextNil sets the value for ComponentLotText to be an explicit nil

### UnsetComponentLotText
`func (o *IngredientWriteParams) UnsetComponentLotText()`

UnsetComponentLotText ensures that no value is present for ComponentLotText, not even an explicit nil
### GetNotes

`func (o *IngredientWriteParams) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *IngredientWriteParams) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *IngredientWriteParams) SetNotes(v string)`

SetNotes sets Notes field to given value.


### SetNotesNil

`func (o *IngredientWriteParams) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *IngredientWriteParams) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetUnits

`func (o *IngredientWriteParams) GetUnits() IngredientMeasurementUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *IngredientWriteParams) GetUnitsOk() (*IngredientMeasurementUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *IngredientWriteParams) SetUnits(v IngredientMeasurementUnits)`

SetUnits sets Units field to given value.


### SetUnitsNil

`func (o *IngredientWriteParams) SetUnitsNil(b bool)`

 SetUnitsNil sets the value for Units to be an explicit nil

### UnsetUnits
`func (o *IngredientWriteParams) UnsetUnits()`

UnsetUnits ensures that no value is present for Units, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


