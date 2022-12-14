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

// AutomationProgressStats Processing progress information.
type AutomationProgressStats struct {
	RowsFailed *int32 `json:"rowsFailed,omitempty"`
	RowsSucceeded *int32 `json:"rowsSucceeded,omitempty"`
	RowsUnprocessed *int32 `json:"rowsUnprocessed,omitempty"`
}

// NewAutomationProgressStats instantiates a new AutomationProgressStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationProgressStats() *AutomationProgressStats {
	this := AutomationProgressStats{}
	return &this
}

// NewAutomationProgressStatsWithDefaults instantiates a new AutomationProgressStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationProgressStatsWithDefaults() *AutomationProgressStats {
	this := AutomationProgressStats{}
	return &this
}

// GetRowsFailed returns the RowsFailed field value if set, zero value otherwise.
func (o *AutomationProgressStats) GetRowsFailed() int32 {
	if o == nil || isNil(o.RowsFailed) {
		var ret int32
		return ret
	}
	return *o.RowsFailed
}

// GetRowsFailedOk returns a tuple with the RowsFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationProgressStats) GetRowsFailedOk() (*int32, bool) {
	if o == nil || isNil(o.RowsFailed) {
    return nil, false
	}
	return o.RowsFailed, true
}

// HasRowsFailed returns a boolean if a field has been set.
func (o *AutomationProgressStats) HasRowsFailed() bool {
	if o != nil && !isNil(o.RowsFailed) {
		return true
	}

	return false
}

// SetRowsFailed gets a reference to the given int32 and assigns it to the RowsFailed field.
func (o *AutomationProgressStats) SetRowsFailed(v int32) {
	o.RowsFailed = &v
}

// GetRowsSucceeded returns the RowsSucceeded field value if set, zero value otherwise.
func (o *AutomationProgressStats) GetRowsSucceeded() int32 {
	if o == nil || isNil(o.RowsSucceeded) {
		var ret int32
		return ret
	}
	return *o.RowsSucceeded
}

// GetRowsSucceededOk returns a tuple with the RowsSucceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationProgressStats) GetRowsSucceededOk() (*int32, bool) {
	if o == nil || isNil(o.RowsSucceeded) {
    return nil, false
	}
	return o.RowsSucceeded, true
}

// HasRowsSucceeded returns a boolean if a field has been set.
func (o *AutomationProgressStats) HasRowsSucceeded() bool {
	if o != nil && !isNil(o.RowsSucceeded) {
		return true
	}

	return false
}

// SetRowsSucceeded gets a reference to the given int32 and assigns it to the RowsSucceeded field.
func (o *AutomationProgressStats) SetRowsSucceeded(v int32) {
	o.RowsSucceeded = &v
}

// GetRowsUnprocessed returns the RowsUnprocessed field value if set, zero value otherwise.
func (o *AutomationProgressStats) GetRowsUnprocessed() int32 {
	if o == nil || isNil(o.RowsUnprocessed) {
		var ret int32
		return ret
	}
	return *o.RowsUnprocessed
}

// GetRowsUnprocessedOk returns a tuple with the RowsUnprocessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationProgressStats) GetRowsUnprocessedOk() (*int32, bool) {
	if o == nil || isNil(o.RowsUnprocessed) {
    return nil, false
	}
	return o.RowsUnprocessed, true
}

// HasRowsUnprocessed returns a boolean if a field has been set.
func (o *AutomationProgressStats) HasRowsUnprocessed() bool {
	if o != nil && !isNil(o.RowsUnprocessed) {
		return true
	}

	return false
}

// SetRowsUnprocessed gets a reference to the given int32 and assigns it to the RowsUnprocessed field.
func (o *AutomationProgressStats) SetRowsUnprocessed(v int32) {
	o.RowsUnprocessed = &v
}

func (o AutomationProgressStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RowsFailed) {
		toSerialize["rowsFailed"] = o.RowsFailed
	}
	if !isNil(o.RowsSucceeded) {
		toSerialize["rowsSucceeded"] = o.RowsSucceeded
	}
	if !isNil(o.RowsUnprocessed) {
		toSerialize["rowsUnprocessed"] = o.RowsUnprocessed
	}
	return json.Marshal(toSerialize)
}

type NullableAutomationProgressStats struct {
	value *AutomationProgressStats
	isSet bool
}

func (v NullableAutomationProgressStats) Get() *AutomationProgressStats {
	return v.value
}

func (v *NullableAutomationProgressStats) Set(val *AutomationProgressStats) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomationProgressStats) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomationProgressStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomationProgressStats(val *AutomationProgressStats) *NullableAutomationProgressStats {
	return &NullableAutomationProgressStats{value: val, isSet: true}
}

func (v NullableAutomationProgressStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomationProgressStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


