package server

import (
	"context"
	"testing"

	"github.com/imhshekhar47/ops-agent/config"
	"github.com/imhshekhar47/ops-agent/pb"
	"github.com/imhshekhar47/ops-agent/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	id      string                     = "localhost"
	tConfig *config.AgentConfiguration = &config.AgentConfiguration{
		Core: config.CoreConfiguration{
			Version: "0.0.0",
		},
		Hostname: id,
		Uuid:     id,
	}

	// test agent definition
	tAgent *pb.Agent = &pb.Agent{
		Uuid:    id,
		Address: "localhost:5701",
		Status:  pb.StatusCode_UP,
		Info: &pb.HierarchyInfo{
			Group:       "test",
			Component:   "test",
			Environment: "test",
			Site:        "test",
			Location: &pb.GeoLocation{
				Latitude:  "0",
				Longitude: "0",
			},
		},
		Server: &pb.Server{
			AgentId:   id,
			Uuid:      id,
			Hostname:  id,
			Status:    pb.StatusCode_UP,
			Databases: make([]*pb.Database, 0),
		},
	}

	tService *service.AgentService = service.NewAgentService(tAgent)
	tLogger  *logrus.Logger        = logrus.New()
	tServer  *AgentServer          = NewAgentServer(tLogger, tService)
)

func TestGetAgent(t *testing.T) {
	agent, err := tServer.GetAgent(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Errorf(err.Error())
	}

	if agent.Uuid != id {
		t.Errorf("incorrect response of GetAgent, expected '%s' found '%s'", id, agent.Uuid)
		t.Fail()
	}
}

func TestGetAgentHealth(t *testing.T) {
	health, err := tServer.GetAgentHealth(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Errorf(err.Error())
	}

	if health.Status != pb.StatusCode_UP {
		t.Errorf("incorrrect response of health, expected '%s' found '%s'", pb.StatusCode_UP, health.Status)
		t.Fail()
	}
}

func TestGetAgentServer(t *testing.T) {
	pair, err := tServer.GetAgentServer(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Errorf(err.Error())
	}

	if pair.Value != id {
		t.Errorf("incorrect response, expected '%s' found '%s'", id, pair.Value)
		t.Fail()
	}
}
