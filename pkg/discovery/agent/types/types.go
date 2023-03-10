// Copyright 2023 NJWS Inc.

package types

const (
	App       = "discovery-bmc"
	Namespace = "proxy.foliage"

	// FunctionType will be as default qdsl to function
	FunctionType = FunctionPath
	Description  = "discovery redfish function"
)

const (
	RedfishDevicesContainerID = "types/redfish-device-container"
	FunctionContainerID       = "types/function-container"
	FunctionID                = "types/function"

	RedfishDeviceID = "types/redfish-device"

	RootID = "system/root"
)

const (
	RedfishDevicesLink    = "redfish-devices"
	FunctionContainerLink = "redfish"
	FunctionLink          = "discovery"
)

const (
	RedfishDevicesPath    = "redfish-devices.root"
	FunctionsPath         = "functions.root"
	FunctionContainerPath = "redfish.functions.root"
	FunctionPath          = "discovery.redfish.functions.root"
)
