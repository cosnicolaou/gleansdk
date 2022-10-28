/*
Glean Indexing API

# Introduction In addition to the data sources that Glean has built-in support for, Glean also provides a REST API that enables customers to put arbitrary content in the search index. This is useful, for example, for doing permissions-aware search over content in internal tools that reside on-prem as well as for searching over applications that Glean does not currently support first class. In addition these APIs allow the customer to push organization data (people info, organization structure etc) into Glean.  # Early Access Please note that we are continually evolving the system so these APIs are considered “early access” and are subject to change. 

API version: 0.9.0
Contact: support@glean.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gleansdk

import (
	"encoding/json"
)

// DatasourceGroupDefinition describes a group in the datasource
type DatasourceGroupDefinition struct {
	// name of the group. Should be unique among all groups for the datasource, and cannot have spaces.
	Name string `json:"name"`
}

// NewDatasourceGroupDefinition instantiates a new DatasourceGroupDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasourceGroupDefinition(name string) *DatasourceGroupDefinition {
	this := DatasourceGroupDefinition{}
	this.Name = name
	return &this
}

// NewDatasourceGroupDefinitionWithDefaults instantiates a new DatasourceGroupDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasourceGroupDefinitionWithDefaults() *DatasourceGroupDefinition {
	this := DatasourceGroupDefinition{}
	return &this
}

// GetName returns the Name field value
func (o *DatasourceGroupDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatasourceGroupDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatasourceGroupDefinition) SetName(v string) {
	o.Name = v
}

func (o DatasourceGroupDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDatasourceGroupDefinition struct {
	value *DatasourceGroupDefinition
	isSet bool
}

func (v NullableDatasourceGroupDefinition) Get() *DatasourceGroupDefinition {
	return v.value
}

func (v *NullableDatasourceGroupDefinition) Set(val *DatasourceGroupDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasourceGroupDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasourceGroupDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasourceGroupDefinition(val *DatasourceGroupDefinition) *NullableDatasourceGroupDefinition {
	return &NullableDatasourceGroupDefinition{value: val, isSet: true}
}

func (v NullableDatasourceGroupDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasourceGroupDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


