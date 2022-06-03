package makestr

import (
	"fmt"
	"testing"
)

//func Test_main(t *testing.T) {
//	tests := []struct {
//		name string
//	}{
//		{'wer'}	// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			main1()
//		})
//	}
//}
//
//func Test_main(t *testing.T) {
//	a := joinstring("asda","asdasd")
//	fmt.Println(a)
//}
func Benchmark_Add(b *testing.B) {
	a := joinstring("asda","asdasd")
	fmt.Println(a)
}

