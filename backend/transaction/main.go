package main

import (
	"log"

	"github.com/Ralphbaer/hubla/backend/common"
	"github.com/Ralphbaer/hubla/backend/transaction/gen"
)

func main() {
	common.InitLocalEnvConfig()
	gen.InitializeApp().Run()
	log.Print("transaction service terminated")
}
