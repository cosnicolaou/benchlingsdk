# LocationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**map[string]Field**](Field.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentStorageId** | Pointer to **string** |  | [optional] 

## Methods

### NewLocationUpdate

`func NewLocationUpdate() *LocationUpdate`

NewLocationUpdate instantiates a new LocationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationUpdateWithDefaults

`func NewLocationUpdateWithDefaults() *LocationUpdate`

NewLocationUpdateWithDefaults instantiates a new LocationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *LocationUpdate) GetFields() map[string]Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *LocationUpdate) GetFieldsOk() (*map[string]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *LocationUpdate) SetFields(v map[string]Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *LocationUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetName

`func (o *LocationUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LocationUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentStorageId

`func (o *LocationUpdate) GetParentStorageId() string`

GetParentStorageId returns the ParentStorageId field if non-nil, zero value otherwise.

### GetParentStorageIdOk

`func (o *LocationUpdate) GetParentStorageIdOk() (*string, bool)`

GetParentStorageIdOk returns a tuple with the ParentStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStorageId

`func (o *LocationUpdate) SetParentStorageId(v string)`

SetParentStorageId sets ParentStorageId field to given value.

### HasParentStorageId

`func (o *LocationUpdate) HasParentStorageId() bool`

HasParentStorageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


