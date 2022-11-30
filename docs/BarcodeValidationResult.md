# BarcodeValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | Pointer to **string** | Barcode to validate. | [optional] 
**IsValid** | Pointer to **bool** | Whether the barcode is valid. | [optional] 
**Message** | Pointer to **NullableString** | If barcode is not valid, a message string explaining the error. | [optional] 

## Methods

### NewBarcodeValidationResult

`func NewBarcodeValidationResult() *BarcodeValidationResult`

NewBarcodeValidationResult instantiates a new BarcodeValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBarcodeValidationResultWithDefaults

`func NewBarcodeValidationResultWithDefaults() *BarcodeValidationResult`

NewBarcodeValidationResultWithDefaults instantiates a new BarcodeValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *BarcodeValidationResult) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *BarcodeValidationResult) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *BarcodeValidationResult) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *BarcodeValidationResult) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetIsValid

`func (o *BarcodeValidationResult) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *BarcodeValidationResult) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *BarcodeValidationResult) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *BarcodeValidationResult) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetMessage

`func (o *BarcodeValidationResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BarcodeValidationResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BarcodeValidationResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BarcodeValidationResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BarcodeValidationResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BarcodeValidationResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


