package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go_pull/pkgs/model"
	"go_pull/pkgs/util/aes"
	"go_pull/pkgs/util/check_path"
	"go_pull/pkgs/util/iowrite"
	"go_pull/pkgs/util/logtool"
	"go_pull/pkgs/util/makestr"
	"go_pull/pkgs/util/request"
	"go_pull/pkgs/util/tartool"
	"go_pull/pkgs/util/timetool"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var (
	auth_url    string
	reg_service string
	repository  string
	auth_head   map[string]string
	resp        *resty.Response
	err         error
	registry string
	platform string
	plist bool
)

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.PersistentFlags().StringVarP(&platform,"platform","p","amd64","Select platform system architecture")
	downloadCmd.PersistentFlags().BoolVarP(&plist, "list", "l", false, "list platform system architecture")
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download image only",
	Args: cobra.MinimumNArgs(1),
	Long:  `All software has versions. This is pull's`,
	Run: func(cmd *cobra.Command, args []string) {
		startdownload(args)
	},
}

func startdownload(args []string) {

	// Look for the Docker image to download
	inargv := args
	var img string
	var repo string= "library"
	var tag string = "latest"
	var imlist string 
	var digest string
 	var imgpartstr string
	registry = "registry-1.docker.io"

	if strings.Contains(inargv[0], "@") {
		s := strings.Split(inargv[0], "@")
		imlist, digest = s[0], s[1]
	} else {
		imlist, digest = inargv[0],""
	}
	
	if strings.Contains(inargv[0],":"){
		s := strings.Split(imlist, ":")
		imgpartstr, tag = s[0],s[1]
	}else{
		imgpartstr, tag =imlist,"latest"
	}

	imgpartlist := strings.Split(imgpartstr, "/")
	img = imgpartlist[len(imgpartlist)-1]


	// Docker client doesn't seem to consider the first element as a potential registry unless there is a '.' or ':'
	if len(imgpartlist) > 1 && (strings.Contains(imgpartlist[0], ".") || strings.Contains(imgpartlist[0], ":")) {
		registry = imgpartlist[0]
		repo = strings.Join(imgpartlist[1:len(imgpartlist)-1], "/")
	} else {
		if len(imgpartlist[:len(imgpartlist)-1])!= 0{
			repo = strings.Join(imgpartlist[:len(imgpartlist)-1], "/")
		}
	}
	repository = makestr.Joinstring(repo, "/", img)

	//Get Docker authentication endpoint when it is required
	auth_url = "https://auth.docker.io/token"
	reg_service = "registry.docker.io"
	resp, err = request.Requests(
		makestr.Joinstring("https://", registry, "/v2/")).
		Settls().
		Get()
	logtool.Fatalerror(err)
	if resp.StatusCode() == 401 {
		auth_url = resp.Header()["Www-Authenticate"][0]
		reg_Header_list := strings.Split(auth_url, "\"")
		auth_url = reg_Header_list[1]
		if len(reg_Header_list) > 4 {
			reg_service = reg_Header_list[3]
		} else {
			reg_service = ""
		}
	}
	//Fetch manifest v2 and get image layer digests

	var real_tag string
	if digest != ""{
		real_tag=digest
	}else{
		real_tag=tag
	}
	auth_head = get_auth_head("application/vnd.docker.distribution.manifest.list.v2+json")
	resp, err := request.Requests(makestr.Joinstring("https://", registry, "/v2/", repository, "/manifests/", real_tag)).
		Setheads(auth_head).
		Settls().
		Get()
		logtool.Errorerror(err)
	if resp.StatusCode() != 200 {
		logtool.SugLog.Fatal("[-] Cannot fetch manifest for %v [HTTP %v]", repository, resp.Status())
	}


	resp_json :=request.Parsebody_to_json(resp)

	var platformv_list []string
	var platform_digest string

	if plist{
		data, _ := json.MarshalIndent(resp_json["manifests"],""," ")
		fmt.Println(string(data))
		os.Exit(0)
	}
	for _,v := range resp_json["manifests"].([]interface{} ) {

		platformv := v.(map[string]interface{})["platform"]
		platformv = platformv.(map[string]interface{})["architecture"].(string)
		if platformv == platform{
			platform_digest = v.(map[string]interface{})["digest"].(string)
			break
		}
		platformv_list = append(platformv_list, platformv.(string))


	}

	if platform_digest == "" {
		logtool.SugLog.Fatalf("please use -p %v\n",platformv_list )
	}


	auth_head = get_auth_head("application/vnd.docker.distribution.manifest.v2+json")
	resp, err = request.Requests(makestr.Joinstring("https://", registry, "/v2/", repository, "/manifests/", platform_digest)).
		Setheads(auth_head).
		Settls().
		Get()
	logtool.Errorerror(err)


	rresp := request.Parsebody_to_json(resp)
	layers := rresp["layers"].([]interface{})

	//Create tmp folder that will hold the image
	imgdir := makestr.Joinstring("tmp_", img, "_", strings.ReplaceAll(tag, "@", ""))
	if check_path.Check_path(imgdir).Exists() {
		os.RemoveAll(imgdir)
	}
	os.Mkdir(imgdir, os.ModePerm)
	logtool.SugLog.Infof("Creating image structure in: %v", imgdir)

	config := rresp["config"].(map[string]interface{})["digest"].(string)
	confresp, err := request.Requests(
		makestr.Joinstring("https://", registry, "/v2/", repository, "/blobs/", config)).
		Setheads(auth_head).
		Settls().
		Get()
	logtool.Fatalerror(err)
	f := iowrite.Uflie(makestr.Joinstring(imgdir, "/", config[7:], ".json"))
	f.BufWriter.WriteString(confresp.String())
	f.Close()

	content := model.Contentvar()
	content[0].Config = makestr.Joinstring(config[7:], ".json")

	if len(imgpartlist[:len(imgpartlist)-1]) != 0 {
		content[0].RepoTags = append(
			content[0].RepoTags,
			makestr.Joinstring(strings.Join(imgpartlist[:len(imgpartlist)-1], "/"), "/", img, ":", tag),
		)
	} else {
		content[0].RepoTags = append(
			content[0].RepoTags,
			makestr.Joinstring(img, ":", tag),
		)
	}

	//Build layer folders
	var wg sync.WaitGroup
	wg = sync.WaitGroup{}
	wg.Add(len(layers))

	var parentid string
	var last_fake_layerid string
	for x, layer := range layers {
		ublob := layer.(map[string]interface{})["digest"].(string)
		logtool.SugLog.Info(ublob)
		fake_layerid := aes.Sha256t(makestr.Joinstring(parentid, "\n", ublob, "\n"))
		layerdir := makestr.Joinstring(imgdir, "/", fake_layerid)
		os.Mkdir(layerdir, os.ModePerm)
		go Download_img(layer,layerdir,ublob,&wg)
	
		content[0].Layers = append(content[0].Layers, makestr.Joinstring(fake_layerid, "/layer_gzip.tar"))
		//Creating json file
		f2 := iowrite.Uflie(makestr.Joinstring(layerdir, "/json"))
		//last layer = config manifest - history - rootfs
		var json_obj map[string]interface{}
		if layers[len(layers)-1].(map[string]interface{})["digest"].(string) ==
			layer.(map[string]interface{})["digest"].(string) {
			json_obj = request.Parsebody_to_json(confresp)
			delete(json_obj, "history")
			if _, ok := json_obj["rootfs"]; ok {
				//存在
				delete(json_obj, "rootfs")
			} else if _, ok := json_obj["rootfS"]; ok {
				delete(json_obj, "rootfS")
			}
		} else {
			json_obj = model.Empty_config()
		}
		json_obj["id"] = fake_layerid

		if parentid != "" {
			json_obj["parent"] = parentid
		}
		parentid = fake_layerid
		data, _ := json.Marshal(json_obj)
		f2.BufWriter.Write(data)
		f2.Close()

		if x == len(layers){
			last_fake_layerid = fake_layerid
		}

	}
	wg.Wait()

	f3 := iowrite.Uflie(makestr.Joinstring(imgdir, "/manifest.json"))
	data, _ := json.Marshal(content)
	f3.BufWriter.Write(data)
	f3.Close()


	var content1 map[string](map[string]string)
	if len(imgpartlist[:len(imgpartlist)-1]) !=0{

		content1 =  map[string](map[string]string){
			makestr.Joinstring( strings.Join(imgpartlist[:len(imgpartlist)-1], "/") ,img) : map[string]string{tag: last_fake_layerid },
		}
	}else{
		content1 =  map[string](map[string]string){
			img: map[string]string{tag: last_fake_layerid},
		}

	}
	
	f5 := iowrite.Uflie(makestr.Joinstring(imgdir, "/repositories"))
	data1, _ := json.Marshal(content1)
	f5.BufWriter.Write(data1)
	f5.Close()	

	//Create image tar and clean tmp folder
	//docker_tar:= makestr.Joinstring(
	//strings.ReplaceAll(repo,"/", "_"),
	//"_",img,".tar")
	fmt.Print("Creating archive...")
	os.Stdout.Sync()

	if check_path.Check_path(img+".tar").Exists() {
		os.Remove(img+".tar")
	}
	tartool.TarGz(img+".tar", imgdir)
	os.RemoveAll(imgdir)
	fmt.Printf("打包完成，生成文件 %v\n",img+".tar" )
}



