package main

import (
	"fmt"
)

type MyInterface1 interface {
	MyMethod1()
}

type MyType1 struct{}

func (m *MyType1) MyMethod1() {
	fmt.Println("Method called on", m)
}

// Вопрос: Что будет выведено при выполнении этого кода?
func exr1() {
	var i MyInterface1
	var t *MyType1 = nil
	i = t
	i.MyMethod1()
}

type MyInterface2 interface {
	MyMethod2()
}

func DoSomething(i MyInterface2) {
	if i == nil {
		fmt.Println("nil interface")
	} else {
		fmt.Println("non-nil interface")
	}
}

// Вопрос: Что будет выведено при выполнении этого кода?
func exr2() {
	var i MyInterface2
	DoSomething(i)

	var t *MyType2 = nil
	i = t
	DoSomething(i)
}

type MyType2 struct{}

func (m *MyType2) MyMethod2() {}

type Reader interface {
	Read() string
}

type Writer interface {
	Write(string)
}

type ReadWriter interface {
	Reader
	Writer
}

type myBuffer struct {
	content string
}

func (b *myBuffer) Read() string {
	return b.content
}

func (b *myBuffer) Write(s string) {
	b.content = s
}

func useInterface(i interface{}) {
	if rw, ok := i.(ReadWriter); ok {
		fmt.Println("ReadWriter with content:", rw.Read())
	} else if r, ok := i.(Reader); ok {
		fmt.Println("Reader with content:", r.Read())
	} else {
		fmt.Println("Unknown type")
	}
}

// Вопрос: Что будет выведено при выполнении этого кода?
func exr3() {
	var b myBuffer
	b.Write("Hello, Interface!")

	var r Reader = &b
	var w Writer = &b
	var rw ReadWriter = &b

	useInterface(r)
	useInterface(w)
	useInterface(rw)
}
