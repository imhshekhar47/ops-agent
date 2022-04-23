/*
Copyright © 2022 Himanshu Shekhar <himanshu.kiit@gmail.com>
Code ownership is with Himanshu Shekhar. Use without modifications.
*/
package cmd

import (
	"github.com/imhshekhar47/ops-agent/util"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "agent status",
	Long:  `Shows agent application status`,
	Run:   runStatusCmd,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func runStatusCmd(cmd *cobra.Command, args []string) {
	util.Logger.Traceln("entry: runStatusCmd")
	util.Logger.Traceln("ext: runStatusCmd")
}
