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

// vmstopCmd represents the vmstop command
var vmstopCmd = &cobra.Command{
	Use:   "vmstop",
	Short: "Stop all VMs for a resource-group or K8 Cluster",
	Long: `Stop all VMs for a resource-group or K8 Cluster. For example:

goaz vmstop (--resource-group myresourcegroup) and (--cluster-name myclustername) for AKS`,
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()

		// This is AKS so let's generate the resourceGroup that contains the  VMs
		if clusterName != "" {

			//Instantiate the AKS Client
			aksClient = getAKSClient()

			managedCluster, err = aksClient.Get(ctx, resourceGroup, clusterName)
			resourceGroup = "MC_" + resourceGroup + "_" + clusterName + "_" + *managedCluster.Location

			if err != nil {
				log.Printf("Could not retrieve AKS clusters for ResourceGroup %q with error %q", resourceGroup, err)
			}

		}

		vmResultPage, err = computeClient.List(ctx, resourceGroup)

		if err != nil {
			log.Printf("Could not retrieve vms for ResourceGroup %q with error %q", resourceGroup, err)
		}


		for _, vm = range vmResultPage.Values() {
			_, err := computeClient.Deallocate(ctx, resourceGroup, *vm.Name)
			if err != nil {
				log.Printf("Could not deallocate vm %q for ResourceGroup %q with error %q", *vm.Name, resourceGroup, err)
			} else {
				log.Printf("Deallocate vm %q for ResourceGroup %q", *vm.Name, resourceGroup)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(vmstopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vmstopCmd.PersistentFlags().String("foo", "", "A help for foo")
	vmstopCmd.Flags().StringVarP(&resourceGroup, "resource-group", "g", "", "Resource Group that you want to stop VMs for if you are running ACS-Engine")
	vmstopCmd.Flags().StringVarP(&clusterName, "cluster-name", "n", "", "Cluster that you want to stop VMs for if you are running AKS")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vmstopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	if authorizer == nil {
		LoadCredential()
	}
	computeClient = getComputeClient()
}
