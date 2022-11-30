# Entity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** |  | [optional] 
**Annotations** | Pointer to [**[]AaAnnotation**](AaAnnotation.md) | Array of annotation objects on the AA sequence. | [optional] 
**ApiURL** | Pointer to **string** | The canonical url of the Custom Entity in the API. | [optional] [readonly] 
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**Bases** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Creator** | Pointer to [**CustomEntityCreator**](CustomEntityCreator.md) |  | [optional] 
**CustomFields** | Pointer to [**map[string]CustomField**](CustomField.md) |  | [optional] 
**EntityRegistryId** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**FolderId** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsCircular** | Pointer to **bool** |  | [optional] 
**Length** | Pointer to **int32** | Number of bases in the Oligo. | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Primers** | Pointer to [**[]Primer**](Primer.md) |  | [optional] 
**RegistrationOrigin** | Pointer to [**NullableAaSequenceRegistrationOrigin**](AaSequenceRegistrationOrigin.md) |  | [optional] 
**RegistryId** | Pointer to **NullableString** |  | [optional] 
**Schema** | Pointer to [**NullableCustomEntitySchema**](CustomEntitySchema.md) |  | [optional] 
**Translations** | Pointer to [**[]Translation**](Translation.md) |  | [optional] 
**WebURL** | Pointer to **string** |  | [optional] [readonly] 
**AminoAcids** | Pointer to **string** | Amino acids of the AA sequence. | [optional] 
**AllowMeasuredIngredients** | Pointer to **bool** | Derived from the mixture&#39;s schema. | [optional] [readonly] 
**Amount** | Pointer to **string** | The positive numerical amount value of this mixture in string format (to preserve full precision). Pair with &#x60;units&#x60;. Supports scientific notation (1.23e4). | [optional] 
**Authors** | Pointer to [**[]UserSummary**](UserSummary.md) |  | [optional] 
**Ingredients** | Pointer to [**[]Ingredient**](Ingredient.md) | List of ingredients on this mixture. | [optional] 
**Units** | Pointer to [**NullableMixtureMeasurementUnits**](MixtureMeasurementUnits.md) |  | [optional] 
**NucleotideType** | Pointer to **string** |  | [optional] 

## Methods

### NewEntity

`func NewEntity() *Entity`

NewEntity instantiates a new Entity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityWithDefaults

`func NewEntityWithDefaults() *Entity`

NewEntityWithDefaults instantiates a new Entity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *Entity) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Entity) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *Entity) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *Entity) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAnnotations

`func (o *Entity) GetAnnotations() []AaAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Entity) GetAnnotationsOk() (*[]AaAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Entity) SetAnnotations(v []AaAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Entity) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetApiURL

`func (o *Entity) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *Entity) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *Entity) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *Entity) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetArchiveRecord

`func (o *Entity) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *Entity) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *Entity) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *Entity) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *Entity) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *Entity) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetBases

`func (o *Entity) GetBases() string`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *Entity) GetBasesOk() (*string, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *Entity) SetBases(v string)`

SetBases sets Bases field to given value.

### HasBases

`func (o *Entity) HasBases() bool`

HasBases returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Entity) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Entity) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Entity) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Entity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *Entity) GetCreator() CustomEntityCreator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Entity) GetCreatorOk() (*CustomEntityCreator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Entity) SetCreator(v CustomEntityCreator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Entity) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetCustomFields

`func (o *Entity) GetCustomFields() map[string]CustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Entity) GetCustomFieldsOk() (*map[string]CustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Entity) SetCustomFields(v map[string]CustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Entity) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *Entity) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *Entity) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *Entity) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *Entity) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.

### SetEntityRegistryIdNil

`func (o *Entity) SetEntityRegistryIdNil(b bool)`

 SetEntityRegistryIdNil sets the value for EntityRegistryId to be an explicit nil

### UnsetEntityRegistryId
`func (o *Entity) UnsetEntityRegistryId()`

UnsetEntityRegistryId ensures that no value is present for EntityRegistryId, not even an explicit nil
### GetFields

`func (o *Entity) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Entity) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Entity) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Entity) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *Entity) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *Entity) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *Entity) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *Entity) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *Entity) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *Entity) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetId

`func (o *Entity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Entity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCircular

`func (o *Entity) GetIsCircular() bool`

GetIsCircular returns the IsCircular field if non-nil, zero value otherwise.

### GetIsCircularOk

`func (o *Entity) GetIsCircularOk() (*bool, bool)`

GetIsCircularOk returns a tuple with the IsCircular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCircular

`func (o *Entity) SetIsCircular(v bool)`

SetIsCircular sets IsCircular field to given value.

### HasIsCircular

`func (o *Entity) HasIsCircular() bool`

HasIsCircular returns a boolean if a field has been set.

### GetLength

`func (o *Entity) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *Entity) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *Entity) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *Entity) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Entity) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Entity) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Entity) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Entity) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *Entity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Entity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Entity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Entity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimers

`func (o *Entity) GetPrimers() []Primer`

GetPrimers returns the Primers field if non-nil, zero value otherwise.

### GetPrimersOk

`func (o *Entity) GetPrimersOk() (*[]Primer, bool)`

GetPrimersOk returns a tuple with the Primers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimers

`func (o *Entity) SetPrimers(v []Primer)`

SetPrimers sets Primers field to given value.

### HasPrimers

`func (o *Entity) HasPrimers() bool`

HasPrimers returns a boolean if a field has been set.

