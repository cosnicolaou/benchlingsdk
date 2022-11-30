# FoldersArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderIds** | **[]string** | A list of folder IDs to archive. | 
**Reason** | **string** | The reason for archiving the provided folders. Accepted reasons may differ based on tenant configuration.  | 

## Methods

### NewFoldersArchive

`func NewFoldersArchive(folderIds []string, reason string, ) *FoldersArchive`

NewFoldersArchive instantiates a new FoldersArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoldersArchiveWithDefaults

`func NewFoldersArchiveWithDefaults() *FoldersArchive`

NewFoldersArchiveWithDefaults instantiates a new FoldersArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderIds

`func (o *FoldersArchive) GetFolderIds() []string`

GetFolderIds returns the FolderIds field if non-nil, zero value otherwise.

### GetFolderIdsOk

`func (o *FoldersArchive) GetFolderIdsOk() (*[]string, bool)`

GetFolderIdsOk returns a tuple with the FolderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderIds

`func (o *FoldersArchive) SetFolderIds(v []string)`

SetFolderIds sets FolderIds field to given value.


### GetReason

`func (o *FoldersArchive) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *FoldersArchive) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *FoldersArchive) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


