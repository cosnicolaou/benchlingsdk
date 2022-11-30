# AssayResultSchemaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModifiedAt** | Pointer to **time.Time** | DateTime the Assay Result Schema was last modified | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAssayResultSchemaAllOf

`func NewAssayResultSchemaAllOf() *AssayResultSchemaAllOf`

NewAssayResultSchemaAllOf instantiates a new AssayResultSchemaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssayResultSchemaAllOfWithDefaults

`func NewAssayResultSchemaAllOfWithDefaults() *AssayResultSchemaAllOf`

NewAssayResultSchemaAllOfWithDefaults instantiates a new AssayResultSchemaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModifiedAt

`func (o *AssayResultSchemaAllOf) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AssayResultSchemaAllOf) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AssayResultSchemaAllOf) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AssayResultSchemaAllOf) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetType

`func (o *AssayResultSchemaAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssayResultSchemaAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssayResultSchemaAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssayResultSchemaAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


