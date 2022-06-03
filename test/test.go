package test

import (
	"fmt"
)
type Member struct{
	Name string
}

func main() {
    m1 := Member{}
    m2 := new(Member)
    Change(m1,m2)
    fmt.Println(m1,m2)
}

func Change(m1 Member,m2 *Member){
    m1.Name = "小明"
    m2.Name = "小红"
}

