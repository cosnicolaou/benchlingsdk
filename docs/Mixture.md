# Mixture

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

## Methods

### NewMixture

`func NewMixture() *Mixture`

NewMixture instantiates a new Mixture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMixtureWithDefaults

`func NewMixtureWithDefaults() *Mixture`

NewMixtureWithDefaults instantiates a new Mixture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *Mixture) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Mixture) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *Mixture) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *Mixture) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAllowMeasuredIngredients

`func (o *Mixture) GetAllowMeasuredIngredients() bool`

GetAllowMeasuredIngredients returns the AllowMeasuredIngredients field if non-nil, zero value otherwise.

### GetAllowMeasuredIngredientsOk

`func (o *Mixture) GetAllowMeasuredIngredientsOk() (*bool, bool)`

GetAllowMeasuredIngredientsOk returns a tuple with the AllowMeasuredIngredients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMeasuredIngredients

`func (o *Mixture) SetAllowMeasuredIngredients(v bool)`

SetAllowMeasuredIngredients sets AllowMeasuredIngredients field to given value.

### HasAllowMeasuredIngredients

`func (o *Mixture) HasAllowMeasuredIngredients() bool`

HasAllowMeasuredIngredients returns a boolean if a field has been set.

### GetAmount

`func (o *Mixture) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Mixture) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Mixture) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Mixture) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetApiURL

`func (o *Mixture) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *Mixture) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *Mixture) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *Mixture) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetArchiveRecord

`func (o *Mixture) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *Mixture) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *Mixture) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *Mixture) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *Mixture) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *Mixture) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetAuthors

`func (o *Mixture) GetAuthors() []UserSummary`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *Mixture) GetAuthorsOk() (*[]UserSummary, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *Mixture) SetAuthors(v []UserSummary)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *Mixture) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Mixture) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Mixture) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Mixture) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Mixture) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *Mixture) GetCreator() CustomEntityCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Mixture) GetCreatorOk() (*CustomEntityCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Mixture) SetCreator(v CustomEntityCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Mixture) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetCustomFields

`func (o *Mixture) GetCustomFields() map[string]CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Mixture) GetCustomFieldsOk() (*map[string]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Mixture) SetCustomFields(v map[string]CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Mixture) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *Mixture) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *Mixture) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *Mixture) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *Mixture) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.

### SetEntityRegistryIdNil

`func (o *Mixture) SetEntityRegistryIdNil(b bool)`

 SetEntityRegistryIdNil sets the value for EntityRegistryId to be an explicit nil

### UnsetEntityRegistryId
`func (o *Mixture) UnsetEntityRegistryId()`

UnsetEntityRegistryId ensures that no value is present for EntityRegistryId, not even an explicit nil
### GetFields

`func (o *Mixture) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Mixture) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Mixture) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Mixture) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *Mixture) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *Mixture) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *Mixture) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *Mixture) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *Mixture) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *Mixture) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetId

`func (o *Mixture) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Mixture) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Mixture) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Mixture) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIngredients

`func (o *Mixture) GetIngredients() []Ingredient`

GetIngredients returns the Ingredients field if non-nil, zero value otherwise.

### GetIngredientsOk

`func (o *Mixture) GetIngredientsOk() (*[]Ingredient, bool)`

GetIngredientsOk returns a tuple with the Ingredients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngredients

`func (o *Mixture) SetIngredients(v []Ingredient)`

SetIngredients sets Ingredients field to given value.

### HasIngredients

`func (o *Mixture) HasIngredients() bool`

HasIngredients returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Mixture) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Mixture) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Mixture) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Mixture) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *Mixture) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Mixture) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Mixture) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Mixture) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegistrationOrigin

`func (o *Mixture) GetRegistrationOrigin() AaSequenceRegistrationOrigin`

GetRegistrationOrigin returns the RegistrationOrigin field if non-nil, zero value otherwise.

### GetRegistrationOriginOk

`func (o *Mixture) GetRegistrationOriginOk() (*AaSequenceRegistrationOrigin, bool)`

GetRegistrationOriginOk returns a tuple with the RegistrationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationOrigin

`func (o *Mixture) SetRegistrationOrigin(v AaSequenceRegistrationOrigin)`

SetRegistrationOrigin sets RegistrationOrigin field to given value.

### HasRegistrationOrigin

`func (o *Mixture) HasRegistrationOrigin() bool`

HasRegistrationOrigin returns a boolean if a field has been set.

### SetRegistrationOriginNil

`func (o *Mixture) SetRegistrationOriginNil(b bool)`

 SetRegistrationOriginNil sets the value for RegistrationOrigin to be an explicit nil

### UnsetRegistrationOrigin
`func (o *Mixture) UnsetRegistrationOrigin()`

UnsetRegistrationOrigin ensures that no value is present for RegistrationOrigin, not even an explicit nil
### GetRegistryId

`func (o *Mixture) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *Mixture) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *Mixture) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *Mixture) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### SetRegistryIdNil

`func (o *Mixture) SetRegistryIdNil(b bool)`

 SetRegistryIdNil sets the value for RegistryId to be an explicit nil

### UnsetRegistryId
`func (o *Mixture) UnsetRegistryId()`

UnsetRegistryId ensures that no value is present for RegistryId, not even an explicit nil
### GetSchema

`func (o *Mixture) GetSchema() MixtureSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Mixture) GetSchemaOk() (*MixtureSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Mixture) SetSchema(v MixtureSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Mixture) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetUnits

`func (o *Mixture) GetUnits() MixtureMeasurementUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Mixture) GetUnitsOk() (*MixtureMeasurementUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Mixture) SetUnits(v MixtureMeasurementUnits)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *Mixture) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnitsNil

`func (o *Mixture) SetUnitsNil(b bool)`

 SetUnitsNil sets the value for Units to be an explicit nil

### UnsetUnits
`func (o *Mixture) UnsetUnits()`

UnsetUnits ensures that no value is present for Units, not even an explicit nil
### GetWebURL

`func (o *Mixture) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *Mixture) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *Mixture) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *Mixture) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


