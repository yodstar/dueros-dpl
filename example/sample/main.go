package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/yodstar/goutil/logger"

	"sample/config"
	"sample/startup"
)

var (
	CONF = config.CONF
	LOG  = logger.LOG
)

func main() {
	logpath := flag.String("o", "./sample.out", "Stdlog file path")
	cfgpath := flag.String("c", "./sample.conf", "Config file path")
	flag.Parse()

	//  Change working directory to the path where the program is running.
	if exefile, err := os.Executable(); err != nil {
		panic(err.Error())
	} else if err = os.Chdir(filepath.Dir(exefile)); err != nil {
		panic(err.Error())
	}

	// Redirect log output to file.
	logfile, err := os.OpenFile(*logpath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	log.SetOutput(logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Load config file
	if err = config.LoadFile(*cfgpath); err != nil {
		panic(err.Error())
	}

	// logger
	LOG.SetLevel(CONF.Logger.Level)
	LOG.SetOutFile(CONF.Logger.Outfile, "200601")
	LOG.SetFilter(CONF.Logger.Filter, func(s string) { log.Output(4, s) })
	// - LOG.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	startup.Run()
}
