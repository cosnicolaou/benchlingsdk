/*
Benchling API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package benchlingsdk

import (
	"encoding/json"
)

// TeamsPaginatedList struct for TeamsPaginatedList
type TeamsPaginatedList struct {
	NextToken *string `json:"nextToken,omitempty"`
	Teams []Team `json:"teams,omitempty"`
}

// NewTeamsPaginatedList instantiates a new TeamsPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamsPaginatedList() *TeamsPaginatedList {
	this := TeamsPaginatedList{}
	return &this
}

// NewTeamsPaginatedListWithDefaults instantiates a new TeamsPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamsPaginatedListWithDefaults() *TeamsPaginatedList {
	this := TeamsPaginatedList{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *TeamsPaginatedList) GetNextToken() string {
	if o == nil || isNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsPaginatedList) GetNextTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextToken) {
    return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *TeamsPaginatedList) HasNextToken() bool {
	if o != nil && !isNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *TeamsPaginatedList) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *TeamsPaginatedList) GetTeams() []Team {
	if o == nil || isNil(o.Teams) {
		var ret []Team
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsPaginatedList) GetTeamsOk() ([]Team, bool) {
	if o == nil || isNil(o.Teams) {
    return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *TeamsPaginatedList) HasTeams() bool {
	if o != nil && !isNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []Team and assigns it to the Teams field.
func (o *TeamsPaginatedList) SetTeams(v []Team) {
	o.Teams = v
}

func (o TeamsPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !isNil(o.Teams) {
		toSerialize["teams"] = o.Teams
	}
	return json.Marshal(toSerialize)
}

type NullableTeamsPaginatedList struct {
	value *TeamsPaginatedList
	isSet bool
}

func (v NullableTeamsPaginatedList) Get() *TeamsPaginatedList {
	return v.value
}

func (v *NullableTeamsPaginatedList) Set(val *TeamsPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamsPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamsPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamsPaginatedList(val *TeamsPaginatedList) *NullableTeamsPaginatedList {
	return &NullableTeamsPaginatedList{value: val, isSet: true}
}

func (v NullableTeamsPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamsPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


