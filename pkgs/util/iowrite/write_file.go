package iowrite

import (
	"bufio"
	"io"
	"fmt"
	"log"
	"os"
	//"strconv"
	//"time"
)
type usefile struct{
	BufWriter *bufio.Writer
	bufReader *bufio.Reader
	dstFile  *os.File
}

func Uflie(filepath string) *usefile{
	dstFile, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file failed, err:%v", err)
	}
	bufWriter := bufio.NewWriter(dstFile)
	//defer func() {
	//	bufWriter.Flush()
	//	dstFile.Close()
	//	fmt.Println("文件写入耗时：", time.Now().Sub(st).Seconds(), "s")
	//}()
	return &usefile{
		BufWriter: bufWriter,
		dstFile: dstFile}
}

func (f *usefile) Close() {
	f.BufWriter.Flush()
	f.dstFile.Close()
}

func (f *usefile) Readio_to_file(r io.ReadCloser) {
	reader:=bufio.NewReader(r)
	buf := make([]byte,4*1024)
	for {
		n, err := reader.Read(buf)
		f.BufWriter.Write(buf[:n])
		//line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF{
				fmt.Println("写完了")
            }
			fmt.Println(err,"ioerr")
			break
		}
	}
	f.Close()
}