package config

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
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

func (c *Controller) GetSubnetByCIDR(cidr string) (out Subnet, err error) {
	err = c.SendRequest("GET", fmt.Sprintf("/subnets/cidr/%s/", cidr), &struct{}{}, &out)
	return
}
