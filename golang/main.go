package main

// #cgo LDFLAGS: -lrustpythoncbinding
// #cgo LDFLAGS: -L/Users/mmhh/study/RustPythonMini/target/debug
// void test_main(void);
// void *rpy_compile_code(void *vm, const char *code);
// void *rpy_new_scope_with_builtins(void *vm);
// void rpy_run_code_obj(void *vm, void *code_obj, void *scope);
// void *rpy_vm_new(void);
import "C"

import (
	"fmt"
	"time"
)

func test() {

	vm := C.rpy_vm_new()
	scope := C.rpy_new_scope_with_builtins(vm)
	code_obj := C.rpy_compile_code(vm, C.CString(
		`for x in range(100):
	pass`,
	))

	t := time.Now().UnixNano()
	for x := 0; x < 10000; x++ {
		C.rpy_run_code_obj(vm, code_obj, scope)
	}

	fmt.Println(time.Now().UnixNano() - t)

}

func main() {
	test()

}
