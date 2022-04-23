package service

import (
	"testing"

	"github.com/imhshekhar47/ops-agent/config"
	"github.com/imhshekhar47/ops-agent/pb"
)

var (
	testConfig *config.AgentConfiguration = &config.AgentConfiguration{
		Core: config.CoreConfiguration{
			Version: "0.0.0",
		},
		Hostname: "localhost",
		Uuid:     "localhost",
		Address:  "localhost:5702",
	}

	s *AgentService = NewAgentService(testConfig)
)

func TestGet(t *testing.T) {

	if s.Get().Uuid != "localhost" {
		t.Logf("failed to get agent")
		t.Fail()
	}
}

func TestGetHealth(t *testing.T) {

	if s.GetHealth().GetStatus() != pb.StatusCode_UP {
		t.Logf("failed to get agent health")
		t.Fail()
	}
}
