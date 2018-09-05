package cmd

import (
	// "fmt"
	// "strings"
	"github.com/spf13/cobra"
	"github.com/HammerZ3it/clipam/config"
	"github.com/fatih/color"
	"os"
	"fmt"
	// "github.com/HammerZ3it/clipam/phpipam/client"
	// "github.com/HammerZ3it/clipam/phpipam"
	// "github.com/HammerZ3it/clipam/phpipam/session"
)

var hostName string
var subnetCIDR string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {

		//var sub config.Subnet
		
		sub, err := cli.GetSubnetByCIDR(subnetCIDR)
		if err != nil {
			color.Red(fmt.Sprintf("%s", err))
			os.Exit(1)
		}

		var AddressInput = config.Address{
			SubnetID:	sub.ID,
			IPAddress:	"",
			Description:	"Created with clipam tool",
			Hostname:	hostName,
		}

		cli.CreateAddress(AddressInput)

	},
}

//type Control config.Controller

func init() {
	addCmd.Flags().StringVar(
		&hostName, "name", "h", "Hostname of the new added server")
	addCmd.Flags().StringVar(
		&subnetCIDR, "subnet", "n", "Subnet of the new added server")

	rootCmd.AddCommand(addCmd)
}
