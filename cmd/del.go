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
	"github.com/fatih/color"
	"os"
	"fmt"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete an entry of a host into ipam",
	Long: `Usage of clipam del
	example : - clipam del --h ServerName --subnet "10.103.0.128/25"
						- clipam del --h ServerName --all`,
	Run: func(cmd *cobra.Command, args []string) {

		sess := sessionConfig()
		cli := NewController(sess)

		var addressesIDtoDel []int

		addrs, err := cli.GetAdressesID(hostName)
		if err != nil {
			color.Red(fmt.Sprintf("%s", err))
			os.Exit(1)
		}
		fmt.Println(subnetCIDR)
		if subnetCIDR == "" {
			for i := 0; i < len(addrs); i++ {
			// for i, addr := range addrs {
      	addressesIDtoDel = append(addressesIDtoDel, addrs[i].ID)
    	}
		} else {
			sub, err := cli.GetSubnetInfo(subnetCIDR)
			if err != nil {
				color.Red(fmt.Sprintf("%s", err))
				os.Exit(1)
			}
			for i := 0; i < len(addrs); i++ {
				if sub[0].ID == addrs[i].SubnetID {
					addressesIDtoDel = append(addressesIDtoDel, addrs[i].ID)
				}
    	}
		}
		cli.DeleteAddresses(addressesIDtoDel)
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
