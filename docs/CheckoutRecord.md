# CheckoutRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to [**NullableCheckoutRecordAssignee**](CheckoutRecordAssignee.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCheckoutRecord

`func NewCheckoutRecord() *CheckoutRecord`

NewCheckoutRecord instantiates a new CheckoutRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutRecordWithDefaults

`func NewCheckoutRecordWithDefaults() *CheckoutRecord`

NewCheckoutRecordWithDefaults instantiates a new CheckoutRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *CheckoutRecord) GetAssignee() CheckoutRecordAssignee`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *CheckoutRecord) GetAssigneeOk() (*CheckoutRecordAssignee, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *CheckoutRecord) SetAssignee(v CheckoutRecordAssignee)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *CheckoutRecord) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *CheckoutRecord) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *CheckoutRecord) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetComment

`func (o *CheckoutRecord) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CheckoutRecord) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CheckoutRecord) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CheckoutRecord) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetModifiedAt

`func (o *CheckoutRecord) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *CheckoutRecord) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *CheckoutRecord) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *CheckoutRecord) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetStatus

`func (o *CheckoutRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckoutRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckoutRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CheckoutRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


