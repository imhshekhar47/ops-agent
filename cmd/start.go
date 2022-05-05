/*
Copyright © 2022 Himanshu Shekhar <himanshu.kiit@gmail.com>
Code ownership is with Himanshu Shekhar. Use without modifications.
*/
package cmd

import (
	"context"
	"fmt"
	"net"

	"github.com/imhshekhar47/ops-agent/pb"
	"github.com/imhshekhar47/ops-agent/server"
	"github.com/imhshekhar47/ops-agent/service"
	"github.com/imhshekhar47/ops-agent/util"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	startCmd.Flags().StringVarP(&argStartAdminAddress, "admin-address", "a", "", "[optional] Address of admin server")
}

func runGrpc(
	listener net.Listener,
	aAgentServer *server.AgentServer,
) error {

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterOpsAgentServiceServer(grpcServer, aAgentServer)
	util.Logger.Debugln("Launching grpc on ", listener.Addr())
	return grpcServer.Serve(listener)
}

func runStartCmd(cmd *cobra.Command, args []string) {
	util.Logger.Traceln("entry: runStartCmd()")

	var agent *pb.Agent = &pb.Agent{
		Meta: &pb.Metadata{
			Timestamp: timestamppb.Now(),
			Version:   agentConfiguration.Core.Version,
		},
		Uuid:    agentConfiguration.Uuid,
		Address: agentConfiguration.Address,
		Status:  pb.StatusCode_UP,
		Info: &pb.HierarchyInfo{
			Group:       agentConfiguration.Group,
			Component:   agentConfiguration.Component,
			Environment: agentConfiguration.Environment,
			Site:        agentConfiguration.Site,
			Location: &pb.GeoLocation{
				Latitude:  agentConfiguration.Location.Latitude,
				Longitude: agentConfiguration.Location.Longitude,
			},
		},
		Server: &pb.Server{
			AgentId:   agentConfiguration.Uuid,
			Uuid:      agentConfiguration.Uuid,
			Hostname:  agentConfiguration.Hostname,
			Status:    pb.StatusCode_UP,
			Databases: make([]*pb.Database, 0),
		},
	}

	// services
	agentService = service.NewAgentService(agent)

	// servers
	agentServer = server.NewAgentServer(
		agentConfiguration,
		util.Logger,
		agentService,
	)

	if len(argStartAdminAddress) > 11 {

		util.Logger.Debugf("Connecting to admin server %s", argStartAdminAddress)
		adminConn, err := grpc.Dial(argStartAdminAddress, grpc.WithInsecure())
		if err != nil {
			panic("Could not connect to admin server")
		}
		//defer adminConn.Close()

		ctx := context.Background()
		adminServiceClient = pb.NewOpsAdminServiceClient(adminConn)
		adminServiceClient.Register(ctx, &pb.Agent{
			Meta: &pb.Metadata{
				Timestamp: timestamppb.Now(),
				Version:   agentConfiguration.Core.Version,
			},
			Uuid:    agent.Uuid,
			Address: agent.Address,
		})
	}

	//grpc
	grpcListner, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", argStartGrpcPort))
	if err != nil {
		util.Logger.Errorln("could not create tcp connection", err)
	} else {
		err = runGrpc(grpcListner, agentServer)
		if err != nil {
			util.Logger.Errorln("could not launch grpc server", err)
		}
	}

	util.Logger.Traceln("exit: runStartCmd()")
}
