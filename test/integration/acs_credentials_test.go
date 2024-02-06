package integration

import (
	"context"
	"log"
	"testing"
	"time"

	api "github.com/seamapi/go"
	"github.com/seamapi/go/acs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func test(t *testing.T) {
	t.Parallel()

	seam, cleanup := newFakeSeam(t)
	defer cleanup()

	systems, sErr := seam.Acs.Systems.List(context.Background(), nil)

	if sErr != nil {
		log.Panic(sErr)
	}

	var visionlineSystem *api.AcsSystem

	for _, s := range systems.AcsSystems {
		if s.ExternalType == "visionline_system" {
			visionlineSystem = s
		}
	}

	users, err := seam.Acs.Users.List(context.Background(), &acs.UsersListRequest{
		AcsSystemId: &visionlineSystem.AcsSystemId,
	})

	isMultiPhoneSyncCredential := true

	require.NoError(t, err)

	startsAt := time.Now()

	endsAt := time.Now().Add(24 * time.Hour)

	credentialResponse, uErr := seam.Acs.Credentials.Create(context.Background(), &acs.CredentialsCreateRequest{
		AcsUserId:                  users.AcsUsers[0].AcsUserId,
		AccessMethod:               "mobile_key",
		IsMultiPhoneSyncCredential: &isMultiPhoneSyncCredential,
		StartsAt:                   &startsAt,
		EndsAt:                     &endsAt,
		VisionlineMetadata: &acs.CredentialsCreateRequestVisionlineMetadata{
			CardFormat: acs.CredentialsCreateRequestVisionlineMetadataCardFormatRfid48.Ptr(),
		},
	})

	require.NoError(t, uErr)
	assert.NotEmpty(t, credentialResponse.AcsCredential)
	assert.True(t, credentialResponse.Ok)
	assert.NotEmpty(t, credentialResponse.AcsCredential.StartsAt)
	assert.NotEmpty(t, credentialResponse.AcsCredential.EndsAt)
}
