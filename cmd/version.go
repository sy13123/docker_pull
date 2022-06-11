package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of pull",
  Long:  `All software has versions. This is pull's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("go_pull version v1")
  },
}