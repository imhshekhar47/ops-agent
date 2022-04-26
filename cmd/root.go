/*
Copyright © 2022 Himanshu Shekhar <himanshu.kiit@gmail.com>
Code ownership is with Himanshu Shekhar. Use without modifications.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/imhshekhar47/ops-agent/config"
	"github.com/imhshekhar47/ops-agent/pb"
	"github.com/imhshekhar47/ops-agent/server"
	"github.com/imhshekhar47/ops-agent/service"
	"github.com/imhshekhar47/ops-agent/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	argStartGrpcPort     uint16
	argStartRestPort     uint16
	argStartAdminAddress string

	// configurations
	agentConfiguration *config.AgentConfiguration

	// services
	agentService *service.AgentService

	// servers
	agentServer *server.AgentServer

	// clients
	adminServiceClient pb.OpsAdminServiceClient
)

var rootCmd = &cobra.Command{
	Use:   "ops-agent",
	Short: "Agent application",
	Long:  `Agent application run on each of the server and reports to the Admin application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ops-agent.yaml)")
}

func initConfig() {
	util.Logger.WithField("origin", "cmd::root").Traceln("entry: initConfig()")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.SetConfigName(".ops-agent")
		viper.SetConfigType("yaml")

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
	}
	util.Logger.WithField("origin", "cmd::root").Traceln("config yaml", cfgFile)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		util.Logger.WithField("origin", "cmd::root").Debugln("Using config file:", viper.ConfigFileUsed())
	}
	loadServerConfig()

	util.Logger.WithField("origin", "cmd::root").Traceln("exit: initConfig()")
}

func loadServerConfig() {
	util.Logger.WithField("origin", "cmd::root").Traceln("entry: loadServerConfig()")
	hostname := viper.GetString("server.hostname")
	if hostname == "" {
		hostname = util.GetHostname()
	}

	uuid := util.NonEmptyOrDefult(os.Getenv("OPS_AGENT_ID"), util.Encode(hostname))

	coreConiguration := config.CoreConfiguration{
		Version: viper.GetString("core.version"),
	}

	agentConfiguration = &config.AgentConfiguration{
		Core:     coreConiguration,
		Hostname: hostname,
		Uuid:     uuid,
		Address:  fmt.Sprintf("%s:%d", hostname, argStartGrpcPort),
	}

	//util.Logger.WithField("origin", "cmd::root").Debugln("agent_configuration", agentConfiguration)
	util.Logger.WithField("origin", "cmd::root").Traceln("exit: loadServerConfig()")
}