### GetRegistrationOrigin

`func (o *Entity) GetRegistrationOrigin() AaSequenceRegistrationOrigin`

GetRegistrationOrigin returns the RegistrationOrigin field if non-nil, zero value otherwise.

### GetRegistrationOriginOk

`func (o *Entity) GetRegistrationOriginOk() (*AaSequenceRegistrationOrigin, bool)`

GetRegistrationOriginOk returns a tuple with the RegistrationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationOrigin

`func (o *Entity) SetRegistrationOrigin(v AaSequenceRegistrationOrigin)`

SetRegistrationOrigin sets RegistrationOrigin field to given value.

### HasRegistrationOrigin

`func (o *Entity) HasRegistrationOrigin() bool`

HasRegistrationOrigin returns a boolean if a field has been set.

### SetRegistrationOriginNil

`func (o *Entity) SetRegistrationOriginNil(b bool)`

 SetRegistrationOriginNil sets the value for RegistrationOrigin to be an explicit nil

### UnsetRegistrationOrigin
`func (o *Entity) UnsetRegistrationOrigin()`

UnsetRegistrationOrigin ensures that no value is present for RegistrationOrigin, not even an explicit nil
### GetRegistryId

`func (o *Entity) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *Entity) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *Entity) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *Entity) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### SetRegistryIdNil

`func (o *Entity) SetRegistryIdNil(b bool)`

 SetRegistryIdNil sets the value for RegistryId to be an explicit nil

### UnsetRegistryId
`func (o *Entity) UnsetRegistryId()`

UnsetRegistryId ensures that no value is present for RegistryId, not even an explicit nil
### GetSchema

`func (o *Entity) GetSchema() CustomEntitySchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Entity) GetSchemaOk() (*CustomEntitySchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Entity) SetSchema(v CustomEntitySchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Entity) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *Entity) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *Entity) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetTranslations

`func (o *Entity) GetTranslations() []Translation`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *Entity) GetTranslationsOk() (*[]Translation, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *Entity) SetTranslations(v []Translation)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *Entity) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.

### GetWebURL

`func (o *Entity) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *Entity) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *Entity) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *Entity) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetAminoAcids

`func (o *Entity) GetAminoAcids() string`

GetAminoAcids returns the AminoAcids field if non-nil, zero value otherwise.

### GetAminoAcidsOk

`func (o *Entity) GetAminoAcidsOk() (*string, bool)`

GetAminoAcidsOk returns a tuple with the AminoAcids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcids

`func (o *Entity) SetAminoAcids(v string)`

SetAminoAcids sets AminoAcids field to given value.

### HasAminoAcids

`func (o *Entity) HasAminoAcids() bool`

HasAminoAcids returns a boolean if a field has been set.

### GetAllowMeasuredIngredients

`func (o *Entity) GetAllowMeasuredIngredients() bool`

GetAllowMeasuredIngredients returns the AllowMeasuredIngredients field if non-nil, zero value otherwise.

### GetAllowMeasuredIngredientsOk

`func (o *Entity) GetAllowMeasuredIngredientsOk() (*bool, bool)`

GetAllowMeasuredIngredientsOk returns a tuple with the AllowMeasuredIngredients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMeasuredIngredients

`func (o *Entity) SetAllowMeasuredIngredients(v bool)`

SetAllowMeasuredIngredients sets AllowMeasuredIngredients field to given value.

### HasAllowMeasuredIngredients

`func (o *Entity) HasAllowMeasuredIngredients() bool`

HasAllowMeasuredIngredients returns a boolean if a field has been set.

### GetAmount

`func (o *Entity) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Entity) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Entity) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Entity) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAuthors

`func (o *Entity) GetAuthors() []UserSummary`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *Entity) GetAuthorsOk() (*[]UserSummary, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *Entity) SetAuthors(v []UserSummary)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *Entity) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetIngredients

`func (o *Entity) GetIngredients() []Ingredient`

GetIngredients returns the Ingredients field if non-nil, zero value otherwise.

### GetIngredientsOk

`func (o *Entity) GetIngredientsOk() (*[]Ingredient, bool)`

GetIngredientsOk returns a tuple with the Ingredients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngredients

`func (o *Entity) SetIngredients(v []Ingredient)`

SetIngredients sets Ingredients field to given value.

### HasIngredients

`func (o *Entity) HasIngredients() bool`

HasIngredients returns a boolean if a field has been set.

### GetUnits

`func (o *Entity) GetUnits() MixtureMeasurementUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Entity) GetUnitsOk() (*MixtureMeasurementUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Entity) SetUnits(v MixtureMeasurementUnits)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *Entity) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnitsNil

`func (o *Entity) SetUnitsNil(b bool)`

 SetUnitsNil sets the value for Units to be an explicit nil

### UnsetUnits
`func (o *Entity) UnsetUnits()`

UnsetUnits ensures that no value is present for Units, not even an explicit nil
### GetNucleotideType

`func (o *Entity) GetNucleotideType() string`

GetNucleotideType returns the NucleotideType field if non-nil, zero value otherwise.

### GetNucleotideTypeOk

`func (o *Entity) GetNucleotideTypeOk() (*string, bool)`

GetNucleotideTypeOk returns a tuple with the NucleotideType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNucleotideType

`func (o *Entity) SetNucleotideType(v string)`

SetNucleotideType sets NucleotideType field to given value.

### HasNucleotideType

`func (o *Entity) HasNucleotideType() bool`

HasNucleotideType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


