# InaccessibleResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InaccessibleId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The type of this inaccessible item. Example values: \&quot;custom_entity\&quot;, \&quot;container\&quot;, \&quot;plate\&quot;, \&quot;dna_sequence\&quot;  | [optional] 

## Methods

### NewInaccessibleResource

`func NewInaccessibleResource() *InaccessibleResource`

NewInaccessibleResource instantiates a new InaccessibleResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInaccessibleResourceWithDefaults

`func NewInaccessibleResourceWithDefaults() *InaccessibleResource`

NewInaccessibleResourceWithDefaults instantiates a new InaccessibleResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInaccessibleId

`func (o *InaccessibleResource) GetInaccessibleId() string`

GetInaccessibleId returns the InaccessibleId field if non-nil, zero value otherwise.

### GetInaccessibleIdOk

`func (o *InaccessibleResource) GetInaccessibleIdOk() (*string, bool)`

GetInaccessibleIdOk returns a tuple with the InaccessibleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInaccessibleId

`func (o *InaccessibleResource) SetInaccessibleId(v string)`

SetInaccessibleId sets InaccessibleId field to given value.

### HasInaccessibleId

`func (o *InaccessibleResource) HasInaccessibleId() bool`

HasInaccessibleId returns a boolean if a field has been set.

### GetType

`func (o *InaccessibleResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InaccessibleResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InaccessibleResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InaccessibleResource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


