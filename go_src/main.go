package main

/*
// Note the -std=gnu99. Using -std=c99 will not work.
#cgo CFLAGS: -std=gnu99
#include <stdint.h>
*/
import "C"
import (
	"./go_test_cpp"
	"fmt"
	"time"
	"unsafe"
)

type my_callback struct {
	go_test_cpp.ICallback
}

func (this my_callback) Notify(arg2 string) {
	fmt.Printf("c++ Notify:%s\n", arg2)
}

func main() {

	cb := go_test_cpp.NewDirectorICallback(my_callback{})

	test := go_test_cpp.TestCallCreate()
	test.SetCallback(cb)

	res_ptr := test.Test("Hello World!").Swigcptr()
	res := *(*int32)(unsafe.Pointer(res_ptr))
	if res == 0 {
		fmt.Println("Test success!")
	} else {
		fmt.Println("init failed[", res, "]!")
	}
	go_test_cpp.Swig_free(res_ptr)

	var a, b uint32
	a = 1
	b = 2
	res_ptr = test.Add((go_test_cpp.SwigcptrUint32_t)(unsafe.Pointer(&a)), (go_test_cpp.SwigcptrUint32_t)(unsafe.Pointer(&b))).Swigcptr()
	res = *(*int32)(unsafe.Pointer(res_ptr))
	fmt.Println("Test success!", res)
	go_test_cpp.Swig_free(res_ptr)

	time.Sleep(time.Second * 1)
	fmt.Println("end")
}
