package service

import (
	"fmt"
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

	tAgent *pb.Agent = &pb.Agent{
		Uuid:    id,
		Address: fmt.Sprintf("%s:5702", id),
	}

	s *AgentService = NewAgentService(tAgent)
)

func TestGet(t *testing.T) {

	if s.GetAgent().Uuid != id {
		t.Error("failed to get agent")
		t.Fail()
	}
}

// func TestGetHealth(t *testing.T) {

// 	if s.GetAgent() != pb.StatusCode_UP {
// 		t.Error("failed to get agent health")
// 		t.Fail()
// 	}
// }
