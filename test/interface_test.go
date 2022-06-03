package test

import (
    "testing"
    "fmt"
)

type Phone interface {
    call()
}

type NokiaPhone struct {
    b string
}

func (nokiaPhone *NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

func return_f() Phone {
    var a *NokiaPhone
    fmt.Println(a)

    return a
}

func  Test_interface(t *testing.T) {
    b:=return_f()
    b.call()
}
