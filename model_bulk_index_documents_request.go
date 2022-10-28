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

// BulkIndexDocumentsRequest Describes the request body of the /bulkindexdocuments API call
type BulkIndexDocumentsRequest struct {
	// Unique id that must be used for this instance of datasource employees upload
	UploadId string `json:"uploadId"`
	// true if this is the first page of the upload. Defaults to false
	IsFirstPage *bool `json:"isFirstPage,omitempty"`
	// true if this is the last page of the upload. Defaults to false
	IsLastPage *bool `json:"isLastPage,omitempty"`
	// Flag to discard previous upload attempts and start from scratch. Must be specified with isFirstPage=true
	ForceRestartUpload *bool `json:"forceRestartUpload,omitempty"`
	// Datasource of the documents
	Datasource string `json:"datasource"`
	// Batch of documents for the datasource
	Documents []DocumentDefinition `json:"documents"`
	// True if older documents need to be force deleted after the upload completes. Defaults to older documents being deleted asynchronously. This must only be set when `isLastPage = true`
	DisableStaleDocumentDeletionCheck *bool `json:"disableStaleDocumentDeletionCheck,omitempty"`
}

// NewBulkIndexDocumentsRequest instantiates a new BulkIndexDocumentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkIndexDocumentsRequest(uploadId string, datasource string, documents []DocumentDefinition) *BulkIndexDocumentsRequest {
	this := BulkIndexDocumentsRequest{}
	this.UploadId = uploadId
	this.Datasource = datasource
	this.Documents = documents
	return &this
}

// NewBulkIndexDocumentsRequestWithDefaults instantiates a new BulkIndexDocumentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkIndexDocumentsRequestWithDefaults() *BulkIndexDocumentsRequest {
	this := BulkIndexDocumentsRequest{}
	return &this
}

// GetUploadId returns the UploadId field value
func (o *BulkIndexDocumentsRequest) GetUploadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value
// and a boolean to check if the value has been set.
func (o *BulkIndexDocumentsRequest) GetUploadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadId, true
}

// SetUploadId sets field value
func (o *BulkIndexDocumentsRequest) SetUploadId(v string) {
	o.UploadId = v
}

// GetIsFirstPage returns the IsFirstPage field value if set, zero value otherwise.
func (o *BulkIndexDocumentsRequest) GetIsFirstPage() bool {
	if o == nil || o.IsFirstPage == nil {
		var ret bool
		return ret
	}
	return *o.IsFirstPage
}

// GetIsFirstPageOk returns a tuple with the IsFirstPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkIndexDocumentsRequest) GetIsFirstPageOk() (*bool, bool) {
	if o == nil || o.IsFirstPage == nil {
		return nil, false
	}
	return o.IsFirstPage, true
}

// HasIsFirstPage returns a boolean if a field has been set.
func (o *BulkIndexDocumentsRequest) HasIsFirstPage() bool {
	if o != nil && o.IsFirstPage != nil {
		return true
	}

	return false
}

// SetIsFirstPage gets a reference to the given bool and assigns it to the IsFirstPage field.
func (o *BulkIndexDocumentsRequest) SetIsFirstPage(v bool) {
	o.IsFirstPage = &v
}

// GetIsLastPage returns the IsLastPage field value if set, zero value otherwise.
func (o *BulkIndexDocumentsRequest) GetIsLastPage() bool {
	if o == nil || o.IsLastPage == nil {
		var ret bool
		return ret
	}
	return *o.IsLastPage
}

// GetIsLastPageOk returns a tuple with the IsLastPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkIndexDocumentsRequest) GetIsLastPageOk() (*bool, bool) {
	if o == nil || o.IsLastPage == nil {
		return nil, false
	}
	return o.IsLastPage, true
}

// HasIsLastPage returns a boolean if a field has been set.
func (o *BulkIndexDocumentsRequest) HasIsLastPage() bool {
	if o != nil && o.IsLastPage != nil {
		return true
	}

	return false
}

// SetIsLastPage gets a reference to the given bool and assigns it to the IsLastPage field.
func (o *BulkIndexDocumentsRequest) SetIsLastPage(v bool) {
	o.IsLastPage = &v
}

// GetForceRestartUpload returns the ForceRestartUpload field value if set, zero value otherwise.
func (o *BulkIndexDocumentsRequest) GetForceRestartUpload() bool {
	if o == nil || o.ForceRestartUpload == nil {
		var ret bool
		return ret
	}
	return *o.ForceRestartUpload
}

