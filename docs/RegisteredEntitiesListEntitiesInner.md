# RegisteredEntitiesListEntitiesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Array of aliases | [optional] 
**Annotations** | Pointer to [**[]AaAnnotation**](AaAnnotation.md) | Array of annotation objects on the AA sequence. | [optional] 
**ApiURL** | Pointer to **string** |  | [optional] 
**ArchiveRecord** | Pointer to [**NullableAaSequenceArchiveRecord**](AaSequenceArchiveRecord.md) |  | [optional] 
**Bases** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | DateTime the Oligo was created. | [optional] [readonly] 
**Creator** | Pointer to [**UserSummary**](UserSummary.md) |  | [optional] 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields set on the Oligo. | [optional] 
**EntityRegistryId** | Pointer to **NullableString** | Registry ID of the Oligo if registered. | [optional] 
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**FolderId** | Pointer to **NullableString** | ID of the folder that contains the Oligo. | [optional] 
**Id** | Pointer to **string** | ID of the Oligo. | [optional] 
**IsCircular** | Pointer to **bool** |  | [optional] 
**Length** | Pointer to **int32** | Number of bases in the Oligo. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the Oligo was last modified. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the Oligo. | [optional] 
**Primers** | Pointer to [**[]Primer**](Primer.md) |  | [optional] 
**RegistrationOrigin** | Pointer to [**NullableAaSequenceRegistrationOrigin**](AaSequenceRegistrationOrigin.md) |  | [optional] 
**RegistryId** | Pointer to **NullableString** | Registry the Oligo is registered in. | [optional] 
**Schema** | Pointer to [**NullableAaSequenceSchema**](AaSequenceSchema.md) |  | [optional] 
**Translations** | Pointer to [**[]Translation**](Translation.md) |  | [optional] 
**WebURL** | Pointer to **string** | URL of the Oligo. | [optional] [readonly] 
**EntityType** | Pointer to **string** |  | [optional] 
**Authors** | Pointer to [**[]UserSummary**](UserSummary.md) |  | [optional] 
**AminoAcids** | Pointer to **string** | Amino acids of the AA sequence. | [optional] 
**AllowMeasuredIngredients** | Pointer to **bool** | Derived from the mixture&#39;s schema. | [optional] [readonly] 
**Amount** | Pointer to **string** | The positive numerical amount value of this mixture in string format (to preserve full precision). Pair with &#x60;units&#x60;. Supports scientific notation (1.23e4). | [optional] 
**Ingredients** | Pointer to [**[]Ingredient**](Ingredient.md) | List of ingredients on this mixture. | [optional] 
**Units** | Pointer to [**NullableMixtureMeasurementUnits**](MixtureMeasurementUnits.md) |  | [optional] 
**NucleotideType** | Pointer to **string** |  | [optional] 

## Methods

### NewRegisteredEntitiesListEntitiesInner

`func NewRegisteredEntitiesListEntitiesInner() *RegisteredEntitiesListEntitiesInner`

NewRegisteredEntitiesListEntitiesInner instantiates a new RegisteredEntitiesListEntitiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisteredEntitiesListEntitiesInnerWithDefaults

`func NewRegisteredEntitiesListEntitiesInnerWithDefaults() *RegisteredEntitiesListEntitiesInner`

NewRegisteredEntitiesListEntitiesInnerWithDefaults instantiates a new RegisteredEntitiesListEntitiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *RegisteredEntitiesListEntitiesInner) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *RegisteredEntitiesListEntitiesInner) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *RegisteredEntitiesListEntitiesInner) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *RegisteredEntitiesListEntitiesInner) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAnnotations

`func (o *RegisteredEntitiesListEntitiesInner) GetAnnotations() []AaAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *RegisteredEntitiesListEntitiesInner) GetAnnotationsOk() (*[]AaAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *RegisteredEntitiesListEntitiesInner) SetAnnotations(v []AaAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *RegisteredEntitiesListEntitiesInner) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetApiURL

`func (o *RegisteredEntitiesListEntitiesInner) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *RegisteredEntitiesListEntitiesInner) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *RegisteredEntitiesListEntitiesInner) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *RegisteredEntitiesListEntitiesInner) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetArchiveRecord

`func (o *RegisteredEntitiesListEntitiesInner) GetArchiveRecord() AaSequenceArchiveRecord`

GetArchiveRecord returns the ArchiveRecord field if non-nil, zero value otherwise.

### GetArchiveRecordOk

