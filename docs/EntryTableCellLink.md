# EntryTableCellLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | For linked Benchling resources, this will be the ID of that resource (e.g., &#39;seq_RhYGVnHF&#39;). Omitted for \&quot;link\&quot; types.  | [optional] 
**Type** | Pointer to **string** | The type of resource being linked. For hyperlinks: &#39;link&#39;. For linked Benchling resources, one of: &#39;user&#39;, &#39;request&#39;, &#39;entry&#39;, &#39;stage_entry&#39;, &#39;protocol&#39;, &#39;workflow&#39;, &#39;custom_entity&#39;, &#39;aa_sequence&#39;, &#39;dna_sequence&#39;, &#39;batch&#39;, &#39;box&#39;, &#39;container&#39;, &#39;location&#39;, &#39;plate&#39;.  | [optional] 
**WebURL** | Pointer to **NullableString** | Canonical URL of the linked Benchling resource (if you have at least READ authorization for that resource), or the explicit URL provided as hyperlink for \&quot;link\&quot; types. Note: locations do not currently have a URL.  | [optional] 

## Methods

### NewEntryTableCellLink

`func NewEntryTableCellLink() *EntryTableCellLink`

NewEntryTableCellLink instantiates a new EntryTableCellLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryTableCellLinkWithDefaults

`func NewEntryTableCellLinkWithDefaults() *EntryTableCellLink`

NewEntryTableCellLinkWithDefaults instantiates a new EntryTableCellLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntryTableCellLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntryTableCellLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntryTableCellLink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntryTableCellLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *EntryTableCellLink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntryTableCellLink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntryTableCellLink) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntryTableCellLink) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWebURL

`func (o *EntryTableCellLink) GetWebURL() string`

GetWebURL returns the WebURL field if non-nil, zero value otherwise.

### GetWebURLOk

`func (o *EntryTableCellLink) GetWebURLOk() (*string, bool)`

GetWebURLOk returns a tuple with the WebURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebURL

`func (o *EntryTableCellLink) SetWebURL(v string)`

SetWebURL sets WebURL field to given value.

### HasWebURL

`func (o *EntryTableCellLink) HasWebURL() bool`

HasWebURL returns a boolean if a field has been set.

### SetWebURLNil

`func (o *EntryTableCellLink) SetWebURLNil(b bool)`

 SetWebURLNil sets the value for WebURL to be an explicit nil

### UnsetWebURL
`func (o *EntryTableCellLink) UnsetWebURL()`

UnsetWebURL ensures that no value is present for WebURL, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


