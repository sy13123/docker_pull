package cmd
import (
	"fmt"
	"github.com/spf13/cobra"
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
		fmt.Println(err)
		os.Exit(1)
	}
}