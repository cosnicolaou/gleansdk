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

// GreenlistUsersRequest Describes the request body of the /betausers API call
type GreenlistUsersRequest struct {
	// Datasource which needs to be made visible to users specified in the `emails` field.
	Datasource string `json:"datasource"`
	// The emails of the beta users
	Emails []string `json:"emails"`
}

// NewGreenlistUsersRequest instantiates a new GreenlistUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGreenlistUsersRequest(datasource string, emails []string) *GreenlistUsersRequest {
	this := GreenlistUsersRequest{}
	this.Datasource = datasource
	this.Emails = emails
	return &this
}

// NewGreenlistUsersRequestWithDefaults instantiates a new GreenlistUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGreenlistUsersRequestWithDefaults() *GreenlistUsersRequest {
	this := GreenlistUsersRequest{}
	return &this
}

// GetDatasource returns the Datasource field value
func (o *GreenlistUsersRequest) GetDatasource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datasource
}

// GetDatasourceOk returns a tuple with the Datasource field value
// and a boolean to check if the value has been set.
func (o *GreenlistUsersRequest) GetDatasourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datasource, true
}

// SetDatasource sets field value
func (o *GreenlistUsersRequest) SetDatasource(v string) {
	o.Datasource = v
}

// GetEmails returns the Emails field value
func (o *GreenlistUsersRequest) GetEmails() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *GreenlistUsersRequest) GetEmailsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails, true
}

// SetEmails sets field value
func (o *GreenlistUsersRequest) SetEmails(v []string) {
	o.Emails = v
}

func (o GreenlistUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["datasource"] = o.Datasource
	}
	if true {
		toSerialize["emails"] = o.Emails
	}
	return json.Marshal(toSerialize)
}

type NullableGreenlistUsersRequest struct {
	value *GreenlistUsersRequest
	isSet bool
}

func (v NullableGreenlistUsersRequest) Get() *GreenlistUsersRequest {
	return v.value
}

func (v *NullableGreenlistUsersRequest) Set(val *GreenlistUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGreenlistUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGreenlistUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGreenlistUsersRequest(val *GreenlistUsersRequest) *NullableGreenlistUsersRequest {
	return &NullableGreenlistUsersRequest{value: val, isSet: true}
}

func (v NullableGreenlistUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGreenlistUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


