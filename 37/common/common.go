package common

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type OpFunc func() error

func Execute(op OpFunc, times int)(err error){
	if op == nil {
		return errors.New("nil operation function")
	}

	if times <= 0 {
		return fmt.Errorf("invalid times: %d", times)
	}
	//var t1 time.Timer // ide 检查不出来
	var t1 time.Time
	defer func() {
		diff := time.Now().Sub(t1)
		fmt.Printf("elapased time: %s\n", diff)
		if p:= recover();p!= nil{
			err = fmt.Errorf("fatal error: %v",p)
		}
	}()
	t1 = time.Now()
	for i := 0; i < times; i++ {
	  if err  = op(); err != nil {
	  	return
	  }
	  time.Sleep(time.Microsecond)
	}
	return
}

func CreateFile(dir, name string) (*os.File, error) {
	if dir == "" {
		var err error
		dir, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	}
	path := filepath.Join(dir, name)
	return os.Create(path)
}