package main

import (
	"log"

	"github.com/Ralphbaer/hubla/backend/auth/gen"
	"github.com/Ralphbaer/hubla/backend/common"
)

func main() {
	common.InitLocalEnvConfig()
	gen.InitializeApp().Run()
	log.Print("auth service terminated")
}
