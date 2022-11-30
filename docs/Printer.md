# Printer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Web address of the printer (either IP address or URL). | [optional] 
**Description** | Pointer to **NullableString** | Short description of the printer. | [optional] 
**Id** | Pointer to **string** | ID of the printer. | [optional] 
**Name** | Pointer to **string** | Name of the printer. | [optional] 
**Port** | Pointer to **NullableInt32** | Port to reach the printer at. | [optional] 
**RegistryId** | Pointer to **string** | ID of the registry associated with this printer. | [optional] 

## Methods

### NewPrinter

`func NewPrinter() *Printer`

NewPrinter instantiates a new Printer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrinterWithDefaults

`func NewPrinterWithDefaults() *Printer`

NewPrinterWithDefaults instantiates a new Printer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Printer) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Printer) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Printer) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Printer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDescription

`func (o *Printer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Printer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Printer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Printer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Printer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Printer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *Printer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Printer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Printer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Printer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Printer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Printer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Printer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Printer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *Printer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Printer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Printer) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Printer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *Printer) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *Printer) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetRegistryId

`func (o *Printer) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *Printer) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *Printer) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *Printer) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


