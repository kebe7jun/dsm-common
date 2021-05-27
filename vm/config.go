package vm

type DSMAgentConfig struct {
	// Token defines the auth info to dsm server.
	Token string `json:"token,omitempty"`

	// Fetch config from dsm server.
	FromServer bool `json:"from_server,omitempty"`

	// The pilot IP for XDS.
	PilotIP string `json:"pilot_ip,omitempty" yaml:"pilot_ip,omitempty"`

	// The pilot IP for XDS.
	PilotHost string `json:"pilot_host,omitempty" yaml:"pilot_host,omitempty"`

	// Namespace of this vm.
	Namespace string `json:"namespace,omitempty"`
	Service   string `json:"service,omitempty"`
	Network   string `json:"network,omitempty"`

	// Only capture the traffic to these ports.
	Ports []uint16 `json:"ports,omitempty"`

	// Only capture the traffic to this IP,
	// If your VM have many IPs, this options may useful.
	// Empty for all IPs.
	IP string `json:"ip,omitempty"`
}
