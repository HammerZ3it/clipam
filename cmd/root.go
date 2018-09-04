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
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/fatih/color"
	// "github.com/HammerZ3it/clipam/phpipam/client"
	// "github.com/HammerZ3it/clipam/phpipam/session"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "clipam",
	Short: "A brief description of your application",
	Long: `A longer description`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// load config
		if err := initConfig(cfgFile); err != nil {
			return err
		}

		addr, err := viper.GetString("Host")
		if err != nil {
			return errors.New("IPAM Server address must be given")
		}

		appid, err := viper.GetString("AppID")
		if err != nil {
			return errors.New("IPAM appId must be given")
		}

		user, err := viper.GetString("User")
		if err != nil {
			return errors.New("IPAM user must be given")
		}

		password, err := viper.GetString("Password")
		if err != nil {
			return errors.New("IPAM user's password must be given")
		}

		sess := sessionConfig()
		client := NewController(sess)

	},
	RunE: func(cmd *cobra.Command, args []string) error {
		olor.Green("\nYour SSH Key is successfully signed")
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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.clipam.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func initConfig() {
	// expand ~ in file path
	expandedCfgFile, err := homedir.Expand(cfgFile)
	if err != nil {
		return err
	}
	// Use default config file if exists
	if _, err := os.Stat("/etc/clipam/config.yml"); err == nil {
		viper.SetConfigFile("/etc/clipam/config.yml")
		return viper.ReadInConfig()
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func NewController(sess *session.Session) *config.Controller {
	c := &config.Controller{
		Client: *client.NewClient(sess),
	}
	return c
}

func sessionConfig() *session.Session {
	return &session.Session{
		Config: config.Config{
			Host: viper.GetString("Host"),
			AppID: viper.GetString("AppID"),
			Username: viper.GetString("User"),
			Password: viper.GetString("Password"),
		},
		Token: session.Token{
			String: IpamAuthentification(Config),
		},
	}
}

func IpamAuthentification(cfg Config) string {
  var ret config.APIResponse
  var retAuth config.AuthAPIResponse

  client := &http.Client{}
  req, err := http.NewRequest("POST", config.Host + "/api/" + cfg.AppID + "/user/", nil)
  req.SetBasicAuth(cfg.User, cfg.Password)
  resp, err := client.Do(req)
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
