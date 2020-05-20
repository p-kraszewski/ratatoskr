package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/p-kraszewski/ratatoskr/application"
	"github.com/p-kraszewski/ratatoskr/config"
	"github.com/p-kraszewski/ratatoskr/mgmt"
)

var (
	confFile = flag.StringP("config", "c", "ratatosk.toml", "Configuration file")
	logLevel = flag.IntP("verbosity", "V", 0, "Log verbosity")
)

func main() {
	flag.Lookup("verbosity").NoOptDefVal = "1"
	flag.Parse()

	conf, err := config.Load(*confFile)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}

	m := &mgmt.Mgmt{}

	fmt.Println(
		application.
			Run(conf, m).
			Wait(),
	)
}
