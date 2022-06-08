package cmd
import (
	"github.com/spf13/cobra"
	"go_pull/pkgs/util/logtool"

	"os"
)
var rootCmd = &cobra.Command{
	Use:   "gohttp",
	Short: "pull a image",
	Long: `pull a image!`,
//	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("hello chenqionghe1")
//	},
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logtool.Fatalerror(err)
		os.Exit(1)
	}
}