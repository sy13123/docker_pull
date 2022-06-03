package check_path

import (
	"os"
)

type filec struct{
    filetype os.FileInfo
    err error
}

func Check_path(path string) (*filec) {
    f, err := os.Stat(path)
    return &filec{filetype: f,
                  err:err}
    //if os.IsNotExist(err) { return false, nil }
    //return false, err
}

func (f *filec) Exists() bool {
    if f.err !=nil{
        return false
    }
    return true
}

func (f *filec) Adir() bool {
    if !f.Exists(){
        return false
    }
    return f.filetype.IsDir()
}

func (f *filec) Afile() bool {
    return !f.Adir()
}