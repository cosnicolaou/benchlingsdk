# EntitySchemaAppConfigItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiURL** | Pointer to **string** |  | [optional] [readonly] 
**App** | Pointer to [**AppConfigItemApiMixinApp**](AppConfigItemApiMixinApp.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | DateTime the app config item was last modified | [optional] [readonly] 
**Path** | Pointer to **[]string** | Array-based representation of config item&#39;s location in the tree in order from top to bottom. | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RequiredConfig** | Pointer to **bool** |  | [optional] 
**LinkedResource** | Pointer to [**NullableLinkedAppConfigResourceMixinLinkedResource**](LinkedAppConfigResourceMixinLinkedResource.md) |  | [optional] 
**Subtype** | Pointer to [**SchemaDependencySubtypes**](SchemaDependencySubtypes.md) |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEntitySchemaAppConfigItem

`func NewEntitySchemaAppConfigItem() *EntitySchemaAppConfigItem`

NewEntitySchemaAppConfigItem instantiates a new EntitySchemaAppConfigItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySchemaAppConfigItemWithDefaults

`func NewEntitySchemaAppConfigItemWithDefaults() *EntitySchemaAppConfigItem`

NewEntitySchemaAppConfigItemWithDefaults instantiates a new EntitySchemaAppConfigItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *EntitySchemaAppConfigItem) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *EntitySchemaAppConfigItem) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *EntitySchemaAppConfigItem) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *EntitySchemaAppConfigItem) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetApp

`func (o *EntitySchemaAppConfigItem) GetApp() AppConfigItemApiMixinApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *EntitySchemaAppConfigItem) GetAppOk() (*AppConfigItemApiMixinApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *EntitySchemaAppConfigItem) SetApp(v AppConfigItemApiMixinApp)`

SetApp sets App field to given value.

### HasApp

`func (o *EntitySchemaAppConfigItem) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetId

`func (o *EntitySchemaAppConfigItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitySchemaAppConfigItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitySchemaAppConfigItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitySchemaAppConfigItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *EntitySchemaAppConfigItem) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *EntitySchemaAppConfigItem) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *EntitySchemaAppConfigItem) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *EntitySchemaAppConfigItem) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetPath

`func (o *EntitySchemaAppConfigItem) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EntitySchemaAppConfigItem) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EntitySchemaAppConfigItem) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *EntitySchemaAppConfigItem) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *EntitySchemaAppConfigItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitySchemaAppConfigItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitySchemaAppConfigItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntitySchemaAppConfigItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *EntitySchemaAppConfigItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitySchemaAppConfigItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitySchemaAppConfigItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitySchemaAppConfigItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequiredConfig

`func (o *EntitySchemaAppConfigItem) GetRequiredConfig() bool`

GetRequiredConfig returns the RequiredConfig field if non-nil, zero value otherwise.

### GetRequiredConfigOk

`func (o *EntitySchemaAppConfigItem) GetRequiredConfigOk() (*bool, bool)`

GetRequiredConfigOk returns a tuple with the RequiredConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredConfig

`func (o *EntitySchemaAppConfigItem) SetRequiredConfig(v bool)`

SetRequiredConfig sets RequiredConfig field to given value.

### HasRequiredConfig

`func (o *EntitySchemaAppConfigItem) HasRequiredConfig() bool`

HasRequiredConfig returns a boolean if a field has been set.

### GetLinkedResource

`func (o *EntitySchemaAppConfigItem) GetLinkedResource() LinkedAppConfigResourceMixinLinkedResource`

GetLinkedResource returns the LinkedResource field if non-nil, zero value otherwise.

### GetLinkedResourceOk

`func (o *EntitySchemaAppConfigItem) GetLinkedResourceOk() (*LinkedAppConfigResourceMixinLinkedResource, bool)`

GetLinkedResourceOk returns a tuple with the LinkedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedResource

`func (o *EntitySchemaAppConfigItem) SetLinkedResource(v LinkedAppConfigResourceMixinLinkedResource)`

SetLinkedResource sets LinkedResource field to given value.

### HasLinkedResource

`func (o *EntitySchemaAppConfigItem) HasLinkedResource() bool`

HasLinkedResource returns a boolean if a field has been set.

### SetLinkedResourceNil

`func (o *EntitySchemaAppConfigItem) SetLinkedResourceNil(b bool)`

 SetLinkedResourceNil sets the value for LinkedResource to be an explicit nil

### UnsetLinkedResource
`func (o *EntitySchemaAppConfigItem) UnsetLinkedResource()`

UnsetLinkedResource ensures that no value is present for LinkedResource, not even an explicit nil
### GetSubtype

`func (o *EntitySchemaAppConfigItem) GetSubtype() SchemaDependencySubtypes`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *EntitySchemaAppConfigItem) GetSubtypeOk() (*SchemaDependencySubtypes, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *EntitySchemaAppConfigItem) SetSubtype(v SchemaDependencySubtypes)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *EntitySchemaAppConfigItem) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetValue

`func (o *EntitySchemaAppConfigItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EntitySchemaAppConfigItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EntitySchemaAppConfigItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EntitySchemaAppConfigItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *EntitySchemaAppConfigItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *EntitySchemaAppConfigItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


