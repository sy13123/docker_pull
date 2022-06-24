package request

import (
	"crypto/tls"
	"encoding/json"

	//"fmt"
	//"os"
	"github.com/go-resty/resty/v2"
)

type reqr struct {
	Client  *resty.Client
	Clientr *resty.Request
	resp    *resty.Response
	Url     string
}

func (c *reqr) sethead(q string, a string) *reqr {
	c.Clientr.SetHeader(q, a)
	return c
}

func (c *reqr) Setheads(k map[string]string) *reqr {
	c.Clientr.SetHeaders(k)
	return c
}

func (c *reqr) Notparse() *reqr {
	c.Clientr.SetDoNotParseResponse(true)
	return c
}

func (c *reqr) Settls() *reqr {
	c.Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	return c
}

func (c *reqr) Get() (*resty.Response, error) {
	return c.Clientr.Get(c.Url)
}

func (c *reqr) setresult(v *interface{}) *reqr {
	//var v interface{}
	c.Clientr.SetResult(v)
	return c
}

func (c *reqr) post(i ...string) (*resty.Response, error) {
	return c.Clientr.Post(c.Url)
}

func Requests(url string) *reqr {
	client := resty.New()
	return &reqr{Client: client,
		Clientr: client.R(),
		Url:     url}

	//
	//if err != nil {
	//fmt.Print(err)
	//	return err
	//  //log.Fatal(err)
	//}
	//return resp
	//fmt.Println("Response Info:",AuthSuccess)
	//fmt.Println("Status Code:", resp.StatusCode())
	//fmt.Println("Status:", resp.Status())
	//fmt.Println("Proto:", resp.Proto())
	//fmt.Println("Time:", resp.Time())
	//fmt.Println("Received At:", resp.ReceivedAt())
	//fmt.Println("Size:", resp.Size())
	//fmt.Println("Headers:")
	//for key, value := range resp.Header() {
	//  fmt.Println(key, "=", value)
	//}
	//fmt.Println("Cookies:")
	//for i, cookie := range resp.Cookies() {
	//  fmt.Printf("cookie%d: name:%s value:%s\n", i, cookie.Name, cookie.Value)
	//}
}

func Parsebody_to_json(resp *resty.Response) map[string]interface{} {
	//if err != nil{
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	var v interface{}
	json.Unmarshal(resp.Body(), &v)
	return v.(map[string]interface{})
}

func Ecocde_json(v any) ([]byte, error) {
	e, err := json.Marshal(v)
	return e, err
}

//func tojson(resp *resty.Response,err error) interface{} {
//	if err != nil{
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	fmt.Println(resp.RawBody())
//	var v interface{}
//	json.NewDecoder(resp.RawBody()).Decode(&v)
//	return v
//}
