# Ingredient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableString** | The amount value of this ingredient in its mixture, in string format (to preserve full precision). Pair with &#x60;units&#x60;. Supports scientific notation (1.23e4). One ingredient on this mixture can have an amount value of &#x60;\&quot;QS\&quot;&#x60;.  | [optional] 
**CatalogIdentifier** | Pointer to **NullableString** |  | [optional] 
**ComponentEntity** | Pointer to [**IngredientComponentEntity**](IngredientComponentEntity.md) |  | [optional] 
**ComponentLotContainer** | Pointer to [**NullableIngredientComponentLotContainer**](IngredientComponentLotContainer.md) |  | [optional] 
**ComponentLotEntity** | Pointer to [**NullableIngredientComponentLotEntity**](IngredientComponentLotEntity.md) |  | [optional] 
**ComponentLotText** | Pointer to **NullableString** | Text representing the component lot for this ingredient. This is only present if the mixture schema supports component lots in \&quot;text\&quot; format. | [optional] 
**HasParent** | Pointer to **bool** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**TargetAmount** | Pointer to **NullableString** | The target amount for this ingredient such that this ingredient&#39;s proportion in its mixture would preserve the equivalent ingredient&#39;s proportion in the parent mixture. Pair with &#x60;units&#x60;. | [optional] [readonly] 
**Units** | Pointer to [**NullableIngredientMeasurementUnits**](IngredientMeasurementUnits.md) |  | [optional] 

## Methods

### NewIngredient

`func NewIngredient() *Ingredient`

NewIngredient instantiates a new Ingredient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngredientWithDefaults

`func NewIngredientWithDefaults() *Ingredient`

NewIngredientWithDefaults instantiates a new Ingredient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Ingredient) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Ingredient) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Ingredient) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Ingredient) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *Ingredient) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *Ingredient) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCatalogIdentifier

`func (o *Ingredient) GetCatalogIdentifier() string`

GetCatalogIdentifier returns the CatalogIdentifier field if non-nil, zero value otherwise.

### GetCatalogIdentifierOk

`func (o *Ingredient) GetCatalogIdentifierOk() (*string, bool)`

GetCatalogIdentifierOk returns a tuple with the CatalogIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogIdentifier

`func (o *Ingredient) SetCatalogIdentifier(v string)`

SetCatalogIdentifier sets CatalogIdentifier field to given value.

### HasCatalogIdentifier

`func (o *Ingredient) HasCatalogIdentifier() bool`

HasCatalogIdentifier returns a boolean if a field has been set.

### SetCatalogIdentifierNil

`func (o *Ingredient) SetCatalogIdentifierNil(b bool)`

 SetCatalogIdentifierNil sets the value for CatalogIdentifier to be an explicit nil

### UnsetCatalogIdentifier
`func (o *Ingredient) UnsetCatalogIdentifier()`

UnsetCatalogIdentifier ensures that no value is present for CatalogIdentifier, not even an explicit nil
### GetComponentEntity

`func (o *Ingredient) GetComponentEntity() IngredientComponentEntity`

GetComponentEntity returns the ComponentEntity field if non-nil, zero value otherwise.

### GetComponentEntityOk

`func (o *Ingredient) GetComponentEntityOk() (*IngredientComponentEntity, bool)`

GetComponentEntityOk returns a tuple with the ComponentEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentEntity

`func (o *Ingredient) SetComponentEntity(v IngredientComponentEntity)`

SetComponentEntity sets ComponentEntity field to given value.

### HasComponentEntity

`func (o *Ingredient) HasComponentEntity() bool`

HasComponentEntity returns a boolean if a field has been set.

### GetComponentLotContainer

`func (o *Ingredient) GetComponentLotContainer() IngredientComponentLotContainer`

GetComponentLotContainer returns the ComponentLotContainer field if non-nil, zero value otherwise.

### GetComponentLotContainerOk

`func (o *Ingredient) GetComponentLotContainerOk() (*IngredientComponentLotContainer, bool)`

GetComponentLotContainerOk returns a tuple with the ComponentLotContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLotContainer

`func (o *Ingredient) SetComponentLotContainer(v IngredientComponentLotContainer)`

SetComponentLotContainer sets ComponentLotContainer field to given value.

### HasComponentLotContainer

`func (o *Ingredient) HasComponentLotContainer() bool`

HasComponentLotContainer returns a boolean if a field has been set.

### SetComponentLotContainerNil

`func (o *Ingredient) SetComponentLotContainerNil(b bool)`

 SetComponentLotContainerNil sets the value for ComponentLotContainer to be an explicit nil

