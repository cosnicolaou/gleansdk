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

// BulkIndexTeamsRequestAllOf struct for BulkIndexTeamsRequestAllOf
type BulkIndexTeamsRequestAllOf struct {
	// Batch of team information
	Teams []TeamInfoDefinition `json:"teams"`
}

// NewBulkIndexTeamsRequestAllOf instantiates a new BulkIndexTeamsRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkIndexTeamsRequestAllOf(teams []TeamInfoDefinition) *BulkIndexTeamsRequestAllOf {
	this := BulkIndexTeamsRequestAllOf{}
	this.Teams = teams
	return &this
}

// NewBulkIndexTeamsRequestAllOfWithDefaults instantiates a new BulkIndexTeamsRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkIndexTeamsRequestAllOfWithDefaults() *BulkIndexTeamsRequestAllOf {
	this := BulkIndexTeamsRequestAllOf{}
	return &this
}

// GetTeams returns the Teams field value
func (o *BulkIndexTeamsRequestAllOf) GetTeams() []TeamInfoDefinition {
	if o == nil {
		var ret []TeamInfoDefinition
		return ret
	}

	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value
// and a boolean to check if the value has been set.
func (o *BulkIndexTeamsRequestAllOf) GetTeamsOk() ([]TeamInfoDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Teams, true
}

// SetTeams sets field value
func (o *BulkIndexTeamsRequestAllOf) SetTeams(v []TeamInfoDefinition) {
	o.Teams = v
}

func (o BulkIndexTeamsRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["teams"] = o.Teams
	}
	return json.Marshal(toSerialize)
}

type NullableBulkIndexTeamsRequestAllOf struct {
	value *BulkIndexTeamsRequestAllOf
	isSet bool
}

func (v NullableBulkIndexTeamsRequestAllOf) Get() *BulkIndexTeamsRequestAllOf {
	return v.value
}

func (v *NullableBulkIndexTeamsRequestAllOf) Set(val *BulkIndexTeamsRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkIndexTeamsRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkIndexTeamsRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkIndexTeamsRequestAllOf(val *BulkIndexTeamsRequestAllOf) *NullableBulkIndexTeamsRequestAllOf {
	return &NullableBulkIndexTeamsRequestAllOf{value: val, isSet: true}
}

func (v NullableBulkIndexTeamsRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkIndexTeamsRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


