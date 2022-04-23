/*
Copyright © 2022 Himanshu Shekhar <himanshu.kiit@gmail.com>
Code ownership is with Himanshu Shekhar. Use without modifications.
*/
package cmd

import (
	"fmt"
	"net"

	"github.com/imhshekhar47/ops-agent/pb"
	"github.com/imhshekhar47/ops-agent/server"
	"github.com/imhshekhar47/ops-agent/service"
	"github.com/imhshekhar47/ops-agent/util"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start agent",
	Long:  `Start the agent application.`,
	Run:   runStartCmd,
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().Uint16VarP(&argStartGrpcPort, "grpc-port", "g", 5702, "gRPC api port")
	startCmd.Flags().Uint16VarP(&argStartRestPort, "rest-port", "r", 8082, "[optional] Rest api port. Specify if you need to expose the agent APIs")
}

func runGrpc(
	listener net.Listener,
	agentServer *server.AgentServer,
) error {

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterOpsAgentServiceServer(grpcServer, agentServer)
	util.Logger.Debugln("Launching grpc on ", listener.Addr())
	return grpcServer.Serve(listener)
}

func runStartCmd(cmd *cobra.Command, args []string) {
	util.Logger.Traceln("entry: runStartCmd()")

	// services
	agentService = service.NewAgentService(agentConfiguration)
	// servers
	agentServer = server.NewAgentServer(
		util.GetLogger("server::AgentServer"),
		agentService,
	)

	//grpc
	grpcListner, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", argStartGrpcPort))
	if err != nil {
		util.Logger.Errorln("error", err)
	}
	runGrpc(grpcListner, agentServer)

	util.Logger.Traceln("exit: runStartCmd()")
}