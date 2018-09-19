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

// gatewayCmd represents the gateway command
var gatewayCmd = &cobra.Command{
	Use: "clipam subnet gateway --subnet \"networkCIDR\"",
	Short: "Allow to retreive the IP Gateway of the given subnet",
	Long: `Usage of clipam subnet gateway
	example : - clipam subnet gateway --subnet "10.103.0.128/25"`,
	Run: func(cmd *cobra.Command, args []string) {

		sess := sessionConfig()
		cli := NewController(sess)

		sub, err := cli.GetSubnetInfo(subnetCIDR)
		if err != nil {
			color.Red(fmt.Sprintf("%s", err))
			os.Exit(1)
		}

		res := `{"subnet_gateway": "` + sub[0].Gateway + `"}`
		fmt.Println(res)
	},
}

func init() {
	subnetCmd.AddCommand(gatewayCmd)
}
