// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/seamapi/go/core"
	time "time"
)

type DevicesDeleteRequest struct {
	DeviceId string `json:"device_id" url:"device_id"`
}

type DevicesGetRequest struct {
	DeviceId *string `json:"device_id,omitempty" url:"device_id,omitempty"`
	Name     *string `json:"name,omitempty" url:"name,omitempty"`
}

type DevicesListRequest struct {
	// List all devices owned by this connected account
	ConnectedAccountId  *string                                              `json:"connected_account_id,omitempty" url:"connected_account_id,omitempty"`
	ConnectedAccountIds []string                                             `json:"connected_account_ids,omitempty" url:"connected_account_ids,omitempty"`
	ConnectWebviewId    *string                                              `json:"connect_webview_id,omitempty" url:"connect_webview_id,omitempty"`
	DeviceType          *DeviceType                                          `json:"device_type,omitempty" url:"device_type,omitempty"`
	DeviceTypes         []DeviceType                                         `json:"device_types,omitempty" url:"device_types,omitempty"`
	Manufacturer        *Manufacturer                                        `json:"manufacturer,omitempty" url:"manufacturer,omitempty"`
	DeviceIds           []string                                             `json:"device_ids,omitempty" url:"device_ids,omitempty"`
	Limit               *float64                                             `json:"limit,omitempty" url:"limit,omitempty"`
	CreatedBefore       *time.Time                                           `json:"created_before,omitempty" url:"created_before,omitempty"`
	UserIdentifierKey   *string                                              `json:"user_identifier_key,omitempty" url:"user_identifier_key,omitempty"`
	CustomMetadataHas   map[string]*DevicesListRequestCustomMetadataHasValue `json:"custom_metadata_has,omitempty" url:"custom_metadata_has,omitempty"`
	IncludeIf           []DevicesListRequestIncludeIfItem                    `json:"include_if,omitempty" url:"include_if,omitempty"`
	ExcludeIf           []DevicesListRequestExcludeIfItem                    `json:"exclude_if,omitempty" url:"exclude_if,omitempty"`
}

func (d *DevicesListRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler DevicesListRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*d = DevicesListRequest(body)
	return nil
}

func (d *DevicesListRequest) MarshalJSON() ([]byte, error) {
	type embed DevicesListRequest
	var marshaler = struct {
		embed
		CreatedBefore *core.DateTime `json:"created_before,omitempty"`
	}{
		embed:         embed(*d),
		CreatedBefore: core.NewOptionalDateTime(d.CreatedBefore),
	}
	return json.Marshal(marshaler)
}

type DevicesListDeviceProvidersRequest struct {
	ProviderCategory *ProviderCategory `json:"provider_category,omitempty" url:"provider_category,omitempty"`
}

