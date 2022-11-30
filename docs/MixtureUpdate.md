# MixtureUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Aliases to add to the mixture | [optional] 
**Amount** | Pointer to **string** | The positive numerical amount value of this mixture in string format (to preserve full precision). Pair with &#x60;units&#x60;. Supports scientific notation (1.23e4). | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the mixture&#39;s authors. | [optional] 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the mixture. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**EntityRegistryId** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Schema fields to set on the mixture. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field. If you are setting the parent mixture field here, you must also specify &#x60;ingredients&#x60;  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder that the entity is moved into | [optional] 
**Ingredients** | Pointer to [**[]IngredientWriteParams**](IngredientWriteParams.md) | Desired final state for the ingredients on this mixture. Each ingredient you specify will be matched with the existing ingredients on the mixture based on the component entity, and Benchling will create, update, or delete this mixture&#39;s ingredients so that the final state of this mixture&#39;s ingredients matches your request. Benchling will recognize that any ingredients you specify that match ingredients on the parent mixture (based on component entity) are inherited. This can be seen on the returned &#x60;ingredients[i].hasParent&#x60; attribute.  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SchemaId** | Pointer to **string** |  | [optional] 
**Units** | Pointer to [**NullableMixtureMeasurementUnits**](MixtureMeasurementUnits.md) |  | [optional] 

## Methods

### NewMixtureUpdate

`func NewMixtureUpdate() *MixtureUpdate`

NewMixtureUpdate instantiates a new MixtureUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMixtureUpdateWithDefaults

`func NewMixtureUpdateWithDefaults() *MixtureUpdate`

NewMixtureUpdateWithDefaults instantiates a new MixtureUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *MixtureUpdate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *MixtureUpdate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *MixtureUpdate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *MixtureUpdate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAmount

`func (o *MixtureUpdate) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MixtureUpdate) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MixtureUpdate) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *MixtureUpdate) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAuthorIds

`func (o *MixtureUpdate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *MixtureUpdate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *MixtureUpdate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *MixtureUpdate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetCustomFields

`func (o *MixtureUpdate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *MixtureUpdate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *MixtureUpdate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *MixtureUpdate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *MixtureUpdate) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *MixtureUpdate) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *MixtureUpdate) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *MixtureUpdate) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.

### GetFields

`func (o *MixtureUpdate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MixtureUpdate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MixtureUpdate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MixtureUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *MixtureUpdate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *MixtureUpdate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *MixtureUpdate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *MixtureUpdate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetIngredients

`func (o *MixtureUpdate) GetIngredients() []IngredientWriteParams`

GetIngredients returns the Ingredients field if non-nil, zero value otherwise.

### GetIngredientsOk

`func (o *MixtureUpdate) GetIngredientsOk() (*[]IngredientWriteParams, bool)`

GetIngredientsOk returns a tuple with the Ingredients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngredients

`func (o *MixtureUpdate) SetIngredients(v []IngredientWriteParams)`

SetIngredients sets Ingredients field to given value.

### HasIngredients

`func (o *MixtureUpdate) HasIngredients() bool`

HasIngredients returns a boolean if a field has been set.

### GetName

`func (o *MixtureUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MixtureUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MixtureUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MixtureUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchemaId

`func (o *MixtureUpdate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *MixtureUpdate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *MixtureUpdate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *MixtureUpdate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetUnits

`func (o *MixtureUpdate) GetUnits() MixtureMeasurementUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *MixtureUpdate) GetUnitsOk() (*MixtureMeasurementUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *MixtureUpdate) SetUnits(v MixtureMeasurementUnits)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *MixtureUpdate) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnitsNil

`func (o *MixtureUpdate) SetUnitsNil(b bool)`

 SetUnitsNil sets the value for Units to be an explicit nil

### UnsetUnits
`func (o *MixtureUpdate) UnsetUnits()`

UnsetUnits ensures that no value is present for Units, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


