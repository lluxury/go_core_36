package main

import (
	"errors"
	"fmt"
	"github.com/lluxury/go_core_36/37/common"
	"github.com/lluxury/go_core_36/37/common/op"
	"os"
	"runtime/pprof"
)

var (
	profileName = "cpuprofile.out"
)
func main() {
	f, err := common.CreateFile("", profileName)

	if err != nil {
		fmt.Printf("CPU profile creation error: %v\n,", err )
		return
	}
	defer f.Close()

	//if err := StartCPUProfile(f); err != nil {
	if err := startCPUProfile(f); err != nil {
		fmt.Printf("CPU profile start error: %v\n", err)
		return
	}

	if err = common.Execute(op.CPUProfile,10);err != nil {
		fmt.Printf("execute error: %v\n", err)
		return

	}
	//StopCPUProfile()
	stopCPUProfile()
}

func startCPUProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	return pprof.StartCPUProfile(f)
}

func stopCPUProfile() {
	pprof.StopCPUProfile()
}