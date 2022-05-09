package server

import (
	"context"
	"time"

	"github.com/imhshekhar47/ops-agent/pb"
	"github.com/imhshekhar47/ops-agent/service"
	"github.com/imhshekhar47/ops-agent/task"
	"github.com/imhshekhar47/ops-agent/util"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AgentServer struct {
	pb.UnimplementedOpsAgentServiceServer
	log        *logrus.Entry
	schedulers []task.ScheduledTask

	agentService *service.AgentService
}

func NewAgentServer(
	aLogger *logrus.Logger,
	agentService *service.AgentService,
) *AgentServer {
	return &AgentServer{
		log:          aLogger.WithField("origin", "server:AgentServer"),
		schedulers:   make([]task.ScheduledTask, 0),
		agentService: agentService,
	}
}

func (s *AgentServer) AddScheduledTask(task task.ScheduledTask) {
	s.log.Tracef("entry: AddScheduledTask(%s)", task.Name())
	s.schedulers = append(s.schedulers, task)
	s.log.Traceln("ext: AddScheduledTask()")
}

func (s *AgentServer) RunTasks() {
	s.log.Trace("entry: RunTasks()")
	for _, sched := range s.schedulers {
		err := sched.Run()
		if err != nil {
			s.log.Errorf("failed to run %s", sched.Name())
		}
	}
	s.log.Trace("exit: RunTasks()")
}

func (s *AgentServer) GetAgent(context.Context, *emptypb.Empty) (*pb.Agent, error) {
	defer util.Timer(time.Now(), "OpsAgentServer/GetAgent")
	s.log.Traceln("entry: GetAgent()")
	agent := s.agentService.GetAgent()
	s.log.Tracef("exit: GetAgent() :%s", util.GetHash(agent))
	return agent, nil
}

func (s *AgentServer) GetAgentHealth(context.Context, *emptypb.Empty) (*pb.Health, error) {
	defer util.Timer(time.Now(), "OpsAgentServer/GetAgentHealth")
	s.log.Traceln("entry: GetAgentHealth()")
	a := s.agentService.GetAgent()

	s.log.Traceln("exit: GetAgentHealth()")
	return &pb.Health{
		Timestamp: timestamppb.Now(),
		Status:    a.Status,
	}, nil
}

func (s *AgentServer) GetAgentServer(context.Context, *emptypb.Empty) (*pb.Pair, error) {
	defer util.Timer(time.Now(), "OpsAgentServer/GetAgentServer")
	s.log.Traceln("entry: GetAgentServer()")
	server := s.agentService.GetAgent().Server
	s.log.Traceln("entry: GetAgentServer()")
	return &pb.Pair{
		Key:   server.Uuid,
		Value: server.Hostname,
	}, nil
}