// GetForceRestartUploadOk returns a tuple with the ForceRestartUpload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkIndexDocumentsRequest) GetForceRestartUploadOk() (*bool, bool) {
	if o == nil || o.ForceRestartUpload == nil {
		return nil, false
	}
	return o.ForceRestartUpload, true
}

// HasForceRestartUpload returns a boolean if a field has been set.
func (o *BulkIndexDocumentsRequest) HasForceRestartUpload() bool {
	if o != nil && o.ForceRestartUpload != nil {
		return true
	}

	return false
}

// SetForceRestartUpload gets a reference to the given bool and assigns it to the ForceRestartUpload field.
func (o *BulkIndexDocumentsRequest) SetForceRestartUpload(v bool) {
	o.ForceRestartUpload = &v
}

// GetDatasource returns the Datasource field value
func (o *BulkIndexDocumentsRequest) GetDatasource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datasource
}

// GetDatasourceOk returns a tuple with the Datasource field value
// and a boolean to check if the value has been set.
func (o *BulkIndexDocumentsRequest) GetDatasourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datasource, true
}

// SetDatasource sets field value
func (o *BulkIndexDocumentsRequest) SetDatasource(v string) {
	o.Datasource = v
}

// GetDocuments returns the Documents field value
func (o *BulkIndexDocumentsRequest) GetDocuments() []DocumentDefinition {
	if o == nil {
		var ret []DocumentDefinition
		return ret
	}

	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value
// and a boolean to check if the value has been set.
func (o *BulkIndexDocumentsRequest) GetDocumentsOk() ([]DocumentDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Documents, true
}

// SetDocuments sets field value
func (o *BulkIndexDocumentsRequest) SetDocuments(v []DocumentDefinition) {
	o.Documents = v
}

// GetDisableStaleDocumentDeletionCheck returns the DisableStaleDocumentDeletionCheck field value if set, zero value otherwise.
func (o *BulkIndexDocumentsRequest) GetDisableStaleDocumentDeletionCheck() bool {
	if o == nil || o.DisableStaleDocumentDeletionCheck == nil {
		var ret bool
		return ret
	}
	return *o.DisableStaleDocumentDeletionCheck
}

// GetDisableStaleDocumentDeletionCheckOk returns a tuple with the DisableStaleDocumentDeletionCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkIndexDocumentsRequest) GetDisableStaleDocumentDeletionCheckOk() (*bool, bool) {
	if o == nil || o.DisableStaleDocumentDeletionCheck == nil {
		return nil, false
	}
	return o.DisableStaleDocumentDeletionCheck, true
}

// HasDisableStaleDocumentDeletionCheck returns a boolean if a field has been set.
func (o *BulkIndexDocumentsRequest) HasDisableStaleDocumentDeletionCheck() bool {
	if o != nil && o.DisableStaleDocumentDeletionCheck != nil {
		return true
	}

	return false
}

// SetDisableStaleDocumentDeletionCheck gets a reference to the given bool and assigns it to the DisableStaleDocumentDeletionCheck field.
func (o *BulkIndexDocumentsRequest) SetDisableStaleDocumentDeletionCheck(v bool) {
	o.DisableStaleDocumentDeletionCheck = &v
}

func (o BulkIndexDocumentsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uploadId"] = o.UploadId
	}
	if o.IsFirstPage != nil {
		toSerialize["isFirstPage"] = o.IsFirstPage
	}
	if o.IsLastPage != nil {
		toSerialize["isLastPage"] = o.IsLastPage
	}
	if o.ForceRestartUpload != nil {
		toSerialize["forceRestartUpload"] = o.ForceRestartUpload
	}
	if true {
		toSerialize["datasource"] = o.Datasource
	}
	if true {
		toSerialize["documents"] = o.Documents
	}
	if o.DisableStaleDocumentDeletionCheck != nil {
		toSerialize["disableStaleDocumentDeletionCheck"] = o.DisableStaleDocumentDeletionCheck
	}
	return json.Marshal(toSerialize)
}

type NullableBulkIndexDocumentsRequest struct {
	value *BulkIndexDocumentsRequest
	isSet bool
}

func (v NullableBulkIndexDocumentsRequest) Get() *BulkIndexDocumentsRequest {
	return v.value
}

func (v *NullableBulkIndexDocumentsRequest) Set(val *BulkIndexDocumentsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkIndexDocumentsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkIndexDocumentsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkIndexDocumentsRequest(val *BulkIndexDocumentsRequest) *NullableBulkIndexDocumentsRequest {
	return &NullableBulkIndexDocumentsRequest{value: val, isSet: true}
}

func (v NullableBulkIndexDocumentsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkIndexDocumentsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