func Download_img(layer interface{},layerdir string,ublob string,w *sync.WaitGroup){
	
	//Creating VERSION file
	f := iowrite.Uflie(makestr.Joinstring(layerdir, "/VERSION"))
	f.BufWriter.WriteString("1.0")
	f.Close()

	// Creating layer.tar file
	logtool.SugLog.Infof("%v%v", ublob[7:19], ": Downloading...")
	os.Stdout.Sync()
	auth_head := get_auth_head("application/vnd.docker.distribution.manifest.v2+json",auth_head)
	bresp, err := request.Requests(
		makestr.Joinstring("https://", registry, "/v2/", repository, "/blobs/", ublob)).
		Notparse().
		Setheads(auth_head).
		Settls().
		Get()

	logtool.Fatalerror(err)
	if bresp.StatusCode() != 200 {
		logtool.SugLog.Fatal(layer.(map[string]interface{}))
		bresp, _ := request.Requests(layer.(map[string]interface{})["urls"].([]string)[0]).
			Setheads(auth_head).
			Settls().
			Get()
		if bresp.StatusCode() != 200 {
			fmt.Printf("\rERROR: Cannot download layer %v [HTTP %v %v]", ublob[7:19], bresp.StatusCode(), bresp.Header()["Content-Length"])
			logtool.SugLog.Info(bresp)
			os.Exit(1)
		}

	} else if bresp.StatusCode() == 200 {
		goto statusok
	} else {
		logtool.SugLog.Info("bad request")
	}
	statusok:
		//Stream download and follow the progress
		unit, _ := strconv.Atoi(bresp.Header()["Content-Length"][0])
		unit = unit / 50
		acc := 0
		nb_traits := 0
		progress_bar(ublob, nb_traits)
		f1 := iowrite.Uflie(makestr.Joinstring(layerdir, "/layer_gzip.tar"))
		buf := make([]byte, 8192)
		reader := bufio.NewReader(bresp.RawBody())

		for {
			n, err := reader.Read(buf)
			f1.BufWriter.Write(buf[:n])
			acc = acc + n
			if acc > unit {
				nb_traits = nb_traits + 1
				progress_bar(ublob, nb_traits)
				acc = 0
			}
			
			//line, err := reader.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					fmt.Println("")
					//logtool.SugLog.Info("ioFinish")
				} else {
					logtool.SugLog.Fatal(err, "ioerr")
				}
				break
			}
		}
		f1.Close()
		fmt.Printf("%v: Extracting...%v", ublob[7:19], strings.Repeat(" ", 50))
		os.Stdout.Sync()

		fmt.Printf("%v: Pull complete [%v]\n",
			ublob[7:19], bresp.Header()["Content-Length"])
		(*w).Done()
}



