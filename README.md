# app-service-examples

## Overview

This repo contains various examples of Application Services based on the App Functions SDK. See the App Functions SDK for [README](https://github.com/edgexfoundry/app-functions-sdk-go/blob/v1.0.0-dev.2/README.md) for complete details on the SDK.

## Build Prerequisites

Please see the [edgex-go README](https://github.com/edgexfoundry/edgex-go/blob/master/README.md).

## Quick Start

If you want to see how all of EdgeX works - you can leverage your own Azure account and deploy EdgeX to the cloud. This template leverages Azure Container Instances and will deploy a single group called "edgex-example" with 12 services deployed with 2.4 vCPUs allocated (0.2 vCPUs per service) and 6GB of RAM allocated (0.5 GB per service) with an estimated cost of $0.14904 / hour or $3.57696 / day. 

1 container groups * 3600 seconds * 2.4 vCPU * $0.0000135 per vCPU-s  = ~$0.11664

1 container groups * 3600 seconds * 6 GB * $0.0000015 per GB-s  = $0.0324

memory($0.0324) + cpu($0.11664) = $0.14904 / hour
= $3.57696 / day

[![Deploy to Azure](https://aka.ms/deploytoazurebutton)](https://portal.azure.com/#create/Microsoft.Template/uri/https%3A%2F%2Fraw.githubusercontent.com%2Fedgexfoundry-holding%2Fapp-service-examples%2Fmaster%2Ftemplates%2Fazuredeploy.json).



## Building Examples

The `makefile` is designed to build all/any examples under the `app-services` folder. Thus the makefile does not need to be updated when a new example is added to the `app-services` folder

​	run `make build` to build all examples.

​	run `make app-services/<example name>` to build a specific example, i.e. `make app-services/simple-filter-xml` 

For simplicity, the executable create for each example is named `app-service` and placed in that examples sub-folder.

## Running an Example

After building examples you simply cd to the folder for the example you want to run and run the executable for that example with or without any of the supported command line options.

The following commands will run the `simple-filter-xml` example

​	run `cd app-services/simple-filter-xml`

​	run `./app-service`

## Building App Service Docker Image

The  `simple-filter-xml` example contains an example `Dockerfile` to demonstrate how to build a **Docker Image** for your Application Service. 

The makefile also contains the `docker` target which will build the **Docker Image** for the `simple-filter-xml` example.

​	run `make docker`

> *Note that Application Services no longer use docker profiles. They use Environment Overrides in the docker compose file to make the necessary changes to the configuration for running in Docker. See the **Environment Variable Overrides For Docker** section in [App Service Configurable's README](https://github.com/edgexfoundry/app-service-configurable/blob/master/README.md#environment-variable-overrides-for-docker)* for more details and an example. 



## Profiles

The profiles folder contains example profiles for use with App Service Configurable. 