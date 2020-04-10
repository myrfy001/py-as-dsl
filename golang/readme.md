https://stackoverflow.com/questions/55282165/dylib-cannot-load-libstd-when-compiled-in-a-workspace


source "$HOME/.cargo/env"
export RUST_SRC_PATH="$(rustc --print sysroot)/lib/rustlib/src/rust/src"
export DYLD_LIBRARY_PATH="$(rustc --print sysroot)/lib:$DYLD_LIBRARY_PATH"