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
	"fmt"
	"errors"
	"math/rand"
	"time"
	"github.com/spf13/cobra"
	"github.com/HammerZ3it/clipam/config"
	"github.com/HammerZ3it/clipam/phpipam/session"
)

var hostName string
var subnetCIDR string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {

		// load config
		if err := initConfig(serverCfgFile); err != nil {
			return err
		}

		var AddressInput = config.Address{
			SubnetID:    client.GetSubnetsByCIDR(subnetCIDR).ID,
			IPAddress:   "",
			Description: "Created with clipam tool",
			Hostname: hostName,
			Tag: "Added by clipam",
		}

		client.CreateAddress(AddressInput)

	},
}

type Control config.Controller

func (c *Control) CreateAddress(in config.Address) (message string, err error) {
	if in.IPAddress == "" && in.SubnetID == 0 {
		return message, errors.New("ip address or subnet id must be defined")
	}

	if in.IPAddress == "" {
		// Retry
		for i := 0; i <= 5; i++ {
			err = c.SendRequest("POST", "/addresses/first_free", &in, &message)
			if err == nil {
				break
			}
			r := rand.Intn(500)
			time.Sleep(time.Duration(r) * time.Microsecond)
		}
	} else {
		err = c.SendRequest("POST", "/addresses/", &in, &message)
	}

	return
}

func (c *Control) GetSubnetsByCIDR(cidr string) (out []config.Subnet, err error) {
	err = c.SendRequest("GET", fmt.Sprintf("/subnets/cidr/%s/", cidr), &struct{}{}, &out)
	return
}

func init() {
	addCmd.Flags.StringVar(
		&hostName, "name", "h", "Hostname of the new added server")
	addCmd.Flags.StringVar(
		&subnetCIDR, "subnet", "n", "network", "Subnet of the new added server")

	rootCmd.AddCommand(addCmd)
}
