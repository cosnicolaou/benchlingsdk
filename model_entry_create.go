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

// EntryCreate struct for EntryCreate
type EntryCreate struct {
	AuthorIds *EntryCreateAuthorIds `json:"authorIds,omitempty"`
	// Custom fields to add to the entry
	CustomFields *ModelMap `json:"customFields,omitempty"`
	// ID of the template to clone the entry from
	EntryTemplateId *string `json:"entryTemplateId,omitempty"`
	// Fields to set on the entry. Must correspond with the schema's field definitions. 
	Fields *ModelMap `json:"fields,omitempty"`
	// ID of the folder that will contain the entry
	FolderId string `json:"folderId"`
	// An array of table API IDs and blob id pairs to seed tables from the template while creating the entry. The entryTemplateId parameter must be set to use this parameter. The table API IDs should be the API Identifiers of the tables in the given template. - If a template table has one row, the values in that row act as default values for cloned entries. - If a template table has multiple rows, there is no default value and those rows are added to the cloned entry along with the provided csv data. - If a table has default values, they will be populated in any respective undefined columns in the csv data. - If a table has no default values, undefined columns from csv data will be empty. - If no csv data is provided for a table, the table in the entry will be populated with whatever values are in the respective template table. 
	InitialTables []InitialTable `json:"initialTables,omitempty"`
	// Name of the entry
	Name string `json:"name"`
	// ID of the entry's schema
	SchemaId *string `json:"schemaId,omitempty"`
}

// NewEntryCreate instantiates a new EntryCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntryCreate(folderId string, name string) *EntryCreate {
	this := EntryCreate{}
	this.FolderId = folderId
	this.Name = name
	return &this
}

// NewEntryCreateWithDefaults instantiates a new EntryCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryCreateWithDefaults() *EntryCreate {
	this := EntryCreate{}
	return &this
}

// GetAuthorIds returns the AuthorIds field value if set, zero value otherwise.
func (o *EntryCreate) GetAuthorIds() EntryCreateAuthorIds {
	if o == nil || isNil(o.AuthorIds) {
		var ret EntryCreateAuthorIds
		return ret
	}
	return *o.AuthorIds
}

// GetAuthorIdsOk returns a tuple with the AuthorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryCreate) GetAuthorIdsOk() (*EntryCreateAuthorIds, bool) {
	if o == nil || isNil(o.AuthorIds) {
    return nil, false
	}
	return o.AuthorIds, true
}

// HasAuthorIds returns a boolean if a field has been set.
func (o *EntryCreate) HasAuthorIds() bool {
	if o != nil && !isNil(o.AuthorIds) {
		return true
	}

	return false
}

// SetAuthorIds gets a reference to the given EntryCreateAuthorIds and assigns it to the AuthorIds field.
func (o *EntryCreate) SetAuthorIds(v EntryCreateAuthorIds) {
	o.AuthorIds = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *EntryCreate) GetCustomFields() ModelMap {
	if o == nil || isNil(o.CustomFields) {
		var ret ModelMap
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryCreate) GetCustomFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.CustomFields) {
    return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *EntryCreate) HasCustomFields() bool {
	if o != nil && !isNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given ModelMap and assigns it to the CustomFields field.
func (o *EntryCreate) SetCustomFields(v ModelMap) {
	o.CustomFields = &v
}

// GetEntryTemplateId returns the EntryTemplateId field value if set, zero value otherwise.
func (o *EntryCreate) GetEntryTemplateId() string {
	if o == nil || isNil(o.EntryTemplateId) {
		var ret string
		return ret
	}
	return *o.EntryTemplateId
}

// GetEntryTemplateIdOk returns a tuple with the EntryTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryCreate) GetEntryTemplateIdOk() (*string, bool) {
	if o == nil || isNil(o.EntryTemplateId) {
    return nil, false
	}
	return o.EntryTemplateId, true
}

// HasEntryTemplateId returns a boolean if a field has been set.
func (o *EntryCreate) HasEntryTemplateId() bool {
	if o != nil && !isNil(o.EntryTemplateId) {
		return true
	}

	return false
}

// SetEntryTemplateId gets a reference to the given string and assigns it to the EntryTemplateId field.
func (o *EntryCreate) SetEntryTemplateId(v string) {
	o.EntryTemplateId = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *EntryCreate) GetFields() ModelMap {
	if o == nil || isNil(o.Fields) {
		var ret ModelMap
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryCreate) GetFieldsOk() (*ModelMap, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *EntryCreate) HasFields() bool {
	if o != nil && !isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ModelMap and assigns it to the Fields field.
func (o *EntryCreate) SetFields(v ModelMap) {
	o.Fields = &v
}

// GetFolderId returns the FolderId field value
func (o *EntryCreate) GetFolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value
// and a boolean to check if the value has been set.
func (o *EntryCreate) GetFolderIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FolderId, true
}

// SetFolderId sets field value
func (o *EntryCreate) SetFolderId(v string) {
	o.FolderId = v
}

// GetInitialTables returns the InitialTables field value if set, zero value otherwise.
func (o *EntryCreate) GetInitialTables() []InitialTable {
	if o == nil || isNil(o.InitialTables) {
		var ret []InitialTable
		return ret
	}
	return o.InitialTables
}

// GetInitialTablesOk returns a tuple with the InitialTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryCreate) GetInitialTablesOk() ([]InitialTable, bool) {
	if o == nil || isNil(o.InitialTables) {
    return nil, false
	}
	return o.InitialTables, true
}

// HasInitialTables returns a boolean if a field has been set.
func (o *EntryCreate) HasInitialTables() bool {
	if o != nil && !isNil(o.InitialTables) {
		return true
	}

	return false
}

// SetInitialTables gets a reference to the given []InitialTable and assigns it to the InitialTables field.
func (o *EntryCreate) SetInitialTables(v []InitialTable) {
	o.InitialTables = v
}

// GetName returns the Name field value
func (o *EntryCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntryCreate) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntryCreate) SetName(v string) {
	o.Name = v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *EntryCreate) GetSchemaId() string {
	if o == nil || isNil(o.SchemaId) {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryCreate) GetSchemaIdOk() (*string, bool) {
	if o == nil || isNil(o.SchemaId) {
    return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *EntryCreate) HasSchemaId() bool {
	if o != nil && !isNil(o.SchemaId) {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *EntryCreate) SetSchemaId(v string) {
	o.SchemaId = &v
}

func (o EntryCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AuthorIds) {
		toSerialize["authorIds"] = o.AuthorIds
	}
	if !isNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !isNil(o.EntryTemplateId) {
		toSerialize["entryTemplateId"] = o.EntryTemplateId
	}
	if !isNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if true {
		toSerialize["folderId"] = o.FolderId
	}
	if !isNil(o.InitialTables) {
		toSerialize["initialTables"] = o.InitialTables
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.SchemaId) {
		toSerialize["schemaId"] = o.SchemaId
	}
	return json.Marshal(toSerialize)
}

type NullableEntryCreate struct {
	value *EntryCreate
	isSet bool
}

func (v NullableEntryCreate) Get() *EntryCreate {
	return v.value
}

func (v *NullableEntryCreate) Set(val *EntryCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryCreate(val *EntryCreate) *NullableEntryCreate {
	return &NullableEntryCreate{value: val, isSet: true}
}

func (v NullableEntryCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


