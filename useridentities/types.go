// This file was auto-generated by Fern from our API Definition.

package useridentities

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/seamapi/go/core"
	time "time"
)

type EnrollmentAutomationsLaunchResponseEnrollmentAutomation struct {
	AcsCredentialProvisioningAutomationId string    `json:"acs_credential_provisioning_automation_id" url:"acs_credential_provisioning_automation_id"`
	CredentialManagerAcsSystemId          string    `json:"credential_manager_acs_system_id" url:"credential_manager_acs_system_id"`
	UserIdentityId                        string    `json:"user_identity_id" url:"user_identity_id"`
	CreatedAt                             time.Time `json:"created_at" url:"created_at"`
	WorkspaceId                           string    `json:"workspace_id" url:"workspace_id"`
	EnrollmentAutomationId                string    `json:"enrollment_automation_id" url:"enrollment_automation_id"`

	_rawJSON json.RawMessage
}

func (e *EnrollmentAutomationsLaunchResponseEnrollmentAutomation) UnmarshalJSON(data []byte) error {
	type embed EnrollmentAutomationsLaunchResponseEnrollmentAutomation
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at"`
	}{
		embed: embed(*e),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*e = EnrollmentAutomationsLaunchResponseEnrollmentAutomation(unmarshaler.embed)
	e.CreatedAt = unmarshaler.CreatedAt.Time()
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EnrollmentAutomationsLaunchResponseEnrollmentAutomation) MarshalJSON() ([]byte, error) {
	type embed EnrollmentAutomationsLaunchResponseEnrollmentAutomation
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at"`
	}{
		embed:     embed(*e),
		CreatedAt: core.NewDateTime(e.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (e *EnrollmentAutomationsLaunchResponseEnrollmentAutomation) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}
