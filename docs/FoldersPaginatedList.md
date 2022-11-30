# FoldersPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Folders** | Pointer to [**[]Folder**](Folder.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewFoldersPaginatedList

`func NewFoldersPaginatedList() *FoldersPaginatedList`

NewFoldersPaginatedList instantiates a new FoldersPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoldersPaginatedListWithDefaults

`func NewFoldersPaginatedListWithDefaults() *FoldersPaginatedList`

NewFoldersPaginatedListWithDefaults instantiates a new FoldersPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolders

`func (o *FoldersPaginatedList) GetFolders() []Folder`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *FoldersPaginatedList) GetFoldersOk() (*[]Folder, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *FoldersPaginatedList) SetFolders(v []Folder)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *FoldersPaginatedList) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### GetNextToken

`func (o *FoldersPaginatedList) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *FoldersPaginatedList) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *FoldersPaginatedList) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *FoldersPaginatedList) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


