package cmd

import (
	"go_pull/pkgs/util/logtool"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gohttp",
	Short: "get a image",
	Long:  `get a image!`,
	//	Run: func(cmd *cobra.Command, args []string) {
	//		logtool.InitEvent()
	//	},
}
func init(){
	logtool.InitEvent()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logtool.SugLog.Fatal(err)
	}
}
