package main

// #cgo LDFLAGS: -lrustpythoncbinding
// #cgo LDFLAGS: -L/Users/mmhh/study/RustPythonMini/target/debug
// void test_main(void);
import "C"

import (
	"fmt"
	"time"
)

func test(x int) {
	for {
		st := time.Now().UnixNano()
		C.test_main()
		fmt.Println(x, time.Now().UnixNano() - st)
	}
}

func main() {
	for x:=0; x<4; x++ {
		go test(x)
	}
	time.Sleep(time.Second*1000)

}