package server

import (
	"context"
	"time"

	"github.com/imhshekhar47/ops-agent/config"
	"github.com/imhshekhar47/ops-agent/pb"
	"github.com/imhshekhar47/ops-agent/service"
	"github.com/imhshekhar47/ops-agent/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AgentServer struct {
	pb.UnimplementedOpsAgentServiceServer
	config *config.AgentConfiguration

	log *logrus.Entry

	service *service.AgentService
}

func NewAgentServer(
	config *config.AgentConfiguration,
	logger *logrus.Logger,
	agentService *service.AgentService,
) *AgentServer {
	return &AgentServer{
		config:  config,
		log:     logger.WithField("origin", "server:AgentServer"),
		service: agentService,
	}
}

func (s *AgentServer) GetAgent(context.Context, *emptypb.Empty) (*pb.Agent, error) {
	defer util.Timer(time.Now(), "OpsAgentServer/GetAgent")
	s.log.Traceln("entry: GetAgent()")
	agent := s.service.GetAgent()
	s.log.Tracef("exit: GetAgent() :%s", util.GetHash(agent))
	return agent, nil
}

func (s *AgentServer) GetAgentHealth(context.Context, *emptypb.Empty) (*pb.Health, error) {
	defer util.Timer(time.Now(), "OpsAgentServer/GetAgentHealth")
	s.log.Traceln("entry: GetAgentHealth()")
	a := s.service.GetAgent()

	s.log.Traceln("exit: GetAgentHealth()")
	return &pb.Health{
		Timestamp: timestamppb.Now(),
		Status:    a.Status,
	}, nil
}

func (s *AgentServer) GetAgentServer(context.Context, *emptypb.Empty) (*pb.Pair, error) {
	defer util.Timer(time.Now(), "OpsAgentServer/GetAgentServer")
	s.log.Traceln("entry: GetAgentServer()")
	server := s.service.GetAgent().Server
	s.log.Traceln("entry: GetAgentServer()")
	return &pb.Pair{
		Key:   server.Uuid,
		Value: server.Hostname,
	}, nil
}
