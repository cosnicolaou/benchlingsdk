# InitialTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsvData** | Pointer to **string** | blobId of an uploaded csv blob. The CSV should be formatted with column headers of &#x60;columnId&#x60; which can be found in the [EntryTemplate](#/components/schemas/EntryTemplate). For more information on uploading a blob, [click here](https://docs.benchling.com/docs/uploading-a-blob-to-benchling). | [optional] 
**TemplateTableID** | Pointer to **string** | Template table API ID | [optional] 

## Methods

### NewInitialTable

`func NewInitialTable() *InitialTable`

NewInitialTable instantiates a new InitialTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitialTableWithDefaults

`func NewInitialTableWithDefaults() *InitialTable`

NewInitialTableWithDefaults instantiates a new InitialTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsvData

`func (o *InitialTable) GetCsvData() string`

GetCsvData returns the CsvData field if non-nil, zero value otherwise.

### GetCsvDataOk

`func (o *InitialTable) GetCsvDataOk() (*string, bool)`

GetCsvDataOk returns a tuple with the CsvData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvData

`func (o *InitialTable) SetCsvData(v string)`

SetCsvData sets CsvData field to given value.

### HasCsvData

`func (o *InitialTable) HasCsvData() bool`

HasCsvData returns a boolean if a field has been set.

### GetTemplateTableID

`func (o *InitialTable) GetTemplateTableID() string`

GetTemplateTableID returns the TemplateTableID field if non-nil, zero value otherwise.

### GetTemplateTableIDOk

`func (o *InitialTable) GetTemplateTableIDOk() (*string, bool)`

GetTemplateTableIDOk returns a tuple with the TemplateTableID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateTableID

`func (o *InitialTable) SetTemplateTableID(v string)`

SetTemplateTableID sets TemplateTableID field to given value.

### HasTemplateTableID

`func (o *InitialTable) HasTemplateTableID() bool`

HasTemplateTableID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


