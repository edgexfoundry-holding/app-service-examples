//
// Copyright (c) 2020 Intel Corporation
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"fmt"
	"os"

	"github.com/edgexfoundry/app-functions-sdk-go/appcontext"
)

type getSecretsTestData struct {
	testName    string
	path        string
	keys        []string
	errExpected bool
}

var secureSecretsData = []getSecretsTestData{
	getSecretsTestData{
		testName:    "Empty path, valid key",
		path:        "",
		keys:        []string{"username"},
		errExpected: false,
	},
	getSecretsTestData{
		testName:    "Fake path, return error",
		path:        "/fakepath",
		keys:        []string{"username"},
		errExpected: true,
	},
	getSecretsTestData{
		testName:    "Valid path, empty keys list, return all secrets",
		path:        "/mongodb",
		keys:        []string{},
		errExpected: false,
	},
	getSecretsTestData{
		testName:    "Valid path, empty string keys, return error",
		path:        "/mongodb",
		keys:        []string{"", ""},
		errExpected: true,
	},
	getSecretsTestData{
		testName:    "Valid path, valid key",
		path:        "/mongodb",
		keys:        []string{"username"},
		errExpected: false,
	},
}

var insecureSecretsData = []getSecretsTestData{
	getSecretsTestData{
		testName:    "Empty path, valid key",
		path:        "",
		keys:        []string{"password"},
		errExpected: false,
	},
	getSecretsTestData{
		testName:    "Fake path, return error",
		path:        "/fakepath",
		keys:        []string{"password"},
		errExpected: true,
	},
	getSecretsTestData{
		testName:    "Valid path, empty keys list, return all secrets",
		path:        "/aws",
		keys:        []string{},
		errExpected: false,
	},
	getSecretsTestData{
		testName:    "Valid path, empty string keys, return error",
		path:        "/aws",
		keys:        []string{"", ""},
		errExpected: true,
	},
	getSecretsTestData{
		testName:    "Valid path, valid key",
		path:        "/aws",
		keys:        []string{"password"},
		errExpected: false,
	},
}

func getSecrets(edgexcontext *appcontext.Context, data []getSecretsTestData, securityEnabled bool) error {

	var origEnv string
	var err error

	if !securityEnabled {
		if origEnv, err = disableSecureStore(edgexcontext); err != nil {
			edgexcontext.LoggingClient.Error("Failed to set env variable: EDGEX_SECURITY_SECRET_STORE")
			return fmt.Errorf("Failed to set env variable: EDGEX_SECURITY_SECRET_STORE")
		}
		edgexcontext.LoggingClient.Debug("Secure Store DISABLED")

		defer resetSecureStoreEnv(edgexcontext, origEnv)
	} else {
		edgexcontext.LoggingClient.Debug("Secure Store ENABLED")
	}

	for _, test := range data {
		edgexcontext.LoggingClient.Info(fmt.Sprintf("------- Test: %v -------", test.testName))
		secrets, err := edgexcontext.SecretProvider.GetSecrets(test.path, test.keys...)
		if test.errExpected {
			if err == nil {
				edgexcontext.LoggingClient.Error("Expected error but got no error")
			}

			edgexcontext.LoggingClient.Error(fmt.Sprintf("Couldn't get secrets: %v", err.Error()))
		} else {
			if err != nil {
				edgexcontext.LoggingClient.Error(fmt.Sprintf("Got error when expected NO error: %v", err.Error()))
			}

			for k, v := range secrets {
				edgexcontext.LoggingClient.Info(fmt.Sprintf("key:%v, value:%v", k, v))
			}
		}
	}

	return nil
}

func disableSecureStore(edgexcontext *appcontext.Context) (origEnv string, err error) {

	origEnv = os.Getenv("EDGEX_SECURITY_SECRET_STORE")
	err = os.Setenv("EDGEX_SECURITY_SECRET_STORE", "false")

	return origEnv, err
}

func resetSecureStoreEnv(edgexcontext *appcontext.Context, origEnv string) {
	if err := os.Setenv("EDGEX_SECURITY_SECRET_STORE", origEnv); err != nil {
		edgexcontext.LoggingClient.Error("Failed to set env variable: EDGEX_SECURITY_SECRET_STORE back to original value")
	}
}
