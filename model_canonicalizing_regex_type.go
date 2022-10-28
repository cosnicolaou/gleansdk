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

// CanonicalizingRegexType Regular expression to apply to an arbitrary string to transform it into a canonical string.
type CanonicalizingRegexType struct {
	// Regular expression to match to an arbitrary string.
	MatchRegex *string `json:"matchRegex,omitempty"`
	// Regular expression to transform into a canonical string.
	RewriteRegex *string `json:"rewriteRegex,omitempty"`
}

// NewCanonicalizingRegexType instantiates a new CanonicalizingRegexType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCanonicalizingRegexType() *CanonicalizingRegexType {
	this := CanonicalizingRegexType{}
	return &this
}

// NewCanonicalizingRegexTypeWithDefaults instantiates a new CanonicalizingRegexType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCanonicalizingRegexTypeWithDefaults() *CanonicalizingRegexType {
	this := CanonicalizingRegexType{}
	return &this
}

// GetMatchRegex returns the MatchRegex field value if set, zero value otherwise.
func (o *CanonicalizingRegexType) GetMatchRegex() string {
	if o == nil || o.MatchRegex == nil {
		var ret string
		return ret
	}
	return *o.MatchRegex
}

// GetMatchRegexOk returns a tuple with the MatchRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanonicalizingRegexType) GetMatchRegexOk() (*string, bool) {
	if o == nil || o.MatchRegex == nil {
		return nil, false
	}
	return o.MatchRegex, true
}

// HasMatchRegex returns a boolean if a field has been set.
func (o *CanonicalizingRegexType) HasMatchRegex() bool {
	if o != nil && o.MatchRegex != nil {
		return true
	}

	return false
}

// SetMatchRegex gets a reference to the given string and assigns it to the MatchRegex field.
func (o *CanonicalizingRegexType) SetMatchRegex(v string) {
	o.MatchRegex = &v
}

// GetRewriteRegex returns the RewriteRegex field value if set, zero value otherwise.
func (o *CanonicalizingRegexType) GetRewriteRegex() string {
	if o == nil || o.RewriteRegex == nil {
		var ret string
		return ret
	}
	return *o.RewriteRegex
}

// GetRewriteRegexOk returns a tuple with the RewriteRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanonicalizingRegexType) GetRewriteRegexOk() (*string, bool) {
	if o == nil || o.RewriteRegex == nil {
		return nil, false
	}
	return o.RewriteRegex, true
}

// HasRewriteRegex returns a boolean if a field has been set.
func (o *CanonicalizingRegexType) HasRewriteRegex() bool {
	if o != nil && o.RewriteRegex != nil {
		return true
	}

	return false
}

// SetRewriteRegex gets a reference to the given string and assigns it to the RewriteRegex field.
func (o *CanonicalizingRegexType) SetRewriteRegex(v string) {
	o.RewriteRegex = &v
}

func (o CanonicalizingRegexType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MatchRegex != nil {
		toSerialize["matchRegex"] = o.MatchRegex
	}
	if o.RewriteRegex != nil {
		toSerialize["rewriteRegex"] = o.RewriteRegex
	}
	return json.Marshal(toSerialize)
}

type NullableCanonicalizingRegexType struct {
	value *CanonicalizingRegexType
	isSet bool
}

func (v NullableCanonicalizingRegexType) Get() *CanonicalizingRegexType {
	return v.value
}

func (v *NullableCanonicalizingRegexType) Set(val *CanonicalizingRegexType) {
	v.value = val
	v.isSet = true
}

func (v NullableCanonicalizingRegexType) IsSet() bool {
	return v.isSet
}

func (v *NullableCanonicalizingRegexType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCanonicalizingRegexType(val *CanonicalizingRegexType) *NullableCanonicalizingRegexType {
	return &NullableCanonicalizingRegexType{value: val, isSet: true}
}

func (v NullableCanonicalizingRegexType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCanonicalizingRegexType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

