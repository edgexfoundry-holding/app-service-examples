# App Functions SDK Azure Extension
## Overview
This is an extension to the App Functions SDK providing easiler integration with Azure IoT hub. It provides one transform and one export fucntion specific to the Azure IoT.

## The Azure Transform Function
This transform function can be used in the middle stage of the the function pipeline as other transforms. It will transform EdgeX events to name-value format to be consumed by Azure IoT Hub as the one below:
```json
{  
   "RandomValue_Int16":"32417"
}
```

## The Azure MQTT Export Function
This function can export transformed readings to Azure IoT via MQTT using some simple configurations. It also supports loading the required key/certificate pair from Vault (the distributed secret store used in the EdgeX architecture) in addition to key/certificate files from local file system. Just include a section of related configurations under ApplicationSettings in the app-functions standard toml configuration file as following:
```
IoTHub         = "EdgeX"
IoTDevice      = "MyDevice"
TokenPath      = "/vault/config/assets/resp-init.json"
VaultHost      = "localhost"
VaultPort      = "8200"
CertPath       = "v1/secret/edgex/pki/tls/azure"
MQTTCert       = "/secret/rsa_cert.pem"
MQTTKey        = "/secret/rsa_private.pem"
```

- `IoTHub` - Name of Azure IoT Hub 
- `IoTDevice` - Name on Azure IoT Device
- `TokenPath`, `VaultHost`, `VaultPort` and `CertPath` - Info for retrieving key/cert from Vault secret store
- `MQTTCert` and `MQTTKey` - Key/Certificate pait from loca file system. The pair will be used when the function fails to load a pair from Vault.

You can use LoadAzureMQTTConfig() to load the configurations to struct AzureMQTTConfig or create one yourself if the standard app-function sdk configuration file does not suit your needs.

See the code under `examples/azure-export` for a complete working example exporting reading to Azure IoT Hub.
