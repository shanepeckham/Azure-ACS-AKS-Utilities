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
	"reflect"
	"testing"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-04-01/compute"
)

func Test_getComputeClient(t *testing.T) {
	tests := []struct {
		name string
		want compute.VirtualMachinesClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getComputeClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getComputeClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
