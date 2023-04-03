package main

import (
	"log"

	"github.com/Ralphbaer/hubla/common"
	"github.com/Ralphbaer/hubla/transaction/gen"
)

func main() {
	common.InitLocalEnvConfig()
	gen.InitializeApp().Run()
	log.Print("transaction service terminated")
}
