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
	"os"
	"encoding/json"
	"io/ioutil"
  "net/http"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/fatih/color"
	"github.com/HammerZ3it/clipam/phpipam"
	"github.com/HammerZ3it/clipam/phpipam/client"
	"github.com/HammerZ3it/clipam/phpipam/session"
	"github.com/HammerZ3it/clipam/config"
)

var configFile string
var cli config.Controller
var sess session.Session

var hostName string
var subnetCIDR string
var IPaddr string
var all bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "clipam",
	Short: "clipam allows to interact with the API of a PHPIPAM server",
	Long: `Instead of directly call the API, clipam offers a way to interact easier with it
	example : - ...
						- ...
						-	...`,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// load config
		if err := initConfig(configFile); err != nil {
					return err
		}

		addr := viper.GetString("phpipam_server")
		if addr == "" {
			return errors.New("IPAM Server address must be given")
		}

		appid := viper.GetString("phpipam_appid")
		if appid == "" {
			return errors.New("IPAM appId must be given")
		}

		user := viper.GetString("phpipam_user")
		if user == "" {
			return errors.New("IPAM user must be given")
		}

		password := viper.GetString("phpipam_password")
		if password == "" {
			return errors.New("IPAM user's password must be given")
		}

		// sess := sessionConfig()
		// cli := NewController(sess)
		//
		// fmt.Println(cli)

		return nil

	},
	RunE: func(cmd *cobra.Command, args []string) error {
		color.Green("\nYour SSH Key is successfully signed")
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
			color.Red(fmt.Sprintf("Error: %s", err))
			os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/clipam.yml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVar(
		&hostName, "name", "", "Hostname of the server")
	rootCmd.PersistentFlags().StringVar(
		&subnetCIDR, "subnet", "", "Subnet of the server")
	rootCmd.PersistentFlags().StringVar(
		&IPaddr, "ip", "", "Ip of the entry you want to delete")
	rootCmd.PersistentFlags().BoolVarP(
		&all, "all", "a", true, "Delete all entry of the selected hostname in the IPAM database (serveur with multiple nic)")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subnetNetmaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subnetNetmaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func initConfig(cfgFile string) error {
	// viper.SetEnvPrefix("clipam")
	// viper.AutomaticEnv()

	// expand ~ in file path
	expandedCfgFile, err := homedir.Expand(cfgFile)
	if err != nil {
		return err
	}
	// Use default config file if exists
	if _, err := os.Stat(expandedCfgFile); err == nil {
		viper.SetConfigFile(expandedCfgFile)
		return viper.ReadInConfig()
	}

	// Use default config file if exists
	if _, err := os.Stat("/etc/clipam/config.yml"); err == nil {
		viper.SetConfigFile("/etc/clipam/config.yml")
		return viper.ReadInConfig()
	}
	return nil
}

func NewController(sess *session.Session) *config.Controller {
	c := &config.Controller{
		Client: *client.NewClient(sess),
	}
	return c
}

func sessionConfig() *session.Session {
	return &session.Session{
		Config: phpipam.Config{
			Endpoint: viper.GetString("phpipam_server"),
			AppID: viper.GetString("phpipam_appid"),
			Username: viper.GetString("phpipam_user"),
			Password: viper.GetString("phpipam_password"),
		},
		Token: session.Token{
			String: IpamAuthentification(),
		},
	}
}

func IpamAuthentification() string {
  var ret config.APIResponse
  var retAuth config.AuthAPIResponse

  client := &http.Client{}
  req, err := http.NewRequest("POST", viper.GetString("phpipam_server") + "/" + viper.GetString("phpipam_appid") + "/user/", nil)
  req.SetBasicAuth(viper.GetString("phpipam_user"), viper.GetString("phpipam_password"))
  resp, err := client.Do(req)
	// if err != nil {
  //   return nil, err
	// }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  err = json.Unmarshal(body, &ret)
  if err != nil {
    panic(err)
  }
  bytes := []byte(string(ret.Data))
  err = json.Unmarshal(bytes, &retAuth)
  if err != nil {
    panic(err)
  }
  return retAuth.Token
}
