package main

import (
	"os"

	"github.com/unicell/fastexec"
)

func main() {
	fastexec.InitFlags()

	fastexec.StartWorkers()
	fastexec.InitJobPool(os.Stdin)

	fastexec.WaitToFinish()
}
