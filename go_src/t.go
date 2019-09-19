package main

import "fmt"

/*
// Note the -std=gnu99. Using -std=c99 will not work.
#cgo CFLAGS: -std=gnu99
#include <stdint.h>

uint32_t add (uint32_t a,uint32_t b){
	return a+b;
}

void cMultiply(int len, uint32_t *input, uint32_t *output) {
    for (int i = 0; i < len; i++) {
	        output[i] = input[i] * 2;
			    }
				}
*/
import "C"

func multiply(input []uint32) []uint32 {
	output := make([]uint32, len(input))
	C.cMultiply(C.int(len(input)),
		(*C.uint32_t)(&input[0]),
		(*C.uint32_t)(&output[0]))
	return output
}

func main() {

	a := C.add(1, 2)
	fmt.Println("a:", a)

	list := []uint32{23, 42, 17}
	list = multiply(list)
	for idx, val := range list {
		fmt.Printf("index %d: value %d\n", idx, val)
	}
}

/*
func main() {
	ch := make(chan string)

	go func() {
		for m := range ch {
			fmt.Println("Processed:", m)
			time.Sleep(10 * time.Second) // 模拟需要长时间运行的操作
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2" // 不会被接收处理
}

func main() {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- idx
		}(i)
	}

	fmt.Println(<-ch)           // 输出第一个发送的值
	close(ch)                   // 不能关闭，还有其他的 sender
	time.Sleep(2 * time.Second) // 模拟做其他的操作
}

func main() {
	done := false

	go func() {
		done = true
	}()

	for !done {
		println("not done !") // 并不内联执行
	}

	println("done !")
}
*/
/*
func main() {
	var names = []string{"aa", "bb", "cc"}
	for _, name := range names {
		go func() {
			var name2 = name
			println(name2)
		}()
	}
	runtime.Gosched()
}
*/
/*
func main() {
	a := ""
	b := "10:42:23"
	c := "10:41:24"
	d := "10:41:35"

	if a[:5] == b[:5] {
		println(a[:4])
		println("a==b")
	} else {
		println("a!=b")
	}
	if a[:5] == c[:5] {
		println("a==c")
	} else {
		println("a!=c")
	}

	if a[:5] == d[:5] {
		println("a==d")
	} else {
		println("a!=d")
	}
}
*/

/*
func test(i *int) int {
	return *i
}

func main() {
	var i = 1

	// defer定义的时候test(&i)的值就已经定了，是1，后面就不会变了
	defer fmt.Println("i1 =", test(&i))
	i++

	// defer定义的时候test(&i)的值就已经定了，是2，后面就不会变了
	defer fmt.Println("i2 =", test(&i))

	// defer定义的时候，i就已经确定了是一个指针类型，地址上的值变了，这里跟着变
	defer func(i *int) {
		fmt.Println("i3 =", *i)
	}(&i)

	// defer定义的时候i的值就已经定了，是2，后面就不会变了
	defer func(i int) {
		//defer 在定义的时候就定了
		fmt.Println("i4 =", i)
	}(i)

	defer func() {
		// 地址，所以后续跟着变
		var c = &i
		fmt.Println("i5 =", *c)
	}()

	// 执行了 i=11 后才调用，此时i值已是11
	defer func() {
		fmt.Println("i6 =", i)
	}()

	i = 11
}
*/
