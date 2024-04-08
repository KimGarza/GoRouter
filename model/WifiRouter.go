package model

type Router struct {
	NAT                    string
	PAT                    string
	Wireless               string
	DHCP                   string
	Firewall               string
	VPN                    string
	QoS                    string
	NetworkInterfacePorts  string
	GuestAccess            string
	Firmware               string
	SecurityEncryption     string // (WPA, WPA2, WPA3)
	IPv4                   string
	IPv6                   string
	ConfigurationInterface string
	Antennas               string
	DMZ                    string
	RouteTable             string
}
