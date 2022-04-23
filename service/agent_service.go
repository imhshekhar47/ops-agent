package service

import (
	"github.com/imhshekhar47/ops-agent/config"
	"github.com/imhshekhar47/ops-agent/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AgentService struct {
	agentConfig *config.AgentConfiguration

	agent *pb.Agent
}

func NewAgentService(config *config.AgentConfiguration) *AgentService {
	return &AgentService{
		agentConfig: config,
		agent: &pb.Agent{
			Meta: &pb.Metadata{
				Timestamp: timestamppb.Now(),
				Version:   config.Core.Version,
			},
			Uuid:    config.Uuid,
			Address: config.Address,
		},
	}
}

func (s *AgentService) Get() *pb.Agent {
	return s.agent
}

func (s *AgentService) GetHealth() *pb.Health {
	return &pb.Health{
		Timestamp: timestamppb.Now(),
		Status:    pb.StatusCode_UP,
	}
}
