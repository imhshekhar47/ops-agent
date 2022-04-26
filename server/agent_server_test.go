package server

import (
	"context"
	"testing"

	"github.com/imhshekhar47/ops-agent/config"
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
	tService *service.AgentService = service.NewAgentService(tConfig)
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
