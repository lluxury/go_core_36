package main

import (
	"errors"
	"fmt"
	"github.com/lluxury/go_core_36/37/common"
	"github.com/lluxury/go_core_36/37/common/op"
	"os"
	"runtime"
	"runtime/pprof"
)

var (
	profileName      = "blockprofile.out"
	blockProfileRate =2
	debug 			=0
)

func main() {
	f, err := common.CreateFile("", profileName)
	if err != nil {
		fmt.Printf("block profile cretion error: %v\n", err)
		return
	}
	defer f.Close()
	startBlockProfile()
	if err = common.Execute(op.BlockProfile, 10); err != nil{
		fmt.Printf("execute error: %v\n", err)
		return
	}
	if err := stopBlockProfile(f); err != nil {
		fmt.Printf("block profile stop error: %v\n", err)
		return
	}
}

func startBlockProfile()  {
	runtime.SetBlockProfileRate(blockProfileRate)
}

func stopBlockProfile(f *os.File) error {
	if f == nil{
		return errors.New("nil file")
	}
	return pprof.Lookup("block").WriteTo(f, debug)
}