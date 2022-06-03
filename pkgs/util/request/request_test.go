package request

import (
	"fmt"
	//"go_pull/pkgs/util/iowrite"
	"strconv"
	"testing"
	 "encoding/json"

)

func Test_request(t *testing.T) {
	tests := []struct {
		name string
	}{
		//{name:"asd"},
		{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//url:="http://jsonplaceholder.typicode.com/posts/2"
			//url := "https://www.baidu.com"
			url := "http://wh.mapall.cn:41810/index"

			aa, err := Requests(url).Settls().
				//notparse().
				sethead("token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkIjoid2Vkb3JhIn19.3-0GxVTviaGCRisQKs1wYWcovpUKH8roCiub46oRrCU").
				Get()
			//aa,err := requests("https://www.baidu.com").get()
			//bb :=aa.settls()
			if err != nil {
				fmt.Println(err)
			}
			//f1 := iowrite.Uflie("/tmp/bao.tar")
			//f1.Readio_to_file(aa.RawBody())
			//reader:=bufio.NewReader(aa.RawBody())
			//body, err :=reader.re()
			fmt.Println(aa)
			var bb map[string]interface{}
			
			json.Unmarshal(aa.Body(),&bb)
			fmt.Println(bb)




			b, _ := strconv.Atoi(aa.Header()["Content-Length"][0])
			fmt.Println(b)
			//reader := bufio.NewReader(aa.RawBody())
			//buf := make([]byte,4*1024)
			//for {
			//	n, err := reader.Read(buf)
			//	f1.bufWriter.Write(buf[:n])
			//	//line, err := reader.ReadBytes('\n')
			//	if err != nil {
			//		fmt.Println(err,"err")
			//		break
			//	}
			//}
			//f1.close()
			//fmt.Println(aa.Header())
			//json := Parsebody_to_json(aa,err)
			//fmt.Println(json["id"])
			//fmt.Printf("%T\n",aa)
		})
	}
}