### UnsetComponentLotContainer
`func (o *Ingredient) UnsetComponentLotContainer()`

UnsetComponentLotContainer ensures that no value is present for ComponentLotContainer, not even an explicit nil
### GetComponentLotEntity

`func (o *Ingredient) GetComponentLotEntity() IngredientComponentLotEntity`

GetComponentLotEntity returns the ComponentLotEntity field if non-nil, zero value otherwise.

### GetComponentLotEntityOk

`func (o *Ingredient) GetComponentLotEntityOk() (*IngredientComponentLotEntity, bool)`

GetComponentLotEntityOk returns a tuple with the ComponentLotEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLotEntity

`func (o *Ingredient) SetComponentLotEntity(v IngredientComponentLotEntity)`

SetComponentLotEntity sets ComponentLotEntity field to given value.

### HasComponentLotEntity

`func (o *Ingredient) HasComponentLotEntity() bool`

HasComponentLotEntity returns a boolean if a field has been set.

### SetComponentLotEntityNil

`func (o *Ingredient) SetComponentLotEntityNil(b bool)`

 SetComponentLotEntityNil sets the value for ComponentLotEntity to be an explicit nil

### UnsetComponentLotEntity
`func (o *Ingredient) UnsetComponentLotEntity()`

UnsetComponentLotEntity ensures that no value is present for ComponentLotEntity, not even an explicit nil
### GetComponentLotText

`func (o *Ingredient) GetComponentLotText() string`

GetComponentLotText returns the ComponentLotText field if non-nil, zero value otherwise.

### GetComponentLotTextOk

`func (o *Ingredient) GetComponentLotTextOk() (*string, bool)`

GetComponentLotTextOk returns a tuple with the ComponentLotText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLotText

`func (o *Ingredient) SetComponentLotText(v string)`

SetComponentLotText sets ComponentLotText field to given value.

### HasComponentLotText

`func (o *Ingredient) HasComponentLotText() bool`

HasComponentLotText returns a boolean if a field has been set.

### SetComponentLotTextNil

`func (o *Ingredient) SetComponentLotTextNil(b bool)`

 SetComponentLotTextNil sets the value for ComponentLotText to be an explicit nil

### UnsetComponentLotText
`func (o *Ingredient) UnsetComponentLotText()`

UnsetComponentLotText ensures that no value is present for ComponentLotText, not even an explicit nil
### GetHasParent

`func (o *Ingredient) GetHasParent() bool`

GetHasParent returns the HasParent field if non-nil, zero value otherwise.

### GetHasParentOk

`func (o *Ingredient) GetHasParentOk() (*bool, bool)`

GetHasParentOk returns a tuple with the HasParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasParent

`func (o *Ingredient) SetHasParent(v bool)`

SetHasParent sets HasParent field to given value.

### HasHasParent

`func (o *Ingredient) HasHasParent() bool`

HasHasParent returns a boolean if a field has been set.

### GetNotes

`func (o *Ingredient) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Ingredient) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Ingredient) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Ingredient) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Ingredient) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Ingredient) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetTargetAmount

`func (o *Ingredient) GetTargetAmount() string`

GetTargetAmount returns the TargetAmount field if non-nil, zero value otherwise.

### GetTargetAmountOk

`func (o *Ingredient) GetTargetAmountOk() (*string, bool)`

GetTargetAmountOk returns a tuple with the TargetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAmount

`func (o *Ingredient) SetTargetAmount(v string)`

SetTargetAmount sets TargetAmount field to given value.

### HasTargetAmount

`func (o *Ingredient) HasTargetAmount() bool`

HasTargetAmount returns a boolean if a field has been set.

### SetTargetAmountNil

`func (o *Ingredient) SetTargetAmountNil(b bool)`

 SetTargetAmountNil sets the value for TargetAmount to be an explicit nil

### UnsetTargetAmount
`func (o *Ingredient) UnsetTargetAmount()`

UnsetTargetAmount ensures that no value is present for TargetAmount, not even an explicit nil
### GetUnits

`func (o *Ingredient) GetUnits() IngredientMeasurementUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Ingredient) GetUnitsOk() (*IngredientMeasurementUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Ingredient) SetUnits(v IngredientMeasurementUnits)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *Ingredient) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnitsNil

`func (o *Ingredient) SetUnitsNil(b bool)`

 SetUnitsNil sets the value for Units to be an explicit nil

### UnsetUnits
`func (o *Ingredient) UnsetUnits()`

UnsetUnits ensures that no value is present for Units, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


