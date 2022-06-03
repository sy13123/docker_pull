package model



type Container_config struct {
	Hostname     string
	Domainname   string
	User         string
	AttachStdin  bool
	AttachStdout bool
	AttachStderr bool
	Tty          bool
	OpenStdin    bool
	StdinOnce    bool
	Env          string
	Cmd          string
	Image        string
	Volumes      string
	WorkingDir   string
	Entrypoint   string
	OnBuild      string
	Labels       string
}

//type Empty struct {
//	Created          string            `json:"created"`
//	Container_config Container_config `json:"container_config"`
//}
func Empty_config() map[string]interface{} {
	Container_config_json := Container_config{
		Hostname:     "",
		Domainname:   "",
		User:         "",
		AttachStdin:  false,
		AttachStdout: false,
		AttachStderr: false,
		Tty:          false,
		OpenStdin:    false,
		StdinOnce:    false,
		//		Env:          nil,
		//		Cmd:          nil,
		//		Image:        nil,
		//		Volumes:      nil,
		//		WorkingDir:   nil,
		//		Entrypoint:   nil,
		//		OnBuild:      nil,
		//		Labels:       nil,
	}	

	out := map[string]interface{}{
		"Created" : "1970-01-01T00:00:00Z",
		"Container_config": Container_config_json,
	}

	return out
}