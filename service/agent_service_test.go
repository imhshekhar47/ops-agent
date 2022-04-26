package service

import (
	"testing"

	"github.com/imhshekhar47/ops-agent/config"
	"github.com/imhshekhar47/ops-agent/pb"
)

var (
	id         string                     = "localhost"
	testConfig *config.AgentConfiguration = &config.AgentConfiguration{
		Core: config.CoreConfiguration{
			Version: "0.0.0",
		},
		Hostname: id,
		Uuid:     id,
		Address:  "localhost:5702",
	}

	s *AgentService = NewAgentService(testConfig)
)

func TestGet(t *testing.T) {

	if s.Get().Uuid != id {
		t.Error("failed to get agent")
		t.Fail()
	}
}

func TestGetHealth(t *testing.T) {

	if s.GetHealth().GetStatus() != pb.StatusCode_UP {
		t.Error("failed to get agent health")
		t.Fail()
	}
}
