package config

import (
	"fmt"
	"strconv"
	"errors"
	"math/rand"
	"time"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/spf13/viper"
	"github.com/HammerZ3it/clipam/phpipam/client"
)

type Controller struct {
	client.Client
}

func (c *Controller) CreateAddress(in Address) (message string, err error) {
	if in.IPAddress == "" && in.SubnetID == 0 {
		return message, errors.New("ip address or subnet id must be defined")
	}

	if in.IPAddress == "" {
		// Retry
		for i := 0; i <= 5; i++ {
			//err = c.SendRequest("POST", fmt.Sprintf("/addresses/first_free/%v/", in.SubnetID), &in, &message)
			err = c.SendRequest("POST", "/addresses/first_free/", &in, &message)
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

func (c *Controller) GetSubnetInfo(cidr string) (out []Subnet, err error) {
	err = c.SendRequest("GET", fmt.Sprintf("/subnets/cidr/%s/", cidr), &struct{}{}, &out)
	return
}

func (c *Controller) GetAdressesID(hostname string) (out []Address, err error) {
	err = c.SendRequest("GET", fmt.Sprintf("/addresses/search_hostname/%s/", hostname), &struct{}{}, &out)
	return
}

func (c *Controller) GetFirstFreeIP(id int) string {
	var ret APIResponse
	// var retIP AuthAPIResponse

	client := &http.Client{}
  req, err := http.NewRequest("GET", viper.GetString("phpipam_server") + "/" + viper.GetString("phpipam_appid") + "/addresses/first_free/" + strconv.Itoa(id) + "/", nil)

	req.Header.Add("phpipam-token", c.Client.Session.Token.String)

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

	var s string
  err = json.Unmarshal(ret.Data, &s)
	if err != nil {
		panic(err)
	}
	return s

}

func (c *Controller) DeleteAddresses(adresses []int) (message string, err error) {
	for _, addr := range adresses {
		err = c.SendRequest("DELETE", fmt.Sprintf("/addresses/%d/", addr), &struct{}{}, &message)
	}
	return
}
