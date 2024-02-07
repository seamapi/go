// This file was auto-generated by Fern from our API Definition.

package acs

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/seamapi/go/core"
	time "time"
)

type CredentialsAssignRequest struct {
	AcsUserId       string `json:"acs_user_id"`
	AcsCredentialId string `json:"acs_credential_id"`
}

type CredentialsCreateRequest struct {
	AcsUserId                  string                                      `json:"acs_user_id"`
	AccessMethod               CredentialsCreateRequestAccessMethod        `json:"access_method,omitempty"`
	Code                       *string                                     `json:"code,omitempty"`
	IsMultiPhoneSyncCredential *bool                                       `json:"is_multi_phone_sync_credential,omitempty"`
	ExternalType               *CredentialsCreateRequestExternalType       `json:"external_type,omitempty"`
	VisionlineMetadata         *CredentialsCreateRequestVisionlineMetadata `json:"visionline_metadata,omitempty"`
	StartsAt                   *string                                     `json:"starts_at,omitempty"`
	EndsAt                     *string                                     `json:"ends_at,omitempty"`
}

type CredentialsDeleteRequest struct {
	AcsCredentialId string `json:"acs_credential_id"`
}

type CredentialsGetRequest struct {
	AcsCredentialId string `json:"acs_credential_id"`
}

type CredentialsAssignResponse struct {
	AcsCredential *CredentialsAssignResponseAcsCredential `json:"acs_credential,omitempty"`
	Ok            bool                                    `json:"ok"`

	_rawJSON json.RawMessage
}

