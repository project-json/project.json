package cmd

import (
	"fmt"
	"github.com/projectcfg/projectcfg"
	"github.com/spf13/cobra"
)

var request *Request

type FlagName string

const (
	FilepathFlag FlagName = "file"
)

type FlagChar string

const (
	FilepathChar FlagChar = "f"
)

var RootCmd = &cobra.Command{
	Use:   "projectcfg",
	Short: "Display and manage your Project Config file(s)",
	PreRun: func(cmd *cobra.Command, args []string) {
		filepath := GetPersistentStringFlag(cmd, FilepathFlag)
		projectcfg.Initialize(filepath)
		//projectcfg.Instance.LoadFrameworks()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(projectcfg.Instance.String())
	},
}

func GetPersistentStringFlag(cmd *cobra.Command, name FlagName) string {
	value, _ := cmd.PersistentFlags().GetString(string(name))
	return value
}

func init() {
	request = &Request{}
	RootCmd.PersistentFlags().StringVarP(
		&request.Filepath,
		"file",
		"f",
		projectcfg.DefaultFilename,
		"Project Config filepath",
	)
}
