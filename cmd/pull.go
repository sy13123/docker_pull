package cmd

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"go_pull/pkgs/util/logtool"
	"io"
	"go_pull/pkgs/util/makestr"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
	"regexp"
)


var (
	username string
	password string
)


func init() {
	rootCmd.AddCommand(pullcmd)
	pullcmd.PersistentFlags().StringVarP(&username,"user","u","","for docker repository auth username")
	pullcmd.PersistentFlags().StringVarP(&password,"password","p","","for docker repository auth password")

}
//NoArgs - 如果存在任何位置参数，该命令将报错
//ArbitraryArgs - 该命令会接受任何位置参数
//OnlyValidArgs - 如果有任何位置参数不在命令的 ValidArgs 字段中，该命令将报错
//MinimumNArgs(int) - 至少要有 N 个位置参数，否则报错
//MaximumNArgs(int) - 如果位置参数超过 N 个将报错
//ExactArgs(int) - 必须有 N 个位置参数，否则报错
//ExactValidArgs(int) 必须有 N 个位置参数，且都在命令的 ValidArgs 字段中，否则报错
//RangeArgs(min, max) - 如果位置参数的个数不在区间 min 和 max 之中，报错
var pullcmd = &cobra.Command{
	Use:   "pull [images]",
	Short: "Download and import the image",
	Args: cobra.RangeArgs(1, 1),
	Long:  `All software has versions. This is pull's`,
	Run: func(cmd *cobra.Command, args []string) {
		startpull(args)
	},
}

func startpull(img []string) {
	if len(img) > 1{
		logtool.SugLog.Fatal("Parameter parsing exception , plase use \"gohttp help\"")
	}
	imageName := regexp.MustCompile(`^[^@]+`).FindString(img[0])
	if imageName=="" {
		logtool.SugLog.Fatal("image 名称不合法")
	}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	logtool.Fatalerror(err)

	authConfig := types.AuthConfig{
		Username: username,
		Password: password,
	}
	encodedJSON, err := json.Marshal(authConfig)
	logtool.Fatalerror(err)
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{RegistryAuth: authStr})

	logtool.Fatalerror(err)

    d := json.NewDecoder(out)

	defer out.Close()

	type Event struct {
       Status         string `json:"status"`
	   Id				string `json:"id"`
       //Error          string `json:"error"`
       Progress       string `json:"progress"`
       //ProgressDetail struct {
       //    Current float64 `json:"current"`
       //    Total   float64 `json:"total"`
       //} `json:"progressDetail"`
   }


	var event *Event
   	for {
   	   if err := d.Decode(&event); err != nil {
   	       if err == io.EOF {
   	           break
   	       }

   	       logtool.Fatalerror(err)
   	   }	   
	   Id := String_lengthening(event.Id,15)
	   if event.Progress   == ""{
		logtool.SugLog.Infof("%v%v",Id,event.Status)
	   }else{
		logtool.SugLog.Infof("%v%v",Id,event.Progress)
	   }
   	}
	   logtool.SugLog.Infof("%v download complete",imageName)
}

func String_lengthening(v string,c int) string{
	if len(v) > c {
		return v
	}
	l := c-len(v)
	return makestr.Repeat(v," ",l)	
}