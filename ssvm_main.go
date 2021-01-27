package main

import (
	"fmt"
	"strconv"
)

func main() {
	envs := map[string]string {
		"LD_LIBRARY_PATH": "/opt",
	}
	opts := SSVMOptions{
		reactor: true,
		dirs: map[string]string {
			"host_path": "guest_path",
		},
		envs: map[string]string {
			"APP_TOKEN": "sample case",
		},
	}
	result := Run(
		"/root/yomo-flow-ssvm-example/triple/pkg/triple_bg.wasm",
		envs,
		opts,
		[]string{"+6.913980e+015"})
	if s, err := strconv.ParseFloat(string(result), 64); err == nil {
		fmt.Println(s)
	}
}