// Get Docker token (this function is useless for unauthenticated registries like Microsoft)
func get_auth_head(qtype string,a ...any ) map[string]string {
	if len(a) != 0  {
			t := a[0].(map[string]string)
			tm:= t["expires_in"]
			st:= timetool.Strtorime(tm,"UTC")
			if st.Add(-2*time.Second).After(time.Now().UTC()){
				t["Accept"]=qtype
				return t
			}

	}
	resp, err := request.Requests(
		makestr.Joinstring(auth_url, "?service=", reg_service, "&scope=repository:", repository, ":pull")).
		Settls().
		Get()
	logtool.Fatalerror(err)
	resp_json := request.Parsebody_to_json(resp)
	expires_in:= int(resp_json["expires_in"].(float64))
	issued_at := resp_json["issued_at"].(string)

	expires_time := timetool.Strtorime(issued_at,"UTC").
					Add(time.Duration(expires_in) * (time.Hour)).
					Format("2006-01-02 15:04:05")

	auth_head := map[string]string{"Authorization": makestr.Joinstring("Bearer ", resp_json["token"].(string)), 
								   "Accept": qtype,
								   "expires_in" : expires_time,
								   }
	return auth_head
	
}

//Docker style progress bar
func progress_bar(ublob string, nb_traits int) {
	fmt.Print(makestr.Joinstring("", ublob[7:19], ": Downloading ["))
	for i := 0; i < nb_traits; i++ {
		if i == nb_traits-1 {
			fmt.Print(">")
		} else {
			fmt.Print("=")
		}
	}
	for i := 0; i < 49-nb_traits; i++ {
		fmt.Print(" ")
	}
	fmt.Print("]\n")
	os.Stdout.Sync()
}
