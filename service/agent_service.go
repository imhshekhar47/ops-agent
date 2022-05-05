package service

import (
	"github.com/imhshekhar47/ops-agent/pb"
)

type AgentService struct {
	agent *pb.Agent
}

func NewAgentService(agent *pb.Agent) *AgentService {
	return &AgentService{
		agent: agent,
	}
}

func (s *AgentService) GetAgent() *pb.Agent {
	return s.agent
}

func (s *AgentService) GetAgentServer() *pb.Server {
	return s.agent.Server
}
