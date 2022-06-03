package cmd

import (
	"fmt"
	"testing"
	"strconv"
	"time"
)

func Test_request(t *testing.T) {
	tests := []struct {
		name string
	}{
		//{name:"asd"},
		{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aa := requests().get("http://www.baidu.com", map[string]interface{}{
				"page_no": "1",
				"limit":   "20",
				"sort":    "name",
				"order":   "asc",
				"random":  strconv.FormatInt(time.Now().Unix(), 10),
			})
			fmt.Println(aa)
		})
	}
}
