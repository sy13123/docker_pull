package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestContentvar(t *testing.T) {
	tests := []struct {
		name string
		want []m1
	}{
		{},// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			 got := Contentvar() 
			 got[0].Config="666"

			 fmt.Println(got)
			b,a :=json.Marshal(got)
			fmt.Println(string(b),a)
		})
	}
}
