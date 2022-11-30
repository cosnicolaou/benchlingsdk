# MixtureWithEntityType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** |  | [optional] 
**AllowMeasuredIngredients** | Pointer to **bool** | Derived from the mixture&#39;s schema. | [optional] [readonly] 
**Amount** | Pointer to **string** | The positive numerical amount value of this mixture in string format (to preserve full precision). Pair with &#x60;units&#x60;. Supports scientific notation (1.23e4). | [optional] 
**ApiURL** | Pointer to **string** | The canonical url of the Mixture in the API. | [optional] [readonly] 
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**Authors** | Pointer to [**[]UserSummary**](UserSummary.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Creator** | Pointer to [**CustomEntityCreator**](CustomEntityCreator.md) |  | [optional] 
**CustomFields** | Pointer to [**map[string]CustomField**](CustomField.md) |  | [optional] 
**EntityRegistryId** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Mixtures can have up to one parent mixture field. | [optional] 
**FolderId** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Ingredients** | Pointer to [**[]Ingredient**](Ingredient.md) | List of ingredients on this mixture. | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**RegistrationOrigin** | Pointer to [**NullableAaSequenceRegistrationOrigin**](AaSequenceRegistrationOrigin.md) |  | [optional] 
**RegistryId** | Pointer to **NullableString** |  | [optional] 
**Schema** | Pointer to [**MixtureSchema**](MixtureSchema.md) |  | [optional] 
**Units** | Pointer to [**NullableMixtureMeasurementUnits**](MixtureMeasurementUnits.md) |  | [optional] 
**WebURL** | Pointer to **string** |  | [optional] [readonly] 
**EntityType** | Pointer to **string** |  | [optional] 

## Methods

### NewMixtureWithEntityType

`func NewMixtureWithEntityType() *MixtureWithEntityType`

NewMixtureWithEntityType instantiates a new MixtureWithEntityType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMixtureWithEntityTypeWithDefaults

`func NewMixtureWithEntityTypeWithDefaults() *MixtureWithEntityType`

NewMixtureWithEntityTypeWithDefaults instantiates a new MixtureWithEntityType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *MixtureWithEntityType) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *MixtureWithEntityType) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *MixtureWithEntityType) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *MixtureWithEntityType) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAllowMeasuredIngredients

`func (o *MixtureWithEntityType) GetAllowMeasuredIngredients() bool`

GetAllowMeasuredIngredients returns the AllowMeasuredIngredients field if non-nil, zero value otherwise.

### GetAllowMeasuredIngredientsOk

`func (o *MixtureWithEntityType) GetAllowMeasuredIngredientsOk() (*bool, bool)`

GetAllowMeasuredIngredientsOk returns a tuple with the AllowMeasuredIngredients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMeasuredIngredients

`func (o *MixtureWithEntityType) SetAllowMeasuredIngredients(v bool)`

SetAllowMeasuredIngredients sets AllowMeasuredIngredients field to given value.

### HasAllowMeasuredIngredients

`func (o *MixtureWithEntityType) HasAllowMeasuredIngredients() bool`

HasAllowMeasuredIngredients returns a boolean if a field has been set.

### GetAmount

`func (o *MixtureWithEntityType) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MixtureWithEntityType) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MixtureWithEntityType) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *MixtureWithEntityType) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetApiURL

`func (o *MixtureWithEntityType) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *MixtureWithEntityType) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *MixtureWithEntityType) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *MixtureWithEntityType) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetArchiveRecord

`func (o *MixtureWithEntityType) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *MixtureWithEntityType) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *MixtureWithEntityType) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *MixtureWithEntityType) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *MixtureWithEntityType) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *MixtureWithEntityType) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetAuthors

`func (o *MixtureWithEntityType) GetAuthors() []UserSummary`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *MixtureWithEntityType) GetAuthorsOk() (*[]UserSummary, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *MixtureWithEntityType) SetAuthors(v []UserSummary)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *MixtureWithEntityType) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MixtureWithEntityType) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MixtureWithEntityType) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MixtureWithEntityType) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MixtureWithEntityType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *MixtureWithEntityType) GetCreator() CustomEntityCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *MixtureWithEntityType) GetCreatorOk() (*CustomEntityCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *MixtureWithEntityType) SetCreator(v CustomEntityCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *MixtureWithEntityType) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetCustomFields

`func (o *MixtureWithEntityType) GetCustomFields() map[string]CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *MixtureWithEntityType) GetCustomFieldsOk() (*map[string]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *MixtureWithEntityType) SetCustomFields(v map[string]CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *MixtureWithEntityType) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *MixtureWithEntityType) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *MixtureWithEntityType) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *MixtureWithEntityType) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *MixtureWithEntityType) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.

### SetEntityRegistryIdNil

