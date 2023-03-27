package main

import (
	"log"

	"github.com/Ralphbaer/hubla/common"
	"github.com/Ralphbaer/hubla/sales/gen"
)

func main() {
	common.InitLocalEnvConfig()
	gen.InitializeApp().Run()
	log.Print("sales service terminated")
}
