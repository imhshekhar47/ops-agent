package task

import (
	"context"

	"github.com/imhshekhar47/ops-agent/pb"
	"github.com/procyon-projects/chrono"
	"github.com/sirupsen/logrus"
)

type PingAdminScheduledTask struct {
	log      *logrus.Entry
	task     *chrono.ScheduledTask
	cronExpr string

	agent       *pb.Agent
	adminClient *pb.OpsAdminServiceClient
}

func NewPingAdminScheduledTask(expr string, admin *pb.OpsAdminServiceClient, anAgent *pb.Agent, aLogger *logrus.Logger) *PingAdminScheduledTask {
	return &PingAdminScheduledTask{
		log:         aLogger.WithField("origin", "task/PingAdminScheduledTask"),
		cronExpr:    expr,
		agent:       anAgent,
		adminClient: admin,
	}
}

func (s *PingAdminScheduledTask) Name() string {
	return "PingAdminScheduledTask"
}

func (s *PingAdminScheduledTask) Run() error {
	s.log.Traceln("entry: Run()")

	scheduledTask, err := TaskScheduler.ScheduleWithCron(func(ctx context.Context) {
		s.log.Traceln("running Task")
		// Register with the admin
		_, adminErr := (*s.adminClient).Register(context.Background(), &pb.Agent{
			Meta:    s.agent.Meta,
			Uuid:    s.agent.Uuid,
			Address: s.agent.Address,
			Status:  s.agent.Status,
			Info:    s.agent.Info,
		})
		if adminErr != nil {
			s.log.Errorf(adminErr.Error())
		}

	}, s.cronExpr)

	if err != nil {
		return err
	}

	s.task = &scheduledTask
	s.log.Traceln("exit: Run()")
	return nil
}

func (s *PingAdminScheduledTask) Cancel() {
	s.log.Traceln("entry: Cancel()")
	if s.task != nil {
		(*s.task).Cancel()
	}
	s.log.Traceln("exit: Cancel()")
}
