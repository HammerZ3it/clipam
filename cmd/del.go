package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var IPaddr string

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "A brief description of your command",
	Long: `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("del called")

		// client.DeleteAddress(client.GetSubnetIDByIP(IPaddr).ID, IPaddr)

	},
}

func init() {

	addCmd.Flags().StringVar(
		&IPaddr, "ip", "h", "Ip of the entry you want to delete")

	rootCmd.AddCommand(delCmd)
}

// func (c *Control) GetSubnetIDByIP(ipaddr string) (out []config.Address, err error) {
// 	err = c.SendRequest("GET", fmt.Sprintf("/addresses/search/%s/", ipaddr), &struct{}{}, &out)
// 	return
// }
//
// func (c *Control) DeleteAddress(id int, ipaddr string) (message string, err error) {
// 	var input = config.Address{
// 		SubnetID:    id,
// 		IPAddress:   ipaddr,
// 	}
// 	err = c.SendRequest("DELETE", fmt.Sprintf("/addresses/%s/%d/", ipaddr, id), &input, &message)
// 	return
// }
