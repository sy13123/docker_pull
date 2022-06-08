package makestr

import (
	"strings"
)

func Joinstring(strs ...string ) string {
	var b strings.Builder
	for _, v:= range strs{
        b.WriteString(v)
	}
	return b.String()
}


//go test  -bench=. -benchmem