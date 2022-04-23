package config

import (
	"testing"
)

var (
	testConfig *AgentConfiguration = &AgentConfiguration{
		Core: CoreConfiguration{
			Version: "0.0.0.0",
		},
		Hostname: "localhost",
		Uuid:     "localhost",
		Address:  "localhost:5702",
	}
)

func TestCoreConfig(t *testing.T) {
	if nil == testConfig {
		t.Fail()
	}

	if testConfig.Core.Version != "0.0.0.0" {
		t.Errorf("failed to get test confgration")
	}
}
