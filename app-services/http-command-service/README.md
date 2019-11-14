# HTTP Command Service

We sometimes want to trigger device actions on the edge from the cloud. EdgeX provide a comprehensive set of APIs for you to do that : https://github.com/edgexfoundry/edgex-go/blob/master/api/raml/core-command.raml. But often times you don't want to expose the entire API and you want finer-grained control over which and how the APIs should be exposed. For example, you want to control which command on which device is allowed to receive commands from the outside of EdgeX or want to only allow certain values for a command. 

HTTP Command Service is an example of doing just that. It use the simple device service from device SDK as an example. Instead of exposing the entire device APIs of the device, We want to allow a simple json document to set status of the switch in the device as:
```json
{
    "status" : "off"
}
```
The HTTP Command Service exposes a HTTP service for the client to switch on / off of the device without knowing the underlying EdgeX APIs.

