package config

type CoreConfiguration struct {
	Version string `json:"version"`
}

type AgentConfiguration struct {
	Core CoreConfiguration `json:"core"`

	Hostname string `json:"hostname" yaml:"hostname"`
	Uuid     string `json:"uuid" yaml:"uuid"`
	Address  string `json:"address" yaml:"address"`
}

func (c *AgentConfiguration) UpdateAddress(address string) {
	c.Address = address
}
