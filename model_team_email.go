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

// TeamEmail Information about a team's email
type TeamEmail struct {
	// An email address
	Email string `json:"email"`
	// An enum of `PRIMARY`, `SECONDARY`, `ONCALL`, `OTHER`
	Type string `json:"type"`
}

// NewTeamEmail instantiates a new TeamEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamEmail(email string, type_ string) *TeamEmail {
	this := TeamEmail{}
	this.Email = email
	this.Type = type_
	return &this
}

// NewTeamEmailWithDefaults instantiates a new TeamEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamEmailWithDefaults() *TeamEmail {
	this := TeamEmail{}
	var type_ string = "OTHER"
	this.Type = type_
	return &this
}

// GetEmail returns the Email field value
func (o *TeamEmail) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *TeamEmail) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *TeamEmail) SetEmail(v string) {
	o.Email = v
}

// GetType returns the Type field value
func (o *TeamEmail) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamEmail) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TeamEmail) SetType(v string) {
	o.Type = v
}

func (o TeamEmail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTeamEmail struct {
	value *TeamEmail
	isSet bool
}

func (v NullableTeamEmail) Get() *TeamEmail {
	return v.value
}

func (v *NullableTeamEmail) Set(val *TeamEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamEmail(val *TeamEmail) *NullableTeamEmail {
	return &NullableTeamEmail{value: val, isSet: true}
}

func (v NullableTeamEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