`func (o *MixtureWithEntityType) SetEntityRegistryIdNil(b bool)`

 SetEntityRegistryIdNil sets the value for EntityRegistryId to be an explicit nil

### UnsetEntityRegistryId
`func (o *MixtureWithEntityType) UnsetEntityRegistryId()`

UnsetEntityRegistryId ensures that no value is present for EntityRegistryId, not even an explicit nil
### GetFields

`func (o *MixtureWithEntityType) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MixtureWithEntityType) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MixtureWithEntityType) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MixtureWithEntityType) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *MixtureWithEntityType) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *MixtureWithEntityType) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *MixtureWithEntityType) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *MixtureWithEntityType) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *MixtureWithEntityType) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *MixtureWithEntityType) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetId

`func (o *MixtureWithEntityType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MixtureWithEntityType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MixtureWithEntityType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MixtureWithEntityType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIngredients

`func (o *MixtureWithEntityType) GetIngredients() []Ingredient`

GetIngredients returns the Ingredients field if non-nil, zero value otherwise.

### GetIngredientsOk

`func (o *MixtureWithEntityType) GetIngredientsOk() (*[]Ingredient, bool)`

GetIngredientsOk returns a tuple with the Ingredients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngredients

`func (o *MixtureWithEntityType) SetIngredients(v []Ingredient)`

SetIngredients sets Ingredients field to given value.

### HasIngredients

`func (o *MixtureWithEntityType) HasIngredients() bool`

HasIngredients returns a boolean if a field has been set.

### GetModifiedAt

`func (o *MixtureWithEntityType) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *MixtureWithEntityType) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *MixtureWithEntityType) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *MixtureWithEntityType) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *MixtureWithEntityType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MixtureWithEntityType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MixtureWithEntityType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MixtureWithEntityType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegistrationOrigin

`func (o *MixtureWithEntityType) GetRegistrationOrigin() AaSequenceRegistrationOrigin`

GetRegistrationOrigin returns the RegistrationOrigin field if non-nil, zero value otherwise.

### GetRegistrationOriginOk

`func (o *MixtureWithEntityType) GetRegistrationOriginOk() (*AaSequenceRegistrationOrigin, bool)`

GetRegistrationOriginOk returns a tuple with the RegistrationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationOrigin

`func (o *MixtureWithEntityType) SetRegistrationOrigin(v AaSequenceRegistrationOrigin)`

SetRegistrationOrigin sets RegistrationOrigin field to given value.

### HasRegistrationOrigin

`func (o *MixtureWithEntityType) HasRegistrationOrigin() bool`

HasRegistrationOrigin returns a boolean if a field has been set.

### SetRegistrationOriginNil

`func (o *MixtureWithEntityType) SetRegistrationOriginNil(b bool)`

 SetRegistrationOriginNil sets the value for RegistrationOrigin to be an explicit nil

### UnsetRegistrationOrigin
`func (o *MixtureWithEntityType) UnsetRegistrationOrigin()`

UnsetRegistrationOrigin ensures that no value is present for RegistrationOrigin, not even an explicit nil
### GetRegistryId

`func (o *MixtureWithEntityType) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *MixtureWithEntityType) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *MixtureWithEntityType) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *MixtureWithEntityType) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### SetRegistryIdNil

`func (o *MixtureWithEntityType) SetRegistryIdNil(b bool)`

 SetRegistryIdNil sets the value for RegistryId to be an explicit nil

### UnsetRegistryId
`func (o *MixtureWithEntityType) UnsetRegistryId()`

UnsetRegistryId ensures that no value is present for RegistryId, not even an explicit nil
### GetSchema

`func (o *MixtureWithEntityType) GetSchema() MixtureSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *MixtureWithEntityType) GetSchemaOk() (*MixtureSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *MixtureWithEntityType) SetSchema(v MixtureSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *MixtureWithEntityType) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetUnits

`func (o *MixtureWithEntityType) GetUnits() MixtureMeasurementUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *MixtureWithEntityType) GetUnitsOk() (*MixtureMeasurementUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *MixtureWithEntityType) SetUnits(v MixtureMeasurementUnits)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *MixtureWithEntityType) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnitsNil

`func (o *MixtureWithEntityType) SetUnitsNil(b bool)`

 SetUnitsNil sets the value for Units to be an explicit nil

### UnsetUnits
`func (o *MixtureWithEntityType) UnsetUnits()`

UnsetUnits ensures that no value is present for Units, not even an explicit nil
### GetWebURL

`func (o *MixtureWithEntityType) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *MixtureWithEntityType) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *MixtureWithEntityType) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *MixtureWithEntityType) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetEntityType

`func (o *MixtureWithEntityType) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *MixtureWithEntityType) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *MixtureWithEntityType) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *MixtureWithEntityType) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


