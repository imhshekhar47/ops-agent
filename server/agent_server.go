package server

import (
	"context"
	"time"

	"github.com/imhshekhar47/ops-agent/pb"
	"github.com/imhshekhar47/ops-agent/service"
	"github.com/imhshekhar47/ops-agent/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AgentServer struct {
	pb.UnimplementedOpsAgentServiceServer
	log *logrus.Entry

	service *service.AgentService
}

func NewAgentServer(
	logger *logrus.Entry,
	agentService *service.AgentService,
) *AgentServer {
	return &AgentServer{
		log:     logger,
		service: agentService,
	}
}

func (s *AgentServer) GetAgent(context.Context, *emptypb.Empty) (*pb.Agent, error) {
	s.log.Traceln("entry: GetAgent()")
	s.log.Traceln("exit: GetAgent()")
	defer util.Timer(time.Now(), "server::AgentServer::GetAgent")
	return s.service.Get(), nil
}

func (s *AgentServer) GetAgentHealth(context.Context, *emptypb.Empty) (*pb.Health, error) {
	s.log.Traceln("entry: GetAgentHealth()")
	s.log.Traceln("exit: GetAgentHealth()")
	defer util.Timer(time.Now(), "server::AgentServer::GetAgentHealth")
	return s.service.GetHealth(), nil
}