func (c *CredentialsAssignResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CredentialsAssignResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CredentialsAssignResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CredentialsAssignResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CredentialsCreateRequestAccessMethod string

const (
	CredentialsCreateRequestAccessMethodCode      CredentialsCreateRequestAccessMethod = "code"
	CredentialsCreateRequestAccessMethodCard      CredentialsCreateRequestAccessMethod = "card"
	CredentialsCreateRequestAccessMethodMobileKey CredentialsCreateRequestAccessMethod = "mobile_key"
)

func NewCredentialsCreateRequestAccessMethodFromString(s string) (CredentialsCreateRequestAccessMethod, error) {
	switch s {
	case "code":
		return CredentialsCreateRequestAccessMethodCode, nil
	case "card":
		return CredentialsCreateRequestAccessMethodCard, nil
	case "mobile_key":
		return CredentialsCreateRequestAccessMethodMobileKey, nil
	}
	var t CredentialsCreateRequestAccessMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CredentialsCreateRequestAccessMethod) Ptr() *CredentialsCreateRequestAccessMethod {
	return &c
}

type CredentialsCreateRequestExternalType string

const (
	CredentialsCreateRequestExternalTypePtiCard         CredentialsCreateRequestExternalType = "pti_card"
	CredentialsCreateRequestExternalTypeBrivoCredential CredentialsCreateRequestExternalType = "brivo_credential"
	CredentialsCreateRequestExternalTypeHidCredential   CredentialsCreateRequestExternalType = "hid_credential"
	CredentialsCreateRequestExternalTypeVisionlineCard  CredentialsCreateRequestExternalType = "visionline_card"
)

func NewCredentialsCreateRequestExternalTypeFromString(s string) (CredentialsCreateRequestExternalType, error) {
	switch s {
	case "pti_card":
		return CredentialsCreateRequestExternalTypePtiCard, nil
	case "brivo_credential":
		return CredentialsCreateRequestExternalTypeBrivoCredential, nil
	case "hid_credential":
		return CredentialsCreateRequestExternalTypeHidCredential, nil
	case "visionline_card":
		return CredentialsCreateRequestExternalTypeVisionlineCard, nil
	}
	var t CredentialsCreateRequestExternalType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CredentialsCreateRequestExternalType) Ptr() *CredentialsCreateRequestExternalType {
	return &c
}

type CredentialsCreateRequestVisionlineMetadata struct {
	AssaAbloyCredentialServiceMobileEndpointId *string                                               `json:"assa_abloy_credential_service_mobile_endpoint_id,omitempty"`
	CardFormat                                 *CredentialsCreateRequestVisionlineMetadataCardFormat `json:"card_format,omitempty"`
	IsOverrideKey                              *bool                                                 `json:"is_override_key,omitempty"`
	JoinerAcsCredentialIds                     []string                                              `json:"joiner_acs_credential_ids,omitempty"`

	_rawJSON json.RawMessage
}

func (c *CredentialsCreateRequestVisionlineMetadata) UnmarshalJSON(data []byte) error {
	type unmarshaler CredentialsCreateRequestVisionlineMetadata
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CredentialsCreateRequestVisionlineMetadata(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CredentialsCreateRequestVisionlineMetadata) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CredentialsCreateResponse struct {
	AcsCredential *CredentialsCreateResponseAcsCredential `json:"acs_credential,omitempty"`
	Ok            bool                                    `json:"ok"`

	_rawJSON json.RawMessage
}

func (c *CredentialsCreateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CredentialsCreateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CredentialsCreateResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CredentialsCreateResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CredentialsDeleteResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (c *CredentialsDeleteResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CredentialsDeleteResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CredentialsDeleteResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CredentialsDeleteResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CredentialsGetResponse struct {
	AcsCredential *CredentialsGetResponseAcsCredential `json:"acs_credential,omitempty"`
	Ok            bool                                 `json:"ok"`

	_rawJSON json.RawMessage
}

func (c *CredentialsGetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CredentialsGetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CredentialsGetResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CredentialsGetResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CredentialsListRequest struct {
	typeName                             string
	CredentialsListRequestZero           *CredentialsListRequestZero
	CredentialsListRequestOne            *CredentialsListRequestOne
	CredentialsListRequestTwo            *CredentialsListRequestTwo
	CredentialsListRequestUserIdentityId *CredentialsListRequestUserIdentityId
}

func NewCredentialsListRequestFromCredentialsListRequestZero(value *CredentialsListRequestZero) *CredentialsListRequest {
	return &CredentialsListRequest{typeName: "credentialsListRequestZero", CredentialsListRequestZero: value}
}

func NewCredentialsListRequestFromCredentialsListRequestOne(value *CredentialsListRequestOne) *CredentialsListRequest {
	return &CredentialsListRequest{typeName: "credentialsListRequestOne", CredentialsListRequestOne: value}
}

func NewCredentialsListRequestFromCredentialsListRequestTwo(value *CredentialsListRequestTwo) *CredentialsListRequest {
	return &CredentialsListRequest{typeName: "credentialsListRequestTwo", CredentialsListRequestTwo: value}
}

func NewCredentialsListRequestFromCredentialsListRequestUserIdentityId(value *CredentialsListRequestUserIdentityId) *CredentialsListRequest {
	return &CredentialsListRequest{typeName: "credentialsListRequestUserIdentityId", CredentialsListRequestUserIdentityId: value}
}

func (c *CredentialsListRequest) UnmarshalJSON(data []byte) error {
	valueCredentialsListRequestZero := new(CredentialsListRequestZero)
	if err := json.Unmarshal(data, &valueCredentialsListRequestZero); err == nil {
		c.typeName = "credentialsListRequestZero"
		c.CredentialsListRequestZero = valueCredentialsListRequestZero
		return nil
	}
	valueCredentialsListRequestOne := new(CredentialsListRequestOne)
	if err := json.Unmarshal(data, &valueCredentialsListRequestOne); err == nil {
		c.typeName = "credentialsListRequestOne"
		c.CredentialsListRequestOne = valueCredentialsListRequestOne
		return nil
	}
	valueCredentialsListRequestTwo := new(CredentialsListRequestTwo)
	if err := json.Unmarshal(data, &valueCredentialsListRequestTwo); err == nil {
		c.typeName = "credentialsListRequestTwo"
		c.CredentialsListRequestTwo = valueCredentialsListRequestTwo
		return nil
	}
	valueCredentialsListRequestUserIdentityId := new(CredentialsListRequestUserIdentityId)
	if err := json.Unmarshal(data, &valueCredentialsListRequestUserIdentityId); err == nil {
		c.typeName = "credentialsListRequestUserIdentityId"
		c.CredentialsListRequestUserIdentityId = valueCredentialsListRequestUserIdentityId
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CredentialsListRequest) MarshalJSON() ([]byte, error) {
	switch c.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "credentialsListRequestZero":
		return json.Marshal(c.CredentialsListRequestZero)
	case "credentialsListRequestOne":
		return json.Marshal(c.CredentialsListRequestOne)
	case "credentialsListRequestTwo":
		return json.Marshal(c.CredentialsListRequestTwo)
	case "credentialsListRequestUserIdentityId":
		return json.Marshal(c.CredentialsListRequestUserIdentityId)
	}
}

type CredentialsListRequestVisitor interface {
	VisitCredentialsListRequestZero(*CredentialsListRequestZero) error
	VisitCredentialsListRequestOne(*CredentialsListRequestOne) error
	VisitCredentialsListRequestTwo(*CredentialsListRequestTwo) error
	VisitCredentialsListRequestUserIdentityId(*CredentialsListRequestUserIdentityId) error
}

func (c *CredentialsListRequest) Accept(visitor CredentialsListRequestVisitor) error {
	switch c.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "credentialsListRequestZero":
		return visitor.VisitCredentialsListRequestZero(c.CredentialsListRequestZero)
	case "credentialsListRequestOne":
		return visitor.VisitCredentialsListRequestOne(c.CredentialsListRequestOne)
	case "credentialsListRequestTwo":
		return visitor.VisitCredentialsListRequestTwo(c.CredentialsListRequestTwo)
	case "credentialsListRequestUserIdentityId":
		return visitor.VisitCredentialsListRequestUserIdentityId(c.CredentialsListRequestUserIdentityId)
	}
}

type CredentialsListResponse struct {
	AcsCredentials []*CredentialsListResponseAcsCredentialsItem `json:"acs_credentials,omitempty"`
	Ok             bool                                         `json:"ok"`

	_rawJSON json.RawMessage
}

func (c *CredentialsListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CredentialsListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CredentialsListResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CredentialsListResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CredentialsUnassignResponse struct {
	AcsCredential *CredentialsUnassignResponseAcsCredential `json:"acs_credential,omitempty"`
	Ok            bool                                      `json:"ok"`

	_rawJSON json.RawMessage
}

func (c *CredentialsUnassignResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CredentialsUnassignResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CredentialsUnassignResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CredentialsUnassignResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CredentialsUpdateResponse struct {
	AcsCredential *CredentialsUpdateResponseAcsCredential `json:"acs_credential,omitempty"`
	Ok            bool                                    `json:"ok"`

	_rawJSON json.RawMessage
}

func (c *CredentialsUpdateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CredentialsUpdateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CredentialsUpdateResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CredentialsUpdateResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CredentialsUnassignRequest struct {
	AcsUserId       string `json:"acs_user_id"`
	AcsCredentialId string `json:"acs_credential_id"`
}

type CredentialsUpdateRequest struct {
	AcsCredentialId string `json:"acs_credential_id"`
	Code            string `json:"code"`
}
