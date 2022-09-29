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
	"fmt"
)

// ConnectorNamespaceTenantKind the model 'ConnectorNamespaceTenantKind'
type ConnectorNamespaceTenantKind string

// List of ConnectorNamespaceTenantKind
const (
	CONNECTORNAMESPACETENANTKIND_USER ConnectorNamespaceTenantKind = "user"
	CONNECTORNAMESPACETENANTKIND_ORGANISATION ConnectorNamespaceTenantKind = "organisation"
)

var allowedConnectorNamespaceTenantKindEnumValues = []ConnectorNamespaceTenantKind{
	"user",
	"organisation",
}

func (v *ConnectorNamespaceTenantKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectorNamespaceTenantKind(value)
	for _, existing := range allowedConnectorNamespaceTenantKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectorNamespaceTenantKind", value)
}

// NewConnectorNamespaceTenantKindFromValue returns a pointer to a valid ConnectorNamespaceTenantKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectorNamespaceTenantKindFromValue(v string) (*ConnectorNamespaceTenantKind, error) {
	ev := ConnectorNamespaceTenantKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectorNamespaceTenantKind: valid values are %v", v, allowedConnectorNamespaceTenantKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectorNamespaceTenantKind) IsValid() bool {
	for _, existing := range allowedConnectorNamespaceTenantKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectorNamespaceTenantKind value
func (v ConnectorNamespaceTenantKind) Ptr() *ConnectorNamespaceTenantKind {
	return &v
}

type NullableConnectorNamespaceTenantKind struct {
	value *ConnectorNamespaceTenantKind
	isSet bool
}

func (v NullableConnectorNamespaceTenantKind) Get() *ConnectorNamespaceTenantKind {
	return v.value
}

func (v *NullableConnectorNamespaceTenantKind) Set(val *ConnectorNamespaceTenantKind) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorNamespaceTenantKind) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorNamespaceTenantKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorNamespaceTenantKind(val *ConnectorNamespaceTenantKind) *NullableConnectorNamespaceTenantKind {
	return &NullableConnectorNamespaceTenantKind{value: val, isSet: true}
}

func (v NullableConnectorNamespaceTenantKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorNamespaceTenantKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

