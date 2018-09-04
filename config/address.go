package config

import (
	"github.com/HammerZ3it/clipam/phpipam"
)

type Address struct {
	// The ID of the IP address entry within PHPIPAM.
	ID int `json:"id,string,omitempty"`
	// The ID of the subnet that the address belongs to.
	SubnetID int `json:"subnetId,string,omitempty"`
	// The IP address, without a CIDR subnet mask.
	IPAddress string `json:"ip,omitempty"`
	// true if this IP address is a gateway address.
	IsGateway phpipam.BoolIntString `json:"is_gateway,omitempty"`
	// A detailed description of the IP address entry.
	Description string `json:"description,omitempty"`
	// A hostname for the IP address.
	Hostname string `json:"hostname,omitempty"`
	// The MAC address for the IP.
	MACAddress string `json:"mac,omitempty"`
	// The address owner (customer, hostname, application, etc).
	Owner string `json:"owner,omitempty"`
	// The tag ID for the IP address.
	Tag int `json:"tag,string,omitempty"`
	// true if PTR records should not be created for this IP address.
	PTRIgnore phpipam.BoolIntString `json:"PTRIgnore,omitempty"`
	// The ID of a PowerDNS PTR record.
	PTRRecordID int `json:"PTR,string,omitempty"`
	// An ID of a device that this address belongs to.
	DeviceID int `json:"deviceId,string,omitempty"`
	// A switchport number/label that this IP address belongs to.
	Port string `json:"port,omitempty"`
	// A note for this IP address, detailing state information not sutiable for
	// entering in the description.
	Note string `json:"note,omitempty"`
	// A timestamp for when the address was last seen with ping.
	LastSeen string `json:"lastSeen,omitempty"`
	// true if you want to exclude this address from ping scans.
	ExcludePing phpipam.BoolIntString `json:"excludePing,omitempty"`
	// The date of the last edit to this resource.
	EditDate string `json:"editDate,omitempty"`
	// A map[string]interface{} of custom fields to set on the resource. Note
	// that this functionality requires PHPIPAM 1.3 or higher with the "Nest
	// custom fields" flag set on the specific API integration. If this is not
	// enabled, this map will be nil on GETs and POSTs and PATCHes with this
	// field set will fail. Use the explicit custom field functions instead.
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}
