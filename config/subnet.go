package config

import (
	"github.com/HammerZ3it/clipam/phpipam"
)

type Subnet struct {
	// The subnet ID.
	ID int `json:"id,string,omitempty"`
	// The subnet address, in dotted quad format (i.e. A.B.C.D).
	SubnetAddress string `json:"subnet,omitempty"`
	// The subnet's mask in number of bits (i.e. 24).
	Mask phpipam.JSONIntString `json:"mask,omitempty"`
	// A detailed description of the subnet.
	Description string `json:"description,omitempty"`
	// The section ID to add the subnet to (required when adding).
	SectionID int `json:"sectionId,string,omitempty"`
	// The ID of a linked IPv6 subnet.
	LinkedSubnet int `json:"linked_subnet,string,omitempty"`
	// The ID of the VLAN that this subnet belongs to.
	VLANID int `json:"vlanId,string,omitempty"`
	// The ID of the VRF this subnet belongs to.
	VRFID int `json:"vrfId,string,omitempty"`
	// The parent subnet ID if this is a nested subnet.
	MasterSubnetID int `json:"masterSubnetId,string,omitempty"`
	// The ID of the nameserver to attache the subnet to.
	NameserverID int `json:"nameserverId,string,omitempty"`
	// true if the name should be displayed in listing instead of the subnet
	// address.
	ShowName phpipam.BoolIntString `json:"showName,omitempty"`
	// A JSON object, stringified, that represents the permissions for this
	// section.
	Permissions string `json:"permissions,omitempty"`
	// Controls if PTR records should be created for the subnet.
	DNSRecursive phpipam.BoolIntString `json:"DNSrecursive,omitempty"`
	// Controls if DNS hostname records are displayed.
	DNSRecords phpipam.BoolIntString `json:"DNSrecords,omitempty"`
	// Controls if IP requests are allowed for the subnet.
	AllowRequests phpipam.BoolIntString `json:"allowRequests,omitempty"`
	// The ID of the scan agent to use for the subnet.
	ScanAgent int `json:"scanAgent,string,omitempty"`
	// Controls if the subnet should be included in status checks.
	PingSubnet phpipam.BoolIntString `json:"pingSubnet,omitempty"`
	// Controls if new hosts should be discovered for new host scans.
	DiscoverSubnet phpipam.BoolIntString `json:"discoverSubnet,omitempty"`
	// Controls if we are adding a subnet or folder.
	IsFolder phpipam.BoolIntString `json:"isFolder,omitempty"`
	// Marks the subnet as used.
	IsFull phpipam.BoolIntString `json:"isFull,omitempty"`
	// The threshold of the subnet.
	Threshold int `json:"threshold,string,omitempty"`
	// The location index of the subnet.
	Location int `json:"location,string,omitempty"`
	// The date of the last edit to this resource.
	EditDate string `json:"editDate,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}
