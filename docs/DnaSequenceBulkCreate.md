# DnaSequenceBulkCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Aliases to add to the DNA sequence | [optional] 
**Annotations** | Pointer to [**[]DnaAnnotation**](DnaAnnotation.md) | Annotations to create on the DNA sequence.  | [optional] 
**AuthorIds** | Pointer to **[]string** | IDs of users to set as the DNA sequence&#39;s authors. | [optional] 
**Bases** | **string** | Base pairs for the DNA sequence.  | 
**CustomFields** | Pointer to [**ModelMap**](map.md) | Custom fields to add to the DNA sequence. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**Fields** | Pointer to [**ModelMap**](map.md) | Fields to set on the DNA sequence. Must correspond with the schema&#39;s field definitions. Every field should have its name as a key, mapping to an object with information about the value of the field.  | [optional] 
**FolderId** | Pointer to **string** | ID of the folder containing the entity. Can be left empty when registryId is present. | [optional] 
**IsCircular** | **bool** | Whether the DNA sequence is circular or linear.  | 
**Name** | **string** | Name of the DNA sequence.  | 
**Primers** | Pointer to [**[]Primer**](Primer.md) |  | [optional] 
**SchemaId** | Pointer to **string** | ID of the DNA sequence&#39;s schema.  | [optional] 
**Translations** | Pointer to [**[]Translation**](Translation.md) | Translations to create on the DNA sequence. Translations are specified by either a combination of &#39;start&#39; and &#39;end&#39; fields, or a list of regions. Both cannot be provided.  | [optional] 
**EntityRegistryId** | Pointer to **string** | Entity registry ID to set for the registered entity. Cannot specify both entityRegistryId and namingStrategy at the same time.  | [optional] 
**NamingStrategy** | Pointer to [**NamingStrategy**](NamingStrategy.md) |  | [optional] 
**RegistryId** | Pointer to **string** | Registry ID into which entity should be registered. this is the ID of the registry which was configured for a particular organization To get available registryIds, use the [/registries endpoint](#/Registry/listRegistries)  Required in order for entities to be created directly in the registry.  | [optional] 

## Methods

### NewDnaSequenceBulkCreate

`func NewDnaSequenceBulkCreate(bases string, isCircular bool, name string, ) *DnaSequenceBulkCreate`

NewDnaSequenceBulkCreate instantiates a new DnaSequenceBulkCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnaSequenceBulkCreateWithDefaults

`func NewDnaSequenceBulkCreateWithDefaults() *DnaSequenceBulkCreate`

NewDnaSequenceBulkCreateWithDefaults instantiates a new DnaSequenceBulkCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *DnaSequenceBulkCreate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *DnaSequenceBulkCreate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *DnaSequenceBulkCreate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *DnaSequenceBulkCreate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAnnotations

`func (o *DnaSequenceBulkCreate) GetAnnotations() []DnaAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *DnaSequenceBulkCreate) GetAnnotationsOk() (*[]DnaAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *DnaSequenceBulkCreate) SetAnnotations(v []DnaAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *DnaSequenceBulkCreate) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetAuthorIds

`func (o *DnaSequenceBulkCreate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *DnaSequenceBulkCreate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *DnaSequenceBulkCreate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *DnaSequenceBulkCreate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetBases

`func (o *DnaSequenceBulkCreate) GetBases() string`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *DnaSequenceBulkCreate) GetBasesOk() (*string, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *DnaSequenceBulkCreate) SetBases(v string)`

SetBases sets Bases field to given value.


### GetCustomFields

`func (o *DnaSequenceBulkCreate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DnaSequenceBulkCreate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DnaSequenceBulkCreate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DnaSequenceBulkCreate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *DnaSequenceBulkCreate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DnaSequenceBulkCreate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DnaSequenceBulkCreate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *DnaSequenceBulkCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *DnaSequenceBulkCreate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *DnaSequenceBulkCreate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *DnaSequenceBulkCreate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *DnaSequenceBulkCreate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetIsCircular

`func (o *DnaSequenceBulkCreate) GetIsCircular() bool`

GetIsCircular returns the IsCircular field if non-nil, zero value otherwise.

### GetIsCircularOk

`func (o *DnaSequenceBulkCreate) GetIsCircularOk() (*bool, bool)`

GetIsCircularOk returns a tuple with the IsCircular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCircular

`func (o *DnaSequenceBulkCreate) SetIsCircular(v bool)`

SetIsCircular sets IsCircular field to given value.


### GetName

`func (o *DnaSequenceBulkCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnaSequenceBulkCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnaSequenceBulkCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPrimers

`func (o *DnaSequenceBulkCreate) GetPrimers() []Primer`

GetPrimers returns the Primers field if non-nil, zero value otherwise.

### GetPrimersOk

`func (o *DnaSequenceBulkCreate) GetPrimersOk() (*[]Primer, bool)`

GetPrimersOk returns a tuple with the Primers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimers

`func (o *DnaSequenceBulkCreate) SetPrimers(v []Primer)`

SetPrimers sets Primers field to given value.

### HasPrimers

`func (o *DnaSequenceBulkCreate) HasPrimers() bool`

HasPrimers returns a boolean if a field has been set.

### GetSchemaId

`func (o *DnaSequenceBulkCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *DnaSequenceBulkCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *DnaSequenceBulkCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *DnaSequenceBulkCreate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetTranslations

`func (o *DnaSequenceBulkCreate) GetTranslations() []Translation`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *DnaSequenceBulkCreate) GetTranslationsOk() (*[]Translation, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *DnaSequenceBulkCreate) SetTranslations(v []Translation)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *DnaSequenceBulkCreate) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *DnaSequenceBulkCreate) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *DnaSequenceBulkCreate) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *DnaSequenceBulkCreate) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *DnaSequenceBulkCreate) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.

### GetNamingStrategy

`func (o *DnaSequenceBulkCreate) GetNamingStrategy() NamingStrategy`

GetNamingStrategy returns the NamingStrategy field if non-nil, zero value otherwise.

### GetNamingStrategyOk

`func (o *DnaSequenceBulkCreate) GetNamingStrategyOk() (*NamingStrategy, bool)`

GetNamingStrategyOk returns a tuple with the NamingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingStrategy

`func (o *DnaSequenceBulkCreate) SetNamingStrategy(v NamingStrategy)`

SetNamingStrategy sets NamingStrategy field to given value.

### HasNamingStrategy

`func (o *DnaSequenceBulkCreate) HasNamingStrategy() bool`

HasNamingStrategy returns a boolean if a field has been set.

### GetRegistryId

`func (o *DnaSequenceBulkCreate) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *DnaSequenceBulkCreate) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *DnaSequenceBulkCreate) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *DnaSequenceBulkCreate) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


