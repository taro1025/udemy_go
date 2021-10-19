package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	//(file, reading and writing, create, append, permission)
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file=logfile err=%s", err.Error())
	}

	//Stdout...near equal terminal
	multiLogFile := io.MultiWriter(os.Stdout, logfile)

	//set format
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//destination out put
	log.SetOutput(multiLogFile)
}
