package logger

import (
	"fmt"
	"log/syslog"
	"os"
	"path"

	"github.com/op/go-logging"
)

var (
	l       *logging.Logger
	idLong  string
	idShort string
)

func init() {

	idShort, _ = os.Executable()
	idShort = path.Base(idShort)

	idLong = fmt.Sprintf("%s[%d]", idShort, os.Getpid())

	l = logging.MustGetLogger(idLong)
}

func Get() *logging.Logger { return l }

func Set(doSyslog bool, verbosity int) {

	var (
		err     error
		format  logging.Formatter
		backend logging.Backend
		sink    logging.Backend
	)

	verb := logging.Level(verbosity)

	if doSyslog {
		format = logging.MustStringFormatter(`%{message}`)

		backend, err = logging.NewSyslogBackendPriority(idShort, syslog.LOG_INFO|syslog.LOG_USER)
		if err != nil {
			panic(err)
		}
	} else {
		if verb == logging.DEBUG {
			format = logging.MustStringFormatter(
				`%{color}%{time:20060102150405}|%{program}.%{pid}|%{level:.4s}%{color:reset}|%{message}|%{shortfile}`,
			)
		} else {
			format = logging.MustStringFormatter(
				`%{color}%{time:20060102150405}|%{program}.%{pid}|%{level:.4s}%{color:reset}|%{message}`,
			)
		}
		backend = logging.NewLogBackend(os.Stderr, "", 0)
	}

	sink = logging.NewBackendFormatter(backend, format)
	logging.SetBackend(sink)

}
