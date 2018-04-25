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
	"context"
	"log"

	"github.com/spf13/cobra"
)

var cerr error

var resourceGroup string

// vmstartCmd represents the vmstart command
var vmstartCmd = &cobra.Command{
	Use:   "vmstart",
	Short: "Start all VMs for a ResourceGroup or K8 Cluster",
	Long: `Start all VMs for a ResourceGroup or K8 Cluster. For example:

goaz vmstart --ResourceGroup myresourcegroup`,
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()
		vmResultPage, err = computeClient.List(ctx, resourceGroup)

		if err != nil {
			log.Printf("Could not retrieve vms for ResourceGroup %q with error %q", resourceGroup, err)
		}

		for _, vm = range vmResultPage.Values() {
			_, err := computeClient.Start(ctx, resourceGroup, *vm.Name)
			if err != nil {
				log.Printf("Could not start vm %q for ResourceGroup %q with error %q", *vm.Name, resourceGroup, err)
			} else {
				log.Printf("Started vm %q for ResourceGroup %q", *vm.Name, resourceGroup)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(vmstartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//vmstartCmd.PersistentFlags().String("ResourceGroup", "", "Resource Group that you want to stop VMs for")
	vmstartCmd.Flags().StringVarP(&resourceGroup, "ResourceGroup", "", "", "Resource Group that you want to start VMs for if you are running ACS-Engine")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vmstartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	computeClient = getComputeClient()
}
