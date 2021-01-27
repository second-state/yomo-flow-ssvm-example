package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type SSVMOptions struct {
	reactor bool
	dirs map[string]string
	envs map[string]string
}

func Run(wasm string, cmdEnvs map[string]string, options SSVMOptions, args []string) []byte {
	cmdArgs := prepareCmdArgs(wasm, options)
	cmd := exec.Command("ssvm", cmdArgs...)
	cmd.Env = os.Environ()
	for ek, ev := range cmdEnvs {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s='%s'", ek, ev))
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		x, _ := json.Marshal(args)
		io.WriteString(stdin, string(x))
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	return out
}

func prepareCmdArgs(wasm string, options SSVMOptions) []string {
	cmdArgs := []string{wasm}
	if options.reactor {
		cmdArgs = append(cmdArgs, "--reactor")
	}
	for dk, dv := range options.dirs {
		cmdArgs = append(cmdArgs, "--dir", fmt.Sprintf("%s:'%s'", dk, dv))
	}
	for ek, ev := range options.envs {
		cmdArgs = append(cmdArgs, "--env", fmt.Sprintf("%s='%s'", ek, ev))
	}
	return cmdArgs
}