`func (o *RegisteredEntitiesListEntitiesInner) GetArchiveRecordOk() (*AaSequenceArchiveRecord, bool)`

GetArchiveRecordOk returns a tuple with the ArchiveRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveRecord

`func (o *RegisteredEntitiesListEntitiesInner) SetArchiveRecord(v AaSequenceArchiveRecord)`

SetArchiveRecord sets ArchiveRecord field to given value.

### HasArchiveRecord

`func (o *RegisteredEntitiesListEntitiesInner) HasArchiveRecord() bool`

HasArchiveRecord returns a boolean if a field has been set.

### SetArchiveRecordNil

`func (o *RegisteredEntitiesListEntitiesInner) SetArchiveRecordNil(b bool)`

 SetArchiveRecordNil sets the value for ArchiveRecord to be an explicit nil

### UnsetArchiveRecord
`func (o *RegisteredEntitiesListEntitiesInner) UnsetArchiveRecord()`

UnsetArchiveRecord ensures that no value is present for ArchiveRecord, not even an explicit nil
### GetBases

`func (o *RegisteredEntitiesListEntitiesInner) GetBases() string`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *RegisteredEntitiesListEntitiesInner) GetBasesOk() (*string, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *RegisteredEntitiesListEntitiesInner) SetBases(v string)`

SetBases sets Bases field to given value.

### HasBases

`func (o *RegisteredEntitiesListEntitiesInner) HasBases() bool`

HasBases returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RegisteredEntitiesListEntitiesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RegisteredEntitiesListEntitiesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RegisteredEntitiesListEntitiesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RegisteredEntitiesListEntitiesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *RegisteredEntitiesListEntitiesInner) GetCreator() UserSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RegisteredEntitiesListEntitiesInner) GetCreatorOk() (*UserSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RegisteredEntitiesListEntitiesInner) SetCreator(v UserSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RegisteredEntitiesListEntitiesInner) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetCustomFields

`func (o *RegisteredEntitiesListEntitiesInner) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RegisteredEntitiesListEntitiesInner) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RegisteredEntitiesListEntitiesInner) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RegisteredEntitiesListEntitiesInner) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *RegisteredEntitiesListEntitiesInner) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *RegisteredEntitiesListEntitiesInner) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *RegisteredEntitiesListEntitiesInner) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *RegisteredEntitiesListEntitiesInner) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.

### SetEntityRegistryIdNil

`func (o *RegisteredEntitiesListEntitiesInner) SetEntityRegistryIdNil(b bool)`

 SetEntityRegistryIdNil sets the value for EntityRegistryId to be an explicit nil

### UnsetEntityRegistryId
`func (o *RegisteredEntitiesListEntitiesInner) UnsetEntityRegistryId()`

UnsetEntityRegistryId ensures that no value is present for EntityRegistryId, not even an explicit nil
### GetFields

`func (o *RegisteredEntitiesListEntitiesInner) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RegisteredEntitiesListEntitiesInner) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RegisteredEntitiesListEntitiesInner) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RegisteredEntitiesListEntitiesInner) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *RegisteredEntitiesListEntitiesInner) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *RegisteredEntitiesListEntitiesInner) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *RegisteredEntitiesListEntitiesInner) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *RegisteredEntitiesListEntitiesInner) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *RegisteredEntitiesListEntitiesInner) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *RegisteredEntitiesListEntitiesInner) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetId

`func (o *RegisteredEntitiesListEntitiesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisteredEntitiesListEntitiesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisteredEntitiesListEntitiesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegisteredEntitiesListEntitiesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCircular

`func (o *RegisteredEntitiesListEntitiesInner) GetIsCircular() bool`

GetIsCircular returns the IsCircular field if non-nil, zero value otherwise.

### GetIsCircularOk

`func (o *RegisteredEntitiesListEntitiesInner) GetIsCircularOk() (*bool, bool)`

GetIsCircularOk returns a tuple with the IsCircular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCircular

`func (o *RegisteredEntitiesListEntitiesInner) SetIsCircular(v bool)`

SetIsCircular sets IsCircular field to given value.

### HasIsCircular

`func (o *RegisteredEntitiesListEntitiesInner) HasIsCircular() bool`

HasIsCircular returns a boolean if a field has been set.

### GetLength

`func (o *RegisteredEntitiesListEntitiesInner) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *RegisteredEntitiesListEntitiesInner) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *RegisteredEntitiesListEntitiesInner) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *RegisteredEntitiesListEntitiesInner) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetModifiedAt

`func (o *RegisteredEntitiesListEntitiesInner) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *RegisteredEntitiesListEntitiesInner) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *RegisteredEntitiesListEntitiesInner) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *RegisteredEntitiesListEntitiesInner) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *RegisteredEntitiesListEntitiesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisteredEntitiesListEntitiesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisteredEntitiesListEntitiesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegisteredEntitiesListEntitiesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimers

