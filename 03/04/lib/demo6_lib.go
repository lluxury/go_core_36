package lib

import (
	"os"
	in "github.com/lluxury/go_core_36/03/04/lib/internal"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}

//in 是别名,指再里面一个包