将Python作为DSL语言使用。

本项目配合本人略微修改的RustPython分支使用，用于演示在Golang中调用编译为动态库的RustPython

修改版本的RustPython在本人Fork的RustPython代码库的mini-rust-python分支：
https://github.com/myrfy001/RustPython/tree/mini-rust-python



编译前请指定CGO的动态库链接路径，参见：
https://stackoverflow.com/questions/55282165/dylib-cannot-load-libstd-when-compiled-in-a-workspace

```
source "$HOME/.cargo/env"
export RUST_SRC_PATH="$(rustc --print sysroot)/lib/rustlib/src/rust/src"
export DYLD_LIBRARY_PATH="$(rustc --print sysroot)/lib:$DYLD_LIBRARY_PATH"
```