# Measurement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Units** | **NullableString** | Can only be null if value is also null | 
**Value** | **NullableFloat32** | Can only be null if units is also null | 

## Methods

### NewMeasurement

`func NewMeasurement(units NullableString, value NullableFloat32, ) *Measurement`

NewMeasurement instantiates a new Measurement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementWithDefaults

`func NewMeasurementWithDefaults() *Measurement`

NewMeasurementWithDefaults instantiates a new Measurement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnits

`func (o *Measurement) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Measurement) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Measurement) SetUnits(v string)`

SetUnits sets Units field to given value.


### SetUnitsNil

`func (o *Measurement) SetUnitsNil(b bool)`

 SetUnitsNil sets the value for Units to be an explicit nil

### UnsetUnits
`func (o *Measurement) UnsetUnits()`

UnsetUnits ensures that no value is present for Units, not even an explicit nil
### GetValue

`func (o *Measurement) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Measurement) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Measurement) SetValue(v float32)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *Measurement) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Measurement) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


