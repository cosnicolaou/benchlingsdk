# EntryTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnLabels** | Pointer to **[]string** | Array of strings, with one item per column. Defaults to null, if the user is using the default, but is set if the user has given a custom name to the column.  | [optional] 
**Name** | Pointer to **string** | Name of the table - defaults to e.g. Table1 but can be renamed.  | [optional] 
**Rows** | Pointer to [**[]EntryTableRow**](EntryTableRow.md) | Array of row objects. | [optional] 

## Methods

### NewEntryTable

`func NewEntryTable() *EntryTable`

NewEntryTable instantiates a new EntryTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryTableWithDefaults

`func NewEntryTableWithDefaults() *EntryTable`

NewEntryTableWithDefaults instantiates a new EntryTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnLabels

`func (o *EntryTable) GetColumnLabels() []*string`

GetColumnLabels returns the ColumnLabels field if non-nil, zero value otherwise.

### GetColumnLabelsOk

`func (o *EntryTable) GetColumnLabelsOk() (*[]*string, bool)`

GetColumnLabelsOk returns a tuple with the ColumnLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnLabels

`func (o *EntryTable) SetColumnLabels(v []*string)`

SetColumnLabels sets ColumnLabels field to given value.

### HasColumnLabels

`func (o *EntryTable) HasColumnLabels() bool`

HasColumnLabels returns a boolean if a field has been set.

### GetName

`func (o *EntryTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntryTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntryTable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntryTable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRows

`func (o *EntryTable) GetRows() []EntryTableRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *EntryTable) GetRowsOk() (*[]EntryTableRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *EntryTable) SetRows(v []EntryTableRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *EntryTable) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


