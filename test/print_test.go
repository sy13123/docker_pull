package test

// 向标准输出写入内容
import (
	"testing"
	"fmt"
	"os"
)

type aa struct {
    bb string

}
func Test_main(t *testing.T) {
	fmt.Fprint(os.Stdout, "\r向标准输出写入内容1")
	fmt.Printf("\r向标准输出写入内容2")
	a := &aa{
		bb: "666",
	}
	fmt.Println(a)

	//fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//if err != nil {
	//	fmt.Println("打开文件出错，err:", err)
	//	return
	//}
	//name := "沙河小王子"
	//// 向打开的文件句柄中写入内容
	//fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
}