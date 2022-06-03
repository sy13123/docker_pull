package test

// 向标准输出写入内容
import (
	"testing"
	"fmt"
	"os"
)
func Test_path(t *testing.T) {
	
}


func exists(path string) (bool, error) {
    a, err := os.Stat(path)
	a.IsDir()
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}