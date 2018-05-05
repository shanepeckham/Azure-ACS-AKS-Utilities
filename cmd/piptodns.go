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
	"github.com/spf13/cobra"
	"log"
	"os"
	"context"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-04-01/network"
)

// piptodnsCmd represents the piptodns command
var piptodnsCmd = &cobra.Command{
	Use:   "piptodns",
	Short: "Assign PIP to DNS",
	Long: `Assign an Azure Public IP Address to a FQDN. Just provide the prefix domain label
	e.g. mydns will become mydns.<location>.cloudapp.azure.com`,
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

		ctx = context.Background()
		pipResultPage, err = pipClient.List(ctx, resourceGroup)

		if err != nil {
			log.Printf("Could not find Public IP for IP Address %q with error %q", publicIP, err)
			os.Exit(1)
		}

		for _, ips = range pipResultPage.Values() {

			if err != nil {
				log.Printf("Could not start vm %q for ResourceGroup %q with error %q", *vm.Name, resourceGroup, err)
			} else {
				dnsSettings := network.PublicIPAddressDNSSettings{}
				dnsSettings.Fqdn = &dnsName
				dnsSettings.DomainNameLabel = &dnsName

				ips.DNSSettings = &dnsSettings
				_, err = pipClient.CreateOrUpdate(ctx, resourceGroup, *ips.Name, ips)

				if err != nil {
					log.Printf("Could not update DNS %q for IP Address %q with error %q", dnsName, publicIP, err)
					os.Exit(1)
				}

				log.Printf("Associated DNS %q with Public IP %q and address %q", dnsName, *ips.Name, publicIP)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(piptodnsCmd)

	// Here you will define your flags and configuration settings.
	piptodnsCmd.Flags().StringVarP(&resourceGroup, "resource-group", "g", "", "Resource Group that you want to start VMs for if you are running ACS-Engine")
	piptodnsCmd.Flags().StringVarP(&publicIP, "pipAddress", "p", "", "The Public IP Address of your Ingress")
	piptodnsCmd.Flags().StringVarP(&clusterName, "cluster-name", "n", "", "Cluster that you want to stop VMs for if you are running AKS")
	piptodnsCmd.Flags().StringVarP(&dnsName, "dns-name", "d", "", "Your unique dns prefix")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// piptodnsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// piptodnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	if authorizer == nil {
		LoadCredential()
	}
	pipClient = getPIPClient()
}
