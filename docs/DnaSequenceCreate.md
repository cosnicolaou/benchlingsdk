# DnaSequenceCreate

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

### NewDnaSequenceCreate

`func NewDnaSequenceCreate(bases string, isCircular bool, name string, ) *DnaSequenceCreate`

NewDnaSequenceCreate instantiates a new DnaSequenceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnaSequenceCreateWithDefaults

`func NewDnaSequenceCreateWithDefaults() *DnaSequenceCreate`

NewDnaSequenceCreateWithDefaults instantiates a new DnaSequenceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *DnaSequenceCreate) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *DnaSequenceCreate) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *DnaSequenceCreate) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *DnaSequenceCreate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAnnotations

`func (o *DnaSequenceCreate) GetAnnotations() []DnaAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *DnaSequenceCreate) GetAnnotationsOk() (*[]DnaAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *DnaSequenceCreate) SetAnnotations(v []DnaAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *DnaSequenceCreate) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetAuthorIds

`func (o *DnaSequenceCreate) GetAuthorIds() []string`

GetAuthorIds returns the AuthorIds field if non-nil, zero value otherwise.

### GetAuthorIdsOk

`func (o *DnaSequenceCreate) GetAuthorIdsOk() (*[]string, bool)`

GetAuthorIdsOk returns a tuple with the AuthorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIds

`func (o *DnaSequenceCreate) SetAuthorIds(v []string)`

SetAuthorIds sets AuthorIds field to given value.

### HasAuthorIds

`func (o *DnaSequenceCreate) HasAuthorIds() bool`

HasAuthorIds returns a boolean if a field has been set.

### GetBases

`func (o *DnaSequenceCreate) GetBases() string`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *DnaSequenceCreate) GetBasesOk() (*string, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *DnaSequenceCreate) SetBases(v string)`

SetBases sets Bases field to given value.


### GetCustomFields

`func (o *DnaSequenceCreate) GetCustomFields() ModelMap`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DnaSequenceCreate) GetCustomFieldsOk() (*ModelMap, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DnaSequenceCreate) SetCustomFields(v ModelMap)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DnaSequenceCreate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFields

`func (o *DnaSequenceCreate) GetFields() ModelMap`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DnaSequenceCreate) GetFieldsOk() (*ModelMap, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DnaSequenceCreate) SetFields(v ModelMap)`

SetFields sets Fields field to given value.

### HasFields

`func (o *DnaSequenceCreate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFolderId

`func (o *DnaSequenceCreate) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *DnaSequenceCreate) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *DnaSequenceCreate) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *DnaSequenceCreate) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetIsCircular

`func (o *DnaSequenceCreate) GetIsCircular() bool`

GetIsCircular returns the IsCircular field if non-nil, zero value otherwise.

### GetIsCircularOk

`func (o *DnaSequenceCreate) GetIsCircularOk() (*bool, bool)`

GetIsCircularOk returns a tuple with the IsCircular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCircular

`func (o *DnaSequenceCreate) SetIsCircular(v bool)`

SetIsCircular sets IsCircular field to given value.


### GetName

`func (o *DnaSequenceCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnaSequenceCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnaSequenceCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPrimers

`func (o *DnaSequenceCreate) GetPrimers() []Primer`

GetPrimers returns the Primers field if non-nil, zero value otherwise.

### GetPrimersOk

`func (o *DnaSequenceCreate) GetPrimersOk() (*[]Primer, bool)`

GetPrimersOk returns a tuple with the Primers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimers

`func (o *DnaSequenceCreate) SetPrimers(v []Primer)`

SetPrimers sets Primers field to given value.

### HasPrimers

`func (o *DnaSequenceCreate) HasPrimers() bool`

HasPrimers returns a boolean if a field has been set.

### GetSchemaId

`func (o *DnaSequenceCreate) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *DnaSequenceCreate) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *DnaSequenceCreate) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *DnaSequenceCreate) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetTranslations

`func (o *DnaSequenceCreate) GetTranslations() []Translation`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *DnaSequenceCreate) GetTranslationsOk() (*[]Translation, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *DnaSequenceCreate) SetTranslations(v []Translation)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *DnaSequenceCreate) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.

### GetEntityRegistryId

`func (o *DnaSequenceCreate) GetEntityRegistryId() string`

GetEntityRegistryId returns the EntityRegistryId field if non-nil, zero value otherwise.

### GetEntityRegistryIdOk

`func (o *DnaSequenceCreate) GetEntityRegistryIdOk() (*string, bool)`

GetEntityRegistryIdOk returns a tuple with the EntityRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRegistryId

`func (o *DnaSequenceCreate) SetEntityRegistryId(v string)`

SetEntityRegistryId sets EntityRegistryId field to given value.

### HasEntityRegistryId

`func (o *DnaSequenceCreate) HasEntityRegistryId() bool`

HasEntityRegistryId returns a boolean if a field has been set.

### GetNamingStrategy

`func (o *DnaSequenceCreate) GetNamingStrategy() NamingStrategy`

GetNamingStrategy returns the NamingStrategy field if non-nil, zero value otherwise.

### GetNamingStrategyOk

`func (o *DnaSequenceCreate) GetNamingStrategyOk() (*NamingStrategy, bool)`

GetNamingStrategyOk returns a tuple with the NamingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingStrategy

`func (o *DnaSequenceCreate) SetNamingStrategy(v NamingStrategy)`

SetNamingStrategy sets NamingStrategy field to given value.

### HasNamingStrategy

`func (o *DnaSequenceCreate) HasNamingStrategy() bool`

HasNamingStrategy returns a boolean if a field has been set.

### GetRegistryId

`func (o *DnaSequenceCreate) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *DnaSequenceCreate) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *DnaSequenceCreate) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *DnaSequenceCreate) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


