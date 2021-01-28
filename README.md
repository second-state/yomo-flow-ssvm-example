![ssvm-yomo](https://raw.githubusercontent.com/yomorun/yomo-flow-ssvm-example/main/yomo-ssvm-2.png)

# yomo-flow-ssvm-example

This examples represents how to write a [yomo-flow](https://yomo.run/flow) with [ssvm](https://www.secondstate.io/)

## Prerequisites

Please follow [ssvm](https://github.com/second-state/ssvm) to build `ssvm` first.
Then link ssvm binary to system $PATH.

## Compile wasm file

```bash
# Install ssvmup tool
curl https://raw.githubusercontent.com/second-state/ssvmup/master/installer/init.sh -sSf | sh

cd triple
ssvmup build
```

will get the wasm file at `triple/pkg/` directory

```bash
Ξ tmp/triple git:(master) ▶ ssvmup build
[INFO]: 🎯  Checking for the Wasm target...
[INFO]: 🌀  Compiling to Wasm...
   Compiling a22 v0.1.0 (/Users/fanweixiao/tmp/triple)
    Finished release [optimized] target(s) in 0.29s
⚠️   [WARN]: origin crate has no README
[INFO]: ⬇️  Installing wasm-bindgen...
[INFO]: Optimizing wasm binaries with `wasm-opt`...
[INFO]: Optional fields missing from Cargo.toml: 'description', 'repository', and 'license'. These are not necessary, but recommended
[INFO]: ✨   Done in 1.28s
[INFO]: 📦   Your wasm pkg is ready to publish at /Users/fanweixiao/tmp/triple/pkg.
Ξ tmp/triple git:(master) ▶
```

## Run this flow

1. Run this `yomo-flow`

```bash
go run main.go ssvm_wrapper.go
```

Then start your [yomo source](https://yomo.run/source) and [yomo zipper](https://yomo.run/zipper)

2. Run `yomo-zipper`

```bash
yomo wf run workflow.yaml
```

3. Run `yomo-source-example`

```bash
gh repo clone yomorun/yomo-source-example
cd yomo-source-example
YOMO_ZIPPER_ENDPOINT=localhost:9999 go run main.go
```

you'll get: 

```bash
Ξ _wrk/yomo-sink-ssvm → go run main.go
2021/01/26 00:14:20 Starting sink server...
2021/01/26 00:14:20 ✅ Listening on 0.0.0.0:4141
+6.913980e+015
+9.036794e+015
^Csignal: interrupt
```

## Notes

`triple/src/main.rs` will be built to pkg/triple.wasm which will be passed to `ssvm` command.

`ssvm_wrapper.go` wraps the command call for you.
You need to pass
1. `wasm` file path
2. `env` for command eg. LD_LIBRARY_PATH
3. `ssvm` [options](https://github.com/second-state/ssvm#run-ssvm-ssvm-with-general-wasm-runtime)
   * --reactor
   * --dir
   * --env
4. `args` for wasm

`ssvm_main.go` is a sample. Besides triple, it contains a tensorflow demo that require you to
follow [ssvm-tensorflow](https://github.com/second-state/ssvm-tensorflow) to build `ssvm-tensorflow` first.
Don't forget to link ssvm-tensorflow binary to system $PATH and download the dependent library.
