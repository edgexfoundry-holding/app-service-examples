package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	azureTransforms "github.com/IOTechSystems/app-functions-azure/pkg/transforms"
	"github.com/edgexfoundry/app-functions-sdk-go/appsdk"
	"github.com/edgexfoundry/app-functions-sdk-go/pkg/transforms"
)

const (
	serviceKey           = "AzureExport"
	appConfigDeviceNames = "DeviceNames"
)

var counter int

func main() {
	// 1) First thing to do is to create an instance of the EdgeX SDK and initialize it.
	edgexSdk := &appsdk.AppFunctionsSDK{ServiceKey: serviceKey}
	if err := edgexSdk.Initialize(); err != nil {
		edgexSdk.LoggingClient.Error(fmt.Sprintf("SDK initialization failed: %v\n", err))
		os.Exit(-1)
	}

	// 2) Since our DeviceNameFilter Function requires the list of device names we would
	// like to search for, we'll go ahead and define that now.
	deviceNames, err := loadDeviceNames(edgexSdk)

	if err != nil {
		edgexSdk.LoggingClient.Error(fmt.Sprintf("Failed to load device names: %v\n", err))
		os.Exit(-1)
	}

	// 3) This is our pipeline configuration, the collection of functions to
	// execute every time an event is triggered.

	// Load Azure-specific MQTT configuration from App SDK
	// You can also create AzureMQTTConfig struct yourself
	config, err := azureTransforms.LoadAzureMQTTConfig(edgexSdk)

	if err != nil {
		edgexSdk.LoggingClient.Error(fmt.Sprintf("Failed to load Azure MQTT configurations: %v\n", err))
		os.Exit(-1)
	}

	edgexSdk.SetFunctionsPipeline(
		transforms.NewFilter(deviceNames).FilterByDeviceName,
		azureTransforms.NewConversion().TransformToAzure,
		azureTransforms.NewAzureMQTTSender(edgexSdk.LoggingClient, config).MQTTSend,
	)

	// 5) Lastly, we'll go ahead and tell the SDK to "start" and begin listening for events
	// to trigger the pipeline.
	err = edgexSdk.MakeItRun()
	if err != nil {
		edgexSdk.LoggingClient.Error("MakeItRun returned error: ", err.Error())
		os.Exit(-1)
	}

	// Do any required cleanup here

	os.Exit(0)
}

func loadDeviceNames(edgexSdk *appsdk.AppFunctionsSDK) ([]string, error) {
	if value, ok := edgexSdk.ApplicationSettings()[appConfigDeviceNames]; ok {
		deviceNames := strings.Split(value, ",")

		if len(deviceNames) < 1 {
			return nil, errors.New("No device is configured")
		}

		for index, name := range deviceNames {
			deviceNames[index] = strings.TrimSpace(name)
		}

		return deviceNames, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Couldn't find '%s' in configuration file: %v\n", appConfigDeviceNames, ok))
	}
}
