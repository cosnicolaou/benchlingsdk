# LinkedAppConfigResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The API ID of the linked resource | [optional] 
**Name** | Pointer to **string** | The name of the resource in Benchling | [optional] 
**InaccessibleId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The type of this inaccessible item. Example values: \&quot;custom_entity\&quot;, \&quot;container\&quot;, \&quot;plate\&quot;, \&quot;dna_sequence\&quot;  | [optional] 

## Methods

### NewLinkedAppConfigResource

`func NewLinkedAppConfigResource() *LinkedAppConfigResource`

NewLinkedAppConfigResource instantiates a new LinkedAppConfigResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedAppConfigResourceWithDefaults

`func NewLinkedAppConfigResourceWithDefaults() *LinkedAppConfigResource`

NewLinkedAppConfigResourceWithDefaults instantiates a new LinkedAppConfigResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkedAppConfigResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedAppConfigResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedAppConfigResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LinkedAppConfigResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LinkedAppConfigResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinkedAppConfigResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinkedAppConfigResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LinkedAppConfigResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInaccessibleId

`func (o *LinkedAppConfigResource) GetInaccessibleId() string`

GetInaccessibleId returns the InaccessibleId field if non-nil, zero value otherwise.

### GetInaccessibleIdOk

`func (o *LinkedAppConfigResource) GetInaccessibleIdOk() (*string, bool)`

GetInaccessibleIdOk returns a tuple with the InaccessibleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInaccessibleId

`func (o *LinkedAppConfigResource) SetInaccessibleId(v string)`

SetInaccessibleId sets InaccessibleId field to given value.

### HasInaccessibleId

`func (o *LinkedAppConfigResource) HasInaccessibleId() bool`

HasInaccessibleId returns a boolean if a field has been set.

### GetType

`func (o *LinkedAppConfigResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedAppConfigResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedAppConfigResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkedAppConfigResource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


