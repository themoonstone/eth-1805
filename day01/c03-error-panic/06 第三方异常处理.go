package main

import (
	"fmt"
	"github.com/chai2010/errors"
)

var (
	e0 = errors.New("err:init0")
	e1 error
	e2 error
)

func init() {
	e1 = errors.New("err:init1")
}

func init() {
	e2 = errors.New("err:init2")
}

func main() {
	//var e3 = errors.New("err:main3")
	//var e4 = func() error {
	//	return errors.New("err:main4")
	//}()
	var e5 = errors.Wrap(e1, "err:main5")
	//var e6 = errors.Wrap(e5, "err:main6")
	//var e7 = errors.Wrap(e6, "err:main7")

	//fmt.Println(e0, e0.(errors.Error).Caller())
	//fmt.Println(e1, e1.(errors.Error).Caller())
	//fmt.Println(e2, e2.(errors.Error).Caller())
	//fmt.Println(e3, e3.(errors.Error).Caller())
	//fmt.Println(e4, e4.(errors.Error).Caller())
	fmt.Println(e5, e5.(errors.Error).Caller())
	//fmt.Println(e6, e6.(errors.Error).Caller())
	//
	//for i, e := range e7.(errors.Error).Wraped() {
	//	fmt.Printf("err7: wraped(%d): %v\n", i, e)
	//}
	//
	//fmt.Println("err7:", e7.(fmt.Stringer).String())
}