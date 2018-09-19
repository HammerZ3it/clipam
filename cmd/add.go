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
	"github.com/HammerZ3it/clipam/config"
	"github.com/fatih/color"
	"os"
	"fmt"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an entry of a host into ipam",
	Long: `Usage of clipam add
	example : - clipam add --name ServerName --subnet "10.103.0.128/25"`,
	Run: func(cmd *cobra.Command, args []string) {

		sess := sessionConfig()
		cli := NewController(sess)

		sub, err := cli.GetSubnetInfo(subnetCIDR)
		if err != nil {
			color.Red(fmt.Sprintf("%s", err))
			os.Exit(1)
		}
		cli.GetFirstFreeIP(sub[0].ID)
		reservedIP := cli.GetFirstFreeIP(sub[0].ID)
		// fmt.Printf(reservedIP)
		var AddressInput = config.Address{
			SubnetID:	sub[0].ID,
			IPAddress:	"",
			Description:	"Created with clipam tool",
			Hostname:	hostName,
		}

		cli.CreateAddress(AddressInput)

		res := `{"reserved_ip": "` + reservedIP + `"}`
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
