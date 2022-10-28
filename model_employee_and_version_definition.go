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

// EmployeeAndVersionDefinition describes info about an employee and optional version for that info
type EmployeeAndVersionDefinition struct {
	Employee *EmployeeInfoDefinition `json:"employee,omitempty"`
	// Version number for the employee object. If absent or 0 then no version checks are done
	Version *int64 `json:"version,omitempty"`
}

// NewEmployeeAndVersionDefinition instantiates a new EmployeeAndVersionDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployeeAndVersionDefinition() *EmployeeAndVersionDefinition {
	this := EmployeeAndVersionDefinition{}
	return &this
}

// NewEmployeeAndVersionDefinitionWithDefaults instantiates a new EmployeeAndVersionDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeAndVersionDefinitionWithDefaults() *EmployeeAndVersionDefinition {
	this := EmployeeAndVersionDefinition{}
	return &this
}

// GetEmployee returns the Employee field value if set, zero value otherwise.
func (o *EmployeeAndVersionDefinition) GetEmployee() EmployeeInfoDefinition {
	if o == nil || o.Employee == nil {
		var ret EmployeeInfoDefinition
		return ret
	}
	return *o.Employee
}

// GetEmployeeOk returns a tuple with the Employee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeAndVersionDefinition) GetEmployeeOk() (*EmployeeInfoDefinition, bool) {
	if o == nil || o.Employee == nil {
		return nil, false
	}
	return o.Employee, true
}

// HasEmployee returns a boolean if a field has been set.
func (o *EmployeeAndVersionDefinition) HasEmployee() bool {
	if o != nil && o.Employee != nil {
		return true
	}

	return false
}

// SetEmployee gets a reference to the given EmployeeInfoDefinition and assigns it to the Employee field.
func (o *EmployeeAndVersionDefinition) SetEmployee(v EmployeeInfoDefinition) {
	o.Employee = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EmployeeAndVersionDefinition) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeAndVersionDefinition) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EmployeeAndVersionDefinition) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *EmployeeAndVersionDefinition) SetVersion(v int64) {
	o.Version = &v
}

func (o EmployeeAndVersionDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Employee != nil {
		toSerialize["employee"] = o.Employee
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableEmployeeAndVersionDefinition struct {
	value *EmployeeAndVersionDefinition
	isSet bool
}

func (v NullableEmployeeAndVersionDefinition) Get() *EmployeeAndVersionDefinition {
	return v.value
}

func (v *NullableEmployeeAndVersionDefinition) Set(val *EmployeeAndVersionDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployeeAndVersionDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployeeAndVersionDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployeeAndVersionDefinition(val *EmployeeAndVersionDefinition) *NullableEmployeeAndVersionDefinition {
	return &NullableEmployeeAndVersionDefinition{value: val, isSet: true}
}

func (v NullableEmployeeAndVersionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployeeAndVersionDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

