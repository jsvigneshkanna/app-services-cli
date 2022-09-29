/*
 * Apicurio Registry API [v2]
 *
 * Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 
 *
 * API version: 2.2.5.Final
 * Contact: apicurio@lists.jboss.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryinstanceclient

import (
	"encoding/json"
)

// UpdateConfigurationProperty struct for UpdateConfigurationProperty
type UpdateConfigurationProperty struct {
	Value string `json:"value"`
}

// NewUpdateConfigurationProperty instantiates a new UpdateConfigurationProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateConfigurationProperty(value string) *UpdateConfigurationProperty {
	this := UpdateConfigurationProperty{}
	this.Value = value
	return &this
}

// NewUpdateConfigurationPropertyWithDefaults instantiates a new UpdateConfigurationProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateConfigurationPropertyWithDefaults() *UpdateConfigurationProperty {
	this := UpdateConfigurationProperty{}
	return &this
}

// GetValue returns the Value field value
func (o *UpdateConfigurationProperty) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UpdateConfigurationProperty) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UpdateConfigurationProperty) SetValue(v string) {
	o.Value = v
}

func (o UpdateConfigurationProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateConfigurationProperty struct {
	value *UpdateConfigurationProperty
	isSet bool
}

func (v NullableUpdateConfigurationProperty) Get() *UpdateConfigurationProperty {
	return v.value
}

func (v *NullableUpdateConfigurationProperty) Set(val *UpdateConfigurationProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateConfigurationProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateConfigurationProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateConfigurationProperty(val *UpdateConfigurationProperty) *NullableUpdateConfigurationProperty {
	return &NullableUpdateConfigurationProperty{value: val, isSet: true}
}

func (v NullableUpdateConfigurationProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateConfigurationProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


