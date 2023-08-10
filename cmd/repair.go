package cmd

import (
	"flag"

	"github.com/ivanlebron/ledisdb/config"
	"github.com/ivanlebron/ledisdb/store"
)

func Repair() {
	var fileName = flag.String("config", "", "ledisdb config file")

	flag.Parse()

	if len(*fileName) == 0 {
		println("need ledis config file")
		return
	}

	cfg, err := config.NewConfigWithFile(*fileName)

	if err != nil {
		println(err.Error())
		return
	}

	if len(cfg.DataDir) == 0 {
		println("must set data dir")
		return
	}

	if err = store.Repair(cfg); err != nil {
		println("repair error: ", err.Error())
	}
}
