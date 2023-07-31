// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type EventsListRequestEventTypesItem uint8

const (
	EventsListRequestEventTypesItemDeviceConnected EventsListRequestEventTypesItem = iota + 1
	EventsListRequestEventTypesItemDeviceUnmanagedConnected
	EventsListRequestEventTypesItemDeviceDisconnected
	EventsListRequestEventTypesItemDeviceUnmanagedDisconnected
	EventsListRequestEventTypesItemDeviceConvertedToUnmanaged
	EventsListRequestEventTypesItemDeviceUnmanagedConvertedToManaged
	EventsListRequestEventTypesItemDeviceRemoved
	EventsListRequestEventTypesItemDeviceTampered
	EventsListRequestEventTypesItemDeviceLowBattery
	EventsListRequestEventTypesItemDeviceBatteryStatusChanged
	EventsListRequestEventTypesItemAccessCodeCreated
	EventsListRequestEventTypesItemAccessCodeChanged
	EventsListRequestEventTypesItemAccessCodeScheduledOnDevice
	EventsListRequestEventTypesItemAccessCodeSetOnDevice
	EventsListRequestEventTypesItemAccessCodeDeleted
	EventsListRequestEventTypesItemAccessCodeRemovedFromDevice
	EventsListRequestEventTypesItemAccessCodeFailedToSetOnDevice
	EventsListRequestEventTypesItemAccessCodeDelayInSettingOnDevice
	EventsListRequestEventTypesItemAccessCodeFailedToRemoveFromDevice
	EventsListRequestEventTypesItemAccessCodeDelayInRemovingFromDevice
	EventsListRequestEventTypesItemAccessCodeUnmanagedConvertedToManaged
	EventsListRequestEventTypesItemAccessCodeUnmanagedFailedToConvertToManaged
	EventsListRequestEventTypesItemAccessCodeUnmanagedCreated
	EventsListRequestEventTypesItemAccessCodeUnmanagedRemoved
	EventsListRequestEventTypesItemLockLocked
	EventsListRequestEventTypesItemLockUnlocked
	EventsListRequestEventTypesItemConnectedAccountConnected
	EventsListRequestEventTypesItemConnectedAccountCreated
	EventsListRequestEventTypesItemConnectedAccountDisconnected
	EventsListRequestEventTypesItemConnectedAccountCompletedFirstSync
	EventsListRequestEventTypesItemNoiseSensorNoiseThresholdTriggered
	EventsListRequestEventTypesItemAccessCodeBackupAccessCodePulled
)

func (e EventsListRequestEventTypesItem) String() string {
	switch e {
	default:
		return strconv.Itoa(int(e))
	case EventsListRequestEventTypesItemDeviceConnected:
		return "device.connected"
	case EventsListRequestEventTypesItemDeviceUnmanagedConnected:
		return "device.unmanaged.connected"
	case EventsListRequestEventTypesItemDeviceDisconnected:
		return "device.disconnected"
	case EventsListRequestEventTypesItemDeviceUnmanagedDisconnected:
		return "device.unmanaged.disconnected"
	case EventsListRequestEventTypesItemDeviceConvertedToUnmanaged:
		return "device.converted_to_unmanaged"
	case EventsListRequestEventTypesItemDeviceUnmanagedConvertedToManaged:
		return "device.unmanaged.converted_to_managed"
	case EventsListRequestEventTypesItemDeviceRemoved:
		return "device.removed"
	case EventsListRequestEventTypesItemDeviceTampered:
		return "device.tampered"
	case EventsListRequestEventTypesItemDeviceLowBattery:
		return "device.low_battery"
	case EventsListRequestEventTypesItemDeviceBatteryStatusChanged:
		return "device.battery_status_changed"
	case EventsListRequestEventTypesItemAccessCodeCreated:
		return "access_code.created"
	case EventsListRequestEventTypesItemAccessCodeChanged:
		return "access_code.changed"
	case EventsListRequestEventTypesItemAccessCodeScheduledOnDevice:
		return "access_code.scheduled_on_device"
	case EventsListRequestEventTypesItemAccessCodeSetOnDevice:
		return "access_code.set_on_device"
	case EventsListRequestEventTypesItemAccessCodeDeleted:
		return "access_code.deleted"
	case EventsListRequestEventTypesItemAccessCodeRemovedFromDevice:
		return "access_code.removed_from_device"
	case EventsListRequestEventTypesItemAccessCodeFailedToSetOnDevice:
		return "access_code.failed_to_set_on_device"
	case EventsListRequestEventTypesItemAccessCodeDelayInSettingOnDevice:
		return "access_code.delay_in_setting_on_device"
	case EventsListRequestEventTypesItemAccessCodeFailedToRemoveFromDevice:
		return "access_code.failed_to_remove_from_device"
	case EventsListRequestEventTypesItemAccessCodeDelayInRemovingFromDevice:
		return "access_code.delay_in_removing_from_device"
	case EventsListRequestEventTypesItemAccessCodeUnmanagedConvertedToManaged:
		return "access_code.unmanaged.converted_to_managed"
	case EventsListRequestEventTypesItemAccessCodeUnmanagedFailedToConvertToManaged:
		return "access_code.unmanaged.failed_to_convert_to_managed"
	case EventsListRequestEventTypesItemAccessCodeUnmanagedCreated:
		return "access_code.unmanaged.created"
	case EventsListRequestEventTypesItemAccessCodeUnmanagedRemoved:
		return "access_code.unmanaged.removed"
	case EventsListRequestEventTypesItemLockLocked:
		return "lock.locked"
	case EventsListRequestEventTypesItemLockUnlocked:
		return "lock.unlocked"
	case EventsListRequestEventTypesItemConnectedAccountConnected:
		return "connected_account.connected"
	case EventsListRequestEventTypesItemConnectedAccountCreated:
		return "connected_account.created"
	case EventsListRequestEventTypesItemConnectedAccountDisconnected:
		return "connected_account.disconnected"
	case EventsListRequestEventTypesItemConnectedAccountCompletedFirstSync:
		return "connected_account.completed_first_sync"
	case EventsListRequestEventTypesItemNoiseSensorNoiseThresholdTriggered:
		return "noise_sensor.noise_threshold_triggered"
	case EventsListRequestEventTypesItemAccessCodeBackupAccessCodePulled:
		return "access_code.backup_access_code_pulled"
	}
}

