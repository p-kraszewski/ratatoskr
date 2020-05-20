package main

import (
	"os"

	flag "github.com/spf13/pflag"

	"github.com/p-kraszewski/ratatoskr/application"
	"github.com/p-kraszewski/ratatoskr/config"
	"github.com/p-kraszewski/ratatoskr/logger"
	"github.com/p-kraszewski/ratatoskr/mgmt"
)

var (
	confFile = flag.StringP("config", "c", "ratatosk.toml", "Configuration file")
	logLevel = flag.IntP("verbosity", "V", 2, "Log verbosity (0=critical to 5=debug)")
	doSysLog = flag.BoolP("syslog", "s", false, "Log to syslog")
	writeCfg = flag.BoolP("writecfg", "w", false, "(Re)write config)")
	log      = logger.Get()
)

func main() {
	flag.Lookup("verbosity").NoOptDefVal = "5"
	flag.Parse()

	conf, err := config.Load(*confFile)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Panicf("Configuration load failed: %s", err)
			os.Exit(1)
		}
	}

	if *writeCfg {
		err := conf.Save(*confFile)
		if err != nil {
			log.Errorf("Error writing config: %s", err)
		} else {
			log.Info("Config written")
		}
		return
	}

	m := &mgmt.Mgmt{}
	mb := &mgmt.MgmtB{}

	os.Exit(
		application.
			Run(conf, m, mb).
			Wait().
			Log(),
	)
}
