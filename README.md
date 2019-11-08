# app-service-examples

### Overview

This repo contains various examples of Application Services based on the App Functions SDK. See the App Functions SDK for [README](https://github.com/edgexfoundry/app-functions-sdk-go/blob/v1.0.0-dev.2/README.md) for complete details on the SDK.

### Build Prerequisites

Please see the [edgex-go README](https://github.com/edgexfoundry/edgex-go/blob/master/README.md).

### Building Examples

The `makefile` is designed to build all/any examples under the `app-services` folder. Thus the makefile does not need to be updated when a new example is added to the `app-services` folder

​	run `make build` to build all examples.

​	run `make app-services/<example name>` to build a specific example, i.e. `make app-services/simple-filter-xml` 

For simplicity, the executable create for each example is named `app-service` and placed in that examples sub-folder.

### Running an Example

After building examples you simply cd to the folder for the example you want to run and run the executable for that example with or without any of the supported command line options.

The following commands will run the `simple-filter-xml` example

​	run `cd app-services/simple-filter-xml`

​	run `./app-service`

### Building App Service Docker Image

The  `simple-filter-xml` example contains an example `Dockerfile` to demonstrate how to build a **Docker Image** for your Application Service. 

The makefile also contains the `docker` target which will build the **Docker Image** for the `simple-filter-xml` example.

​	run `make docker`

> *Note that Application Services no longer use docker profiles. They use Environment Overrides in the docker compose file to make the necessary changes to the configuration for running in Docker. See the **Environment Variable Overrides For Docker** section in [App Service Configurable's README](https://github.com/edgexfoundry/app-service-configurable/blob/master/README.md#environment-variable-overrides-for-docker)* for more details and an example. 

### Profiles

The profiles folder contains example profiles for use with App Service Configurable. 