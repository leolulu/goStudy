package main

import (
	"log"
	"myWeatherAnalysis/util"
	"myWeatherAnalysis/webserver"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		util.LogseqBaseDirPath = os.Args[1]
	}

	r := webserver.InitRouter()
	webserver.AddRouterForLogseqJournal(r)
	err := r.Run(":8082")
	if err != nil {
		log.Panicln(err)
	}
}
