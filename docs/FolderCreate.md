# FolderCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new folder. | 
**ParentFolderId** | **string** | The ID of the parent folder. | 

## Methods

### NewFolderCreate

`func NewFolderCreate(name string, parentFolderId string, ) *FolderCreate`

NewFolderCreate instantiates a new FolderCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderCreateWithDefaults

`func NewFolderCreateWithDefaults() *FolderCreate`

NewFolderCreateWithDefaults instantiates a new FolderCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FolderCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FolderCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FolderCreate) SetName(v string)`

SetName sets Name field to given value.


### GetParentFolderId

`func (o *FolderCreate) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *FolderCreate) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *FolderCreate) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


