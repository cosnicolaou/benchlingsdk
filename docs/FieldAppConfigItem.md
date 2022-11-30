# FieldAppConfigItem

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
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFieldAppConfigItem

`func NewFieldAppConfigItem() *FieldAppConfigItem`

NewFieldAppConfigItem instantiates a new FieldAppConfigItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldAppConfigItemWithDefaults

`func NewFieldAppConfigItemWithDefaults() *FieldAppConfigItem`

NewFieldAppConfigItemWithDefaults instantiates a new FieldAppConfigItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *FieldAppConfigItem) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *FieldAppConfigItem) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *FieldAppConfigItem) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *FieldAppConfigItem) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetApp

`func (o *FieldAppConfigItem) GetApp() AppConfigItemApiMixinApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *FieldAppConfigItem) GetAppOk() (*AppConfigItemApiMixinApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *FieldAppConfigItem) SetApp(v AppConfigItemApiMixinApp)`

SetApp sets App field to given value.

### HasApp

`func (o *FieldAppConfigItem) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetId

`func (o *FieldAppConfigItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FieldAppConfigItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FieldAppConfigItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FieldAppConfigItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *FieldAppConfigItem) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *FieldAppConfigItem) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *FieldAppConfigItem) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *FieldAppConfigItem) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetPath

`func (o *FieldAppConfigItem) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FieldAppConfigItem) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FieldAppConfigItem) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FieldAppConfigItem) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *FieldAppConfigItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldAppConfigItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldAppConfigItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FieldAppConfigItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *FieldAppConfigItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FieldAppConfigItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FieldAppConfigItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FieldAppConfigItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequiredConfig

`func (o *FieldAppConfigItem) GetRequiredConfig() bool`

GetRequiredConfig returns the RequiredConfig field if non-nil, zero value otherwise.

### GetRequiredConfigOk

`func (o *FieldAppConfigItem) GetRequiredConfigOk() (*bool, bool)`

GetRequiredConfigOk returns a tuple with the RequiredConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredConfig

`func (o *FieldAppConfigItem) SetRequiredConfig(v bool)`

SetRequiredConfig sets RequiredConfig field to given value.

### HasRequiredConfig

`func (o *FieldAppConfigItem) HasRequiredConfig() bool`

HasRequiredConfig returns a boolean if a field has been set.

### GetLinkedResource

`func (o *FieldAppConfigItem) GetLinkedResource() LinkedAppConfigResourceMixinLinkedResource`

GetLinkedResource returns the LinkedResource field if non-nil, zero value otherwise.

### GetLinkedResourceOk

`func (o *FieldAppConfigItem) GetLinkedResourceOk() (*LinkedAppConfigResourceMixinLinkedResource, bool)`

GetLinkedResourceOk returns a tuple with the LinkedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedResource

`func (o *FieldAppConfigItem) SetLinkedResource(v LinkedAppConfigResourceMixinLinkedResource)`

SetLinkedResource sets LinkedResource field to given value.

### HasLinkedResource

`func (o *FieldAppConfigItem) HasLinkedResource() bool`

HasLinkedResource returns a boolean if a field has been set.

### SetLinkedResourceNil

`func (o *FieldAppConfigItem) SetLinkedResourceNil(b bool)`

 SetLinkedResourceNil sets the value for LinkedResource to be an explicit nil

### UnsetLinkedResource
`func (o *FieldAppConfigItem) UnsetLinkedResource()`

UnsetLinkedResource ensures that no value is present for LinkedResource, not even an explicit nil
### GetValue

`func (o *FieldAppConfigItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FieldAppConfigItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FieldAppConfigItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FieldAppConfigItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *FieldAppConfigItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FieldAppConfigItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