`func (o *RegisteredEntitiesListEntitiesInner) GetPrimers() []Primer`

GetPrimers returns the Primers field if non-nil, zero value otherwise.

### GetPrimersOk

`func (o *RegisteredEntitiesListEntitiesInner) GetPrimersOk() (*[]Primer, bool)`

GetPrimersOk returns a tuple with the Primers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimers

`func (o *RegisteredEntitiesListEntitiesInner) SetPrimers(v []Primer)`

SetPrimers sets Primers field to given value.

### HasPrimers

`func (o *RegisteredEntitiesListEntitiesInner) HasPrimers() bool`

HasPrimers returns a boolean if a field has been set.

### GetRegistrationOrigin

`func (o *RegisteredEntitiesListEntitiesInner) GetRegistrationOrigin() AaSequenceRegistrationOrigin`

GetRegistrationOrigin returns the RegistrationOrigin field if non-nil, zero value otherwise.

### GetRegistrationOriginOk

`func (o *RegisteredEntitiesListEntitiesInner) GetRegistrationOriginOk() (*AaSequenceRegistrationOrigin, bool)`

GetRegistrationOriginOk returns a tuple with the RegistrationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationOrigin

`func (o *RegisteredEntitiesListEntitiesInner) SetRegistrationOrigin(v AaSequenceRegistrationOrigin)`

SetRegistrationOrigin sets RegistrationOrigin field to given value.

### HasRegistrationOrigin

`func (o *RegisteredEntitiesListEntitiesInner) HasRegistrationOrigin() bool`

HasRegistrationOrigin returns a boolean if a field has been set.

### SetRegistrationOriginNil

`func (o *RegisteredEntitiesListEntitiesInner) SetRegistrationOriginNil(b bool)`

 SetRegistrationOriginNil sets the value for RegistrationOrigin to be an explicit nil

### UnsetRegistrationOrigin
`func (o *RegisteredEntitiesListEntitiesInner) UnsetRegistrationOrigin()`

UnsetRegistrationOrigin ensures that no value is present for RegistrationOrigin, not even an explicit nil
### GetRegistryId

`func (o *RegisteredEntitiesListEntitiesInner) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *RegisteredEntitiesListEntitiesInner) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *RegisteredEntitiesListEntitiesInner) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *RegisteredEntitiesListEntitiesInner) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### SetRegistryIdNil

`func (o *RegisteredEntitiesListEntitiesInner) SetRegistryIdNil(b bool)`

 SetRegistryIdNil sets the value for RegistryId to be an explicit nil

### UnsetRegistryId
`func (o *RegisteredEntitiesListEntitiesInner) UnsetRegistryId()`

UnsetRegistryId ensures that no value is present for RegistryId, not even an explicit nil
### GetSchema

`func (o *RegisteredEntitiesListEntitiesInner) GetSchema() AaSequenceSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *RegisteredEntitiesListEntitiesInner) GetSchemaOk() (*AaSequenceSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *RegisteredEntitiesListEntitiesInner) SetSchema(v AaSequenceSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *RegisteredEntitiesListEntitiesInner) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *RegisteredEntitiesListEntitiesInner) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *RegisteredEntitiesListEntitiesInner) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetTranslations

`func (o *RegisteredEntitiesListEntitiesInner) GetTranslations() []Translation`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *RegisteredEntitiesListEntitiesInner) GetTranslationsOk() (*[]Translation, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *RegisteredEntitiesListEntitiesInner) SetTranslations(v []Translation)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *RegisteredEntitiesListEntitiesInner) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.

### GetWebURL

`func (o *RegisteredEntitiesListEntitiesInner) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *RegisteredEntitiesListEntitiesInner) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *RegisteredEntitiesListEntitiesInner) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *RegisteredEntitiesListEntitiesInner) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### GetEntityType

`func (o *RegisteredEntitiesListEntitiesInner) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *RegisteredEntitiesListEntitiesInner) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *RegisteredEntitiesListEntitiesInner) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *RegisteredEntitiesListEntitiesInner) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetAuthors

