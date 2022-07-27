package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/fyzercmd/myGoServer/internal/apiserver"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	apiserver.NewApp("iam-apiserver").Run()
}
