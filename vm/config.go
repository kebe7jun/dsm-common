package vm

type DSMAgentConfig struct {
	// DSM Server address, like http://10.31.11.1:30801/api
	Server string `json:"server,omitempty"`

	// Token defines the auth info to dsm server.
	Token string `json:"token,omitempty"`

	// Fetch config from dsm server.
	FromServer bool `json:"from_server,omitempty"`

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
