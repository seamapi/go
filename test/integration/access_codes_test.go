package integration

import (
	"context"
	"testing"

	seamgo "github.com/seamapi/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccessCodes(t *testing.T) {
	t.Parallel()

	seam, cleanup := newFakeSeam(t)
	defer cleanup()

	ctx := context.Background()
	device := getTestDevice(t, seam)
	initialAccessCodes, err := seam.AccessCodes.List(
		ctx,
		&seamgo.AccessCodesListRequest{
			DeviceId: device.DeviceId,
		},
	)
	require.NoError(t, err)

	createdAccessCode, err := seam.AccessCodes.Create(
		ctx,
		&seamgo.AccessCodesCreateRequest{
			DeviceId: device.DeviceId,
			Name:     seamgo.String("Test code"),
			Code:     seamgo.String("4444"),
		},
	)
	require.NoError(t, err)
	assert.NotNil(t, createdAccessCode.Name)
	assert.NotNil(t, createdAccessCode.Code)
	assert.Equal(t, "Test code", *createdAccessCode.Name)
	assert.Equal(t, "4444", *createdAccessCode.Code)
	assert.Equal(t, seamgo.AccessCodeStatusSetting, createdAccessCode.Status)

	_, err = seam.AccessCodes.Create(
		ctx,
		&seamgo.AccessCodesCreateRequest{
			DeviceId: device.DeviceId,
			Name:     seamgo.String("Test code 2"),
			Code:     seamgo.String("5555"),
		},
	)
	require.NoError(t, err)

	accessCodes, err := seam.AccessCodes.List(
		ctx,
		&seamgo.AccessCodesListRequest{
			DeviceId: device.DeviceId,
		},
	)
	require.NoError(t, err)
	assert.Len(t, accessCodes, len(initialAccessCodes)+2)

	accessCodes, err = seam.AccessCodes.List(
		ctx,
		&seamgo.AccessCodesListRequest{
			DeviceId:      device.DeviceId,
			AccessCodeIds: []string{createdAccessCode.AccessCodeId},
		},
	)
	require.NoError(t, err)
	assert.Len(t, accessCodes, 1)

	gotAccessCode, err := seam.AccessCodes.Get(
		ctx,
		&seamgo.AccessCodesGetRequest{
			AccessCodeId: &createdAccessCode.AccessCodeId,
		},
	)
	require.NoError(t, err)
	assert.NotNil(t, gotAccessCode.Code)
	assert.Equal(t, "4444", *gotAccessCode.Code)

	// Fake Seam does not dedupe access codes.
	// _, err = seam.AccessCodes.Create(
	// 	ctx,
	// 	&seamgo.AccessCodesCreateRequest{
	// 		DeviceId: device.DeviceId,
	// 		Name:     seamgo.String("Duplicate Access Code"),
	// 		Code:     seamgo.String("4444"),
	// 	},
	// )
	// require.Error(t, err)

	_, err = seam.AccessCodes.Update(
		ctx,
		&seamgo.AccessCodesUpdateRequest{
			AccessCodeId: createdAccessCode.AccessCodeId,
			Name:         seamgo.String("Updated Name"),
		},
	)
	require.NoError(t, err)

	updatedAccessCode, err := seam.AccessCodes.Get(
		ctx,
		&seamgo.AccessCodesGetRequest{
			AccessCodeId: &createdAccessCode.AccessCodeId,
		},
	)
	require.NoError(t, err)
	assert.NotNil(t, updatedAccessCode.Name)
	assert.Equal(t, "Updated Name", *updatedAccessCode.Name)

	// Fake Seam does not like access code update.
	// _, err = seam.AccessCodes.Update(
	// 	ctx,
	// 	&seamgo.AccessCodesUpdateRequest{
	// 		AccessCodeId: createdAccessCode.AccessCodeId,
	// 		Type:         seamgo.AccessCodesUpdateRequestTypeTimeBound.Ptr(),
	// 		StartsAt:     seamgo.String("3001-01-01"),
	// 		EndsAt:       seamgo.String("3001-01-03"),
	// 	},
	// )
	// require.NoError(t, err)

	// accessCode, err := seam.AccessCodes.Get(
	// 	ctx,
	// 	&seamgo.AccessCodesGetRequest{
	// 		AccessCodeId: &createdAccessCode.AccessCodeId,
	// 	},
	// )
	// require.NoError(t, err)
	// assert.Equal(t, seamgo.AccessCodeTypeTimeBound, accessCode.Type)

	deleteAcitonAttempt, err := seam.AccessCodes.Delete(
		ctx,
		&seamgo.AccessCodesDeleteRequest{
			AccessCodeId: createdAccessCode.AccessCodeId,
		},
	)
	require.NoError(t, err)
	assert.NotNil(t, deleteAcitonAttempt.Success)

	deviceIds := make([]string, 0, len(accessCodes))
	for _, accessCode := range accessCodes {
		deviceIds = append(deviceIds, accessCode.DeviceId)
	}
	createdAccessCodes, err := seam.AccessCodes.CreateMultiple(
		ctx,
		&seamgo.AccessCodesCreateMultipleRequest{
			DeviceIds: deviceIds,
		},
	)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(createdAccessCodes), 1)
	assert.Equal(t, len(deviceIds), len(createdAccessCodes))

	commonCodes := make(map[string]struct{})
	for _, createdAccessCode := range createdAccessCodes {
		if commonCodeKey := createdAccessCode.CommonCodeKey; commonCodeKey != nil {
			commonCodes[*commonCodeKey] = struct{}{}
		}
	}
	assert.Len(t, commonCodes, 1)
}
