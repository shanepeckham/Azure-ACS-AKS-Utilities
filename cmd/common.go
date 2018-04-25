// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/Azure/go-autorest/autorest"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-04-01/compute"
	//"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2017-03-30/compute"
	auth "github.com/Azure/go-autorest/autorest/azure/auth"
)

var computeClient compute.VirtualMachinesClient
var authorizer autorest.Authorizer
var vmResultPage compute.VirtualMachineListResultPage
var err error

var vms func() []compute.VirtualMachine
var vm compute.VirtualMachine

// AcsCredential represents the credential file for GoAz
type AcsCredential struct {
	Cloud          string `json:"cloud"`
	TenantID       string `json:"tenantId"`
	SubscriptionID string `json:"subscriptionId"`
	ClientID       string `json:"clientId"`
	ClientSecret   string `json:"clientSecret"`
}

func getComputeClient() compute.VirtualMachinesClient {

	computeClient = compute.NewVirtualMachinesClient(os.Getenv("AZURE_SUBSCRIPTION_ID"))
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	//authorizer, err := auth.NewAuthorizerFromFile(os.Getenv("AZURE_AUTH_LOCATION"))

	if err == nil {
		computeClient.Authorizer = authorizer
	} else {
		println("Could not authenticate", err)
	}
	return computeClient
}

// LoadCredential returns an Credential struct from file path
func LoadCredential() {

	if os.Getenv("AZURE_AUTH_LOCATION") == "" {
		println("AZURE_AUTH_LOCATION is not set")
	}

	if os.Getenv("AZURE_TENANT_ID") == "" {
		println("AZURE_TENANT_ID is not set")
	}

	if os.Getenv("AZURE_CLIENT_ID") == "" {
		println("AZURE_CLIENT_ID is not set")
	}

	if os.Getenv("AZURE_CLIENT_SECRET") == "" {
		println("AZURE_CLIENT_SECRET is not set")
	}

	filepath := os.Getenv("AZURE_AUTH_LOCATION")
	log.Printf("Reading credential file %q", filepath)

	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		println("Reading credential file failed", filepath, err)
	}

	// Unmarshal the authentication file.
	var cred AcsCredential
	if err := json.Unmarshal(b, &cred); err != nil {
		println("Reading credential file failed", filepath, err)
	}
	os.Setenv("AZURE_TENANT_ID", cred.TenantID)
	os.Setenv("AZURE_CLIENT_ID", cred.ClientID)
	os.Setenv("AZURE_CLIENT_SECRET", cred.ClientSecret)
	os.Setenv("AZURE_SUBSCRIPTION_ID", cred.SubscriptionID)

	log.Printf("Load credential file %q successfully and set env vars", filepath)
}