type DevicesDeleteResponse struct {
	Ok bool `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (d *DevicesDeleteResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DevicesDeleteResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DevicesDeleteResponse(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DevicesDeleteResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DevicesGetResponse struct {
	Device *Device `json:"device,omitempty" url:"device,omitempty"`
	Ok     bool    `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (d *DevicesGetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DevicesGetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DevicesGetResponse(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DevicesGetResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DevicesListDeviceProvidersResponse struct {
	DeviceProviders []*DeviceProvider `json:"device_providers,omitempty" url:"device_providers,omitempty"`
	Ok              bool              `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (d *DevicesListDeviceProvidersResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DevicesListDeviceProvidersResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DevicesListDeviceProvidersResponse(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DevicesListDeviceProvidersResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DevicesListRequestCustomMetadataHasValue struct {
	typeName string
	String   string
	Boolean  bool
}

func NewDevicesListRequestCustomMetadataHasValueFromString(value string) *DevicesListRequestCustomMetadataHasValue {
	return &DevicesListRequestCustomMetadataHasValue{typeName: "string", String: value}
}

func NewDevicesListRequestCustomMetadataHasValueFromBoolean(value bool) *DevicesListRequestCustomMetadataHasValue {
	return &DevicesListRequestCustomMetadataHasValue{typeName: "boolean", Boolean: value}
}

func (d *DevicesListRequestCustomMetadataHasValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		d.typeName = "string"
		d.String = valueString
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		d.typeName = "boolean"
		d.Boolean = valueBoolean
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, d)
}

func (d DevicesListRequestCustomMetadataHasValue) MarshalJSON() ([]byte, error) {
	switch d.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", d.typeName, d)
	case "string":
		return json.Marshal(d.String)
	case "boolean":
		return json.Marshal(d.Boolean)
	}
}

type DevicesListRequestCustomMetadataHasValueVisitor interface {
	VisitString(string) error
	VisitBoolean(bool) error
}

func (d *DevicesListRequestCustomMetadataHasValue) Accept(visitor DevicesListRequestCustomMetadataHasValueVisitor) error {
	switch d.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", d.typeName, d)
	case "string":
		return visitor.VisitString(d.String)
	case "boolean":
		return visitor.VisitBoolean(d.Boolean)
	}
}

type DevicesListRequestExcludeIfItem string

const (
	DevicesListRequestExcludeIfItemCanRemotelyUnlock            DevicesListRequestExcludeIfItem = "can_remotely_unlock"
	DevicesListRequestExcludeIfItemCanRemotelyLock              DevicesListRequestExcludeIfItem = "can_remotely_lock"
	DevicesListRequestExcludeIfItemCanProgramOfflineAccessCodes DevicesListRequestExcludeIfItem = "can_program_offline_access_codes"
	DevicesListRequestExcludeIfItemCanProgramOnlineAccessCodes  DevicesListRequestExcludeIfItem = "can_program_online_access_codes"
	DevicesListRequestExcludeIfItemCanSimulateRemoval           DevicesListRequestExcludeIfItem = "can_simulate_removal"
	DevicesListRequestExcludeIfItemCanSimulateConnection        DevicesListRequestExcludeIfItem = "can_simulate_connection"
	DevicesListRequestExcludeIfItemCanSimulateDisconnection     DevicesListRequestExcludeIfItem = "can_simulate_disconnection"
)

func NewDevicesListRequestExcludeIfItemFromString(s string) (DevicesListRequestExcludeIfItem, error) {
	switch s {
	case "can_remotely_unlock":
		return DevicesListRequestExcludeIfItemCanRemotelyUnlock, nil
	case "can_remotely_lock":
		return DevicesListRequestExcludeIfItemCanRemotelyLock, nil
	case "can_program_offline_access_codes":
		return DevicesListRequestExcludeIfItemCanProgramOfflineAccessCodes, nil
	case "can_program_online_access_codes":
		return DevicesListRequestExcludeIfItemCanProgramOnlineAccessCodes, nil
	case "can_simulate_removal":
		return DevicesListRequestExcludeIfItemCanSimulateRemoval, nil
	case "can_simulate_connection":
		return DevicesListRequestExcludeIfItemCanSimulateConnection, nil
	case "can_simulate_disconnection":
		return DevicesListRequestExcludeIfItemCanSimulateDisconnection, nil
	}
	var t DevicesListRequestExcludeIfItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DevicesListRequestExcludeIfItem) Ptr() *DevicesListRequestExcludeIfItem {
	return &d
}

type DevicesListRequestIncludeIfItem string

const (
	DevicesListRequestIncludeIfItemCanRemotelyUnlock            DevicesListRequestIncludeIfItem = "can_remotely_unlock"
	DevicesListRequestIncludeIfItemCanRemotelyLock              DevicesListRequestIncludeIfItem = "can_remotely_lock"
	DevicesListRequestIncludeIfItemCanProgramOfflineAccessCodes DevicesListRequestIncludeIfItem = "can_program_offline_access_codes"
	DevicesListRequestIncludeIfItemCanProgramOnlineAccessCodes  DevicesListRequestIncludeIfItem = "can_program_online_access_codes"
	DevicesListRequestIncludeIfItemCanSimulateRemoval           DevicesListRequestIncludeIfItem = "can_simulate_removal"
	DevicesListRequestIncludeIfItemCanSimulateConnection        DevicesListRequestIncludeIfItem = "can_simulate_connection"
	DevicesListRequestIncludeIfItemCanSimulateDisconnection     DevicesListRequestIncludeIfItem = "can_simulate_disconnection"
)

func NewDevicesListRequestIncludeIfItemFromString(s string) (DevicesListRequestIncludeIfItem, error) {
	switch s {
	case "can_remotely_unlock":
		return DevicesListRequestIncludeIfItemCanRemotelyUnlock, nil
	case "can_remotely_lock":
		return DevicesListRequestIncludeIfItemCanRemotelyLock, nil
	case "can_program_offline_access_codes":
		return DevicesListRequestIncludeIfItemCanProgramOfflineAccessCodes, nil
	case "can_program_online_access_codes":
		return DevicesListRequestIncludeIfItemCanProgramOnlineAccessCodes, nil
	case "can_simulate_removal":
		return DevicesListRequestIncludeIfItemCanSimulateRemoval, nil
	case "can_simulate_connection":
		return DevicesListRequestIncludeIfItemCanSimulateConnection, nil
	case "can_simulate_disconnection":
		return DevicesListRequestIncludeIfItemCanSimulateDisconnection, nil
	}
	var t DevicesListRequestIncludeIfItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DevicesListRequestIncludeIfItem) Ptr() *DevicesListRequestIncludeIfItem {
	return &d
}

type DevicesListResponse struct {
	Devices []*Device `json:"devices,omitempty" url:"devices,omitempty"`
	Ok      bool      `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (d *DevicesListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DevicesListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DevicesListResponse(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DevicesListResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DevicesUpdateRequestCustomMetadataValue struct {
	typeName       string
	String         string
	Boolean        bool
	StringOptional *string
}

func NewDevicesUpdateRequestCustomMetadataValueFromString(value string) *DevicesUpdateRequestCustomMetadataValue {
	return &DevicesUpdateRequestCustomMetadataValue{typeName: "string", String: value}
}

func NewDevicesUpdateRequestCustomMetadataValueFromBoolean(value bool) *DevicesUpdateRequestCustomMetadataValue {
	return &DevicesUpdateRequestCustomMetadataValue{typeName: "boolean", Boolean: value}
}

func NewDevicesUpdateRequestCustomMetadataValueFromStringOptional(value *string) *DevicesUpdateRequestCustomMetadataValue {
	return &DevicesUpdateRequestCustomMetadataValue{typeName: "stringOptional", StringOptional: value}
}

func (d *DevicesUpdateRequestCustomMetadataValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		d.typeName = "string"
		d.String = valueString
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		d.typeName = "boolean"
		d.Boolean = valueBoolean
		return nil
	}
	var valueStringOptional *string
	if err := json.Unmarshal(data, &valueStringOptional); err == nil {
		d.typeName = "stringOptional"
		d.StringOptional = valueStringOptional
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, d)
}

func (d DevicesUpdateRequestCustomMetadataValue) MarshalJSON() ([]byte, error) {
	switch d.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", d.typeName, d)
	case "string":
		return json.Marshal(d.String)
	case "boolean":
		return json.Marshal(d.Boolean)
	case "stringOptional":
		return json.Marshal(d.StringOptional)
	}
}

type DevicesUpdateRequestCustomMetadataValueVisitor interface {
	VisitString(string) error
	VisitBoolean(bool) error
	VisitStringOptional(*string) error
}

func (d *DevicesUpdateRequestCustomMetadataValue) Accept(visitor DevicesUpdateRequestCustomMetadataValueVisitor) error {
	switch d.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", d.typeName, d)
	case "string":
		return visitor.VisitString(d.String)
	case "boolean":
		return visitor.VisitBoolean(d.Boolean)
	case "stringOptional":
		return visitor.VisitStringOptional(d.StringOptional)
	}
}

type DevicesUpdateRequestProperties struct {
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	_rawJSON json.RawMessage
}

func (d *DevicesUpdateRequestProperties) UnmarshalJSON(data []byte) error {
	type unmarshaler DevicesUpdateRequestProperties
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DevicesUpdateRequestProperties(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DevicesUpdateRequestProperties) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DevicesUpdateResponse struct {
	Ok bool `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (d *DevicesUpdateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DevicesUpdateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DevicesUpdateResponse(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DevicesUpdateResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DevicesUpdateRequest struct {
	DeviceId       string                                              `json:"device_id" url:"device_id"`
	Properties     *DevicesUpdateRequestProperties                     `json:"properties,omitempty" url:"properties,omitempty"`
	Name           *string                                             `json:"name,omitempty" url:"name,omitempty"`
	IsManaged      *bool                                               `json:"is_managed,omitempty" url:"is_managed,omitempty"`
	CustomMetadata map[string]*DevicesUpdateRequestCustomMetadataValue `json:"custom_metadata,omitempty" url:"custom_metadata,omitempty"`
}
