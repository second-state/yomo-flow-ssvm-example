![ssvm-yomo](https://github.com/yomorun/yomo-flow-ssvm-example/blob/main/yomo-ssvm-2.jpg?raw=true)

# yomo-flow-ssvm-example

This examples represents how to write a [yomo-flow](https://yomo.run/flow) with [ssvm](https://www.secondstate.io/)

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
go run main.go
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
