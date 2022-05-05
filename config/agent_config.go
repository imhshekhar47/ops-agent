package config

type CoreConfiguration struct {
	Version string `json:"version"`
}

type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type AgentConfiguration struct {
	Core CoreConfiguration `json:"core"`

	Hostname string `json:"hostname" yaml:"hostname"`
	Uuid     string `json:"uuid" yaml:"uuid"`
	Address  string `json:"address" yaml:"address"`

	Group       string   `json:"group" yaml:"group"`
	Component   string   `json:"component" yaml:"component"`
	Environment string   `json:"environment" yaml:"environment"`
	Site        string   `json:"site" yaml:"site"`
	Location    Location `json:"location" yaml:"location"`
}

func (c *AgentConfiguration) UpdateAddress(address string) {
	c.Address = address
}
