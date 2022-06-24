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



func Repeat(v string,f string,l int ) string {
	var b strings.Builder
	b.WriteString(v)
	for i := 0; i <= l; i++{
		b.WriteString(f)
	}
	return b.String()
}



//go test  -bench=. -benchmem