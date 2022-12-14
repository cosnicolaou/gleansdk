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

// DeleteDocumentRequest Describes the request body of the /deletedocument API call
type DeleteDocumentRequest struct {
	// Version number for document for optimistic concurrency control. If absent or 0 then no version checks are done.
	Version *int64 `json:"version,omitempty"`
	// datasource of the document
	Datasource string `json:"datasource"`
	// object type of the document
	ObjectType string `json:"objectType"`
	// The id of the document
	Id string `json:"id"`
}

// NewDeleteDocumentRequest instantiates a new DeleteDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDocumentRequest(datasource string, objectType string, id string) *DeleteDocumentRequest {
	this := DeleteDocumentRequest{}
	this.Datasource = datasource
	this.ObjectType = objectType
	this.Id = id
	return &this
}

// NewDeleteDocumentRequestWithDefaults instantiates a new DeleteDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDocumentRequestWithDefaults() *DeleteDocumentRequest {
	this := DeleteDocumentRequest{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DeleteDocumentRequest) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDocumentRequest) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DeleteDocumentRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *DeleteDocumentRequest) SetVersion(v int64) {
	o.Version = &v
}

// GetDatasource returns the Datasource field value
func (o *DeleteDocumentRequest) GetDatasource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datasource
}

// GetDatasourceOk returns a tuple with the Datasource field value
// and a boolean to check if the value has been set.
func (o *DeleteDocumentRequest) GetDatasourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datasource, true
}

// SetDatasource sets field value
func (o *DeleteDocumentRequest) SetDatasource(v string) {
	o.Datasource = v
}

// GetObjectType returns the ObjectType field value
func (o *DeleteDocumentRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *DeleteDocumentRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *DeleteDocumentRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetId returns the Id field value
func (o *DeleteDocumentRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeleteDocumentRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeleteDocumentRequest) SetId(v string) {
	o.Id = v
}

func (o DeleteDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["datasource"] = o.Datasource
	}
	if true {
		toSerialize["objectType"] = o.ObjectType
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteDocumentRequest struct {
	value *DeleteDocumentRequest
	isSet bool
}

func (v NullableDeleteDocumentRequest) Get() *DeleteDocumentRequest {
	return v.value
}

func (v *NullableDeleteDocumentRequest) Set(val *DeleteDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDocumentRequest(val *DeleteDocumentRequest) *NullableDeleteDocumentRequest {
	return &NullableDeleteDocumentRequest{value: val, isSet: true}
}

func (v NullableDeleteDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


