package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	y3 "github.com/yomorun/y3-codec-golang"
	"github.com/yomorun/yomo/pkg/quic"
)

var (
	addr = "0.0.0.0:4141"
)

func main() {
	go serveSinkServer(addr)
	select {}
}

// serveSinkServer serves the Sink server over QUIC.
func serveSinkServer(addr string) {
	log.Print("Starting sink server...")
	quicServer := quic.NewServer(&quicServerHandler{})
	err := quicServer.ListenAndServe(context.Background(), addr)
	if err != nil {
		log.Printf("❌ Serve the yomo-ssvm-example on %s failure with err: %v", addr, err)
	}
}

type quicServerHandler struct {
}

func (s *quicServerHandler) Listen() error {
	return nil
}

func (s *quicServerHandler) Read(st quic.Stream) error {
	// decode the data via Y3 Codec.
	ch := y3.
		FromStream(st).
		Subscribe(0x10).
		OnObserve(onObserve)

	go func() {
		for item := range ch {
			// Invoke wasm
			val := triple_ssvm(item.(noiseData).Noise)
			println(val)
		}
	}()

	return nil
}

type noiseData struct {
	Noise float64 `yomo:"0x11" fauna:"noise"` // Noise value
	Time  int64   `yomo:"0x12" fauna:"time"`  // Timestamp (ms)
	From  string  `yomo:"0x13" fauna:"from"`  // Source IP
}

func onObserve(v []byte) (interface{}, error) {
	// decode the data via Y3 Codec.
	data := noiseData{}
	err := y3.ToObject(v, &data)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return data, nil
}

func triple_ssvm(i float64) float64 {
	f := fmt.Sprintf("%f", i)
	result := Run(
		"/root/yomo-flow-ssvm-example/triple/pkg/triple_bg.wasm",
		map[string]string{},
		SSVMOptions{},
		[]string{f},
	)
	s, _ := strconv.ParseFloat(string(result), 64)
	return s
}

func triple_wasmer(i float64) float64 {
	// Reads the WebAssembly module as bytes.
	bytes, _ := wasm.ReadBytes("triple/pkg/triple_lib_bg.wasm")
	// Instantiates the WebAssembly module.
	instance, _ := wasm.NewInstance(bytes)
	defer instance.Close()
	// Gets the `sum` exported function from the WebAssembly instance.
	sum := instance.Exports["triple"]
	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, _ := sum(i)
	return result.ToF64()
}
