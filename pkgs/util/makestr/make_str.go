package makestr

import (
	"fmt"
	"strings"
)

func Joinstring(strs ...string ) string {
	var b strings.Builder
	fmt.Println(strs)
	for _, v:= range strs{
        b.WriteString(v)
	}
	return b.String()
}


//go test  -bench=. -benchmem