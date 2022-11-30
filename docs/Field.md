# Field

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayValue** | Pointer to **NullableString** |  | [optional] [readonly] 
**IsMulti** | Pointer to **bool** |  | [optional] [readonly] 
**TextValue** | Pointer to **NullableString** |  | [optional] [readonly] 
**Type** | Pointer to [**FieldType**](FieldType.md) |  | [optional] [readonly] 
**Value** | [**NullableFieldValue**](FieldValue.md) |  | 

## Methods

### NewField

`func NewField(value NullableFieldValue, ) *Field`

NewField instantiates a new Field object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldWithDefaults

`func NewFieldWithDefaults() *Field`

NewFieldWithDefaults instantiates a new Field object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayValue

`func (o *Field) GetDisplayValue() string`

GetDisplayValue returns the DisplayValue field if non-nil, zero value otherwise.

### GetDisplayValueOk

`func (o *Field) GetDisplayValueOk() (*string, bool)`

GetDisplayValueOk returns a tuple with the DisplayValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValue

`func (o *Field) SetDisplayValue(v string)`

SetDisplayValue sets DisplayValue field to given value.

### HasDisplayValue

`func (o *Field) HasDisplayValue() bool`

HasDisplayValue returns a boolean if a field has been set.

### SetDisplayValueNil

`func (o *Field) SetDisplayValueNil(b bool)`

 SetDisplayValueNil sets the value for DisplayValue to be an explicit nil

### UnsetDisplayValue
`func (o *Field) UnsetDisplayValue()`

UnsetDisplayValue ensures that no value is present for DisplayValue, not even an explicit nil
### GetIsMulti

`func (o *Field) GetIsMulti() bool`

GetIsMulti returns the IsMulti field if non-nil, zero value otherwise.

### GetIsMultiOk

`func (o *Field) GetIsMultiOk() (*bool, bool)`

GetIsMultiOk returns a tuple with the IsMulti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMulti

`func (o *Field) SetIsMulti(v bool)`

SetIsMulti sets IsMulti field to given value.

### HasIsMulti

`func (o *Field) HasIsMulti() bool`

HasIsMulti returns a boolean if a field has been set.

### GetTextValue

`func (o *Field) GetTextValue() string`

GetTextValue returns the TextValue field if non-nil, zero value otherwise.

### GetTextValueOk

`func (o *Field) GetTextValueOk() (*string, bool)`

GetTextValueOk returns a tuple with the TextValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextValue

`func (o *Field) SetTextValue(v string)`

SetTextValue sets TextValue field to given value.

### HasTextValue

`func (o *Field) HasTextValue() bool`

HasTextValue returns a boolean if a field has been set.

### SetTextValueNil

`func (o *Field) SetTextValueNil(b bool)`

 SetTextValueNil sets the value for TextValue to be an explicit nil

### UnsetTextValue
`func (o *Field) UnsetTextValue()`

UnsetTextValue ensures that no value is present for TextValue, not even an explicit nil
### GetType

`func (o *Field) GetType() FieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Field) GetTypeOk() (*FieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Field) SetType(v FieldType)`

SetType sets Type field to given value.

### HasType

`func (o *Field) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *Field) GetValue() FieldValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Field) GetValueOk() (*FieldValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Field) SetValue(v FieldValue)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *Field) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Field) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