func (e EventsListRequestEventTypesItem) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", e.String())), nil
}

func (e *EventsListRequestEventTypesItem) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "device.connected":
		value := EventsListRequestEventTypesItemDeviceConnected
		*e = value
	case "device.unmanaged.connected":
		value := EventsListRequestEventTypesItemDeviceUnmanagedConnected
		*e = value
	case "device.disconnected":
		value := EventsListRequestEventTypesItemDeviceDisconnected
		*e = value
	case "device.unmanaged.disconnected":
		value := EventsListRequestEventTypesItemDeviceUnmanagedDisconnected
		*e = value
	case "device.converted_to_unmanaged":
		value := EventsListRequestEventTypesItemDeviceConvertedToUnmanaged
		*e = value
	case "device.unmanaged.converted_to_managed":
		value := EventsListRequestEventTypesItemDeviceUnmanagedConvertedToManaged
		*e = value
	case "device.removed":
		value := EventsListRequestEventTypesItemDeviceRemoved
		*e = value
	case "device.tampered":
		value := EventsListRequestEventTypesItemDeviceTampered
		*e = value
	case "device.low_battery":
		value := EventsListRequestEventTypesItemDeviceLowBattery
		*e = value
	case "device.battery_status_changed":
		value := EventsListRequestEventTypesItemDeviceBatteryStatusChanged
		*e = value
	case "access_code.created":
		value := EventsListRequestEventTypesItemAccessCodeCreated
		*e = value
	case "access_code.changed":
		value := EventsListRequestEventTypesItemAccessCodeChanged
		*e = value
	case "access_code.scheduled_on_device":
		value := EventsListRequestEventTypesItemAccessCodeScheduledOnDevice
		*e = value
	case "access_code.set_on_device":
		value := EventsListRequestEventTypesItemAccessCodeSetOnDevice
		*e = value
	case "access_code.deleted":
		value := EventsListRequestEventTypesItemAccessCodeDeleted
		*e = value
	case "access_code.removed_from_device":
		value := EventsListRequestEventTypesItemAccessCodeRemovedFromDevice
		*e = value
	case "access_code.failed_to_set_on_device":
		value := EventsListRequestEventTypesItemAccessCodeFailedToSetOnDevice
		*e = value
	case "access_code.delay_in_setting_on_device":
		value := EventsListRequestEventTypesItemAccessCodeDelayInSettingOnDevice
		*e = value
	case "access_code.failed_to_remove_from_device":
		value := EventsListRequestEventTypesItemAccessCodeFailedToRemoveFromDevice
		*e = value
	case "access_code.delay_in_removing_from_device":
		value := EventsListRequestEventTypesItemAccessCodeDelayInRemovingFromDevice
		*e = value
	case "access_code.unmanaged.converted_to_managed":
		value := EventsListRequestEventTypesItemAccessCodeUnmanagedConvertedToManaged
		*e = value
	case "access_code.unmanaged.failed_to_convert_to_managed":
		value := EventsListRequestEventTypesItemAccessCodeUnmanagedFailedToConvertToManaged
		*e = value
	case "access_code.unmanaged.created":
		value := EventsListRequestEventTypesItemAccessCodeUnmanagedCreated
		*e = value
	case "access_code.unmanaged.removed":
		value := EventsListRequestEventTypesItemAccessCodeUnmanagedRemoved
		*e = value
	case "lock.locked":
		value := EventsListRequestEventTypesItemLockLocked
		*e = value
	case "lock.unlocked":
		value := EventsListRequestEventTypesItemLockUnlocked
		*e = value
	case "connected_account.connected":
		value := EventsListRequestEventTypesItemConnectedAccountConnected
		*e = value
	case "connected_account.created":
		value := EventsListRequestEventTypesItemConnectedAccountCreated
		*e = value
	case "connected_account.disconnected":
		value := EventsListRequestEventTypesItemConnectedAccountDisconnected
		*e = value
	case "connected_account.completed_first_sync":
		value := EventsListRequestEventTypesItemConnectedAccountCompletedFirstSync
		*e = value
	case "noise_sensor.noise_threshold_triggered":
		value := EventsListRequestEventTypesItemNoiseSensorNoiseThresholdTriggered
		*e = value
	case "access_code.backup_access_code_pulled":
		value := EventsListRequestEventTypesItemAccessCodeBackupAccessCodePulled
		*e = value
	}
	return nil
}