`func (o *RegisteredEntitiesListEntitiesInner) GetAuthors() []UserSummary`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *RegisteredEntitiesListEntitiesInner) GetAuthorsOk() (*[]UserSummary, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *RegisteredEntitiesListEntitiesInner) SetAuthors(v []UserSummary)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *RegisteredEntitiesListEntitiesInner) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetAminoAcids

`func (o *RegisteredEntitiesListEntitiesInner) GetAminoAcids() string`

GetAminoAcids returns the AminoAcids field if non-nil, zero value otherwise.

### GetAminoAcidsOk

`func (o *RegisteredEntitiesListEntitiesInner) GetAminoAcidsOk() (*string, bool)`

GetAminoAcidsOk returns a tuple with the AminoAcids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAminoAcids

`func (o *RegisteredEntitiesListEntitiesInner) SetAminoAcids(v string)`

SetAminoAcids sets AminoAcids field to given value.

### HasAminoAcids

`func (o *RegisteredEntitiesListEntitiesInner) HasAminoAcids() bool`

HasAminoAcids returns a boolean if a field has been set.

### GetAllowMeasuredIngredients

`func (o *RegisteredEntitiesListEntitiesInner) GetAllowMeasuredIngredients() bool`

GetAllowMeasuredIngredients returns the AllowMeasuredIngredients field if non-nil, zero value otherwise.

### GetAllowMeasuredIngredientsOk

`func (o *RegisteredEntitiesListEntitiesInner) GetAllowMeasuredIngredientsOk() (*bool, bool)`

GetAllowMeasuredIngredientsOk returns a tuple with the AllowMeasuredIngredients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMeasuredIngredients

`func (o *RegisteredEntitiesListEntitiesInner) SetAllowMeasuredIngredients(v bool)`

SetAllowMeasuredIngredients sets AllowMeasuredIngredients field to given value.

### HasAllowMeasuredIngredients

`func (o *RegisteredEntitiesListEntitiesInner) HasAllowMeasuredIngredients() bool`

HasAllowMeasuredIngredients returns a boolean if a field has been set.

### GetAmount

`func (o *RegisteredEntitiesListEntitiesInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RegisteredEntitiesListEntitiesInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RegisteredEntitiesListEntitiesInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RegisteredEntitiesListEntitiesInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetIngredients

`func (o *RegisteredEntitiesListEntitiesInner) GetIngredients() []Ingredient`

GetIngredients returns the Ingredients field if non-nil, zero value otherwise.

### GetIngredientsOk

`func (o *RegisteredEntitiesListEntitiesInner) GetIngredientsOk() (*[]Ingredient, bool)`

GetIngredientsOk returns a tuple with the Ingredients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngredients

`func (o *RegisteredEntitiesListEntitiesInner) SetIngredients(v []Ingredient)`

SetIngredients sets Ingredients field to given value.

### HasIngredients

`func (o *RegisteredEntitiesListEntitiesInner) HasIngredients() bool`

HasIngredients returns a boolean if a field has been set.

### GetUnits

`func (o *RegisteredEntitiesListEntitiesInner) GetUnits() MixtureMeasurementUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *RegisteredEntitiesListEntitiesInner) GetUnitsOk() (*MixtureMeasurementUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *RegisteredEntitiesListEntitiesInner) SetUnits(v MixtureMeasurementUnits)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *RegisteredEntitiesListEntitiesInner) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### SetUnitsNil

`func (o *RegisteredEntitiesListEntitiesInner) SetUnitsNil(b bool)`

 SetUnitsNil sets the value for Units to be an explicit nil

### UnsetUnits
`func (o *RegisteredEntitiesListEntitiesInner) UnsetUnits()`

UnsetUnits ensures that no value is present for Units, not even an explicit nil
### GetNucleotideType

`func (o *RegisteredEntitiesListEntitiesInner) GetNucleotideType() string`

GetNucleotideType returns the NucleotideType field if non-nil, zero value otherwise.

### GetNucleotideTypeOk

`func (o *RegisteredEntitiesListEntitiesInner) GetNucleotideTypeOk() (*string, bool)`

GetNucleotideTypeOk returns a tuple with the NucleotideType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNucleotideType

`func (o *RegisteredEntitiesListEntitiesInner) SetNucleotideType(v string)`

SetNucleotideType sets NucleotideType field to given value.

### HasNucleotideType

`func (o *RegisteredEntitiesListEntitiesInner) HasNucleotideType() bool`

HasNucleotideType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


