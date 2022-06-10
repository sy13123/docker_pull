package tartool

import (
  "fmt"
  "os"
  "io"
  "log"
  "strings"
  "archive/tar"
  "compress/gzip"
)

func handleError( _e error ) {
  if _e != nil {
    log.Fatal( _e )
  }
}

func TarGzWrite( _dpath,_spath string, tw *tar.Writer, fi os.FileInfo ) {
  fr, err := os.Open( _dpath+"/"+_spath )
  handleError( err )
  defer fr.Close()

  h := new( tar.Header )
  //_p := strings.Split(_dpath,"/")
  //h.Name =_p[len(_p)-1]+"/"+_spath
  h.Name =_spath
  h.Size = fi.Size()
  h.Mode = int64( fi.Mode() )
  h.ModTime = fi.ModTime()
  err = tw.WriteHeader( h )
  handleError( err )

  _, err = io.Copy( tw, fr )
  handleError( err )
}

func IterDirectory( dirPath,subpath string, tw *tar.Writer ) {
  dir, err := os.Open( dirPath +"/"+subpath )
  handleError( err )
  defer dir.Close()
  fis, err := dir.Readdir( 0 )
  handleError( err )
  for _, fi := range fis {
    var curpath string
    if subpath == ""{
        curpath =  fi.Name()
    }else{
        curpath = subpath + "/" + fi.Name()
    }

    if fi.IsDir() {
      //TarGzWrite( curPath, tw, fi )
      IterDirectory( dirPath,curpath, tw )
    } else {
      fmt.Printf( "adding... %s\n", dirPath+"/"+curpath )
      TarGzWrite( dirPath,curpath, tw, fi )
    }
  }
}

func TarGz( outFilePath string, inPath string) {
  inPath = strings.TrimRight( inPath, "/" ) 
  // file write
  fw, err := os.Create( outFilePath )
  handleError( err )
  defer fw.Close()

  // gzip write
  gw := gzip.NewWriter(fw)
  defer gw.Close()

  // tar write
  tw := tar.NewWriter( gw )
  defer tw.Close()

  IterDirectory( inPath,"", tw )
}

