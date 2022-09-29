/*
 * Connector Management API
 *
 * Connector Management API is a REST API to manage connectors.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
)

// ConnectorNamespaceListAllOf struct for ConnectorNamespaceListAllOf
type ConnectorNamespaceListAllOf struct {
	Items *[]ConnectorNamespace `json:"items,omitempty"`
}

// NewConnectorNamespaceListAllOf instantiates a new ConnectorNamespaceListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorNamespaceListAllOf() *ConnectorNamespaceListAllOf {
	this := ConnectorNamespaceListAllOf{}
	return &this
}

// NewConnectorNamespaceListAllOfWithDefaults instantiates a new ConnectorNamespaceListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorNamespaceListAllOfWithDefaults() *ConnectorNamespaceListAllOf {
	this := ConnectorNamespaceListAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ConnectorNamespaceListAllOf) GetItems() []ConnectorNamespace {
	if o == nil || o.Items == nil {
		var ret []ConnectorNamespace
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorNamespaceListAllOf) GetItemsOk() (*[]ConnectorNamespace, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ConnectorNamespaceListAllOf) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ConnectorNamespace and assigns it to the Items field.
func (o *ConnectorNamespaceListAllOf) SetItems(v []ConnectorNamespace) {
	o.Items = &v
}

func (o ConnectorNamespaceListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorNamespaceListAllOf struct {
	value *ConnectorNamespaceListAllOf
	isSet bool
}

func (v NullableConnectorNamespaceListAllOf) Get() *ConnectorNamespaceListAllOf {
	return v.value
}

func (v *NullableConnectorNamespaceListAllOf) Set(val *ConnectorNamespaceListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorNamespaceListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorNamespaceListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorNamespaceListAllOf(val *ConnectorNamespaceListAllOf) *NullableConnectorNamespaceListAllOf {
	return &NullableConnectorNamespaceListAllOf{value: val, isSet: true}
}

func (v NullableConnectorNamespaceListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorNamespaceListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


